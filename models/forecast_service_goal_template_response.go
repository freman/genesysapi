// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ForecastServiceGoalTemplateResponse forecast service goal template response
//
// swagger:model ForecastServiceGoalTemplateResponse
type ForecastServiceGoalTemplateResponse struct {

	// The abandon rate goal for this forecast
	AbandonRate *ForecastAbandonRateResponse `json:"abandonRate,omitempty"`

	// The average speed of answer goal for this forecast
	AverageSpeedOfAnswer *ForecastAverageSpeedOfAnswerResponse `json:"averageSpeedOfAnswer,omitempty"`

	// The service level goal for this forecast
	ServiceLevel *ForecastServiceLevelResponse `json:"serviceLevel,omitempty"`
}

// Validate validates this forecast service goal template response
func (m *ForecastServiceGoalTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbandonRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAverageSpeedOfAnswer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ForecastServiceGoalTemplateResponse) validateAbandonRate(formats strfmt.Registry) error {
	if swag.IsZero(m.AbandonRate) { // not required
		return nil
	}

	if m.AbandonRate != nil {
		if err := m.AbandonRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abandonRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abandonRate")
			}
			return err
		}
	}

	return nil
}

func (m *ForecastServiceGoalTemplateResponse) validateAverageSpeedOfAnswer(formats strfmt.Registry) error {
	if swag.IsZero(m.AverageSpeedOfAnswer) { // not required
		return nil
	}

	if m.AverageSpeedOfAnswer != nil {
		if err := m.AverageSpeedOfAnswer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("averageSpeedOfAnswer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("averageSpeedOfAnswer")
			}
			return err
		}
	}

	return nil
}

func (m *ForecastServiceGoalTemplateResponse) validateServiceLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceLevel) { // not required
		return nil
	}

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this forecast service goal template response based on the context it is used
func (m *ForecastServiceGoalTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAbandonRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAverageSpeedOfAnswer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ForecastServiceGoalTemplateResponse) contextValidateAbandonRate(ctx context.Context, formats strfmt.Registry) error {

	if m.AbandonRate != nil {
		if err := m.AbandonRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abandonRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abandonRate")
			}
			return err
		}
	}

	return nil
}

func (m *ForecastServiceGoalTemplateResponse) contextValidateAverageSpeedOfAnswer(ctx context.Context, formats strfmt.Registry) error {

	if m.AverageSpeedOfAnswer != nil {
		if err := m.AverageSpeedOfAnswer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("averageSpeedOfAnswer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("averageSpeedOfAnswer")
			}
			return err
		}
	}

	return nil
}

func (m *ForecastServiceGoalTemplateResponse) contextValidateServiceLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ForecastServiceGoalTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForecastServiceGoalTemplateResponse) UnmarshalBinary(b []byte) error {
	var res ForecastServiceGoalTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
