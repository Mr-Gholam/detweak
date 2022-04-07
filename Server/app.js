const express = require('express')

// importing routers
const mainRouter = require('./router/main')
const authRouter = require('./router/auth')
//importing database 
const mongoConnect = require('./database/mongodb').mongoConnect
//importing custom middleware
const core = require('./middleware/core')

// Using express middleware
const app = express()
app.use(express.json())

//Using custom Middleware
app.use(core.core)

//using routes 
app.use(authRouter)
app.use(mainRouter)


mongoConnect(() => {
    console.log('mongodb connected')
    app.listen('8585')
})