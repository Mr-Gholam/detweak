// import user model 
const User = require('../model/user')
const { Op } = require("sequelize");
const Friend = require('../model/friend')
const FriendReq = require('../model/friendReq')
// get search controller
exports.getSearch = async (req, res, next) => {
    let usersFound = []
    const userInput = req.params.userInput
    const users = await User.findAll(
        {
            where: { [Op.or]: [{ username: { [Op.like]: userInput } }, { firstName: { [Op.like]: userInput } }, { lastName: { [Op.like]: userInput } }] }
            , attributes: ['firstName', 'username', 'lastName', 'profileImgUrl', 'id']
        })
    for (let i = 0; i < users.length; i++) {
        const user = users[i];
        let userFound
        if (req.UserId) {
            const userId = req.UserId
            let myProfile = false
            if (userId === user.id) myProfile = true
            const isFriend = await Friend.findOne({ where: { [Op.or]: [{ [Op.and]: [{ userId }, { targetId: user.id }] }, { [Op.and]: [{ userId: user.id }, { targetId: userId }] }] } })
            const friendReq = await FriendReq.findOne({ where: { [Op.and]: [{ userId }, { targetId: user.id }] } })
            if (isFriend) {
                userFound = {
                    username: user.username,
                    firstName: user.firstName,
                    lastName: user.lastName,
                    profileImg: user.profileImgUrl,
                    id: user.id,
                    myProfile,
                    isFriend: true,
                    sentRequest: false
                }
                usersFound.push(userFound)
            } else if (friendReq) {
                userFound = {
                    username: user.username,
                    firstName: user.firstName,
                    lastName: user.lastName,
                    profileImg: user.profileImgUrl,
                    id: user.id,
                    myProfile,
                    isFriend: false,
                    sentRequest: true
                }
                usersFound.push(userFound)
            } else {
                userFound = {
                    username: user.username,
                    firstName: user.firstName,
                    lastName: user.lastName,
                    profileImg: user.profileImgUrl,
                    id: user.id,
                    myProfile,
                    isFriend: false,
                    sentRequest: false
                }
                usersFound.push(userFound)
            }
        } else {
            userFound = {
                username: user.username,
                firstName: user.firstName,
                lastName: user.lastName,
                profileImg: user.profileImgUrl,
                id: user.id,
                myProfile,
                isFriend: false,
                sentRequest: false
            }
            usersFound.push(userFound)
        }

    }
    res.json({ usersFound })
}
// get main controller
exports.getMain = (req, res, next) => {
    res.status(200).json({
        post: 'this is the test version'
    })
}
// post main controller
exports.postMain = (req, res, next) => {
    if (req.cookies.jwt) {
        res.status(200).end()
    } else {
        res.status(400).end()
    }
}
