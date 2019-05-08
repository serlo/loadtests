# Loadtests for serlo.org

All loadtests for serlo.org are based on K6 (https://k6.io/), Grafana (https://grafana.com/) and InfluxDB (https://www.influxdata.com/).

The project initially followed the tutorial at https://docs.k6.io/docs/influxdb-grafana.


## Quickstart for developing loadtests (using the example of athene2)
```
# 1 Clone loadtest project
git clone https://github.com/serlo/loadtests.git

# 2 Change into project directory 
cd loadtests

# 3 Run Grafana, InfluxDB and K6 via docker compose
docker-compose up -d

# 4 Run athene2 loadtest 
docker-compose run -v $PWD/scripts:/scripts k6 run /scripts/athene2.js

# 5 Access grafana at localhost, port 80
http://localhost

# 6 Import dashboard from file /dashboard/k6_load_testing_results.json
```



## Quickstart for running loadtests in Google Cloud Platform (using the example of athene2)
```
# 1 Start a VM-instance in Googlge Cloud with the following options: 
- choose the same region/zone as the serlo cluster
- min 1 vCPU (for informations regarding throughput cpas see https://cloud.google.com/compute/docs/networks-and-firewalls#egress_throughput_caps)
- boot image: Containter Optimized OS 
- firewall: allow HTTP Traffic

# 2 Connect to your new instance via SSH

# 3 Install docker-compose via docker (see https://cloud.google.com/community/tutorials/docker-compose-on-container-optimized-os)
docker run docker/compose:1.24.0 version
echo alias docker-compose="'"'docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "$PWD:$PWD" -w="$PWD" docker/compose:1.24.0'"'" >> ~/.bashrc
source ~/.bashrc

# 4 Clone loadtest project
git clone https://github.com/serlo/loadtests.git

# 5 Change into project directory 
cd loadtests

# 6 Run Grafana, InfluxDB and K6 via docker compose
docker-compose up -d

# 7 Access grafana at external-IP of VM, port 80
http://EXTERNAL_IP_VM

# 8 Import dashboard from file /dashboard/k6_load_testing_results.json

# 9 Run athene2 loadtest
docker-compose run -v $PWD/scripts:/scripts k6 run /scripts/athene2.js

# 10 Delete VM-instance
```

