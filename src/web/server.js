#!/usr/bin/env node

const apm = require('elastic-apm-node').start()
const port = process.env.PORT || 9012;

const Koa = require('koa');
const Router = require('koa-router');
const auth = require('./lib/auth');
const db = require('./lib/db');
const pets = require('./lib/pets');

const app = new Koa();
const router = new Router({ prefix: "/api/v1" });

router.get("/pets", auth, pets.allPets);

app
	.use(router.routes())
	.use(router.allowedMethods())

db()
	.then(users => {
		app.context.users = users;
		app.listen(port, () => { console.log(`web vX.Y.Z listening on port ${ port }`) })
	})
	.catch(err => {
		console.error(err);
		process.exit(0); // do not restart from unrecoverable err
	})
