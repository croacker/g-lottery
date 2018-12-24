import BaseItemDataMapper from './base-item-data-mapper'

export default class NominationResultDataMapper extends BaseItemDataMapper {
  insertResult = null

  buildInsertJobs(payload) {
    return []
  }

  get fieldMap() {
    return new Map([
      ['ID', 'ID'],
    //   ['CreatedAt', 'CreatedAt'],
    //   ['UpdatedAt', 'UpdatedAt'],
    //   ['DeletedAt', 'DeletedAt'],
    //   ['Name', 'Name'],
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

  async playNomination(nominationId){
    const u = `nominationresult/${nominationId}`
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

  async deleteNominationResult(nominationId){
    const u = `nominationresult/${nominationId}`
    const response = await this.sendRequest({
      url: u,
      method: 'delete',
    })

    const data = response.data
    let result = null
    if (data) {
      result = this.afterMap(this.map(data, 0))
    }
    return result
  }

  map(payload, i) {
    return payload
  }

  get url() {
    return '/nominationresultbynomination'
  }

  get urlItems() {
    return '/nominationresultbynomination'
  }
}
