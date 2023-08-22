// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer Infrastructure deployments
// and scheduled upgrades. For more information see
// Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/Content/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCccUpgradeScheduleDetails Defines a schedule for times when automated Compute Cloud@Customer upgrades are preferred.
// A created upgrade schedule must supply events with a minimum frequency and duration or
// the schedule will be rejected. Upgrades may impact performance of
// Compute Cloud@Customer infrastructures when they are being applied.
type CreateCccUpgradeScheduleDetails struct {

	// Compute Cloud@Customer upgrade schedule display name.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the
	// Compute Cloud@Customer Upgrade Schedule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of preferred times for Compute Cloud@Customer infrastructure to be upgraded.
	Events []CreateCccScheduleEvent `mandatory:"true" json:"events"`

	// An optional description of the Compute Cloud@Customer upgrade schedule.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCccUpgradeScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCccUpgradeScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
