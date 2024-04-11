#!/bin/bash

# cd to the directory of this script so that this can be run from anywhere
script_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

# build abigen-with-interfaces docker image if it doesn't exist
if [[ "$(docker images -q abigen-with-interfaces 2> /dev/null)" == "" ]]; then
    docker build -t abigen-with-interfaces -f abigen-with-interfaces.Dockerfile $script_path
fi

# TODO: refactor this function.. it got really ugly with the dockerizing of abigen
function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo "creating bindings for $contract..."
    mkdir -p $binding_dir/${contract}
    contract_json="$contract_dir/out/${contract}.sol/${contract}.json"
    solc_abi=$(cat ${contract_json} | jq -r '.abi')
    solc_bin=$(cat ${contract_json} | jq -r '.bytecode.object')

    mkdir -p data
    echo ${solc_abi} >data/tmp.abi
    echo ${solc_bin} >data/tmp.bin

    rm -f $binding_dir/${contract}/binding.go
    docker run --rm -v $(realpath $binding_dir):/home/binding_dir -v .:/home/repo abigen-with-interfaces --bin=/home/repo/data/tmp.bin --abi=/home/repo/data/tmp.abi --pkg=contract${contract} --out=/home/binding_dir/${contract}/binding.go
    rm -rf data/tmp.abi data/tmp.bin
}

cd $script_path
cd ../../eoracle-middleware
# you might want to run forge clean if the contracts have changed
echo "Building eoracle-middleware"
forge build
cd $script_path

# No idea why but ordering of the contracts matters here... when I move them around sometimes bindings fail
avs_contracts="EORegistryCoordinator EOStakeRegistry OperatorStateRetriever BLSApkRegistry IBLSSignatureChecker ServiceManagerBase"
for contract in $avs_contracts; do
    create_binding ../../eoracle-middleware $contract ./bindings
    sleep 1
done

cd $script_path
cd ../../eoracle-middleware/lib/eigenlayer-contracts
# you might want to run forge clean if the contracts have changed
echo "Building eoracle-middleware"
forge build
cd $script_path

# No idea why but the ordering of the contracts matters, and for some orderings abigen fails...
el_contracts="DelegationManager ISlasher StrategyManager IStrategy IERC20 AVSDirectory"
for contract in $el_contracts; do
    create_binding ../../eoracle-middleware/lib/eigenlayer-contracts $contract ./bindings
    sleep 1
done

# Create bindings for the eoconfig contract in eochain
eo_contracts="EOConfig"
for contract in $eo_contracts; do
    create_binding ../../lb-contracts $contract ./bindings
    sleep 1
done