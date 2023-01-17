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

// ExpandableWebDeployment Details about a Web Deployment
//
// swagger:model ExpandableWebDeployment
type ExpandableWebDeployment struct {

	// Property indicates whether all domains are allowed or not. allowedDomains must be empty when this is set as true.
	AllowAllDomains bool `json:"allowAllDomains"`

	// The list of domains that are approved to use this deployment; the list will be added to CORS headers for ease of web use.
	AllowedDomains []string `json:"allowedDomains"`

	// The config version this deployment uses
	// Required: true
	Configuration *WebDeploymentConfigurationVersionResponse `json:"configuration"`

	// The date the deployment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date the deployment was most recently modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The description of the config
	Description string `json:"description,omitempty"`

	// A reference to the inboundshortmessage flow used by this deployment
	Flow *DomainEntityRef `json:"flow,omitempty"`

	// The deployment ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// A reference to the user who most recently modified the deployment
	// Read Only: true
	LastModifiedUser *AddressableEntityRef `json:"lastModifiedUser,omitempty"`

	// The deployment name
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Javascript snippet used to load the config
	// Read Only: true
	Snippet string `json:"snippet,omitempty"`

	// The current status of the deployment
	// Enum: [Pending Active Inactive Error Deleting]
	Status string `json:"status,omitempty"`
}

// Validate validates this expandable web deployment
func (m *ExpandableWebDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ExpandableWebDeployment) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandableWebDeployment) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) validateFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.Flow) { // not required
		return nil
	}

	if m.Flow != nil {
		if err := m.Flow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandableWebDeployment) validateLastModifiedUser(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedUser) { // not required
		return nil
	}

	if m.LastModifiedUser != nil {
		if err := m.LastModifiedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastModifiedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastModifiedUser")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandableWebDeployment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var expandableWebDeploymentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Active","Inactive","Error","Deleting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		expandableWebDeploymentTypeStatusPropEnum = append(expandableWebDeploymentTypeStatusPropEnum, v)
	}
}

const (

	// ExpandableWebDeploymentStatusPending captures enum value "Pending"
	ExpandableWebDeploymentStatusPending string = "Pending"

	// ExpandableWebDeploymentStatusActive captures enum value "Active"
	ExpandableWebDeploymentStatusActive string = "Active"

	// ExpandableWebDeploymentStatusInactive captures enum value "Inactive"
	ExpandableWebDeploymentStatusInactive string = "Inactive"

	// ExpandableWebDeploymentStatusError captures enum value "Error"
	ExpandableWebDeploymentStatusError string = "Error"

	// ExpandableWebDeploymentStatusDeleting captures enum value "Deleting"
	ExpandableWebDeploymentStatusDeleting string = "Deleting"
)

// prop value enum
func (m *ExpandableWebDeployment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, expandableWebDeploymentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExpandableWebDeployment) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this expandable web deployment based on the context it is used
func (m *ExpandableWebDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastModifiedUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnippet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpandableWebDeployment) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.Configuration != nil {
		if err := m.Configuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.Flow != nil {
		if err := m.Flow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateLastModifiedUser(ctx context.Context, formats strfmt.Registry) error {

	if m.LastModifiedUser != nil {
		if err := m.LastModifiedUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastModifiedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastModifiedUser")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *ExpandableWebDeployment) contextValidateSnippet(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snippet", "body", string(m.Snippet)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpandableWebDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpandableWebDeployment) UnmarshalBinary(b []byte) error {
	var res ExpandableWebDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}