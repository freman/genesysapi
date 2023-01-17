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

// CampaignInteraction campaign interaction
//
// swagger:model CampaignInteraction
type CampaignInteraction struct {

	// Boolean value if there is an active preview call on the interaction
	ActivePreviewCall bool `json:"activePreviewCall"`

	// agent
	Agent *DomainEntityRef `json:"agent,omitempty"`

	// The time when the agent or system places the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CallPlacedTime strfmt.DateTime `json:"callPlacedTime,omitempty"`

	// The time when the agent was connected to the call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CallRoutedTime strfmt.DateTime `json:"callRoutedTime,omitempty"`

	// caller address
	CallerAddress string `json:"callerAddress,omitempty"`

	// caller name
	CallerName string `json:"callerName,omitempty"`

	// campaign
	Campaign *DomainEntityRef `json:"campaign,omitempty"`

	// contact
	Contact *DomainEntityRef `json:"contact,omitempty"`

	// conversation
	Conversation *ConversationBasic `json:"conversation,omitempty"`

	// The time when dialer created the interaction. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creationTime,omitempty"`

	// destination address
	DestinationAddress string `json:"destinationAddress,omitempty"`

	// conversation participant id that is the dialer system participant to monitor the call from dialer perspective
	DialerSystemParticipantID string `json:"dialerSystemParticipantId,omitempty"`

	// dialing mode
	DialingMode string `json:"dialingMode,omitempty"`

	// Describes what happened with call analysis for instance: disposition.classification.callable.person, disposition.classification.callable.noanswer
	// Enum: [DISCONNECT LIVE_VOICE BUSY MACHINE NO_ANSWER SIT_CALLABLE SIT_UNCALLABLE FAX]
	Disposition string `json:"disposition,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// The time when the last preview of the interaction was wrapped up. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	LastActivePreviewWrapupTime strfmt.DateTime `json:"lastActivePreviewWrapupTime,omitempty"`

	// The time when the customer and routing participant are connected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	PreviewConnectedTime strfmt.DateTime `json:"previewConnectedTime,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	PreviewPopDeliveredTime strfmt.DateTime `json:"previewPopDeliveredTime,omitempty"`

	// queue
	Queue *DomainEntityRef `json:"queue,omitempty"`

	// script
	Script *DomainEntityRef `json:"script,omitempty"`

	// Any skills that are attached to the call for routing
	// Unique: true
	Skills []*DomainEntityRef `json:"skills"`
}

// Validate validates this campaign interaction
func (m *CampaignInteraction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallPlacedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallRoutedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActivePreviewWrapupTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviewConnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviewPopDeliveredTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkills(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignInteraction) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) validateCallPlacedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CallPlacedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("callPlacedTime", "body", "date-time", m.CallPlacedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignInteraction) validateCallRoutedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CallRoutedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("callRoutedTime", "body", "date-time", m.CallRoutedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignInteraction) validateCampaign(formats strfmt.Registry) error {
	if swag.IsZero(m.Campaign) { // not required
		return nil
	}

	if m.Campaign != nil {
		if err := m.Campaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("campaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("campaign")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) validateContact(formats strfmt.Registry) error {
	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) validateConversation(formats strfmt.Registry) error {
	if swag.IsZero(m.Conversation) { // not required
		return nil
	}

	if m.Conversation != nil {
		if err := m.Conversation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTime", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var campaignInteractionTypeDispositionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISCONNECT","LIVE_VOICE","BUSY","MACHINE","NO_ANSWER","SIT_CALLABLE","SIT_UNCALLABLE","FAX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignInteractionTypeDispositionPropEnum = append(campaignInteractionTypeDispositionPropEnum, v)
	}
}

const (

	// CampaignInteractionDispositionDISCONNECT captures enum value "DISCONNECT"
	CampaignInteractionDispositionDISCONNECT string = "DISCONNECT"

	// CampaignInteractionDispositionLIVEVOICE captures enum value "LIVE_VOICE"
	CampaignInteractionDispositionLIVEVOICE string = "LIVE_VOICE"

	// CampaignInteractionDispositionBUSY captures enum value "BUSY"
	CampaignInteractionDispositionBUSY string = "BUSY"

	// CampaignInteractionDispositionMACHINE captures enum value "MACHINE"
	CampaignInteractionDispositionMACHINE string = "MACHINE"

	// CampaignInteractionDispositionNOANSWER captures enum value "NO_ANSWER"
	CampaignInteractionDispositionNOANSWER string = "NO_ANSWER"

	// CampaignInteractionDispositionSITCALLABLE captures enum value "SIT_CALLABLE"
	CampaignInteractionDispositionSITCALLABLE string = "SIT_CALLABLE"

	// CampaignInteractionDispositionSITUNCALLABLE captures enum value "SIT_UNCALLABLE"
	CampaignInteractionDispositionSITUNCALLABLE string = "SIT_UNCALLABLE"

	// CampaignInteractionDispositionFAX captures enum value "FAX"
	CampaignInteractionDispositionFAX string = "FAX"
)

// prop value enum
func (m *CampaignInteraction) validateDispositionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignInteractionTypeDispositionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CampaignInteraction) validateDisposition(formats strfmt.Registry) error {
	if swag.IsZero(m.Disposition) { // not required
		return nil
	}

	// value enum
	if err := m.validateDispositionEnum("disposition", "body", m.Disposition); err != nil {
		return err
	}

	return nil
}

func (m *CampaignInteraction) validateLastActivePreviewWrapupTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActivePreviewWrapupTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastActivePreviewWrapupTime", "body", "date-time", m.LastActivePreviewWrapupTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignInteraction) validatePreviewConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PreviewConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("previewConnectedTime", "body", "date-time", m.PreviewConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignInteraction) validatePreviewPopDeliveredTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PreviewPopDeliveredTime) { // not required
		return nil
	}

	if err := validate.FormatOf("previewPopDeliveredTime", "body", "date-time", m.PreviewPopDeliveredTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignInteraction) validateQueue(formats strfmt.Registry) error {
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

func (m *CampaignInteraction) validateScript(formats strfmt.Registry) error {
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

func (m *CampaignInteraction) validateSkills(formats strfmt.Registry) error {
	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	if err := validate.UniqueItems("skills", "body", m.Skills); err != nil {
		return err
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this campaign interaction based on the context it is used
func (m *CampaignInteraction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCampaign(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSkills(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignInteraction) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {
		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) contextValidateCampaign(ctx context.Context, formats strfmt.Registry) error {

	if m.Campaign != nil {
		if err := m.Campaign.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("campaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("campaign")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) contextValidateContact(ctx context.Context, formats strfmt.Registry) error {

	if m.Contact != nil {
		if err := m.Contact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) contextValidateConversation(ctx context.Context, formats strfmt.Registry) error {

	if m.Conversation != nil {
		if err := m.Conversation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignInteraction) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CampaignInteraction) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CampaignInteraction) contextValidateSkills(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Skills); i++ {

		if m.Skills[i] != nil {
			if err := m.Skills[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignInteraction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignInteraction) UnmarshalBinary(b []byte) error {
	var res CampaignInteraction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
