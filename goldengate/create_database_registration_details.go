// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateDatabaseRegistrationDetails The information about a new DatabaseRegistration.
type CreateDatabaseRegistrationDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"true" json:"fqdn"`

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"true" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, and so on.
	Password *string `mandatory:"true" json:"password"`

	// Credential store alias.
	AliasName *string `mandatory:"true" json:"aliasName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The private IP address in the customer's VCN of the customer's endpoint, typically a database.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database being referenced.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// Connect descriptor or Easy Connect Naming method that Oracle GoldenGate uses to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The mode of the database connection session to be established by the data client. REDIRECT - for a RAC database, DIRECT - for a non-RAC database. Connection to a RAC database involves a redirection received from the SCAN listeners to the database node to connect to. By default the mode would be DIRECT.
	SessionMode CreateDatabaseRegistrationDetailsSessionModeEnum `mandatory:"false" json:"sessionMode,omitempty"`

	// The wallet contents Oracle GoldenGate uses to make connections to a database.  This attribute is expected to be base64 encoded.
	Wallet *string `mandatory:"false" json:"wallet"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer vault being referenced. If provided, this will reference a vault which the customer will be required to ensure the policies are established to permit the GoldenGate Service to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer "Master" key being referenced. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this key to manage secrets.
	KeyId *string `mandatory:"false" json:"keyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the the GGS Secret will be created. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this Compartment in which to create a Secret.
	SecretCompartmentId *string `mandatory:"false" json:"secretCompartmentId"`
}

func (m CreateDatabaseRegistrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseRegistrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingCreateDatabaseRegistrationDetailsSessionModeEnum[string(m.SessionMode)]; !ok && m.SessionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionMode: %s. Supported values are: %s.", m.SessionMode, strings.Join(GetCreateDatabaseRegistrationDetailsSessionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseRegistrationDetailsSessionModeEnum Enum with underlying type: string
type CreateDatabaseRegistrationDetailsSessionModeEnum string

// Set of constants representing the allowable values for CreateDatabaseRegistrationDetailsSessionModeEnum
const (
	CreateDatabaseRegistrationDetailsSessionModeDirect   CreateDatabaseRegistrationDetailsSessionModeEnum = "DIRECT"
	CreateDatabaseRegistrationDetailsSessionModeRedirect CreateDatabaseRegistrationDetailsSessionModeEnum = "REDIRECT"
)

var mappingCreateDatabaseRegistrationDetailsSessionModeEnum = map[string]CreateDatabaseRegistrationDetailsSessionModeEnum{
	"DIRECT":   CreateDatabaseRegistrationDetailsSessionModeDirect,
	"REDIRECT": CreateDatabaseRegistrationDetailsSessionModeRedirect,
}

// GetCreateDatabaseRegistrationDetailsSessionModeEnumValues Enumerates the set of values for CreateDatabaseRegistrationDetailsSessionModeEnum
func GetCreateDatabaseRegistrationDetailsSessionModeEnumValues() []CreateDatabaseRegistrationDetailsSessionModeEnum {
	values := make([]CreateDatabaseRegistrationDetailsSessionModeEnum, 0)
	for _, v := range mappingCreateDatabaseRegistrationDetailsSessionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseRegistrationDetailsSessionModeEnumStringValues Enumerates the set of values in String for CreateDatabaseRegistrationDetailsSessionModeEnum
func GetCreateDatabaseRegistrationDetailsSessionModeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"REDIRECT",
	}
}
