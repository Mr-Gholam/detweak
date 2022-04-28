const JWT = require('jsonwebtoken')


module.exports = (req, res, next) => {
    const token = req.cookies.jwt
    let decodedToken
    try {
        decodedToken = JWT.verify(token, 'superSecret')
    } catch (err) {
        console.log(err)
    }
    if (!decodedToken) {
        const error = new Error('Not authenticated')
        error.statusCode = 401
        throw error
    }
    req.userId = decodedToken.userId
    req.email = decodedToken.email
    next()
}