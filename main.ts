import Vue from 'vue'
import Buefy from 'buefy';

import { library} from '@fortawesome/fontawesome-svg-core'
import { faSearch } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import App from './App.vue'
import router from './router'
import store from './store'

library.add(faSearch);

Vue.component('font-awesome-icon', FontAwesomeIcon);

Vue.config.productionTip = false

Vue.use(Buefy, {
    defaultIconPack: 'fas',
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
