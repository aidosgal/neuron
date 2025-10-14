package main

import (
	"fmt"

	"github.com/aidosgal/neuron/services/clips/engine"
	"github.com/aidosgal/neuron/services/clips/entity"
)

func main() {
	engine := engine.NewEngine()
	defer engine.Destroy()

	engine.LoadRules([]string{
		"services/clips/rules/facts.clp",
		"services/clips/rules/congestion.clp",
		"services/clips/rules/incidents.clp",
		"services/clips/rules/weather.clp",
		"services/clips/rules/policy.clp",
		"services/clips/rules/routing.clp",
		"services/clips/rules/coordination.clp",
	})

	input := entity.Input{
		Timestamp: "2025-10-14T08:20:00Z",
		Location:  "Downtown",
		RoadSegments: []entity.RoadSegment{
			{
				ID:           "R1",
				VehicleCount: 150,
				AvgSpeed:     12.5,
				Capacity:     300,
				Type:         "arterial",
			},
			{
				ID:           "R2",
				VehicleCount: 80,
				AvgSpeed:     35.0,
				Capacity:     200,
				Type:         "highway",
			},
			{
				ID:           "R3",
				VehicleCount: 60,
				AvgSpeed:     18.0,
				Capacity:     150,
				Type:         "local",
			},
		},
		Incidents: []entity.Incident{
			{
				SegmentID: "R1",
				Type:      "accident",
				Severity:  2,
			},
			{
				SegmentID: "R2",
				Type:      "construction",
				Severity:  1,
			},
		},
		Weather: entity.Weather{
			Condition:  "rain",
			Visibility: 0.7,
		},
		Policy: entity.Policy{
			CongestionCharge:        true,
			PublicTransportPriority: true,
		},
	}

	output := engine.Infer(input)

	fmt.Println("Decisions:", output.Decisions)
	fmt.Println("Explanations:", output.Explanations)
}
