#!/usr/bin/env bash

# A bash script to scrape problem data from the Project Euler website and create the necessary files
if [ "$1" = "" ]
then
  echo "Usage: $0 [problem-number]"
  exit 1
fi

# chained since they kinda rely on the previous steps working
mkdir p$1 && cd p$1 && echo "# Problem $1" >> README.md && curl https://projecteuler.net/minimal=$1 >> README.md && go mod init euler.com/p$1 

