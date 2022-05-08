const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')
// create comment in database
const Chat = sequelize.define('chat', {
    id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
    massage: {
        type: Sequelize.STRING,
        allowNull: false
    },
    targetId: {
        type: Sequelize.INTEGER,
        allowNull: false
    }
})
// export model
module.exports = Chat