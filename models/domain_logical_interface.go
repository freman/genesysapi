// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainLogicalInterface domain logical interface
//
// swagger:model DomainLogicalInterface
type DomainLogicalInterface struct {

	// The list of IP addresses on this interface.  Priority of dns addresses are based on order in the list.
	Addresses []*DomainNetworkAddress `json:"addresses"`

	// command responses
	CommandResponses []*DomainNetworkCommandResponse `json:"commandResponses"`

	// The ID of the user that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	CreatedByApp string `json:"createdByApp,omitempty"`

	// current state
	// Enum: [INIT CREATING UPDATING OK EXCEPTION DELETING]
	CurrentState string `json:"currentState,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The resource's description.
	Description string `json:"description,omitempty"`

	// edge assigned Id
	EdgeAssignedID string `json:"edgeAssignedId,omitempty"`

	// edge Uri
	// Format: uri
	EdgeURI strfmt.URI `json:"edgeUri,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// External trunk base settings to use for external communication from this interface.
	ExternalTrunkBaseAssignments []*TrunkBaseAssignment `json:"externalTrunkBaseAssignments"`

	// Friendly Name
	// Required: true
	FriendlyName *string `json:"friendlyName"`

	// Hardware Address
	// Required: true
	HardwareAddress *string `json:"hardwareAddress"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// if status
	IfStatus string `json:"ifStatus,omitempty"`

	// The IPv4 phone trunk base assignment will be inherited from the Edge Group.
	InheritPhoneTrunkBasesIPV4 bool `json:"inheritPhoneTrunkBasesIPv4"`

	// The IPv6 phone trunk base assignment will be inherited from the Edge Group.
	InheritPhoneTrunkBasesIPV6 bool `json:"inheritPhoneTrunkBasesIPv6"`

	// The type of this network interface.
	// Read Only: true
	// Enum: [DIAGNOSTIC SYSTEM]
	InterfaceType string `json:"interfaceType,omitempty"`

	// IPv4 interface settings.
	IPV4Capabilities *DomainCapabilities `json:"ipv4Capabilities,omitempty"`

	// IPv6 interface settings.
	IPV6Capabilities *DomainCapabilities `json:"ipv6Capabilities,omitempty"`

	// last modified correlation Id
	LastModifiedCorrelationID string `json:"lastModifiedCorrelationId,omitempty"`

	// last modified user Id
	LastModifiedUserID string `json:"lastModifiedUserId,omitempty"`

	// The ID of the user that last modified the resource.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// Phone trunk base settings to use for phone communication from this interface.  These settings will be ignored when "inheritPhoneTrunkBases" is true.
	PhoneTrunkBaseAssignments []*TrunkBaseAssignment `json:"phoneTrunkBaseAssignments"`

	// Physical Adapter Id
	// Required: true
	PhysicalAdapterID *string `json:"physicalAdapterId"`

	// IPv4 NENT IP Address
	PublicNatAddressIPV4 string `json:"publicNatAddressIpV4,omitempty"`

	// IPv6 NENT IP Address
	PublicNatAddressIPV6 string `json:"publicNatAddressIpV6,omitempty"`

	// The list of routes assigned to this interface.
	Routes []*DomainNetworkRoute `json:"routes"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// trace enabled
	TraceEnabled bool `json:"traceEnabled"`

	// Site Interconnects using the "Cloud Proxy" method will broker the connection between them with a Cloud Proxy. This method is required for connections between one or more Sites using Cloud Media, but can optionally be used between two premises Sites if Direct or Indirect are not an option.
	UseForCloudProxyEdgeCommunication bool `json:"useForCloudProxyEdgeCommunication"`

	// Site Interconnects using the "Indirect" method will communicate using the Public IP Address specified on the interface. Use this option when a NAT enabled firewall is between the Edge and the far end.
	UseForIndirectEdgeCommunication bool `json:"useForIndirectEdgeCommunication"`

	// This interface will be used for all internal edge-to-edge communication using settings from the edgeTrunkBaseAssignment on the Edge Group.
	UseForInternalEdgeCommunication bool `json:"useForInternalEdgeCommunication"`

	// This interface will be used for all communication with the internet.
	// Read Only: true
	UseForWanInterface *bool `json:"useForWanInterface"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`

	// vlan tag Id
	VlanTagID int32 `json:"vlanTagId,omitempty"`
}

// Validate validates this domain logical interface
func (m *DomainLogicalInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommandResponses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalTrunkBaseAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFriendlyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardwareAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4Capabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6Capabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneTrunkBaseAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalAdapterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainLogicalInterface) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainLogicalInterface) validateCommandResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.CommandResponses) { // not required
		return nil
	}

	for i := 0; i < len(m.CommandResponses); i++ {
		if swag.IsZero(m.CommandResponses[i]) { // not required
			continue
		}

		if m.CommandResponses[i] != nil {
			if err := m.CommandResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commandResponses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var domainLogicalInterfaceTypeCurrentStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INIT","CREATING","UPDATING","OK","EXCEPTION","DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainLogicalInterfaceTypeCurrentStatePropEnum = append(domainLogicalInterfaceTypeCurrentStatePropEnum, v)
	}
}

const (

	// DomainLogicalInterfaceCurrentStateINIT captures enum value "INIT"
	DomainLogicalInterfaceCurrentStateINIT string = "INIT"

	// DomainLogicalInterfaceCurrentStateCREATING captures enum value "CREATING"
	DomainLogicalInterfaceCurrentStateCREATING string = "CREATING"

	// DomainLogicalInterfaceCurrentStateUPDATING captures enum value "UPDATING"
	DomainLogicalInterfaceCurrentStateUPDATING string = "UPDATING"

	// DomainLogicalInterfaceCurrentStateOK captures enum value "OK"
	DomainLogicalInterfaceCurrentStateOK string = "OK"

	// DomainLogicalInterfaceCurrentStateEXCEPTION captures enum value "EXCEPTION"
	DomainLogicalInterfaceCurrentStateEXCEPTION string = "EXCEPTION"

	// DomainLogicalInterfaceCurrentStateDELETING captures enum value "DELETING"
	DomainLogicalInterfaceCurrentStateDELETING string = "DELETING"
)

// prop value enum
func (m *DomainLogicalInterface) validateCurrentStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainLogicalInterfaceTypeCurrentStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainLogicalInterface) validateCurrentState(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentState) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrentStateEnum("currentState", "body", m.CurrentState); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateEdgeURI(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeURI) { // not required
		return nil
	}

	if err := validate.FormatOf("edgeUri", "body", "uri", m.EdgeURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateExternalTrunkBaseAssignments(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalTrunkBaseAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalTrunkBaseAssignments); i++ {
		if swag.IsZero(m.ExternalTrunkBaseAssignments[i]) { // not required
			continue
		}

		if m.ExternalTrunkBaseAssignments[i] != nil {
			if err := m.ExternalTrunkBaseAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalTrunkBaseAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainLogicalInterface) validateFriendlyName(formats strfmt.Registry) error {

	if err := validate.Required("friendlyName", "body", m.FriendlyName); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateHardwareAddress(formats strfmt.Registry) error {

	if err := validate.Required("hardwareAddress", "body", m.HardwareAddress); err != nil {
		return err
	}

	return nil
}

var domainLogicalInterfaceTypeInterfaceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIAGNOSTIC","SYSTEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainLogicalInterfaceTypeInterfaceTypePropEnum = append(domainLogicalInterfaceTypeInterfaceTypePropEnum, v)
	}
}

const (

	// DomainLogicalInterfaceInterfaceTypeDIAGNOSTIC captures enum value "DIAGNOSTIC"
	DomainLogicalInterfaceInterfaceTypeDIAGNOSTIC string = "DIAGNOSTIC"

	// DomainLogicalInterfaceInterfaceTypeSYSTEM captures enum value "SYSTEM"
	DomainLogicalInterfaceInterfaceTypeSYSTEM string = "SYSTEM"
)

// prop value enum
func (m *DomainLogicalInterface) validateInterfaceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainLogicalInterfaceTypeInterfaceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainLogicalInterface) validateInterfaceType(formats strfmt.Registry) error {

	if swag.IsZero(m.InterfaceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateInterfaceTypeEnum("interfaceType", "body", m.InterfaceType); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateIPV4Capabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4Capabilities) { // not required
		return nil
	}

	if m.IPV4Capabilities != nil {
		if err := m.IPV4Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4Capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *DomainLogicalInterface) validateIPV6Capabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6Capabilities) { // not required
		return nil
	}

	if m.IPV6Capabilities != nil {
		if err := m.IPV6Capabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6Capabilities")
			}
			return err
		}
	}

	return nil
}

func (m *DomainLogicalInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validatePhoneTrunkBaseAssignments(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneTrunkBaseAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.PhoneTrunkBaseAssignments); i++ {
		if swag.IsZero(m.PhoneTrunkBaseAssignments[i]) { // not required
			continue
		}

		if m.PhoneTrunkBaseAssignments[i] != nil {
			if err := m.PhoneTrunkBaseAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneTrunkBaseAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainLogicalInterface) validatePhysicalAdapterID(formats strfmt.Registry) error {

	if err := validate.Required("physicalAdapterId", "body", m.PhysicalAdapterID); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainLogicalInterface) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainLogicalInterface) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var domainLogicalInterfaceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainLogicalInterfaceTypeStatePropEnum = append(domainLogicalInterfaceTypeStatePropEnum, v)
	}
}

const (

	// DomainLogicalInterfaceStateActive captures enum value "active"
	DomainLogicalInterfaceStateActive string = "active"

	// DomainLogicalInterfaceStateInactive captures enum value "inactive"
	DomainLogicalInterfaceStateInactive string = "inactive"

	// DomainLogicalInterfaceStateDeleted captures enum value "deleted"
	DomainLogicalInterfaceStateDeleted string = "deleted"
)

// prop value enum
func (m *DomainLogicalInterface) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainLogicalInterfaceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainLogicalInterface) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainLogicalInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainLogicalInterface) UnmarshalBinary(b []byte) error {
	var res DomainLogicalInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
