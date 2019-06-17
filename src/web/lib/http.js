const axios = require('axios');
const petsHost = process.env.SERVICE_PETS_HOST;
const petsPort = process.env.SERVICE_PETS_PORT;

module.exports = axios.create({
  baseURL: `http://${ petsHost }:${ petsPort }`,
  timeout: 3000
});
