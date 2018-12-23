import ParticipantDataMapper from '../data/participant-data-mapper'

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
    },
    async insertParticipant({commit}, filter) {
      const mapper = new ParticipantDataMapper()
      await mapper.insert(filter)
    },
    async insertParticipants({commit}, items) {
      const mapper = new ParticipantDataMapper()
      await mapper.insertItems(items)
    },
    async deleteParticipant({commit}, filter) {
      const mapper = new ParticipantDataMapper()
      await mapper.delete(filter)
    },
    async getByNomination({commit}, filter) {
      const mapper = new ParticipantDataMapper()
      const result = await mapper.byNomination(filter)
      commit('FETCH_PARTICIPANTS', result)
    },
    async deleteByNomination({commit}, filter) {
      const mapper = new ParticipantDataMapper()
      await mapper.deleteByNomination(filter)
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
  