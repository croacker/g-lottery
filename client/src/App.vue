<template>
<div id="app">
    <div id="nav">
        <div>
            <b-nav pills>
                <b-nav-item>
                    <router-link class="buttons-color" to="/">Домой</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-best-hunter">Best hunter</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-creative-class">Креативный класс</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-menthor">Наставник</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-certificate">Сертификаты</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-thanks">Спасибо</router-link>
                </b-nav-item>
                <b-nav-item>
                    <router-link class="buttons-color" to="/nomination-transformation">Трансформация</router-link>
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

body.nomination-list {
    background-image: url('/img/bg.png');
    background-size: cover;
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

.nomination-container {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 110vh;
}

.participant-name-container {
    border-image: url('/img/common/square-border.png') 30 round round;
    border-width: 30px;
    border-style: solid;
    background-size: cover;
    width: 1350px;
    height: 250px;
    position: absolute;
    top: 55%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.participant-name-background {
    position: absolute;
    left: -32px;
    top: -20px;
    opacity: 0.5;
    width: 1354px;
    height: 230px;
}

.participant-name-table {
    font-size: 90px;
    font-family: 'BebasNeueRegular';
    color: aliceblue;
    font-weight: normal;
    font-style: normal;
    border: none;
    position: absolute;
    top: 48%;
    left: 50%;
    height: 100%;
    width: 100%;
    transform: translate(-50%, -50%);
}

.participant-name-td {
    line-height: 0px;
    padding: 0px;
    border: 0px;
    width: 40px;
    text-transform: uppercase;
    font-weight: 550;
    text-shadow: 0px 0px 3px #fff, 0px 0px 5px #fff;
}

.participant-name-tr-animated {
    animation: blur 1.75s ease-out infinite;
    text-shadow: 0px 0px 5px #fff, 0px 0px 7px #fff;
}

@keyframes blur {
    from {
        text-shadow: 0px 0px 10px #fff,
            0px 0px 10px #fff,
            0px 0px 25px #fff,
            0px 0px 25px #fff,
            0px 0px 25px #fff,
            0px 0px 25px #fff,
            0px 0px 25px #fff,
            0px 0px 25px #fff,
            0px 0px 50px #fff,
            0px 0px 50px #fff,
            0px 0px 50px #7B96B8,
            0px 0px 150px #7B96B8,
            0px 10px 100px #7B96B8,
            0px 10px 100px #7B96B8,
            0px 10px 100px #7B96B8,
            0px 10px 100px #7B96B8,
            0px -10px 100px #7B96B8,
            0px -10px 100px #7B96B8;
    }
}
</style>
