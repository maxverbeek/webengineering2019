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
var bar = new Vue({
  el: '#bar',
  data: {
    searchType: 'Songs',
    yearField: '',
    genreField: '',
    artistFilter: '',
    sortBy: 'None',
    searchField: '',
    limit: 50,
    page: 0
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
          name: encodeURI(this.searchField),
          genre: encodeURI(this.genreField.toLowerCase()),
          year: encodeURI(this.yearField),
          artist: encodeURI(this.artistFilter),
          sort: encodeURI(this.sortBy.toLowerCase()),
          limit: encodeURI(this.limit),
          page: encodeURI(this.page),
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
            response.data.data.forEach(data => {
              table.rows.push([
                [data.title, "/song.html?link=" + encodeURI(data.links.self)],
                [data.ArtistId, null],
                [data.duration, null],
                [data.year, null],
              ]);
            })
          } else {
            table.headers = [
              'Name',
              'ArtistId',
              'Genre'
            ]
            response.data.data.forEach(data => {
              table.rows.push([
              [data.ArtistName, "/artist.html?link=" + encodeURI(data.links.self)],
              [data.ArtistId, null],
              [data.ArtistTerms, null]
            ]);
          })
        }
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
          name: encodeURI(this.searchField),
          genre: encodeURI(this.genreField.toLowerCase()),
          year: encodeURI(this.yearField),
          artist: encodeURI(this.artistFilter),
          sort: encodeURI(this.sortBy.toLowerCase()),
          limit: encodeURI(this.limit),
          page: encodeURI(this.page),
        }
      }).then( response => {
        var newWindow = window.open();
        newWindow.document.write(response.data);
	console.log(response);
      });
    }
  }
})

if(getParameterByName('type') != null){
  bar.searchType = getParameterByName('type');
}
if(getParameterByName('search') != null ){
  bar.searchField = getParameterByName('search');
  bar.search();
}
if(getParameterByName('artist') != null ){
  bar.artistFilter = getParameterByName('artist');
  bar.search();
}
