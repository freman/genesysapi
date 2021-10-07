// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TranscriptAggregationView transcript aggregation view
//
// swagger:model TranscriptAggregationView
type TranscriptAggregationView struct {

	// Type of view you wish to create
	// Required: true
	// Enum: [rangeBound]
	Function *string `json:"function"`

	// A unique name for this view. Must be distinct from other views and built-in metric names.
	// Required: true
	Name *string `json:"name"`

	// Range of numbers for slicing up data
	Range *AggregationRange `json:"range,omitempty"`

	// Target metric name
	// Required: true
	// Enum: [nTopicCommunications oCustomerSentiment oSentimentScore]
	Target *string `json:"target"`
}

// Validate validates this transcript aggregation view
func (m *TranscriptAggregationView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var transcriptAggregationViewTypeFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rangeBound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptAggregationViewTypeFunctionPropEnum = append(transcriptAggregationViewTypeFunctionPropEnum, v)
	}
}

const (

	// TranscriptAggregationViewFunctionRangeBound captures enum value "rangeBound"
	TranscriptAggregationViewFunctionRangeBound string = "rangeBound"
)

// prop value enum
func (m *TranscriptAggregationView) validateFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptAggregationViewTypeFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptAggregationView) validateFunction(formats strfmt.Registry) error {

	if err := validate.Required("function", "body", m.Function); err != nil {
		return err
	}

	// value enum
	if err := m.validateFunctionEnum("function", "body", *m.Function); err != nil {
		return err
	}

	return nil
}

func (m *TranscriptAggregationView) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TranscriptAggregationView) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

var transcriptAggregationViewTypeTargetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nTopicCommunications","oCustomerSentiment","oSentimentScore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptAggregationViewTypeTargetPropEnum = append(transcriptAggregationViewTypeTargetPropEnum, v)
	}
}

const (

	// TranscriptAggregationViewTargetNTopicCommunications captures enum value "nTopicCommunications"
	TranscriptAggregationViewTargetNTopicCommunications string = "nTopicCommunications"

	// TranscriptAggregationViewTargetOCustomerSentiment captures enum value "oCustomerSentiment"
	TranscriptAggregationViewTargetOCustomerSentiment string = "oCustomerSentiment"

	// TranscriptAggregationViewTargetOSentimentScore captures enum value "oSentimentScore"
	TranscriptAggregationViewTargetOSentimentScore string = "oSentimentScore"
)

// prop value enum
func (m *TranscriptAggregationView) validateTargetEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptAggregationViewTypeTargetPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptAggregationView) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnum("target", "body", *m.Target); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TranscriptAggregationView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptAggregationView) UnmarshalBinary(b []byte) error {
	var res TranscriptAggregationView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
