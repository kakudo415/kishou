#!/bin/sh
cd $(cd $(dirname $0); pwd)
git pull
make
env PORT=10200 bin/kishow