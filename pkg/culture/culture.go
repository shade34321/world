package culture

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/drink"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/music"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/species"
)

// Culture is a fantasy culture
type Culture struct {
	Adjective       string
	AlcoholicDrinks []drink.Drink
	AttributeMax    int
	Attributes      struct {
		Aggression      int
		Curiosity       int
		MagicPrevalence int
		MagicStrength   int
		Rigidity        int
		Superstition    int
	}
	BuildingStyle     buildings.BuildingStyle
	ClothingStyle     clothing.Style
	CommonFamilyNames []string
	CommonFemaleNames []string
	CommonMaleNames   []string
	FoodStyle         food.Style
	HomeClimate       climate.Climate
	Language          language.Language
	MusicStyle        music.Style
	Name              string
	PrimaryRace       species.Species
	Religion          religion.Religion
	Views             []string
}

// Generate generates a culture
func Generate(homeClimate climate.Climate) (Culture, error) {
	culture := Culture{}
	culture.HomeClimate = homeClimate

	cultureLanguage, err := language.Generate()
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.Language = cultureLanguage
	wordList, err := culture.createWordList()
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.Language.WordList = wordList

	maleNames, err := culture.Language.GenerateNameList("male")
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.CommonMaleNames = maleNames
	femaleNames, err := culture.Language.GenerateNameList("female")
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.CommonFemaleNames = femaleNames
	familyNames, err := culture.Language.GenerateNameList("family")
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.CommonFamilyNames = familyNames

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective

	instruments, err := music.GenerateInstruments(culture.HomeClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	musicStyle, err := music.GenerateStyle(instruments)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.MusicStyle = musicStyle

	buildingStyle, err := buildings.GenerateStyle()
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.BuildingStyle = buildingStyle
	clothingStyle, err := clothing.GenerateStyle(culture.HomeClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.ClothingStyle = clothingStyle
	foodStyle, err := food.GenerateStyle(culture.HomeClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.FoodStyle = foodStyle
	drinks, err := drink.RandomSet(3, culture.HomeClimate.Resources)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.AlcoholicDrinks = drinks

	culture.AttributeMax = 100
	culture.Attributes.Aggression = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.Curiosity = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.MagicPrevalence = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.MagicStrength = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.Rigidity = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.Superstition = rand.Intn(culture.AttributeMax) + 1

	primaryRace, err := race.RandomWeighted()
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.PrimaryRace = primaryRace

	cultureReligion, err := religion.Generate(culture.Language)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture.Religion = cultureReligion

	culture.Views = culture.getViews()

	return culture, nil
}

// Random returns a completely random culture
func Random() (Culture, error) {
	homeClimate, err := climate.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	culture, err := Generate(homeClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate culture: %w", err)
		return Culture{}, err
	}
	return culture, nil
}

func (culture Culture) createWordList() (map[string]string, error) {
	list := culture.Language.WordList

	for _, i := range culture.HomeClimate.Animals {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Fish {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Gems {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Insects {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Metals {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.OtherMinerals {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Plants {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Seasons {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Stones {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Fish {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range culture.HomeClimate.Trees {
		if !language.IsInWordList(i.Name, list) {
			modifiedList, err := culture.Language.AddNounToWordList(i.Name)
			if err != nil {
				err = fmt.Errorf("Could not generate culture: %w", err)
				return list, err
			}
			list = modifiedList
		}
	}

	return list, nil
}
