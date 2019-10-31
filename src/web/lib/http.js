const axios = require('axios');
const petsHost = process.env.SERVICE_PETS_HOST || "localhost"; // for the testing mock
const petsPort = process.env.SERVICE_PETS_PORT || 8989;

module.exports = axios.create({
  baseURL: `http://${ petsHost }:${ petsPort }`,
  timeout: 3000
});
