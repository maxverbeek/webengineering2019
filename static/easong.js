var delCount = 5;

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
    song: {
      artist_id: null,
      release_id: null,
      id: null,
      title: null,
      year: null,
      duration: null,
      artist_mbtags: null,
      artist_mbtags_count: null,
      bars_confidence: null,
      bars_start: null,
      beats_confidence: null,
      beats_start: null,
      end_of_fade_in: null,
      hotttnesss: null,
      key: null,
      loudness: null,
      mode: null,
      mode_confidence: null,
      start_of_fade_out: null,
      tatums_confidence: null,
      tatums_start: null,
      tempo: null,
      time_signature: null,
      time_signature_confidence: null
    },
    pholders: {
      artist_id: null,
      release_id: null,
      id: null,
      title: null,
      year: null,
      duration: null,
      artist_mbtags: null,
      artist_mbtags_count: null,
      bars_confidence: null,
      bars_start: null,
      beats_confidence: null,
      beats_start: null,
      end_of_fade_in: null,
      hotttnesss: null,
      key: null,
      loudness: null,
      mode: null,
      mode_confidence: null,
      start_of_fade_out: null,
      tatums_confidence: null,
      tatums_start: null,
      tempo: null,
      time_signature: null,
      time_signature_confidence: null
    },
    link: null,
    code: 200,
    message: ''
  }
});

var bar = new Vue({
  el: '#bar',
  data: {
    edit: false
  }
});

var footer = new Vue({
  el: '#footer',
  data: {
    edit: false
  },
  methods: {
    addSong: function(){
      axios({
        method: 'post',
        url: '/api/v1/songs',
        data: song.song,
        headers: {
          'Content-Type': 'application/json'
        }
      }).then( response => {
        alert("Song added!");
      }).catch( error => {
        alert(error.response.status + ": " + error.response.data.message);
      })
    },
    editSong: function(){
      var dat = {};
      for( k in song.song ){
        console.log(k);
        if(song.song[k] != null){
          switch(k){
            case 'artist_id':
            case 'id':
            case 'title':
              dat[k] = song.song[k];
              break;
            case 'release_id':
            case 'year':
            case 'mode':
              dat[k] = parseInt(song.song[k]);
              break;
            default:
              dat[k] = parseFloat(song.song[k]);
          }
        }
      }
      axios({
        method: 'put',
        url: song.link,
        data: dat,
        headers: {
          'Content-Type': 'application/json'
        }
      }).then( response => {
        alert("Song updated!");
      }).catch( error => {
        alert(error.response.status + ": " + error.response.data.message);
      })
    }
  }
});
(async () => {
  if(getParameterByName('link') != null){
    footer.edit = bar.edit = true;
    try{
      response = await axios.get(getParameterByName('link'));
      song.pholders = response.data.data;
      song.link = getParameterByName('link');
      song.code = response.status;
    } catch(e) {
      song.code = e.response.status;
      song.message = e.response.data.message;
      footer.show = false;
    }
  }
})();
