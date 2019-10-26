#!/bin/bash

docker run -it --rm \
-u root \
-v ~/.m2:/root/.m2 \
-v $(pwd)/java:/usr/src/mymaven \
-w /usr/src/mymaven \
maven:3.5.3 \
mvn clean install -Dmaven.test.skip=true
