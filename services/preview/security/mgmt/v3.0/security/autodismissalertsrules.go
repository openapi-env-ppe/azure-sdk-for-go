package security

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

// AutoDismissAlertsRulesClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AutoDismissAlertsRulesClient struct {
	BaseClient
}

// NewAutoDismissAlertsRulesClient creates an instance of the AutoDismissAlertsRulesClient client.
func NewAutoDismissAlertsRulesClient(subscriptionID string, ascLocation string) AutoDismissAlertsRulesClient {
	return NewAutoDismissAlertsRulesClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAutoDismissAlertsRulesClientWithBaseURI creates an instance of the AutoDismissAlertsRulesClient client.
func NewAutoDismissAlertsRulesClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AutoDismissAlertsRulesClient {
	return AutoDismissAlertsRulesClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Delete delete dismiss alert rule for this subscription.
// Parameters:
// autoDismissAlertsRuleName - the unique name of the auto dismiss alert rule
func (client AutoDismissAlertsRulesClient) Delete(ctx context.Context, autoDismissAlertsRuleName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoDismissAlertsRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AutoDismissAlertsRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, autoDismissAlertsRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AutoDismissAlertsRulesClient) DeletePreparer(ctx context.Context, autoDismissAlertsRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoDismissAlertsRuleName": autorest.Encode("path", autoDismissAlertsRuleName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoDismissAlertsRules/{autoDismissAlertsRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AutoDismissAlertsRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AutoDismissAlertsRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get dismiss rule, with name: {autoDismissAlertsRuleName}, for the given subscription
// Parameters:
// autoDismissAlertsRuleName - the unique name of the auto dismiss alert rule
func (client AutoDismissAlertsRulesClient) Get(ctx context.Context, autoDismissAlertsRuleName string) (result AutoDismissAlertsRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoDismissAlertsRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AutoDismissAlertsRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, autoDismissAlertsRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AutoDismissAlertsRulesClient) GetPreparer(ctx context.Context, autoDismissAlertsRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoDismissAlertsRuleName": autorest.Encode("path", autoDismissAlertsRuleName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoDismissAlertsRules/{autoDismissAlertsRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AutoDismissAlertsRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AutoDismissAlertsRulesClient) GetResponder(resp *http.Response) (result AutoDismissAlertsRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list of all the dismiss rules for the given subscription
// Parameters:
// alertType - type of the alert to get rules for
func (client AutoDismissAlertsRulesClient) List(ctx context.Context, alertType string) (result AutoDismissAlertsRulesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoDismissAlertsRulesClient.List")
		defer func() {
			sc := -1
			if result.adarl.Response.Response != nil {
				sc = result.adarl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AutoDismissAlertsRulesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, alertType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.adarl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "List", resp, "Failure sending request")
		return
	}

	result.adarl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AutoDismissAlertsRulesClient) ListPreparer(ctx context.Context, alertType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(alertType) > 0 {
		queryParameters["AlertType"] = autorest.Encode("query", alertType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoDismissAlertsRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AutoDismissAlertsRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AutoDismissAlertsRulesClient) ListResponder(resp *http.Response) (result AutoDismissAlertsRulesList, err error) {
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
func (client AutoDismissAlertsRulesClient) listNextResults(ctx context.Context, lastResults AutoDismissAlertsRulesList) (result AutoDismissAlertsRulesList, err error) {
	req, err := lastResults.autoDismissAlertsRulesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AutoDismissAlertsRulesClient) ListComplete(ctx context.Context, alertType string) (result AutoDismissAlertsRulesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoDismissAlertsRulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, alertType)
	return
}

// Update update existing rule or create new rule if it doesn't exist
// Parameters:
// autoDismissAlertsRuleName - the unique name of the auto dismiss alert rule
// autoDismissAlertsRule - auto dismiss rule object
func (client AutoDismissAlertsRulesClient) Update(ctx context.Context, autoDismissAlertsRuleName string, autoDismissAlertsRule AutoDismissAlertsRule) (result AutoDismissAlertsRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoDismissAlertsRulesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: autoDismissAlertsRule,
			Constraints: []validation.Constraint{{Target: "autoDismissAlertsRule.AutoDismissAlertsRuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "autoDismissAlertsRule.AutoDismissAlertsRuleProperties.AlertType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "autoDismissAlertsRule.AutoDismissAlertsRuleProperties.Reason", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "autoDismissAlertsRule.AutoDismissAlertsRuleProperties.AutoDismissAlertsScope", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "autoDismissAlertsRule.AutoDismissAlertsRuleProperties.AutoDismissAlertsScope.AllOf", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("security.AutoDismissAlertsRulesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, autoDismissAlertsRuleName, autoDismissAlertsRule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AutoDismissAlertsRulesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AutoDismissAlertsRulesClient) UpdatePreparer(ctx context.Context, autoDismissAlertsRuleName string, autoDismissAlertsRule AutoDismissAlertsRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoDismissAlertsRuleName": autorest.Encode("path", autoDismissAlertsRuleName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoDismissAlertsRules/{autoDismissAlertsRuleName}", pathParameters),
		autorest.WithJSON(autoDismissAlertsRule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AutoDismissAlertsRulesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AutoDismissAlertsRulesClient) UpdateResponder(resp *http.Response) (result AutoDismissAlertsRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
