import RemoteDataMapper from './remote-data-mapper'

export default class ParticipantDataMapper extends RemoteDataMapper {
  insertResult = null

  buildInsertJobs(payload) {
    return []
  }

  get fieldMap() {
    return new Map([
      ['ID', 'ID'],
      ['CreatedAt', 'CreatedAt'],
      ['UpdatedAt', 'UpdatedAt'],
      ['DeletedAt', 'DeletedAt'],
      ['Surname', 'Surname'],
      ['Name', 'Name'],
      ['Chance', 'Chance'],
      ['Nomination', 'Nomination'],
      ['NominationID', 'NominationID'],
    ])
  }

  get url() {
    return '/participant'
  }
}
