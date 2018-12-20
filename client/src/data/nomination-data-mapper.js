import RemoteDataMapper from './remote-data-mapper'

export default class NominationDataMapper extends RemoteDataMapper {
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
      ['Name', 'Name'],
    ])
  }

  get url() {
    return '/participants'
  }
}
