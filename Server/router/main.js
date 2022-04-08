const express = require('express')

const mainController = require('../controller/main')

const router = express.Router()


router.post('/create-post', mainController.postCreatePost)


router.get('/', mainController.getMain)



module.exports = router