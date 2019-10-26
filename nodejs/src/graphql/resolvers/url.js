const db = require('../../db')

module.exports = {
  Mutation: {
    async shorten (_, { url }) {
      if (!url.match(/^http/)) {
        throw new Error('invalid url')
      }
      const shortId = new Date().getTime().toString(36) + ('0000' + (Math.random() * Math.pow(36, 4) << 0).toString(36)).slice(-4)

      const find = await db.shortened_urls.findOne({
        where: {
          url
        }
      })

      if (find) {
        return find
      }

      await db.shortened_urls.create({
        url,
        shortId
      })

      return db.shortened_urls.findOne({
        where: {
          shortId
        }
      })
    }
  }
}
