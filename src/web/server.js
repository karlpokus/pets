const apm = require('elastic-apm-node').start({ // note: swallows exceptions
	transactionSampleRate: 0.2
});

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
	.get("/stats", routes.stats)
	.get("/ping", routes.ping);

app
	.use(bodyparser())
	.use(router.routes())
	.use(router.allowedMethods())

module.exports = {app, db};
