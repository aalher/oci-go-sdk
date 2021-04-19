// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v40/common"
	"net/http"
)

// ChangeTsigKeyCompartmentRequest wrapper for the ChangeTsigKeyCompartment operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeTsigKeyCompartment.go.html to see an example of how to use ChangeTsigKeyCompartmentRequest.
type ChangeTsigKeyCompartmentRequest struct {

	// The OCID of the target TSIG key.
	TsigKeyId *string `mandatory:"true" contributesTo:"path" name:"tsigKeyId"`

	// Details for moving a TSIG key into a different compartment.
	ChangeTsigKeyCompartmentDetails `contributesTo:"body"`

	// The `If-Match` header field makes the request method conditional on the
	// existence of at least one current representation of the target resource,
	// when the field-value is `*`, or having a current representation of the
	// target resource that has an entity-tag matching a member of the list of
	// entity-tags provided in the field-value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated before
	// then due to conflicting operations (for example, if a resource has been
	// deleted and purged from the system, then a retry of the original creation
	// request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ChangeTsigKeyCompartmentScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeTsigKeyCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeTsigKeyCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ChangeTsigKeyCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeTsigKeyCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ChangeTsigKeyCompartmentResponse wrapper for the ChangeTsigKeyCompartment operation
type ChangeTsigKeyCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangeTsigKeyCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeTsigKeyCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ChangeTsigKeyCompartmentScopeEnum Enum with underlying type: string
type ChangeTsigKeyCompartmentScopeEnum string

// Set of constants representing the allowable values for ChangeTsigKeyCompartmentScopeEnum
const (
	ChangeTsigKeyCompartmentScopeGlobal  ChangeTsigKeyCompartmentScopeEnum = "GLOBAL"
	ChangeTsigKeyCompartmentScopePrivate ChangeTsigKeyCompartmentScopeEnum = "PRIVATE"
)

var mappingChangeTsigKeyCompartmentScope = map[string]ChangeTsigKeyCompartmentScopeEnum{
	"GLOBAL":  ChangeTsigKeyCompartmentScopeGlobal,
	"PRIVATE": ChangeTsigKeyCompartmentScopePrivate,
}

// GetChangeTsigKeyCompartmentScopeEnumValues Enumerates the set of values for ChangeTsigKeyCompartmentScopeEnum
func GetChangeTsigKeyCompartmentScopeEnumValues() []ChangeTsigKeyCompartmentScopeEnum {
	values := make([]ChangeTsigKeyCompartmentScopeEnum, 0)
	for _, v := range mappingChangeTsigKeyCompartmentScope {
		values = append(values, v)
	}
	return values
}
