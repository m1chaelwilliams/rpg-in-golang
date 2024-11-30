package entities

import "rpg-tutorial/components"

type Enemy struct {
	*Sprite
	FollowsPlayer bool
	CombatComp    *components.EnemyCombat
}
