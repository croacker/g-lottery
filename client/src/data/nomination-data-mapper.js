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

  async getItem(code) {
    const u = `nominationbycode/${code}`
    const response = await this.sendRequest({
      url: u,
      method: 'get',
    })

    const data = response.data
    let result = null
    if (data) {
      result = this.afterMap(this.map(data, 0))
    }
    return result
  }

  get url() {
    return '/nomination'
  }

  get urlItems() {
    return '/nomination'
  }
}
