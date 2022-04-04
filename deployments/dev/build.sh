#!/bin/bash

#################################################
# Clone the repository project from Github
#
# ARGUMENTS:
#   PROJECT: the project name
#   REPOSITORY: destination folder
#   BRANCH: branch code to work on
#
# NOTE: 
#   1.- You must provide the authorized ssh keys: id_rsa, id_rsa.pub
#   2.- Save them under src/.ssh folder
#   3.- Use make tool to build the image  
#################################################

# ARGUMENTS
PROJECT=capstone
REPOSITORY=/go/src/${PROJECT}
BRANCH=development

if [ -d $REPOSITORY ]
then
    echo "WARNING: sources ${REPOSITORY} already exists"
else
    # FIX: Host key verification failed.
    ssh-keyscan github.com > /home/${PROJECT}/.ssh/known_hosts
    git clone \
        --verbose --branch ${BRANCH} \
        git@github.com:marcos-wz/${PROJECT}.git ${REPOSITORY}
fi
