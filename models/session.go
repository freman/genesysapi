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

// Session session
//
// swagger:model Session
type Session struct {

	// List-type attributes projected from the session's event stream.
	AttributeLists map[string]CustomEventAttributeList `json:"attributeLists,omitempty"`

	// Attributes projected from the session's event stream.
	Attributes map[string]CustomEventAttribute `json:"attributes,omitempty"`

	// Indicates whether or not the session is authenticated.
	Authenticated bool `json:"authenticated"`

	// Timestamp indicating when the visitor should be considered as away. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	AwayDate strfmt.DateTime `json:"awayDate,omitempty"`

	// Customer's browser.
	Browser *Browser `json:"browser,omitempty"`

	// The conversation for this session.
	// Read Only: true
	Conversation *AddressableEntityRef `json:"conversation,omitempty"`

	// Represents the channels used for this conversation.
	ConversationChannels []*ConversationChannel `json:"conversationChannels"`

	// The subject for the conversation, for example an email subject.
	ConversationSubject string `json:"conversationSubject,omitempty"`

	// Timestamp indicating when the session was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Primary identifier of the customer in the source where the events for the session originate from.
	CustomerID string `json:"customerId,omitempty"`

	// Type of source customer identifier (e.g. cookie, email, phone).
	CustomerIDType string `json:"customerIdType,omitempty"`

	// Customer's device.
	Device *Device `json:"device,omitempty"`

	// Indicates how long the session has been active (valid for an individual device).
	DurationInSeconds int32 `json:"durationInSeconds,omitempty"`

	// Timestamp indicating when the session was ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndedDate strfmt.DateTime `json:"endedDate,omitempty"`

	// The count of all events performed during the session.
	EventCount int32 `json:"eventCount,omitempty"`

	// The external contact associated with this session.
	// Read Only: true
	ExternalContact *AddressableEntityRef `json:"externalContact,omitempty"`

	// Unique identifier in the external system where the events for the session originate from.
	ExternalID string `json:"externalId,omitempty"`

	// A URL that identifies an external system-of-record resource that may have more detailed information on the session.
	ExternalURL string `json:"externalUrl,omitempty"`

	// Customer's geolocation.
	Geolocation *JourneyGeolocation `json:"geolocation,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Timestamp indicating when the visitor should be considered as idle. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	IdleDate strfmt.DateTime `json:"idleDate,omitempty"`

	// Customer's IP address.
	IPAddress string `json:"ipAddress,omitempty"`

	// Customer's IP-based organization or ISP name.
	IPOrganization string `json:"ipOrganization,omitempty"`

	// Last ACD outcome for the conversation.
	// Enum: [Unknown Abandon Answered FlowOut]
	LastAcdOutcome string `json:"lastAcdOutcome,omitempty"`

	// The last queue connected to this session.
	LastConnectedQueue *ConnectedQueue `json:"lastConnectedQueue,omitempty"`

	// The last user connected to this session.
	LastConnectedUser *ConnectedUser `json:"lastConnectedUser,omitempty"`

	// Information about the most recent event in this session.
	LastEvent *SessionLastEvent `json:"lastEvent,omitempty"`

	// The webpage where the customer's last web interaction occurred.
	LastPage *JourneyPage `json:"lastPage,omitempty"`

	// Disconnect reason for the last user connected to the conversation.
	// Enum: [Unknown Endpoint Client System Transfer Error Peer Other Spam Timeout TransportFailure ConferenceTransfer ConsultTransfer ForwardTransfer NoAnswerTransfer NotAvailableTransfer Uncallable]
	LastUserDisconnectType string `json:"lastUserDisconnectType,omitempty"`

	// The last user disposition connected to this session.
	LastUserDisposition *ConversationUserDisposition `json:"lastUserDisposition,omitempty"`

	// Marketing / traffic source information.
	MktCampaign *JourneyCampaign `json:"mktCampaign,omitempty"`

	// The original direction of the conversation.
	// Enum: [Unknown Inbound Outbound]
	OriginatingDirection string `json:"originatingDirection,omitempty"`

	// List of the outcome achievements by the customer in this session.
	OutcomeAchievements []*OutcomeAchievement `json:"outcomeAchievements"`

	// The count of all pageviews performed during the session.
	PageviewCount int32 `json:"pageviewCount,omitempty"`

	// Identifies the page URL that originally generated the request for the current page being viewed.
	Referrer *Referrer `json:"referrer,omitempty"`

	// The count of all screenviews performed during the session.
	ScreenviewCount int32 `json:"screenviewCount,omitempty"`

	// Search terms associated with the session.
	SearchTerms []string `json:"searchTerms"`

	// List of the segment assignments to the customer in this session.
	SegmentAssignments []*SessionSegmentAssignment `json:"segmentAssignments"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Shortened numeric identifier of 4-6 digits.
	ShortID string `json:"shortId,omitempty"`

	// Session types indicate the type or category of sessions (e.g. web, ticket, delivery, atm).
	Type string `json:"type,omitempty"`

	// String identifying the user agent.
	UserAgentString string `json:"userAgentString,omitempty"`
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeLists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwayDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBrowser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeolocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdleDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastAcdOutcome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastConnectedQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastConnectedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUserDisconnectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUserDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMktCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginatingDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutcomeAchievements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferrer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) validateAttributeLists(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeLists) { // not required
		return nil
	}

	for k := range m.AttributeLists {

		if err := validate.Required("attributeLists"+"."+k, "body", m.AttributeLists[k]); err != nil {
			return err
		}
		if val, ok := m.AttributeLists[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeLists" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeLists" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for k := range m.Attributes {

		if err := validate.Required("attributes"+"."+k, "body", m.Attributes[k]); err != nil {
			return err
		}
		if val, ok := m.Attributes[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) validateAwayDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AwayDate) { // not required
		return nil
	}

	if err := validate.FormatOf("awayDate", "body", "date-time", m.AwayDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateBrowser(formats strfmt.Registry) error {
	if swag.IsZero(m.Browser) { // not required
		return nil
	}

	if m.Browser != nil {
		if err := m.Browser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("browser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("browser")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateConversation(formats strfmt.Registry) error {
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

func (m *Session) validateConversationChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationChannels) { // not required
		return nil
	}

	for i := 0; i < len(m.ConversationChannels); i++ {
		if swag.IsZero(m.ConversationChannels[i]) { // not required
			continue
		}

		if m.ConversationChannels[i] != nil {
			if err := m.ConversationChannels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversationChannels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conversationChannels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateEndedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endedDate", "body", "date-time", m.EndedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateExternalContact(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalContact) { // not required
		return nil
	}

	if m.ExternalContact != nil {
		if err := m.ExternalContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateGeolocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Geolocation) { // not required
		return nil
	}

	if m.Geolocation != nil {
		if err := m.Geolocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geolocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geolocation")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateIdleDate(formats strfmt.Registry) error {
	if swag.IsZero(m.IdleDate) { // not required
		return nil
	}

	if err := validate.FormatOf("idleDate", "body", "date-time", m.IdleDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var sessionTypeLastAcdOutcomePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Abandon","Answered","FlowOut"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sessionTypeLastAcdOutcomePropEnum = append(sessionTypeLastAcdOutcomePropEnum, v)
	}
}

const (

	// SessionLastAcdOutcomeUnknown captures enum value "Unknown"
	SessionLastAcdOutcomeUnknown string = "Unknown"

	// SessionLastAcdOutcomeAbandon captures enum value "Abandon"
	SessionLastAcdOutcomeAbandon string = "Abandon"

	// SessionLastAcdOutcomeAnswered captures enum value "Answered"
	SessionLastAcdOutcomeAnswered string = "Answered"

	// SessionLastAcdOutcomeFlowOut captures enum value "FlowOut"
	SessionLastAcdOutcomeFlowOut string = "FlowOut"
)

// prop value enum
func (m *Session) validateLastAcdOutcomeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sessionTypeLastAcdOutcomePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Session) validateLastAcdOutcome(formats strfmt.Registry) error {
	if swag.IsZero(m.LastAcdOutcome) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastAcdOutcomeEnum("lastAcdOutcome", "body", m.LastAcdOutcome); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateLastConnectedQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.LastConnectedQueue) { // not required
		return nil
	}

	if m.LastConnectedQueue != nil {
		if err := m.LastConnectedQueue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastConnectedQueue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastConnectedQueue")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateLastConnectedUser(formats strfmt.Registry) error {
	if swag.IsZero(m.LastConnectedUser) { // not required
		return nil
	}

	if m.LastConnectedUser != nil {
		if err := m.LastConnectedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastConnectedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastConnectedUser")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateLastEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.LastEvent) { // not required
		return nil
	}

	if m.LastEvent != nil {
		if err := m.LastEvent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastEvent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastEvent")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateLastPage(formats strfmt.Registry) error {
	if swag.IsZero(m.LastPage) { // not required
		return nil
	}

	if m.LastPage != nil {
		if err := m.LastPage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastPage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastPage")
			}
			return err
		}
	}

	return nil
}

var sessionTypeLastUserDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Endpoint","Client","System","Transfer","Error","Peer","Other","Spam","Timeout","TransportFailure","ConferenceTransfer","ConsultTransfer","ForwardTransfer","NoAnswerTransfer","NotAvailableTransfer","Uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sessionTypeLastUserDisconnectTypePropEnum = append(sessionTypeLastUserDisconnectTypePropEnum, v)
	}
}

const (

	// SessionLastUserDisconnectTypeUnknown captures enum value "Unknown"
	SessionLastUserDisconnectTypeUnknown string = "Unknown"

	// SessionLastUserDisconnectTypeEndpoint captures enum value "Endpoint"
	SessionLastUserDisconnectTypeEndpoint string = "Endpoint"

	// SessionLastUserDisconnectTypeClient captures enum value "Client"
	SessionLastUserDisconnectTypeClient string = "Client"

	// SessionLastUserDisconnectTypeSystem captures enum value "System"
	SessionLastUserDisconnectTypeSystem string = "System"

	// SessionLastUserDisconnectTypeTransfer captures enum value "Transfer"
	SessionLastUserDisconnectTypeTransfer string = "Transfer"

	// SessionLastUserDisconnectTypeError captures enum value "Error"
	SessionLastUserDisconnectTypeError string = "Error"

	// SessionLastUserDisconnectTypePeer captures enum value "Peer"
	SessionLastUserDisconnectTypePeer string = "Peer"

	// SessionLastUserDisconnectTypeOther captures enum value "Other"
	SessionLastUserDisconnectTypeOther string = "Other"

	// SessionLastUserDisconnectTypeSpam captures enum value "Spam"
	SessionLastUserDisconnectTypeSpam string = "Spam"

	// SessionLastUserDisconnectTypeTimeout captures enum value "Timeout"
	SessionLastUserDisconnectTypeTimeout string = "Timeout"

	// SessionLastUserDisconnectTypeTransportFailure captures enum value "TransportFailure"
	SessionLastUserDisconnectTypeTransportFailure string = "TransportFailure"

	// SessionLastUserDisconnectTypeConferenceTransfer captures enum value "ConferenceTransfer"
	SessionLastUserDisconnectTypeConferenceTransfer string = "ConferenceTransfer"

	// SessionLastUserDisconnectTypeConsultTransfer captures enum value "ConsultTransfer"
	SessionLastUserDisconnectTypeConsultTransfer string = "ConsultTransfer"

	// SessionLastUserDisconnectTypeForwardTransfer captures enum value "ForwardTransfer"
	SessionLastUserDisconnectTypeForwardTransfer string = "ForwardTransfer"

	// SessionLastUserDisconnectTypeNoAnswerTransfer captures enum value "NoAnswerTransfer"
	SessionLastUserDisconnectTypeNoAnswerTransfer string = "NoAnswerTransfer"

	// SessionLastUserDisconnectTypeNotAvailableTransfer captures enum value "NotAvailableTransfer"
	SessionLastUserDisconnectTypeNotAvailableTransfer string = "NotAvailableTransfer"

	// SessionLastUserDisconnectTypeUncallable captures enum value "Uncallable"
	SessionLastUserDisconnectTypeUncallable string = "Uncallable"
)

// prop value enum
func (m *Session) validateLastUserDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sessionTypeLastUserDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Session) validateLastUserDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUserDisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastUserDisconnectTypeEnum("lastUserDisconnectType", "body", m.LastUserDisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateLastUserDisposition(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUserDisposition) { // not required
		return nil
	}

	if m.LastUserDisposition != nil {
		if err := m.LastUserDisposition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUserDisposition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastUserDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateMktCampaign(formats strfmt.Registry) error {
	if swag.IsZero(m.MktCampaign) { // not required
		return nil
	}

	if m.MktCampaign != nil {
		if err := m.MktCampaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mktCampaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mktCampaign")
			}
			return err
		}
	}

	return nil
}

var sessionTypeOriginatingDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Inbound","Outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sessionTypeOriginatingDirectionPropEnum = append(sessionTypeOriginatingDirectionPropEnum, v)
	}
}

const (

	// SessionOriginatingDirectionUnknown captures enum value "Unknown"
	SessionOriginatingDirectionUnknown string = "Unknown"

	// SessionOriginatingDirectionInbound captures enum value "Inbound"
	SessionOriginatingDirectionInbound string = "Inbound"

	// SessionOriginatingDirectionOutbound captures enum value "Outbound"
	SessionOriginatingDirectionOutbound string = "Outbound"
)

// prop value enum
func (m *Session) validateOriginatingDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sessionTypeOriginatingDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Session) validateOriginatingDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginatingDirection) { // not required
		return nil
	}

	// value enum
	if err := m.validateOriginatingDirectionEnum("originatingDirection", "body", m.OriginatingDirection); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateOutcomeAchievements(formats strfmt.Registry) error {
	if swag.IsZero(m.OutcomeAchievements) { // not required
		return nil
	}

	for i := 0; i < len(m.OutcomeAchievements); i++ {
		if swag.IsZero(m.OutcomeAchievements[i]) { // not required
			continue
		}

		if m.OutcomeAchievements[i] != nil {
			if err := m.OutcomeAchievements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outcomeAchievements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outcomeAchievements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) validateReferrer(formats strfmt.Registry) error {
	if swag.IsZero(m.Referrer) { // not required
		return nil
	}

	if m.Referrer != nil {
		if err := m.Referrer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referrer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referrer")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateSegmentAssignments(formats strfmt.Registry) error {
	if swag.IsZero(m.SegmentAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.SegmentAssignments); i++ {
		if swag.IsZero(m.SegmentAssignments[i]) { // not required
			continue
		}

		if m.SegmentAssignments[i] != nil {
			if err := m.SegmentAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segmentAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segmentAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this session based on the context it is used
func (m *Session) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributeLists(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBrowser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeolocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastConnectedQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastConnectedUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastPage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUserDisposition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMktCampaign(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutcomeAchievements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferrer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSegmentAssignments(ctx, formats); err != nil {
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

func (m *Session) contextValidateAttributeLists(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.AttributeLists {

		if val, ok := m.AttributeLists[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Session) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Attributes {

		if val, ok := m.Attributes[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Session) contextValidateBrowser(ctx context.Context, formats strfmt.Registry) error {

	if m.Browser != nil {
		if err := m.Browser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("browser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("browser")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateConversation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Session) contextValidateConversationChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConversationChannels); i++ {

		if m.ConversationChannels[i] != nil {
			if err := m.ConversationChannels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversationChannels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conversationChannels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateExternalContact(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalContact != nil {
		if err := m.ExternalContact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateGeolocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Geolocation != nil {
		if err := m.Geolocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geolocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geolocation")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateLastConnectedQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.LastConnectedQueue != nil {
		if err := m.LastConnectedQueue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastConnectedQueue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastConnectedQueue")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateLastConnectedUser(ctx context.Context, formats strfmt.Registry) error {

	if m.LastConnectedUser != nil {
		if err := m.LastConnectedUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastConnectedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastConnectedUser")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateLastEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.LastEvent != nil {
		if err := m.LastEvent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastEvent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastEvent")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateLastPage(ctx context.Context, formats strfmt.Registry) error {

	if m.LastPage != nil {
		if err := m.LastPage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastPage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastPage")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateLastUserDisposition(ctx context.Context, formats strfmt.Registry) error {

	if m.LastUserDisposition != nil {
		if err := m.LastUserDisposition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUserDisposition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastUserDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateMktCampaign(ctx context.Context, formats strfmt.Registry) error {

	if m.MktCampaign != nil {
		if err := m.MktCampaign.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mktCampaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mktCampaign")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateOutcomeAchievements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OutcomeAchievements); i++ {

		if m.OutcomeAchievements[i] != nil {
			if err := m.OutcomeAchievements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outcomeAchievements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outcomeAchievements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) contextValidateReferrer(ctx context.Context, formats strfmt.Registry) error {

	if m.Referrer != nil {
		if err := m.Referrer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referrer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("referrer")
			}
			return err
		}
	}

	return nil
}

func (m *Session) contextValidateSegmentAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SegmentAssignments); i++ {

		if m.SegmentAssignments[i] != nil {
			if err := m.SegmentAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segmentAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segmentAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Session) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
