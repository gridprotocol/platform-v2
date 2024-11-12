#!/bin/bash
ps -ef | grep platform-v2 | grep -v 'color' | awk '{print $2}' | xargs kill -9
