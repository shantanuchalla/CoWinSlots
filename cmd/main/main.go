package main

import (
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
	"github.com/shantanuchalla/awesomeProject/pkg/service"
	"net/http"
	"time"

	"github.com/shantanuchalla/awesomeProject/pkg/client"
)

const (
	baseUrl = "https://cdn-api.co-vin.in"
	timeout = 10 * time.Second
)

func main() {
	cowinClient := &client.CowinClinet{
		Client:  &http.Client{Timeout: timeout},
		BaseUrl: baseUrl,
	}

	checkService := &service.CowinSlotChecker{
		CowinClient: cowinClient,
		Location: contracts.Location{
			State:      "Karnataka",
			City:       "BBMP",
			DistrictId: "294",
		},
		PollInterval: 10 * time.Second,
	}

	go checkService.InitSlotPoller()
	checkService.InitPollListener()

	//handler := http.NewServeMux()
	//handler.

	//server := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        server.InitHandler(),
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 0,
	//}
}
