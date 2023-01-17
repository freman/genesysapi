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

// Flow flow
//
// swagger:model Flow
type Flow struct {

	// active
	Active bool `json:"active"`

	// checked in version
	CheckedInVersion *FlowVersion `json:"checkedInVersion,omitempty"`

	// Compatible flow types designate which flow types are allowed to embed a flow’s configuration within their own flow configuration.  Currently the only flows that can be embedded are Common Module flows and the embedding flow can invoke them using the Call Common Module action.
	// Read Only: true
	CompatibleFlowTypes []string `json:"compatibleFlowTypes"`

	// current operation
	CurrentOperation *Operation `json:"currentOperation,omitempty"`

	// debug version
	DebugVersion *FlowVersion `json:"debugVersion,omitempty"`

	// deleted
	Deleted bool `json:"deleted"`

	// description
	Description string `json:"description,omitempty"`

	// The division to which this entity belongs.
	Division *WritableDivision `json:"division,omitempty"`

	// The flow identifier
	ID string `json:"id,omitempty"`

	// json schema describing the inputs for the flow
	InputSchema interface{} `json:"inputSchema,omitempty"`

	// OAuth client that has the flow locked.
	LockedClient *DomainEntityRef `json:"lockedClient,omitempty"`

	// User that has the flow locked.
	LockedUser *User `json:"lockedUser,omitempty"`

	// The flow name
	// Required: true
	Name *string `json:"name"`

	// Information about the natural language understanding configuration for the published version of the flow
	// Read Only: true
	NluInfo *NluInfo `json:"nluInfo,omitempty"`

	// json schema describing the outputs for the flow
	OutputSchema interface{} `json:"outputSchema,omitempty"`

	// published by
	PublishedBy *User `json:"publishedBy,omitempty"`

	// published version
	PublishedVersion *FlowVersion `json:"publishedVersion,omitempty"`

	// saved version
	SavedVersion *FlowVersion `json:"savedVersion,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// List of supported languages for the published version of the flow.
	// Read Only: true
	SupportedLanguages []*SupportedLanguage `json:"supportedLanguages"`

	// system
	System bool `json:"system"`

	// type
	// Enum: [BOT COMMONMODULE DIGITALBOT INBOUNDCALL INBOUNDCHAT INBOUNDEMAIL INBOUNDSHORTMESSAGE INQUEUECALL INQUEUEEMAIL INQUEUESHORTMESSAGE OUTBOUNDCALL SECURECALL SPEECH SURVEYINVITE VOICE VOICEMAIL WORKFLOW WORKITEM]
	Type string `json:"type,omitempty"`
}

// Validate validates this flow
func (m *Flow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckedInVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompatibleFlowTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebugVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockedClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNluInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSavedVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flow) validateCheckedInVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckedInVersion) { // not required
		return nil
	}

	if m.CheckedInVersion != nil {
		if err := m.CheckedInVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkedInVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkedInVersion")
			}
			return err
		}
	}

	return nil
}

var flowCompatibleFlowTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BOT","COMMONMODULE","DIGITALBOT","INBOUNDCALL","INBOUNDCHAT","INBOUNDEMAIL","INBOUNDSHORTMESSAGE","INQUEUECALL","INQUEUEEMAIL","INQUEUESHORTMESSAGE","OUTBOUNDCALL","SECURECALL","SPEECH","SURVEYINVITE","VOICE","VOICEMAIL","WORKFLOW","WORKITEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowCompatibleFlowTypesItemsEnum = append(flowCompatibleFlowTypesItemsEnum, v)
	}
}

func (m *Flow) validateCompatibleFlowTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowCompatibleFlowTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Flow) validateCompatibleFlowTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.CompatibleFlowTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.CompatibleFlowTypes); i++ {

		// value enum
		if err := m.validateCompatibleFlowTypesItemsEnum("compatibleFlowTypes"+"."+strconv.Itoa(i), "body", m.CompatibleFlowTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Flow) validateCurrentOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentOperation) { // not required
		return nil
	}

	if m.CurrentOperation != nil {
		if err := m.CurrentOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentOperation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentOperation")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validateDebugVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.DebugVersion) { // not required
		return nil
	}

	if m.DebugVersion != nil {
		if err := m.DebugVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debugVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("debugVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validateDivision(formats strfmt.Registry) error {
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

func (m *Flow) validateLockedClient(formats strfmt.Registry) error {
	if swag.IsZero(m.LockedClient) { // not required
		return nil
	}

	if m.LockedClient != nil {
		if err := m.LockedClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lockedClient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lockedClient")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validateLockedUser(formats strfmt.Registry) error {
	if swag.IsZero(m.LockedUser) { // not required
		return nil
	}

	if m.LockedUser != nil {
		if err := m.LockedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lockedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lockedUser")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Flow) validateNluInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.NluInfo) { // not required
		return nil
	}

	if m.NluInfo != nil {
		if err := m.NluInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nluInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nluInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validatePublishedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishedBy) { // not required
		return nil
	}

	if m.PublishedBy != nil {
		if err := m.PublishedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publishedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validatePublishedVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishedVersion) { // not required
		return nil
	}

	if m.PublishedVersion != nil {
		if err := m.PublishedVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publishedVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validateSavedVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.SavedVersion) { // not required
		return nil
	}

	if m.SavedVersion != nil {
		if err := m.SavedVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("savedVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("savedVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Flow) validateSupportedLanguages(formats strfmt.Registry) error {
	if swag.IsZero(m.SupportedLanguages) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedLanguages); i++ {
		if swag.IsZero(m.SupportedLanguages[i]) { // not required
			continue
		}

		if m.SupportedLanguages[i] != nil {
			if err := m.SupportedLanguages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportedLanguages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supportedLanguages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var flowTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BOT","COMMONMODULE","DIGITALBOT","INBOUNDCALL","INBOUNDCHAT","INBOUNDEMAIL","INBOUNDSHORTMESSAGE","INQUEUECALL","INQUEUEEMAIL","INQUEUESHORTMESSAGE","OUTBOUNDCALL","SECURECALL","SPEECH","SURVEYINVITE","VOICE","VOICEMAIL","WORKFLOW","WORKITEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowTypeTypePropEnum = append(flowTypeTypePropEnum, v)
	}
}

const (

	// FlowTypeBOT captures enum value "BOT"
	FlowTypeBOT string = "BOT"

	// FlowTypeCOMMONMODULE captures enum value "COMMONMODULE"
	FlowTypeCOMMONMODULE string = "COMMONMODULE"

	// FlowTypeDIGITALBOT captures enum value "DIGITALBOT"
	FlowTypeDIGITALBOT string = "DIGITALBOT"

	// FlowTypeINBOUNDCALL captures enum value "INBOUNDCALL"
	FlowTypeINBOUNDCALL string = "INBOUNDCALL"

	// FlowTypeINBOUNDCHAT captures enum value "INBOUNDCHAT"
	FlowTypeINBOUNDCHAT string = "INBOUNDCHAT"

	// FlowTypeINBOUNDEMAIL captures enum value "INBOUNDEMAIL"
	FlowTypeINBOUNDEMAIL string = "INBOUNDEMAIL"

	// FlowTypeINBOUNDSHORTMESSAGE captures enum value "INBOUNDSHORTMESSAGE"
	FlowTypeINBOUNDSHORTMESSAGE string = "INBOUNDSHORTMESSAGE"

	// FlowTypeINQUEUECALL captures enum value "INQUEUECALL"
	FlowTypeINQUEUECALL string = "INQUEUECALL"

	// FlowTypeINQUEUEEMAIL captures enum value "INQUEUEEMAIL"
	FlowTypeINQUEUEEMAIL string = "INQUEUEEMAIL"

	// FlowTypeINQUEUESHORTMESSAGE captures enum value "INQUEUESHORTMESSAGE"
	FlowTypeINQUEUESHORTMESSAGE string = "INQUEUESHORTMESSAGE"

	// FlowTypeOUTBOUNDCALL captures enum value "OUTBOUNDCALL"
	FlowTypeOUTBOUNDCALL string = "OUTBOUNDCALL"

	// FlowTypeSECURECALL captures enum value "SECURECALL"
	FlowTypeSECURECALL string = "SECURECALL"

	// FlowTypeSPEECH captures enum value "SPEECH"
	FlowTypeSPEECH string = "SPEECH"

	// FlowTypeSURVEYINVITE captures enum value "SURVEYINVITE"
	FlowTypeSURVEYINVITE string = "SURVEYINVITE"

	// FlowTypeVOICE captures enum value "VOICE"
	FlowTypeVOICE string = "VOICE"

	// FlowTypeVOICEMAIL captures enum value "VOICEMAIL"
	FlowTypeVOICEMAIL string = "VOICEMAIL"

	// FlowTypeWORKFLOW captures enum value "WORKFLOW"
	FlowTypeWORKFLOW string = "WORKFLOW"

	// FlowTypeWORKITEM captures enum value "WORKITEM"
	FlowTypeWORKITEM string = "WORKITEM"
)

// prop value enum
func (m *Flow) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Flow) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this flow based on the context it is used
func (m *Flow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCheckedInVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompatibleFlowTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrentOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDebugVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLockedClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLockedUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNluInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublishedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublishedVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSavedVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportedLanguages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flow) contextValidateCheckedInVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.CheckedInVersion != nil {
		if err := m.CheckedInVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkedInVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkedInVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateCompatibleFlowTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "compatibleFlowTypes", "body", []string(m.CompatibleFlowTypes)); err != nil {
		return err
	}

	return nil
}

func (m *Flow) contextValidateCurrentOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentOperation != nil {
		if err := m.CurrentOperation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentOperation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentOperation")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateDebugVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.DebugVersion != nil {
		if err := m.DebugVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debugVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("debugVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Flow) contextValidateLockedClient(ctx context.Context, formats strfmt.Registry) error {

	if m.LockedClient != nil {
		if err := m.LockedClient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lockedClient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lockedClient")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateLockedUser(ctx context.Context, formats strfmt.Registry) error {

	if m.LockedUser != nil {
		if err := m.LockedUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lockedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lockedUser")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateNluInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.NluInfo != nil {
		if err := m.NluInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nluInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nluInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidatePublishedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.PublishedBy != nil {
		if err := m.PublishedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publishedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidatePublishedVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.PublishedVersion != nil {
		if err := m.PublishedVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publishedVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateSavedVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.SavedVersion != nil {
		if err := m.SavedVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("savedVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("savedVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Flow) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Flow) contextValidateSupportedLanguages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supportedLanguages", "body", []*SupportedLanguage(m.SupportedLanguages)); err != nil {
		return err
	}

	for i := 0; i < len(m.SupportedLanguages); i++ {

		if m.SupportedLanguages[i] != nil {
			if err := m.SupportedLanguages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportedLanguages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supportedLanguages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Flow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Flow) UnmarshalBinary(b []byte) error {
	var res Flow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
