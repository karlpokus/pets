#!/usr/bin/env node

const secrets = require("./secrets");
const apm = require('elastic-apm-node').start(secrets.apm)
const port = process.env.PORT || 9012;

const Koa = require('koa');
const Router = require('koa-router');
const auth = require('./lib/auth');
const db = require('./lib/db');
const http = require('./lib/http');

const app = new Koa();
const router = new Router({ prefix: "/api/v1" });

function getPets(ctx) { // should be in lib
	return http.get('/pets') // should be in a config
		.then(res => ctx.body = res.data)
		.catch(err => ctx.throw(500, "server error"))
}

router.get("/pets", auth, getPets);

app
	.use(router.routes())
	.use(router.allowedMethods())

db(secrets.mongoOpts)
	.then(users => {
		app.context.users = users;
		app.listen(port, () => { console.log(`pets-web running on port ${ port }`) })
	})
	.catch(err => {
		console.error(err);
		process.exit(1);
	})
