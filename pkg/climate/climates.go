package climate

func getAllClimates() []Climate {
	climates := []Climate{
		{
			Name:               "coniferous forest",
			Adjective:          "coniferous forest",
			Temperature:        4,
			HasDeciduousTrees:  false,
			HasConiferousTrees: true,
			Humidity:           6,
			MaxAnimals:         15,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          15,
			MaxSoils:           4,
			MaxStones:          1,
			MaxTrees:           6,
		},
		{
			Name:               "deciduous forest",
			Adjective:          "deciduous forest",
			Temperature:        5,
			HasDeciduousTrees:  true,
			HasConiferousTrees: false,
			Humidity:           5,
			MaxAnimals:         15,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          15,
			MaxSoils:           4,
			MaxStones:          1,
			MaxTrees:           7,
		},
		{
			Name:               "desert",
			Adjective:          "desert",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           0,
			MaxAnimals:         5,
			MaxMetals:          3,
			MaxFish:            2,
			MaxGems:            2,
			MaxPlants:          3,
			MaxSoils:           2,
			MaxStones:          1,
			MaxTrees:           0,
		},
		{
			Name:               "grassland",
			Adjective:          "grassland",
			Temperature:        5,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         10,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          15,
			MaxSoils:           3,
			MaxStones:          1,
			MaxTrees:           3,
		},
		{
			Name:               "marshland",
			Adjective:          "marshy",
			Temperature:        7,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           9,
			HasWetlands:        true,
			MaxAnimals:         15,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            1,
			MaxPlants:          10,
			MaxSoils:           4,
			MaxStones:          0,
			MaxTrees:           3,
		},
		{
			Name:               "tropical",
			Adjective:          "tropical",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           7,
			MaxAnimals:         16,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            4,
			MaxPlants:          16,
			MaxSoils:           3,
			MaxStones:          1,
			MaxTrees:           3,
		},
		{
			Name:               "mountain",
			Adjective:          "mountainous",
			Temperature:        4,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           4,
			MaxAnimals:         5,
			MaxMetals:          10,
			MaxFish:            3,
			MaxGems:            2,
			MaxPlants:          5,
			MaxSoils:           2,
			MaxStones:          5,
			MaxTrees:           3,
		},
		{
			Name:               "rainforest",
			Adjective:          "rainforest",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           9,
			MaxAnimals:         16,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          16,
			MaxSoils:           3,
			MaxStones:          1,
			MaxTrees:           3,
		},
		{
			Name:               "savanna",
			Adjective:          "savanna",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           5,
			MaxAnimals:         9,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            3,
			MaxPlants:          6,
			MaxSoils:           2,
			MaxStones:          1,
			MaxTrees:           3,
		},
		{
			Name:               "steppe",
			Adjective:          "steppe",
			Temperature:        7,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         9,
			MaxMetals:          3,
			MaxFish:            6,
			MaxGems:            3,
			MaxPlants:          8,
			MaxSoils:           2,
			MaxStones:          3,
			MaxTrees:           3,
		},
		{
			Name:               "taiga",
			Adjective:          "taiga",
			Temperature:        3,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         9,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            3,
			MaxPlants:          8,
			MaxSoils:           2,
			MaxStones:          1,
			MaxTrees:           5,
		},
		{
			Name:               "tundra",
			Adjective:          "tundra",
			Temperature:        1,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         6,
			MaxMetals:          3,
			MaxFish:            6,
			MaxGems:            1,
			MaxPlants:          7,
			MaxSoils:           1,
			MaxStones:          2,
			MaxTrees:           3,
		},
	}

	return climates
}
