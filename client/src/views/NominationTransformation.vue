<template>
<div class="container">
  <b-container>
    <b-alert show>Трансформация</b-alert>
<b-container fluid class="bv-example-row">
  <b-row>
    <b-col md="4">md="4"</b-col>
    <b-col md="4" offset-md="4">md="4" offset-md="4"</b-col>
  </b-row>
  <b-row>
    <b-col md="3" offset-md="3">md="3" offset-md="3"</b-col>
    <b-col md="3" offset-md="3">md="3" offset-md="3"</b-col>
  </b-row>
  <b-row>
    <b-col md="6" offset-md="3">md="6" offset-md="3"</b-col>
  </b-row>
  <b-row>
    <b-col md="6" offset-md="3">md="6" offset-md="3"</b-col>
  </b-row>
  <b-row>
    <b-col md="6" offset-md="3"><br></b-col>
  </b-row>
  <b-row>
    <b-col md="6" offset-md="3"><br><br><br><br><br></b-col>
  </b-row>
</b-container>

    <b-button-group class="mt-2">
      <b-btn href="/nomination-thanks">Previous</b-btn>
    </b-button-group>
  </b-container>
</div>
</template>

<script>
export default {
  name: 'nomination-transformation',
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
