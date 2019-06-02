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
- [mongodb stats](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-mongodb.html)
- [redis stats](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-redis.html)
- [go mem and gc](https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-module-golang.html)

# deploy
```bash
cd deploy
# initial setup
$ ansible-playbook -i hosts user.yaml --limit <host>
# deploy some developer friendly tools
$ ansible-playbook -i hosts tools.yaml --limit <host>
```

# todos
- [x] do private network
- [x] deploy and run mongodb
- [x] mongodb metrics w metricbeat modules System, MongoDB
- [ ] api
- [ ] service
- [x] elastic apm on elastic cloud
- [ ] [centralized logs](https://www.elastic.co/products/beats/filebeat)
- [ ] [ping/heartbeat](https://www.elastic.co/products/beats/heartbeat)
- [ ] elastic apm rum
