// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// NativeShapeField The native shape field object.
type NativeShapeField struct {

	// The type reference.
	Type *interface{} `mandatory:"true" json:"type"`

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The position of the attribute.
	Position *int `mandatory:"false" json:"position"`

	// The default value.
	DefaultValueString *string `mandatory:"false" json:"defaultValueString"`

	// Specifies whether the field is mandatory.
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`
}

//GetKey returns Key
func (m NativeShapeField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m NativeShapeField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m NativeShapeField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m NativeShapeField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m NativeShapeField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m NativeShapeField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m NativeShapeField) GetDescription() *string {
	return m.Description
}

func (m NativeShapeField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NativeShapeField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NativeShapeField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNativeShapeField NativeShapeField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeNativeShapeField
	}{
		"NATIVE_SHAPE_FIELD",
		(MarshalTypeNativeShapeField)(m),
	}

	return json.Marshal(&s)
}