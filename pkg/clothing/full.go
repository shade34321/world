package clothing

import "fmt"

func getRandomDress() (Item, error) {
	template := ItemTemplate{
		Name:         "dress",
		Type:         "full",
		MaterialType: "fabric",
		PrefixModifiers: []string{
			"form-fitting",
			"long-sleeved",
			"loose",
			"short-sleeved",
			"sleeveless",
			"thin-strapped",
			"tight-sleeved",
		},
		SuffixModifiers: []string{
			"with a tight top and loose skirt",
			"with a tight top and flared skirt",
			"with short sleeves",
			"with long sleeves",
			"with ribbed sleeves",
			"with sleeves flared at the wrists",
			"with tight sleeves",
			"with no sleeves",
			"with tapered sleeves",
		},
	}

	dress, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random dress: %w", err)
		return Item{}, err
	}

	return dress, nil
}

func getRandomRobe() (Item, error) {
	template := ItemTemplate{
		Name:         "robe",
		Type:         "full",
		MaterialType: "fabric",
		PrefixModifiers: []string{
			"long-sleeved",
			"loose",
			"short-sleeved",
			"sleeveless",
			"tight-sleeved",
		},
		SuffixModifiers: []string{
			"with short sleeves",
			"with long sleeves",
			"with ribbed sleeves",
			"with sleeves flared at the wrists",
			"with tight sleeves",
			"with no sleeves",
			"with tapered sleeves",
		},
	}

	robe, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random robe: %w", err)
		return Item{}, err
	}

	return robe, nil
}
