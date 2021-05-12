package client

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
)

type CowinClinet struct {
	Client        *http.Client
	BaseUrl       string
	Authorization string
}

func (client CowinClinet) CallCoWinAPI(districtId, date string) (*contracts.SlotResponse, error) {
	requestURL := client.getAvailableSlotRequestURL(districtId, date)

	//optionsReq, err := http.NewRequest(http.MethodOptions, requestURL, nil)
	//if err != nil {
	//	log.Error().Err(err).Msg("error creating request")
	//	return nil, err
	//}
	//optionsResponse, err := client.Client.Do(optionsReq)
	//if err != nil {
	//	log.Error().Err(err).Msg("error calling CoWin API")
	//	return nil, err
	//}
	//defer optionsResponse.Body.Close()
	//
	//bytes, err := ioutil.ReadAll(optionsResponse.Body)
	//if err != nil {
	//	log.Error().Err(err).Msg("error reading response CoWin options API")
	//	return nil, err
	//}
	//
	//log.Info().Msgf("%+v  |  %+v", string(bytes), optionsResponse.Header)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Error().Err(err).Msg("error creating request")
		return nil, err
	}
	//req.Header.Set("Authorization", client.Authorization)

	response, err := client.Client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error calling CoWin API")
		return nil, err
	}
	defer response.Body.Close()

	var slotResponse = new(contracts.SlotResponse)
	err = json.NewDecoder(response.Body).Decode(slotResponse)
	if err != nil {
		log.Error().Err(err).Msg("error decoding response")
		return nil, err
	}

	return slotResponse, nil
}

func (client CowinClinet) getAvailableSlotRequestURL(districtId, date string) string {
	return client.BaseUrl + "/api/v2/appointment/sessions/calendarByDistrict?district_id=" + districtId + "&date=" + date
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
