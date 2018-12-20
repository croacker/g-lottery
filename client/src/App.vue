<template>
  <div id="app">
    <div id="nav">
<div>      
<b-nav pills>
  <b-nav-item active><router-link to="/">Home</router-link></b-nav-item>
  <b-nav-item><router-link to="/about">About</router-link></b-nav-item>
  <b-nav-item><router-link to="/nominations">Номинации</router-link></b-nav-item>
</b-nav>
</div>
      
      
    </div>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
