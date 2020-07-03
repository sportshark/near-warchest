# near warchest bot

This is a tool to manage near staked tokens, refer to [near stakewars challenge004](https://github.com/nearprotocol/stakewars/blob/master/challenges/challenge004.md)

## config

Change the config value in config.go to your's:

```
const (
	POOL_ID        = "c1.hashquark"   // CHANGE TO YOUR POOL ID
	DELEGATOR_ID   = "hashquark"      // CHANGE TO YOU MASTER ACCOUNT
	AGGRESSIVENESS = 0.4
	REAPT_TIME     = 300
	EPOCH_LENGTH   = 10000 // for betanet
	ENDPOINT       = "https://rpc.betanet.near.org"
)
```


## build


```
go build ./
```

the output executable will be named `near-warchest`

## run


Put the executable in the same server and directory as you participate stakewars game, such as ``~/stakewars`.

Make sure `export NODE_ENV=betanet`

Then run with `./near-warchest`.

# result

outputs like below:

```
2020/07/03 21:49:46 --------------------
2020/07/03 21:49:49 current seat price: 108202
2020/07/03 21:49:51 next seat price: 106596
2020/07/03 21:49:53 proposal seat price: 112481
2020/07/03 21:49:54 current epoch percentage: 91.69
2020/07/03 21:49:54 nothing to do now, epoch percentage: 91.69
2020/07/03 21:54:46 --------------------
2020/07/03 21:54:49 current seat price: 108202
2020/07/03 21:54:51 next seat price: 106596
2020/07/03 21:54:53 proposal seat price: 112481
2020/07/03 21:54:54 current epoch percentage: 94.54
2020/07/03 21:54:54 nothing to do now, epoch percentage: 94.54
2020/07/03 21:59:46 --------------------
2020/07/03 21:59:49 current seat price: 108202
2020/07/03 21:59:51 next seat price: 106596
2020/07/03 21:59:53 proposal seat price: 112481
2020/07/03 21:59:54 current epoch percentage: 97.39
2020/07/03 22:00:02 ping contract success.
2020/07/03 22:00:03 contract total stake 157487
2020/07/03 22:00:05 account staked balance: 86978
2020/07/03 22:00:05 try to unstake: 14
2020/07/03 22:00:13 unstake successfully
2020/07/03 22:04:46 --------------------

```


## TODO List
- [ ] refactor messy code
- [ ] add promethus support
- [ ] use rpc instead of cmd?
- [ ] maybe support multi delegators?
- [ ] config adjust strategy
