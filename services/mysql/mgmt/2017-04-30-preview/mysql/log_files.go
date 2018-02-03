package mysql

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

// LogFilesClient is the the Microsoft Azure management API provides create, read, update, and delete functionality for
// Azure MySQL resources including servers, databases, firewall rules, log files and configurations.
type LogFilesClient struct {
	ManagementClient
}

// NewLogFilesClient creates an instance of the LogFilesClient client.
func NewLogFilesClient(p pipeline.Pipeline) LogFilesClient {
	return LogFilesClient{NewManagementClient(p)}
}

// ListByServer list all the log files in a given server.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server.
func (client LogFilesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (*LogFileListResult, error) {
	req, err := client.listByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByServerResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LogFileListResult), err
}

// listByServerPreparer prepares the ListByServer request.
func (client LogFilesClient) listByServerPreparer(resourceGroupName string, serverName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/logFiles"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByServerResponder handles the response to the ListByServer request.
func (client LogFilesClient) listByServerResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LogFileListResult{rawResponse: resp.Response()}
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
