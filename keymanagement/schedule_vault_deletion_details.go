// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleVaultDeletionDetails Details for scheduling vault deletion.
type ScheduleVaultDeletionDetails struct {

	// An optional property indicating when to delete the vault, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format. The specified
	// time must be between 7 and 30 days from the time when the request is received.
	// If this property is missing, it will be set to 30 days from the time of the request
	// by default.
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`
}

func (m ScheduleVaultDeletionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleVaultDeletionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
