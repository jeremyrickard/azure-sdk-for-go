package contentmoderator

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
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ListManagementTermListsClient is the you use the API to scan your content as it is generated. Content Moderator then
// processes your content and sends the results along with relevant information either back to your systems or to the
// built-in review tool. You can use this information to take decisions e.g. take it down, send to human judge, etc.
//
// When using the API, images need to have a minimum of 128 pixels and a maximum file size of 4MB.
// Text can be at most 1024 characters long.
// If the content passed to the text API or the image API exceeds the size limits, the API will return an error code
// that informs about the issue.
//
// This API is currently available in:
//
// * West US - westus.api.cognitive.microsoft.com
// * East US 2 - eastus2.api.cognitive.microsoft.com
// * West Central US - westcentralus.api.cognitive.microsoft.com
// * West Europe - westeurope.api.cognitive.microsoft.com
// * Southeast Asia - southeastasia.api.cognitive.microsoft.com .
type ListManagementTermListsClient struct {
	ManagementClient
}

// NewListManagementTermListsClient creates an instance of the ListManagementTermListsClient client.
func NewListManagementTermListsClient(url url.URL, p pipeline.Pipeline) ListManagementTermListsClient {
	return ListManagementTermListsClient{NewManagementClient(url, p)}
}

// Create creates a Term List
//
// contentType is the content type. body is schema of the body.
func (client ListManagementTermListsClient) Create(ctx context.Context, contentType string, body Body) (*TermList, error) {
	req, err := client.createPreparer(contentType, body)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TermList), err
}

// createPreparer prepares the Create request.
func (client ListManagementTermListsClient) createPreparer(contentType string, body Body) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("POST", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", contentType)
	b, err := json.Marshal(body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createResponder handles the response to the Create request.
func (client ListManagementTermListsClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &TermList{rawResponse: resp.Response()}
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

// Delete deletes term list with the list Id equal to list Id passed.
//
// listID is list Id of the image list.
func (client ListManagementTermListsClient) Delete(ctx context.Context, listID string) (*DeleteResponse, error) {
	req, err := client.deletePreparer(listID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DeleteResponse), err
}

// deletePreparer prepares the Delete request.
func (client ListManagementTermListsClient) deletePreparer(listID string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("DELETE", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ListManagementTermListsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &DeleteResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetAllTermLists gets all the Term Lists
func (client ListManagementTermListsClient) GetAllTermLists(ctx context.Context) (*GetAllTermListsResponse, error) {
	req, err := client.getAllTermListsPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getAllTermListsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetAllTermListsResponse), err
}

// getAllTermListsPreparer prepares the GetAllTermLists request.
func (client ListManagementTermListsClient) getAllTermListsPreparer() (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getAllTermListsResponder handles the response to the GetAllTermLists request.
func (client ListManagementTermListsClient) getAllTermListsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetAllTermListsResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetDetails returns list Id details of the term list with list Id equal to list Id passed.
//
// listID is list Id of the image list.
func (client ListManagementTermListsClient) GetDetails(ctx context.Context, listID string) (*TermList, error) {
	req, err := client.getDetailsPreparer(listID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getDetailsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TermList), err
}

// getDetailsPreparer prepares the GetDetails request.
func (client ListManagementTermListsClient) getDetailsPreparer(listID string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getDetailsResponder handles the response to the GetDetails request.
func (client ListManagementTermListsClient) getDetailsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &TermList{rawResponse: resp.Response()}
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

// RefreshIndexMethod refreshes the index of the list with list Id equal to list ID passed.
//
// listID is list Id of the image list. language is language of the terms.
func (client ListManagementTermListsClient) RefreshIndexMethod(ctx context.Context, listID string, language string) (*RefreshIndex, error) {
	req, err := client.refreshIndexMethodPreparer(listID, language)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.refreshIndexMethodResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RefreshIndex), err
}

// refreshIndexMethodPreparer prepares the RefreshIndexMethod request.
func (client ListManagementTermListsClient) refreshIndexMethodPreparer(listID string, language string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("POST", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("language", language)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// refreshIndexMethodResponder handles the response to the RefreshIndexMethod request.
func (client ListManagementTermListsClient) refreshIndexMethodResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RefreshIndex{rawResponse: resp.Response()}
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

// Update updates an Term List.
//
// listID is list Id of the image list. contentType is the content type. body is schema of the body.
func (client ListManagementTermListsClient) Update(ctx context.Context, listID string, contentType string, body Body) (*TermList, error) {
	req, err := client.updatePreparer(listID, contentType, body)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*TermList), err
}

// updatePreparer prepares the Update request.
func (client ListManagementTermListsClient) updatePreparer(listID string, contentType string, body Body) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", contentType)
	b, err := json.Marshal(body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateResponder handles the response to the Update request.
func (client ListManagementTermListsClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &TermList{rawResponse: resp.Response()}
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
