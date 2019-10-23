function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, '\\$&');
    var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, ' '));
}

var song = new Vue({
  el: '#song',
  data: {
    song: '',
    artist: '',
    links: ''
  }
});

var bar = new Vue({
  el: '#bar',
  data: {
    title: ''
  }
});

(async () => {
  if(getParameterByName('link') != null){
    response = await axios.get(getParameterByName('link'));
    song.song = response.data.data;
    song.links = response.data.links;
    response = await axios.get(song.links.artist);
    song.artist = response.data.data;
    bar.title = song.song.title;
  }
})();
