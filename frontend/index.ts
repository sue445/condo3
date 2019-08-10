import Vue from 'vue';
import App from './components/App.vue';

require('./scss/mystyles.scss');

let v = new Vue({
  el: "#app",
  components: { App },
  template: "<app></app>"
});
