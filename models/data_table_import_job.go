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

// DataTableImportJob State information for an import job of rows to a datatable
//
// swagger:model DataTableImportJob
type DataTableImportJob struct {

	// The current count of the number of records deleted
	CountRecordsDeleted int32 `json:"countRecordsDeleted,omitempty"`

	// The current count of the number of records that failed to import
	CountRecordsFailed int32 `json:"countRecordsFailed,omitempty"`

	// The current count of the number of records processed
	CountRecordsUpdated int32 `json:"countRecordsUpdated,omitempty"`

	// The timestamp of when the import stopped (either successfully or unsuccessfully). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// The timestamp of when the import began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Any error information, or null of the processing is not in an error state
	ErrorInformation *ErrorBody `json:"errorInformation,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The indication of whether the processing should remove rows that don't appear in the import file
	// Enum: [ReplaceAll Append]
	ImportMode string `json:"importMode,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The PureCloud user who started the import job
	Owner *AddressableEntityRef `json:"owner,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the import job
	// Required: true
	// Enum: [WaitingForUpload Processing Failed Succeeded]
	Status *string `json:"status"`

	// The URL of the location at which the caller can upload the file to be imported
	// Format: uri
	UploadURI strfmt.URI `json:"uploadURI,omitempty"`
}

// Validate validates this data table import job
func (m *DataTableImportJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTableImportJob) validateDateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCompleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCompleted", "body", "date-time", m.DateCompleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataTableImportJob) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataTableImportJob) validateErrorInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorInformation) { // not required
		return nil
	}

	if m.ErrorInformation != nil {
		if err := m.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInformation")
			}
			return err
		}
	}

	return nil
}

var dataTableImportJobTypeImportModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ReplaceAll","Append"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataTableImportJobTypeImportModePropEnum = append(dataTableImportJobTypeImportModePropEnum, v)
	}
}

const (

	// DataTableImportJobImportModeReplaceAll captures enum value "ReplaceAll"
	DataTableImportJobImportModeReplaceAll string = "ReplaceAll"

	// DataTableImportJobImportModeAppend captures enum value "Append"
	DataTableImportJobImportModeAppend string = "Append"
)

// prop value enum
func (m *DataTableImportJob) validateImportModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataTableImportJobTypeImportModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataTableImportJob) validateImportMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateImportModeEnum("importMode", "body", m.ImportMode); err != nil {
		return err
	}

	return nil
}

func (m *DataTableImportJob) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *DataTableImportJob) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var dataTableImportJobTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WaitingForUpload","Processing","Failed","Succeeded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataTableImportJobTypeStatusPropEnum = append(dataTableImportJobTypeStatusPropEnum, v)
	}
}

const (

	// DataTableImportJobStatusWaitingForUpload captures enum value "WaitingForUpload"
	DataTableImportJobStatusWaitingForUpload string = "WaitingForUpload"

	// DataTableImportJobStatusProcessing captures enum value "Processing"
	DataTableImportJobStatusProcessing string = "Processing"

	// DataTableImportJobStatusFailed captures enum value "Failed"
	DataTableImportJobStatusFailed string = "Failed"

	// DataTableImportJobStatusSucceeded captures enum value "Succeeded"
	DataTableImportJobStatusSucceeded string = "Succeeded"
)

// prop value enum
func (m *DataTableImportJob) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataTableImportJobTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataTableImportJob) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DataTableImportJob) validateUploadURI(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadURI) { // not required
		return nil
	}

	if err := validate.FormatOf("uploadURI", "body", "uri", m.UploadURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this data table import job based on the context it is used
func (m *DataTableImportJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTableImportJob) contextValidateErrorInformation(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorInformation != nil {
		if err := m.ErrorInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInformation")
			}
			return err
		}
	}

	return nil
}

func (m *DataTableImportJob) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DataTableImportJob) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *DataTableImportJob) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTableImportJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTableImportJob) UnmarshalBinary(b []byte) error {
	var res DataTableImportJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
