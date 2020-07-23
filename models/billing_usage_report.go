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

// BillingUsageReport billing usage report
//
// swagger:model BillingUsageReport
type BillingUsageReport struct {

	// The period end date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The period start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate"`

	// Generation status of report
	// Enum: [InProgress Complete]
	Status string `json:"status,omitempty"`

	// The usages for the given period.
	// Required: true
	Usages []*BillingUsage `json:"usages"`
}

// Validate validates this billing usage report
func (m *BillingUsageReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingUsageReport) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingUsageReport) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingUsageReport) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var billingUsageReportTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InProgress","Complete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingUsageReportTypeStatusPropEnum = append(billingUsageReportTypeStatusPropEnum, v)
	}
}

const (

	// BillingUsageReportStatusInProgress captures enum value "InProgress"
	BillingUsageReportStatusInProgress string = "InProgress"

	// BillingUsageReportStatusComplete captures enum value "Complete"
	BillingUsageReportStatusComplete string = "Complete"
)

// prop value enum
func (m *BillingUsageReport) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billingUsageReportTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillingUsageReport) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BillingUsageReport) validateUsages(formats strfmt.Registry) error {

	if err := validate.Required("usages", "body", m.Usages); err != nil {
		return err
	}

	for i := 0; i < len(m.Usages); i++ {
		if swag.IsZero(m.Usages[i]) { // not required
			continue
		}

		if m.Usages[i] != nil {
			if err := m.Usages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingUsageReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingUsageReport) UnmarshalBinary(b []byte) error {
	var res BillingUsageReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
