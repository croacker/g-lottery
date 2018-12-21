import BaseDataMapper from './base-data-mapper'
import Axios from 'axios'
import {ui} from '../service/ui-utils'
import config from '../config/prod-config'
// import store from 'vuex-store'

const networkAddress = config.NETWORKS.network.url

export default class RemoteDataMapper extends BaseDataMapper {
  get url() {
    throw new Error('Not implemented RemoteDataMapper.url')
  }

  parseLookupValue(lookupValue) {
    if (!ui.stringIsNullOrEmpty(lookupValue) && lookupValue.indexOf('#') > 0) {
      return decodeURIComponent(lookupValue.split('#')[1])
    }

    return lookupValue
  }

  createLookupValue(ref, value) {
    if (null != value) {
      return `${ref}#${value}`
    }
    return null
  }

  async getItems(filter) {
    const response = await this.sendRequest({
      url: this.url,
      method: 'get',
      data: filter
    })

    const data = response.data

    return this.convertCollection(data)
  }

  async getItem(filter) {
    const u = `${ui.stringTrimEnd(this.url, '/')}/${filter.id}`
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

  insert(payload) {
    return this.sendRequest({
      url: this.url,
      method: 'post',
      data: this.unmap(payload)
    })
  }

  afterUnmap(payload) {
    return payload
  }

  afterMap(payload) {
    return payload
  }

  update(payload) {
    const _url = `${ui.stringTrimEnd(this.url, '/')}/${payload.id}`
    const _data = this.afterUnmap(this.unmap(payload))
    return this.sendRequest({
      url: _url,
      method: 'put',
      data: _data
    })
  }

  convertCollection(data) {
    let result = []

    if (data && data.length) {
      for (const i in data) {
        const item = data[i]
        result.push(this.afterMap(this.map(item, i)))
      }
    }
    return result
  }

  delete(id) {
    const _url = `${ui.stringTrimEnd(this.url, '/')}/${id}`

    return this.sendRequest({
      url: _url,
      method: 'delete'
    })
  }

  sendRequest(props) {
    const url = props.url
    const method = props.method
    const formData = props.data

    let accessTokenId = ''
    // if (store.getters.accessToken) {
    //   accessTokenId = store.getters.accessToken.id
    // }

    if (!(url && url.length)) {
      throw new RangeError('Url null or empty')
    }

    let targetUrl = `${ui.stringTrimEnd(networkAddress, '/')}/${ui.stringTrimStart(url, '/')}`
    let bodyFormData = null

    if (method.toLowerCase() !== 'get') {
      bodyFormData = formData
    } else {
      targetUrl = targetUrl + '?'

      for (const p in formData) {
        targetUrl = targetUrl + p + '=' + encodeURIComponent(formData[p]) + '&'
      }
      targetUrl = ui.stringTrim(targetUrl, '&')
      targetUrl = ui.stringTrim(targetUrl, '?')
    }

    return Axios({
      method: method,
      url: targetUrl,
      data: bodyFormData,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Cache-Control': 'no-cache, no-store, must-revalidate',
        'X-Access-Token': accessTokenId,
      }
    })
  }
}
