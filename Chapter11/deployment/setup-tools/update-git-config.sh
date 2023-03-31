#!/bin/bash
shopt -s expand_aliases
alias deploytools='docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v ~/.mallbots:/root deploytools'

deploytools git config --global --add safe.directory /mallbots/deployment/.current/.terraform/modules/db
deploytools git config --global --add safe.directory /mallbots/deployment/.current/.terraform/modules/eks
deploytools git config --global --add safe.directory /mallbots/deployment/.current/.terraform/modules/eks.kms
deploytools git config --global --add safe.directory /mallbots/deployment/.current/.terraform/modules/vpc
deploytools git config --global --add safe.directory /mallbots/deployment/.current/.terraform/modules/vpc_cni_irsa
deploytools git config --global --add safe.directory /mallbots/deployment/.current/.terraform/modules/security_group

deploytools git config --global --list

# sudo rm ~/.mallbots/.gitconfig
