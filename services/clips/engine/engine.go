package engine

/*
#cgo LDFLAGS: -lclips -lm
#include "clips.h"
#include "clips_helper.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/aidosgal/neuron/services/clips/entity"
)

type Engine struct {
	env *C.Environment
}

// NewEngine creates a new CLIPS environment
func NewEngine() *Engine {
	env := C.CreateEnvironment()
	return &Engine{env: env}
}

// Destroy cleans up
func (e *Engine) Destroy() {
	C.DestroyEnvironment(e.env)
}

// LoadRules loads multiple CLIPS rule files
func (e *Engine) LoadRules(paths []string) error {
	for _, path := range paths {
		cpath := C.CString(path)
		defer C.free(unsafe.Pointer(cpath))
		if C.Load(e.env, cpath) != C.LE_NO_ERROR {
			return fmt.Errorf("failed to load CLIPS file: %s", path)
		}
	}
	return nil
}

// InjectInput asserts JSON input as CLIPS facts
func (e *Engine) InjectInput(input entity.Input) {
	// Road segments
	for _, seg := range input.RoadSegments {
		fact := fmt.Sprintf(
			`(roadSegment (id "%s") (vehicleCount %d) (avgSpeed %.1f) (capacity %d) (type "%s"))`,
			seg.ID, seg.VehicleCount, seg.AvgSpeed, seg.Capacity, seg.Type)
		cfact := C.CString(fact)
		C.AssertString(e.env, cfact)
		C.free(unsafe.Pointer(cfact))
	}

	// Incidents
	for _, inc := range input.Incidents {
		fact := fmt.Sprintf(
			`(incident (segmentId "%s") (type "%s") (severity %d))`,
			inc.SegmentID, inc.Type, inc.Severity)
		cfact := C.CString(fact)
		C.AssertString(e.env, cfact)
		C.free(unsafe.Pointer(cfact))
	}

	// Weather
	weatherFact := fmt.Sprintf(
		`(weather (condition "%s") (visibility %.2f))`,
		input.Weather.Condition, input.Weather.Visibility)
	cfact := C.CString(weatherFact)
	C.AssertString(e.env, cfact)
	C.free(unsafe.Pointer(cfact))

	// Policy
	policyFact := fmt.Sprintf(
		`(policy (congestionCharge %t) (publicTransportPriority %t))`,
		input.Policy.CongestionCharge, input.Policy.PublicTransportPriority)
	cfact = C.CString(policyFact)
	C.AssertString(e.env, cfact)
	C.free(unsafe.Pointer(cfact))
}

// RunInference executes CLIPS rules
func (e *Engine) RunInference() {
	C.Reset(e.env)
	C.Run(e.env, -1)
}

// Infer runs everything and collects output
func (e *Engine) Infer(input entity.Input) entity.Output {
	e.InjectInput(input)
	e.RunInference()

	decisions := []entity.Decision{}
	explanations := []string{}

	// Iterate all facts
	var fact *C.Fact
	for fact = C.GetNextFact(e.env, nil); fact != nil; fact = C.GetNextFact(e.env, fact) {
		// Get template name
		deftemplate := C.FactDeftemplate(fact)
		templateName := C.GoString(C.DeftemplateName(deftemplate))

		switch templateName {
		case "decision":
			d := entity.Decision{}
			for _, slot := range []string{"action", "segment", "from", "to", "reason", "newCycle", "area", "priority"} {
				slotC := C.CString(slot)
				val := C.GetFactSlotString(fact, slotC)
				if val != nil {
					value := C.GoString(val)
					switch slot {
					case "action":
						d.Action = value
					case "segment":
						d.Segment = value
					case "from":
						d.From = value
					case "to":
						d.To = value
					case "reason":
						d.Reason = value
					case "newCycle":
						d.NewCycle = value
					case "area":
						d.Area = value
					case "priority":
						fmt.Sscanf(value, "%d", &d.Priority)
					}
				}
				C.free(unsafe.Pointer(slotC))
			}
			decisions = append(decisions, d)
		case "explanation":
			slotC := C.CString("text")
			val := C.GetFactSlotString(fact, slotC)
			if val != nil {
				explanations = append(explanations, C.GoString(val))
			}
			C.free(unsafe.Pointer(slotC))
		}
	}

	return entity.Output{
		Decisions:    decisions,
		Explanations: explanations,
	}
}
