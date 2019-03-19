import Vue from 'vue'
// import Buefy from 'buefy';

const { Input } = require('buefy/dist/components/input');
const { Field } = require('buefy/dist/components/field');

import { library } from '@fortawesome/fontawesome-svg-core'
import { faSearch } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import App from './App.vue';
import router from './router';
import store from './store/index';

library.add(faSearch);

Vue.component('font-awesome-icon', FontAwesomeIcon);
Vue.component('b-field', Field);
Vue.component('b-input', Input)

Vue.config.productionTip = false

// Vue.use(Buefy, {
//     defaultIconPack: 'fas',
// });

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
