import { createApp } from 'vue';
import ClipboardJS from 'clipboard';
import App from './components/App.vue';
import 'bulma/css/bulma.css';

const app = createApp(App);
app.mount('#app');
app.component('App', App);

new ClipboardJS('.button');
