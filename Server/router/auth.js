const express = require('express')
// importing express-validator
const { body } = require('express-validator')


const authController = require('../controller/auth')

const router = express.Router()

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
    , authController.postLogin)

module.exports = router