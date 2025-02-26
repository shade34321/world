package world

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/worldmap"
)

// World is a fantasy world
type World struct {
	Name                string
	WorldMap            worldmap.WorldMap
	Countries           []country.Country
	OccupiedCoordinates []grid.Coordinate
}

// Generate procedurally generates a world
func Generate() (World, error) {
	var boundary worldmap.Boundary
	var boundaries []worldmap.Boundary
	var activeTile worldmap.Tile
	var activeTileCoordinates grid.Coordinate
	var homeTile worldmap.Tile
	var homeTileCoordinates grid.Coordinate
	var newRegions []region.Region
	var numRegions int
	var regionCoordinates []grid.Coordinate
	var contiguousAvailableCoords []grid.Coordinate
	var contiguousAvailableRegionCoords []grid.Coordinate
	var adjacentTiles []worldmap.Tile
	var adjacentRegionTiles []worldmap.Tile
	var pattern string
	var style string

	world := World{}

	world.Name = "WORLD"
	world.WorldMap = worldmap.Generate(50, 80)

	numCountries := rand.Intn(20) + 10
	baseRegions := 40 - numCountries
	availableCoords := world.WorldMap.AllLandCoordinates()

	for i := 0; i < numCountries; i++ {
		newCountry, err := country.Generate()
		if err != nil {
			err = fmt.Errorf("Could not generate world: %w", err)
			return World{}, err
		}
		numRegions = rand.Intn(12) + baseRegions
		newRegions = []region.Region{}
		regionCoordinates = []grid.Coordinate{}

		homeTileCoordinates = availableCoords[rand.Intn(len(availableCoords))]
		homeTile = world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X]
		homeTile.IsInhabited = true
		world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X] = homeTile
		homeClimate, err := climate.Generate(homeTile.TileType)

		newRegion, err := region.Generate(homeClimate, newCountry.DominantCulture)
		if err != nil {
			err = fmt.Errorf("Could not generate world: %w", err)
			return World{}, err
		}
		newCountry.DominantCulture = newRegion.Culture
		regionCoordinates = append(regionCoordinates, homeTileCoordinates)
		activeTile = homeTile
		for j := 1; j < newRegion.Class.NumberOfTiles; j++ {
			adjacentTiles = world.WorldMap.GetAdjacentTiles(activeTile)
			contiguousAvailableCoords = []grid.Coordinate{}
			for _, t := range adjacentTiles {
				if !t.IsOcean && !t.IsInhabited {
					contiguousAvailableCoords = append(contiguousAvailableCoords, t.Coordinate)
				}
			}
			if len(contiguousAvailableCoords) > 0 {
				if len(contiguousAvailableCoords) == 1 {
					activeTileCoordinates = contiguousAvailableCoords[0]
				} else {
					activeTileCoordinates = contiguousAvailableCoords[rand.Intn(len(contiguousAvailableCoords))]
				}
				activeTile = world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X]
				activeTile.IsInhabited = true
				world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X] = activeTile
				regionCoordinates = append(regionCoordinates, activeTileCoordinates)
			}
		}
		newRegion = newRegion.AssignTiles(regionCoordinates)
		world.OccupiedCoordinates = append(world.OccupiedCoordinates, regionCoordinates...)
		availableCoords = grid.RemoveCoordinatesFromSlice(world.OccupiedCoordinates, availableCoords)
		newRegions = append(newRegions, newRegion)

		for n := 0; n < numRegions; n++ {
			regionCoordinates = []grid.Coordinate{}
			contiguousAvailableCoords = []grid.Coordinate{}
			adjacentTiles = world.WorldMap.GetAdjacentTiles(homeTile)
			for _, t := range adjacentTiles {
				if !t.IsOcean && !t.IsInhabited {
					contiguousAvailableCoords = append(contiguousAvailableCoords, t.Coordinate)
				}
			}
			if len(contiguousAvailableCoords) > 0 {
				if len(contiguousAvailableCoords) == 1 {
					homeTileCoordinates = contiguousAvailableCoords[0]
				} else {
					homeTileCoordinates = contiguousAvailableCoords[rand.Intn(len(contiguousAvailableCoords))]
				}
				homeTile = world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X]
				homeTile.IsInhabited = true
				world.WorldMap.Tiles[homeTileCoordinates.Y][homeTileCoordinates.X] = homeTile
				homeClimate, err := climate.Generate(homeTile.TileType)

				newRegion, err := region.Generate(homeClimate, newCountry.DominantCulture)
				if err != nil {
					err = fmt.Errorf("Could not generate world: %w", err)
					return World{}, err
				}
				regionCoordinates = append(regionCoordinates, homeTileCoordinates)
				activeTile = homeTile
				for j := 1; j < newRegion.Class.NumberOfTiles; j++ {
					adjacentRegionTiles = world.WorldMap.GetAdjacentTiles(activeTile)
					contiguousAvailableRegionCoords = []grid.Coordinate{}
					for _, t := range adjacentRegionTiles {
						if !t.IsOcean && !t.IsInhabited {
							contiguousAvailableRegionCoords = append(contiguousAvailableRegionCoords, t.Coordinate)
						}
					}
					if len(contiguousAvailableRegionCoords) > 0 {
						if len(contiguousAvailableRegionCoords) == 1 {
							activeTileCoordinates = contiguousAvailableRegionCoords[0]
						} else {
							activeTileCoordinates = contiguousAvailableRegionCoords[rand.Intn(len(contiguousAvailableRegionCoords))]
						}
						activeTile = world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X]
						activeTile.IsInhabited = true
						world.WorldMap.Tiles[activeTileCoordinates.Y][activeTileCoordinates.X] = activeTile
						regionCoordinates = append(regionCoordinates, activeTileCoordinates)
					}
				}
				newRegion = newRegion.AssignTiles(regionCoordinates)
				world.OccupiedCoordinates = append(world.OccupiedCoordinates, regionCoordinates...)
				availableCoords = grid.RemoveCoordinatesFromSlice(world.OccupiedCoordinates, availableCoords)
				newRegions = append(newRegions, newRegion)
			}
		}
		newCountry.Regions = newRegions
		world.Countries = append(world.Countries, newCountry)
	}

	for _, c := range world.Countries {
		pattern = ""

		boundary = worldmap.CreateBoundary(c.Name, style, pattern, c.GetAllTiles(world.WorldMap))
		boundary.Edges = world.WorldMap.CalculateBoundaryEdges(boundary)
		boundary.Vertices = world.WorldMap.CalculateBoundaryVertices(boundary)
		boundaries = append(boundaries, boundary)
	}
	world.WorldMap.Boundaries = boundaries
	world.WorldMap.SVG = world.WorldMap.RenderAsSVG()

	return world, nil
}
