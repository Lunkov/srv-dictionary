#!/bin/bash

#####################################
# COLORS BLOCK
RED="\\033[1;31m"
BLUE="\\033[1;34m"
YELLOW="\033[1;33m"
GREEN="\033[0;32m"
RED="\033[41m\033[1;33m"
NC="\033[0m\n" # No Color

#####################################
# SHOW HELP
show_help() {
  echo "-----------------------------------------------------------------------"
  echo "                      Available commands                              -"
  echo "-----------------------------------------------------------------------"
  echo -e -n "$BLUE"
  echo "   > build - To build the Docker image"
  echo "   > install - To execute full install at once"
  echo "   > push - To push container"
  echo "   > pull - To pull container"
  echo "   > stop - To stop container"
  echo "   > start - To start container"
  echo "   > remove - Remove container"
  echo "   > help - Display this help"
  echo -e -n "$NC"
  echo "-----------------------------------------------------------------------"
}

if [ "$1" == "" ] || [ "$1" == "help" ] || [ "$1" == "?" ]; then
    show_help
    exit 1
fi

#####################################
# CHECK SUDO

function check_sudo () {
  if [ "$(id -u)" != "0" ]; then
    printf "$RED Sorry, you are not root.$NC"
    exit 1
  fi
}

#####################################
# Load environment variables
export $(cat .env | xargs)

#####################################
# LOG MESSAGE
log() {
  printf "$BLUE > $1 $NORMAL \n"
}

#####################################
# ERROR MESSAGE
error() {
  printf ""
  printf "$RED >>> ERROR - $1$NORMAL \n"
}

#####################################
# REMOVE CONTAINER
remove() {
  log "DELETE $CONTAINER_NAME"
  docker stop $CONTAINER_NAME
  docker rm --force $CONTAINER_NAME
  docker rmi --force $CONTAINER_NAME

  docker stop $DOCKER_SRC
  docker rm --force $DOCKER_SRC
  docker rmi --force $DOCKER_SRC
}

#####################################
# STOP CONTAINER
stop() {
  log "STOP $CONTAINER_NAME"
  docker stop $CONTAINER_NAME
}

#####################################
# START CONTAINER
start() {
  log "START $CONTAINER_NAME"
  docker start $CONTAINER_NAME
}

#####################################
# COMPILE CONTAINER
compile() {
  log "COMPILE $CONTAINER_NAME"
  DOCKERFILE=$(pwd)/docker/Dockerfile.builder
  
  docker build --force-rm=true --no-cache -f $DOCKERFILE -t $DOCKER_SRC .

  if [ $? -eq 0 ]; then
      log "OK"
  else
      error "FAIL"
      exit 1
  fi
}

#####################################
# BUILD CONTAINER
build() {
  log "BUILD $CONTAINER_NAME"
  DOCKERFILE=$(pwd)/docker/Dockerfile

  # Compile SRC
  compile

  if [ $? -eq 0 ]; then
      log "OK"
  else
      error "FAIL"
      exit 1
  fi
}

#####################################
# PUSH CONTAINER
push() {
  log "PUSH $CONTAINER_NAME"
  docker tag $CONTAINER_NAME $DOCKER_SRC
  docker push $DOCKER_SRC
}

#####################################
# PULL CONTAINER
pull() {
  log "PULL $CONTAINER_NAME"
  docker pull $DOCKER_SRC
}

#####################################
# CHECK IF THE CONTAINER IS STARTED
wait_for_running() {
  until [ "`docker inspect -f {{.State.Running}} $CONTAINER_NAME`"=="true" ]; do
      sleep 0.1;
  done;
}

#####################################
# INSTAL/UPGRADE CONTAINER
install() {
  SET_NETWORK=""
  if [ ! -z "$NETWORK_NAME" ]; then
    if [ ! "$(docker network ls | grep $NETWORK_NAME)" ]; then
      log "Creating '$NETWORK_NAME' network ..."
      docker network create $NETWORK_NAME
    else
      log "'$NETWORK_NAME' network exists."
    fi
    SET_NETWORK="--net $NETWORK_NAME"
  fi

  pull
  if [ ! "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=$CONTAINER_NAME)" ]; then
      # cleanup
      log "DELETE $CONTAINER_NAME"
      docker rm $CONTAINER_NAME
    fi
  else
    remove
  fi

  log "INSTALL $CONTAINER_NAME"

  docker run \
             -v /etc/localtime:/etc/localtime:ro \
             -v "$PWD"/etc:/app/etc/ \
             $SET_NETWORK \
             --name=$CONTAINER_NAME $CONTAINER_RESTART \
             -dit $DOCKER_SRC

  #sleep 5s

  #docker cp "$PWD"/etc/ $CONTAINER_NAME:/app/

  if [ $? -eq 0 ]; then
    log "OK"
  else
    error "FAIL"
    exit 1
  fi

}

log "START\n"
$1
log "FINISH\n"

