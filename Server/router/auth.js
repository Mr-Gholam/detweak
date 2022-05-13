const express = require('express')
// importing express-validator
const { body } = require('express-validator')
// importing is auth middleware 
const isAuth = require('../middleware/is-auth')
// importing auth main controller
const authController = require('../controller/auth')



const router = express.Router()
// get jwt
router.get('/jwt', authController.getJWT)
// post signup
router.post('/signup',
    [
        body('email').isEmail().normalizeEmail().notEmpty(),
        body('username').trim().isLength({ min: 3 }).notEmpty(),
        body('password').trim().isLength({ min: 8 }).isAlphanumeric().notEmpty(),
        body('confirmPassword').trim().custom((value, { req }) => {
            if (value !== req.body.password) {
                throw new Error('Password confirmation does not match password');
            }
            return true
        })
    ]
    , authController.postSignup)

// post login
router.post('/login',
    [
        body('email').isEmail().normalizeEmail().notEmpty(),
        body('password').trim().isAlphanumeric().isLength({ min: 8 })
    ]
    , isAuth, authController.postLogin)

// post set profile
router.post('/set-profile', isAuth, authController.postSetProfile)
// Get UserInfo 
router.get('/userInfo', isAuth, authController.getUserInfo)
// post Update profile
router.post('/update-profile', isAuth, authController.postUpdateProfile)
// post update account 
router.post('/update-account', isAuth, authController.postUpdateAccount)
// post logout 
router.post('/logout', authController.postLogout)

module.exports = router