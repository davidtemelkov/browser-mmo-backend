package fightsimulator

import (
	"browser-mmo-backend/internal/data"
	"fmt"
	"math/rand"
)

type Fighter struct {
	name          string
	health        float32
	damageMin     float32
	damageMax     float32
	critChance    float32
	magicDamage   float32
	hitFirstIndex float32
}

func (f *Fighter) dealDamage() float32 {
	return f.damageMin + rand.Float32()*(f.damageMax-f.damageMin)
}

func (f *Fighter) criticalHit() bool {
	return rand.Float32() < f.critChance
}

// Get player and enemy, simulate the fight and return a fight log
func Simulate(playerData data.User, enemyData data.User) string {
	var fightLog string
	player := NewFighter(playerData)
	enemy := NewFighter(enemyData)

	// Initial magic damage
	fightLog += "Initial Magic Damage:\n"
	if player.magicDamage > enemy.magicDamage {
		enemy.health -= player.magicDamage
		fightLog += fmt.Sprintf("%s deals %f magic damage to %s. %s's health is now %f.\n", player.name, player.magicDamage, enemy.name, enemy.name, enemy.health)
	} else {
		player.health -= enemy.magicDamage
		fightLog += fmt.Sprintf("%s deals %f magic damage to %s. %s's health is now %f.\n", enemy.name, enemy.magicDamage, player.name, player.name, player.health)
	}

	// Determine who hits first
	first, second := player, enemy
	if enemy.hitFirstIndex > player.hitFirstIndex {
		first, second = enemy, player
	}

	// Fight rounds
	round := 1
	for player.health > 0 && enemy.health > 0 {
		fightLog += fmt.Sprintf("Round %d:\n", round)

		// First attack
		damage := first.dealDamage()
		if first.criticalHit() {
			damage *= 2
			fightLog += fmt.Sprintf("Critical hit! ")
		}
		second.health -= damage
		fightLog += fmt.Sprintf("%s deals %f damage to %s. %s's health is now %f.\n", first.name, damage, second.name, second.name, second.health)

		if second.health <= 0 {
			fightLog += fmt.Sprintf("%s is defeated!\n", second.name)
			break
		}

		// Second attack
		damage = second.dealDamage()
		if second.criticalHit() {
			damage *= 2
			fightLog += fmt.Sprintf("Critical hit! ")
		}
		first.health -= damage
		fightLog += fmt.Sprintf("%s deals %f damage to %s. %s's health is now %f.\n", second.name, damage, first.name, first.name, first.health)

		if first.health <= 0 {
			fightLog += fmt.Sprintf("%s is defeated!\n", first.name)
			break
		}

		round++
	}

	return fightLog
}

func NewFighter(playerData data.User) *Fighter {
	return &Fighter{
		name:          playerData.Name,
		health:        float32(playerData.Constitution) + 100,
		damageMin:     float32(playerData.Strength) / 2,
		damageMax:     float32(playerData.Strength),
		critChance:    float32(playerData.Dexterity) * 0.01,
		magicDamage:   float32(playerData.Intelligence),
		hitFirstIndex: float32(playerData.Dexterity) + float32(playerData.Level),
	}
}
