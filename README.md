# Eoracle Operator Setup
This guide will walk you through the process of registering as an operator to Eoracle AVS.

## Prerequisites
1. **Registered Eigenlayer Operator Account:** Ensure you have a fully registered Eigenlayer operator account. If you don't have one, follow the steps in the [Eigenlayer User Guide](https://docs.eigenlayer.xyz/restaking-guides/restaking-user-guide) to create and fund your account.

## Software/Hardware Requirement 
* Operating System: linux amd x64
* vCPUs: 2
* Memory: 4GiB 
* Storage: 100GB
* EC2 Equivalent: m5.large
* Expected Network Utilization:
  * Total download bandwidth usage: 1 Mbps
  * Upload bandwidth usage: 1 Mbps 
* Open Ports:
  * 3000 Grafana dashboards
  * 9090 Prometheus 

## Operator Setup
### ​Prepare Local Eoracle data validator files
Clone this [repo](https://github.com/Eoracle/Eoracle-operator-setup) and execute the following commands
```bash
git clone https://github.com/Eoracle/Eoracle-operator-setup.git
cd Eoracle-operator-setup
cp data-validator/.example_env data-validator/.env
```
Copy `Eoracle-operator-setup/data-validator/.example_env` into `Eoracle-operator-setup/data-validator/.env`.  
Edit the `Eoracle-operator-setup/data-validator/.env` and update the values for your setup

### Generate a BLS pair (recommended)
The register process requires two sets of private keys: an ecdsa private key and a bls private key,  
We recommend creating a new BLS pair for security reasons.
If you want to create a new BLS pair, you can generate a new BLS pair that will be dedicated to Eoracle
```bash
./run.sh generate-bls-key
```

### Encrypt your private keys (recommended)
Encrypt your private keys. The encrypted private keys will be stored using the `EO_KEYSTORE_PATH` field.
This is the recommended approach. if you encrypt a pasted private key it will never be saved as is anywhere.
```bash
./run.sh encrypt <your ecdsa private key> <your bls private key>
```

### Work with plain text private keys (discouraged)
If you don't want to encrypt your private keys, update them in the `data-validator/.env` file  
This approach is highly discouraged. We recommend encrypting the private keys and never saving them anywhere on any machine.  
```bash
EO_BLS_PRIVATE_KEY=<your ecdsa private key>
EO_ECDSA_PRIVATE_KEY=<your bls private key>
```

### Register with Eoracle AVS
Operators need to have a minimum of 32 ETH delegated to them to opt-in to Eoracle. Execute the following command 
```bash
./run.sh register
```

The output should look like
```
Registering operator with address <your address> 
Opted into slashing with Service Manager at address 0x9546536FdAb1903e9c263923106eA5a4A44C4159
Registered Operator with Coordinator at address 0xd8eA2939cE17316b7CA2F86b121EB9c7c94c39c0 with tx hash 0x28b4b463762d917c514b0e0a5b6b8e65187a4ea43f2ab146672ba22d6b7aaa81
Awaiting finalization of registration
Awaiting finalization of registration
.
.
.
Awaiting finalization of registration
Successfully registered operator
```

### Troubleshooting the register command.
salt already spent - if you get the following error:
```
Failed to create RegisterOperator transaction execution reverted: AVSDirectory.registerOperatorToAVS: salt already spent
```
Please add EO_SALT=<salt_in_hex> field to your .env file and retry runnning register.  
(*) the EO_SALT should be in the following format EO_SALT=0x04 (even length hex number, and could be any number but must be even length)

### Checking the status of Eoracle operator AVS
The following command will print the status of the operator
```bash
./run.sh print-status
```

The output should look like
```
docker-entrypoint-oprcli.sh: Starting oprcli print-status 
{"level":"info","ts":1712824061.311895,"caller":"logging/zap_logger.go:49","msg":"Operator Status","status":"REGISTERED"}
{"level":"info","ts":1712824061.3466434,"caller":"logging/zap_logger.go:49","msg":"Operator stake update","stake":"1000","block number":1253026}
```
