package rest

import (
	"TicketReservation/src/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Client struct {
	serverURL string
}

func NewClient(serverURL string) Client {
	return Client{
		serverURL: "http://" + serverURL,
	}
}

func (c Client) GetEvents() (events []model.Event, msg string, err error) {
	client := &http.Client{}
	log.SetPrefix("[Client] (getEvents)")

	req, err := http.NewRequest("GET", c.serverURL+apiGetEvents, nil)
	if err != nil {
		log.Println("Error preparing request:", err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Error response status code:", resp.StatusCode)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&events) // TODO: to get msg, we need to define a communication struct
	if err != nil {
		log.Println("Error decoding response body:", err.Error())
		return
	}

	return
}

func (c Client) BookTicket(eventId string, quantity int) (ticketIDs []string, msg string, err error) {
	client := &http.Client{}
	log.SetPrefix("[Client] (bookTicket)")

	reservationData := map[string]interface{}{
		"EventID":    eventId,
		"numTickets": quantity,
	}

	reqBody, err := json.Marshal(reservationData)
	if err != nil {
		log.Println("Error preparing request body:", err.Error())
		return
	}

	req, err := http.NewRequest("POST", c.serverURL+apiSetReservation, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error preparing request:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error response status code:", resp.StatusCode)
		return
	}

	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		log.Println("Error decoding response body:", err.Error())
		return
	}
	ticketIDs = responseData["ticketIDs"].([]string)

	return
}

func (c Client) CreateEvent(name string, date time.Time, totalTickets int) (eventID, msg string, err error) {
	client := &http.Client{}
	log.SetPrefix("[Client] (createEvent)")

	eventData := map[string]interface{}{
		"Name":         name,
		"Date":         date,
		"totalTickets": totalTickets,
	}
	reqBody, err := json.Marshal(eventData)
	if err != nil {
		log.Println("Error preparing request body:", err.Error())
		return
	}

	req, err := http.NewRequest("POST", c.serverURL+apiCreateEvent, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error preparing request:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error response status code:", resp.StatusCode)
		return
	}

	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding response body:", err.Error())
		return
	}
	eventID = responseData["EventID"].(string)

	return
}

func (c Client) BurstTest() (string, error) {
	//TODO implement me
	panic("implement me")
}