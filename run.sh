#!/bin/sh
git pull && make && echo "BUILD COMPLETE"
env PORT=10200 bin/kishow