const http = require('./http');

module.exports = {
  allPets: ctx => {
    return http.get('/pets') // should be in a config
  		.then(res => ctx.body = res.data)
  		.catch(err => ctx.throw(500, "server error"))
  }
};
