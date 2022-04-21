// importing bcrypt 
const bcrypt = require('bcryptjs')
//importing user model
const User = require('../model/user')
// Post Sign up controller
exports.postSignup = (req, res, next) => {
    const email = req.body.email
    const username = req.body.username
    const password = req.body.password
    // Looking for similar email
    User.findOne({ where: { email: email } }).then(result => {
        if (!result) {
            // Looking for similar username
            User.findOne({ where: { username } }).then(result => {
                if (!result) {
                    // hashing the password
                    bcrypt.hash(password, 12).then(hashedPassword => {
                        // creating the new user
                        User.create({
                            email: email,
                            username: username,
                            password: hashedPassword
                        }).then(result => {
                            // console.log(result)
                            res.status(301).json({

                            })
                        })
                    })

                }
                // found similar username in database
                else {
                    res.status(403).json({
                        UsernameErr: 'This Username has been used before'
                    })
                }
            })

        }
        // found similar Email in data base
        else {
            res.status(403).json({
                emailErr: 'This Email has been used before'
            })
        }
    }).catch(err => {
        console.log(err)
    })
}
//  Post Login controller
exports.postLogin = (req, res, next) => {
    const email = req.body.email
    const password = req.body.password
    // Looking for email in data base
    User.findOne({ where: { email } }).then(result => {
        // email found 
        if (result) {
            // comparing the password 
            bcrypt.compare(password, result.password).then(doMatch => {
                // passwords match
                if (doMatch) {
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
            )
        }
        // email not found
        else {
            res.status(403).json({
                emailErr: 'Email not found'
            })
        }
    }
    )
}