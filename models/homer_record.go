// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HomerRecord homer record
//
// swagger:model HomerRecord
type HomerRecord struct {

	// metadata associated to the SIP calls
	Auth string `json:"auth,omitempty"`

	// metadata associated to the SIP calls
	AuthUser string `json:"authUser,omitempty"`

	// metadata associated to the SIP calls
	Callid string `json:"callid,omitempty"`

	// metadata associated to the SIP calls
	CallidAleg string `json:"callidAleg,omitempty"`

	// metadata associated to the SIP calls
	ContactIP string `json:"contactIp,omitempty"`

	// metadata associated to the SIP calls
	ContactPort string `json:"contactPort,omitempty"`

	// metadata associated to the SIP calls
	ContactUser string `json:"contactUser,omitempty"`

	// metadata associated to the SIP calls
	ContentType string `json:"contentType,omitempty"`

	// metadata associated to the SIP calls
	ConversationID string `json:"conversationId,omitempty"`

	// metadata associated to the SIP calls
	CorrelationID string `json:"correlationId,omitempty"`

	// metadata associated to the SIP calls
	Cseq string `json:"cseq,omitempty"`

	// metadata associated to the SIP calls. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// metadata associated to the SIP calls
	Dbnode string `json:"dbnode,omitempty"`

	// metadata associated to the SIP calls
	DestinationAlias string `json:"destinationAlias,omitempty"`

	// metadata associated to the SIP calls
	DestinationIP string `json:"destinationIp,omitempty"`

	// metadata associated to the SIP calls
	DestinationPort string `json:"destinationPort,omitempty"`

	// metadata associated to the SIP calls
	Diversion string `json:"diversion,omitempty"`

	// metadata associated to the SIP calls
	Family string `json:"family,omitempty"`

	// metadata associated to the SIP calls
	FromDomain string `json:"fromDomain,omitempty"`

	// metadata associated to the SIP calls
	FromTag string `json:"fromTag,omitempty"`

	// metadata associated to the SIP calls
	FromUser string `json:"fromUser,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// metadata associated to the SIP calls
	Method string `json:"method,omitempty"`

	// metadata associated to the SIP calls
	MicroTs string `json:"microTs,omitempty"`

	// metadata associated to the SIP calls
	MilliTs string `json:"milliTs,omitempty"`

	// metadata associated to the SIP calls
	Msg string `json:"msg,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// metadata associated to the SIP calls
	Node string `json:"node,omitempty"`

	// metadata associated to the SIP calls
	OriginatorIP string `json:"originatorIp,omitempty"`

	// metadata associated to the SIP calls
	OriginatorPort string `json:"originatorPort,omitempty"`

	// metadata associated to the SIP calls
	ParticipantID string `json:"participantId,omitempty"`

	// metadata associated to the SIP calls
	PidUser string `json:"pidUser,omitempty"`

	// metadata associated to the SIP calls
	Proto string `json:"proto,omitempty"`

	// metadata associated to the SIP calls
	Reason string `json:"reason,omitempty"`

	// metadata associated to the SIP calls
	ReplyReason string `json:"replyReason,omitempty"`

	// metadata associated to the SIP calls
	RtpStat string `json:"rtpStat,omitempty"`

	// metadata associated to the SIP calls
	Ruri string `json:"ruri,omitempty"`

	// metadata associated to the SIP calls
	RuriDomain string `json:"ruriDomain,omitempty"`

	// metadata associated to the SIP calls
	RuriUser string `json:"ruriUser,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// metadata associated to the SIP calls
	SourceAlias string `json:"sourceAlias,omitempty"`

	// metadata associated to the SIP calls
	SourceIP string `json:"sourceIp,omitempty"`

	// metadata associated to the SIP calls
	SourcePort string `json:"sourcePort,omitempty"`

	// metadata associated to the SIP calls
	ToDomain string `json:"toDomain,omitempty"`

	// metadata associated to the SIP calls
	ToTag string `json:"toTag,omitempty"`

	// metadata associated to the SIP calls
	ToUser string `json:"toUser,omitempty"`

	// metadata associated to the SIP calls
	Trans string `json:"trans,omitempty"`

	// metadata associated to the SIP calls
	Type string `json:"type,omitempty"`

	// metadata associated to the SIP calls
	UserAgent string `json:"userAgent,omitempty"`

	// metadata associated to the SIP calls
	Via1 string `json:"via1,omitempty"`

	// metadata associated to the SIP calls
	Via1Branch string `json:"via1Branch,omitempty"`
}

// Validate validates this homer record
func (m *HomerRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
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

func (m *HomerRecord) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HomerRecord) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this homer record based on the context it is used
func (m *HomerRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
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

func (m *HomerRecord) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *HomerRecord) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HomerRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HomerRecord) UnmarshalBinary(b []byte) error {
	var res HomerRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
