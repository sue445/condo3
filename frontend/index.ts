import { createApp } from 'vue';
import ClipboardJS from 'clipboard';
import App from './components/App.vue';
import './scss/mystyles.scss';

const app = createApp(App);
app.mount('#app');
app.component('App', App);

new ClipboardJS('.button');
