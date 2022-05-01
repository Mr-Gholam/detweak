// importing post model
const Post = require('../model/post')
// import user model 
const User = require('../model/user')
// importing friend model 
const Friend = require('../model/friend')
// GET for available Posts
exports.getAvailablePosts = async (req, res, next) => {
    const userId = req.UserId
    const friendTargetIds = await Friend.findAll({ where: { userId }, attributes: ['targetId'] })
    const friendUserIds = await Friend.findAll({ where: { targetId: userId }, attributes: ['userId'] })
    const ids = []
    let availablePosts = []
    if (friendTargetIds.length != 0) {
        for (let i = 0; i < friendTargetIds.length; i++) {
            const id = friendTargetIds[i];
            ids.push(id.targetId)
        }
    }
    if (friendUserIds.length != 0) {
        for (let i = 0; i < friendUserIds.length; i++) {
            const id = friendUserIds[i];
            ids.push(id.userId)
        }
    }
    ids.push(userId)
    for (let i = 0; i < ids.length; i++) {
        const id = ids[i];
        const Posts = await Post.findAll({ where: { userId: id }, attributes: ['description', 'imageUrl', 'createdAt', 'id'] })
        const user = await User.findOne({ where: { id }, attributes: ['firstName', 'username', 'lastName', 'profileImgUrl', 'onlineTime'] })
        for (let i = 0; i < Posts.length; i++) {
            const post = Posts[i];
            const pushedPost = {
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
            availablePosts.push(pushedPost)
        }
        availablePosts.sort((a, b) => b.createdAt - a.createdAt)
    }
    res.status(200).json({ availablePosts })
}

// POST for creat post
exports.postCreatePost = async (req, res, next) => {
    let imageUrl
    if (req.file) {
        imageUrl = req.file.path
    }
    const description = req.body.description
    const userId = req.UserId
    Post.create({
        description,
        imageUrl,
        userId
    }).then(() => {

        res.status(200).json({ message: 'post created ' })
    }).catch(err => console.log(err))
}