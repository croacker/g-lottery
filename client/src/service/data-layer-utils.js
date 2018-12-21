import Axios from 'axios'

const dataLayerUtils = {
  all: function (jobs) {
    return Axios.all(jobs)
  },
  getMapKeyByValue(map, value) {
    let op = [...map.entries()].find(entry => entry[1] === value)
    if (op == null || !op.length) {
      return null
    }

    return op[0]
  },
  getMapValueByKey(map, key) {
    if (map.has(key)) {
      return map.get(key)
    }

    return null
  },
  generateGuid: function () {
    function s4() {
      return Math.floor((1 + Math.random()) * 0x10000).toString(16).substring(1)
    }

    return s4() + s4() + s4() + s4() + s4() + s4() + s4() + s4()
  },
  transformArray(inspectionItems) {
    // console.log(inspectionItems)
    const uniqueEquipments = this.getColumnSet(inspectionItems, 'equipmentId')
    const uniqueParams = this.getColumnSet(inspectionItems, 'param')

    const result = []
    if (null != uniqueEquipments && uniqueEquipments.length) {
      for (const i in uniqueEquipments) {
        const uniqueEquipment = uniqueEquipments[i]

        const equipment = inspectionItems.find((element, index, array) => {
          return element.equipmentId === uniqueEquipment
        })

        const row = {
          equipmentId: equipment.equipmentId,
          equipmentType: equipment.equipmentType
        }
        for (const j in uniqueParams) {
          const uniqueParam = uniqueParams[j]
          const p = inspectionItems.find((element, index, array) => {
            return element.equipmentId === uniqueEquipment && element.param === uniqueParam
          })

          row[uniqueParam] = undefined !== p ? this.parseValue(p.value) : null
        }
        result.push(row)
      }
    }

    return result
  },
  parseValue(v) {
    if ('true' === v) {
      v = 'Да'
    }
    if ('false' === v) {
      v = 'Нет'
    }
    return v
  },
  getColumnSet(array, columnName) {
    const tmp = []
    for (const i in array) {
      const row = array[i]
      const val = row[columnName]
      if (null != val) {
        tmp.push(val)
      }
    }
    return [...new Set(tmp)]
  }
}

export {
  dataLayerUtils
}
