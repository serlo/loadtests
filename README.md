#TEMP FILE!!!

based on https://k6.io/
and
https://docs.k6.io/docs/influxdb-grafana


#how to run locally
git clone TODO
cd TODO
docker-compose up -d
docker-compose run -v $PWD/scripts:/scripts k6 run /scripts/athene2-default.js

Import Dashboard:
https://grafana.com/dashboards/2587 

Access Grafana
http://localhost:3000

#run in gcp
startup VM in same zone!
install docker
sudo apt-get update

informations regarding traffic bandwith:
https://cloud.google.com/compute/docs/networks-and-firewalls#egress_throughput_caps
