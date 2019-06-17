const apm = require('elastic-apm-node').start();
const port = process.env.HTTP_PORT;
const host = process.env.HTTP_HOST;
const version = process.env.npm_package_version;

const Koa = require('koa');
const Router = require('koa-router');
const bodyparser = require('koa-bodyparser');
const auth = require('./lib/auth');
const db = require('./lib/db');
const pets = require('./lib/pets');

const app = new Koa();
const router = new Router({ prefix: "/api/v1" });

router
	.get("/pets", auth, pets.allPets)
	.post("/pet", auth, pets.addPet);

app
	.use(bodyparser())
	.use(router.routes())
	.use(router.allowedMethods())

db()
	.then(users => {
		app.context.users = users;
		app.listen(port, host, () => {
			console.log(`web ${ version } listening on ${ host }:${ port }`)
		})
	})
	.catch(err => {
		console.error(err);
		process.exit(0); // do not restart from unrecoverable err
	})
