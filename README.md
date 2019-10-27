# pets
A silly pet store to demonstrate SRE metrics with elastic apm. This repo used to contain deploy instructions with ansible for all apps and apm components. Now that everything is containerized on the master branch - we keep the old stuff in the native-systemd branch.

# usage

run everything in dev mode
```bash
$ docker-compose -f docker-compose-dev.yml -d up|down [-v]
```

run only pets service native
```bash
# mongo is required of course
$ docker run -d -p 127.0.0.1:27017:27017 --rm \
-v mongo:/data -v `pwd`/mongo-seed.js:/docker-entrypoint-initdb.d/mongo-seed.js \
--name mongo mongo:4.0.3
# cd src/pets
$ go run ./cmd/pets -n
```

# build

building images requires passing tests
```bash
# cd src/web - grabs version from package.json
$ ./build.sh
# cd src/pets - grabs version from file
$ ./build.sh
```

# api

```bash
# GET /api/v1/pets
$ curl -u <credentials> localhost:9012/api/v1/pets
# POST /api/v1/pet
$ curl -u <credentials> localhost:9012/api/v1/pet -d '{"name":"pet"}' -H "Content-Type:application/json"
```

# todos

- [x] do private network
- [x] deploy and run mongodb
- [x] mongodb metrics w metricbeat modules System, MongoDB
- [x] hostmetrics playbook
- [x] web-api
- [x] service
- [x] elastic apm on elastic cloud
- [x] add apm agents
- [x] deploy web and service
- [x] let web stats end point collect stats from pets
- [ ] [centralized logs](https://www.elastic.co/products/beats/filebeat)
- [x] [ping/heartbeat](https://www.elastic.co/products/beats/heartbeat)
- [ ] elastic apm rum
- [x] [kibana dashboards](https://www.elastic.co/guide/en/kibana/7.1/dashboard.html)
- [x] sre dashboard in kibana
- [x] graceful exits for http Servers and db connections
- [x] containerize apps
- [x] docker-composed dev
- [x] build scripts with tests
- [ ] run on k8s

# license
MIT
