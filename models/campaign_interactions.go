// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CampaignInteractions campaign interactions
//
// swagger:model CampaignInteractions
type CampaignInteractions struct {

	// campaign
	Campaign *DomainEntityRef `json:"campaign,omitempty"`

	// interacting interactions
	InteractingInteractions []*CampaignInteraction `json:"interactingInteractions"`

	// pending interactions
	PendingInteractions []*CampaignInteraction `json:"pendingInteractions"`

	// previewing interactions
	PreviewingInteractions []*CampaignInteraction `json:"previewingInteractions"`

	// proceeding interactions
	ProceedingInteractions []*CampaignInteraction `json:"proceedingInteractions"`

	// scheduled interactions
	ScheduledInteractions []*CampaignInteraction `json:"scheduledInteractions"`
}

// Validate validates this campaign interactions
func (m *CampaignInteractions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInteractingInteractions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingInteractions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviewingInteractions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProceedingInteractions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledInteractions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignInteractions) validateCampaign(formats strfmt.Registry) error {

	if swag.IsZero(m.Campaign) { // not required
		return nil
	}

	if m.Campaign != nil {
		if err := m.Campaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("campaign")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteractions) validateInteractingInteractions(formats strfmt.Registry) error {

	if swag.IsZero(m.InteractingInteractions) { // not required
		return nil
	}

	for i := 0; i < len(m.InteractingInteractions); i++ {
		if swag.IsZero(m.InteractingInteractions[i]) { // not required
			continue
		}

		if m.InteractingInteractions[i] != nil {
			if err := m.InteractingInteractions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interactingInteractions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CampaignInteractions) validatePendingInteractions(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingInteractions) { // not required
		return nil
	}

	for i := 0; i < len(m.PendingInteractions); i++ {
		if swag.IsZero(m.PendingInteractions[i]) { // not required
			continue
		}

		if m.PendingInteractions[i] != nil {
			if err := m.PendingInteractions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pendingInteractions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CampaignInteractions) validatePreviewingInteractions(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviewingInteractions) { // not required
		return nil
	}

	for i := 0; i < len(m.PreviewingInteractions); i++ {
		if swag.IsZero(m.PreviewingInteractions[i]) { // not required
			continue
		}

		if m.PreviewingInteractions[i] != nil {
			if err := m.PreviewingInteractions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("previewingInteractions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CampaignInteractions) validateProceedingInteractions(formats strfmt.Registry) error {

	if swag.IsZero(m.ProceedingInteractions) { // not required
		return nil
	}

	for i := 0; i < len(m.ProceedingInteractions); i++ {
		if swag.IsZero(m.ProceedingInteractions[i]) { // not required
			continue
		}

		if m.ProceedingInteractions[i] != nil {
			if err := m.ProceedingInteractions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proceedingInteractions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CampaignInteractions) validateScheduledInteractions(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledInteractions) { // not required
		return nil
	}

	for i := 0; i < len(m.ScheduledInteractions); i++ {
		if swag.IsZero(m.ScheduledInteractions[i]) { // not required
			continue
		}

		if m.ScheduledInteractions[i] != nil {
			if err := m.ScheduledInteractions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scheduledInteractions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignInteractions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignInteractions) UnmarshalBinary(b []byte) error {
	var res CampaignInteractions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
