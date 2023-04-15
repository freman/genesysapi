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

// OpenNormalizedMessage Open Messaging rich media message structure
//
// swagger:model OpenNormalizedMessage
type OpenNormalizedMessage struct {

	// Channel-specific information that describes the message and the message channel/provider.
	// Required: true
	Channel *OpenMessagingChannel `json:"channel"`

	// List of content elements.
	Content []*OpenMessageContent `json:"content"`

	// The direction of the message.
	// Enum: [Inbound Outbound]
	Direction string `json:"direction,omitempty"`

	// Unique ID of the message. This ID is generated by Messaging Platform. Message receipts will have the same ID as the message they reference, as such should only be set when sending a message receipt.
	ID string `json:"id,omitempty"`

	// Indicates if this is the last message receipt for this message, or if another message receipt can be expected.
	IsFinalReceipt bool `json:"isFinalReceipt"`

	// Additional metadata about this message.
	Metadata map[string]string `json:"metadata,omitempty"`

	// List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
	Reasons []*ConversationReason `json:"reasons"`

	// Message receipt status, only used with type Receipt.
	// Enum: [Sent Delivered Read Failed Published Removed]
	Status string `json:"status,omitempty"`

	// Message text.
	Text string `json:"text,omitempty"`

	// Message type.
	// Required: true
	// Enum: [Text Receipt]
	Type *string `json:"type"`
}

// Validate validates this open normalized message
func (m *OpenNormalizedMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *OpenNormalizedMessage) validateChannel(formats strfmt.Registry) error {

	if err := validate.Required("channel", "body", m.Channel); err != nil {
		return err
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *OpenNormalizedMessage) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var openNormalizedMessageTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Inbound","Outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openNormalizedMessageTypeDirectionPropEnum = append(openNormalizedMessageTypeDirectionPropEnum, v)
	}
}

const (

	// OpenNormalizedMessageDirectionInbound captures enum value "Inbound"
	OpenNormalizedMessageDirectionInbound string = "Inbound"

	// OpenNormalizedMessageDirectionOutbound captures enum value "Outbound"
	OpenNormalizedMessageDirectionOutbound string = "Outbound"
)

// prop value enum
func (m *OpenNormalizedMessage) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openNormalizedMessageTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenNormalizedMessage) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *OpenNormalizedMessage) validateReasons(formats strfmt.Registry) error {
	if swag.IsZero(m.Reasons) { // not required
		return nil
	}

	for i := 0; i < len(m.Reasons); i++ {
		if swag.IsZero(m.Reasons[i]) { // not required
			continue
		}

		if m.Reasons[i] != nil {
			if err := m.Reasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var openNormalizedMessageTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sent","Delivered","Read","Failed","Published","Removed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openNormalizedMessageTypeStatusPropEnum = append(openNormalizedMessageTypeStatusPropEnum, v)
	}
}

const (

	// OpenNormalizedMessageStatusSent captures enum value "Sent"
	OpenNormalizedMessageStatusSent string = "Sent"

	// OpenNormalizedMessageStatusDelivered captures enum value "Delivered"
	OpenNormalizedMessageStatusDelivered string = "Delivered"

	// OpenNormalizedMessageStatusRead captures enum value "Read"
	OpenNormalizedMessageStatusRead string = "Read"

	// OpenNormalizedMessageStatusFailed captures enum value "Failed"
	OpenNormalizedMessageStatusFailed string = "Failed"

	// OpenNormalizedMessageStatusPublished captures enum value "Published"
	OpenNormalizedMessageStatusPublished string = "Published"

	// OpenNormalizedMessageStatusRemoved captures enum value "Removed"
	OpenNormalizedMessageStatusRemoved string = "Removed"
)

// prop value enum
func (m *OpenNormalizedMessage) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openNormalizedMessageTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenNormalizedMessage) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var openNormalizedMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text","Receipt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openNormalizedMessageTypeTypePropEnum = append(openNormalizedMessageTypeTypePropEnum, v)
	}
}

const (

	// OpenNormalizedMessageTypeText captures enum value "Text"
	OpenNormalizedMessageTypeText string = "Text"

	// OpenNormalizedMessageTypeReceipt captures enum value "Receipt"
	OpenNormalizedMessageTypeReceipt string = "Receipt"
)

// prop value enum
func (m *OpenNormalizedMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openNormalizedMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenNormalizedMessage) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this open normalized message based on the context it is used
func (m *OpenNormalizedMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReasons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenNormalizedMessage) contextValidateChannel(ctx context.Context, formats strfmt.Registry) error {

	if m.Channel != nil {
		if err := m.Channel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *OpenNormalizedMessage) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Content); i++ {

		if m.Content[i] != nil {
			if err := m.Content[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenNormalizedMessage) contextValidateReasons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Reasons); i++ {

		if m.Reasons[i] != nil {
			if err := m.Reasons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenNormalizedMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenNormalizedMessage) UnmarshalBinary(b []byte) error {
	var res OpenNormalizedMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
