// importing User model
const User = require('../model/user')
// importing Post model
const Post = require('../model/post')
// importing friend model 
const Friend = require('../model/friend')
// importing friend Req model
const FriendReq = require('../model/friendReq')
const res = require('express/lib/response')

exports.postAddFriend = async (req, res, next) => {
    const targetUsername = req.body.targetUsername
    try {
        const target = await User.findOne({ where: { username: targetUsername }, attributes: ['id'] })
        const userId = req.UserId
        FriendReq.create({
            targetId: target.id,
            userId
        })
        res.status(200).end()

    } catch (err) {
        console.log(err)
    }
}
exports.getFriendRequests = async (req, res, next) => {
    const userId = req.UserId
    try {
        const requests = await FriendReq.findAll({ where: { targetId: userId } })
        const friendRequests = []
        for (let i = 0; i < requests.length; i++) {
            const request = requests[i];
            let ownerId = request.userId
            const owners = await User.findOne({
                where: {
                    id: ownerId
                },
                attributes: ['id', 'firstName', 'lastName', 'profileImgUrl', 'onlineTime', 'username'],
            })
            const owner = {
                requestId: request.id,
                id: owners.id,
                firstName: owners.firstName,
                username: owners.username,
                lastName: owners.lastName,
                profileImg: owners.profileImgUrl,
                onlineTime: owners.onlineTime
            }
            friendRequests.push(owner)
        }
        res.json({ friendRequests })

    } catch (err) {
        console.log(err)
    }
}
exports.postAcceptRequest = async (req, re, next) => {
    const requestId = req.body.requestId
    const request = await FriendReq.findOne({ where: { id: requestId }, attributes: ['userId', 'targetId',] })
    await Friend.create({
        targetId: request.targetId,
        userId: request.userId
    })
    FriendReq.destroy({
        where: { id: requestId }
    }).then(result => {
        res.status(200)
    })
}