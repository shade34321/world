package food

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func randomEatingTraits() ([]string, error) {
	var traits []string
	var trait string
	var typesOfTraits []string

	utensils := []string{
		"chopsticks",
		"fork",
		"hands",
		"knife and fork",
		"knife",
		"spoon",
	}

	spices := []string{
		"heavy",
		"moderate",
		"minimal",
	}

	heatLevels := []string{
		"hot",
		"warm",
		"cold",
		"chilled",
	}

	customs := []string{
		"eat communal meals",
		"eat alone",
		"belch after eating",
		"thank the cook after eating",
	}

	potentialTraits := []string{
		"utensils",
		"spices",
		"heat",
		"customs",
	}

	for i := 0; i < 2; i++ {
		t, err := random.String(potentialTraits)
		if err != nil {
			err = fmt.Errorf("Could not generate food traits: %w", err)
			return []string{}, err
		}
		if !slices.StringIn(t, typesOfTraits) {
			typesOfTraits = append(typesOfTraits, t)
			if t == "utensils" {
				utensil, err := random.String(utensils)
				if err != nil {
					err = fmt.Errorf("Could not generate food traits: %w", err)
					return []string{}, err
				}
				trait = "eat with " + utensil
			} else if t == "spices" {
				spice, err := random.String(spices)
				if err != nil {
					err = fmt.Errorf("Could not generate food traits: %w", err)
					return []string{}, err
				}
				trait = "use " + spice + " spice"
			} else if t == "heat" {
				heat, err := random.String(heatLevels)
				if err != nil {
					err = fmt.Errorf("Could not generate food traits: %w", err)
					return []string{}, err
				}
				trait = "serve food " + heat
			} else if t == "customs" {
				custom, err := random.String(customs)
				if err != nil {
					err = fmt.Errorf("Could not generate food traits: %w", err)
					return []string{}, err
				}
				trait = custom
			}
			traits = append(traits, trait)
		}
	}

	return traits, nil
}
