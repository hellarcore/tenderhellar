#!/bin/bash
set -exo pipefail

#- kvstore over socket, curl

# TODO: install everything

export PATH="$GOBIN:$PATH"
export TMHOME=$HOME/.tenderhellar_app

function init_validator() {
    rm -rf -- "$TMHOME"
    tenderhellar init validator

    # The default configuration sets a null indexer, but these tests require
    # indexing to be enabled. Rewrite the config file to set the "kv" indexer
    # before starting up the node.
    sed -i'' -e '/indexer = \["null"\]/c\
indexer = ["kv"]' "$TMHOME/config/config.toml"
}

function kvstore_over_socket() {
    init_validator
    echo "Starting kvstore_over_socket"
    abci-cli kvstore > /dev/null &
    pid_kvstore=$!
    tenderhellar start --mode validator > tenderhellar.log &
    pid_tenderhellar=$!
    sleep 5

    echo "running test"
    bash test/app/kvstore_test.sh "KVStore over Socket"

    kill -9 $pid_kvstore $pid_tenderhellar
}

# start tendermint first
function kvstore_over_socket_reorder() {
    init_validator
    echo "Starting kvstore_over_socket_reorder (ie. start tenderhellar first)"
    tenderhellar start --mode validator > tenderhellar.log &
    pid_tendermint=$!
    sleep 2
    abci-cli kvstore > /dev/null &
    pid_kvstore=$!
    sleep 5

    echo "running test"
    bash test/app/kvstore_test.sh "KVStore over Socket"

    kill -9 $pid_kvstore $pid_tenderhellar
}

function counter_over_socket() {
    rm -rf $TMHOME
    tenderhellar init validator
    echo "Starting counter_over_socket"
    abci-cli counter --serial > /dev/null &
    pid_counter=$!
    tenderhellar start --mode validator > tenderhellar.log &
    pid_tenderhellar=$!
    sleep 5

    echo "running test"
    bash test/app/counter_test.sh "Counter over Socket"

    kill -9 $pid_counter $pid_tenderhellar
}

function counter_over_grpc() {
    rm -rf $TMHOME
    tenderhellar init validator
    echo "Starting counter_over_grpc"
    abci-cli counter --serial --abci grpc > /dev/null &
    pid_counter=$!
    tenderhellar start --mode validator --abci grpc > tenderhellar.log &
    pid_tenderhellar=$!
    sleep 5

    echo "running test"
    bash test/app/counter_test.sh "Counter over GRPC"

    kill -9 $pid_counter $pid_tenderhellar
}

function counter_over_grpc_grpc() {
    rm -rf $TMHOME
    tenderhellar init validator
    echo "Starting counter_over_grpc_grpc (ie. with grpc broadcast_tx)"
    abci-cli counter --serial --abci grpc > /dev/null &
    pid_counter=$!
    sleep 1
    GRPC_PORT=36656
    tenderhellar start --mode validator --abci grpc --rpc.grpc_laddr tcp://localhost:$GRPC_PORT > tenderhellar.log &
    pid_tenderhellar=$!
    sleep 5

    echo "running test"
    GRPC_BROADCAST_TX=true bash test/app/counter_test.sh "Counter over GRPC via GRPC BroadcastTx"

    kill -9 $pid_counter $pid_tenderhellar
}

case "$1" in
    "kvstore_over_socket")
    kvstore_over_socket
    ;;
		"kvstore_over_socket_reorder")
    kvstore_over_socket_reorder
    ;;
*)
    echo "Running all"
    kvstore_over_socket
    echo ""
    kvstore_over_socket_reorder
    echo ""
esac
