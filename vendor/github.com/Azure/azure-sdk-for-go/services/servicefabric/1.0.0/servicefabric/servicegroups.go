package servicefabric

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ServiceGroupsClient is the client for the ServiceGroups methods of the Servicefabric service.
type ServiceGroupsClient struct {
	BaseClient
}

// NewServiceGroupsClient creates an instance of the ServiceGroupsClient client.
func NewServiceGroupsClient(timeout *int32) ServiceGroupsClient {
	return NewServiceGroupsClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewServiceGroupsClientWithBaseURI creates an instance of the ServiceGroupsClient client.
func NewServiceGroupsClientWithBaseURI(baseURI string, timeout *int32) ServiceGroupsClient {
	return ServiceGroupsClient{NewWithBaseURI(baseURI, timeout)}
}

// Create create service groups
//
// applicationName is the name of the service group createServiceGroupDescription is the description of the service
// group
func (client ServiceGroupsClient) Create(ctx context.Context, applicationName string, createServiceGroupDescription BasicCreateServiceGroupDescription) (result String, err error) {
	req, err := client.CreatePreparer(ctx, applicationName, createServiceGroupDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ServiceGroupsClient) CreatePreparer(ctx context.Context, applicationName string, createServiceGroupDescription BasicCreateServiceGroupDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/GetServices/$/CreateServiceGroup", pathParameters),
		autorest.WithJSON(createServiceGroupDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceGroupsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServiceGroupsClient) CreateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove remove service groups
//
// applicationName is the name of the application serviceName is the name of the service
func (client ServiceGroupsClient) Remove(ctx context.Context, applicationName string, serviceName string) (result String, err error) {
	req, err := client.RemovePreparer(ctx, applicationName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client ServiceGroupsClient) RemovePreparer(ctx context.Context, applicationName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
		"serviceName":     serviceName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/GetServiceGroups/{serviceName}/$/Delete", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceGroupsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client ServiceGroupsClient) RemoveResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update service groups
//
// applicationName is the name of the application serviceName is the name of the service
// updateServiceGroupDescription is the description of the service group update
func (client ServiceGroupsClient) Update(ctx context.Context, applicationName string, serviceName string, updateServiceGroupDescription BasicUpdateServiceGroupDescription) (result String, err error) {
	req, err := client.UpdatePreparer(ctx, applicationName, serviceName, updateServiceGroupDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServiceGroupsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServiceGroupsClient) UpdatePreparer(ctx context.Context, applicationName string, serviceName string, updateServiceGroupDescription BasicUpdateServiceGroupDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
		"serviceName":     serviceName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/GetServices/{serviceName}/$/UpdateServiceGroup", pathParameters),
		autorest.WithJSON(updateServiceGroupDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServiceGroupsClient) UpdateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
