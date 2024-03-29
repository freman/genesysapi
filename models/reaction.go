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

// Reaction reaction
//
// swagger:model Reaction
type Reaction struct {

	// Parameter for this reaction. For transfer_flow, this would be the outbound flow id.
	Data string `json:"data,omitempty"`

	// Name of the parameter for this reaction. For transfer_flow, this would be the outbound flow name.
	Name string `json:"name,omitempty"`

	// The reaction to take for a given call analysis result.
	// Required: true
	// Enum: [hangup transfer transfer_flow play_file]
	ReactionType *string `json:"reactionType"`
}

// Validate validates this reaction
func (m *Reaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReactionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reactionTypeReactionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hangup","transfer","transfer_flow","play_file"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reactionTypeReactionTypePropEnum = append(reactionTypeReactionTypePropEnum, v)
	}
}

const (

	// ReactionReactionTypeHangup captures enum value "hangup"
	ReactionReactionTypeHangup string = "hangup"

	// ReactionReactionTypeTransfer captures enum value "transfer"
	ReactionReactionTypeTransfer string = "transfer"

	// ReactionReactionTypeTransferFlow captures enum value "transfer_flow"
	ReactionReactionTypeTransferFlow string = "transfer_flow"

	// ReactionReactionTypePlayFile captures enum value "play_file"
	ReactionReactionTypePlayFile string = "play_file"
)

// prop value enum
func (m *Reaction) validateReactionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reactionTypeReactionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Reaction) validateReactionType(formats strfmt.Registry) error {

	if err := validate.Required("reactionType", "body", m.ReactionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateReactionTypeEnum("reactionType", "body", *m.ReactionType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this reaction based on context it is used
func (m *Reaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Reaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reaction) UnmarshalBinary(b []byte) error {
	var res Reaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
