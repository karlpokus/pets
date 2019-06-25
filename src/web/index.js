const { app, db } = require('./server');

const port = process.env.HTTP_PORT;
const host = process.env.HTTP_HOST;
const version = process.env.npm_package_version;

db()
	.then(mongoClient => {
		app.context.mongoClient = mongoClient;
		app.listen(port, host, () => {
			console.log(`web ${ version } listening on ${ host }:${ port }`)
		})
	})
	.catch(err => {
		console.error(`Attempt to connect to db failed: ${ err }`);
		process.exit(0);
	})
