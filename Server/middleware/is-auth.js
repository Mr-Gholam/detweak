const JWT = require('jsonwebtoken')


module.exports = (req, res, next) => {
    let token
    let decodedToken
    if (req.cookies.jwt) {
        token = req.cookies.jwt
    }
    try {
        if (token) {
            decodedToken = JWT.verify(token, 'superSecret')
            if (!decodedToken) {
                const error = new Error('Not authenticated')
                error.statusCode = 401
                throw error
            }
            req.email = decodedToken.email
        }
    } catch (err) {
        console.log(err)
    }
    next()
}