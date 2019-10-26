const sequelize = require('sequelize')
let DataTypes = sequelize

module.exports = {
  timestamps: {
    created_at: { allowNull: false, type: DataTypes.DATE },
    updated_at: { allowNull: false, type: DataTypes.DATE }
  },
  paranoid: {
    deleted_at: DataTypes.DATE
  }
}
