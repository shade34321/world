package music

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"text/template"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

// Instrument is a musical instrument
type Instrument struct {
	Name                   string
	Description            string
	Type                   string
	BaseMaterialOptions    []string
	SupportMaterialOptions []string
	BaseMaterial           string
	SupportMaterial        string
	DescriptionTemplate    string
}

// GenerateInstruments generates a set of musical instruments for a climate
func GenerateInstruments(originClimate climate.Climate) ([]Instrument, error) {
	var instrument Instrument
	var availableBaseMaterials []string
	var availableSupportMaterials []string
	var woodName string

	availableHides := []string{}
	availableMetals := []string{}
	availableWoods := []string{}
	availableMaterials := []string{}

	allInstruments := getAllInstruments()
	availableInstruments := []Instrument{}
	instruments := []Instrument{}

	for _, i := range originClimate.Metals {
		availableMetals = append(availableMetals, i.Name)
	}

	for _, i := range originClimate.Trees {
		woodName = i.Name
		if !strings.HasSuffix(i.Name, "wood") {
			woodName += "-wood"
		}
		availableWoods = append(availableWoods, woodName)
	}

	hides := resource.ByTag("hide", originClimate.Resources)
	for _, h := range hides {
		availableHides = append(availableHides, h.Name)
	}

	if len(availableHides) > 0 {
		availableMaterials = append(availableMaterials, "hide")
	}
	if len(availableMetals) > 0 {
		availableMaterials = append(availableMaterials, "metal")
	}
	if len(availableWoods) > 0 {
		availableMaterials = append(availableMaterials, "wood")
	}

	for _, i := range allInstruments {
		if slices.StringSlicePartlyWithin(i.BaseMaterialOptions, availableMaterials) {
			if slices.StringSlicePartlyWithin(i.SupportMaterialOptions, availableMaterials) {
				availableInstruments = append(availableInstruments, i)
			}
		}
	}

	numberOfInstruments := rand.Intn(3) + 1

	for i := 0; i < numberOfInstruments; i++ {
		instrument = availableInstruments[rand.Intn(len(availableInstruments))]
		availableBaseMaterials = []string{}
		availableSupportMaterials = []string{}

		for _, m := range instrument.BaseMaterialOptions {
			if slices.StringIn(m, availableMaterials) {
				availableBaseMaterials = append(availableBaseMaterials, m)
			}
		}

		for _, m := range instrument.SupportMaterialOptions {
			if slices.StringIn(m, availableMaterials) {
				availableSupportMaterials = append(availableSupportMaterials, m)
			}
		}

		materialType, err := random.String(availableBaseMaterials)
		if err != nil {
			err = fmt.Errorf("Could not generate instruments: %w", err)
			return []Instrument{}, err
		}
		if materialType == "hide" {
			hide, err := random.String(availableHides)
			if err != nil {
				err = fmt.Errorf("Could not generate instruments: %w", err)
				return []Instrument{}, err
			}
			instrument.BaseMaterial = hide
		} else if materialType == "metal" {
			metal, err := random.String(availableMetals)
			if err != nil {
				err = fmt.Errorf("Could not generate instruments: %w", err)
				return []Instrument{}, err
			}
			instrument.BaseMaterial = metal
		} else if materialType == "wood" {
			wood, err := random.String(availableWoods)
			if err != nil {
				err = fmt.Errorf("Could not generate instruments: %w", err)
				return []Instrument{}, err
			}
			instrument.BaseMaterial = wood
		}

		materialType, err = random.String(availableSupportMaterials)
		if err != nil {
			err = fmt.Errorf("Could not generate instruments: %w", err)
			return []Instrument{}, err
		}
		if materialType == "hide" {
			hide, err := random.String(availableHides)
			if err != nil {
				err = fmt.Errorf("Could not generate instruments: %w", err)
				return []Instrument{}, err
			}
			instrument.SupportMaterial = hide
		} else if materialType == "metal" {
			metal, err := random.String(availableMetals)
			if err != nil {
				err = fmt.Errorf("Could not generate instruments: %w", err)
				return []Instrument{}, err
			}
			instrument.SupportMaterial = metal
		} else if materialType == "wood" {
			wood, err := random.String(availableWoods)
			if err != nil {
				err = fmt.Errorf("Could not generate instruments: %w", err)
				return []Instrument{}, err
			}
			instrument.SupportMaterial = wood
		}

		instrument.Description = instrument.getDescription()

		instruments = append(instruments, instrument)
	}

	return instruments, nil
}

func (instrument Instrument) getDescription() string {
	t := template.New("instrument description")

	var err error
	t, err = t.Parse(instrument.DescriptionTemplate)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, instrument); err != nil {
		panic(err)
	}

	result := tpl.String()

	return result
}

func getAllInstruments() []Instrument {
	instruments := []Instrument{
		{
			Name:                   "short flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		{
			Name:                   "long flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		{
			Name:                   "twin flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		{
			Name:                   "short harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "long harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "full harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "lyre",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "lijerica",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "long-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "pierced lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "short-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		{
			Name:                   "single-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drone",
		},
		{
			Name:                   "multiple-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drones",
		},
		{
			Name:                   "hand drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		{
			Name:                   "short drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		{
			Name:                   "walking drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		{
			Name:                   "heavy drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
	}

	return instruments
}
