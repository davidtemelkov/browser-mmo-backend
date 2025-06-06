package fightsimulator

import (
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/gamecontent"
	"browser-mmo-backend/internal/monsters"
	"fmt"
	"math"
	"math/rand"
)

// TODO: Add armour amount and blockchance and use actual user's stats
type Fighter struct {
	ID            string
	Name          string
	Health        float32
	DamageMin     float32
	DamageMax     float32
	CritChance    float32
	MagicDamage   float32
	HitFirstIndex float32
	// BlockChance float32
	// ArmourAmount float32
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

	first, second := player, enemy
	if enemy.HitFirstIndex > player.HitFirstIndex {
		first, second = enemy, player
	}

	round := 1
	for player.Health > 0 && enemy.Health > 0 {
		fightLog += fmt.Sprintf("Round %d:\n", round)

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
			return fightLog, first.ID == player.ID
		}

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
			return fightLog, second.ID == player.ID // TODO: Remove this todo if fighting isn't broken frontend
		}

		round++
	}

	return fightLog, false
}

func NewFighterFromUser(playerData data.User) Fighter {
	return Fighter{
		ID:            playerData.ID,
		Name:          playerData.Name,
		Health:        float32(playerData.Constitution) + 100,
		DamageMin:     float32(playerData.DamageMin) + float32(playerData.Strength)/2,
		DamageMax:     float32(playerData.DamageMax) + float32(playerData.Strength),
		CritChance:    float32(playerData.Dexterity) * 0.01,
		MagicDamage:   float32(playerData.Intelligence),
		HitFirstIndex: float32(playerData.Dexterity) + float32(playerData.Lvl),
	}
}

func NewFighterFromMonster(monster monsters.GeneratedMonster) Fighter {
	return Fighter{
		ID:            monster.ID,
		Name:          monster.Name,
		Health:        float32(monster.Constitution) + 100,
		DamageMin:     float32(monster.Strength) / 2,
		DamageMax:     float32(monster.Strength),
		CritChance:    float32(monster.Dexterity) * 0.01,
		MagicDamage:   float32(monster.Intelligence),
		HitFirstIndex: float32(monster.Dexterity) + float32(monster.Lvl),
	}
}

func NewFighterFromBoss(boss gamecontent.Boss) Fighter {
	return Fighter{
		ID:            boss.ID,
		Name:          boss.Name,
		Health:        float32(boss.Constitution) + 100,
		DamageMin:     float32(boss.Strength / 2),
		DamageMax:     float32(boss.Strength),
		CritChance:    float32(boss.Dexterity) * 0.01,
		MagicDamage:   float32(boss.Intelligence),
		HitFirstIndex: float32(boss.Dexterity) + float32(boss.Lvl),
	}
}
