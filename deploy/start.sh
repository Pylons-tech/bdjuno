#!/bin/bash

#sed -i 's/search_string/replace_string/' /bdjuno/config.yaml
#CMD [ "/usr/bin/bdjuno", "parse", "--home", "/bdjuno/" ]
/usr/bin/bdjuno --home /bdjuno hasura-actions --rpc http://rpc.testnet.pylons.tech:26657 --grpc http://grpc.testnet.pylons.tech:9090