$ docker-machine start default
$ docker-machine env default
$ eval $(docker-machine env default)
$ env | grep DOCKER

docker run -p 8080:8080 ladislavlisy/payroll-maestro-080318:master

curl http://192.168.99.100:8080 
