var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!'
  }
})
var bar = new Vue({
  el: '#bar',
  data: {
    searchType: 'songs',
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
          + this.searchType
          + "?name=" + encodeURI(this.searchField)
          + "&genre=" + encodeURI(this.genreField)
          + "&year=" + encodeURI(this.yearField)
          + "&sort=" + encodeURI(this.sortBy)
          + "&limit=" + encodeURI(this.limit)
          + "&page=" + encodeURI(this.page)
        )
      .then(response => (app.message = response))
    }
  }
})
