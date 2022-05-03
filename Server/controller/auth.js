
// importing bcrypt 
const bcrypt = require('bcryptjs')
// importing JWT
const jwt = require('jsonwebtoken')
// importing validation result
const { validationResult } = require('express-validator')
//importing user model
const User = require('../model/user')
// get Sign up 

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
        const emailFound = await User.findOne({ where: { email } })
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
                const newUser = await User.findOne({ where: { email } })
                const token = jwt.sign({ id: newUser.id, email, username }, 'superSecret')
                res.cookie('jwt', token, { maxAge: 1000 * 3600 })
                res.status(200).json({
                    username
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
        const matchPassword = await bcrypt.compare(password, emailFound.password)
        // passwords match
        if (matchPassword) {
            const token = jwt.sign({ id: emailFound.id, email, username: emailFound.username, profileImg: emailFound.profileImgUrl }, 'superSecret')
            res.cookie('jwt', token)
            res.status(200).json({
                username: emailFound.username,
                profileImg: emailFound.profileImgUrl
            })
        }
        // passwords do not match
        else {
            res.status(403).json({
                passwordErr: 'Password is incorrect'
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

// Post set profile
exports.postSetProfile = async (req, res, next) => {
    const email = req.email
    const firstName = req.body.firstName
    const lastName = req.body.lastName
    const bio = req.body.bio
    const onlineTime = req.body.onlineTime
    const birthday = req.body.birthday
    const location = req.body.location
    let imageUrl
    imageUrl = req.body.imageUrl
    if (req.file) {
        imageUrl = req.file.path
    }
    await User.update({
        firstName,
        lastName,
        bio,
        onlineTime,
        location,
        birthday,
        profileImgUrl: imageUrl
    }, { where: { email } })
    res.status(200).json({
        msg: 'user updated'
    })
}
// get Jwt 
exports.getJWT = async (req, res) => {
    try {
        if (req.cookies?.jwt) {
            const token = req.cookies.jwt
            const { id, username, email, profileImg } = jwt.verify(
                token,
                'superSecret'
            )
            res.status(200).json({
                id,
                username,
                email,
                profileImg
            })
        }

    } catch (error) {
        console.log(error)
    }
}
// post logout
exports.postLogout = async (req, res, next) => {
    res.cookie('jwt', '', { maxAge: -1 }).end()
}