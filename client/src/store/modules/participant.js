const state = {
    participants: [],
    participant: null
}

const mutations = {
    FETCH_PARTICIPANTS(state, participants) {
      state.participants = participants
    },
    FETCH_PARTICIPANT(state, participant) {
      state.participant = participant
    }
  }
  
  const actions = {
    async fetchParticipants({commit}, filter) {
      const mapper = new ParticipantDataMapper()
      const result = await mapper.getItems(filter)
      commit('FETCH_PARTICIPANTS', result)
    },
    async fetchParticipant({commit}, filter) {
      const mapper = new ParticipantDataMapper()
      const result = await mapper.getItem(filter)
      commit('FETCH_PARTICIPANT', result)
    }
  }
  
  const getters = {
    participants(state) {
      return state.participants
    },
    participant(state) {
      return state.participant
    }
  }
  
  export default {
    state,
    mutations,
    actions,
    getters
  }
  