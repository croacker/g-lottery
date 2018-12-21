import BaseItemDataMapper from './base-item-data-mapper'

export default class NominationDataMapper extends BaseItemDataMapper {
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
    return '/nomination'
  }
}
