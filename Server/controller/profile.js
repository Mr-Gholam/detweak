// importing User model
const User = require('../model/user')
// importing Post model
const Post = require('../model/post')
// importing friend model 
const Friend = require('../model/friend')



// get profile controller
exports.getProfile = async (req, res, next) => {
    const username = req.params.username
    try {
        const user = await User.findOne({ where: { username } })
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
            availablePosts

        })
    } catch (err) { console.log(err) }

}