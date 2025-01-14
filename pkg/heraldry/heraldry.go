package heraldry

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ironarachne/world/pkg/words"
)

// Device is the entire coat of arms
type Device struct {
	Blazon string
	Field
	ImageURL string
	FileName string
	GUID     string
}

// Generate procedurally generates a random heraldic device and returns it.
func Generate() (Device, error) {
	field, err := randomField()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	d := Device{
		Field: field,
	}

	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"
	d.Blazon = d.RenderToBlazon()
	d.Blazon = words.CapitalizeFirst(d.Blazon)
	imageURL, err := d.RenderToPNG()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}
	d.ImageURL = imageURL

	return d, nil
}

// GenerateByFieldName generates a random heraldic device with a given field shape.
func GenerateByFieldName(name string) (Device, error) {
	field, err := fieldByName(name)
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	d := Device{
		Field: field,
	}

	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"
	d.Blazon = d.RenderToBlazon()
	d.Blazon = words.CapitalizeFirst(d.Blazon)
	imageURL, err := d.RenderToPNG()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}
	d.ImageURL = imageURL

	return d, nil
}
