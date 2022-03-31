package interaction

import (
	"fmt"
	"github.com/bhavanichandra/golang-learning/monsterslayer/actions"
)

type RoundData struct {
	Action              string
	PlayerAttackDamage  int
	PlayerHealValue     int
	MonsterAttackDamage int
	PlayerHealth        int
	MonsterHealth       int
}

func (roundData *RoundData) PrintRoundStatistics() {
	fmt.Println(roundData.LogInfo())
}

func (roundData RoundData) LogInfo() string {
	log := ""
	if roundData.Action == actions.ATTACK_ACTION {
		log += fmt.Sprintf("Player attacked monster for %v damange.\n", roundData.MonsterAttackDamage)
	} else if roundData.Action == actions.SPECIAL_ATTACK_ACTION {
		log += fmt.Sprintf("Player performed a strong attack monster for %v damange.\n", roundData.MonsterAttackDamage)
	} else {
		log += fmt.Sprintf("Player healed for %v .\n", roundData.PlayerHealValue)
	}
	log += fmt.Sprintf("Monster attacked player for %v damange.\n", roundData.PlayerAttackDamage)
	log += fmt.Sprintf("Player Health: %v\n", roundData.PlayerHealth)
	log += fmt.Sprintf("Monster Health: %v\n", roundData.MonsterHealth)
	return log
}

func (roundData *RoundData) LogMap(currentRound int) string {
	logEntry := map[string]string{
		"Round":                 fmt.Sprint(currentRound + 1),
		"Action":                roundData.Action,
		"Player Attack Damage":  fmt.Sprint(roundData.PlayerAttackDamage),
		"Monster Attack Damage": fmt.Sprint(roundData.MonsterAttackDamage),
		"Player Health":         fmt.Sprint(roundData.PlayerHealth),
		"Monster Health":        fmt.Sprint(roundData.MonsterHealth),
	}
	logLine := fmt.Sprintln(logEntry)
	return logLine
}
