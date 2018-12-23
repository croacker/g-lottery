<template>
  <div class="home">  
    <Welcome msg="Welcome"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Welcome from '@/components/Welcome.vue'

export default {
  name: 'home',
  components: {
    Welcome
  },
  beforeCreate: function() {
            document.body.className = 'home';
  }
}
</script>
