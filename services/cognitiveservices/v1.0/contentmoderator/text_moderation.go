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
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
	"net/url"
)

// TextModerationClient is the you use the API to scan your content as it is generated. Content Moderator then
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
type TextModerationClient struct {
	ManagementClient
}

// NewTextModerationClient creates an instance of the TextModerationClient client.
func NewTextModerationClient(url url.URL, p pipeline.Pipeline) TextModerationClient {
	return TextModerationClient{NewManagementClient(url, p)}
}

// DetectLanguage this operation will detect the language of given input content. Returns the <a
// href="http://www-01.sil.org/iso639-3/codes.asp">ISO 639-3 code</a> for the predominant language comprising the
// submitted text. Over 110 languages supported.
//
// textContentType is the content type. textContent is content to screen.
func (client TextModerationClient) DetectLanguage(ctx context.Context, textContentType string, textContent string) (*DetectedLanguage, error) {
	req, err := client.detectLanguagePreparer(textContentType, textContent)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.detectLanguageResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DetectedLanguage), err
}

// detectLanguagePreparer prepares the DetectLanguage request.
func (client TextModerationClient) detectLanguagePreparer(textContentType string, textContent string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("POST", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", fmt.Sprintf("%v", textContentType))
	b, err := json.Marshal(textContent)
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

// detectLanguageResponder handles the response to the DetectLanguage request.
func (client TextModerationClient) detectLanguageResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &DetectedLanguage{rawResponse: resp.Response()}
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

// ScreenText detects profanity in more than 100 languages and match against custom and shared blacklists.
//
// language is language of the terms. textContentType is the content type. textContent is content to screen.
// autocorrect is autocorrect text. pii is detect personal identifiable information. listID is the list Id. classify is
// classify input.
func (client TextModerationClient) ScreenText(ctx context.Context, language string, textContentType string, textContent string, autocorrect *bool, pii *bool, listID *string, classify *bool) (*Screen, error) {
	req, err := client.screenTextPreparer(language, textContentType, textContent, autocorrect, pii, listID, classify)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.screenTextResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Screen), err
}

// screenTextPreparer prepares the ScreenText request.
func (client TextModerationClient) screenTextPreparer(language string, textContentType string, textContent string, autocorrect *bool, pii *bool, listID *string, classify *bool) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("POST", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("language", language)
	if autocorrect != nil {
		params.Set("autocorrect", fmt.Sprintf("%v", *autocorrect))
	}
	if pii != nil {
		params.Set("PII", fmt.Sprintf("%v", *pii))
	}
	if listID != nil {
		params.Set("listId", *listID)
	}
	if classify != nil {
		params.Set("classify", fmt.Sprintf("%v", *classify))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", fmt.Sprintf("%v", textContentType))
	b, err := json.Marshal(textContent)
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

// screenTextResponder handles the response to the ScreenText request.
func (client TextModerationClient) screenTextResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Screen{rawResponse: resp.Response()}
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
