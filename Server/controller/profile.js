// importing User model
const User = require('../model/user')
// importing Post model
const Post = require('../model/post')
// importing friend model 
const Friend = require('../model/friend')
// importing friend Req model 
const FriendReq = require('../model/friendReq')

const { Op } = require("sequelize");



// get profile controller
exports.getProfile = async (req, res, next) => {
    const username = req.params.username
    let userId
    let isFriend
    let sentRequest

    try {
        const user = await User.findOne({ where: { username }, attributes: ['id', 'username', 'firstName', 'lastName', 'profileImgUrl', 'onlineTime', 'bio', 'birthday', 'location'] })
        const Posts = await Post.findAll({ where: { userId: user.id } })
        const friend = await Friend.findAll({ where: { userId: user.id } })
        const friendCount = friend.length
        const postCount = Posts.length
        const availablePosts = Posts.map(
            post => {
                return {
                    description: post.description
                    , postImg: post.imageUrl
                    , username: user.username
                    , firstName: user.firstName
                    , lastName: user.lastName
                    , profileImg: user.profileImgUrl
                    , onlineTime: user.onlineTime
                    , createdAt: post.createdAt
                    , postId: post.id
                }
            })
        if (req.UserId) {
            userId = req.UserId
            const friendReq = await FriendReq.findOne({ where: { [Op.and]: [{ userId }, { targetId: user.id }] } })
            const IsFriend = await Friend.findOne({ where: { [Op.or]: [{ [Op.and]: [{ userId }, { targetId: user.id }] }, { [Op.and]: [{ userId: user.id }, { targetId: userId }] }] } })
            if (friendReq) {
                sentRequest = true
                isFriend = false
            } else if (IsFriend) {
                isFriend = true
                sentRequest = false
            } else {
                isFriend = false,
                    sentRequest = false
            }

        }
        res.status(200).json({
            username: username,
            firstName: user.firstName,
            lastName: user.lastName,
            profileImg: user.profileImgUrl,
            bio: user.bio,
            onlineTime: user.onlineTime,
            birthday: user.birthday,
            location: user.location,
            postCount: postCount,
            friendCount: friendCount,
            sentRequest,
            isFriend,
            availablePosts

        })
    } catch (err) { console.log(err) }

}