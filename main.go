package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type OrderHistory struct {
	OrderId           string    `json:"orderId"`
	BusinessPartnerId string    `json:"businessPartnerId"`
	OrderDate         time.Time `json:"orderDate"`
	DeliveryDate      time.Time `json:"deliveryDate"`
	DeliveryStatus    string    `json: deliveryStatus`
	WasOrderDelayed   bool      `json: wasOrderDelayed`
}

var OrderHistoryItems []OrderHistory

func returnAllOrderHistoryItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: orderhistory")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	urlParameter := r.URL.Query().Get("bpid")

	if urlParameter != "" {

		var foundOrderHistoryItems []OrderHistory

		for _, orderhistoryitem := range OrderHistoryItems {
			if orderhistoryitem.BusinessPartnerId == urlParameter {
				foundOrderHistoryItems = append(foundOrderHistoryItems, orderhistoryitem)
			}
		}

		json.NewEncoder(w).Encode(foundOrderHistoryItems)

	} else {

		json.NewEncoder(w).Encode(OrderHistoryItems)

	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/orderhistory", returnAllOrderHistoryItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	OrderHistoryItems = []OrderHistory{
		{OrderId: "76272720", BusinessPartnerId: "1245783695", OrderDate: time.Date(2020, time.October, 12, 0, 0, 0, 0, time.UTC), DeliveryDate: time.Date(2020, time.November, 05, 0, 0, 0, 0, time.UTC), DeliveryStatus: "C", WasOrderDelayed: true},
		{OrderId: "76272715", BusinessPartnerId: "1245783695", OrderDate: time.Date(2020, time.August, 8, 0, 0, 0, 0, time.UTC), DeliveryDate: time.Date(2020, time.August, 30, 0, 0, 0, 0, time.UTC), DeliveryStatus: "C", WasOrderDelayed: true},
		{OrderId: "76272710", BusinessPartnerId: "1245783695", OrderDate: time.Date(2020, time.July, 18, 0, 0, 0, 0, time.UTC), DeliveryDate: time.Date(2020, time.August, 5, 0, 0, 0, 0, time.UTC), DeliveryStatus: "C", WasOrderDelayed: true},
	}
	handleRequests()
}
