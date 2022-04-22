const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')
// create comment in database
const Comment = sequelize.define('comments', {
    id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
    comment: {
        type: Sequelize.STRING
    }
})
// export model
module.exports = Comment