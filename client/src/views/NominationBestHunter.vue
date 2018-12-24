<template>
<div class="container">
  <div class="action">
    <div class="participant-name-container-inner">
      <div class="call-button-container participant-name-container">
        <span  @click="startPlayNomination" class="participant-name-text">{{caption}}</span>
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
  methods: {
    startPlayNomination() {
      const participants = this.$store.getters.participants
      this.caption = participants[0].description

      const component = this
      let i = 1;
      let participantsCount = participants.length
      let maxLength = Math.max.apply(Math, participants.map(function (el) { return el.description.length }));

      (function myLoop() {
        setTimeout(function () {
          var rand = Math.random() * participantsCount
          rand = Math.round(rand);
          let participant = participants[rand]
          
          rand = Math.random() * maxLength
          rand = Math.round(rand);

          let symb = participant.description[rand]
          let caption = component.caption.substr(0, rand - 1)
            + symb + component.caption.substr(rand + 1, component.caption.length)
          component.caption = caption            
          if (i < participants.length){
            i++
            myLoop(i);
          }else{
            i = 0
          } 
        }, 100)
      })(10);
    }
  }
}
</script>

<style scoped>
span {
  display: block;
}

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
  width: 700px;
  height: 190px;
}

.participant-name-container-inner {
  background-image: url('/img/best-hunter/square-background.png');
  background-size: cover;
  width: 700px;
  height: 190px;
}

.participant-name-text {
  font-size: 70px;
}

.buttons-color {
  opacity: 0.05;
}
</style>
