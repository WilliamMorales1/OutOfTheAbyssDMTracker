package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Character represents a player character in the fight
type Character struct {
	Name        string
	MaxHP       int
	CurrentHP   int
	AC          int    // Armor Class
	AttackBonus int    // Bonus to attack rolls
	DamageDie   string // e.g., "1d8+2"
}

// Monster represents a monster in the fight
type Monster struct {
	Name        string
	MaxHP       int
	CurrentHP   int
	AC          int
	AttackBonus int
	DamageDie   string
}

func rollDie(sides int) int {
	return rand.Intn(sides) + 1
}

// parseDamageDie parses a damage die string (e.g., "1d8+2") into number of dice and bonus
func parseDamageDie(damageDie string) (int, int, int) {
	var numDice, dieSides, bonus int
	_, err := fmt.Sscanf(damageDie, "%dd%d+%d", &numDice, &dieSides, &bonus)
	if err != nil {
		// Try without bonus
		_, err := fmt.Sscanf(damageDie, "%dd%d", &numDice, &dieSides)
		if err != nil {
			return 0, 0, 0 // Should handle error more gracefully in a real app
		}
	}
	return numDice, dieSides, bonus
}

// calculateDamage calculates the total damage based on a damage die string
func calculateDamage(damageDie string) int {
	numDice, dieSides, bonus := parseDamageDie(damageDie)

	damage := 0
	for range numDice {
		damage += rollDie(dieSides)
	}
	return damage + bonus
}

// makeAttack determines if an attack hits
func makeAttack(attackerAttackBonus int, targetAC int) bool {
	attackRoll := rollDie(20) // d20 for attack roll
	fmt.Printf("Attack roll: %d + %d (bonus) = %d\n", attackRoll, attackerAttackBonus, attackRoll+attackerAttackBonus)
	return (attackRoll + attackerAttackBonus) >= targetAC
}

// rollInitiative rolls for initiative (d20)
func rollInitiative() int {
	return rollDie(20)
}

func main() {
	fmt.Println("D&D Group Fight Simulator")

	// --- Simulation Setup ---
	// 4 Player Characters
	players := []*Character{
		{
			Name: "Hero 1", MaxHP: 40, CurrentHP: 40, AC: 16, AttackBonus: 6, DamageDie: "1d8+4"},
		{
			Name: "Hero 2", MaxHP: 35, CurrentHP: 35, AC: 14, AttackBonus: 5, DamageDie: "1d6+3"},
		{
			Name: "Hero 3", MaxHP: 30, CurrentHP: 30, AC: 15, AttackBonus: 7, DamageDie: "2d6+2"},
		{
			Name: "Hero 4", MaxHP: 45, CurrentHP: 45, AC: 17, AttackBonus: 5, DamageDie: "1d10+3"},
	}

	// Variable Monster Group (e.g., 3 Goblins)
	monsters := []*Monster{
		{Name: "Goblin 1", MaxHP: 10, CurrentHP: 10, AC: 13, AttackBonus: 4, DamageDie: "1d6+2"},
		{Name: "Goblin 2", MaxHP: 10, CurrentHP: 10, AC: 13, AttackBonus: 4, DamageDie: "1d6+2"},
		{Name: "Goblin 3", MaxHP: 10, CurrentHP: 10, AC: 13, AttackBonus: 4, DamageDie: "1d6+2"},
	}

	fmt.Println("\n--- Combat Starts ---")

	turnOrder := []any{} // Can hold *Character or *Monster

	// Roll initiative for players
	for _, p := range players {
		initiative := rollInitiative()
		fmt.Printf("%s (Player) rolls initiative: %d\n", p.Name, initiative)
		turnOrder = append(turnOrder, struct {
			combatant  any
			initiative int
		}{p, initiative})
	}

	// Roll initiative for monsters
	for _, m := range monsters {
		initiative := rollInitiative()
		fmt.Printf("%s (Monster) rolls initiative: %d\n", m.Name, initiative)
		turnOrder = append(turnOrder, struct {
			combatant  any
			initiative int
		}{m, initiative})
	}

	// Sort turn order (descending initiative)
	sort.Slice(turnOrder, func(i, j int) bool {
		return turnOrder[i].(struct {
			combatant  any
			initiative int
		}).initiative > turnOrder[j].(struct {
			combatant  any
			initiative int
		}).initiative
	})

	// Simulation Loop
	round := 0
	for {
		round++
		fmt.Printf("\n--- Round %d ---\n", round)

		for _, entry := range turnOrder {
			combatant := entry.(struct {
				combatant  any
				initiative int
			}).combatant

			switch c := combatant.(type) {
			case *Character:
				if c.CurrentHP <= 0 {
					continue // Skip defeated character
				}
				// Character's turn
				fmt.Printf("%s's turn (Player)\n", c.Name)

				// Find a living monster to attack
				livingMonsters := []*Monster{}
				for _, m := range monsters {
					if m.CurrentHP > 0 {
						livingMonsters = append(livingMonsters, m)
					}
				}

				if len(livingMonsters) == 0 {
					break // All monsters defeated
				}

				// Target a random living monster
				target := livingMonsters[rand.Intn(len(livingMonsters))]
				fmt.Printf("%s attacks %s!\n", c.Name, target.Name)
				if makeAttack(c.AttackBonus, target.AC) {
					damage := calculateDamage(c.DamageDie)
					target.CurrentHP -= damage
					fmt.Printf("%s hits %s for %d damage! %s HP: %d/%d\n", c.Name, target.Name, damage, target.Name, target.CurrentHP, target.MaxHP)
					if target.CurrentHP <= 0 {
						fmt.Printf("%s has been defeated!\n", target.Name)
					}
				} else {
					fmt.Printf("%s misses %s.\n", c.Name, target.Name)
				}

			case *Monster:
				if c.CurrentHP <= 0 {
					continue // Skip defeated monster
				}
				// Monster's turn
				fmt.Printf("%s's turn (Monster)\n", c.Name)

				// Find a living player to attack
				livingPlayers := []*Character{}
				for _, p := range players {
					if p.CurrentHP > 0 {
						livingPlayers = append(livingPlayers, p)
					}
				}

				if len(livingPlayers) == 0 {
					break // All players defeated
				}

				// Target a random living player
				target := livingPlayers[rand.Intn(len(livingPlayers))]
				fmt.Printf("%s attacks %s!\n", c.Name, target.Name)
				if makeAttack(c.AttackBonus, target.AC) {
					damage := calculateDamage(c.DamageDie)
					target.CurrentHP -= damage
					fmt.Printf("%s hits %s for %d damage! %s HP: %d/%d\n", c.Name, target.Name, damage, target.Name, target.CurrentHP, target.MaxHP)
					if target.CurrentHP <= 0 {
						fmt.Printf("%s has been defeated!\n", target.Name)
					}
				} else {
					fmt.Printf("%s misses %s.\n", c.Name, target.Name)
				}
			}
		}

		// Check for end of combat
		allMonstersDefeated := true
		for _, m := range monsters {
			if m.CurrentHP > 0 {
				allMonstersDefeated = false
				break
			}
		}

		allPlayersDefeated := true
		for _, p := range players {
			if p.CurrentHP > 0 {
				allPlayersDefeated = false
				break
			}
		}

		if allMonstersDefeated {
			fmt.Println("\n--- Combat Ends: Players Win! ---")
			break
		}

		if allPlayersDefeated {
			fmt.Println("\n--- Combat Ends: Monsters Win! ---")
			break
		}
	}

	fmt.Printf("\n--- Simulation Results (Total Rounds: %d) ---\n", round)

	fmt.Println("Players:")
	for _, p := range players {
		if p.CurrentHP > 0 {
			fmt.Printf("- %s: %d/%d HP\n", p.Name, p.CurrentHP, p.MaxHP)
		} else {
			fmt.Printf("- %s: Defeated\n", p.Name)
		}
	}

	fmt.Println("\nMonsters:")
	for _, m := range monsters {
		if m.CurrentHP > 0 {
			fmt.Printf("- %s: %d/%d HP\n", m.Name, m.CurrentHP, m.MaxHP)
		} else {
			fmt.Printf("- %s: Defeated\n", m.Name)
		}
	}
	fmt.Println("-------------------------------------")
}
