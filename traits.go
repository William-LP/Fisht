package main

const (
	BASELINE_HEALTHPOINT = 100
	BASELINE_MAGICPOINT  = 100
	MAGE_MAGIC_BONUS     = 0
	WARRIOR_HEALTH_BONUS = 0
)

type Trait struct {
	HealthPoint,
	MagicPoint int
}

var TraitBonuses = map[ClassName]Trait{
	MAGE:    {HealthPoint: 0, MagicPoint: MAGE_MAGIC_BONUS},
	WARRIOR: {HealthPoint: WARRIOR_HEALTH_BONUS, MagicPoint: 0},
}

func modifyPoints(class ClassName) Trait {
	bonuses, ok := TraitBonuses[class]
	if !ok {
		return Trait{}
	}
	return Trait{
		HealthPoint: BASELINE_HEALTHPOINT + bonuses.HealthPoint,
		MagicPoint:  BASELINE_MAGICPOINT + bonuses.MagicPoint,
	}
}
