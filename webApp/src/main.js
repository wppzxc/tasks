import Vue from 'vue'
import VueRouter from 'vue-router'
import Vant from 'vant'
import 'vant/lib/index.css'
import App from './App.vue'
import routes from './js/router/router'
import store from './js/store/store'
import VueClipboard from 'vue-clipboard2'
import axios from 'axios'

Vue.use(Vant);
Vue.use(VueRouter);
Vue.use(VueClipboard);
Vue.prototype.$axios = axios;

const router = new VueRouter({routes});

new Vue({
    router,
    store,
    el: '#app',
    render: h => h(App),
});
