package mineral

import "github.com/ironarachne/world/pkg/resource"

// Metals returns all metals
func Metals() []Mineral {
	metals := []Mineral{
		{
			Name:         "adamantine",
			PluralName:   "adamantine",
			Hardness:     9,
			Malleability: 1,
			Commonality:  1,
			Resources: []resource.Resource{
				{
					Name:         "adamantine",
					Origin:       "adamantine",
					MainMaterial: "adamantine",
					Tags: []string{
						"metal ore",
						"hard metal ore",
						"ore",
					},
					Commonality: 1,
					Value:       150,
				},
			},
		},
		{
			Name:         "cold iron",
			PluralName:   "cold iron",
			Hardness:     5,
			Malleability: 4,
			Commonality:  2,
			Resources: []resource.Resource{
				{
					Name:         "cold iron",
					Origin:       "cold iron",
					MainMaterial: "cold iron",
					Tags: []string{
						"metal ore",
						"hard metal ore",
						"ore",
					},
					Commonality: 3,
					Value:       20,
				},
			},
		},
		{
			Name:         "copper",
			PluralName:   "copper",
			Hardness:     5,
			Malleability: 5,
			Commonality:  8,
			Resources: []resource.Resource{
				{
					Name:         "copper",
					Origin:       "copper",
					MainMaterial: "copper",
					Tags: []string{
						"metal ore",
						"hard metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "gold",
			PluralName:   "gold",
			Hardness:     2,
			Malleability: 8,
			Commonality:  5,
			Resources: []resource.Resource{
				{
					Name:         "gold",
					Origin:       "gold",
					MainMaterial: "gold",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 1,
					Value:       50,
				},
			},
		},
		{
			Name:         "iron",
			PluralName:   "iron",
			Hardness:     8,
			Malleability: 4,
			Commonality:  20,
			Resources: []resource.Resource{
				{
					Name:         "iron",
					Origin:       "iron",
					MainMaterial: "iron",
					Tags: []string{
						"metal ore",
						"hard metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       2,
				},
			},
		},
		{
			Name:         "lead",
			PluralName:   "lead",
			Hardness:     1,
			Malleability: 9,
			Commonality:  6,
			Resources: []resource.Resource{
				{
					Name:         "lead",
					Origin:       "lead",
					MainMaterial: "lead",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "mithril",
			PluralName:   "mithril",
			Hardness:     5,
			Malleability: 9,
			Commonality:  2,
			Resources: []resource.Resource{
				{
					Name:         "mithril",
					Origin:       "mithril",
					MainMaterial: "mithril",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 1,
					Value:       120,
				},
			},
		},
		{
			Name:         "nickel",
			PluralName:   "nickel",
			Hardness:     3,
			Malleability: 4,
			Commonality:  3,
			Resources: []resource.Resource{
				{
					Name:         "nickel",
					Origin:       "nickel",
					MainMaterial: "nickel",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "orichalcum",
			PluralName:   "orichalcum",
			Hardness:     1,
			Malleability: 9,
			Commonality:  2,
			Resources: []resource.Resource{
				{
					Name:         "orichalcum",
					Origin:       "orichalcum",
					MainMaterial: "orichalcum",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       80,
				},
			},
		},
		{
			Name:         "silver",
			PluralName:   "silver",
			Hardness:     2,
			Malleability: 7,
			Commonality:  5,
			Resources: []resource.Resource{
				{
					Name:         "silver",
					Origin:       "silver",
					MainMaterial: "silver",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 3,
					Value:       30,
				},
			},
		},
		{
			Name:         "tin",
			PluralName:   "tin",
			Hardness:     3,
			Malleability: 8,
			Commonality:  6,
			Resources: []resource.Resource{
				{
					Name:         "tin",
					Origin:       "tin",
					MainMaterial: "tin",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "zinc",
			PluralName:   "zinc",
			Hardness:     3,
			Malleability: 4,
			Commonality:  5,
			Resources: []resource.Resource{
				{
					Name:         "zinc",
					Origin:       "zinc",
					MainMaterial: "zinc",
					Tags: []string{
						"metal ore",
						"soft metal ore",
						"ore",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return metals
}
