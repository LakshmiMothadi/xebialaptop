package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GetThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type Client struct {
	Hostname   string
	HTTPClient http.Client
}

func NewThingClient(hostname string) *Client {
	return &Client{
		Hostname:   hostname,
		HTTPClient: http.Client{Timeout: 30 * time.Second},
	}
}

func (cl *Client) GetThingOnUUID(thingUUID string) (*GetThingResponse, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/thing/%s", cl.Hostname, thingUUID))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching thing: http-status %d", httpResponse.StatusCode)
	}
	var resp GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type Data map[string]interface{}

func (cl *Client) CreateThing(name string, value string) (*GetThingResponse, error) {
	Data2 := Data{"name": "Lakshmi",
		"value": "XI2471"}

	postData, err := json.Marshal(Data2)
	if err != nil {
		fmt.Println(err)
	}
	responseData := bytes.NewBuffer(postData)

	httpResponse, err := cl.HTTPClient.Post(cl.Hostname+"/thing/new ", "application/json", responseData)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching thing: http-status %d", httpResponse.StatusCode)
	}
	var resp2 GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp2)
	if err != nil {
		return nil, err
	}

	return &resp2, nil
}

type UpdateData struct {
	Value string `json:"value"`
}

func (cl *Client) UpdateThing(uuid string, value string) (GetThingResponse, error) {
	Update := UpdateData{Value: value}

	UpdateData, err := json.Marshal(Update)
	if err != nil {
		return GetThingResponse{}, err
	}
	response := bytes.NewBuffer(UpdateData)
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/thing/%s", cl.Hostname, uuid), response)
	if err != nil {
		return GetThingResponse{}, err
	}

	httpResponse, err := cl.HTTPClient.Do(request)
	if err != nil {
		return GetThingResponse{}, err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return GetThingResponse{}, fmt.Errorf("error updating thing: http-status %d", httpResponse.StatusCode)
	}

	var updated GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&updated)
	if err != nil {
		return GetThingResponse{}, err
	}
	return updated, nil
}

/*
func (cl *Client) DeleteThing(uuid string) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/thing/%s", cl.Hostname, uuid), nil)
	if err != nil {
		return err
	}

	httpResponse, err := cl.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting thing: http-status %d", httpResponse.StatusCode)
	}
	return nil
}
*/

func main() {
	client := NewThingClient("https://api-ldej-nl.el.r.appspot.com")
	resp, err := client.GetThingOnUUID("19b19046-5a1a-43c6-a4fa-6020b0810996")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v", resp)

	resp2, err := client.CreateThing("lakshmi", "done")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v", resp2)
	var updating GetThingResponse
	Update, err := client.UpdateThing(updating.UUID, "updated")
	fmt.Printf("UpdatedThing: %v\n", Update)

}
