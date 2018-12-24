<template>
<div id="app">
    <!-- <link rel="stylesheet" media="screen" href="https://fontlibrary.org/face/bebasneueregular" type="text/css"/> -->
    <div id="nav">
        <div>
            <b-nav pills>
                <b-nav-item>
                    <router-link class="buttons-color" to="/">Home</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-best-hunter">Best hunter</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-creative-class">Creative class</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-menthor">Menthor</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-certificate">Certificate</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-thanks">Thanks</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-transformation">Transformation</router-link>
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
body.home {
    background-image: url('/img/bg.png');
    background-size: cover;
    overflow: hidden;
}
body.best-hunter {
    background-image: url('/img/best-hunter/background.jpg');
    background-size: cover;
    overflow: hidden;
}
body.creative-class {
    background-image: url('/img/creative-class/background.jpg');
    background-size: cover;
    overflow: hidden;
}
body.menthor {
    background-image: url('/img/menthor/background.jpg');
    background-size: cover;
    overflow: hidden;
}
body.certificate {
    background-image: url('/img/certificate/background.jpg');
    background-size: cover;
    overflow: hidden;
}
body.thanks {
    background-image: url('/img/thanks/background.jpg');
    background-size: cover;
    overflow: hidden;
}
body.transformation {
    background-image: url('/img/transformation/background.jpg');
    background-size: cover;
    overflow: hidden;
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
</style>
