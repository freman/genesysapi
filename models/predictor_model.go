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

// PredictorModel predictor model
//
// swagger:model PredictorModel
type PredictorModel struct {

	// DateTime indicating when the model was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// DateTime indicating when the model was last trained. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateTrained strfmt.DateTime `json:"dateTrained,omitempty"`

	// features
	Features []*PredictorModelFeature `json:"features"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The key performance indicator used in the model.
	// Read Only: true
	Kpi string `json:"kpi,omitempty"`

	// The media type of the model.
	// Read Only: true
	// Enum: [voice email message]
	MediaType string `json:"mediaType,omitempty"`

	// The List of Queues that are assessed for Predictive Routing.
	// Read Only: true
	Queues []*AddressableEntityRef `json:"queues"`
}

// Validate validates this predictor model
func (m *PredictorModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTrained(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PredictorModel) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PredictorModel) validateDateTrained(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTrained) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTrained", "body", "date-time", m.DateTrained.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PredictorModel) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {
		if swag.IsZero(m.Features[i]) { // not required
			continue
		}

		if m.Features[i] != nil {
			if err := m.Features[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var predictorModelTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["voice","email","message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		predictorModelTypeMediaTypePropEnum = append(predictorModelTypeMediaTypePropEnum, v)
	}
}

const (

	// PredictorModelMediaTypeVoice captures enum value "voice"
	PredictorModelMediaTypeVoice string = "voice"

	// PredictorModelMediaTypeEmail captures enum value "email"
	PredictorModelMediaTypeEmail string = "email"

	// PredictorModelMediaTypeMessage captures enum value "message"
	PredictorModelMediaTypeMessage string = "message"
)

// prop value enum
func (m *PredictorModel) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, predictorModelTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PredictorModel) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

func (m *PredictorModel) validateQueues(formats strfmt.Registry) error {

	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PredictorModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PredictorModel) UnmarshalBinary(b []byte) error {
	var res PredictorModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
