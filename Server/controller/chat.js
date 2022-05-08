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
    res.json({
        username: target.username,
        firstName: target.firstName,
        lastName: target.lastName,
        profileImg: target.profileImgUrl,
        chatRoom: id.chatRoomId
    })
}