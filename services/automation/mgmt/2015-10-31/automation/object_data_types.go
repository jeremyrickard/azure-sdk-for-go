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

// ObjectDataTypesClient is the automation Client
type ObjectDataTypesClient struct {
	ManagementClient
}

// NewObjectDataTypesClient creates an instance of the ObjectDataTypesClient client.
func NewObjectDataTypesClient(p pipeline.Pipeline) ObjectDataTypesClient {
	return ObjectDataTypesClient{NewManagementClient(p)}
}

// ListFieldsByModuleAndType retrieve a list of fields of a given type identified by module name.
//
// automationAccountName is the automation account name. moduleName is the name of module. typeName is the name of
// type.
func (client ObjectDataTypesClient) ListFieldsByModuleAndType(ctx context.Context, automationAccountName string, moduleName string, typeName string) (*TypeFieldListResult, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.listFieldsByModuleAndTypePreparer(automationAccountName, moduleName, typeName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listFieldsByModuleAndTypeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TypeFieldListResult), err
}

// listFieldsByModuleAndTypePreparer prepares the ListFieldsByModuleAndType request.
func (client ObjectDataTypesClient) listFieldsByModuleAndTypePreparer(automationAccountName string, moduleName string, typeName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/objectDataTypes/{typeName}/fields"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listFieldsByModuleAndTypeResponder handles the response to the ListFieldsByModuleAndType request.
func (client ObjectDataTypesClient) listFieldsByModuleAndTypeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &TypeFieldListResult{rawResponse: resp.Response()}
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

// ListFieldsByType retrieve a list of fields of a given type across all accessible modules.
//
// automationAccountName is the automation account name. typeName is the name of type.
func (client ObjectDataTypesClient) ListFieldsByType(ctx context.Context, automationAccountName string, typeName string) (*TypeFieldListResult, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.listFieldsByTypePreparer(automationAccountName, typeName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listFieldsByTypeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TypeFieldListResult), err
}

// listFieldsByTypePreparer prepares the ListFieldsByType request.
func (client ObjectDataTypesClient) listFieldsByTypePreparer(automationAccountName string, typeName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/objectDataTypes/{typeName}/fields"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listFieldsByTypeResponder handles the response to the ListFieldsByType request.
func (client ObjectDataTypesClient) listFieldsByTypeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &TypeFieldListResult{rawResponse: resp.Response()}
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
