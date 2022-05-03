const express = require('express')

const router = express.Router()
// importing Is auth controller
const isAuth = require('../middleware/is-auth')
// importing post controller
const postController = require('../controller/post')

router.get('/availablePosts', isAuth, postController.getAvailablePosts)

//  post create post
router.post('/create-post', isAuth, postController.postCreatePost)
// post like post 
router.post('/like-post', isAuth, postController.postLikePost)

module.exports = router