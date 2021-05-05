package contracts

type SlotResponse struct {
	Centres []*Centre `json:"centers"`
}

type Session struct {
	Date     string  `json:"date"`
	Capacity float64 `json:"available_capacity"`
	AgeLimit int     `json:"min_age_limit"`
	Vaccine  string  `json:"vaccine"`
}

type StateInfo struct {
	StateId   int    `json:"state_id"`
	StateName string `json:"state_name"`
}

type StateResponse struct {
	States []StateInfo `json:"states"`
}

type DistrictInfo struct {
	DistrictId   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}

type DistrictResponse struct {
	Districts []DistrictInfo `json:"districts"`
}

type Centre struct {
	Name     string     `json:"name"`
	Pin      int        `json:"pincode"`
	From     string     `json:"from"`
	To       string     `json:"to"`
	FeeType  string     `json:"fee_type"`
	Sessions []*Session `json:"sessions"`
}
