const http = require('./http');

// TODO: put service endpoints in a conf
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
  }
};
