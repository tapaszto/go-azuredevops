// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides functions for validating payloads from GitHub Webhooks.
// GitHub API docs: https://developer.github.com/webhooks/securing/#validating-payloads-from-github

// Adapted for Azure Devops

package azuredevops

import (
	"encoding/json"
	"net/http"
)

const (
	activityIDHeader     = "X-VSS-ActivityId"
	subscriptionIDHeader = "X-VSS-SubscriptionId"
	// requestIDHeader is the Azure Devops header key used to pass the unique ID for the webhook event.
	requestIDHeader = "Request-Id"
)

// GetActivityID returns the value of the activityIDHeader webhook header.
//
// Haven't found vendor documentation yet.  This could be a GUID that identifies
// the webhook request ID.  A different GUID is also present in the body of
// webhook requests.
func GetActivityID(r *http.Request) string {
	return r.Header.Get(activityIDHeader)
}

// GetRequestID returns the value of the requestIDHeader webhook header.
//
// Haven't found vendor documentation yet.  This could be a GUID that identifies
// the webhook request ID.  A different GUID is also present in the body of
// webhook requests.
func GetRequestID(r *http.Request) string {
	return r.Header.Get(requestIDHeader)
}

// GetSubscriptionID returns the value of the subscriptionIDHeader webhook header.
//
// Haven't found vendor documentation yet.  This could be a GUID that identifies
// the webhook event type and settings in the Azure Devops tenant
func GetSubscriptionID(r *http.Request) string {
	return r.Header.Get(subscriptionIDHeader)
}

// ParseWebHook parses the event payload into a corresponding struct.
// An error will be returned for unrecognized event types.
//
// https://docs.microsoft.com/en-us/azure/devops/service-hooks/events?toc=/azure/devops/integrate/toc.json&bc=/azure/devops/integrate/breadcrumb/toc.json&view=azure-devops
//
func ParseWebHook(payload []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(payload, &event)
	if err != nil {
		return nil, err
	}
	if event.EventType != nil {
		_, err = event.ParsePayload()
	}

	return event, err
}
