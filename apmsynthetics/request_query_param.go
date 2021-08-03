// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// RequestQueryParam Information about request query parameters.
type RequestQueryParam struct {

	// Name of request query parameter.
	ParamName *string `mandatory:"true" json:"paramName"`

	// Value of request query parameter.
	ParamValue *string `mandatory:"false" json:"paramValue"`
}

func (m RequestQueryParam) String() string {
	return common.PointerString(m)
}