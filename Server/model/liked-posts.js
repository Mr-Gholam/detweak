const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')
//  creating post in database
const LikedPosts = sequelize.define('liked-posts', {
    id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
})
// exporting model
module.exports = LikedPosts