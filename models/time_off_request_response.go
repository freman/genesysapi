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

// TimeOffRequestResponse time off request response
//
// swagger:model TimeOffRequestResponse
type TimeOffRequestResponse struct {

	// The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
	ActivityCodeID string `json:"activityCodeId,omitempty"`

	// The daily duration of this time off request in minutes
	DailyDurationMinutes int32 `json:"dailyDurationMinutes,omitempty"`

	// A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.  Will be not empty if isFullDayRequest == true
	// Unique: true
	FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Whether this is a full day request (false means partial day)
	IsFullDayRequest bool `json:"isFullDayRequest"`

	// Whether this request has been marked as read by the agent
	MarkedAsRead bool `json:"markedAsRead"`

	// The version metadata of the time off request
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// The user who last modified this TimeOffRequestResponse
	ModifiedBy *UserReference `json:"modifiedBy,omitempty"`

	// The timestamp when this request was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// Notes about the time off request
	Notes string `json:"notes,omitempty"`

	// A set of start date-times in ISO-8601 format for partial day requests.  Will be not empty if isFullDayRequest == false
	// Unique: true
	PartialDayStartDateTimes []strfmt.DateTime `json:"partialDayStartDateTimes"`

	// The user who reviewed this time off request
	ReviewedBy *UserReference `json:"reviewedBy,omitempty"`

	// The timestamp when this request was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReviewedDate strfmt.DateTime `json:"reviewedDate,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of this time off request
	// Enum: [PENDING APPROVED DENIED CANCELED]
	Status string `json:"status,omitempty"`

	// The user who submitted this time off request
	SubmittedBy *UserReference `json:"submittedBy,omitempty"`

	// The timestamp when this request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	SubmittedDate strfmt.DateTime `json:"submittedDate,omitempty"`

	// The user associated with this time off request
	User *UserReference `json:"user,omitempty"`
}

// Validate validates this time off request response
func (m *TimeOffRequestResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullDayManagementUnitDates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartialDayStartDateTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeOffRequestResponse) validateFullDayManagementUnitDates(formats strfmt.Registry) error {

	if swag.IsZero(m.FullDayManagementUnitDates) { // not required
		return nil
	}

	if err := validate.UniqueItems("fullDayManagementUnitDates", "body", m.FullDayManagementUnitDates); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequestResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *TimeOffRequestResponse) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TimeOffRequestResponse) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequestResponse) validatePartialDayStartDateTimes(formats strfmt.Registry) error {

	if swag.IsZero(m.PartialDayStartDateTimes) { // not required
		return nil
	}

	if err := validate.UniqueItems("partialDayStartDateTimes", "body", m.PartialDayStartDateTimes); err != nil {
		return err
	}

	for i := 0; i < len(m.PartialDayStartDateTimes); i++ {

		if err := validate.FormatOf("partialDayStartDateTimes"+"."+strconv.Itoa(i), "body", "date-time", m.PartialDayStartDateTimes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *TimeOffRequestResponse) validateReviewedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ReviewedBy) { // not required
		return nil
	}

	if m.ReviewedBy != nil {
		if err := m.ReviewedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reviewedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TimeOffRequestResponse) validateReviewedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReviewedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("reviewedDate", "body", "date-time", m.ReviewedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequestResponse) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var timeOffRequestResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","APPROVED","DENIED","CANCELED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeOffRequestResponseTypeStatusPropEnum = append(timeOffRequestResponseTypeStatusPropEnum, v)
	}
}

const (

	// TimeOffRequestResponseStatusPENDING captures enum value "PENDING"
	TimeOffRequestResponseStatusPENDING string = "PENDING"

	// TimeOffRequestResponseStatusAPPROVED captures enum value "APPROVED"
	TimeOffRequestResponseStatusAPPROVED string = "APPROVED"

	// TimeOffRequestResponseStatusDENIED captures enum value "DENIED"
	TimeOffRequestResponseStatusDENIED string = "DENIED"

	// TimeOffRequestResponseStatusCANCELED captures enum value "CANCELED"
	TimeOffRequestResponseStatusCANCELED string = "CANCELED"
)

// prop value enum
func (m *TimeOffRequestResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, timeOffRequestResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TimeOffRequestResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequestResponse) validateSubmittedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedBy) { // not required
		return nil
	}

	if m.SubmittedBy != nil {
		if err := m.SubmittedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("submittedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TimeOffRequestResponse) validateSubmittedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("submittedDate", "body", "date-time", m.SubmittedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffRequestResponse) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeOffRequestResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOffRequestResponse) UnmarshalBinary(b []byte) error {
	var res TimeOffRequestResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
