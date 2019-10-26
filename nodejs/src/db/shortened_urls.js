module.exports = function (sequelize, DataTypes) {
  const tableName = 'shortened_urls'
  let model = sequelize.define(tableName, {
    id: { type: DataTypes.INTEGER.UNSIGNED, primaryKey: true, autoIncrement: true },
    url: {
      type: DataTypes.STRING
    },
    shortId: {
      type: DataTypes.STRING,
      unique: true,
      field: 'shortid'
    }
  })
  return model
}
