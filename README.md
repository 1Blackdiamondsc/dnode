# Wings Blockchain / Relay Part

[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)
[![Gitter](https://badges.gitter.im/WingsChat/community.svg)](https://gitter.im/WingsChat/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

**THIS IS VERY EARLY WORK IN PROGRESS, NOT FOR TESTNET/PRODUCTION USAGE**

Wings Blockchain peg zone implementation based on [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

This is work in progress, but still it already support next features:

* Proof Of Authority (PoA) Validators mechanism
* **N/2+1** confirmations model
* Multisignature based on PoA validators
* Managing of PoA validators state by PoA consensus
* Execution of messages (transactions) based on PoA consensus
* Issuing/destroying new coins based on PoA consensus
* 86400 blocks interval to confirm call execution under multisig

Motivation is allowing to moving tokens/currencies between different blockchains and Wings blockchain.

Additional information could be found in other, that presents part of Wings peg zones.

Other repositories related to peg zones could be found:

* [Ethereum Peg Zone](https://github.com/WingsDao/eth-peg-zone)

# Installation

Before we start you should have a correct 'GOPATH', 'GOROOT' environment variables.

To install fetch this repository:

    git clone git@github.com:WingsDao/blockchain-relay-layer.git

And let's build both daemon and cli:

    GO111MODULE=on go build cmd/wbd/main.go
    GO111MODULE=on go build cmd/wbcli/main.go

Both commands must execute fine, after it you can run both daemon and cli:

    GO111MODULE=on go run cmd/wbd/main.go
    GO111MODULE=on go run cmd/wbcli/main.go

## Install as binary

To install both cli and daemon as binaries you can use Makefile:

    make install

So after this command both `wbd` and `wbcli` will be available from console.

# Usage

First of all we need to create genesis configuration.

    wbd init --chain-id wings-testnet

Then let's create 4 accounts, one to store coins, the rest for PoA validators:

    wbcli keys add bank
    wbcli keys add validator1
    wbcli keys add validator2
    wbcli keys add validator3

Copy addresses and private keys from output, we will need them in the future.

As you see we create one account calling `bank` where we will be store all generated **WINGS** coins for start,
and then 3 accounts to make them PoA validators, we indeed 3 because by default it's minimum amount of PoA validators
to has.

Now let's add genesis account and initiate genesis poa validators:

    wbd add-genesis-account <bank-address> 10000,wings

    wbd add-poa-validator <validator-1-address>
    wbd add-poa-validator <validator-2-address>
    wbd add-poa-validator <validator-3-address>

Replace expressions in brackets with correct addresses, configure chain by Cosmos SDK documentation:

    wbcli config chain-id wings-testnet
    wbcli config output json
    wbcli config indent true
    wbcli config trust-node true

Now we are ready to launch testnet:

    wbd start

## Add/remove/replace validator by multisig

Before we start managing validators by PoA, let's remember that minimum amount of validators is 3, maximum is 11.

To add new validator use next command:

    wbcli tx validators ms-add-validator [validator_address] [eth_address] --validator-1

Where:

* [validator_address] - cosmos bench32 validator address
* [eth_address]       - validator ethereum address

To remove:

    wbcli tx validators ms-remove-validator [validator_address] --from validator1

To replace:

    wbcli tx validators ms-replace-validator [old_address] [new_address] [eth_address] --from validator-1

## Confirm multisignature call

To confirm multisignature call you need to extract call id from transaction execution output and confirm this call
by other validators:

    wbcli tx multisig confirm-call [call_id]

Once call submited under multisignature, there is `86400` blocks interval to confirm it by other validators, if call
not confirmed by that time, it will be marked as rejected.

To revoke confirmation from call:

    wbcli tx multisig revoke-confirm [call_id]

Once call reaches **N/2+1** amount of confirmations, message inside call will be executed.

# Docs

In progress.

# Tests

In progress.

# Contributors

Current project is under development and going to evolve together with other parts of Wings blockchain as
**Relay Layer** and Wings blockchain itself, anyway we have
planned things to:

* More Tests Coverage
* Refactoring
* Generate docs
* PoS government implementation instead of PoA

This project has the [following contributors](https://github.com/WingsDao/griffin-consensus-poc/graphs/contributors).

# License

Copyright © 2019 Wings Foundation

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the [GNU General Public License](https://github.com/WingsDAO/griffin-consensus-poc/tree/master/LICENSE) along with this program.  If not, see <http://www.gnu.org/licenses/>.