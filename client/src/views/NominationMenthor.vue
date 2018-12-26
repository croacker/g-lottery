<template>
<div class="action">
    <div class="participant-name-container">
        <img class="participant-name-background" src="/img/common/square-background.png" alt="">
        <div>
            <table class="participant-name-table">
                <tr class="participant-name-tr">
                    <td class="participant-name-td" v-for="s in columnsData" v-bind:id="columnsData" @click="startPlayNomination">
                        {{s}}
                </td>
                </tr>
            </table>
        </div>
    </div>
</div>
</template>

<script>
const TITLE = "Н\xa0\xa0А\xa0\xa0Ч\xa0\xa0И\xa0\xa0Н\xa0\xa0А\xa0\xa0Е\xa0\xa0М\xa0\xa0!".padStart(31, '\xa0').padEnd(35, '\xa0')
const MAX_LENGTH = 28
const NOMINATION_CODE = 'menthor'
const POSSIBLE = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ0123456789";
export default {
    beforeCreate: function () {
        document.body.className = 'menthor';
    },
    name: 'nomination-menthor',
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
    mounted() {
        const component = this
        window.addEventListener('keyup', function (event) {
            if (event.keyCode === 13) {
                component.startPlayNomination();
            }
        });
    },
    data() {
        return {
            nomination: null,
            columnsData: TITLE.split(''),
            participantDescriptions: [],
            winner: null,
        }
    },
    methods: {
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
            description = description.padStart(leftPad + description.length, '\xa0')
            description = description.padEnd(rightPad + description.length, '\xa0')
            description = description.toUpperCase()
            return description
        },
        showWinner(winnerResult) {
            const participant = winnerResult.Participant
            let surname = participant.Surname || ''
            let name = participant.Name || ''
            let department = participant.Department || ''

            let description = surname + '\xa0' + name + '\xa0' + department
            description = this.formatDescription(description)

            this.columnsData = description.slice('')
            this.winner = participant
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

            let symbolNumber = Math.random() * (POSSIBLE.length - 1)
            symbolNumber = Math.round(symbolNumber);
            let symbol = POSSIBLE[symbolNumber]

            let position = Math.random() * MAX_LENGTH
            position = Math.round(position);
            return {
                number: position,
                symbol: symbol
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

.participant-name-td {
    line-height: 0px;
    padding: 0px;
    border: 0px;
    width: 42px;
}

.buttons-color {
    opacity: 0.05;
}
</style>
