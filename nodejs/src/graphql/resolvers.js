const path = require('path')
const fs = require('fs')
const { merge } = require('lodash')
const GraphQLDateTime = require('graphql-type-datetime')

let resolvers = {
  DateTime: GraphQLDateTime
}

const rootPath = path.join(__dirname, './resolvers')

fs.readdirSync(rootPath).forEach(resolver => {
  merge(resolvers, require(path.join(rootPath, resolver)))
})

module.exports = resolvers
