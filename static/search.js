function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, '\\$&');
    var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, ' '));
}

var table = new Vue({
  el: '#table',
  data: {
    response: '',
    headers: [],
    rows: []
  }
})

var footer = new Vue({
  el: '#footer',
  data: {
    prev: null,
    limit: 50,
    page: 0,
    next: null
  },
  methods: {
    previousPage: function(){
      this.page--;
      bar.search();
    },
    currentPage: function(){
      bar.search();
    },
    nextPage: function(){
      this.page++;
      bar.search();
    },
    search: function(){
      bar.search();
    },
    getCsv: function(){
      bar.getCsv();
    }
  }
});

var bar = new Vue({
  el: '#bar',
  data: {
    searchType: 'Songs',
    yearField: '',
    genreField: '',
    artistFilter: '',
    sortBy: 'None',
    searchField: '',
  },
  methods: {
    search: function(){
      axios({
        method: 'get',
        url: '/api/v1/' + this.searchType.toLowerCase(),
        headers: {
          'Accept': 'application/json'
        },
        params: {
          name: this.searchField,
          genre: this.genreField.toLowerCase(),
          year: this.yearField,
          artist: this.artistFilter,
          sort: this.sortBy.toLowerCase(),
          limit: footer.limit,
          page: footer.page
        }
      }).then( response => {
          table.response = response;
          table.rows = [];
          if(this.searchType == 'Songs'){
            table.headers = [
              'Title',
              'ArtistId',
              'Duration',
              'Year'
            ]
            response.data.data.forEach(song => {
              table.rows.push([
                [song.title, "/song.html?link=" + encodeURI(song.links.self)],
                [song.artist_id, null],
                [song.duration, null],
                [song.year, null],
              ]);
            })
          } else {
            table.headers = [
              'Name',
              'ArtistId',
              'Genre'
            ]
            response.data.data.forEach(artist => {
              table.rows.push([
              [artist.name, "/artist.html?link=" + encodeURI(artist.links.self)],
              [artist.id, null],
              [artist.terms, null]
            ]);
          })
        }
        footer.prev = response.data.links.prev;
        footer.next = response.data.links.next;
      })
    },
    getCsv: function(){
      axios({
        method: 'get',
        url: '/api/v1/' + this.searchType.toLowerCase(),
        headers: {
          'Accept': 'text/csv'
        },
        params: {
          name: this.searchField,
          genre: this.genreField.toLowerCase(),
          year: this.yearField,
          artist: this.artistFilter,
          sort: this.sortBy.toLowerCase(),
          limit: footer.limit,
          page: footer.page
        }
      }).then( response => {
        console.log(response.data);
        var blob = new Blob([response.data]);
        var a = window.document.createElement("a");
        a.href = window.URL.createObjectURL(blob, {type: "text/plain"});
        a.download = this.searchType.toLowerCase() + ".csv";
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
      });
    }
  }
})

if(getParameterByName('type') != null){
  bar.searchType = getParameterByName('type');
}
if(getParameterByName('page') != null){
  footer.page = getParameterByName('page');
}
if(getParameterByName('search') != null ){
  bar.searchField = getParameterByName('search');
  bar.search();
}
if(getParameterByName('artist') != null ){
  bar.artistFilter = getParameterByName('artist');
  bar.search();
}
