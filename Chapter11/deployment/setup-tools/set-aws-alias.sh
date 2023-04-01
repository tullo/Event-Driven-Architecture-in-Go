#!/bin/bash

alias aws='docker run --rm -ti -v ~/.aws:/root/.aws -v $(pwd):/aws -v ~/.kube:/root/.kube amazon/aws-cli:2.11.8'

echo "---"
echo
echo "Usage: aws <command> help"
echo "aws --version"

# . set-aws-alias.sh
