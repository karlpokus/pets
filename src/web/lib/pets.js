const http = require('./http');

// TODO: put service endpoints in a conf
module.exports = {
  allPets: ctx => {
    return http.get('/pets')
  		.then(res => ctx.body = res.data)
  		.catch(err => ctx.throw(500, "server error"))
  },
  addPet: ctx => {
    return http.post('/pet', ctx.request.body)
  		.then(res => ctx.body = res.data)
  		.catch(err => ctx.throw(500, "server error"))
  }
};
