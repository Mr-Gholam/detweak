// importing bcrypt 
const bcrypt = require('bcryptjs')
// importing validation result
const { validationResult } = require('express-validator')
//importing user model
const User = require('../model/user')
// Post Sign up controller
exports.postSignup = async (req, res, next) => {
    // Looking for validation errors
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
        return res.status(400).json({ errors: errors.array() });
    }
    const email = req.body.email
    const username = req.body.username
    const password = req.body.password
    try {
        // Looking for similar email
        const emailFound = await User.findOne({ where: { email: email } })
        if (!emailFound) {
            // Looking for similar username
            const usernameFound = await User.findOne({ where: { username } })
            if (!usernameFound) {
                // hashing the password
                const hashedPassword = await bcrypt.hash(password, 12)
                // creating the new user
                await User.create({
                    email: email,
                    username: username,
                    password: hashedPassword,
                })
                res.status(301).json({

                })

            }
            // found similar username in database
            else {
                res.status(403).json({
                    UsernameErr: 'This Username has been used before'
                })
            }

        }
        // found similar Email in data base
        else {
            res.status(403).json({
                emailErr: 'This Email has been used before'
            })
        }
    } catch (err) {
        console.log(err)
    }


}
//  Post Login controller
exports.postLogin = async (req, res, next) => {
    const errors = validationResult(req)
    if (!errors.isEmpty()) {
        return res.status(400).json({ errors: errors.array() });
    }
    const email = req.body.email
    const password = req.body.password
    // Looking for email in data base
    const emailFound = await User.findOne({ where: { email } })
    // email found 
    if (emailFound) {
        // comparing the password 
        const matchPassword = await bcrypt.compare(password, result.password)
        // passwords match
        if (matchPassword) {
            res.status(200).json({
                IsLoggedIn: true
            })
        }
        // passwords do not match
        else {
            res.status(403).json({
                IsLoggedIn: false
            })
        }

    }

    // email not found
    else {
        res.status(403).json({
            emailErr: 'Email not found'
        })
    }
}
