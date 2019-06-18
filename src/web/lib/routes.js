const http = require('./http');

module.exports = {
  allPets: ctx => {
    return http.get('/api/v1/pets')
  		.then(res => ctx.body = res.data)
  		.catch(err => ctx.throw(500, "server error"))
  },
  addPet: ctx => {
    return http.post('/api/v1/pet', ctx.request.body)
  		.then(res => ctx.body = res.data)
  		.catch(err => ctx.throw(500, "server error"))
  },
  stat: ctx => {
    return ctx.body = JSON.stringify({
      name: "pets web",
      version: process.env.npm_package_version,
      uptime: process.uptime(),
      node_version: process.version, // node version
      os: process.platform,
      v8: process.versions.v8
    });
  }
};
