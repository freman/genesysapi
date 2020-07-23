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

// FlowRuntimeExecution Details about the current state of a Flow execution
//
// swagger:model FlowRuntimeExecution
type FlowRuntimeExecution struct {

	// The completion reason set at the flow completion time, if applicable.
	CompletionReason string `json:"completionReason,omitempty"`

	// The conversation to which this Flow execution is related
	Conversation *DomainEntityRef `json:"conversation,omitempty"`

	// The time the flow completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// The time the flow was launched. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Required: true
	// Format: date-time
	DateLaunched *strfmt.DateTime `json:"dateLaunched"`

	// Additional information if the flow is in error
	FlowErrorInfo *ErrorBody `json:"flowErrorInfo,omitempty"`

	// The Version of the flow definition of the flow execution.
	// Required: true
	FlowVersion *FlowVersion `json:"flowVersion"`

	// The flow execution ID
	ID string `json:"id,omitempty"`

	// The flow execution name.
	Name string `json:"name,omitempty"`

	// List of the flow's output variables, if any. Output variables are only supplied for Completed flows.
	OutputData map[string]interface{} `json:"outputData,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The flow's running status, which indicates whether the flow is running normally or completed, etc.
	// Required: true
	// Enum: [UNKNOWN RUNNING ERROR TERMINATED COMPLETED FAILED]
	Status *string `json:"status"`
}

// Validate validates this flow runtime execution
func (m *FlowRuntimeExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateLaunched(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowErrorInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *FlowRuntimeExecution) validateConversation(formats strfmt.Registry) error {

	if swag.IsZero(m.Conversation) { // not required
		return nil
	}

	if m.Conversation != nil {
		if err := m.Conversation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *FlowRuntimeExecution) validateDateCompleted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCompleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCompleted", "body", "date-time", m.DateCompleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FlowRuntimeExecution) validateDateLaunched(formats strfmt.Registry) error {

	if err := validate.Required("dateLaunched", "body", m.DateLaunched); err != nil {
		return err
	}

	if err := validate.FormatOf("dateLaunched", "body", "date-time", m.DateLaunched.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FlowRuntimeExecution) validateFlowErrorInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowErrorInfo) { // not required
		return nil
	}

	if m.FlowErrorInfo != nil {
		if err := m.FlowErrorInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowErrorInfo")
			}
			return err
		}
	}

	return nil
}

func (m *FlowRuntimeExecution) validateFlowVersion(formats strfmt.Registry) error {

	if err := validate.Required("flowVersion", "body", m.FlowVersion); err != nil {
		return err
	}

	if m.FlowVersion != nil {
		if err := m.FlowVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowVersion")
			}
			return err
		}
	}

	return nil
}

func (m *FlowRuntimeExecution) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var flowRuntimeExecutionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","RUNNING","ERROR","TERMINATED","COMPLETED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowRuntimeExecutionTypeStatusPropEnum = append(flowRuntimeExecutionTypeStatusPropEnum, v)
	}
}

const (

	// FlowRuntimeExecutionStatusUNKNOWN captures enum value "UNKNOWN"
	FlowRuntimeExecutionStatusUNKNOWN string = "UNKNOWN"

	// FlowRuntimeExecutionStatusRUNNING captures enum value "RUNNING"
	FlowRuntimeExecutionStatusRUNNING string = "RUNNING"

	// FlowRuntimeExecutionStatusERROR captures enum value "ERROR"
	FlowRuntimeExecutionStatusERROR string = "ERROR"

	// FlowRuntimeExecutionStatusTERMINATED captures enum value "TERMINATED"
	FlowRuntimeExecutionStatusTERMINATED string = "TERMINATED"

	// FlowRuntimeExecutionStatusCOMPLETED captures enum value "COMPLETED"
	FlowRuntimeExecutionStatusCOMPLETED string = "COMPLETED"

	// FlowRuntimeExecutionStatusFAILED captures enum value "FAILED"
	FlowRuntimeExecutionStatusFAILED string = "FAILED"
)

// prop value enum
func (m *FlowRuntimeExecution) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowRuntimeExecutionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FlowRuntimeExecution) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowRuntimeExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowRuntimeExecution) UnmarshalBinary(b []byte) error {
	var res FlowRuntimeExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
