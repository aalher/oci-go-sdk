// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary Specifies the instance group blue-green deployment load balancer traffic shift stage.
type ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the upstream compute instance group blue-green deployment stage in this pipeline.
	ComputeInstanceGroupBlueGreenDeploymentDeployStageId *string `mandatory:"true" json:"computeInstanceGroupBlueGreenDeploymentDeployStageId"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the deployment stage.
	LifecycleState DeployStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetProjectId returns ProjectId
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetProjectId() *string {
	return m.ProjectId
}

// GetDeployPipelineId returns DeployPipelineId
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetCompartmentId returns CompartmentId
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeployStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary
	}{
		"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT",
		(MarshalTypeComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary)(m),
	}

	return json.Marshal(&s)
}
