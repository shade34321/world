package insect

import "github.com/ironarachne/world/pkg/resource"

func getBeetles() []Insect {
	insects := []Insect{
		{
			Name:           "rhino beetle",
			PluralName:     "rhino beetles",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "rhino beetle carapace",
					Origin:       "rhino beetle",
					MainMaterial: "rhino beetle",
					Tags: []string{
						"carapace",
					},
					Commonality: 3,
					Value:       2,
				},
			},
			Tags: []string{
				"beetle",
				"herbivore",
			},
		},
	}

	return insects
}
