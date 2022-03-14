// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// AssociationSummary The details of the association.
type AssociationSummary struct {

	// The OCID of the association.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name generated by the service for the association. Name format follows the pattern [certificatesResourceEntityType]-[associatedResourceEntityType]-UUID.
	Name *string `mandatory:"true" json:"name"`

	// A property indicating when the association was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the association.
	LifecycleState AssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the certificate-related resource associated with another Oracle Cloud Infrastructure resource.
	CertificatesResourceId *string `mandatory:"true" json:"certificatesResourceId"`

	// The OCID of the associated resource.
	AssociatedResourceId *string `mandatory:"true" json:"associatedResourceId"`

	// The compartment OCID of the association. This is strongly tied to the compartment OCID of the certificate-related resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the association.
	AssociationType AssociationTypeEnum `mandatory:"true" json:"associationType"`
}

func (m AssociationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssociationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAssociationTypeEnum(string(m.AssociationType)); !ok && m.AssociationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationType: %s. Supported values are: %s.", m.AssociationType, strings.Join(GetAssociationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
