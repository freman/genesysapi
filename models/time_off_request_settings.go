// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeOffRequestSettings Time Off Request Settings
//
// swagger:model TimeOffRequestSettings
type TimeOffRequestSettings struct {

	// The earliest number of days from now for which an agent can submit a time off request.  Use negative numbers to indicate days in the past
	SubmissionEarliestDaysFromNow int32 `json:"submissionEarliestDaysFromNow,omitempty"`

	// The latest number of days from now for which an agent can submit a time off request
	SubmissionLatestDaysFromNow int32 `json:"submissionLatestDaysFromNow,omitempty"`

	// Whether to enforce a submission range for agent time off requests
	SubmissionRangeEnforced bool `json:"submissionRangeEnforced,omitempty"`
}

// Validate validates this time off request settings
func (m *TimeOffRequestSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeOffRequestSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOffRequestSettings) UnmarshalBinary(b []byte) error {
	var res TimeOffRequestSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
