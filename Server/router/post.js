const express = require('express')

const router = express.Router()

// importing post controller
const postController = require('../controller/post')

router.get('/availablePosts', postController.getAvailablePosts)

//  post create post
router.post('/create-post', postController.postCreatePost)

module.exports = router