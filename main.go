package main

import (
	"encoding/xml"
	"net/http"
)

type MySOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MySOAPBody
}

type MySOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	MyAction MySOAPAction
}

type MySOAPAction struct {
	XMLName xml.Name `xml:"my_action"`
	Value   string   `xml:"value"`
}

func soapHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Decode the incoming SOAP message
	decoder := xml.NewDecoder(r.Body)
	var envelope MySOAPEnvelope
	err := decoder.Decode(&envelope)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a SOAP response
	response := MySOAPEnvelope{
		Body: MySOAPBody{
			MyAction: MySOAPAction{
				Value: "Hello, " + envelope.Body.MyAction.Value,
			},
		},
	}

	// Encode and send the SOAP response
	w.Header().Add("Content-Type", "text/xml")
	encoder := xml.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/soap", soapHandler)
	http.ListenAndServe(":8080", nil)
}
