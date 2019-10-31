// grabs user and pwd from authorization header
// returns an empty array on failure
function basicAuth(authHeader) {
	if (/^Basic/.test(authHeader)) {
		const b64 = authHeader.split(" ")[1];
		return Buffer.from(b64, 'base64').toString().split(":");
	}
	return [];
}

function auth(ctx, next) {
	const [name, pwd] = basicAuth(ctx.request.headers.authorization);
	if (name && pwd) {
		const users = ctx.mongoClient.db('pets').collection('users');

		return users.findOne({ name, pwd })
			.then(doc => {
				if (!doc) {
					return Promise.reject(new Error(`ERR no match for ${ name }:${ pwd }`));
				}
				return next();
			})
			.catch(err => {
				return ctx.throw(401, "Unauthorized\n") // this sadly also catches thrown errors from the next handler
			});
	}
	return ctx.throw(401, "Unauthorized\n")
}

module.exports = auth;
