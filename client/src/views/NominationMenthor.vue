<template>
<div class="action">
    <div class="participant-name-container">
        <img class="participant-name-background" src="/img/common/square-background.png" alt="">
        <div>
            <table class="participant-name-table">
                <tr class="participant-name-tr">
                    <td class="participant-name-td" v-for="(symbol, idx) in columnsData" v-bind:id="'symbol-' + idx" v-bind:key="'key-' + idx" @click="startPlayNomination">
                        {{symbol}}
                    </td>
                </tr>
            </table>
        </div>
    </div>
</div>
</template>

<script>
import Vue from 'vue'
import Component from 'vue-class-component'
import common from '../service/common'

const TITLE = common.NOMINATION_TITLE
const NOMINATION_CODE = 'menthor'

@Component
class MenthorNomination extends Vue {
    beforeCreate() {
        document.body.className = NOMINATION_CODE;
    }
    name = 'nomination-menthor'
    created() {
        const component = this
        this.$store.dispatch('fetchNominationsResults').then((result) => {
            let winnerResult = component.getExistsWinner()
            if (!winnerResult) {
                this.$store.dispatch('fetchNomination', NOMINATION_CODE).then((result) => {
                    const nomination = this.$store.getters.nomination
                    this.$store.dispatch('getByNomination', nomination)
                }).catch(error => {
                    console.log(error)
                })
            } else {
                component.showWinner(winnerResult)
            }
        })
    }
    mounted() {
        const component = this
        window.addEventListener('keyup', function (event) {
            if (event.keyCode === 13) {
                component.startPlayNomination();
            }
        });
    }
    data() {
        return {
            nomination: null,
            columnsData: TITLE.split(''),
            winner: null,
        }
    }
    startPlayNomination() {
        if (!this.winner) {
            const animateTimerId = this.animateTitle()
            let nomination = this.$store.getters.nomination
            this.$store.dispatch('playNomination', nomination.ID)
            setTimeout(() => {
                clearInterval(animateTimerId);
                let winnerResult = this.getExistsWinner()
                this.showWinner(winnerResult)
            }, 10000)
        }
    }
    showWinner(winnerResult) {
        const participant = winnerResult.Participant
        this.columnsData = common.getParticipantDescription(participant).slice('')
        this.winner = participant
    }
    getExistsWinner() {
        const existsResults = this.$store.getters.nominationsResults
        return existsResults.find(el => {
            return el.Nomination.Code === NOMINATION_CODE
        })
    }
    animateTitle() {
        const component = this
        component.columnsData = common.START_TITLE.split('')
        let animateFn = function () {
            let randomSymbol = common.getRandomSymbol(component.participantDescriptions)
            let tmpArray = []
            component.columnsData.forEach(el => {
                tmpArray.push(el)
            })
            tmpArray[randomSymbol.number] = randomSymbol.symbol
            component.columnsData = tmpArray
        }
        return setInterval(animateFn, 50)
    }
}

export default MenthorNomination
</script>

<style scoped>
</style>
