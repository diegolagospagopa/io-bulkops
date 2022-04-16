#!/bin/bash

docker run -it --rm -p 3000:3000 -v $PWD/src/go/mock-server/:/home -w /home friendsofgo/killgrave -host 0.0.0.0
