import Vue from 'vue'
import Vuex from 'vuex'
// import VuexI18n from 'vuex-i18n'
import VuexPersistence from 'vuex-persist'

import nomination from './store/nomination'
import participant from './store/participant'

Vue.use(Vuex)

const vuexSessionStorage = new VuexPersistence({
  storage: window.sessionStorage
})

export default new Vuex.Store({
  strict: true,
  getters: {},
  modules: {
    nomination,
    participant
  },
  state: {},
  mutations: {},
  plugins: [vuexSessionStorage.plugin]
})