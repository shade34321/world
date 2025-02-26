package food

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/slices"
)

// Style is a cultural food style
type Style struct {
	Breads             []string
	CommonBases        []string
	CommonDessertBases []string
	CommonDesserts     []string
	CommonMainDishes   []string
	CommonSpices       []string
	CookingTechniques  []string
	EatingTraits       []string
	DessertTypes       []string
	Noodles            []string
}

// GenerateStyle procedurally generates a style of food
func GenerateStyle(originClimate climate.Climate) (Style, error) {
	chanceForGoldFlakes := 0
	style := Style{}

	for _, r := range originClimate.Resources {
		if r.HasTag("meat") {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.HasTag("spice") || r.HasTag("herb") {
			if r.Name == "gold" {
				chanceForGoldFlakes = rand.Intn(100)
				if chanceForGoldFlakes > 89 {
					style.CommonSpices = append(style.CommonSpices, "gold flakes")
				}
			} else {
				style.CommonSpices = append(style.CommonSpices, r.Name)
			}
		} else if r.HasTag("egg") {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.HasTag("milk") {
			if !slices.StringIn("pudding", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "pudding")
			}
		} else if r.HasTag("fruit") {
			style.CommonDessertBases = append(style.CommonDessertBases, r.Name)
			if !slices.StringIn("fruit", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "fruit")
			}
		} else if r.HasTag("vegetable") || r.HasTag("squash") {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.HasTag("flour") {
			style.Noodles = append(style.Noodles, getNoodles(r.Name)...)
			if !slices.StringIn("pie", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "pie")
			}
			if !slices.StringIn("crisp", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "crisp")
			}
			if !slices.StringIn("tart", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "tart")
			}
			if !slices.StringIn("cake", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "cake")
			}
		}
	}

	if !slices.StringIn("salt", style.CommonSpices) {
		style.CommonSpices = append(style.CommonSpices, "salt")
	}

	techniques, err := randomTechniques(3)
	if err != nil {
		err = fmt.Errorf("Could not generate food style: %w", err)
		return Style{}, err
	}
	style.CookingTechniques = techniques

	desserts, err := style.randomDesserts()
	if err != nil {
		err = fmt.Errorf("Could not generate food style: %w", err)
		return Style{}, err
	}
	style.CommonDesserts = desserts
	mainDishes, err := style.randomMainDishes()
	if err != nil {
		err = fmt.Errorf("Could not generate food style: %w", err)
		return Style{}, err
	}
	style.CommonMainDishes = mainDishes
	breads, err := randomBreads(originClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate food style: %w", err)
		return Style{}, err
	}
	style.Breads = breads
	traits, err := randomEatingTraits()
	if err != nil {
		err = fmt.Errorf("Could not generate food style: %w", err)
		return Style{}, err
	}
	style.EatingTraits = traits

	return style, nil
}

// Random generates a completely random style of food
func Random() (Style, error) {
	originClimate, err := climate.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate random food style: %w", err)
		return Style{}, err
	}
	style, err := GenerateStyle(originClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate random food style: %w", err)
		return Style{}, err
	}

	return style, nil
}
