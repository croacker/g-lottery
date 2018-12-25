import RemoteDataMapper from './remote-data-mapper'

export default class ParticipantDataMapper extends RemoteDataMapper {
  insertResult = null

  buildInsertJobs(payload) {
    return []
  }

  get fieldMap() {
    return new Map([
      ['ID', 'ID'],
      ['surname', 'Surname'],
      ['name', 'Name'],
      ['department', 'Department'],
      ['chance', 'Chance'],
      ['nominationID', 'NominationID'],
    ])
  }

  async byNomination(nomination){
    const response = await this.sendRequest({
      url: this.urlItems + '/' + nomination.ID,
      method: 'get',
      data: ''
    })

    const data = response.data

    return this.convertCollection(data)
  }

  deleteByNomination(nomination){
    const _url = this.urlItems + '/' + nomination.ID
    return this.sendRequest({
      url: _url,
      method: 'delete'
    })
  }

  insertItems(payloads) {
    let items = []

    payloads.forEach(element => {
      items.push(this.unmap(element))  
    });

    return this.sendRequest({
      url: this.url,
      method: 'post',
      data: items
    })
  }

  map(payload, i) {
    let item = {}
    for (const [internalKey, externalKey] of this.fieldMap.entries()) {
      let value = null
      if (null != externalKey && typeof externalKey === 'function') {
        value = externalKey(payload, internalKey)
      } else {
        value = payload[externalKey]
      }

      item[internalKey] = value
    }
    item.surname = item.surname || '' 
    item.name = item.name || '' 
    item.description = item.surname + ' ' + item.name
    return item
  }

  unmap(payload) {
    let item = {}
    for (const [internalKey, externalKey] of this.fieldMap.entries()) {
      if (typeof externalKey !== 'function') {
        const value = payload[internalKey]
        item[externalKey] = value
      }
    }
    return item
  }

  get url() {
    return '/participant'
  }

  get urlItems() {
    return '/participantbynomination'
  }
}
