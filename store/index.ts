import Vue from 'vue';
import Vuex from 'vuex';

import { search } from './search/index';

import { RootState } from './types';

Vue.use(Vuex)

export default new Vuex.Store<RootState>({
  state: {
      verison: '0.0.1',
  },
  modules: {
      search,
  },
  mutations: {

  },
  actions: {

  }
})
