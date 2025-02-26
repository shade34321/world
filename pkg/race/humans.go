package race

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"
)

func getHumans() []species.Species {
	common := getHumanCommonTraitTemplates()
	possible := getHumanPossibleTraitTemplates()
	ageCategories := getHumanAgeCategories()

	races := []species.Species{
		{
			Name:           "human",
			PluralName:     "humans",
			Adjective:      "human",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    100,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"human",
			},
		},
	}

	return races
}

func getHumanAgeCategories() []age.Category {
	heightDice := dice.Dice{Number: 2, Sides: 8}
	weightDice := dice.Dice{Number: 2, Sides: 10}
	adultSizeCategory := size.GetCategoryByName("medium")
	childSizeCategory := size.GetCategoryByName("small")
	infantSizeCategory := size.GetCategoryByName("tiny")

	categories := []age.Category{
		{
			Name:             "elderly",
			MinAge:           70,
			MaxAge:           110,
			MaleHeightBase:   62,
			FemaleHeightBase: 57,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   140,
			FemaleWeightBase: 105,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      30,
		},
		{
			Name:             "adult",
			MinAge:           26,
			MaxAge:           69,
			MaleHeightBase:   62,
			FemaleHeightBase: 57,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   140,
			FemaleWeightBase: 105,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      150,
		},
		{
			Name:             "young adult",
			MinAge:           20,
			MaxAge:           25,
			MaleHeightBase:   62,
			FemaleHeightBase: 57,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   140,
			FemaleWeightBase: 105,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      100,
		},
		{
			Name:             "teenager",
			MinAge:           13,
			MaxAge:           19,
			MaleHeightBase:   61,
			FemaleHeightBase: 59,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   92,
			FemaleWeightBase: 89,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      30,
		},
		{
			Name:             "child",
			MinAge:           2,
			MaxAge:           12,
			MaleHeightBase:   42,
			FemaleHeightBase: 42,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   36,
			FemaleWeightBase: 36,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     childSizeCategory,
			Commonality:      10,
		},
		{
			Name:             "infant",
			MinAge:           0,
			MaxAge:           1,
			MaleHeightBase:   18,
			FemaleHeightBase: 18,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   5,
			FemaleWeightBase: 5,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     infantSizeCategory,
			Commonality:      1,
		},
	}

	return categories
}

func getHumanCommonTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "eye color",
			PossibleValues: []string{
				"amber",
				"blue",
				"brown",
				"gold",
				"green",
				"hazel",
				"grey",
			},
			PossibleDescriptors: []string{
				"{{.Value}} eyes",
			},
			Tags: []string{
				"appearance",
				"physical",
				"eyes",
			},
		},
		{
			Name: "hair color",
			PossibleValues: []string{
				"black",
				"auburn",
				"red",
				"grey",
				"brown",
				"dark brown",
				"light brown",
			},
			PossibleDescriptors: []string{
				"{{.Value}} hair",
			},
			Tags: []string{
				"appearance",
				"physical",
				"hair",
			},
		},
		{
			Name: "skin color",
			PossibleValues: []string{
				"light",
				"tan",
				"olive",
				"bronze",
				"brown",
				"dark",
			},
			PossibleDescriptors: []string{
				"{{.Value}} skin",
			},
			Tags: []string{
				"appearance",
				"physical",
				"skin",
			},
		},
	}

	return templates
}

func getHumanPossibleTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "nose shape",
			PossibleValues: []string{
				"aquiline",
				"broad",
				"flat",
				"long",
				"narrow",
				"pointed",
				"upturned",
			},
			PossibleDescriptors: []string{
				"{{.Value}} nose",
			},
			Tags: []string{
				"appearance",
				"physical",
				"nose",
			},
		},
	}

	return templates
}
