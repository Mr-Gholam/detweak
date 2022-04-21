const express = require('express')

const authController = require('../controller/auth')

const router = express.Router()

// post signup
router.post('/signup', authController.postSignup)

// post login
router.post('/login', authController.postLogin)


module.exports = router