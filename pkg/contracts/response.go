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

type Centre struct {
	Name     string     `json:"name"`
	Pin      int        `json:"pincode"`
	From     string     `json:"from"`
	To       string     `json:"to"`
	FeeType  string     `json:"fee_type"`
	Sessions []*Session `json:"sessions"`
}
