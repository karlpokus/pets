# pets
A silly pet store to demonstrate SRE metrics with elastic apm.

Apps-, services and dbs will be deployed on cheap digital ocean droplets with ansible and elastic apm will be hosted on elastic cloud. Apps will be written in go and nodeJS.

Since elastic apm captures redis-, and mongodb traffic ootb we'll include them somehow.

# hosts
- api gateway
- services on a shared host
- redis
- mongodb
- elastic cloud (elastic apm, elasticsearch, kibana)

# metrics
- host system metrics
- [apm agent](https://www.elastic.co/guide/en/apm/agent/index.html)
- [mongodb stats](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-mongodb.html)
- [redis stats](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-redis.html)
- [go mem and gc](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-golang.html)

# usage
```bash
# web
GET /api/v1/pets
```

# deploy
```bash
cd deploy
# initial setup
$ ansible-playbook -i hosts user.yaml tools.yaml --limit <host>
# deploy metrics
$ ansible-playbook -i hosts metricbeat.yaml
```

# todos
- [x] do private network
- [x] deploy and run mongodb
- [x] mongodb metrics w metricbeat modules System, MongoDB
- [x] hostmetrics playbook
- [x] web-api
- [x] service
- [x] elastic apm on elastic cloud
- [ ] [centralized logs](https://www.elastic.co/products/beats/filebeat)
- [ ] [ping/heartbeat](https://www.elastic.co/products/beats/heartbeat)
- [ ] elastic apm rum
- [ ] [kibana dashboards](https://www.elastic.co/guide/en/kibana/7.1/dashboard.html)
