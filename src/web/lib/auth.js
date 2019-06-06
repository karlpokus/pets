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
		return ctx.users.findOne({ name, pwd })
			.then(doc => {
				if (!doc) {
					return Promise.reject(new Error(`ERR no match for ${ name }:${ pwd }`));
				}
				return next();
			})
			.catch(err => {
				return ctx.throw(401, "Unauthorized")
			});
	}
	return ctx.throw(401, "Unauthorized")
}

module.exports = auth;
