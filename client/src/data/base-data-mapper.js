export default class BaseDataMapper {
    getItems(filter) {
      throw new Error('Not implemented BaseDataMapper.getItems')
    }
  
    getItem(filter) {
      throw new Error('Not implemented BaseDataMapper.getItem')
    }
  
    insert(payload) {
      throw new Error('Not implemented BaseDataMapper.insert')
    }
  
    update(payload) {
      throw new Error('Not implemented BaseDataMapper.update')
    }
  
    delete(filter) {
      throw new Error('Not implemented BaseDataMapper.delete')
    }
  
    exist(filter) {
      throw new Error('Not implemented BaseDataMapper.exist')
    }
  
    map(payload, i) {
      throw new Error('Not implemented BaseDataMapper.map')
    }
  
    unmap(payload) {
      throw new Error('Not implemented BaseDataMapper.unmap')
    }
  }
  