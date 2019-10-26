const { timestamps, paranoid } = require('../migrations-model')

const tableName = 'shortened_urls'
module.exports = {
  up: async (queryInterface, Sequelize) => {
    let DataTypes = Sequelize
    return [
      await queryInterface.createTable(tableName, {
        id: { type: DataTypes.INTEGER.UNSIGNED, primaryKey: true, autoIncrement: true },
        url: {
          type: DataTypes.STRING
        },
        shortid: {
          type: DataTypes.STRING,
          unique: true
        },
        ...timestamps,
        ...paranoid
      }, { charset: 'utf8mb4' })
    ]
  },

  down: async (queryInterface, Sequelize) => {
    return [
      await queryInterface.dropTable(tableName)
    ]
  }
}
