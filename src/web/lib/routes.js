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
