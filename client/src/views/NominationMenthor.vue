<template>
<div class="container">
  <b-container>
    <b-alert show>Наставник</b-alert>
    <b-button-group class="mt-2">
      <b-btn href="/nomination-creative-class">Previous</b-btn>
      <b-btn href="/nomination-certificate">Next</b-btn>
    </b-button-group>
  </b-container>
</div>
</template>

<script>
export default {
  name: 'nomination-menthor',
  created: function () {
    this.$store.dispatch('fetchNominations').then((result) => {
      this.updateNominations()
    }).catch(error => {
      console.log(error)
    })
  },
  data() {
    return {
      nomination: null,
      items: []
    }
  },
  methods: {
    updateParticipants(participants) {
      const nomiation = this.selectedNomination
      let participantsItems = []
      participants.forEach((participant, idx) => {
        participantsItems.push({
          id: idx,
          surname: participant.surname,
          name: participant.name,
          chance: participant.chance,
          nominationID: nomiation.ID
        })
      });
      this.items = participantsItems
    }
  }
}
</script>

<style scoped>
.nomination-button-group {
  margin: 0 auto;
}
</style>
