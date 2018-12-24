<template>
  <div class="home">  
    <Welcome msg="Перед началом работы, загрузите участников!"/>
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
