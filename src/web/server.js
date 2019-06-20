const apm = require('elastic-apm-node').start();
const port = process.env.HTTP_PORT;
const host = process.env.HTTP_HOST;
const version = process.env.npm_package_version;

const Koa = require('koa');
const Router = require('koa-router');
const bodyparser = require('koa-bodyparser');
const auth = require('./lib/auth');
const db = require('./lib/db');
const routes = require('./lib/routes');

const app = new Koa();
const router = new Router({ prefix: "/api/v1" });

router
	.get("/pets", auth, routes.allPets)
	.post("/pet", auth, routes.addPet)
	.get("/stats", routes.stat)
	.get("/ping", routes.ping);

app
	.use(bodyparser())
	.use(router.routes())
	.use(router.allowedMethods())

db()
	.then(mongoClient => {
		app.context.mongoClient = mongoClient;
		app.listen(port, host, () => {
			console.log(`web ${ version } listening on ${ host }:${ port }`)
		})
	})
	.catch(err => {
		console.error(err);
		process.exit(0);
	})
