package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

const (
	CURRENT_SEAT_PRICE_CMD   = "near validators current | awk '/price/ {print substr($6, 1, length($6)-2)}'"
	NEXT_SEAT_PRICE_CMD      = "near validators next | awk '/price/ {print substr($7, 1, length($7)-2)}'"
	PROPOSALS_SEAT_PRICE_CMD = "near proposals | awk '/price =/ {print substr($15, 1, length($15)-1)}'"

	STAKE_CMD = "near call %s %s '{\"amount\": \"%s\"}' --accountId %s"

	GET_TOTAL_STAKED_BALANCE     = "near view %s get_total_staked_balance '{\"account_id\": \"%s\"}'"
	GET_ACCOUNT_STAKED_BALANCE   = "near view %s get_account_staked_balance '{\"account_id\": \"%s\"}'"
	GET_ACCOUNT_UNSTAKED_BALANCE = "near view %s get_account_unstaked_balance '{\"account_id\": \"%s\"}'"

	PING_CMD = "near call %s ping '{}' --accountId %s"
)

func runCmd(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to execute command: %s", cmd))
	}
	return string(out), nil
}

func pingContract() (bool, error) {
	command := fmt.Sprintf(PING_CMD, POOL_ID, DELEGATOR_ID)
	_, err := runCmd(command)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getSeatPrice(epoch string) (int, error) {
	var (
		price string
		err   error
	)

	if epoch == "current" {
		price, err = runCmd(CURRENT_SEAT_PRICE_CMD)
	} else if epoch == "next" {
		price, err = runCmd(NEXT_SEAT_PRICE_CMD)
	} else {
		price, err = runCmd(PROPOSALS_SEAT_PRICE_CMD)
	}

	if err != nil {
		return 0, err
	}

	return intFromString(price), nil
}

func getTotalStakedBalance() (int, error) {
	command := fmt.Sprintf(GET_TOTAL_STAKED_BALANCE, POOL_ID, DELEGATOR_ID)
	r, err := runCmd(command)
	if err != nil {
		return 0, err
	}
	return stakeFromNearView(r), nil
}

func getAccountStakedBalance() (int, error) {
	command := fmt.Sprintf(GET_ACCOUNT_STAKED_BALANCE, POOL_ID, DELEGATOR_ID)
	r, err := runCmd(command)
	if err != nil {
		return 0, err
	}
	return stakeFromNearView(r), nil
}

func getAccountUnStakedBalance() (int, error) {
	command := fmt.Sprintf(GET_ACCOUNT_UNSTAKED_BALANCE, POOL_ID, DELEGATOR_ID)
	r, err := runCmd(command)
	if err != nil {
		return 0, err
	}
	return stakeFromNearView(r), nil

}

func runStake(poolId, method, amount, delegatorId string) error {
	_, err := runCmd(fmt.Sprintf(STAKE_CMD, poolId, method, amount, delegatorId))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
