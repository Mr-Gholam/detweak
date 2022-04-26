const express = require('express')

// importing path 
const path = require('path')
// importing multer
const multer = require('multer')

const cookieParser = require('cookie-parser');
// creating multer options
const fileDestination = multer.diskStorage({
    destination: (req, file, cb) => {
        cb(null, 'images')
    },
    filename: (req, file, cb) => {
        cb(null, new Date().toISOString().replace(/:/g, '-') + '-' + file.originalname)
    }
})
const fileFilter = (req, file, cb) => {
    if (
        file.mimetype === 'image/png' ||
        file.mimetype === 'image/jpg' ||
        file.mimetype === 'image/jpeg') {
        cb(null, true)
    } else {
        cb(null, false)
    }
}
// importing routers
const mainRouter = require('./router/main')
const authRouter = require('./router/auth')
const postRouter = require('./router/post')
//importing database 
const sequelize = require('./database/sequelize')
//importing custom middleware
const core = require('./middleware/core')

// importing models 
const User = require('./model/user')
const Post = require('./model/post')
const Comment = require('./model/comment')
const Friend = require('./model/friend')

// Using express middleware
const app = express()
app.use(cookieParser());
app.use(express.json())
// using multer
app.use(multer({ storage: fileDestination, fileFilter: fileFilter }).single('image'))
app.use('/images', express.static(path.join(__dirname, 'images')))



//Using custom Middleware
app.use(core.core)

//using routes 
app.use(authRouter)
app.use(postRouter)
app.use(mainRouter)

// relation between models
User.hasMany(Post)
User.hasMany(Comment)
User.hasMany(Friend)
Post.hasMany(Comment)


//syncing database
sequelize.sync().then(result => {
    // console.log(result)
    app.listen('8585')
}).catch(err => { console.log(err) })
