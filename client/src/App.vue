<template>
<div id="app">
    <div id="nav">
        <div>
            <b-nav pills>
                <b-nav-item>
                    <router-link class="buttons-color" to="/">Home</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/about">About</router-link>
                </b-nav-item>
                <b-navbar-nav class="ml-auto">
                    <b-nav-item right>
                        <router-link class="buttons-color" to="/nominations">
                            <font-awesome-icon icon="cog" />
                        </router-link>
                    </b-nav-item>
                </b-navbar-nav>
            </b-nav>
        </div>
    </div>
    <router-view />
</div>
</template>

<style>
body {
    background-image: url('/img/bg.png');
    background-size: cover;
}
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
}
.buttons-color {
    color: #142F3F;
}
.custom-file-input:lang(en)~.custom-file-label::after {
    content: "Выбрать";
}
.custom-file-input:lang(ru)~.custom-file-label::after {
    content: "Выбрать";
}
</style>
