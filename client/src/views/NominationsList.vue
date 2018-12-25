<template>
<div class="container">
    <b-container>
        <b-card title="Номинации" class="text-left">
            <b-form-select required v-model="selectedNomination" :options="nominations" @change="onSelectedNominationChange" class="mb-3" />
            <b-button-toolbar>
                <b-button-group class="nomination-button-group mx-1">
                    <b-form-file v-model="file" :state="Boolean(file)" @change="onFileSelect" placeholder="Загрузить участников..."></b-form-file>
                    <b-btn variant="success" @click="onFileUload">
                        <font-awesome-icon icon="file-upload" />
                    </b-btn>
                </b-button-group>
                <b-button-group class="nomination-button-group mx-1">
                    <b-btn @click="onLoadAllClick" variant="success">Загрузить</b-btn>
                </b-button-group>
                <b-button-group class="nomination-button-group mx-1">
                    <b-btn @click="onNominationResultDeleteClick" variant="danger">Отменить результат</b-btn>
                </b-button-group>

            </b-button-toolbar>

            <br>
            <b-table striped hover :items="items" :fields="fields"></b-table>
        </b-card>
    </b-container>
</div>
</template>

<script>
import FileService from '../service/FileService'
export default {
    name: 'nominations-list',
    beforeCreate: function () {
        document.body.className = 'nomination-list';
    },
    created: function () {
        this.$store.dispatch('fetchNominations').then((result) => {
            this.updateNominations()
        }).catch(error => {
            console.log(error)
        })
    },
    data() {
        return {
            fields: [{
                key: 'id',
                label: '',
            }, {
                key: 'surname',
                label: 'Фамилия'
            }, {
                key: 'name',
                label: 'Имя'
            }, {
            }, {
                key: 'department',
                label: 'Подразделение'
            }, {
                key: 'chance',
                label: 'Шанс'
            }],
            items: [],
            file: null,
            selectedNomination: null,
            nominations: [{
                value: null,
                text: 'Номинация не указана'
            }]
        }
    },
    methods: {
        onSelectedNominationChange: function (val) {
            this.clearParticipants()
            this.$store.dispatch('getByNomination', val).then((result) => {
                const participants = this.$store.getters.participants
                this.updateParticipants(participants)
            }).catch(error => {
                console.log(error)
            })
        },
        onLoadAllClick: function () {
            const nomiation = this.selectedNomination
            if (nomiation) {
                this.$store.dispatch('deleteByNomination', nomiation)
                this.$store.dispatch('insertParticipants', this.items)
            }
        },
        onFileSelect: function () {
            console.log("onFileSelect")
        },
        onFileUload: function () {
            try {
                FileService.parseXls(this.file).then((result) => {
                    this.updateParticipants(result)
                }).catch(error => {
                    console.log(error)
                })
            } catch (e) {
                console.log(e)
            }
        },
        onFileClear: function () {
            console.log("onFileClear")
        },
        updateNominations: function () {
            const nominations = this.$store.getters.nominations
            let dataNominations = []
            nominations.forEach(nomination => {
                dataNominations.push({
                    value: nomination,
                    text: nomination.Name
                })
            });
            this.nominations = dataNominations
            this.selectedNomination = this.nominations[0]
        },
        clearParticipants() {
            this.items = []
        },
        requestParticipants(nomination) {

        },
        updateParticipants(participants) {
            const nomiation = this.selectedNomination
            let participantsItems = []
            participants.forEach((participant, idx) => {
                participantsItems.push({
                    id: idx,
                    surname: participant.surname,
                    name: participant.name,
                    department: participant.department,
                    chance: participant.chance,
                    nominationID: nomiation.ID
                })
            });
            this.items = participantsItems
        },
        onNominationResultDeleteClick() {
            const nomination = this.selectedNomination
            if (nomination) {
                this.$store.dispatch('deleteNominationResult', nomination.ID)
            }
        }
    }

}
</script>

<style scoped>
.nomination-button-group {
    margin: 0 auto;
}
</style>
