<template>
<div class="container">
    <b-container>
        <b-card title="Номинации" class="text-left">
            <b-form-select v-model="selectedNomination" :options="nominations" @change="onSelectedNominationChange" class="mb-3" />
            <b-button-toolbar>
                <b-button-group class="nomination-button-group mx-1">
                    <b-form-file v-model="file" :state="Boolean(file)" @change="onFileSelect" placeholder="Загрузить участников..."></b-form-file>
                    <b-btn variant="success" @click="onFileUload">
                        <font-awesome-icon icon="file-upload" />
                    </b-btn>
                    <b-btn variant="danger" @click="onFileClear">
                        <font-awesome-icon icon="trash" />
                    </b-btn>
                </b-button-group>
                <b-button-group class="nomination-button-group mx-1">
                    <b-btn @click="onRemoveAllClick" variant="danger">Удалить</b-btn>
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
    created: function () {
        this.$store.dispatch('fetchNominations').then((result) => {
            this.updateNominations()
        }).catch(error => {
            console.log(error)
        })
    },
    data() {
        return {
            fields: ['id', 'surname', 'name', 'chance'],
            items: [{
                    id: 1,
                    chance: 40,
                    name: 'Dickerson',
                    surname: 'Macdonald'
                },
                {
                    id: 2,
                    chance: 21,
                    name: 'Larsen',
                    surname: 'Shaw'
                }
            ],
            file: null,
            selectedNomination: null,
            nominations: [{
                value: null,
                text: 'Номинация не указана'
            }]
        }
    },
    methods: {
        onSelectedNominationChange: function () {
            console.log('changed');
        },
        onRemoveAllClick: function () {
            console.log("onRemoveAllClick")
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
                    value: nomination.ID,
                    text: nomination.Name
                })
            });
            this.nominations = dataNominations
        },
        updateParticipants(participants) {
            let participantsItems = []
            participants.forEach((participant, idx) => {
                participantsItems.push({
                    id: idx,
                    surname: participant.surname,
                    name: participant.name,
                    chance: participant.chance
                })
            });
            this.items = participantsItems
        }
    }

}
</script>

<style scoped>
.nomination-button-group {
    margin: 0 auto;
}
</style>
