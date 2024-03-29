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

// Recipient recipient
//
// swagger:model Recipient
type Recipient struct {

	// User that created this recipient
	CreatedBy *User `json:"createdBy,omitempty"`

	// Date this recipient was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date this recipient was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// An automate flow object which defines the set of actions to be taken, when a message is received by this provisioned phone number.
	Flow *Flow `json:"flow,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The messenger type for this recipient
	// Enum: [sms facebook twitter line whatsapp webmessaging instagram open]
	MessengerType string `json:"messengerType,omitempty"`

	// User that modified this recipient
	ModifiedBy *User `json:"modifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this recipient
func (m *Recipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessengerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipient) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *Recipient) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.Flow) { // not required
		return nil
	}

	if m.Flow != nil {
		if err := m.Flow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

var recipientTypeMessengerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","facebook","twitter","line","whatsapp","webmessaging","instagram","open"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recipientTypeMessengerTypePropEnum = append(recipientTypeMessengerTypePropEnum, v)
	}
}

const (

	// RecipientMessengerTypeSms captures enum value "sms"
	RecipientMessengerTypeSms string = "sms"

	// RecipientMessengerTypeFacebook captures enum value "facebook"
	RecipientMessengerTypeFacebook string = "facebook"

	// RecipientMessengerTypeTwitter captures enum value "twitter"
	RecipientMessengerTypeTwitter string = "twitter"

	// RecipientMessengerTypeLine captures enum value "line"
	RecipientMessengerTypeLine string = "line"

	// RecipientMessengerTypeWhatsapp captures enum value "whatsapp"
	RecipientMessengerTypeWhatsapp string = "whatsapp"

	// RecipientMessengerTypeWebmessaging captures enum value "webmessaging"
	RecipientMessengerTypeWebmessaging string = "webmessaging"

	// RecipientMessengerTypeInstagram captures enum value "instagram"
	RecipientMessengerTypeInstagram string = "instagram"

	// RecipientMessengerTypeOpen captures enum value "open"
	RecipientMessengerTypeOpen string = "open"
)

// prop value enum
func (m *Recipient) validateMessengerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recipientTypeMessengerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Recipient) validateMessengerType(formats strfmt.Registry) error {
	if swag.IsZero(m.MessengerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessengerTypeEnum("messengerType", "body", m.MessengerType); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateModifiedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Recipient) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this recipient based on the context it is used
func (m *Recipient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipient) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {
		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *Recipient) contextValidateFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.Flow != nil {
		if err := m.Flow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *Recipient) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Recipient) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipient) UnmarshalBinary(b []byte) error {
	var res Recipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
