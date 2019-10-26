const Koa = require('koa')
const fs = require('fs')
const bodyParser = require('koa-bodyparser')

const { ApolloServer, gql } = require('apollo-server-koa')
const { makeExecutableSchema } = require('graphql-tools')

const typeDefs = gql(require('./graphql/typeDefs'))
const resolvers = require('./graphql/resolvers')

const db = require('./db')
/**
 * render view
 * @param  {string} page render view/${page}
 * @return {promise}
 */
function render (page) {
  return fs.readFileSync(`./src/view/${page}`, 'utf-8')
}

/**
 * 根據 URL
 * @param  {string} url koa2 上下文的 url, ctx.url
 * @return {string}     获取HTML文件内容
 */
async function route (url) {
  let view = '404.html'
  switch (url) {
    case '':
    case '/':
    case '/index':
    case '/index.html':
      view = 'index.html'
      break
    default:
      break
  }
  let html = await render(view)
  return html
}

const app = new Koa()
  .on('error', (err, ctx) => {
    console.error('server error', err, ctx)
  })
const server = new ApolloServer({
  schema: makeExecutableSchema({
    typeDefs,
    resolvers
  })
})

server.applyMiddleware({ app })

app
  .use(bodyParser())
  .use(async (ctx, next) => {
    if (!['', '/'].includes(ctx.request.url.substring(1))) {
      return next()
    }
    ctx.body = 'server alive!'
  })
  .use(async (ctx, next) => {
    if (!['shorten'].includes(ctx.request.url.substring(1))) {
      return next()
    }
    console.error(ctx.request.rawBody, ctx.request.body)
    const url = ctx.request.body.url

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
      ctx.body = find
      return
    }

    await db.shortened_urls.create({
      url,
      shortId
    })

    ctx.body = await db.shortened_urls.findOne({
      where: {
        shortId
      }
    })
  })
  .use(async (ctx, next) => {
    let shortid = ctx.request.url.substring(1)
    if (shortid.length !== 12) {
      return next()
    }
    const row = await db.shortened_urls.findOne({
      where: {
        shortid
      }
    })
    if (!row) {
      return next()
    }
    ctx.redirect(row.url)
  })
  .use(async (ctx) => {
    let url = ctx.request.url
    let html = await route(url)
    ctx.body = html
  })
module.exports = app
