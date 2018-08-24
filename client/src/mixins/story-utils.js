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
    }
  }
}
