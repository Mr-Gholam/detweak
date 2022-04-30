// importing post model
const Post = require('../model/post')
// import user model 
const User = require('../model/user')
// GET for available Posts
exports.getAvailablePosts = async (req, res, next) => {
    const userId = 1
    const Posts = await Post.findAll()
    const user = await User.findOne({ where: { Id: 1 } })
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