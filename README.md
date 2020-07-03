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
2020/07/03 23:43:04 --------------------
2020/07/03 23:43:06 current seat price: 106596
2020/07/03 23:43:08 next seat price: 112528
2020/07/03 23:43:10 proposal seat price: 106761
2020/07/03 23:43:11 contract total stake 157473
2020/07/03 23:43:13 current epoch percentage: 52.68
2020/07/03 23:43:13 nothing to do now, epoch percentage: 52.68

```

you can check the seat price and stake at the log.

## TODO List
- [ ] refactor messy code
- [ ] add promethus support
- [ ] use rpc instead of cmd?
- [ ] maybe support multi delegators?
- [ ] config adjust strategy
