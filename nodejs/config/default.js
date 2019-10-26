const data = {
  app: {
    env: process.env.NODE_ENV || 'development',
    port: process.env.PORT || 4000,
    debug: true
  },
  // see http://docs.sequelizejs.com/class/lib/sequelize.js~Sequelize.html#instance-constructor-constructor
  db: {
    host: 'mysql', // in container docker-compose service name
    port: 3306,
    username: 'root',
    password: 'root',
    database: 'test',
    dialect: 'mysql',
    charset: 'utf8mb4',
    pool: {
      max: 5,
      min: 0,
      idle: 10000
    },
    // see http://docs.sequelizejs.com/class/lib/model.js~Model.html#static-method-init
    define: {
      paranoid: true,
      underscored: true,
      freezeTableName: true
    },
    benchmark: true
  }
}

module.exports = data
