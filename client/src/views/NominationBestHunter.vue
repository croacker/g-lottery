<template>
<div class="container">
  <div class="action">
    <div class="participant-name-container-inner">
      <div class="call-button-container participant-name-container">
        <span class="participant-name-text">{{caption}}</span>
      </div>
    </div>
  </div>
  <b-container>
    <b-button-group class="mt-2">
      <b-btn class="buttons-color" href="/#">Previous</b-btn>
      <b-btn href="/nomination-creative-class">Next</b-btn>
    </b-button-group>
  </b-container>
</div>
</template>

<script>
export default {
  beforeCreate: function () {
    document.body.className = 'best-hunter';
  },
  name: 'nomination-best-hunter',
  created: function () {
    this.$store.dispatch('fetchNomination', 'best-hunter').then((result) => {
      const nomination = this.$store.getters.nomination
      this.$store.dispatch('getByNomination', nomination).then((result) => {
        const participants = this.$store.getters.participants
      }).catch(error => {
        console.log(error)
      })
    }).catch(error => {
      console.log(error)
    })
  },
  data() {
    return {
      caption: 'Image thumbnails',
      nomination: null,
      items: []
    }
  },
  methods: {}
}
</script>

<style scoped>
/* .nomination-button-group {
  margin: 0 auto;
} */

.action {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.participant-name-container {
  background-image: url('/img/best-hunter/square-border.png');
  background-size: cover;
  padding: 35px;
}

.participant-name-container-inner {
  background-image: url('/img/best-hunter/square-background.png');
  background-size: cover;
}

.participant-name-text {
  font-size: 80px;
}

.buttons-color {
  opacity: 0.05;
}
</style>
