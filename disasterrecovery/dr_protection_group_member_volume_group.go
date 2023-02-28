// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberVolumeGroup Properties for a Volume Group member of a DR Protection Group.
type DrProtectionGroupMemberVolumeGroup struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1.phx.exampleocid1`
	MemberId *string `mandatory:"true" json:"memberId"`
}

//GetMemberId returns MemberId
func (m DrProtectionGroupMemberVolumeGroup) GetMemberId() *string {
	return m.MemberId
}

func (m DrProtectionGroupMemberVolumeGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroupMemberVolumeGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrProtectionGroupMemberVolumeGroup) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrProtectionGroupMemberVolumeGroup DrProtectionGroupMemberVolumeGroup
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeDrProtectionGroupMemberVolumeGroup
	}{
		"VOLUME_GROUP",
		(MarshalTypeDrProtectionGroupMemberVolumeGroup)(m),
	}

	return json.Marshal(&s)
}