const nock = require('nock')

function run() {
  const scope = nock('http://localhost:8989').persist();

  scope
    .get('/api/v1/pets')
    .reply(200, "bixa,rex");

  scope
    .post('/api/v1/pet')
    .reply(200, "bixa created");
}

module.exports = {run}
