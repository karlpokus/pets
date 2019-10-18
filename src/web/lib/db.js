// https://mongodb.github.io/node-mongodb-native/3.2/api/MongoClient.html
const mongo = require('mongodb').MongoClient;
const host = process.env.MONGODB_HOST;
const port = process.env.MONGODB_PORT;
const connString = `mongodb://${ host }:${ port }`;
const opts = {
  useNewUrlParser: true,
  appname: "pets-web"
}

function db() {
  return mongo.connect(connString, opts)
  	.then(client => {
  		console.log('connected to db');
  		process.on('SIGINT', exit.bind(null, client));
      return client;
  	});
};

function exit(client) {
	if (client.isConnected()) {
		client.close();
		console.log('db closed');
	}
	process.exit(0);
};

module.exports = db;
