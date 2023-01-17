// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoryEntry history entry
//
// swagger:model HistoryEntry
type HistoryEntry struct {

	// The action performed
	// Enum: [CHECKIN CHECKOUT CREATE DEACTIVATE DEBUG DELETE PUBLISH REVERT SAVE TRANSCODE UPDATE UPLOAD]
	Action string `json:"action,omitempty"`

	// OAuth client associated with this entry.
	Client *DomainEntityRef `json:"client,omitempty"`

	// For actions performed not on the item itself, but on a sub-item, this field identifies the sub-item by name.  For example, for actions performed on prompt resources, this will be the prompt resource name.
	Resource string `json:"resource,omitempty"`

	// secure
	Secure bool `json:"secure"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// User associated with this entry.
	User *User `json:"user,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this history entry
func (m *HistoryEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var historyEntryTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CHECKIN","CHECKOUT","CREATE","DEACTIVATE","DEBUG","DELETE","PUBLISH","REVERT","SAVE","TRANSCODE","UPDATE","UPLOAD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historyEntryTypeActionPropEnum = append(historyEntryTypeActionPropEnum, v)
	}
}

const (

	// HistoryEntryActionCHECKIN captures enum value "CHECKIN"
	HistoryEntryActionCHECKIN string = "CHECKIN"

	// HistoryEntryActionCHECKOUT captures enum value "CHECKOUT"
	HistoryEntryActionCHECKOUT string = "CHECKOUT"

	// HistoryEntryActionCREATE captures enum value "CREATE"
	HistoryEntryActionCREATE string = "CREATE"

	// HistoryEntryActionDEACTIVATE captures enum value "DEACTIVATE"
	HistoryEntryActionDEACTIVATE string = "DEACTIVATE"

	// HistoryEntryActionDEBUG captures enum value "DEBUG"
	HistoryEntryActionDEBUG string = "DEBUG"

	// HistoryEntryActionDELETE captures enum value "DELETE"
	HistoryEntryActionDELETE string = "DELETE"

	// HistoryEntryActionPUBLISH captures enum value "PUBLISH"
	HistoryEntryActionPUBLISH string = "PUBLISH"

	// HistoryEntryActionREVERT captures enum value "REVERT"
	HistoryEntryActionREVERT string = "REVERT"

	// HistoryEntryActionSAVE captures enum value "SAVE"
	HistoryEntryActionSAVE string = "SAVE"

	// HistoryEntryActionTRANSCODE captures enum value "TRANSCODE"
	HistoryEntryActionTRANSCODE string = "TRANSCODE"

	// HistoryEntryActionUPDATE captures enum value "UPDATE"
	HistoryEntryActionUPDATE string = "UPDATE"

	// HistoryEntryActionUPLOAD captures enum value "UPLOAD"
	HistoryEntryActionUPLOAD string = "UPLOAD"
)

// prop value enum
func (m *HistoryEntry) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historyEntryTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoryEntry) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *HistoryEntry) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryEntry) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoryEntry) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this history entry based on the context it is used
func (m *HistoryEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryEntry) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *HistoryEntry) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoryEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryEntry) UnmarshalBinary(b []byte) error {
	var res HistoryEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
