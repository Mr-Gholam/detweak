// importing post model
const Post = require('../model/post')
// import user model 
const User = require('../model/user')
// importing friend model 
const Friend = require('../model/friend')
// importing like posts model 
const LikedPosts = require('../model/liked-posts')
// importing comment model 
const Comment = require('../model/comment')
const { Op } = require("sequelize");
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
        const Posts = await Post.findAll({ where: { userId: id }, attributes: ['description', 'imageUrl', 'createdAt', 'id', 'likes'] })
        const user = await User.findOne({ where: { id }, attributes: ['firstName', 'username', 'lastName', 'profileImgUrl', 'onlineTime'] })
        for (let i = 0; i < Posts.length; i++) {
            const post = Posts[i];
            const likedPost = await LikedPosts.findOne({ where: { [Op.and]: [{ userId }, { postId: post.id }] } })
            if (likedPost) {
                const pushedPost = {
                    description: post.description
                    , postImg: post.imageUrl
                    , username: user.username
                    , firstName: user.firstName
                    , lastName: user.lastName
                    , profileImg: user.profileImgUrl
                    , onlineTime: user.onlineTime
                    , createdAt: post.createdAt
                    , likes: post.likes
                    , postId: post.id
                    , liked: true
                }
                availablePosts.push(pushedPost)

            } else {
                const pushedPost = {
                    description: post.description
                    , postImg: post.imageUrl
                    , username: user.username
                    , firstName: user.firstName
                    , lastName: user.lastName
                    , profileImg: user.profileImgUrl
                    , onlineTime: user.onlineTime
                    , createdAt: post.createdAt
                    , likes: post.likes
                    , postId: post.id
                    , liked: false
                }
                availablePosts.push(pushedPost)
            }
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
        likes: 0,
        imageUrl,
        userId
    }).then(() => {

        res.status(200).json({ message: 'post created ' })
    }).catch(err => console.log(err))
}
// POST for like post 
exports.postLikePost = async (req, res, next) => {
    if (req.UserId) {
        const userId = req.UserId
        const postId = req.body.postId
        const likedPost = await LikedPosts.findOne({ where: { [Op.and]: [{ userId }, { postId }] } })
        const post = await Post.findOne({ where: { id: postId }, attributes: ['likes'] })

        if (!likedPost) {
            const newLikes = post.likes + 1
            await Post.update({ likes: newLikes }, { where: { id: postId } })
            await LikedPosts.create({
                userId,
                postId
            })
            res.json({ added: true })
        } else {
            let newLikes
            if (post.likes == 0) {
                newLikes = 0
            } else {
                newLikes = post.likes - 1
            }
            await Post.update({ likes: newLikes }, { where: { id: postId } })
            await LikedPosts.destroy({ where: { [Op.and]: [{ userId }, { postId }] } })
            res.json({ added: false })
        }
    } else {
        res.status(401)
    }
}
// POST for Add comment 
exports.postAddComment = async (req, res, next) => {
    if (req.UserId) {
        const userId = req.UserId
        const postId = req.body.postId
        const commentContent = req.body.comment
        const user = await User.findOne({ where: { id: userId }, attributes: ['profileImgUrl', 'username'] })
        const newComment = await Comment.create({
            userId,
            postId,
            comment: commentContent
        })
        res.status(200).json({
            commentContent,
            profileImg: user.profileImgUrl,
            username: user.username,
            createdAt: newComment.createdAt,
            commentId: newComment.id
        })
    } else {
        res.status(401)
    }
}
// GET a single post
exports.getAPost = async (req, res, next) => {
    if (req.UserId) {
        try {

            let comments = []
            let postInfo = []
            const postId = req.params.postId
            const userId = req.UserId
            const post = await Post.findOne({ where: { id: postId }, attributes: ['description', 'imageUrl', 'createdAt', 'id', 'likes', 'userId'] })
            const postOwner = await User.findOne({ where: { id: post.userId } })
            const likedPost = await LikedPosts.findOne({ where: { [Op.and]: [{ userId }, { postId }] } })
            const currentComments = await Comment.findAll({ where: { postId }, attributes: ['userId', 'comment', 'createdAt', 'id'] })
            for (let i = 0; i < currentComments.length; i++) {
                const comment = currentComments[i];
                const user = await User.findOne({ where: { id: comment.userId }, attributes: ['profileImgUrl', 'username'] })
                const pushComment = {
                    commentContent: comment.comment,
                    profileImg: user.profileImgUrl,
                    username: user.username,
                    createdAt: comment.createdAt,
                    commentId: comment.id
                }
                comments.push(pushComment)
            }
            if (likedPost) {
                const pushedPost = {
                    description: post.description
                    , postImg: post.imageUrl
                    , username: postOwner.username
                    , firstName: postOwner.firstName
                    , lastName: postOwner.lastName
                    , profileImg: postOwner.profileImgUrl
                    , onlineTime: postOwner.onlineTime
                    , createdAt: post.createdAt
                    , likes: post.likes
                    , postId: post.id
                    , liked: true
                }
                postInfo.push(pushedPost)
            } else {
                const pushedPost = {
                    description: post.description
                    , postImg: post.imageUrl
                    , username: postOwner.username
                    , firstName: postOwner.firstName
                    , lastName: postOwner.lastName
                    , profileImg: postOwner.profileImgUrl
                    , onlineTime: postOwner.onlineTime
                    , createdAt: post.createdAt
                    , likes: post.likes
                    , postId: post.id
                    , liked: false
                }
                postInfo.push(pushedPost)
            }
            res.json({ postInfo, comments })
        }
        catch (err) { }
    } else {
        res.status(401)
    }
}