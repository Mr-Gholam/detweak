const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')

const FriendReq = sequelize.define('friendRequest', {
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
module.exports = FriendReq