// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"context"
	"fmt"
	"net/http"
)

type ComputeClient struct {
	common.BaseClient
}

// Create a new default Compute client for a given region and default configuration provider
func NewComputeClientForRegion(region common.Region) (client ComputeClient) {
	client = ComputeClient{BaseClient: common.NewClientForRegion(region)}

	client.Host = fmt.Sprintf(common.DefaultHostUrlTemplate, "iaas", string(region))
	client.BasePath = "20160918"
	return
}

// Create a new default Compute client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComputeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComputeClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = ComputeClient{BaseClient: baseClient}
	region, err := configProvider.Region()
	if err != nil {
		return
	}

	client.Host = fmt.Sprintf(common.DefaultHostUrlTemplate, "iaas", string(region))
	client.BasePath = "20160918"
	return
}

// Adds a shape to the compatible shapes list for the image.
func (client ComputeClient) AddImageShapeCompatibilityEntry(ctx context.Context, request AddImageShapeCompatibilityEntryRequest) (response AddImageShapeCompatibilityEntryResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/images/{imageId}/shapes/{shapeName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a secondary VNIC and attaches it to the specified instance.
// For more information about secondary VNICs, see
// [Virtual Network Interface Cards (VNICs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVNICs.htm).
func (client ComputeClient) AttachVnic(ctx context.Context, request AttachVnicRequest) (response AttachVnicResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/vnicAttachments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Attaches the specified storage volume to the specified instance.
func (client ComputeClient) AttachVolume(ctx context.Context, request AttachVolumeRequest) (response AttachVolumeResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/volumeAttachments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &volumeattachment{})
	return
}

// Captures the most recent serial console data (up to a megabyte) for the
// specified instance.
// The `CaptureConsoleHistory` operation works with the other console history operations
// as described below.
// 1. Use `CaptureConsoleHistory` to request the capture of up to a megabyte of the
// most recent console history. This call returns a `ConsoleHistory`
// object. The object will have a state of REQUESTED.
// 2. Wait for the capture operation to succeed by polling `GetConsoleHistory` with
// the identifier of the console history metadata. The state of the
// `ConsoleHistory` object will go from REQUESTED to GETTING-HISTORY and
// then SUCCEEDED (or FAILED).
// 3. Use `GetConsoleHistoryContent` to get the actual console history data (not the
// metadata).
// 4. Optionally, use `DeleteConsoleHistory` to delete the console history metadata
// and the console history data.
func (client ComputeClient) CaptureConsoleHistory(ctx context.Context, request CaptureConsoleHistoryRequest) (response CaptureConsoleHistoryResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/instanceConsoleHistories/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a boot disk image for the specified instance or imports an exported image from the Oracle Cloud Infrastructure Object Storage service.
// When creating a new image, you must provide the OCID of the instance you want to use as the basis for the image, and
// the OCID of the compartment containing that instance. For more information about images,
// see [Managing Custom Images]({{DOC_SERVER_URL}}/Content/Compute/Tasks/managingcustomimages.htm).
// When importing an exported image from Object Storage, you specify the source information
// in ImageSourceDetails.
// When importing an image based on the namespace, bucket name, and object name,
// use ImageSourceViaObjectStorageTupleDetails.
// When importing an image based on the Object Storage URL, use
// ImageSourceViaObjectStorageUriDetails.
// See [Object Storage URLs]({{DOC_SERVER_URL}}/Content/Compute/Tasks/imageimportexport.htm#URLs) and [pre-authenticated requests]({{DOC_SERVER_URL}}/Content/Object/Tasks/managingaccess.htm#pre-auth)
// for constructing URLs for image import/export.
// For more information about importing exported images, see
// [Image Import/Export]({{DOC_SERVER_URL}}/Content/Compute/Tasks/imageimportexport.htm).
// You may optionally specify a *display name* for the image, which is simply a friendly name or description.
// It does not have to be unique, and you can change it. See UpdateImage.
// Avoid entering confidential information.
func (client ComputeClient) CreateImage(ctx context.Context, request CreateImageRequest) (response CreateImageResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/images/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new serial console connection to the specified instance.
// Once the serial console connection has been created and is available,
// you connect to the serial console using an SSH client.
// The default number of enabled serial console connections per tenancy is 10.
// For more information about serial console access, see [Accessing the Serial Console]({{DOC_SERVER_URL}}/Content/Compute/References/serialconsole.htm).
func (client ComputeClient) CreateInstanceConsoleConnection(ctx context.Context, request CreateInstanceConsoleConnectionRequest) (response CreateInstanceConsoleConnectionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/instanceConsoleConnections/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Deletes the specified console history metadata and the console history data.
func (client ComputeClient) DeleteConsoleHistory(ctx context.Context, request DeleteConsoleHistoryRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/instanceConsoleHistories/{instanceConsoleHistoryId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes an image.
func (client ComputeClient) DeleteImage(ctx context.Context, request DeleteImageRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/images/{imageId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified serial console connection.
func (client ComputeClient) DeleteInstanceConsoleConnection(ctx context.Context, request DeleteInstanceConsoleConnectionRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/instanceConsoleConnections/{instanceConsoleConnectionId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Detaches and deletes the specified secondary VNIC.
// This operation cannot be used on the instance's primary VNIC.
// When you terminate an instance, all attached VNICs (primary
// and secondary) are automatically detached and deleted.
// **Important:** If the VNIC has a
// PrivateIp that is the
// [target of a route rule]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingroutetables.htm#privateip),
// deleting the VNIC causes that route rule to blackhole and the traffic
// will be dropped.
func (client ComputeClient) DetachVnic(ctx context.Context, request DetachVnicRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/vnicAttachments/{vnicAttachmentId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Detaches a storage volume from an instance. You must specify the OCID of the volume attachment.
// This is an asynchronous operation. The attachment's `lifecycleState` will change to DETACHING temporarily
// until the attachment is completely removed.
func (client ComputeClient) DetachVolume(ctx context.Context, request DetachVolumeRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/volumeAttachments/{volumeAttachmentId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Exports the specified image to the Oracle Cloud Infrastructure Object Storage service. You can use the Object Storage URL,
// or the namespace, bucket name, and object name when specifying the location to export to.
// For more information about exporting images, see [Image Import/Export]({{DOC_SERVER_URL}}/Content/Compute/Tasks/imageimportexport.htm).
// To perform an image export, you need write access to the Object Storage bucket for the image,
// see [Let Users Write Objects to Object Storage Buckets]({{DOC_SERVER_URL}}/Content/Identity/Concepts/commonpolicies.htm#Let4).
// See [Object Storage URLs]({{DOC_SERVER_URL}}/Content/Compute/Tasks/imageimportexport.htm#URLs) and [pre-authenticated requests]({{DOC_SERVER_URL}}/Content/Object/Tasks/managingaccess.htm#pre-auth)
// for constructing URLs for image import/export.
func (client ComputeClient) ExportImage(ctx context.Context, request ExportImageRequest) (response ExportImageResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/images/{imageId}/actions/export", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Shows the metadata for the specified console history.
// See CaptureConsoleHistory
// for details about using the console history operations.
func (client ComputeClient) GetConsoleHistory(ctx context.Context, request GetConsoleHistoryRequest) (response GetConsoleHistoryResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instanceConsoleHistories/{instanceConsoleHistoryId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the actual console history data (not the metadata).
// See CaptureConsoleHistory
// for details about using the console history operations.
func (client ComputeClient) GetConsoleHistoryContent(ctx context.Context, request GetConsoleHistoryContentRequest) (response GetConsoleHistoryContentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instanceConsoleHistories/{instanceConsoleHistoryId}/data", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified image.
func (client ComputeClient) GetImage(ctx context.Context, request GetImageRequest) (response GetImageResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/images/{imageId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about the specified instance.
func (client ComputeClient) GetInstance(ctx context.Context, request GetInstanceRequest) (response GetInstanceResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instances/{instanceId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified serial console connection's information.
func (client ComputeClient) GetInstanceConsoleConnection(ctx context.Context, request GetInstanceConsoleConnectionRequest) (response GetInstanceConsoleConnectionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instanceConsoleConnections/{instanceConsoleConnectionId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the information for the specified VNIC attachment.
func (client ComputeClient) GetVnicAttachment(ctx context.Context, request GetVnicAttachmentRequest) (response GetVnicAttachmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/vnicAttachments/{vnicAttachmentId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets information about the specified volume attachment.
func (client ComputeClient) GetVolumeAttachment(ctx context.Context, request GetVolumeAttachmentRequest) (response GetVolumeAttachmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/volumeAttachments/{volumeAttachmentId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &volumeattachment{})
	return
}

// Gets the generated credentials for the instance. Only works for Windows instances. The returned credentials
// are only valid for the initial login.
func (client ComputeClient) GetWindowsInstanceInitialCredentials(ctx context.Context, request GetWindowsInstanceInitialCredentialsRequest) (response GetWindowsInstanceInitialCredentialsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instances/{instanceId}/initialCredentials", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Performs one of the power actions (start, stop, softreset, or reset)
// on the specified instance.
// **start** - power on
// **stop** - power off
// **softreset** - ACPI shutdown and power on
// **reset** - power off and power on
// Note that the **stop** state has no effect on the resources you consume.
// Billing continues for instances that you stop, and related resources continue
// to apply against any relevant quotas. You must terminate an instance
// (TerminateInstance)
// to remove its resources from billing and quotas.
func (client ComputeClient) InstanceAction(ctx context.Context, request InstanceActionRequest) (response InstanceActionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/instances/{instanceId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new instance in the specified compartment and the specified Availability Domain.
// For general information about instances, see
// [Overview of the Compute Service]({{DOC_SERVER_URL}}/Content/Compute/Concepts/computeoverview.htm).
// For information about access control and compartments, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about Availability Domains, see
// [Regions and Availability Domains]({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
// To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
// in the Identity and Access Management Service API.
// All Oracle Cloud Infrastructure resources, including instances, get an Oracle-assigned,
// unique ID called an Oracle Cloud Identifier (OCID).
// When you create a resource, you can find its OCID in the response. You can
// also retrieve a resource's OCID by using a List API operation
// on that resource type, or by viewing the resource in the Console.
// When you launch an instance, it is automatically attached to a virtual
// network interface card (VNIC), called the *primary VNIC*. The VNIC
// has a private IP address from the subnet's CIDR. You can either assign a
// private IP address of your choice or let Oracle automatically assign one.
// You can choose whether the instance has a public IP address. To retrieve the
// addresses, use the ListVnicAttachments
// operation to get the VNIC ID for the instance, and then call
// GetVnic with the VNIC ID.
// You can later add secondary VNICs to an instance. For more information, see
// [Virtual Network Interface Cards (VNICs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVNICs.htm).
func (client ComputeClient) LaunchInstance(ctx context.Context, request LaunchInstanceRequest) (response LaunchInstanceResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/instances/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the console history metadata for the specified compartment or instance.
func (client ComputeClient) ListConsoleHistories(ctx context.Context, request ListConsoleHistoriesRequest) (response ListConsoleHistoriesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instanceConsoleHistories/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the available images in the specified compartment. For more
// information about images, see
// [Managing Custom Images]({{DOC_SERVER_URL}}/Content/Compute/Tasks/managingcustomimages.htm).
func (client ComputeClient) ListImages(ctx context.Context, request ListImagesRequest) (response ListImagesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/images/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the serial console connections for the specified compartment or instance.
// For more information about serial console access, see [Accessing the Serial Console]({{DOC_SERVER_URL}}/Content/Compute/References/serialconsole.htm).
func (client ComputeClient) ListInstanceConsoleConnections(ctx context.Context, request ListInstanceConsoleConnectionsRequest) (response ListInstanceConsoleConnectionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instanceConsoleConnections/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the instances in the specified compartment and the specified Availability Domain.
// You can filter the results by specifying an instance name (the list will include all the identically-named
// instances in the compartment).
func (client ComputeClient) ListInstances(ctx context.Context, request ListInstancesRequest) (response ListInstancesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/instances/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the shapes that can be used to launch an instance within the specified compartment. You can
// filter the list by compatibility with a specific image.
func (client ComputeClient) ListShapes(ctx context.Context, request ListShapesRequest) (response ListShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/shapes", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the VNIC attachments in the specified compartment. A VNIC attachment
// resides in the same compartment as the attached instance. The list can be
// filtered by instance, VNIC, or Availability Domain.
func (client ComputeClient) ListVnicAttachments(ctx context.Context, request ListVnicAttachmentsRequest) (response ListVnicAttachmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/vnicAttachments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the volume attachments in the specified compartment. You can filter the
// list by specifying an instance OCID, volume OCID, or both.
// Currently, the only supported volume attachment type is IScsiVolumeAttachment.
func (client ComputeClient) ListVolumeAttachments(ctx context.Context, request ListVolumeAttachmentsRequest) (response ListVolumeAttachmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/volumeAttachments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Removes a shape from the compatible shapes list for the image.
func (client ComputeClient) RemoveImageShapeCompatibilityEntry(ctx context.Context, request RemoveImageShapeCompatibilityEntryRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/images/{imageId}/shapes/{shapeName}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Terminates the specified instance. Any attached VNICs and volumes are automatically detached
// when the instance terminates.
// This is an asynchronous operation. The instance's `lifecycleState` will change to TERMINATING temporarily
// until the instance is completely removed.
func (client ComputeClient) TerminateInstance(ctx context.Context, request TerminateInstanceRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/instances/{instanceId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Updates the specified console history metadata.
func (client ComputeClient) UpdateConsoleHistory(ctx context.Context, request UpdateConsoleHistoryRequest) (response UpdateConsoleHistoryResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/instanceConsoleHistories/{instanceConsoleHistoryId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the display name of the image. Avoid entering confidential information.
func (client ComputeClient) UpdateImage(ctx context.Context, request UpdateImageRequest) (response UpdateImageResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/images/{imageId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the display name of the specified instance. Avoid entering confidential information.
// The OCID of the instance remains the same.
func (client ComputeClient) UpdateInstance(ctx context.Context, request UpdateInstanceRequest) (response UpdateInstanceResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/instances/{instanceId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}