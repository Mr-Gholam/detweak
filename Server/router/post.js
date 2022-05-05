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
// post add comment
router.post('/add-comment', isAuth, postController.postAddComment)
// get a single post 
router.get('/post/:postId', isAuth, postController.getAPost)
// post delete comment
router.post('/delete-Comment', isAuth, postController.postDeleteComment)
// post delete post
router.post('/delete-post', isAuth, postController.postDeletePost)
// post update post 
router.post('/update-post', isAuth, postController.postUpdatePost)

module.exports = router