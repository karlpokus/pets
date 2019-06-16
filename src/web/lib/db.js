// https://mongodb.github.io/node-mongodb-native/3.2/api/MongoClient.html
const mongo = require('mongodb').MongoClient;
const port = process.env.MONGODB_PORT;
const connString = `mongodb://localhost:${ port }`;
const opts = {
  useNewUrlParser: true,
  authSource: process.env.MONGODB_PETS_AUTHSOURCE,
  appname: "pets-web",
  auth: {
    user: process.env.MONGODB_PETS_USER,
    password: process.env.MONGODB_PETS_PWD
  }
}

function db() {
  return mongo.connect(connString, opts)
  	.then(mongoClient => {
  		console.log('connected to db');
  		process.on('SIGINT', exit.bind(null, mongoClient));
  		return mongoClient.db('pets').collection('users');
  	});
};

function exit(mongoClient) {
	if (mongoClient) {
		mongoClient.close();
		console.log('db closed');
	}
	process.exit(0);
};

module.exports = db;
