const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')
// create comment in database
const friend = sequelize.define('friends', {
    id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
    targetId: {
        type: Sequelize.INTEGER,
        allowNull: false
    }
})
// export model
module.exports = friend