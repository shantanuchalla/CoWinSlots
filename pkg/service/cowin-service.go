package service

import (
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shantanuchalla/awesomeProject/pkg/client"
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
)

var poll = make(chan contracts.SlotRequest)

type CowinSlotChecker struct {
	contracts.Location

	CowinClient  *client.CowinClinet
	PollInterval time.Duration
}

func (checker CowinSlotChecker) InitSlotPoller() {
	log.Info().Msgf("Initializing Poller with interval -- %+v", checker.PollInterval)
	poll <- checker.getSlotRequest()
	ticker := time.NewTicker(checker.PollInterval)
	tickChannel := ticker.C
	for tick := range tickChannel {
		log.Info().Msgf("received tick <- +%v", tick)
		poll <- checker.getSlotRequest()
	}
}

func (checker CowinSlotChecker) InitPollListener() {
	for slotRequest := range poll {
		err := checker.processSlots(slotRequest)
		if err != nil {
			log.Error().Err(err)
		}
	}
}

func (checker CowinSlotChecker) getSlotRequest() contracts.SlotRequest {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	return contracts.SlotRequest{
		Location: checker.Location,
		Date:     strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year),
	}
}
