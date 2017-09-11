// Package ifttt provides connectivity to IFTTT for TIBCO Cloud Integration - Web Integrator
// using the WebHooks service from IFTTT (https://ifttt.com/maker_webhooks)
package ifttt

// Imports
import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

// activityMetadata is the metadata of the activity as described in activity.json
// We'll store it as a variable to reuse it across multiple testcases
var activityMetadata *activity.Metadata

// getActivityMetadata reads the activity.json file and sets the activityMetadata variable
// if the variable already contains metadata it simply returns the current value rather than reading the file again
func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

// TestActivityRegistration checks whether the activity can be registered, and is registered in the engine
func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

// TestEval tests the Eval function and sends a message to IFTTT
// Make sure that you have updated the values below
func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	// Generate a connection object
	connectionData := make(map[string]interface{})
	connectionSettings := make([]interface{}, 2)

	// Add a WebHook Key
	webhookKey := make(map[string]interface{})
	webhookKey["name"] = "webhookKey"
	webhookKey["value"] = "<YOUR WEBHOOK KEY>"
	connectionSettings[0] = webhookKey

	// Add an Event Name
	eventName := make(map[string]interface{})
	eventName["name"] = "eventName"
	eventName["value"] = "<YOUR EVENT>"
	connectionSettings[1] = eventName

	// Add the settings to the connection
	connectionData["settings"] = connectionSettings

	// Specify the input values for the activity
	// These are the constants that start with 'iv'
	tc.SetInput(ivConnection, connectionData)
	tc.SetInput(ivValue1, "<VALUE 1>")
	tc.SetInput(ivValue2, "<VALUE 2>")
	tc.SetInput(ivValue3, "<VALUE 3>")

	// Execute the activity
	_, err := act.Eval(tc)

	// We assume there will be no errors
	assert.Nil(t, err)

	// If you've provided correct details the return value from IFTTT should be a 200
	result := tc.GetOutput(ovResult)
	assert.Equal(t, result, "200")
}
