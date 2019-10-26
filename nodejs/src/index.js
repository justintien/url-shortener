const config = require('config')
const app = require('./app.js')

const pkg = require('..//package.json')
const env = config.get('app.env', 'development')
const port = config.get('app.port', 4000)

app.listen(port)

console.log(`Server running at ${port}`)
console.log(`Running in ${env} v${pkg.version}`)
