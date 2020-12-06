#!/bin/bash
server=http://your.domain.without.slash
curl -X POST -d @<(nvidia-smi -q -x) --header "Content-Type:application/xml" $server/gpu
