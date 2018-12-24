<template>
<div class="action">
    <div class="participant-name-container">
        <img class="participant-name-background" src="/img/common/square-background.png" alt="">
        <div>
            <table class="participant-name-table">
                <tr>
                    <td v-for="s in columnsData" v-bind:id="columnsData" @click="startPlayNomination">
                        {{s}}
                </td>
                </tr>
            </table>
        </div>
    </div>
</div>
</template>

<script>
const TITLE = '        НАЧИНАЕМ!        '
const MAX_LENGTH = 25
const NOMINATION_CODE = 'transformation'
export default {
    beforeCreate: function () {
        document.body.className = 'transformation';
    },
    name: 'nomination-transformation',
    created: function () {
        const component = this
        this.$store.dispatch('fetchNominationsResults').then((result) => {
            let winnerResult = component.getExistsWinner()
            if (!winnerResult) {
                this.$store.dispatch('fetchNomination', NOMINATION_CODE).then((result) => {
                    const nomination = this.$store.getters.nomination
                    this.$store.dispatch('getByNomination', nomination).then((result) => {
                        const participants = this.$store.getters.participants
                        component.processParticipants(participants)
                    }).catch(error => {
                        console.log(error)
                    })
                }).catch(error => {
                    console.log(error)
                })
            } else {
                component.showWinner(winnerResult)
            }
        })
    },
    data() {
        return {
            nomination: null,
            columnsData: TITLE.split(''),
            participantDescriptions: [],
        }
    },
    methods: {
        startPlayNomination() {
            if (this.participantDescriptions.length != 0) {
                const animateTimerId = this.animateTitle()
                let nomination = this.$store.getters.nomination
                this.$store.dispatch('playNomination', nomination.ID)
                setTimeout(() => {
                    clearInterval(animateTimerId);
                    let winnerResult = this.getExistsWinner()
                    this.showWinner(winnerResult)
                }, 10000)
            }
        },
        processParticipants(participants) {
            const participantDescriptions = []
            let formatDescription = this.formatDescription
            participants.forEach(el => {
                let description = formatDescription(el.description)
                participantDescriptions.push(description.split(''))
            });
            this.participantDescriptions = participantDescriptions;
        },
        formatDescription(description) {
            const pad = (MAX_LENGTH - description.length) / 2
            let leftPad = pad
            let rightPad = pad
            if ((pad ^ 0) !== pad) {
                leftPad = Math.round(pad)
                rightPad = MAX_LENGTH - description.length - leftPad
            }
            description = description.padStart(leftPad + description.length)
            description = description.padEnd(rightPad + description.length)
            description = description.toUpperCase()
            return description
        },
        showWinner(winnerResult) {
            const participant = winnerResult.Participant
            let surname = participant.Surname || ''
            let name = participant.Name || ''

            let description = surname + '\xa0\xa0' + name
            description = this.formatDescription(description)

            this.columnsData = description.slice('')
        },
        getExistsWinner() {
            const existsResults = this.$store.getters.nominationsResults
            return existsResults.find(el => {
                return el.Nomination.Code === NOMINATION_CODE
            })
        },
        animateTitle() {
            const component = this
            let animateFn = function () {
                let randomSymbol = component.getRandomSymbol(component.participantDescriptions)
                let tmpArray = []
                component.columnsData.forEach(el => {
                    tmpArray.push(el)
                })
                tmpArray[randomSymbol.number] = randomSymbol.symbol
                component.columnsData = tmpArray
            }
            return setInterval(animateFn, 50)
        },

        getRandomSymbol(participantDescriptions) {
            let participantsCount = participantDescriptions.length

            let pNumber = Math.random() * (participantsCount - 1)
            pNumber = Math.round(pNumber);

            let sNumber = Math.random() * MAX_LENGTH
            sNumber = Math.round(sNumber);

            let participantDescription = participantDescriptions[pNumber]
            return {
                number: sNumber,
                symbol: participantDescription[sNumber]
            }
        }
    },
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
    height: 110vh;
}

.participant-name-container {
    border-image: url('/img/best-hunter/square-border.png') 30 round round;
    border-width: 30px;
    border-style: solid;
    background-size: cover;
    width: 1300px;
    height: 250px;
    position: absolute;
    top: 55%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.participant-name-container-inner {
    background-image: url('/img/common/square-background.png');
    background-size: cover;
    width: 1500px;
    height: 250px;
}

.participant-name-background {
    position: absolute;
    left: -30px;
    top: -20px;
    opacity: 0.5;
    width: 1300px;
    height: 230px;
}

.participant-name-text {
    font-size: 125px;
    font-family: 'BebasNeueRegular';
    font-weight: normal;
    font-style: normal;
    position: absolute;
    top: 50%;
    left: 50%;
    width: 1500px;
    transform: translate(-50%, -50%);
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

.buttons-color {
    opacity: 0.05;
}
</style>
