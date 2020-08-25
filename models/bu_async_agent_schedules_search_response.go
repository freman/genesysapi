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

// BuAsyncAgentSchedulesSearchResponse bu async agent schedules search response
//
// swagger:model BuAsyncAgentSchedulesSearchResponse
type BuAsyncAgentSchedulesSearchResponse struct {

	// The URL from which to download the result if it is too large to pass directly
	DownloadURL string `json:"downloadUrl,omitempty"`

	// The ID for the operation
	OperationID string `json:"operationId,omitempty"`

	// Percent progress for the operation
	Progress int32 `json:"progress,omitempty"`

	// The result of the operation.  Null unless status == Complete
	Result *BuAgentSchedulesSearchResponse `json:"result,omitempty"`

	// The status of the operation
	// Enum: [Processing Complete Canceled Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this bu async agent schedules search response
func (m *BuAsyncAgentSchedulesSearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAsyncAgentSchedulesSearchResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

var buAsyncAgentSchedulesSearchResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","Complete","Canceled","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buAsyncAgentSchedulesSearchResponseTypeStatusPropEnum = append(buAsyncAgentSchedulesSearchResponseTypeStatusPropEnum, v)
	}
}

const (

	// BuAsyncAgentSchedulesSearchResponseStatusProcessing captures enum value "Processing"
	BuAsyncAgentSchedulesSearchResponseStatusProcessing string = "Processing"

	// BuAsyncAgentSchedulesSearchResponseStatusComplete captures enum value "Complete"
	BuAsyncAgentSchedulesSearchResponseStatusComplete string = "Complete"

	// BuAsyncAgentSchedulesSearchResponseStatusCanceled captures enum value "Canceled"
	BuAsyncAgentSchedulesSearchResponseStatusCanceled string = "Canceled"

	// BuAsyncAgentSchedulesSearchResponseStatusError captures enum value "Error"
	BuAsyncAgentSchedulesSearchResponseStatusError string = "Error"
)

// prop value enum
func (m *BuAsyncAgentSchedulesSearchResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buAsyncAgentSchedulesSearchResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuAsyncAgentSchedulesSearchResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAsyncAgentSchedulesSearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAsyncAgentSchedulesSearchResponse) UnmarshalBinary(b []byte) error {
	var res BuAsyncAgentSchedulesSearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}