import Vue from 'vue';
import VueI18n from 'vue-i18n';
import en from './en.json';

//TODO: locale mutations with lazy loading https://github.com/kuanhsuh/vue-i18n-sandbox

Vue.use(VueI18n);

const locale = 'en';

const messages = {
  en
};

export default new VueI18n({
    locale,
    messages,
});
