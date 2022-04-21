const Sequelize = require('sequelize')

const sequelize = new Sequelize('secret', 'root', 'mehdi007', {
    host: 'localhost',
    dialect: 'mysql',
    logging: false
})
module.exports = sequelize