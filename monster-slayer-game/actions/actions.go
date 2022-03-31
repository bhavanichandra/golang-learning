package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN
	maxAttackValue := PLAYER_ATTACK_MAX

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN
		maxAttackValue = PLAYER_SPEACIAL_ATTACK_MAX
	}
	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= damageValue
	return damageValue
}

func HealPlayer() int {
	healValue := generateRandBetween(PLAYER_HEAL_MIN, PLAYER_HEAL_MAX)
	healthDiff := PLAYER_HEALTH - currentPlayerHealth
	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() int {
	minAttackValue := MONSTER_ATTACK_MIN
	maxAttackValue := MONSTER_ATTACK_MAX
	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= damageValue
	return damageValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + max
}
