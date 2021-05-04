package contracts

import (
	"github.com/olekukonko/tablewriter"
)

type Location struct {
	State      string
	City       string
	DistrictId string
}

type SlotRequest struct {
	Location Location
	Date     string
}

type AvailData struct {
	Available string
	Vaccine   string
	Color     tablewriter.Colors
}
