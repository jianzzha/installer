package resources

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DeploymentOperationsGroupClient is the provides operations for working with resources and resource groups.
type DeploymentOperationsGroupClient struct {
	BaseClient
}

// NewDeploymentOperationsGroupClient creates an instance of the DeploymentOperationsGroupClient client.
func NewDeploymentOperationsGroupClient(subscriptionID string) DeploymentOperationsGroupClient {
	return NewDeploymentOperationsGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeploymentOperationsGroupClientWithBaseURI creates an instance of the DeploymentOperationsGroupClient client.
func NewDeploymentOperationsGroupClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsGroupClient {
	return DeploymentOperationsGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a deployments operation.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// deploymentName - the name of the deployment.
// operationID - the ID of the operation to get.
func (client DeploymentOperationsGroupClient) Get(ctx context.Context, resourceGroupName string, deploymentName string, operationID string) (result DeploymentOperation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentOperationsGroupClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: deploymentName,
			Constraints: []validation.Constraint{{Target: "deploymentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "deploymentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "deploymentName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentOperationsGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, deploymentName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeploymentOperationsGroupClient) GetPreparer(ctx context.Context, resourceGroupName string, deploymentName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsGroupClient) GetResponder(resp *http.Response) (result DeploymentOperation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAtSubscriptionScope gets a deployments operation.
// Parameters:
// deploymentName - the name of the deployment.
// operationID - the ID of the operation to get.
func (client DeploymentOperationsGroupClient) GetAtSubscriptionScope(ctx context.Context, deploymentName string, operationID string) (result DeploymentOperation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentOperationsGroupClient.GetAtSubscriptionScope")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deploymentName,
			Constraints: []validation.Constraint{{Target: "deploymentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "deploymentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "deploymentName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentOperationsGroupClient", "GetAtSubscriptionScope", err.Error())
	}

	req, err := client.GetAtSubscriptionScopePreparer(ctx, deploymentName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "GetAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtSubscriptionScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "GetAtSubscriptionScope", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtSubscriptionScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "GetAtSubscriptionScope", resp, "Failure responding to request")
	}

	return
}

// GetAtSubscriptionScopePreparer prepares the GetAtSubscriptionScope request.
func (client DeploymentOperationsGroupClient) GetAtSubscriptionScopePreparer(ctx context.Context, deploymentName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName": autorest.Encode("path", deploymentName),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAtSubscriptionScopeSender sends the GetAtSubscriptionScope request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsGroupClient) GetAtSubscriptionScopeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetAtSubscriptionScopeResponder handles the response to the GetAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsGroupClient) GetAtSubscriptionScopeResponder(resp *http.Response) (result DeploymentOperation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all deployments operations for a deployment.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// deploymentName - the name of the deployment with the operation to get.
// top - the number of results to return.
func (client DeploymentOperationsGroupClient) List(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result DeploymentOperationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentOperationsGroupClient.List")
		defer func() {
			sc := -1
			if result.dolr.Response.Response != nil {
				sc = result.dolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: deploymentName,
			Constraints: []validation.Constraint{{Target: "deploymentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "deploymentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "deploymentName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentOperationsGroupClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, deploymentName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.dolr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DeploymentOperationsGroupClient) ListPreparer(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName":    autorest.Encode("path", deploymentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsGroupClient) ListResponder(resp *http.Response) (result DeploymentOperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeploymentOperationsGroupClient) listNextResults(ctx context.Context, lastResults DeploymentOperationsListResult) (result DeploymentOperationsListResult, err error) {
	req, err := lastResults.deploymentOperationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeploymentOperationsGroupClient) ListComplete(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result DeploymentOperationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentOperationsGroupClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, deploymentName, top)
	return
}

// ListAtSubscriptionScope gets all deployments operations for a deployment.
// Parameters:
// deploymentName - the name of the deployment with the operation to get.
// top - the number of results to return.
func (client DeploymentOperationsGroupClient) ListAtSubscriptionScope(ctx context.Context, deploymentName string, top *int32) (result DeploymentOperationsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentOperationsGroupClient.ListAtSubscriptionScope")
		defer func() {
			sc := -1
			if result.dolr.Response.Response != nil {
				sc = result.dolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deploymentName,
			Constraints: []validation.Constraint{{Target: "deploymentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "deploymentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "deploymentName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("resources.DeploymentOperationsGroupClient", "ListAtSubscriptionScope", err.Error())
	}

	result.fn = client.listAtSubscriptionScopeNextResults
	req, err := client.ListAtSubscriptionScopePreparer(ctx, deploymentName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "ListAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAtSubscriptionScopeSender(req)
	if err != nil {
		result.dolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "ListAtSubscriptionScope", resp, "Failure sending request")
		return
	}

	result.dolr, err = client.ListAtSubscriptionScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "ListAtSubscriptionScope", resp, "Failure responding to request")
	}

	return
}

// ListAtSubscriptionScopePreparer prepares the ListAtSubscriptionScope request.
func (client DeploymentOperationsGroupClient) ListAtSubscriptionScopePreparer(ctx context.Context, deploymentName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deploymentName": autorest.Encode("path", deploymentName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAtSubscriptionScopeSender sends the ListAtSubscriptionScope request. The method will close the
// http.Response Body if it receives an error.
func (client DeploymentOperationsGroupClient) ListAtSubscriptionScopeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListAtSubscriptionScopeResponder handles the response to the ListAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (client DeploymentOperationsGroupClient) ListAtSubscriptionScopeResponder(resp *http.Response) (result DeploymentOperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAtSubscriptionScopeNextResults retrieves the next set of results, if any.
func (client DeploymentOperationsGroupClient) listAtSubscriptionScopeNextResults(ctx context.Context, lastResults DeploymentOperationsListResult) (result DeploymentOperationsListResult, err error) {
	req, err := lastResults.deploymentOperationsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "listAtSubscriptionScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAtSubscriptionScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "listAtSubscriptionScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAtSubscriptionScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.DeploymentOperationsGroupClient", "listAtSubscriptionScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAtSubscriptionScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeploymentOperationsGroupClient) ListAtSubscriptionScopeComplete(ctx context.Context, deploymentName string, top *int32) (result DeploymentOperationsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeploymentOperationsGroupClient.ListAtSubscriptionScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAtSubscriptionScope(ctx, deploymentName, top)
	return
}
