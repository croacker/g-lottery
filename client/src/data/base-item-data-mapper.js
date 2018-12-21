import RemoteDataMapper from './remote-data-mapper'
import {dataLayerUtils} from '../service/data-layer-utils'
import {ui} from '../service/ui-utils'

export default class BaseItemDataMapper extends RemoteDataMapper {
    insertResult = null
  
    async insert(payload) {
      if (ui.stringIsNullOrEmpty(payload.id)) {
        payload.id = dataLayerUtils.generateGuid()
      }
      this.insertResult = await super.insert(payload)
      return dataLayerUtils.all(this.buildInsertJobs(payload))
    }
  
    buildInsertJobs(payload) {
      return []
    }
  
    get fieldMap() {
      throw new Error('Not implemented RootItemDataMapper.fieldMap')
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
  }
  