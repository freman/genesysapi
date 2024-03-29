// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Campaign campaign
//
// swagger:model Campaign
type Campaign struct {

	// The targeted abandon rate percentage. Required for progressive, power, and predictive campaigns.
	AbandonRate float64 `json:"abandonRate,omitempty"`

	// Indicates (when true) that the campaign will remain on after contacts are depleted, allowing additional contacts to be appended/added to the contact list and processed by the still-running campaign. The campaign can still be turned off manually.
	AlwaysRunning bool `json:"alwaysRunning"`

	// The language the edge will use to analyze the call.
	CallAnalysisLanguage string `json:"callAnalysisLanguage,omitempty"`

	// The call analysis response set to handle call analysis results from the edge. Required for all dialing modes except preview.
	CallAnalysisResponseSet *DomainEntityRef `json:"callAnalysisResponseSet,omitempty"`

	// The callable time set for this campaign to check before placing a call.
	CallableTimeSet *DomainEntityRef `json:"callableTimeSet,omitempty"`

	// The caller id phone number to be displayed on the outbound call.
	// Example: (555) 555-5555
	// Required: true
	CallerAddress *string `json:"callerAddress"`

	// The caller id name to be displayed on the outbound call.
	// Required: true
	CallerName *string `json:"callerName"`

	// The current status of the Campaign. A Campaign may be turned 'on' or 'off'. Required for updates.
	// Enum: [on stopping off complete invalid forced_off forced_stopping]
	CampaignStatus string `json:"campaignStatus,omitempty"`

	// The ContactList for this Campaign to dial.
	// Required: true
	ContactList *DomainEntityRef `json:"contactList"`

	// Filter to apply to the contact list before dialing. Currently a campaign can only have one filter applied.
	ContactListFilters []*DomainEntityRef `json:"contactListFilters"`

	// The order in which to sort contacts for dialing, based on a column.
	ContactSort *ContactSort `json:"contactSort,omitempty"`

	// The order in which to sort contacts for dialing, based on up to four columns.
	ContactSorts []*ContactSort `json:"contactSorts"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The strategy this Campaign will use for dialing.
	// Required: true
	// Enum: [agentless preview power predictive progressive external]
	DialingMode *string `json:"dialingMode"`

	// The division this campaign belongs to.
	Division *DomainEntityRef `json:"division,omitempty"`

	// DncLists for this Campaign to check before placing a call.
	DncLists []*DomainEntityRef `json:"dncLists"`

	// Settings for dynamic queueing of contacts.
	DynamicContactQueueingSettings *DynamicContactQueueingSettings `json:"dynamicContactQueueingSettings,omitempty"`

	// The EdgeGroup that will place the calls. Required for all dialing modes except preview.
	EdgeGroup *DomainEntityRef `json:"edgeGroup,omitempty"`

	// A list of current error conditions associated with the campaign.
	// Read Only: true
	Errors []*RestErrorDetail `json:"errors"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the Campaign.
	// Required: true
	Name *string `json:"name"`

	// How long to wait before dispositioning a call as 'no-answer'. Default 30 seconds. Only applicable to non-preview campaigns.
	NoAnswerTimeout int32 `json:"noAnswerTimeout,omitempty"`

	// The number of outbound lines to be concurrently dialed. Only applicable to non-preview campaigns; only required for agentless.
	OutboundLineCount int32 `json:"outboundLineCount,omitempty"`

	// The ContactPhoneNumberColumns on the ContactList that this Campaign should dial.
	// Required: true
	PhoneColumns []*PhoneColumn `json:"phoneColumns"`

	// The number of seconds before a call will be automatically placed on a preview. A value of 0 indicates no automatic placement of calls. Only applicable to preview campaigns.
	PreviewTimeOutSeconds int64 `json:"previewTimeOutSeconds,omitempty"`

	// The priority of this campaign relative to other campaigns that are running on the same queue. 5 is the highest priority, 1 the lowest.
	Priority int32 `json:"priority,omitempty"`

	// The Queue for this Campaign to route calls to. Required for all dialing modes except agentless.
	Queue *DomainEntityRef `json:"queue,omitempty"`

	// Rule sets to be applied while this campaign is dialing.
	RuleSets []*DomainEntityRef `json:"ruleSets"`

	// The Script to be displayed to agents that are handling outbound calls. Required for all dialing modes except agentless.
	Script *DomainEntityRef `json:"script,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The identifier of the site to be used for dialing; can be set in place of an edge group.
	Site *DomainEntityRef `json:"site,omitempty"`

	// Whether or not agents can skip previews without placing a call. Only applicable for preview campaigns.
	SkipPreviewDisabled bool `json:"skipPreviewDisabled"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this campaign
func (m *Campaign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallAnalysisResponseSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallableTimeSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallerAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaignStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactListFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactSort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactSorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDialingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDncLists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamicContactQueueingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Campaign) validateCallAnalysisResponseSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CallAnalysisResponseSet) { // not required
		return nil
	}

	if m.CallAnalysisResponseSet != nil {
		if err := m.CallAnalysisResponseSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callAnalysisResponseSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callAnalysisResponseSet")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateCallableTimeSet(formats strfmt.Registry) error {
	if swag.IsZero(m.CallableTimeSet) { // not required
		return nil
	}

	if m.CallableTimeSet != nil {
		if err := m.CallableTimeSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callableTimeSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callableTimeSet")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateCallerAddress(formats strfmt.Registry) error {

	if err := validate.Required("callerAddress", "body", m.CallerAddress); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) validateCallerName(formats strfmt.Registry) error {

	if err := validate.Required("callerName", "body", m.CallerName); err != nil {
		return err
	}

	return nil
}

var campaignTypeCampaignStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","stopping","off","complete","invalid","forced_off","forced_stopping"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignTypeCampaignStatusPropEnum = append(campaignTypeCampaignStatusPropEnum, v)
	}
}

const (

	// CampaignCampaignStatusOn captures enum value "on"
	CampaignCampaignStatusOn string = "on"

	// CampaignCampaignStatusStopping captures enum value "stopping"
	CampaignCampaignStatusStopping string = "stopping"

	// CampaignCampaignStatusOff captures enum value "off"
	CampaignCampaignStatusOff string = "off"

	// CampaignCampaignStatusComplete captures enum value "complete"
	CampaignCampaignStatusComplete string = "complete"

	// CampaignCampaignStatusInvalid captures enum value "invalid"
	CampaignCampaignStatusInvalid string = "invalid"

	// CampaignCampaignStatusForcedOff captures enum value "forced_off"
	CampaignCampaignStatusForcedOff string = "forced_off"

	// CampaignCampaignStatusForcedStopping captures enum value "forced_stopping"
	CampaignCampaignStatusForcedStopping string = "forced_stopping"
)

// prop value enum
func (m *Campaign) validateCampaignStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignTypeCampaignStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Campaign) validateCampaignStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.CampaignStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCampaignStatusEnum("campaignStatus", "body", m.CampaignStatus); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) validateContactList(formats strfmt.Registry) error {

	if err := validate.Required("contactList", "body", m.ContactList); err != nil {
		return err
	}

	if m.ContactList != nil {
		if err := m.ContactList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contactList")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateContactListFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactListFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactListFilters); i++ {
		if swag.IsZero(m.ContactListFilters[i]) { // not required
			continue
		}

		if m.ContactListFilters[i] != nil {
			if err := m.ContactListFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactListFilters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contactListFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) validateContactSort(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactSort) { // not required
		return nil
	}

	if m.ContactSort != nil {
		if err := m.ContactSort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactSort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contactSort")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateContactSorts(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactSorts) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactSorts); i++ {
		if swag.IsZero(m.ContactSorts[i]) { // not required
			continue
		}

		if m.ContactSorts[i] != nil {
			if err := m.ContactSorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactSorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contactSorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var campaignTypeDialingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agentless","preview","power","predictive","progressive","external"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignTypeDialingModePropEnum = append(campaignTypeDialingModePropEnum, v)
	}
}

const (

	// CampaignDialingModeAgentless captures enum value "agentless"
	CampaignDialingModeAgentless string = "agentless"

	// CampaignDialingModePreview captures enum value "preview"
	CampaignDialingModePreview string = "preview"

	// CampaignDialingModePower captures enum value "power"
	CampaignDialingModePower string = "power"

	// CampaignDialingModePredictive captures enum value "predictive"
	CampaignDialingModePredictive string = "predictive"

	// CampaignDialingModeProgressive captures enum value "progressive"
	CampaignDialingModeProgressive string = "progressive"

	// CampaignDialingModeExternal captures enum value "external"
	CampaignDialingModeExternal string = "external"
)

// prop value enum
func (m *Campaign) validateDialingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignTypeDialingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Campaign) validateDialingMode(formats strfmt.Registry) error {

	if err := validate.Required("dialingMode", "body", m.DialingMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateDialingModeEnum("dialingMode", "body", *m.DialingMode); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) validateDivision(formats strfmt.Registry) error {
	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateDncLists(formats strfmt.Registry) error {
	if swag.IsZero(m.DncLists) { // not required
		return nil
	}

	for i := 0; i < len(m.DncLists); i++ {
		if swag.IsZero(m.DncLists[i]) { // not required
			continue
		}

		if m.DncLists[i] != nil {
			if err := m.DncLists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dncLists" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dncLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) validateDynamicContactQueueingSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicContactQueueingSettings) { // not required
		return nil
	}

	if m.DynamicContactQueueingSettings != nil {
		if err := m.DynamicContactQueueingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamicContactQueueingSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamicContactQueueingSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateEdgeGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeGroup) { // not required
		return nil
	}

	if m.EdgeGroup != nil {
		if err := m.EdgeGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) validatePhoneColumns(formats strfmt.Registry) error {

	if err := validate.Required("phoneColumns", "body", m.PhoneColumns); err != nil {
		return err
	}

	for i := 0; i < len(m.PhoneColumns); i++ {
		if swag.IsZero(m.PhoneColumns[i]) { // not required
			continue
		}

		if m.PhoneColumns[i] != nil {
			if err := m.PhoneColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneColumns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phoneColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateRuleSets(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleSets) { // not required
		return nil
	}

	for i := 0; i < len(m.RuleSets); i++ {
		if swag.IsZero(m.RuleSets[i]) { // not required
			continue
		}

		if m.RuleSets[i] != nil {
			if err := m.RuleSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleSets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ruleSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) validateScript(formats strfmt.Registry) error {
	if swag.IsZero(m.Script) { // not required
		return nil
	}

	if m.Script != nil {
		if err := m.Script.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this campaign based on the context it is used
func (m *Campaign) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCallAnalysisResponseSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCallableTimeSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContactList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContactListFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContactSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContactSorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDncLists(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDynamicContactQueueingSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuleSets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Campaign) contextValidateCallAnalysisResponseSet(ctx context.Context, formats strfmt.Registry) error {

	if m.CallAnalysisResponseSet != nil {
		if err := m.CallAnalysisResponseSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callAnalysisResponseSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callAnalysisResponseSet")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateCallableTimeSet(ctx context.Context, formats strfmt.Registry) error {

	if m.CallableTimeSet != nil {
		if err := m.CallableTimeSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callableTimeSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callableTimeSet")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateContactList(ctx context.Context, formats strfmt.Registry) error {

	if m.ContactList != nil {
		if err := m.ContactList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contactList")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateContactListFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContactListFilters); i++ {

		if m.ContactListFilters[i] != nil {
			if err := m.ContactListFilters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactListFilters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contactListFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) contextValidateContactSort(ctx context.Context, formats strfmt.Registry) error {

	if m.ContactSort != nil {
		if err := m.ContactSort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactSort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contactSort")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateContactSorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContactSorts); i++ {

		if m.ContactSorts[i] != nil {
			if err := m.ContactSorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactSorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contactSorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

	if m.Division != nil {
		if err := m.Division.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateDncLists(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DncLists); i++ {

		if m.DncLists[i] != nil {
			if err := m.DncLists[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dncLists" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dncLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) contextValidateDynamicContactQueueingSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.DynamicContactQueueingSettings != nil {
		if err := m.DynamicContactQueueingSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamicContactQueueingSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamicContactQueueingSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateEdgeGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeGroup != nil {
		if err := m.EdgeGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errors", "body", []*RestErrorDetail(m.Errors)); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) contextValidatePhoneColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhoneColumns); i++ {

		if m.PhoneColumns[i] != nil {
			if err := m.PhoneColumns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneColumns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phoneColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {
		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateRuleSets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RuleSets); i++ {

		if m.RuleSets[i] != nil {
			if err := m.RuleSets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleSets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ruleSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Campaign) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

	if m.Script != nil {
		if err := m.Script.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *Campaign) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Campaign) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {
		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Campaign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Campaign) UnmarshalBinary(b []byte) error {
	var res Campaign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
