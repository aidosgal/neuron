package entity

type RoadSegment struct {
    ID           string  `json:"id"`
    VehicleCount int     `json:"vehicleCount"`
    AvgSpeed     float64 `json:"avgSpeed"`
    Capacity     int     `json:"capacity"`
    Type         string  `json:"type"`
}

type Incident struct {
    SegmentID string `json:"segmentId"`
    Type      string `json:"type"`
    Severity  int    `json:"severity"`
}

type Weather struct {
    Condition  string  `json:"condition"`
    Visibility float64 `json:"visibility"`
}

type Policy struct {
    CongestionCharge        bool `json:"congestionCharge"`
    PublicTransportPriority bool `json:"publicTransportPriority"`
}

type Input struct {
    Timestamp    string        `json:"timestamp"`
    Location     string        `json:"location"`
    RoadSegments []RoadSegment `json:"roadSegments"`
    Incidents    []Incident    `json:"incidents"`
    Weather      Weather       `json:"weather"`
    Policy       Policy        `json:"policy"`
}

type Decision struct {
    Action   string `json:"action"`
    Segment  string `json:"segment,omitempty"`
    From     string `json:"from,omitempty"`
    To       string `json:"to,omitempty"`
    Reason   string `json:"reason,omitempty"`
    NewCycle string `json:"newCycle,omitempty"`
    Area     string `json:"area,omitempty"`
    Priority int    `json:"priority,omitempty"`
}

type Output struct {
    Decisions    []Decision `json:"decisions"`
    Explanations []string   `json:"explanations"`
}
