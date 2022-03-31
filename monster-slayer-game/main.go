package main

import (
	"github.com/bhavanichandra/golang-learning/monsterslayer/actions"
	"github.com/bhavanichandra/golang-learning/monsterslayer/interaction"
)

var currentRound = 0

var playerAttackValue = 0
var playerHealValue = 0
var monsterAttackValue = 0

var rounds = []interaction.RoundData{}

func main() {
	startGame()
	winner := ""
	for winner == "" {
		winner = executeRound()
	}
	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isSpecialRound)
	newUserChoice := interaction.GetPlayerChoice(isSpecialRound)

	if newUserChoice == actions.ATTACK_ACTION {
		playerAttackValue = actions.AttackMonster(false)
	} else if newUserChoice == actions.HEAL_ACTION {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackValue = actions.AttackMonster(true)
	}
	monsterAttackValue = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:              newUserChoice,
		PlayerAttackDamage:  playerAttackValue,
		PlayerHealValue:     playerHealValue,
		MonsterAttackDamage: monsterAttackValue,
		PlayerHealth:        playerHealth,
		MonsterHealth:       monsterHealth,
	}
	rounds = append(rounds, roundData)
	roundData.PrintRoundStatistics()

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLog(&rounds)
	interaction.WriteLogMap(&rounds)
}
