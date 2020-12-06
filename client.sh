#!/bin/bash
server=http://localhost
curl -X POST -d @<(nvidia-smi -q -x) --header "Content-Type:application/xml" $server/post
