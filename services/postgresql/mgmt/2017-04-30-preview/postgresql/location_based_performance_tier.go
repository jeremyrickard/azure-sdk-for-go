package postgresql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// LocationBasedPerformanceTierClient is the the Microsoft Azure management API provides create, read, update, and
// delete functionality for Azure PostgreSQL resources including servers, databases, firewall rules, log files and
// configurations.
type LocationBasedPerformanceTierClient struct {
	ManagementClient
}

// NewLocationBasedPerformanceTierClient creates an instance of the LocationBasedPerformanceTierClient client.
func NewLocationBasedPerformanceTierClient(p pipeline.Pipeline) LocationBasedPerformanceTierClient {
	return LocationBasedPerformanceTierClient{NewManagementClient(p)}
}

// List list all the performance tiers at specified location in a given subscription.
//
// locationName is the name of the location.
func (client LocationBasedPerformanceTierClient) List(ctx context.Context, locationName string) (*PerformanceTierListResult, error) {
	req, err := client.listPreparer(locationName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PerformanceTierListResult), err
}

// listPreparer prepares the List request.
func (client LocationBasedPerformanceTierClient) listPreparer(locationName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/performanceTiers"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client LocationBasedPerformanceTierClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PerformanceTierListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
