// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAssignedSubscriptionsRequest wrapper for the ListAssignedSubscriptions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListAssignedSubscriptions.go.html to see an example of how to use ListAssignedSubscriptionsRequest.
type ListAssignedSubscriptionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the subscription to which the tenancy is associated.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListAssignedSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListAssignedSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssignedSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssignedSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssignedSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssignedSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssignedSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListAssignedSubscriptionsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssignedSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListAssignedSubscriptionsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssignedSubscriptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssignedSubscriptionsResponse wrapper for the ListAssignedSubscriptions operation
type ListAssignedSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssignedSubscriptionCollection instances
	AssignedSubscriptionCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAssignedSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssignedSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssignedSubscriptionsSortOrderEnum Enum with underlying type: string
type ListAssignedSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListAssignedSubscriptionsSortOrderEnum
const (
	ListAssignedSubscriptionsSortOrderAsc  ListAssignedSubscriptionsSortOrderEnum = "ASC"
	ListAssignedSubscriptionsSortOrderDesc ListAssignedSubscriptionsSortOrderEnum = "DESC"
)

var mappingListAssignedSubscriptionsSortOrderEnum = map[string]ListAssignedSubscriptionsSortOrderEnum{
	"ASC":  ListAssignedSubscriptionsSortOrderAsc,
	"DESC": ListAssignedSubscriptionsSortOrderDesc,
}

// GetListAssignedSubscriptionsSortOrderEnumValues Enumerates the set of values for ListAssignedSubscriptionsSortOrderEnum
func GetListAssignedSubscriptionsSortOrderEnumValues() []ListAssignedSubscriptionsSortOrderEnum {
	values := make([]ListAssignedSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListAssignedSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssignedSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListAssignedSubscriptionsSortOrderEnum
func GetListAssignedSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListAssignedSubscriptionsSortByEnum Enum with underlying type: string
type ListAssignedSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListAssignedSubscriptionsSortByEnum
const (
	ListAssignedSubscriptionsSortByTimecreated ListAssignedSubscriptionsSortByEnum = "timeCreated"
	ListAssignedSubscriptionsSortByDisplayname ListAssignedSubscriptionsSortByEnum = "displayName"
)

var mappingListAssignedSubscriptionsSortByEnum = map[string]ListAssignedSubscriptionsSortByEnum{
	"timeCreated": ListAssignedSubscriptionsSortByTimecreated,
	"displayName": ListAssignedSubscriptionsSortByDisplayname,
}

// GetListAssignedSubscriptionsSortByEnumValues Enumerates the set of values for ListAssignedSubscriptionsSortByEnum
func GetListAssignedSubscriptionsSortByEnumValues() []ListAssignedSubscriptionsSortByEnum {
	values := make([]ListAssignedSubscriptionsSortByEnum, 0)
	for _, v := range mappingListAssignedSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssignedSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListAssignedSubscriptionsSortByEnum
func GetListAssignedSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}
