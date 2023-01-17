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

// AnalyticsEvaluation analytics evaluation
//
// swagger:model AnalyticsEvaluation
type AnalyticsEvaluation struct {

	// The calibration ID used for the purpose of training evaluators
	CalibrationID string `json:"calibrationId,omitempty"`

	// A unique identifier for an evaluation form, regardless of version
	ContextID string `json:"contextId,omitempty"`

	// Whether the evaluation has been deleted
	Deleted bool `json:"deleted"`

	// Unique identifier for the evaluation
	EvaluationID string `json:"evaluationId,omitempty"`

	// A unique identifier of the user who evaluated the interaction
	EvaluatorID string `json:"evaluatorId,omitempty"`

	// Specifies when an evaluation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EventTime strfmt.DateTime `json:"eventTime,omitempty"`

	// ID of the evaluation form used
	FormID string `json:"formId,omitempty"`

	// Name of the evaluation form used
	FormName string `json:"formName,omitempty"`

	// o total critical score
	OTotalCriticalScore int64 `json:"oTotalCriticalScore,omitempty"`

	// o total score
	OTotalScore int64 `json:"oTotalScore,omitempty"`

	// The ID of the associated queue
	QueueID string `json:"queueId,omitempty"`

	// Whether the evaluation has been released
	Released bool `json:"released"`

	// Whether the evaluation has been rescored at least once
	Rescored bool `json:"rescored"`

	// ID of the agent the evaluation was performed against
	UserID string `json:"userId,omitempty"`
}

// Validate validates this analytics evaluation
func (m *AnalyticsEvaluation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsEvaluation) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this analytics evaluation based on context it is used
func (m *AnalyticsEvaluation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsEvaluation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsEvaluation) UnmarshalBinary(b []byte) error {
	var res AnalyticsEvaluation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
