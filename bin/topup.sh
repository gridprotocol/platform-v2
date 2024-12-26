#!/bin/bash

echo "transfer credit to user"
./platform-v2 topup -a 0x82379862a857C98aB391Fa7F66957AfDE97EF528 -v 1000000000000000000 -c $1

echo "transfer gtoken to provider"
./platform-v2 topup2 -a 0xEf95c72C836605203F7f66788E450Af2a4141957 -v 1000000000000000000 -c $1