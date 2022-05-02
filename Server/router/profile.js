const express = require('express')

const profileController = require('../controller/profile')
const router = express.Router()
const isAuth = require('../middleware/is-auth')

router.get('/profile/:username', isAuth, profileController.getProfile)


module.exports = router