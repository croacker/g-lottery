const state = {
    nominations: [],
    nomination: null
}


const mutations = {
    FETCH_NOMINATIONS(state, nominations) {
      state.nominations = nominations
    },
    FETCH_NOMINATION(state, nomination) {
      state.nomination = nomination
    }
  }
  
  const actions = {
    async fetchNominations({commit}, filter) {
      const mapper = new NominationDataMapper()
      const result = await mapper.getItems(filter)
      commit('FETCH_NOMINATIONS', result)
    },
    async fetchNomination({commit}, filter) {
      const mapper = new NominationDataMapper()
      const result = await mapper.getItem(filter)
      commit('FETCH_NOMINATION', result)
    }
  }
  
  const getters = {
    nominations(state) {
      return state.nominations
    },
    nomination(state) {
      return state.nomination
    }
  }
  
  export default {
    state,
    mutations,
    actions,
    getters
  }
  