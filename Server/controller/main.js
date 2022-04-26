exports.getMain = (req, res, next) => {
    res.status(200).json({
        post: 'this is the test version'
    })
}
