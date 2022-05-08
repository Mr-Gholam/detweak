const express = require('express')
// importing express-validator
const { body } = require('express-validator')
// importing is auth middleware 
const isAuth = require('../middleware/is-auth')
// importing auth main controller
const chatController = require('../controller/chat')

const router = express.Router()


router.get('/load-chatRooms', isAuth, chatController.getChats)

router.post('/create-room', isAuth, chatController.postCreateRoom)



module.exports = router