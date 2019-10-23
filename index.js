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
    sortBy: 'None',
    searchField: '',
    limit: 50,
    page: 0
  },
  methods: {
    search: function(){
      axios
        .get('/api/v1/'
          + this.searchType.toLowerCase()
          + "?name=" + encodeURI(this.searchField)
          + "&genre=" + encodeURI(this.genreField.toLowerCase())
          + "&year=" + encodeURI(this.yearField)
          + "&sort=" + encodeURI(this.sortBy.toLowerCase())
          + "&limit=" + encodeURI(this.limit)
          + "&page=" + encodeURI(this.page)
        )
        .then(response => {
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
                data.title,
                data.ArtistId,
                data.duration,
                data.year
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
                data.ArtistName,
                data.ArtistId,
                data.ArtistTerms
              ]);
            })
          }
        })
    }
  }
})
