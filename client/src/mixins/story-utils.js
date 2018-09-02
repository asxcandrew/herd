import dateFormat from 'dateformat';

export default {
  methods: {
    formatDate(date){
      return dateFormat(date, 'mmm d yyyy');
    },
    storyUri(story){
      return `${story.link}-${story.uid}`;
    },
    uidFromUri(uri){
      return uri.slice(-12);
    },
    countWords(story) {
      var matches = story.html_body.match(/[\w\d\'\`-]+/gi);
      return matches ? matches.length : 0;
    }
  }
}
