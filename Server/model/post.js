const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')
//  creating post in database
const Post = sequelize.define('posts', {
    id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
    imageUrl: {
        type: Sequelize.STRING
    },
    description: {
        type: Sequelize.STRING,
        allowNull: false
    },
    likes: {
        type: Sequelize.INTEGER
    }
})
// exporting model
module.exports = Post
