import NominationResultDataMapper from '../data/nominationresult-data-mapper'

const state = {
    nominationsResults: []
}

const mutations = {
    FETCH_NOMINATIONS_RESULTS(state, nominationsResults) {
      state.nominationsResults = nominationsResults
    }
  }
  
  const actions = {
    async fetchNominationsResults({commit}, filter) {
      const mapper = new NominationResultDataMapper()
      const result = await mapper.getItems(filter)
      commit('FETCH_NOMINATIONS_RESULTS', result)
    },
    async playNomination({commit}, filter) {
      const mapper = new NominationResultDataMapper()
      await mapper.playNomination(filter)
      this.dispatch('fetchNominationsResults')
    },
    async deleteNominationResult({commit}, filter) {
      const mapper = new NominationResultDataMapper()
      await mapper.deleteNominationResult(filter)
      this.dispatch('fetchNominationsResults')
    }
  }
  
  const getters = {
    nominationsResults(state) {
      return state.nominationsResults
    }
  }
  
  export default {
    state,
    mutations,
    actions,
    getters
  }
  