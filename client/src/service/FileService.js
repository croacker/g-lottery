import XLSX from 'xlsx'

export default {

    async parseXls(file) {
        const data = await this.fileToBase64(file, 'BinaryString')
        const workbook = XLSX.read(data, { type: 'binary' })
        const sheet = this.getSheetByIndex(workbook, 0)
        let cells = XLSX.utils.sheet_to_json(sheet, {header: 1})
        cells = cells.filter(el => {return !!el.length != 0})
        let grouped = {}
        cells.forEach(el =>{
            let description = el[0]
            if(description in grouped){
                grouped[description]++
            }else{
                grouped[description] = 1
            }
        })
        return Object.entries(grouped).map(el => {
            let names = el[0].split(" "); 
            let chance = el[1]
            return {
                surname: names[0],
                name:names[1],
                department:names[2],
                chance: chance
            }
        })

    },

    getSheetByIndex(workbook, index) {
        if (workbook.SheetNames != null && workbook.SheetNames.length > index) {
            const firstSheetName = workbook.SheetNames[index]
            if (firstSheetName != null) {
                return workbook.Sheets[firstSheetName]
            }
        }
        throw new Error('Sheet not found')
    },

    getCellValue(sheet, cell, prop) {
        if (sheet == null || cell == null) {
            return null
        }
        if (prop == null) {
            prop = 'w'
        }
        return sheet[cell] != null ? sheet[cell][prop] : null
    },

    async fileToBase64(file, readAs) {
        return new Promise((resolve, reject) => {
          const reader = new FileReader()
          reader.onload = () => {
            resolve(reader.result)
          }
    
          reader.onerror = (e) => {
            reject(e)
          }
    
          if (!readAs) {
            readAs = 'DataURL'
          }
          if (readAs === 'DataURL') {
            reader.readAsDataURL(file)
          } else if (readAs === 'BinaryString') {
            reader.readAsBinaryString(file)
          } else {
            throw new Error('UnSupported')
          }
        })
      },
}