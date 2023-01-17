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

// HistoricalImportStatus historical import status
//
// swagger:model HistoricalImportStatus
type HistoricalImportStatus struct {

	// Whether this historical import is active or not
	// Read Only: true
	Active *bool `json:"active"`

	// Date in which the historical import is initiated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The last day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateImportEnded strfmt.DateTime `json:"dateImportEnded,omitempty"`

	// The first day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateImportStarted strfmt.DateTime `json:"dateImportStarted,omitempty"`

	// Date in which the historical import is modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Error occured if the status of the import is failed
	// Read Only: true
	Error string `json:"error,omitempty"`

	// Request id of the historical import in the organization
	// Read Only: true
	RequestID string `json:"requestId,omitempty"`

	// Status of the historical import in the organization.
	// Read Only: true
	// Enum: [Initiated InProgress Pending Success Failed Cancelled Purged PurgePending]
	Status string `json:"status,omitempty"`

	// Whether this historical import is of type csv or json
	// Read Only: true
	// Enum: [Csv Json]
	Type string `json:"type,omitempty"`
}

// Validate validates this historical import status
func (m *HistoricalImportStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateImportEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateImportStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
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

func (m *HistoricalImportStatus) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) validateDateImportEnded(formats strfmt.Registry) error {
	if swag.IsZero(m.DateImportEnded) { // not required
		return nil
	}

	if err := validate.FormatOf("dateImportEnded", "body", "date-time", m.DateImportEnded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) validateDateImportStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.DateImportStarted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateImportStarted", "body", "date-time", m.DateImportStarted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var historicalImportStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Initiated","InProgress","Pending","Success","Failed","Cancelled","Purged","PurgePending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historicalImportStatusTypeStatusPropEnum = append(historicalImportStatusTypeStatusPropEnum, v)
	}
}

const (

	// HistoricalImportStatusStatusInitiated captures enum value "Initiated"
	HistoricalImportStatusStatusInitiated string = "Initiated"

	// HistoricalImportStatusStatusInProgress captures enum value "InProgress"
	HistoricalImportStatusStatusInProgress string = "InProgress"

	// HistoricalImportStatusStatusPending captures enum value "Pending"
	HistoricalImportStatusStatusPending string = "Pending"

	// HistoricalImportStatusStatusSuccess captures enum value "Success"
	HistoricalImportStatusStatusSuccess string = "Success"

	// HistoricalImportStatusStatusFailed captures enum value "Failed"
	HistoricalImportStatusStatusFailed string = "Failed"

	// HistoricalImportStatusStatusCancelled captures enum value "Cancelled"
	HistoricalImportStatusStatusCancelled string = "Cancelled"

	// HistoricalImportStatusStatusPurged captures enum value "Purged"
	HistoricalImportStatusStatusPurged string = "Purged"

	// HistoricalImportStatusStatusPurgePending captures enum value "PurgePending"
	HistoricalImportStatusStatusPurgePending string = "PurgePending"
)

// prop value enum
func (m *HistoricalImportStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historicalImportStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoricalImportStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var historicalImportStatusTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Csv","Json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historicalImportStatusTypeTypePropEnum = append(historicalImportStatusTypeTypePropEnum, v)
	}
}

const (

	// HistoricalImportStatusTypeCsv captures enum value "Csv"
	HistoricalImportStatusTypeCsv string = "Csv"

	// HistoricalImportStatusTypeJSON captures enum value "Json"
	HistoricalImportStatusTypeJSON string = "Json"
)

// prop value enum
func (m *HistoricalImportStatus) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historicalImportStatusTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoricalImportStatus) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this historical import status based on the context it is used
func (m *HistoricalImportStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateImportEnded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateImportStarted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoricalImportStatus) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateDateImportEnded(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateImportEnded", "body", strfmt.DateTime(m.DateImportEnded)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateDateImportStarted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateImportStarted", "body", strfmt.DateTime(m.DateImportStarted)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "error", "body", string(m.Error)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateRequestID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "requestId", "body", string(m.RequestID)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalImportStatus) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoricalImportStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoricalImportStatus) UnmarshalBinary(b []byte) error {
	var res HistoricalImportStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
