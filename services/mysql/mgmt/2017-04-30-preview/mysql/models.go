package mysql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// CreateModeType enumerates the values for create mode.
type CreateModeType string

const (
	// CreateModeCreateModeDefault ...
	CreateModeCreateModeDefault CreateModeType = "Default"
	// CreateModeCreateModePointInTimeRestore ...
	CreateModeCreateModePointInTimeRestore CreateModeType = "PointInTimeRestore"
)

// OperationOriginType enumerates the values for operation origin.
type OperationOriginType string

const (
	// OperationOriginNone represents an empty OperationOriginType.
	OperationOriginNone OperationOriginType = ""
	// OperationOriginNotSpecified ...
	OperationOriginNotSpecified OperationOriginType = "NotSpecified"
	// OperationOriginSystem ...
	OperationOriginSystem OperationOriginType = "system"
	// OperationOriginUser ...
	OperationOriginUser OperationOriginType = "user"
)

// ServerStateType enumerates the values for server state.
type ServerStateType string

const (
	// ServerStateDisabled ...
	ServerStateDisabled ServerStateType = "Disabled"
	// ServerStateDropping ...
	ServerStateDropping ServerStateType = "Dropping"
	// ServerStateNone represents an empty ServerStateType.
	ServerStateNone ServerStateType = ""
	// ServerStateReady ...
	ServerStateReady ServerStateType = "Ready"
)

// ServerVersionType enumerates the values for server version.
type ServerVersionType string

const (
	// ServerVersionFiveFullStopSeven ...
	ServerVersionFiveFullStopSeven ServerVersionType = "5.7"
	// ServerVersionFiveFullStopSix ...
	ServerVersionFiveFullStopSix ServerVersionType = "5.6"
	// ServerVersionNone represents an empty ServerVersionType.
	ServerVersionNone ServerVersionType = ""
)

// SkuTierType enumerates the values for sku tier.
type SkuTierType string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTierType = "Basic"
	// SkuTierNone represents an empty SkuTierType.
	SkuTierNone SkuTierType = ""
	// SkuTierStandard ...
	SkuTierStandard SkuTierType = "Standard"
)

// SslEnforcementEnumType enumerates the values for ssl enforcement enum.
type SslEnforcementEnumType string

const (
	// SslEnforcementEnumDisabled ...
	SslEnforcementEnumDisabled SslEnforcementEnumType = "Disabled"
	// SslEnforcementEnumEnabled ...
	SslEnforcementEnumEnabled SslEnforcementEnumType = "Enabled"
	// SslEnforcementEnumNone represents an empty SslEnforcementEnumType.
	SslEnforcementEnumNone SslEnforcementEnumType = ""
)

// Configuration - Represents a Configuration.
type Configuration struct {
	rawResponse *http.Response
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Properties - The properties of a configuration.
	*ConfigurationProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (c Configuration) Response() *http.Response {
	return c.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (c Configuration) StatusCode() int {
	return c.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (c Configuration) Status() string {
	return c.rawResponse.Status
}

// ConfigurationListResult - A list of server configurations.
type ConfigurationListResult struct {
	rawResponse *http.Response
	// Value - The list of server configurations.
	Value []Configuration `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (clr ConfigurationListResult) Response() *http.Response {
	return clr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (clr ConfigurationListResult) StatusCode() int {
	return clr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (clr ConfigurationListResult) Status() string {
	return clr.rawResponse.Status
}

// ConfigurationProperties - The properties of a configuration.
type ConfigurationProperties struct {
	// Value - Value of the configuration.
	Value *string `json:"value,omitempty"`
	// Description - Description of the configuration.
	Description *string `json:"description,omitempty"`
	// DefaultValue - Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// DataType - Data type of the configuration.
	DataType *string `json:"dataType,omitempty"`
	// AllowedValues - Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`
	// Source - Source of the configuration.
	Source *string `json:"source,omitempty"`
}

// Database - Represents a Database.
type Database struct {
	rawResponse *http.Response
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Properties - The properties of a database.
	*DatabaseProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (d Database) Response() *http.Response {
	return d.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (d Database) StatusCode() int {
	return d.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (d Database) Status() string {
	return d.rawResponse.Status
}

// DatabaseListResult - A List of databases.
type DatabaseListResult struct {
	rawResponse *http.Response
	// Value - The list of databases housed in a server
	Value []Database `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (dlr DatabaseListResult) Response() *http.Response {
	return dlr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (dlr DatabaseListResult) StatusCode() int {
	return dlr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (dlr DatabaseListResult) Status() string {
	return dlr.rawResponse.Status
}

// DatabaseProperties - The properties of a database.
type DatabaseProperties struct {
	// Charset - The charset of the database.
	Charset *string `json:"charset,omitempty"`
	// Collation - The collation of the database.
	Collation *string `json:"collation,omitempty"`
}

// FirewallRule - Represents a server firewall rule.
type FirewallRule struct {
	rawResponse *http.Response
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Properties - The properties of a firewall rule.
	*FirewallRuleProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (fr FirewallRule) Response() *http.Response {
	return fr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (fr FirewallRule) StatusCode() int {
	return fr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (fr FirewallRule) Status() string {
	return fr.rawResponse.Status
}

// FirewallRuleListResult - A list of firewall rules.
type FirewallRuleListResult struct {
	rawResponse *http.Response
	// Value - The list of firewall rules in a server.
	Value []FirewallRule `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (frlr FirewallRuleListResult) Response() *http.Response {
	return frlr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (frlr FirewallRuleListResult) StatusCode() int {
	return frlr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (frlr FirewallRuleListResult) Status() string {
	return frlr.rawResponse.Status
}

// FirewallRuleProperties - The properties of a server firewall rule.
type FirewallRuleProperties struct {
	// StartIPAddress - The start IP address of the server firewall rule. Must be IPv4 format.
	StartIPAddress string `json:"startIpAddress,omitempty"`
	// EndIPAddress - The end IP address of the server firewall rule. Must be IPv4 format.
	EndIPAddress string `json:"endIpAddress,omitempty"`
}

// LogFile - Represents a log file.
type LogFile struct {
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Properties - The properties of the log file.
	*LogFileProperties `json:"properties,omitempty"`
}

// LogFileListResult - A list of log files.
type LogFileListResult struct {
	rawResponse *http.Response
	// Value - The list of log files.
	Value []LogFile `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (lflr LogFileListResult) Response() *http.Response {
	return lflr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (lflr LogFileListResult) StatusCode() int {
	return lflr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (lflr LogFileListResult) Status() string {
	return lflr.rawResponse.Status
}

// LogFileProperties - The properties of a log file.
type LogFileProperties struct {
	// Name - Log file name.
	Name *string `json:"name,omitempty"`
	// SizeInKB - Size of the log file.
	SizeInKB *int64 `json:"sizeInKB,omitempty"`
	// CreatedTime - Creation timestamp of the log file.
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	// LastModifiedTime - Last modified timestamp of the log file.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	// Type - Type of the log file.
	Type *string `json:"type,omitempty"`
	// URL - The url to download the log file from.
	URL *string `json:"url,omitempty"`
}

// NameAvailability - Represents a resource name availability.
type NameAvailability struct {
	rawResponse *http.Response
	// Message - Error Message.
	Message *string `json:"message,omitempty"`
	// NameAvailable - Indicates whether the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - Reason for name being unavailable.
	Reason *string `json:"reason,omitempty"`
}

// Response returns the raw HTTP response object.
func (na NameAvailability) Response() *http.Response {
	return na.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (na NameAvailability) StatusCode() int {
	return na.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (na NameAvailability) Status() string {
	return na.rawResponse.Status
}

// NameAvailabilityRequest - Request from client to check resource name availability.
type NameAvailabilityRequest struct {
	// Name - Resource name to verify.
	Name string `json:"name,omitempty"`
	// Type - Resource type used for verification.
	Type *string `json:"type,omitempty"`
}

// Operation - REST API operation definition.
type Operation struct {
	// Name - The name of the operation being performed on this particular object.
	Name *string `json:"name,omitempty"`
	// Display - The localized display information for this particular operation or action.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - The intended executor of the operation. Possible values include: 'NotSpecified', 'User', 'System', 'None'
	Origin OperationOriginType `json:"origin,omitempty"`
	// Properties - Additional descriptions for the operation.
	Properties map[string]map[string]interface{} `json:"properties,omitempty"`
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Operation resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Localized friendly name for the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Operation description.
	Description *string `json:"description,omitempty"`
}

// OperationListResult - A list of resource provider operations.
type OperationListResult struct {
	rawResponse *http.Response
	// Value - The list of resource provider operations.
	Value []Operation `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (olr OperationListResult) Response() *http.Response {
	return olr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (olr OperationListResult) StatusCode() int {
	return olr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (olr OperationListResult) Status() string {
	return olr.rawResponse.Status
}

// PerformanceTierListResult - A list of performance tiers.
type PerformanceTierListResult struct {
	rawResponse *http.Response
	// Value - The list of performance tiers
	Value []PerformanceTierProperties `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (ptlr PerformanceTierListResult) Response() *http.Response {
	return ptlr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (ptlr PerformanceTierListResult) StatusCode() int {
	return ptlr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (ptlr PerformanceTierListResult) Status() string {
	return ptlr.rawResponse.Status
}

// PerformanceTierProperties - Performance tier properties
type PerformanceTierProperties struct {
	// ID - ID of the performance tier.
	ID *string `json:"id,omitempty"`
	// BackupRetentionDays - Backup retention in days for the performance tier edition
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`
	// ServiceLevelObjectives - Service level objectives associated with the performance tier
	ServiceLevelObjectives []PerformanceTierServiceLevelObjectives `json:"serviceLevelObjectives,omitempty"`
}

// PerformanceTierServiceLevelObjectives - Service level objectives for performance tier.
type PerformanceTierServiceLevelObjectives struct {
	// ID - ID for the service level objective.
	ID *string `json:"id,omitempty"`
	// Edition - Edition of the performance tier.
	Edition *string `json:"edition,omitempty"`
	// Dtu - Database throughput unit associated with the service level objective
	Dtu *int32 `json:"dtu,omitempty"`
	// StorageMB - Maximum storage in MB associated with the service level objective
	StorageMB *int32 `json:"storageMB,omitempty"`
}

// ProxyResource - Resource properties.
type ProxyResource struct {
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// Server - Represents a server.
type Server struct {
	rawResponse *http.Response
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - The location the resource resides in.
	Location string `json:"location,omitempty"`
	// Tags - Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
	// Sku - The SKU (pricing tier) of the server.
	Sku *Sku `json:"sku,omitempty"`
	// Properties - Properties of the server.
	*ServerProperties `json:"properties,omitempty"`
}

// Response returns the raw HTTP response object.
func (s Server) Response() *http.Response {
	return s.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (s Server) StatusCode() int {
	return s.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (s Server) Status() string {
	return s.rawResponse.Status
}

// ServerForCreate - Represents a server to be created.
type ServerForCreate struct {
	// Sku - The SKU (pricing tier) of the server.
	Sku *Sku `json:"sku,omitempty"`
	// Properties - Properties of the server.
	Properties ServerPropertiesForCreate `json:"properties,omitempty"`
	// Location - The location the resource resides in.
	Location string `json:"location,omitempty"`
	// Tags - Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ServerForCreate struct.
func (sfc *ServerForCreate) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["sku"]
	if v != nil {
		var sku Sku
		err = json.Unmarshal(*m["sku"], &sku)
		if err != nil {
			return err
		}
		sfc.Sku = &sku
	}

	v = m["properties"]
	if v != nil {
		properties, err := unmarshalServerPropertiesForCreate(*m["properties"])
		if err != nil {
			return err
		}
		sfc.Properties = properties
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		sfc.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		sfc.Tags = &tags
	}

	return nil
}

// ServerListResult - A list of servers.
type ServerListResult struct {
	rawResponse *http.Response
	// Value - The list of servers
	Value []Server `json:"value,omitempty"`
}

// Response returns the raw HTTP response object.
func (slr ServerListResult) Response() *http.Response {
	return slr.rawResponse
}

// StatusCode returns the HTTP status code of the response, e.g. 200.
func (slr ServerListResult) StatusCode() int {
	return slr.rawResponse.StatusCode
}

// Status returns the HTTP status message of the response, e.g. "200 OK".
func (slr ServerListResult) Status() string {
	return slr.rawResponse.Status
}

// ServerProperties - The properties of a server.
type ServerProperties struct {
	// AdministratorLogin - The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`
	// StorageMB - The maximum storage allowed for a server.
	StorageMB *int64 `json:"storageMB,omitempty"`
	// Version - Server version. Possible values include: 'FiveFullStopSix', 'FiveFullStopSeven', 'None'
	Version ServerVersionType `json:"version,omitempty"`
	// SslEnforcement - Enable ssl enforcement or not when connect to server. Possible values include: 'Enabled', 'Disabled', 'None'
	SslEnforcement SslEnforcementEnumType `json:"sslEnforcement,omitempty"`
	// UserVisibleState - A state of a server that is visible to user. Possible values include: 'Ready', 'Dropping', 'Disabled', 'None'
	UserVisibleState ServerStateType `json:"userVisibleState,omitempty"`
	// FullyQualifiedDomainName - The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`
}

type ServerPropertiesForCreate interface {
	AsServerPropertiesForDefaultCreate() (*ServerPropertiesForDefaultCreate, bool)
	AsServerPropertiesForRestore() (*ServerPropertiesForRestore, bool)
}

func unmarshalServerPropertiesForCreate(body []byte) (ServerPropertiesForCreate, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["createMode"] {
	case string(CreateModeDefault):
		var spfdc ServerPropertiesForDefaultCreate
		err := json.Unmarshal(body, &spfdc)
		return spfdc, err
	case string(CreateModePointInTimeRestore):
		var spfr ServerPropertiesForRestore
		err := json.Unmarshal(body, &spfr)
		return spfr, err
	default:
		return nil, errors.New("Unsupported type")
	}
}
func unmarshalServerPropertiesForCreateArray(body []byte) ([]ServerPropertiesForCreate, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	spfcArray := make([]ServerPropertiesForCreate, len(rawMessages))

	for index, rawMessage := range rawMessages {
		spfc, err := unmarshalServerPropertiesForCreate(*rawMessage)
		if err != nil {
			return nil, err
		}
		spfcArray[index] = spfc
	}
	return spfcArray, nil
}

// ServerPropertiesForDefaultCreate - The properties used to create a new server.
type ServerPropertiesForDefaultCreate struct {
	// StorageMB - The maximum storage allowed for a server.
	StorageMB *int64 `json:"storageMB,omitempty"`
	// Version - Server version. Possible values include: 'FiveFullStopSix', 'FiveFullStopSeven', 'None'
	Version ServerVersionType `json:"version,omitempty"`
	// SslEnforcement - Enable ssl enforcement or not when connect to server. Possible values include: 'Enabled', 'Disabled', 'None'
	SslEnforcement SslEnforcementEnumType `json:"sslEnforcement,omitempty"`
	// CreateMode - Possible values include: 'CreateModeDefault', 'CreateModePointInTimeRestore'
	CreateMode CreateModeType `json:"createMode,omitempty"`
	// AdministratorLogin - The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin string `json:"administratorLogin,omitempty"`
	// AdministratorLoginPassword - The password of the administrator login.
	AdministratorLoginPassword string `json:"administratorLoginPassword,omitempty"`
}

// MarshalJSON is the custom marshaler for ServerPropertiesForDefaultCreate.
func (spfdc ServerPropertiesForDefaultCreate) MarshalJSON() ([]byte, error) {
	spfdc.CreateMode = CreateModeDefault
	type Alias ServerPropertiesForDefaultCreate
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(spfdc),
	})
}

// AsServerPropertiesForDefaultCreate is the ServerPropertiesForCreate implementation for ServerPropertiesForDefaultCreate.
func (spfdc ServerPropertiesForDefaultCreate) AsServerPropertiesForDefaultCreate() (*ServerPropertiesForDefaultCreate, bool) {
	return &spfdc, true
}

// AsServerPropertiesForRestore is the ServerPropertiesForCreate implementation for ServerPropertiesForDefaultCreate.
func (spfdc ServerPropertiesForDefaultCreate) AsServerPropertiesForRestore() (*ServerPropertiesForRestore, bool) {
	return nil, false
}

// ServerPropertiesForRestore - The properties to a new server by restoring from a backup.
type ServerPropertiesForRestore struct {
	// StorageMB - The maximum storage allowed for a server.
	StorageMB *int64 `json:"storageMB,omitempty"`
	// Version - Server version. Possible values include: 'FiveFullStopSix', 'FiveFullStopSeven', 'None'
	Version ServerVersionType `json:"version,omitempty"`
	// SslEnforcement - Enable ssl enforcement or not when connect to server. Possible values include: 'Enabled', 'Disabled', 'None'
	SslEnforcement SslEnforcementEnumType `json:"sslEnforcement,omitempty"`
	// CreateMode - Possible values include: 'CreateModeDefault', 'CreateModePointInTimeRestore'
	CreateMode CreateModeType `json:"createMode,omitempty"`
	// SourceServerID - The source server id to restore from.
	SourceServerID string `json:"sourceServerId,omitempty"`
	// RestorePointInTime - Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime time.Time `json:"restorePointInTime,omitempty"`
}

// MarshalJSON is the custom marshaler for ServerPropertiesForRestore.
func (spfr ServerPropertiesForRestore) MarshalJSON() ([]byte, error) {
	spfr.CreateMode = CreateModePointInTimeRestore
	type Alias ServerPropertiesForRestore
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(spfr),
	})
}

// AsServerPropertiesForDefaultCreate is the ServerPropertiesForCreate implementation for ServerPropertiesForRestore.
func (spfr ServerPropertiesForRestore) AsServerPropertiesForDefaultCreate() (*ServerPropertiesForDefaultCreate, bool) {
	return nil, false
}

// AsServerPropertiesForRestore is the ServerPropertiesForCreate implementation for ServerPropertiesForRestore.
func (spfr ServerPropertiesForRestore) AsServerPropertiesForRestore() (*ServerPropertiesForRestore, bool) {
	return &spfr, true
}

// ServerUpdateParameters - Parameters allowd to update for a server.
type ServerUpdateParameters struct {
	// Sku - The SKU (pricing tier) of the server.
	Sku *Sku `json:"sku,omitempty"`
	// Properties - The properties that can be updated for a server.
	*ServerUpdateParametersProperties `json:"properties,omitempty"`
	// Tags - Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

// ServerUpdateParametersProperties - The properties that can be updated for a server.
type ServerUpdateParametersProperties struct {
	// StorageMB - The max storage allowed for a server.
	StorageMB *int64 `json:"storageMB,omitempty"`
	// AdministratorLoginPassword - The password of the administrator login.
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`
	// Version - The version of a server. Possible values include: 'FiveFullStopSix', 'FiveFullStopSeven', 'None'
	Version ServerVersionType `json:"version,omitempty"`
	// SslEnforcement - Enable ssl enforcement or not when connect to server. Possible values include: 'Enabled', 'Disabled', 'None'
	SslEnforcement SslEnforcementEnumType `json:"sslEnforcement,omitempty"`
}

// Sku - Billing information related properties of a server.
type Sku struct {
	// Name - The name of the sku, typically, a letter + Number code, e.g. P3.
	Name *string `json:"name,omitempty"`
	// Tier - The tier of the particular SKU, e.g. Basic. Possible values include: 'Basic', 'Standard', 'None'
	Tier SkuTierType `json:"tier,omitempty"`
	// Capacity - The scale up/out capacity, representing server's compute units.
	Capacity *int32 `json:"capacity,omitempty"`
	// Size - The size code, to be interpreted by resource as appropriate.
	Size *string `json:"size,omitempty"`
	// Family - The family of hardware.
	Family *string `json:"family,omitempty"`
}

// TrackedResource - Resource properties including location and tags for track resources.
type TrackedResource struct {
	// ID - Resource ID
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - The location the resource resides in.
	Location string `json:"location,omitempty"`
	// Tags - Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}
