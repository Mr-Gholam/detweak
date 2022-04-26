const Sequelize = require('sequelize')

const sequelize = require('../database/sequelize')

const User = sequelize.define('users', {
    id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
    email: {
        type: Sequelize.STRING,
        allowNull: false
    },
    username: {
        type: Sequelize.STRING,
        allowNull: false
    },
    password: {
        type: Sequelize.STRING,
        allowNull: false
    },
    firstName: {
        type: Sequelize.STRING,
    },
    lastName: {
        type: Sequelize.STRING,
    },
    bio: {
        type: Sequelize.STRING,
    },
    profileImgUrl: {
        type: Sequelize.STRING,
    },
    onlineTime: {
        type: Sequelize.STRING,
    },
    birthday: {
        type: Sequelize.DATE,
    }
})
module.exports = User