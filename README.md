# Photon 

EVM-compatible chain secured by the Sirius consensus algorithm.
With TechPay Photon, you can experience the [fastest blockchain](https://www.thankyoucoin.io/fastest-blockchain) transactions.

## Building the source

Building `thankyoucoin` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make thankyoucoin
```
The build output is ```build/thankyoucoin``` executable.

## Running `thankyoucoin`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `thankyoucoin` instance.

### Launching a network

Launching `thankyoucoin` for a network:

```shell
$ thankyoucoin --genesis /path/to/genesis.g
```

### Configuration

As an alternative to passing the numerous flags to the `thankyoucoin` binary, you can also pass a
configuration file via:

```shell
$ thankyoucoin --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ thankyoucoin --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `thankyoucoin validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ thankyoucoin --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`thankyoucoin` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ thankyoucoin --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ thankyoucoin --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ thankyoucoin --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ thankyoucoin --datadir /path/to/datadir account new # create new account
$ thankyoucoin --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Sirius has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```
ok  	github.com/kalibroida/ThankyouCoin_Node/app	0.033s
?   	github.com/kalibroida/ThankyouCoin_Node/cmd/cmdtest	[no test files]
ok  	github.com/kalibroida/ThankyouCoin_Node/cmd/thankyoucoin	13.890s
?   	github.com/kalibroida/ThankyouCoin_Node/cmd/thankyoucoin/metrics	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/cmd/thankyoucoin/tracing	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/crypto	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/debug	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/ethapi	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/eventcheck	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/eventcheck/basiccheck	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/eventcheck/gaspowercheck	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/eventcheck/heavycheck	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/eventcheck/parentscheck	[no test files]
ok  	github.com/kalibroida/ThankyouCoin_Node/evmcore	6.322s
?   	github.com/kalibroida/ThankyouCoin_Node/gossip	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/gossip/emitter	[no test files]
ok  	github.com/kalibroida/ThankyouCoin_Node/gossip/filters	1.250s
?   	github.com/kalibroida/ThankyouCoin_Node/gossip/gasprice	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/gossip/occuredtxs	[no test files]
?   	github.com/kalibroida/ThankyouCoin_Node/gossip/piecefunc	[no test files]
ok  	github.com/kalibroida/ThankyouCoin_Node/integration	21.640s
```

Also it is tested with [fuzzing](./FUZZING.md).


### Operating a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ thankyoucoin --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ thankyoucoin --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ thankyoucoin --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ thankyoucoin --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the Photon processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```
