
export default {
    MAX_LENGTH: 30,
    NOMINATION_TITLE: 'Н\xa0\xa0А\xa0\xa0Ч\xa0\xa0И\xa0\xa0Н\xa0\xa0А\xa0\xa0Е\xa0\xa0М\xa0\xa0!'.padStart(28, '\xa0').padEnd(30, '\xa0'),
    START_TITLE: '000000000000000000000000000000',
    // POSSIBLE : "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ0123456789",
    POSSIBLE: "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ0123456789",

    getRandomSymbol() {
        let symbolNumber = Math.random() * (this.POSSIBLE.length - 1)
        symbolNumber = Math.round(symbolNumber);
        let symbol = this.POSSIBLE[symbolNumber]

        let position = Math.random() * this.MAX_LENGTH
        position = Math.round(position);
        return {
            number: position,
            symbol: symbol
        }
    },

    formatDescription(description) {
        const pad = (this.MAX_LENGTH - description.length) / 2
        let leftPad = pad
        let rightPad = pad
        if ((pad ^ 0) !== pad) {
            leftPad = Math.round(pad)
            rightPad = this.MAX_LENGTH - description.length - leftPad
        }
        description = description.padStart(leftPad + description.length, '\xa0')
        description = description.padEnd(rightPad + description.length, '\xa0')
        description = description.toUpperCase()
        return description
    },
    
    getParticipantDescription(participant) {
        let surname = participant.Surname || ''
        let name = participant.Name || ''
        let department = participant.Department || ''

        let description = `${surname}\xa0${name}\xa0${department}`
        description = this.formatDescription(description)
        return description.slice('')
    }
}