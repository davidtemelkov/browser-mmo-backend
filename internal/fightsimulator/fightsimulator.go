package fightsimulator

import (
	"browser-mmo-backend/internal/data"
	"fmt"
	"math"
	"math/rand"
)

type Fighter struct {
	Name          string
	Health        float32
	DamageMin     float32
	DamageMax     float32
	CritChance    float32
	MagicDamage   float32
	HitFirstIndex float32
}

func (f *Fighter) dealDamage() float32 {
	return f.DamageMin + rand.Float32()*(f.DamageMax-f.DamageMin)
}

func (f *Fighter) criticalHit() bool {
	return rand.Float32() < f.CritChance
}

// Simulate the fight and return a fight log and a boolean indicating if the player won
func Simulate(player Fighter, enemy Fighter) (string, bool) {
	var fightLog string

	// Initial magic damage
	fightLog += "Initial Magic Damage:\n"
	if player.MagicDamage > enemy.MagicDamage {
		enemy.Health -= player.MagicDamage
		fightLog += fmt.Sprintf("%s deals %d magic damage to %s. %s's health is now %d.\n",
			player.Name, int(player.MagicDamage), enemy.Name, enemy.Name, int(math.Round(float64(enemy.Health))))
	} else {
		player.Health -= enemy.MagicDamage
		fightLog += fmt.Sprintf("%s deals %d magic damage to %s. %s's health is now %d.\n",
			enemy.Name, int(enemy.MagicDamage), player.Name, player.Name, int(math.Round(float64(player.Health))))
	}

	// Determine who hits first
	first, second := player, enemy
	if enemy.HitFirstIndex > player.HitFirstIndex {
		first, second = enemy, player
	}

	// Fight rounds
	round := 1
	for player.Health > 0 && enemy.Health > 0 {
		fightLog += fmt.Sprintf("Round %d:\n", round)

		// First attack
		damage := first.dealDamage()
		if first.criticalHit() {
			damage *= 2
			fightLog += "Critical hit! "
		}
		second.Health -= damage
		fightLog += fmt.Sprintf("%s deals %d damage to %s. %s's health is now %d.\n",
			first.Name, int(math.Round(float64(damage))), second.Name, second.Name, int(math.Round(float64(second.Health))))

		if second.Health <= 0 {
			fightLog += fmt.Sprintf("%s is defeated!\n", second.Name)
			// TODO: This would be more secure with IDs instead of Name
			// Someone can name themselves Slime and break the logic
			return fightLog, first.Name == player.Name
		}

		// Second attack
		damage = second.dealDamage()
		if second.criticalHit() {
			damage *= 2
			fightLog += "Critical hit! "
		}
		first.Health -= damage
		fightLog += fmt.Sprintf("%s deals %d damage to %s. %s's health is now %d.\n",
			second.Name, int(math.Round(float64(damage))), first.Name, first.Name, int(math.Round(float64(first.Health))))

		if first.Health <= 0 {
			fightLog += fmt.Sprintf("%s is defeated!\n", first.Name)
			return fightLog, second.Name == player.Name
		}

		round++
	}

	// If somehow the loop exits without either being defeated
	return fightLog, false
}

func NewFighterFromUser(playerData data.User) Fighter {
	return Fighter{
		Name:          playerData.Name,
		Health:        float32(playerData.Constitution) + 100,
		DamageMin:     float32(playerData.Strength) / 2,
		DamageMax:     float32(playerData.Strength),
		CritChance:    float32(playerData.Dexterity) * 0.01,
		MagicDamage:   float32(playerData.Intelligence),
		HitFirstIndex: float32(playerData.Dexterity) + float32(playerData.Level),
	}
}

func NewFighterFromMonster(monster data.GeneratedMonster) Fighter {
	return Fighter{
		Name:          monster.Name,
		Health:        monster.Constitution + 100,
		DamageMin:     monster.Strength / 2,
		DamageMax:     monster.Strength,
		CritChance:    monster.Dexterity * 0.01,
		MagicDamage:   monster.Intelligence,
		HitFirstIndex: monster.Dexterity + float32(monster.Level),
	}
}
