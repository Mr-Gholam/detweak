
// get main controller
exports.getMain = (req, res, next) => {
    res.sendFile(path.join(__dirname, 'index.html'))
    res.status(200).json({
        post: 'this is the test version'
    })
}
// post main controller
exports.postMain = (req, res, next) => {
    if (req.cookies.jwt) {
        res.status(200).end()
    } else {
        res.status(400).end()
    }
}
