const mongodb = require('mongodb')

const MongoClient = mongodb.MongoClient

let _db
const mongoConnect = (callback) => {
    MongoClient.connect('mongodb+srv://MrGholam:mehdi007@cluster0.bslwn.mongodb.net/Safe-Project?retryWrites=true&w=majority')
        .then(client => {
            _db = client._db
            callback()
        }
        ).catch(err => console.log(err))
}
const getDb = () => {
    if (_db) {
        return _db
    }
    throw 'No database found!'
}
exports.mongoConnect = mongoConnect
exports.getDb = getDb