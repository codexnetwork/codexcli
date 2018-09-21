---
title: "What's Been Added Since the Mainnet Launched"
description: "Since the launch of the EOS mainnet, EOS Canada has been working on eosc, a command line interface tool for EOS. It is a secure and simple-to-use interface for any EOS.IO blockchain."
categories: [blog]
---

Since the launch of the EOS mainnet, [EOS Canada](https://eoscanada.com) has been working on `eosc`, a command line interface tool for EOS. It is a secure and simple-to-use interface for any EOS.IO blockchain.

`eosc` is built using the [`eos-go`](https://www.eoscanada.com/en/tools/eos-go) library that EOS Canada has built, a fully-fledged toolkit for interacting with an EOS.IO blockchain in the modern and efficient Go programming language.

Building on previous `eosc` capabilities, like voting, securing your keys and using them to sign any transaction, here are the new features we are shipping today:

* Interactive prompts for reconfirmation when signing a transaction to mitigate against unauthorized signing requests performed by a malicious actor [(v0.7.0)](https://github.com/eoscanada/eosc/releases/tag/v0.7.0)
* Ability to delegate transaction signature to a wallet running outside of `eosc` by using the global flag --wallet-url, like `keosd` or `eosc vault serve` [(v0.7.3)](https://github.com/eoscanada/eosc/releases/tag/v0.7.3)
* Block Producer functions like `regproducer` `unregproducer` and `claimrewards` (under `eosc system`, since [v0.7.3](https://github.com/eoscanada/eosc/releases/tag/v0.7.3) and [v0.7.7](https://github.com/eoscanada/eosc/releases/tag/v0.7.3))
* A host of multisig functions, such as *propose*, *approve*, *unapprove*, *cancel*, and *exec* [(v0.7.3)](https://github.com/eoscanada/eosc/releases/tag/v0.7.3)
* Ability to bid on a namespace [(v0.7.5)](https://github.com/eoscanada/eosc/releases/tag/v0.7.5)
* Linkauth and Updateauth for setting up custom permissions for an account (under `eosc system`, since [(v0.7.5)](https://github.com/eoscanada/eosc/releases/tag/v0.7.5) and [(v0.7.7)](https://github.com/eoscanada/eosc/releases/tag/v0.7.7))
* `eosio.sudo` wrap using --sudo-wrap flag for easily testing and executing a sudo’d transaction by Block Producers [(v0.7.5)](https://github.com/eoscanada/eosc/releases/tag/v0.7.5)
* Ability to output a raw unsigned transaction to a file to greatly simplify the creation of multisig proposals, that and many `--offline-` options to perform cold storage signature of transactions, from an airtight computer [(v0.7.5)](https://github.com/eoscanada/eosc/releases/tag/v0.7.5)
* Chain freezing tool (`eosc tools chain-freeze`) to easily pause the chain at a certain block, useful for grabbing a specific chain state or table data knowing it is immutable (upgraded in [(v0.7.5)](https://github.com/eoscanada/eosc/releases/tag/v0.7.5))
* Added standard contract developer commands ([(v0.7.6)](https://github.com/eoscanada/eosc/releases/tag/v0.7.6))
* `eosc system setcontract` : Allows a smart contract developer to push WASM code along with the ABI (Application Binary Interface)
* `eosc system setcode` : Allows to push WASM code update without updating the ABI.
* `eosc system etabi`: Allows to push a new ABI without updating the code.
* Ability to atomically sell and transfer account ownership to another entity without the need for a trusted third party (with `eosc tools sell-account`, since [(v0.7.7)](https://github.com/eoscanada/eosc/releases/tag/v0.7.7))
* Bash and zsh shell auto-completion for increased usability ([(v0.7.8)](https://github.com/eoscanada/eosc/releases/tag/v0.7.8))
* Voting and vote tally tooling for the [`eosio.forum`](https://github.com/eoscanada/eosio.forum) referendum system, updated to latest contract revisions ([(v0.7.3)](https://github.com/eoscanada/eosc/releases/tag/v0.7.3) and [(v0.7.8)](https://github.com/eoscanada/eosc/releases/tag/v0.7.8))

As EOS.IO changes, and the needs of the user base change, EOS Canada will constantly be improving and adding new features to `eosc`. We encourage everyone to create Issues and Pull Requests on the [GitHub repository](https://github.com/eoscanada/eosc/) to let us know what new features you’d like to see. Feel free to join us anytime in our [Telegram channel](https://t.me/eoscanada) if you’re ever having any issues.

