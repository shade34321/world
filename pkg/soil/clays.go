package soil

import "github.com/ironarachne/world/pkg/resource"

// Clays returns all clays
func Clays() []Soil {
	soils := []Soil{
		{
			Name:               "earthenware",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "earthenware",
					Origin:       "earthenware",
					MainMaterial: "earthenware",
					Tags: []string{
						"clay",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:               "porcelain",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "porcelain",
					Origin:       "porcelain",
					MainMaterial: "porcelain",
					Tags: []string{
						"clay",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:               "stoneware",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "stoneware",
					Origin:       "stoneware",
					MainMaterial: "stoneware",
					Tags: []string{
						"clay",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:               "terra cotta",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "terra cotta",
					Origin:       "terra cotta",
					MainMaterial: "terra cotta",
					Tags: []string{
						"clay",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return soils
}
