const config = require('config')

module.exports = {
  testing: config.get('db'),
  development: config.get('db'),
  production: config.get('db')
}
