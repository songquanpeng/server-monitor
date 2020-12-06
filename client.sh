#!/bin/bash
server=http://localhost:80
curl -X POST -d @<(nvidia-smi -q -x) --header "Content-Type:application/xml" $server/gpu
