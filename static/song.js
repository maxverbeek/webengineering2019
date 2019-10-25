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
    song: '',
    artist: '',
    links: '',
    code: null,
    message: ''
  }
});

var bar = new Vue({
  el: '#bar',
  data: {
    title: ''
  }
});

var footer = new Vue({
  el: '#footer',
  data:{
    txt: 'Delete',
    link: '',
    show: true
  },
  methods: {
    del: function() {
      if(delCount-- <= 0){
        axios({
          method: 'delete',
          url: '/api/v1/songs/' + song.song.id
        }).then( response => {
          alert("deletion successfull");
        });
      }
      this.txt = this.txt.substr(0, 5-delCount) + this.txt.charAt(5-delCount).toUpperCase() + this.txt.substr(5-delCount+1)
    }
  }
});

(async () => {
  if(getParameterByName('link') != null){
    try{
      response = await axios.get(getParameterByName('link'));
      song.song = response.data.data;
      song.links = response.data.links;
      footer.link = song.links.self;
      response = await axios.get(song.links.artist);
      song.artist = response.data.data;
      bar.title = song.song.title;
      song.code = response.status;
    } catch(e) {
      song.code = e.response.status;
      song.message = e.response.data.message;
      footer.show = false;
    }
  }
})();
