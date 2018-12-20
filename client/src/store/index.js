import Vue from 'vue'
import Vuex from 'vuex'
import VuexI18n from 'vuex-i18n'

import nomination from './modules/nomination'
import participant from './modules/participant'

Vue.use(Vuex)

const vuexSessionStorage = new VuexPersistence({
  storage: window.sessionStorage
})

const store = new Vuex.Store({
  strict: true,
  getters,
  modules: {
    nomination,
    participant
  },
  state: {},
  mutations: {},
  plugins: [vuexSessionStorage.plugin]
})

Vue.use(VuexI18n.plugin, store)

export default store