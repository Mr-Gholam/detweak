const { Op } = require("sequelize");
const Chat = require('../model/chat')
const ChatRoom = require('../model/chatroom')
const User = require('../model/user')
exports.getChats = async (req, res, next) => {
    const userId = req.UserId
    const ids = []
    const contacts = []
    const targetIds = await ChatRoom.findAll({ where: { userId }, attributes: ['targetId', 'id'] })
    const userIds = await ChatRoom.findAll({ where: { targetId: userId }, attributes: ['userId', 'id'] })
    for (let i = 0; i < targetIds.length; i++) {
        const targetId = targetIds[i];
        ids.push({
            id: targetId.targetId,
            chatRoomId: targetId.id
        })
    }
    for (let i = 0; i < userIds.length; i++) {
        const targetId = userIds[i];
        ids.push({
            id: targetId.userId,
            chatRoomId: targetId.id
        })
    }
    for (let i = 0; i < ids.length; i++) {
        const id = ids[i];
        const user = await User.findOne({ where: { id: id.id }, attributes: ['username', 'firstName', 'lastName', 'profileImgUrl'] })
        contacts.push({
            username: user.username,
            firstName: user.firstName,
            lastName: user.lastName,
            profileImg: user.profileImgUrl,
            chatRoomId: id.chatRoomId
        })
    }
    res.json({ contacts })
}
exports.postCreateRoom = async (req, res, next) => {
    const userId = req.UserId
    const targetUsername = req.body.targetUsername
    const target = await User.findOne({ where: { username: targetUsername }, attributes: ['id', 'username', 'firstName', 'lastName', 'profileImgUrl'] })
    await ChatRoom.create({ userId, targetId: target.id })
    const chatRoom = await ChatRoom.findOne({ where: { [Op.and]: [{ userId }, { targetId: target.id }] }, attributes: ['id'] })
    res.json({
        username: target.username,
        firstName: target.firstName,
        lastName: target.lastName,
        profileImg: target.profileImgUrl,
        chatRoomId: chatRoom.id
    })
}
exports.postChat = async (req, res, next) => {
    const userId = req.UserId
    const chatRoomId = req.body.chatRoomId
    const chats = await Chat.findAll({ where: { chatRoomId }, attributes: ['targetId', 'message', 'createdAt'] })
    const chatRoom = await ChatRoom.findOne({ where: { id: chatRoomId } })
    if (chatRoom.userId == userId) {
        let targetId = chatRoom.targetId
        const target = await User.findOne({ where: { id: targetId }, attributes: ['username', 'firstName', 'lastName', 'profileImgUrl', 'onlineTime', 'id'] })
        const chatInfo = {
            username: target.username,
            firstName: target.firstName,
            lastName: target.lastName,
            profileImg: target.profileImgUrl,
            onlineTime: target.onlineTime,
            targetId: target.id,
            chatRoomId
        }
        let currentChat = []
        for (let i = 0; i < chats.length; i++) {
            const chat = chats[i];
            if (chat.targetId == userId) {
                currentChat.push({
                    receive: true,
                    message: chat.message,
                    createdAt: chat.createdAt
                })
            } else {
                currentChat.push({
                    receive: false,
                    message: chat.message,
                    createdAt: chat.createdAt
                })
            }
        }
        currentChat.sort((a, b) => { a.createdAt - b.createdAt })
        res.json({ currentChat, chatInfo })
    } else {
        let targetId = chatRoom.userId
        const target = await User.findOne({ where: { id: targetId }, attributes: ['username', 'firstName', 'lastName', 'profileImgUrl', 'onlineTime', 'id'] })
        const chatInfo = {
            username: target.username,
            firstName: target.firstName,
            lastName: target.lastName,
            profileImg: target.profileImgUrl,
            onlineTime: target.onlineTime,
            targetId: target.id,
            chatRoomId
        }
        let currentChat = []
        for (let i = 0; i < chats.length; i++) {
            const chat = chats[i];
            if (chat.targetId == userId) {
                currentChat.push({
                    receive: true,
                    message: chat.message,
                    createdAt: chat.createdAt
                })
            } else {
                currentChat.push({
                    receive: false,
                    message: chat.message,
                    createdAt: chat.createdAt
                })
            }
        }
        currentChat.sort((a, b) => { a.createdAt - b.createdAt })
        res.json({ currentChat, chatInfo })
    }


}
exports.postCreateMessage = async (req, res, next) => {
    const message = req.body.message
    const targetId = req.body.targetId
    const chatroomId = req.body.chatRoomId
    await Chat.create({
        message,
        targetId,
        chatroomId
    })
    res.status(200).end()
}