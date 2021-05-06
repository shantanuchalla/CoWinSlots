package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shantanuchalla/awesomeProject/pkg/client"
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
)

var poll = make(chan contracts.SlotRequest)

type CowinSlotChecker struct {
	Locations []*contracts.Location

	CowinClient  *client.CowinClinet
	PollInterval time.Duration
}

func (checker CowinSlotChecker) InitSlotPoller() {
	log.Info().Msgf("Initializing Poller with interval -- %+v", checker.PollInterval)

	checker.initializeSlotCheck()
	ticker := time.NewTicker(checker.PollInterval)
	tickChannel := ticker.C
	for tick := range tickChannel {
		log.Info().Msgf("received tick <- +%v", tick)
		checker.initializeSlotCheck()
	}
}

func (checker CowinSlotChecker) initializeSlotCheck() {
	for _, loc := range checker.Locations {
		req, err := checker.getSlotRequest(loc)
		if err != nil {
			continue
		}
		poll <- *req
	}
}

func (checker CowinSlotChecker) InitPollListener() {
	for slotRequest := range poll {
		log.Info().Msgf("Results for state -- %s | region -- %s", slotRequest.Location.State, slotRequest.Location.City)
		err := checker.processSlots(slotRequest)
		if err != nil {
			log.Error().Err(err)
		}
	}
}

func (checker CowinSlotChecker) getSlotRequest(loc *contracts.Location) (*contracts.SlotRequest, error) {
	currentTime := time.Now()
	year, month, day := currentTime.Date()

	stateId, err := checker.getStateId(loc)
	if err != nil {
		return nil, err
	}

	districtId, err := checker.funcName(loc, stateId)
	if err != nil {
		return nil, err
	}

	loc.DistrictId = strconv.Itoa(districtId)

	return &contracts.SlotRequest{
		Location: *loc,
		Date:     fmt.Sprintf("%02d", day) + "-" + fmt.Sprintf("%02d", int(month)) + "-" + fmt.Sprintf("%04d", year),
	}, nil
}

func (checker CowinSlotChecker) funcName(loc *contracts.Location, stateId int) (int, error) {
	districts, err := checker.CowinClient.GetDistricts(stateId)
	if err != nil {
		return -1, err
	}

	districtId := -1
	for _, district := range districts.Districts {
		if strings.ToLower(loc.City) == strings.ToLower(district.DistrictName) {
			districtId = district.DistrictId
		}
	}
	if districtId < 0 {
		return -1, errors.New("city not found")
	}
	return districtId, nil
}

func (checker CowinSlotChecker) getStateId(loc *contracts.Location) (int, error) {
	states, err := checker.CowinClient.GetStates()
	if err != nil {
		return -1, err
	}

	stateId := -1
	for _, state := range states.States {
		if strings.ToLower(loc.State) == strings.ToLower(state.StateName) {
			stateId = state.StateId
		}
	}
	if stateId < 0 {
		return -1, errors.New("state not found")
	}
	return stateId, nil
}
