#! /bin/bash -vex
opts='DOCKER_OPTS=" -H unix:///var/run/docker.sock -H tcp://0.0.0.0:4243"'

echo $opts > /etc/default/docker
