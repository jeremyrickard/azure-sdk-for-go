package automation

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

// TestJobStreamsClient is the automation Client
type TestJobStreamsClient struct {
	ManagementClient
}

// NewTestJobStreamsClient creates an instance of the TestJobStreamsClient client.
func NewTestJobStreamsClient(p pipeline.Pipeline) TestJobStreamsClient {
	return TestJobStreamsClient{NewManagementClient(p)}
}

// Get retrieve a test job streams identified by runbook name and stream id.
//
// automationAccountName is the automation account name. runbookName is the runbook name. jobStreamID is the job stream
// id.
func (client TestJobStreamsClient) Get(ctx context.Context, automationAccountName string, runbookName string, jobStreamID string) (*JobStream, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPreparer(automationAccountName, runbookName, jobStreamID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*JobStream), err
}

// getPreparer prepares the Get request.
func (client TestJobStreamsClient) getPreparer(automationAccountName string, runbookName string, jobStreamID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client TestJobStreamsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &JobStream{rawResponse: resp.Response()}
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

// ListByTestJob retrieve a list of test job streams identified by runbook name.
//
// automationAccountName is the automation account name. runbookName is the runbook name. filter is the filter to apply
// on the operation.
func (client TestJobStreamsClient) ListByTestJob(ctx context.Context, automationAccountName string, runbookName string, filter *string) (*JobStreamListResult, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.listByTestJobPreparer(automationAccountName, runbookName, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByTestJobResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*JobStreamListResult), err
}

// listByTestJobPreparer prepares the ListByTestJob request.
func (client TestJobStreamsClient) listByTestJobPreparer(automationAccountName string, runbookName string, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if filter != nil {
		params.Set("$filter", *filter)
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByTestJobResponder handles the response to the ListByTestJob request.
func (client TestJobStreamsClient) listByTestJobResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &JobStreamListResult{rawResponse: resp.Response()}
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
