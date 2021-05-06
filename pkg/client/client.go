package client

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
)

type CowinClinet struct {
	Client  *http.Client
	BaseUrl string
}

func (client CowinClinet) CallCoWinAPI(districtId, date string) (*contracts.SlotResponse, error) {
	request := client.getAvailableSlotRequest(districtId, date)

	response, err := client.Client.Get(request)
	if err != nil {
		log.Error().Err(err).Msg("error calling CoWin API")
		return nil, err
	}

	var slotResponse = new(contracts.SlotResponse)
	err = json.NewDecoder(response.Body).Decode(slotResponse)
	if err != nil {
		log.Error().Err(err).Msg("error decoding response")
		return nil, err
	}

	return slotResponse, nil
}

func (client CowinClinet) getAvailableSlotRequest(districtId, date string) string {
	return client.BaseUrl + "/api/v2/appointment/sessions/public/calendarByDistrict?district_id=" + districtId + "&date=" + date
}

func (client CowinClinet) GetStates() (*contracts.StateResponse, error) {
	response, err := client.Client.Get(client.BaseUrl + "/api/v2/admin/location/states")
	if err != nil {
		log.Error().Err(err).Msg("error calling CoWin API")
		return nil, err
	}

	var stateResponse = new(contracts.StateResponse)
	err = json.NewDecoder(response.Body).Decode(stateResponse)
	if err != nil {
		log.Error().Err(err).Msg("error decoding response")
		return nil, err
	}

	return stateResponse, nil
}

func (client CowinClinet) GetDistricts(stateId int) (*contracts.DistrictResponse, error) {
	response, err := client.Client.Get(client.BaseUrl + "/api/v2/admin/location/districts/" + strconv.Itoa(stateId))
	if err != nil {
		log.Error().Err(err).Msg("error calling CoWin API")
		return nil, err
	}

	var districtResponse = new(contracts.DistrictResponse)
	err = json.NewDecoder(response.Body).Decode(districtResponse)
	if err != nil {
		log.Error().Err(err).Msg("error decoding response")
		return nil, err
	}

	return districtResponse, nil
}
