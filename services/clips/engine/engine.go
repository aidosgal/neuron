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

func NewEngine() *Engine {
	env := C.CreateEnvironment()
	return &Engine{env: env}
}

func (e *Engine) Destroy() {
	C.DestroyEnvironment(e.env)
}

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

func (e *Engine) InjectInput(input entity.Input) {
	for _, seg := range input.RoadSegments {
		fact := fmt.Sprintf(
			`(roadSegment (id "%s") (vehicleCount %d) (avgSpeed %.1f) (capacity %d) (type "%s"))`,
			seg.ID, seg.VehicleCount, seg.AvgSpeed, seg.Capacity, seg.Type)
		cfact := C.CString(fact)
		C.AssertString(e.env, cfact)
		C.free(unsafe.Pointer(cfact))
	}

	for _, inc := range input.Incidents {
		fact := fmt.Sprintf(
			`(incident (segmentId "%s") (type "%s") (severity %d))`,
			inc.SegmentID, inc.Type, inc.Severity)
		cfact := C.CString(fact)
		C.AssertString(e.env, cfact)
		C.free(unsafe.Pointer(cfact))
	}

	weatherFact := fmt.Sprintf(
		`(weather (condition "%s") (visibility %.2f))`,
		input.Weather.Condition, input.Weather.Visibility)
	cfact := C.CString(weatherFact)
	C.AssertString(e.env, cfact)
	C.free(unsafe.Pointer(cfact))

	policyFact := fmt.Sprintf(
		`(policy (congestionCharge %t) (publicTransportPriority %t))`,
		input.Policy.CongestionCharge, input.Policy.PublicTransportPriority)
	cfact = C.CString(policyFact)
	C.AssertString(e.env, cfact)
	C.free(unsafe.Pointer(cfact))
}

func (e *Engine) RunInference() {
	C.Run(e.env, -1)
}

func (e *Engine) Infer(input entity.Input) entity.Output {
	C.Reset(e.env)
	e.InjectInput(input)
	e.RunInference()

	decisions := []entity.Decision{}
	explanations := []string{}

	var fact *C.Fact
	for fact = C.GetNextFactWrapper(e.env, nil); fact != nil; fact = C.GetNextFactWrapper(e.env, fact) {
		deftemplate := C.GetFactDeftemplate(fact)
		if deftemplate == nil {
			continue
		}
		templateNameC := C.GetDeftemplateName(deftemplate)
		if templateNameC == nil {
			continue
		}
		templateName := C.GoString(templateNameC)

		switch templateName {
		case "decision":
			d := entity.Decision{}
			for _, slot := range []string{"action", "segment", "from", "to", "reason", "newCycle", "area", "priority"} {
				slotC := C.CString(slot)
				val := C.GetFactSlotString(fact, slotC)
				C.free(unsafe.Pointer(slotC))

				if val != nil {
					value := C.GoString(val)
					C.FreeString(val)

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
			}
			decisions = append(decisions, d)
		case "explanation":
			slotC := C.CString("text")
			val := C.GetFactSlotString(fact, slotC)
			C.free(unsafe.Pointer(slotC))

			if val != nil {
				explanations = append(explanations, C.GoString(val))
				C.FreeString(val)
			}
		}
	}

	return entity.Output{
		Decisions:    decisions,
		Explanations: explanations,
	}
}
