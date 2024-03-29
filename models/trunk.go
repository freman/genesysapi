// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Trunk trunk
//
// swagger:model Trunk
type Trunk struct {

	// The connected status of the trunk
	// Read Only: true
	ConnectedStatus *TrunkConnectedStatus `json:"connectedStatus,omitempty"`

	// The ID of the user that created the resource.
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	// Read Only: true
	CreatedByApp string `json:"createdByApp,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The resource's description.
	Description string `json:"description,omitempty"`

	// The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// The Edge using this trunk.
	Edge *DomainEntityRef `json:"edge,omitempty"`

	// The edge group associated with this trunk.
	EdgeGroup *DomainEntityRef `json:"edgeGroup,omitempty"`

	// True if the Edge used by this trunk is in-service
	Enabled bool `json:"enabled"`

	// The IP Network Family of the trunk
	// Read Only: true
	Family int32 `json:"family,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// True if this trunk is in-service.  This comes from the trunk_enabled property of the referenced trunk base.
	// Read Only: true
	InService *bool `json:"inService"`

	// The trunk ipStatus
	// Read Only: true
	IPStatus *TrunkMetricsNetworkTypeIP `json:"ipStatus,omitempty"`

	// The Logical Interface on the Edge to which the trunk is assigned.
	// Read Only: true
	LogicalInterface *DomainEntityRef `json:"logicalInterface,omitempty"`

	// The ID of the user that last modified the resource.
	// Read Only: true
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	// Read Only: true
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// Returns Enabled when the trunk base supports the availability interval and it has a value greater than 0.
	// Read Only: true
	// Enum: [ENABLED DISABLED NOT_SUPPORTED]
	OptionsEnabledStatus string `json:"optionsEnabledStatus,omitempty"`

	// The trunk optionsStatus
	// Read Only: true
	OptionsStatus []*TrunkMetricsOptions `json:"optionsStatus"`

	// The list of proxy addresses (ports if provided) for the trunk
	// Read Only: true
	ProxyAddressList []string `json:"proxyAddressList"`

	// Returns Enabled when the trunk base supports the registration interval and it has a value greater than 0.
	// Read Only: true
	// Enum: [ENABLED DISABLED NOT_SUPPORTED]
	RegistersEnabledStatus string `json:"registersEnabledStatus,omitempty"`

	// The trunk registersStatus
	// Read Only: true
	RegistersStatus []*TrunkMetricsRegisters `json:"registersStatus"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// The trunk base configuration used on this trunk.
	TrunkBase *DomainEntityRef `json:"trunkBase,omitempty"`

	// The metabase used to create this trunk.
	TrunkMetabase *DomainEntityRef `json:"trunkMetabase,omitempty"`

	// The type of this trunk.
	// Enum: [EXTERNAL PHONE EDGE]
	TrunkType string `json:"trunkType,omitempty"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this trunk
func (m *Trunk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionsEnabledStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionsStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistersEnabledStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistersStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrunkBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrunkMetabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrunkType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Trunk) validateConnectedStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedStatus) { // not required
		return nil
	}

	if m.ConnectedStatus != nil {
		if err := m.ConnectedStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectedStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectedStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) validateDivision(formats strfmt.Registry) error {
	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateEdge(formats strfmt.Registry) error {
	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateEdgeGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeGroup) { // not required
		return nil
	}

	if m.EdgeGroup != nil {
		if err := m.EdgeGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateIPStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.IPStatus) { // not required
		return nil
	}

	if m.IPStatus != nil {
		if err := m.IPStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateLogicalInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.LogicalInterface) { // not required
		return nil
	}

	if m.LogicalInterface != nil {
		if err := m.LogicalInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logicalInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var trunkTypeOptionsEnabledStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","NOT_SUPPORTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trunkTypeOptionsEnabledStatusPropEnum = append(trunkTypeOptionsEnabledStatusPropEnum, v)
	}
}

const (

	// TrunkOptionsEnabledStatusENABLED captures enum value "ENABLED"
	TrunkOptionsEnabledStatusENABLED string = "ENABLED"

	// TrunkOptionsEnabledStatusDISABLED captures enum value "DISABLED"
	TrunkOptionsEnabledStatusDISABLED string = "DISABLED"

	// TrunkOptionsEnabledStatusNOTSUPPORTED captures enum value "NOT_SUPPORTED"
	TrunkOptionsEnabledStatusNOTSUPPORTED string = "NOT_SUPPORTED"
)

// prop value enum
func (m *Trunk) validateOptionsEnabledStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trunkTypeOptionsEnabledStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Trunk) validateOptionsEnabledStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OptionsEnabledStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOptionsEnabledStatusEnum("optionsEnabledStatus", "body", m.OptionsEnabledStatus); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) validateOptionsStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OptionsStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.OptionsStatus); i++ {
		if swag.IsZero(m.OptionsStatus[i]) { // not required
			continue
		}

		if m.OptionsStatus[i] != nil {
			if err := m.OptionsStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("optionsStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("optionsStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var trunkTypeRegistersEnabledStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","NOT_SUPPORTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trunkTypeRegistersEnabledStatusPropEnum = append(trunkTypeRegistersEnabledStatusPropEnum, v)
	}
}

const (

	// TrunkRegistersEnabledStatusENABLED captures enum value "ENABLED"
	TrunkRegistersEnabledStatusENABLED string = "ENABLED"

	// TrunkRegistersEnabledStatusDISABLED captures enum value "DISABLED"
	TrunkRegistersEnabledStatusDISABLED string = "DISABLED"

	// TrunkRegistersEnabledStatusNOTSUPPORTED captures enum value "NOT_SUPPORTED"
	TrunkRegistersEnabledStatusNOTSUPPORTED string = "NOT_SUPPORTED"
)

// prop value enum
func (m *Trunk) validateRegistersEnabledStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trunkTypeRegistersEnabledStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Trunk) validateRegistersEnabledStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistersEnabledStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegistersEnabledStatusEnum("registersEnabledStatus", "body", m.RegistersEnabledStatus); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) validateRegistersStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistersStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.RegistersStatus); i++ {
		if swag.IsZero(m.RegistersStatus[i]) { // not required
			continue
		}

		if m.RegistersStatus[i] != nil {
			if err := m.RegistersStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registersStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registersStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Trunk) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var trunkTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trunkTypeStatePropEnum = append(trunkTypeStatePropEnum, v)
	}
}

const (

	// TrunkStateActive captures enum value "active"
	TrunkStateActive string = "active"

	// TrunkStateInactive captures enum value "inactive"
	TrunkStateInactive string = "inactive"

	// TrunkStateDeleted captures enum value "deleted"
	TrunkStateDeleted string = "deleted"
)

// prop value enum
func (m *Trunk) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trunkTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Trunk) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) validateTrunkBase(formats strfmt.Registry) error {
	if swag.IsZero(m.TrunkBase) { // not required
		return nil
	}

	if m.TrunkBase != nil {
		if err := m.TrunkBase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkBase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trunkBase")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) validateTrunkMetabase(formats strfmt.Registry) error {
	if swag.IsZero(m.TrunkMetabase) { // not required
		return nil
	}

	if m.TrunkMetabase != nil {
		if err := m.TrunkMetabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkMetabase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trunkMetabase")
			}
			return err
		}
	}

	return nil
}

var trunkTypeTrunkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXTERNAL","PHONE","EDGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trunkTypeTrunkTypePropEnum = append(trunkTypeTrunkTypePropEnum, v)
	}
}

const (

	// TrunkTrunkTypeEXTERNAL captures enum value "EXTERNAL"
	TrunkTrunkTypeEXTERNAL string = "EXTERNAL"

	// TrunkTrunkTypePHONE captures enum value "PHONE"
	TrunkTrunkTypePHONE string = "PHONE"

	// TrunkTrunkTypeEDGE captures enum value "EDGE"
	TrunkTrunkTypeEDGE string = "EDGE"
)

// prop value enum
func (m *Trunk) validateTrunkTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trunkTypeTrunkTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Trunk) validateTrunkType(formats strfmt.Registry) error {
	if swag.IsZero(m.TrunkType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrunkTypeEnum("trunkType", "body", m.TrunkType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this trunk based on the context it is used
func (m *Trunk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectedStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedByApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogicalInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedByApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptionsEnabledStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptionsStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxyAddressList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistersEnabledStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistersStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrunkBase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrunkMetabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Trunk) contextValidateConnectedStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectedStatus != nil {
		if err := m.ConnectedStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectedStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectedStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateCreatedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdByApp", "body", string(m.CreatedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

	if m.Division != nil {
		if err := m.Division.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateEdge(ctx context.Context, formats strfmt.Registry) error {

	if m.Edge != nil {
		if err := m.Edge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateEdgeGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeGroup != nil {
		if err := m.EdgeGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "family", "body", int32(m.Family)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateInService(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "inService", "body", m.InService); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateIPStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.IPStatus != nil {
		if err := m.IPStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateLogicalInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.LogicalInterface != nil {
		if err := m.LogicalInterface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalInterface")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logicalInterface")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedBy", "body", string(m.ModifiedBy)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateModifiedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedByApp", "body", string(m.ModifiedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateOptionsEnabledStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "optionsEnabledStatus", "body", string(m.OptionsEnabledStatus)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateOptionsStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "optionsStatus", "body", []*TrunkMetricsOptions(m.OptionsStatus)); err != nil {
		return err
	}

	for i := 0; i < len(m.OptionsStatus); i++ {

		if m.OptionsStatus[i] != nil {
			if err := m.OptionsStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("optionsStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("optionsStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Trunk) contextValidateProxyAddressList(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "proxyAddressList", "body", []string(m.ProxyAddressList)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateRegistersEnabledStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "registersEnabledStatus", "body", string(m.RegistersEnabledStatus)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateRegistersStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "registersStatus", "body", []*TrunkMetricsRegisters(m.RegistersStatus)); err != nil {
		return err
	}

	for i := 0; i < len(m.RegistersStatus); i++ {

		if m.RegistersStatus[i] != nil {
			if err := m.RegistersStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registersStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registersStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Trunk) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Trunk) contextValidateTrunkBase(ctx context.Context, formats strfmt.Registry) error {

	if m.TrunkBase != nil {
		if err := m.TrunkBase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkBase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trunkBase")
			}
			return err
		}
	}

	return nil
}

func (m *Trunk) contextValidateTrunkMetabase(ctx context.Context, formats strfmt.Registry) error {

	if m.TrunkMetabase != nil {
		if err := m.TrunkMetabase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkMetabase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trunkMetabase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Trunk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Trunk) UnmarshalBinary(b []byte) error {
	var res Trunk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
