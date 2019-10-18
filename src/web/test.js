const test = require('supertest');
const http = require('http');
const assert = require('assert');
const { app } = require('./server');
const dbmock = require('./lib/dbmock');
const httpmock = require('./lib/httpmock');

const cb = http.createServer(app.callback())
app.context.mongoClient = dbmock;
httpmock.run();

const testTable = [
  {
    method: 'get',
    path: '/api/v1/ping',
    auth: {
      user: "",
      pwd: ""
    },
    status: 200,
    text: 'db connection ok'
  },
  {
    method: 'get',
    path: '/api/v1/pets',
    auth: {
      user: "buck",
      pwd: "nasty"
    },
    status: 200,
    text: 'bixa,rex'
  },
  {
    method: 'post',
    path: '/api/v1/pet',
    auth: {
      user: "buck",
      pwd: "nasty"
    },
    status: 200,
    text: 'bixa created'
  }
];

testTable.forEach(tt => {
  test(cb)[tt.method](tt.path)
    .auth(tt.auth.user, tt.auth.pwd)
    .expect(tt.status)
    .then(res => {
      assert.strictEqual(res.text, tt.text);
    })
    .catch(err => {
      console.error(`${ tt.path } failed: ${ err }`);
			process.exit(1);
    });
});
