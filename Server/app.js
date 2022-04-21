const express = require('express')

// importing routers
const mainRouter = require('./router/main')
const authRouter = require('./router/auth')
//importing database 
const sequelize = require('./database/sequelize')
//importing custom middleware
const core = require('./middleware/core')

// Using express middleware
const app = express()
app.use(express.json())

// using body parser


//Using custom Middleware
app.use(core.core)

//using routes 
app.use(authRouter)
app.use(mainRouter)


//syncing database
sequelize.sync({ forced: true }).then(result => {
    // console.log(result)
    app.listen('8585')
}).catch(err => { console.log(err) })
