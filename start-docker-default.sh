docker-machine start default
docker-machine env default
eval $(docker-machine env default)
env | grep DOCKER
docker-machine ls