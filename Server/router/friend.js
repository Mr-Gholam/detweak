const express = require('express')

const friendController = require('../controller/friend')
const router = express.Router()
const isAuth = require('../middleware/is-auth')
// post add friend
router.post('/add-friend', isAuth, friendController.postAddFriend)
// get friend requests
router.get('/friend-requests', isAuth, friendController.getFriendRequests)

module.exports = router