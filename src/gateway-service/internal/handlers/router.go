package handlers

import (
	"github.com/gorilla/mux"
)

type ServicesStruct struct {
	TicketServiceAddress string
	FlightServiceAddress string
	BonusServiceAddress  string
}

type GatewayService struct {
	Config ServicesStruct
}

func NewGatewayService(config *ServicesStruct) *GatewayService {
	return &GatewayService{Config: *config}
}

func Router() *mux.Router {
	servicesConfig := ServicesStruct{
		TicketServiceAddress: "http://10.112.128.45:8080",
		FlightServiceAddress: "http://10.112.128.43:8080",
		BonusServiceAddress:  "http://10.112.128.42:8080",
	}

	router := mux.NewRouter()
	gs := NewGatewayService(&servicesConfig)

	router.HandleFunc("/api/v1/flights", gs.GetAllFlights).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/me", gs.GetUserInfo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/tickets", gs.GetUserTickets).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/tickets/{ticketUid}", gs.GetUserTicket).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/tickets", gs.BuyTicket).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/tickets/{ticketUid}", gs.CancelTicket).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/v1/privilege", gs.GetPrivilege).Methods("GET", "OPTIONS")

	return router
}
