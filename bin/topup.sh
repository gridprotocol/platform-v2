#!/bin/bash

echo "transfer credit to user"
./platform-v2 topup -a 0x82379862a857C98aB391Fa7F66957AfDE97EF528 -v 400000 -c $1

echo "transfer gtoken to user"
./platform-v2 topup2 -a 0x82379862a857C98aB391Fa7F66957AfDE97EF528 -v 400000 -c $1