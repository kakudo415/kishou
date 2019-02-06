#!/bin/sh
git pull && make
env PORT=10200 bin/kishow