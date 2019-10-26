const fs = require('fs')
const path = require('path')
const config = require('config')
const Sequelize = require('sequelize')

const sequelize = new Sequelize(config.get('db'))

let db = {
  sequelize
}
fs
  .readdirSync(__dirname)
  .filter(file => {
    if (file.indexOf('.') === 0) {
      return false
    }
    if (file === path.basename(__filename)) {
      return false
    }
    if (path.extname(file) !== '.js') {
      return false
    }
    return true
  })
  .forEach(file => {
    let model = sequelize.import(path.join(__dirname, file))
    db[model.name] = model
  })

Object.keys(db).forEach(modelName => {
  if (db[modelName].associate) {
    db[modelName].associate(db)
  }
})

module.exports = db
