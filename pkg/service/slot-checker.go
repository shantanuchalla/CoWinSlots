package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
)

var (
	blankColor         = &tablewriter.Colors{}
	noSlotColor        = &tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.Bold, tablewriter.BgHiRedColor}
	availableSlotColor = &tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold, tablewriter.BgHiGreenColor}
)

func (checker CowinSlotChecker) processSlots(req contracts.SlotRequest) error {
	resp, err := checker.CowinClient.CallCoWinAPI(req.Location.DistrictId, req.Date)
	if err != nil {
		return err
	}

	today := getDateString(time.Now())
	t1 := getDateString(time.Now().AddDate(0, 0, 1))
	t2 := getDateString(time.Now().AddDate(0, 0, 2))
	t3 := getDateString(time.Now().AddDate(0, 0, 3))
	t4 := getDateString(time.Now().AddDate(0, 0, 4))
	t5 := getDateString(time.Now().AddDate(0, 0, 5))
	t6 := getDateString(time.Now().AddDate(0, 0, 6))
	dateArray := []string{today, t1, t2, t3, t4, t5, t6}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "PIN", today, t1, t2, t3, t4, t5, t6})
	table.SetBorder(true)
	table.SetRowLine(true)

	totalAvailableCentres := 0
	totalAvailableSlots := 0
	totalAvailableCapacity := float64(0)

	vaccineCentreCount := make(map[string]int)
	vaccineSlotCount := make(map[string]int)
	vaccineCapacityCount := make(map[string]float64)

	for _, centre := range resp.Centres {
		availability := make(map[string]*contracts.AvailData, 7)
		vaccineCentreAvailability := make(map[string]bool)
		write := false
		available := false

		for _, session := range centre.Sessions {
			if session.AgeLimit == 18 {
				write = true
				vaccineSlotCount[session.Vaccine]++
				vaccineCentreAvailability[session.Vaccine] = true
				totalAvailableSlots++
				color := new(tablewriter.Colors)

				if session.Capacity > float64(0) {
					totalAvailableCapacity += session.Capacity
					vaccineCapacityCount[session.Vaccine] += session.Capacity
					color = availableSlotColor
					available = true
				} else {
					color = noSlotColor
				}

				availability[session.Date] = &contracts.AvailData{
					Available: strconv.Itoa(int(session.Capacity)),
					Vaccine:   session.Vaccine,
					Color:     *color,
				}
			}
		}
		if write {
			totalAvailableCentres++
			for vaccine, avail := range vaccineCentreAvailability {
				if avail {
					vaccineCentreCount[vaccine]++
				}
			}
			appendTableRow(table, availability, centre, dateArray, available)
		}
	}

	table.Append([]string{"", "", "", "", "", "", "", "", ""})

	for vaccine, slotCount := range vaccineSlotCount {
		table.Append([]string{vaccine, "Centres", strconv.Itoa(vaccineCentreCount[vaccine]),
			"Slots", strconv.Itoa(slotCount),
			"Capacity", fmt.Sprintf("%f", vaccineCapacityCount[vaccine]),
			"", ""})
	}

	table.SetFooter([]string{"Total",
		"Centres", strconv.Itoa(totalAvailableCentres),
		"Slots", strconv.Itoa(totalAvailableSlots),
		"Capacity", fmt.Sprintf("%f", totalAvailableCapacity),
		"", ""})

	if totalAvailableSlots > 0 {
		table.Render()
	}
	return nil
}

func appendTableRow(table *tablewriter.Table, availability map[string]*contracts.AvailData, centre *contracts.Centre, dateArray []string, available bool) {
	row := make([]string, 0)
	details := make([]tablewriter.Colors, 0)

	row = append(row, centre.Name)
	row = append(row, strconv.Itoa(centre.Pin))
	details = append(details, *blankColor)
	if available {
		details = append(details, *availableSlotColor)
	} else {
		details = append(details, *blankColor)
	}

	for _, day := range dateArray {
		avail := availability[day]
		if nil == avail {
			row = append(row, "")
			details = append(details, *blankColor)
		} else {
			row = append(row, avail.Vaccine+" @ "+avail.Available)
			details = append(details, avail.Color)
		}
	}
	table.Rich(row, details)
}

func getDateString(time time.Time) string {
	year, month, day := time.Date()
	return fmt.Sprintf("%02d", day) + "-" + fmt.Sprintf("%02d", int(month)) + "-" + fmt.Sprintf("%04d", year)
}
