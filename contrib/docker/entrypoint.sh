#!/bin/bash

# Fail if anything fails!
set -e
BITCOIN_DIR=/root/.bitcoin
BITCOIND_ARGS="-regtest -datadir=$BITCOIN_DIR -fallbackfee=0.0002 -rpcallowip=0.0.0.0/0 -server=1 -rpcbind=0.0.0.0 -txindex=$TXINDEX_ENABLED"
BITCOINCLI_ARGS="-regtest -datadir=$BITCOIN_DIR"

# Start bitcoind pointing into a temporary directory
mkdir -p $BITCOIN_DIR
bitcoind $BITCOIND_ARGS &

# Wait for it to work!
while ! bitcoin-cli $BITCOINCLI_ARGS help > /dev/null 2>&1; do sleep 1; done

# Create a default wallet
bitcoin-cli $BITCOINCLI_ARGS createwallet "default" || true
bitcoin-cli $BITCOINCLI_ARGS loadwallet "default" || true

while sleep 5; do
    bitcoin-cli $BITCOINCLI_ARGS -generate 1
done