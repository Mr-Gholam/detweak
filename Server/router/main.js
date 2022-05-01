const express = require('express')

const mainController = require('../controller/main')
const isAuth = require('../middleware/is-auth')

const router = express.Router()


router.get('/search/:userInput', isAuth, mainController.getSearch)
router.post('/', isAuth, mainController.postMain)

router.get('/', mainController.getMain)




module.exports = router