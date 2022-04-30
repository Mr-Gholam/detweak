// importing User model
const User = require('../model/user')
// importing Post model
const Post = require('../model/post')
// importing friend model 
const Friend = require('../model/friend')
// importing friend Req model
const FriendReq = require('../model/friendReq')

exports.postAddFriend = async (req, res, next) => {
    const targetUsername = req.body.targetUsername
    const target = await User.findOne({ where: { username: targetUsername } })
    const userId = req.UserId
    FriendReq.create({
        targetId: target.id,
        userId
    })
    res.status(200).end()
}
exports.getFriendRequests = async (req, res, next) => {
    const userId = req.UserId
    const requests = await FriendReq.findAll({ where: { targetId: userId } })
    let friendRequests = []
    requests.forEach(async (request) => {
        let ownerId = request.userId
        const owner = await User.findOne({
            where: {
                id: ownerId
            },
            attributes: ['id', 'firstName', 'lastName', 'profileImgUrl', 'onlineTime'],
        })
        friendRequests.push(owner)
    });
    res.json({ friendRequests })
}