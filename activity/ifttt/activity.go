// Package ifttt provides connectivity to IFTTT for TIBCO Cloud Integration - Web Integrator
// using the WebHooks service from IFTTT (https://ifttt.com/maker_webhooks)
package ifttt

// Imports
import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivConnection = "iftttConnection"
	ivValue1     = "value1"
	ivValue2     = "value2"
	ivValue3     = "value3"
	ovResult     = "result"
)

// Payload is used to describe the payload to IFTTT with a maximum of
// three values (this limit is set by IFTTT)
type Payload struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

// activityLog is the default logger for this class
var activityLog = logger.GetLogger("ifttt")

// IFTTTActivity describes the metadata of the activity as found in the activity.json file
type IFTTTActivity struct {
	metadata *activity.Metadata
}

// NewActivity will instantiate a new IFTTTActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &IFTTTActivity{metadata: metadata}
}

// Metadata will return the metadata of the IFTTTActivity
func (a *IFTTTActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval executes the activity and sends a message to IFTTT
func (a *IFTTTActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing IFTTT WebHook")

	// Validates that the connection has been set. The connection is mandatory
	// whereas the three values are optional
	if context.GetInput(ivConnection) == nil {
		return false, activity.NewError("IFTTT connection is not configured", "IFTTT-2000", nil)
	}

	// Get the connection details
	connectionInfo := context.GetInput(ivConnection).(map[string]interface{})
	connectionSettings := connectionInfo["settings"].([]interface{})

	// Build the IFTTT WebHook URL. To trigger the event it will make a POST request to the URL
	var webhookKey string
	var eventName string

	for _, v := range connectionSettings {
		setting := v.(map[string]interface{})
		if setting["name"] == "webhookKey" {
			webhookKey = setting["value"].(string)
		} else if setting["name"] == "eventName" {
			eventName = setting["value"].(string)
		}
	}

	list := []string{"https://maker.ifttt.com/trigger/", eventName, "/with/key/", webhookKey}
	var str bytes.Buffer

	for _, l := range list {
		str.WriteString(l)
	}

	var iftttURL = str.String()
	activityLog.Info("The WebHook URL is set to " + iftttURL)

	// Create JSON payload. The data is completely optional, and you can also pass value1, value2, and value3 as query parameters or form variables. This content will be passed on to the Action in your Recipe.
	var data Payload

	if context.GetInput(ivValue3) == nil {
		data = Payload{Value1: context.GetInput(ivValue1).(string), Value2: context.GetInput(ivValue2).(string)}
	} else if context.GetInput(ivValue2) == nil {
		data = Payload{Value1: context.GetInput(ivValue1).(string)}
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return false, activity.NewError("Error occured while creating JSON Payload", "IFTTT-2001", nil)
	}
	body := bytes.NewReader(payloadBytes)

	// Send the POST message
	req, err := http.NewRequest("POST", iftttURL, body)
	if err != nil {
		return false, activity.NewError("Error occured sending message to IFTTT", "IFTTT-2002", nil)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, activity.NewError("Error occured receiving response from IFTTT", "IFTTT-2003", nil)
	}
	defer resp.Body.Close()

	activityLog.Info("WebHook returned with StatusCode " + strconv.Itoa(resp.StatusCode))

	// Set the return value
	context.SetOutput(ovResult, strconv.Itoa(resp.StatusCode))
	return true, nil
}
