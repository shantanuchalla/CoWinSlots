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
		Locations: []*contracts.Location{
			{State: "Karnataka", City: "BBMP"},
			{State: "Karnataka", City: "Bangalore Rural"},
			{State: "Karnataka", City: "Bangalore Urban"},
			{State: "Telangana", City: "Hyderabad"},
			{State: "Telangana", City: "Rangareddy"},
			{State: "Telangana", City: "Medchal"},
		},
		PollInterval: 10 * time.Minute,
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
