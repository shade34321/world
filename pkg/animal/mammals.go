package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getMammals() []Animal {
	animals := []Animal{
		{
			Name:           "beaver",
			PluralName:     "beavers",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "beaver hide",
					Origin:      "beaver",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "beaver teeth",
					Origin:      "beaver",
					Type:        "teeth",
					Commonality: 5,
				},
				{
					Name:        "beaver meat",
					Origin:      "beaver",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "deer",
			PluralName:     "deer",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "deer hide",
					Origin:      "deer",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "deer teeth",
					Origin:      "deer",
					Type:        "teeth",
					Commonality: 5,
				},
				{
					Name:        "venison",
					Origin:      "deer",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "deer antler",
					Origin:      "deer",
					Type:        "bone",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "squirrel",
			PluralName:     "squirrels",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "squirrel hide",
					Origin:      "squirrel",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "squirrel meat",
					Origin:      "squirrel",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		{
			Name:           "camel",
			PluralName:     "camels",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        true,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "camel hide",
					Origin:      "camel",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "camel teeth",
					Origin:      "camel",
					Type:        "teeth",
					Commonality: 5,
				},
				{
					Name:        "camel meat",
					Origin:      "camel",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "camel milk",
					Origin:      "camel",
					Type:        "milk",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "bison",
			PluralName:     "bison",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "bison hide",
					Origin:      "bison",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "bison teeth",
					Origin:      "bison",
					Type:        "teeth",
					Commonality: 5,
				},
				{
					Name:        "bison",
					Origin:      "bison",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "bison bone",
					Origin:      "bison",
					Type:        "bone",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "cow",
			PluralName:     "cows",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "leather",
					Origin:      "cow",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "cow teeth",
					Origin:      "cow",
					Type:        "teeth",
					Commonality: 5,
				},
				{
					Name:        "beef",
					Origin:      "cow",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "beef loin",
					Origin:      "cow",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "beef ribs",
					Origin:      "cow",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "cow bone",
					Origin:      "cow",
					Type:        "bone",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "elephant",
			PluralName:     "elephants",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        true,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "elephant hide",
					Origin:      "elephant",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "elephant milk",
					Origin:      "elephant",
					Type:        "milk",
					Commonality: 5,
				},
				{
					Name:        "elephant",
					Origin:      "elephant",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "ivory",
					Origin:      "elephant",
					Type:        "bone",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("large"),
		},
		{
			Name:           "goat",
			PluralName:     "goats",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "goat hide",
					Origin:      "goat",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "goat's milk",
					Origin:      "goat",
					Type:        "milk",
					Commonality: 5,
				},
				{
					Name:        "goat",
					Origin:      "goat",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "sheep",
			PluralName:     "sheep",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "wool",
					Origin:      "sheep",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "sheep's milk",
					Origin:      "sheep",
					Type:        "milk",
					Commonality: 5,
				},
				{
					Name:        "lamb",
					Origin:      "sheep",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "alpaca",
			PluralName:     "alpacas",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 4,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:        "alpaca wool",
					Origin:      "alpaca",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "alpaca milk",
					Origin:      "alpaca",
					Type:        "milk",
					Commonality: 5,
				},
				{
					Name:        "alpaca",
					Origin:      "alpaca",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "llama",
			PluralName:     "llamas",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:        "llama wool",
					Origin:      "llama",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "llama",
					Origin:      "llama",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "llama milk",
					Origin:      "llama",
					Type:        "milk",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "hippopotamus",
			PluralName:     "hippopotamuses",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      true,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "hippopotamus hide",
					Origin:      "hippopotamus",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "hippopotamus",
					Origin:      "hippopotamus",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "hippopotamus milk",
					Origin:      "hippopotamus",
					Type:        "milk",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("large"),
		},
		{
			Name:           "antelope",
			PluralName:     "antelopes",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "antelope hide",
					Origin:      "antelope",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "antelope",
					Origin:      "antelope",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "antelope milk",
					Origin:      "antelope",
					Type:        "milk",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "gazelle",
			PluralName:     "gazelles",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:        "gazelle hide",
					Origin:      "gazelle",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "gazelle",
					Origin:      "gazelle",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "gazelle milk",
					Origin:      "gazelle",
					Type:        "milk",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "rabbit",
			PluralName:     "rabbits",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "rabbit hide",
					Origin:      "rabbit",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "rabbit",
					Origin:      "rabbit",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		{
			Name:           "ermine",
			PluralName:     "ermines",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "ermine fur",
					Origin:      "ermine",
					Type:        "fur",
					Commonality: 5,
				},
				{
					Name:        "ermine",
					Origin:      "ermine",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "mink",
			PluralName:     "minks",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:        "mink fur",
					Origin:      "mink",
					Type:        "fur",
					Commonality: 5,
				},
				{
					Name:        "mink",
					Origin:      "mink",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "pig",
			PluralName:     "pigs",
			AnimalType:     "mammal",
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
			Resources: []resource.Resource{
				{
					Name:        "pig hide",
					Origin:      "pig",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "pork",
					Origin:      "pig",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "pork loin",
					Origin:      "pig",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "pork intestine",
					Origin:      "pig",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "pork ribs",
					Origin:      "pig",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "bacon",
					Origin:      "pig",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
		{
			Name:           "raccoon",
			PluralName:     "raccoons",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:        "raccoon hide",
					Origin:      "raccoon",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "raccoon",
					Origin:      "raccoon",
					Type:        "meat",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("small"),
		},
		{
			Name:           "reindeer",
			PluralName:     "reindeer",
			AnimalType:     "mammal",
			EatsAnimals:    false,
			EatsPlants:     true,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   true,
			IsScavenger:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 4,
			Resources: []resource.Resource{
				{
					Name:        "reindeer hide",
					Origin:      "reindeer",
					Type:        "hide",
					Commonality: 5,
				},
				{
					Name:        "reindeer",
					Origin:      "reindeer",
					Type:        "meat",
					Commonality: 5,
				},
				{
					Name:        "reindeer milk",
					Origin:      "reindeer",
					Type:        "milk",
					Commonality: 5,
				},
				{
					Name:        "reindeer bone",
					Origin:      "reindeer",
					Type:        "bone",
					Commonality: 5,
				},
			},
			Size: size.GetCategoryByName("medium"),
		},
	}

	return animals
}
