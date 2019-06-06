// https://mongodb.github.io/node-mongodb-native/3.2/api/MongoClient.html
const mongo = require('mongodb').MongoClient;
const port = 4321;
const connString = `mongodb://localhost:${ port }`;

function db(opts) {
  return mongo.connect(connString, opts)
  	.then(mongoClient => {
  		console.log(`connected to ${ connString }`);
  		process.on('SIGINT', exit.bind(null, mongoClient));
  		return mongoClient.db('pets').collection('users');
  	});
};

function exit(mongoClient) {
	if (mongoClient) {
		mongoClient.close();
		console.log('mongoClient closed');
	}
	process.exit(0);
};

module.exports = db;
