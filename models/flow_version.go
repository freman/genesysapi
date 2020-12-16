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

// FlowVersion flow version
//
// swagger:model FlowVersion
type FlowVersion struct {

	// commit version
	CommitVersion string `json:"commitVersion,omitempty"`

	// configuration Uri
	ConfigurationURI string `json:"configurationUri,omitempty"`

	// configuration version
	ConfigurationVersion string `json:"configurationVersion,omitempty"`

	// created by
	CreatedBy *User `json:"createdBy,omitempty"`

	// created by client
	CreatedByClient *DomainEntityRef `json:"createdByClient,omitempty"`

	// date created
	DateCreated int64 `json:"dateCreated,omitempty"`

	// debug
	Debug bool `json:"debug"`

	// generation Id
	GenerationID string `json:"generationId,omitempty"`

	// The flow version identifier
	ID string `json:"id,omitempty"`

	// input schema
	InputSchema *JSONSchemaDocument `json:"inputSchema,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Information about the natural language understanding configuration for the flow version
	// Read Only: true
	NluInfo *NluInfo `json:"nluInfo,omitempty"`

	// output schema
	OutputSchema *JSONSchemaDocument `json:"outputSchema,omitempty"`

	// publish result Uri
	// Format: uri
	PublishResultURI strfmt.URI `json:"publishResultUri,omitempty"`

	// secure
	Secure bool `json:"secure"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// List of supported languages for this version of the flow
	// Read Only: true
	SupportedLanguages []*SupportedLanguage `json:"supportedLanguages"`

	// type
	// Enum: [PUBLISH CHECKIN SAVE]
	Type string `json:"type,omitempty"`
}

// Validate validates this flow version
func (m *FlowVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedByClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNluInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishResultURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedLanguages(formats); err != nil {
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

func (m *FlowVersion) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *FlowVersion) validateCreatedByClient(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedByClient) { // not required
		return nil
	}

	if m.CreatedByClient != nil {
		if err := m.CreatedByClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdByClient")
			}
			return err
		}
	}

	return nil
}

func (m *FlowVersion) validateInputSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.InputSchema) { // not required
		return nil
	}

	if m.InputSchema != nil {
		if err := m.InputSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inputSchema")
			}
			return err
		}
	}

	return nil
}

func (m *FlowVersion) validateNluInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.NluInfo) { // not required
		return nil
	}

	if m.NluInfo != nil {
		if err := m.NluInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nluInfo")
			}
			return err
		}
	}

	return nil
}

func (m *FlowVersion) validateOutputSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputSchema) { // not required
		return nil
	}

	if m.OutputSchema != nil {
		if err := m.OutputSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outputSchema")
			}
			return err
		}
	}

	return nil
}

func (m *FlowVersion) validatePublishResultURI(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishResultURI) { // not required
		return nil
	}

	if err := validate.FormatOf("publishResultUri", "body", "uri", m.PublishResultURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FlowVersion) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FlowVersion) validateSupportedLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedLanguages) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedLanguages); i++ {
		if swag.IsZero(m.SupportedLanguages[i]) { // not required
			continue
		}

		if m.SupportedLanguages[i] != nil {
			if err := m.SupportedLanguages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportedLanguages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var flowVersionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLISH","CHECKIN","SAVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowVersionTypeTypePropEnum = append(flowVersionTypeTypePropEnum, v)
	}
}

const (

	// FlowVersionTypePUBLISH captures enum value "PUBLISH"
	FlowVersionTypePUBLISH string = "PUBLISH"

	// FlowVersionTypeCHECKIN captures enum value "CHECKIN"
	FlowVersionTypeCHECKIN string = "CHECKIN"

	// FlowVersionTypeSAVE captures enum value "SAVE"
	FlowVersionTypeSAVE string = "SAVE"
)

// prop value enum
func (m *FlowVersion) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowVersionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FlowVersion) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowVersion) UnmarshalBinary(b []byte) error {
	var res FlowVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
