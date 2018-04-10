# Blockchain in Go

A blockchain implementation in Go, as described in these articles:

1. [Basic Prototype](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)
2. [Proof-of-Work](https://jeiwan.cc/posts/building-blockchain-in-go-part-2/)
3. [Persistence and CLI](https://jeiwan.cc/posts/building-blockchain-in-go-part-3/)
4. [Transactions 1](https://jeiwan.cc/posts/building-blockchain-in-go-part-4/)
5. [Addresses](https://jeiwan.cc/posts/building-blockchain-in-go-part-5/)
6. [Transactions 2](https://jeiwan.cc/posts/building-blockchain-in-go-part-6/)
7. [Network](https://jeiwan.cc/posts/building-blockchain-in-go-part-7/)

Result
Let’s play the scenario we defined earlier.

NODE 3000
Create a wallet and a new blockchain:

```bash
source create_blockchain.sh
Your new address: 14LFjRUBPXz4dgarGYqcsmCkydfFKMfeWL
0000261f43701f4304b027dbd1113c7bcace53e7a5425e72136ab6db0b04b2ce

NODE 3001
init wallet node(Create 3 new wallets) WALLET_1, WALLET_2, WALLET_3
```bash
source init_wallet_node.sh
```

NODE 3000
Send some coins to the wallet addresses:
```bash
source send_coin.sh
```
start main node
```bash
source start_main_node.sh
```

NODE 3001
start wallet node
```bash
source start_wallet_node.sh
    Starting node 3001
    Received version command
    Received inv command
    Recevied inventory with 3 block
    Received block command
    Recevied a new block!
    Added block 0000345f8b77688a19e4db8657607eaab4920e93c9265cf64e869499d16e5675
    Received block command
    Recevied a new block!
    Added block 000037d05e444edeaa4130e936cb8f38aaf0b6a575715721cc38c733287be244
    Received block command
    Recevied a new block!
    Added block 00000d01fbd16cd98f15d81733b9b28bd61f06cab413abec484b2c0f13f0d1d1
```

    in NODE 3000, we can see:
    ```
    Received version command
    Received getblocks command
    Received getdata command
    Received getdata command
    Received getdata command
    ```

NODE 3001
after sync block data, list balance of WALLET_1, WALLET_2, WALLET_2, MAIN_NODE
```bash
source listbalance_wallet.sh.sh
    Balance of '17MsiCPJWZ92fyBVQ8iJXZAvSzxQg1mCT': 10
    Balance of '1Bx1Q8bSphw21AB7RBEDfL4BWVrpHPXWRV': 10
    Balance of '1Car8iAHs4UeKXUmhLqRJabh5UtpvXJPz3': 0
    Balance of '15dsN57ANqi6cHQdhW4zzBdrZk21Lu4566': 10

```

NODE 3002
init miner node(generate a new wallet-MINER_WALLET)
```bash
source init_miner_node.sh
```
start miner node
```bash
source start_miner_node.sh
```

NODE 3001
send coins
```bash
source send_coin_3001.sh
```
