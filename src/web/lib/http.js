const axios = require('axios');

module.exports = axios.create({
  baseURL: 'http://localhost:37042/api/v1',
  timeout: 3000,
  //headers: {'X-Custom-Header': 'foobar'} // might be something for apm?
});
