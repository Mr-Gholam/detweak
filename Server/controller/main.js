
// get main controller
exports.getMain = (req, res, next) => {
    res.sendFile(path.join(__dirname, 'index.html'))
    res.status(200).json({
        post: 'this is the test version'
    })
}
// post main controller
exports.postMain = (req, res, next) => {
    if (req.email) {
        res.status(200)
    } else {
        res.status(400)
    }
}
