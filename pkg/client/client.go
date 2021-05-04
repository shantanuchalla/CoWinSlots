package client

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
)

type CowinClinet struct {
	Client  *http.Client
	BaseUrl string
}

func (client CowinClinet) CallCoWinAPI(districtId, date string) (*contracts.SlotResponse, error) {
	request := client.getHttpRequest(districtId, date)
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

func (client CowinClinet) getHttpRequest(districtId, date string) string {
	return client.BaseUrl + "/api/v2/appointment/sessions/public/calendarByDistrict?district_id=" + districtId + "&date=" + date
}
