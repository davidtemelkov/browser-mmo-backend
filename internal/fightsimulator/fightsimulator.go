package fightsimulator

import (
	"browser-mmo-backend/internal/data"
	"fmt"
	"math"
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

type FighterData struct {
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

// Simulate the fight and return a fight log and a boolean indicating if the player won
func Simulate(player *Fighter, enemy *Fighter) (string, bool) {
	var fightLog string

	// Initial magic damage
	fightLog += "Initial Magic Damage:\n"
	if player.magicDamage > enemy.magicDamage {
		enemy.health -= player.magicDamage
		fightLog += fmt.Sprintf("%s deals %d magic damage to %s. %s's health is now %d.\n",
			player.name, int(player.magicDamage), enemy.name, enemy.name, int(math.Round(float64(enemy.health))))
	} else {
		player.health -= enemy.magicDamage
		fightLog += fmt.Sprintf("%s deals %d magic damage to %s. %s's health is now %d.\n",
			enemy.name, int(enemy.magicDamage), player.name, player.name, int(math.Round(float64(player.health))))
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
			fightLog += "Critical hit! "
		}
		second.health -= damage
		fightLog += fmt.Sprintf("%s deals %d damage to %s. %s's health is now %d.\n",
			first.name, int(math.Round(float64(damage))), second.name, second.name, int(math.Round(float64(second.health))))

		if second.health <= 0 {
			fightLog += fmt.Sprintf("%s is defeated!\n", second.name)
			return fightLog, first == player
		}

		// Second attack
		damage = second.dealDamage()
		if second.criticalHit() {
			damage *= 2
			fightLog += "Critical hit! "
		}
		first.health -= damage
		fightLog += fmt.Sprintf("%s deals %d damage to %s. %s's health is now %d.\n",
			second.name, int(math.Round(float64(damage))), first.name, first.name, int(math.Round(float64(first.health))))

		if first.health <= 0 {
			fightLog += fmt.Sprintf("%s is defeated!\n", first.name)
			return fightLog, second == player
		}

		round++
	}

	// If somehow the loop exits without either being defeated
	return fightLog, false
}

func NewFighterFromUser(playerData data.User) *Fighter {
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

func NewFighterFromMonster(monster data.GeneratedMonster) *Fighter {
	return &Fighter{
		name:          monster.Name,
		health:        monster.Constitution + 100,
		damageMin:     monster.Strength / 2,
		damageMax:     monster.Strength,
		critChance:    monster.Dexterity * 0.01,
		magicDamage:   monster.Intelligence,
		hitFirstIndex: monster.Dexterity + float32(monster.Level),
	}
}
