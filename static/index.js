var search = new Vue({
  el: '#search',
  data: {
    searchField: ''
  },
  methods: {
    searchSong: function() {
      if(search.searchField.toLowerCase() == "rickroll"){
        window.location.href = "https://www.youtube.com/watch?v=dQw4w9WgXcQ";
      }else{
        window.location.href = "/search.html?type=Songs&search=" + encodeURI(search.searchField);
      }
    },
    searchArtist: function(){
      window.location.href = "/search.html?type=Artists&search=" + encodeURI(search.searchField);
    }
  }
})
