function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, '\\$&');
    var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, ' '));
}

var artist = new Vue({
  el: '#artist',
  data: {
    songs: '0',
    artist: '',
    links: '',
    stats: {}
  }
});
var footer = new Vue({
  el: '#footer',
  methods: {
    getArtist: function(){
      return(artist.artist.id);
    }
  }
});
(async () => {
  if(getParameterByName('link') != null){
    response = await axios.get(getParameterByName('link'));
    artist.artist = response.data.data;
    [songres, statsres] = await Promise.all([
      axios.get(response.data.links.songs),
      axios.get(response.data.links.stats)
    ])

    artist.songs = songres.data.data
    artist.stats = statsres.data.data
  }
})();
