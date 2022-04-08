exports.getMain = (req, res, next) => {
    res.status(200).json({
        post: 'this is the test version'
    })
}

exports.postCreatePost = (req, res, next) => {
    console.log(req.body)
    res.status(200).json({
        postStatus: 'post created'
    })
}