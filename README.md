#Loadtests for serlo.org

All loadtests for serlo.org are based on K6 (https://k6.io/), Grafana (https://grafana.com/) and InfluxDB (https://www.influxdata.com/).

The project initially followed the tutorial at https://docs.k6.io/docs/influxdb-grafana.


##Quickstart for developing loadtests (using the example of athene2)
```
# 1 clone loaddtest project
git clone https://github.com/serlo/loadtests.git

# 2 change into project directory 
cd loadtests

# 3 run Grafana, InfluxDB and K6 via docker compose
docker-compose up -d

# 4 run athene2 loadtest 
docker-compose run -v $PWD/scripts:/scripts k6 run /scripts/athene2.js

# 5 access grafana at localhost
http://localhost:3000

# 6 import dashboard from file /dashboard/k6_load_testing_results.json
```



##Quickstart for running loadtests in Google Cloud Platform (using the example of athene2)
TODO
startup VM in same zone!
install docker
sudo apt-get update

informations regarding traffic bandwith:
https://cloud.google.com/compute/docs/networks-and-firewalls#egress_throughput_caps
