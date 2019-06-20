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
  stats: ctx => {
    let webStats = {
      name: "pets web",
      version: process.env.npm_package_version,
      uptime: process.uptime(),
      node_version: process.version, // node version
      os: process.platform,
      v8: process.versions.v8
    };

    return http.get('/api/v1/stats')
      .then(res => {
        ctx.set('Content-Type', 'application/json')
        return ctx.body = JSON.stringify([
          webStats,
          res.data
        ]);
      }).catch(err => {
        console.error(`Fetch pets stats failure ${ err }`);
        ctx.set('Content-Type', 'application/json')
        return ctx.body = JSON.stringify([
          webStats,
          {}
        ]);
      });
  },
  ping: ctx => {
    let msg = "";
    if (ctx.mongoClient.isConnected()) {
      msg = "db connection ok";
      ctx.response.status = 200;
    } else {
      msg = "db connection not ok";
      ctx.response.status = 500;
    }
    return ctx.body = msg;
  }
};
