package main

import (
	"github.com/shantanuchalla/awesomeProject/pkg/contracts"
	"github.com/shantanuchalla/awesomeProject/pkg/service"
	"net/http"
	"time"

	"github.com/shantanuchalla/awesomeProject/pkg/client"
)

const (
	baseUrl       = "https://cdn-api.co-vin.in"
	timeout       = 1 * time.Second
	interval      = 10 * time.Second
	authorization = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJlZWM5Zjk2Mi01MzI2LTQ4OWEtOGQzYS01MTYxOTQ2M2VlNGYiLCJ1c2VyX2lkIjoiZWVjOWY5NjItNTMyNi00ODlhLThkM2EtNTE2MTk0NjNlZTRmIiwidXNlcl90eXBlIjoiQkVORUZJQ0lBUlkiLCJtb2JpbGVfbnVtYmVyIjo5OTA4NjAyMjI3LCJiZW5lZmljaWFyeV9yZWZlcmVuY2VfaWQiOjUxOTQzMzg1Nzk0NzIwLCJzZWNyZXRfa2V5IjoiYjVjYWIxNjctNzk3Ny00ZGYxLTgwMjctYTYzYWExNDRmMDRlIiwidWEiOiJNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF8xNV83KSBBcHBsZVdlYktpdC82MDUuMS4xNSAoS0hUTUwsIGxpa2UgR2Vja28pIFZlcnNpb24vMTQuMSBTYWZhcmkvNjA1LjEuMTUiLCJkYXRlX21vZGlmaWVkIjoiMjAyMS0wNS0wN1QxNzoyMTozNi41OTJaIiwiaWF0IjoxNjIwNDA4MDk2LCJleHAiOjE2MjA0MDg5OTZ9.WfAqlH6Ih-fUaB3Khvdgs7qvgdmZXJX6GWb_QNx64aI"
)

func main() {

	cowinClient := &client.CowinClinet{
		Client:        &http.Client{Timeout: timeout, Transport: &client.MyTransport{}},
		BaseUrl:       baseUrl,
		Authorization: authorization,
	}

	checkService := &service.CowinSlotChecker{
		CowinClient: cowinClient,
		Locations: []*contracts.Location{
			{State: "Karnataka", City: "BBMP", DistrictId: "294"},
			{State: "Karnataka", City: "Bangalore Rural", DistrictId: "276"},
			{State: "Karnataka", City: "Bangalore Urban", DistrictId: "265"},
			//{State: "Telangana", City: "Hyderabad"},
			//{State: "Telangana", City: "Rangareddy"},
			//{State: "Telangana", City: "Medchal"},
		},
		PollInterval: interval,
	}

	go checkService.InitSlotPoller(false)
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
