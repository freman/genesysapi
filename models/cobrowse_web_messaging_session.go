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

// CobrowseWebMessagingSession cobrowse web messaging session
//
// swagger:model CobrowseWebMessagingSession
type CobrowseWebMessagingSession struct {

	// CommunicationType for Cobrowse Session
	// Read Only: true
	// Enum: [Call Message]
	CommunicationType string `json:"communicationType,omitempty"`

	// Date when Cobrowse Offer Expires. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateOfferEnds strfmt.DateTime `json:"dateOfferEnds,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Cobrowse session join code
	// Read Only: true
	JoinCode string `json:"joinCode,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// WebSocket URL for the JS client
	// Read Only: true
	WebsocketURL string `json:"websocketUrl,omitempty"`
}

// Validate validates this cobrowse web messaging session
func (m *CobrowseWebMessagingSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunicationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateOfferEnds(formats); err != nil {
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

var cobrowseWebMessagingSessionTypeCommunicationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Call","Message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cobrowseWebMessagingSessionTypeCommunicationTypePropEnum = append(cobrowseWebMessagingSessionTypeCommunicationTypePropEnum, v)
	}
}

const (

	// CobrowseWebMessagingSessionCommunicationTypeCall captures enum value "Call"
	CobrowseWebMessagingSessionCommunicationTypeCall string = "Call"

	// CobrowseWebMessagingSessionCommunicationTypeMessage captures enum value "Message"
	CobrowseWebMessagingSessionCommunicationTypeMessage string = "Message"
)

// prop value enum
func (m *CobrowseWebMessagingSession) validateCommunicationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cobrowseWebMessagingSessionTypeCommunicationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CobrowseWebMessagingSession) validateCommunicationType(formats strfmt.Registry) error {
	if swag.IsZero(m.CommunicationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCommunicationTypeEnum("communicationType", "body", m.CommunicationType); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) validateDateOfferEnds(formats strfmt.Registry) error {
	if swag.IsZero(m.DateOfferEnds) { // not required
		return nil
	}

	if err := validate.FormatOf("dateOfferEnds", "body", "date-time", m.DateOfferEnds.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cobrowse web messaging session based on the context it is used
func (m *CobrowseWebMessagingSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommunicationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateOfferEnds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJoinCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebsocketURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CobrowseWebMessagingSession) contextValidateCommunicationType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "communicationType", "body", string(m.CommunicationType)); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) contextValidateDateOfferEnds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateOfferEnds", "body", strfmt.DateTime(m.DateOfferEnds)); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) contextValidateJoinCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "joinCode", "body", string(m.JoinCode)); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseWebMessagingSession) contextValidateWebsocketURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "websocketUrl", "body", string(m.WebsocketURL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CobrowseWebMessagingSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CobrowseWebMessagingSession) UnmarshalBinary(b []byte) error {
	var res CobrowseWebMessagingSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
