package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var mux sync.Mutex

func getPercentageEpoch() (float64, error) {
	startHeight, err := GetEpochStartHeight()
	if err != nil {
		return 0, err
	}
	latestHeight, err := getLatestBlockHeight()
	if err != nil {
		return 0, err
	}
	return float64(100) * float64(latestHeight-startHeight) / float64(EPOCH_LENGTH), nil

}

func AdaptStake(totalStake, seatPrice int) (error) {

	toPropose := int(float64(seatPrice) * (1 + AGGRESSIVENESS))

	if totalStake < toPropose {
		// low current locked stake. Stake some near token

		accountUnstaked, err := getAccountUnStakedBalance()
		if err != nil {
			log.Println("get account unstaked balance error", err)
			return err
		}
		log.Println("account unstaked balance:", accountUnstaked)

		toStake := min(toPropose-totalStake, accountUnstaked)

		if toStake > 0 {
			log.Println("try to stake:", toStake)
			err = runStake(POOL_ID, "stake", stringFromStake(toStake), DELEGATOR_ID)
			if err != nil {
				log.Println("stake error:", err)
			}
			log.Println("stake successfully")
			return err
		}

		return nil

	} else if totalStake > toPropose {
		accountStaked, err := getAccountStakedBalance()
		if err != nil {
			log.Println("get account staked balance error", err)
			return err
		}
		log.Println("account staked balance:", accountStaked)

		toUnstake := min(totalStake-toPropose, accountStaked)
		if toUnstake > 0 {
			log.Println("try to unstake:", toUnstake)
			err = runStake(POOL_ID, "unstake", stringFromStake(toUnstake), DELEGATOR_ID)
			if err != nil {
				log.Println("stake error:", err)
			}
			log.Println("unstake successfully")
			return err
		}

		return nil

	} else {
		log.Println("Nothing to do now.")
		return nil
	}

}

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(REAPT_TIME) * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("--------------------")

			// Get set Price
			currentSeatPrice, err := getSeatPrice("current")
			if err != nil {
				log.Println("get current set price error", err)
				continue
			}
			log.Println("current seat price:", currentSeatPrice)

			nextSeatPrice, err := getSeatPrice("next")
			if err != nil {
				log.Println("get next set price error", err)
				continue
			}
			log.Println("next seat price:", nextSeatPrice)

			proposalSeatPrice, err := getSeatPrice("proposal")
			if err != nil {
				log.Println("get proposal set price error", err)
				continue
			}
			log.Println("proposal seat price:", proposalSeatPrice)

			// get total stake
			contractTotalStake, err := getTotalStakedBalance()
			if err != nil {
				log.Println("get total staked balance error", err)
				continue
			}
			log.Println("contract total stake", contractTotalStake)

			// Get Percentage Epoch
			percentage, err := getPercentageEpoch()
			if err != nil {
				log.Println("get epoch percentage error", err)
				continue
			}
			log.Println("current epoch percentage:", percentage)

			if percentage <= 95 {
				log.Println("nothing to do now, epoch percentage:", percentage)
				continue
			}

			// Ping contract
			_, err = pingContract()
			if err != nil {
				log.Println("ping contract error", err)
				continue
			}
			log.Println("ping contract success.")

			mux.Lock()
			AdaptStake(contractTotalStake, proposalSeatPrice)
			mux.Unlock()

		case <-sigc:
			log.Println("System kill")
			os.Exit(0)
		}

	}
}
