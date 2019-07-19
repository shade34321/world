package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getBirds() []Animal {
	birds := []Animal{
		Animal{
			Name:           "goose",
			PluralName:     "geese",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "chicken",
			PluralName:     "chickens",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "mudhen",
			PluralName:     "mudhens",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "quail",
			PluralName:     "quails",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "duck",
			PluralName:     "ducks",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "flamingo",
			PluralName:     "flamingos",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    5,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "eagle",
			PluralName:     "eagles",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "falcon",
			PluralName:     "falcons",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "hawk",
			PluralName:     "hawks",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Size:           size.GetCategoryByName("tiny"),
		},
	}

	for _, b := range birds {
		b.AnimalType = "bird"
		b.Resources = []resource.Resource{
			resource.Resource{
				Name:        b.Name + " eggs",
				Origin:      b.Name,
				Type:        "egg",
				Commonality: 4,
			},
			resource.Resource{
				Name:        b.Name + " feathers",
				Origin:      b.Name,
				Type:        "feather",
				Commonality: 4,
			},
		}
	}

	return birds
}
