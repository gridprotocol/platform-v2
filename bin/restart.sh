#!/bin/bash

ps -ef | grep platform-v2 | grep -v 'color' | awk '{print $2}' | xargs kill -9

nohup ./platform-v2 daemon run --chain=$1 > log 2>&1 &

