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

// Evaluation evaluation
//
// swagger:model Evaluation
type Evaluation struct {

	// agent
	Agent *User `json:"agent,omitempty"`

	// agent has read
	AgentHasRead bool `json:"agentHasRead"`

	// answers
	Answers *EvaluationScoringSet `json:"answers,omitempty"`

	// Set to false to unassign the evaluation. This cannot be set to false when assignee is also set.
	Assigned bool `json:"assigned"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	AssignedDate strfmt.DateTime `json:"assignedDate,omitempty"`

	// assignee
	Assignee *User `json:"assignee,omitempty"`

	// List of user authorized actions on evaluation. Possible values: assign, edit, editScore, editAgentSignoff, delete, release, viewAudit
	AuthorizedActions []string `json:"authorizedActions"`

	// calibration
	Calibration *Calibration `json:"calibration,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ChangedDate strfmt.DateTime `json:"changedDate,omitempty"`

	// conversation
	Conversation *ConversationReference `json:"conversation,omitempty"`

	// Date of conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConversationDate strfmt.DateTime `json:"conversationDate,omitempty"`

	// End date of conversation if it had completed before evaluation creation. Null if created before the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConversationEndDate strfmt.DateTime `json:"conversationEndDate,omitempty"`

	// Evaluation form used for evaluation.
	EvaluationForm *EvaluationForm `json:"evaluationForm,omitempty"`

	// The source that created the evaluation.
	// Read Only: true
	EvaluationSource *EvaluationSource `json:"evaluationSource,omitempty"`

	// evaluator
	Evaluator *User `json:"evaluator,omitempty"`

	// Is true when evaluation assistance didn't execute successfully
	HasAssistanceFailed bool `json:"hasAssistanceFailed"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// is scoring index
	IsScoringIndex bool `json:"isScoringIndex"`

	// List of different communication types used in conversation.
	MediaType []string `json:"mediaType"`

	// name
	Name string `json:"name,omitempty"`

	// Signifies if the evaluation is never to be released. This cannot be set true if release date is also set.
	NeverRelease bool `json:"neverRelease"`

	// queue
	Queue *Queue `json:"queue,omitempty"`

	// Is only true when the user making the request does not have sufficient permissions to see evaluation
	Redacted bool `json:"redacted"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReleaseDate strfmt.DateTime `json:"releaseDate,omitempty"`

	// Is only true when evaluation is re-scored.
	Rescore bool `json:"rescore"`

	// Only used for email evaluations. Will be null for all other evaluations.
	ResourceID string `json:"resourceId,omitempty"`

	// The type of resource. Only used for email evaluations. Will be null for evaluations on all other resources.
	// Enum: [EMAIL]
	ResourceType string `json:"resourceType,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// status
	// Enum: [PENDING INPROGRESS FINISHED INREVIEW RETRACTED]
	Status string `json:"status,omitempty"`
}

// Validate validates this evaluation
func (m *Evaluation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalibration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
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

func (m *Evaluation) validateAgent(formats strfmt.Registry) error {
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

func (m *Evaluation) validateAnswers(formats strfmt.Registry) error {
	if swag.IsZero(m.Answers) { // not required
		return nil
	}

	if m.Answers != nil {
		if err := m.Answers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("answers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("answers")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) validateAssignedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AssignedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("assignedDate", "body", "date-time", m.AssignedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Evaluation) validateAssignee(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignee) { // not required
		return nil
	}

	if m.Assignee != nil {
		if err := m.Assignee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) validateCalibration(formats strfmt.Registry) error {
	if swag.IsZero(m.Calibration) { // not required
		return nil
	}

	if m.Calibration != nil {
		if err := m.Calibration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calibration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("calibration")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) validateChangedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("changedDate", "body", "date-time", m.ChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Evaluation) validateConversation(formats strfmt.Registry) error {
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

func (m *Evaluation) validateConversationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationDate", "body", "date-time", m.ConversationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Evaluation) validateConversationEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationEndDate", "body", "date-time", m.ConversationEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Evaluation) validateEvaluationForm(formats strfmt.Registry) error {
	if swag.IsZero(m.EvaluationForm) { // not required
		return nil
	}

	if m.EvaluationForm != nil {
		if err := m.EvaluationForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluationForm")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) validateEvaluationSource(formats strfmt.Registry) error {
	if swag.IsZero(m.EvaluationSource) { // not required
		return nil
	}

	if m.EvaluationSource != nil {
		if err := m.EvaluationSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluationSource")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) validateEvaluator(formats strfmt.Registry) error {
	if swag.IsZero(m.Evaluator) { // not required
		return nil
	}

	if m.Evaluator != nil {
		if err := m.Evaluator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluator")
			}
			return err
		}
	}

	return nil
}

var evaluationMediaTypeItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CALL","CALLBACK","CHAT","COBROWSE","EMAIL","MESSAGE","SOCIAL_EXPRESSION","VIDEO","SCREENSHARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationMediaTypeItemsEnum = append(evaluationMediaTypeItemsEnum, v)
	}
}

func (m *Evaluation) validateMediaTypeItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationMediaTypeItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Evaluation) validateMediaType(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaType); i++ {

		// value enum
		if err := m.validateMediaTypeItemsEnum("mediaType"+"."+strconv.Itoa(i), "body", m.MediaType[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Evaluation) validateQueue(formats strfmt.Registry) error {
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

func (m *Evaluation) validateReleaseDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseDate", "body", "date-time", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var evaluationTypeResourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EMAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationTypeResourceTypePropEnum = append(evaluationTypeResourceTypePropEnum, v)
	}
}

const (

	// EvaluationResourceTypeEMAIL captures enum value "EMAIL"
	EvaluationResourceTypeEMAIL string = "EMAIL"
)

// prop value enum
func (m *Evaluation) validateResourceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationTypeResourceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Evaluation) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceTypeEnum("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *Evaluation) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var evaluationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","INPROGRESS","FINISHED","INREVIEW","RETRACTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationTypeStatusPropEnum = append(evaluationTypeStatusPropEnum, v)
	}
}

const (

	// EvaluationStatusPENDING captures enum value "PENDING"
	EvaluationStatusPENDING string = "PENDING"

	// EvaluationStatusINPROGRESS captures enum value "INPROGRESS"
	EvaluationStatusINPROGRESS string = "INPROGRESS"

	// EvaluationStatusFINISHED captures enum value "FINISHED"
	EvaluationStatusFINISHED string = "FINISHED"

	// EvaluationStatusINREVIEW captures enum value "INREVIEW"
	EvaluationStatusINREVIEW string = "INREVIEW"

	// EvaluationStatusRETRACTED captures enum value "RETRACTED"
	EvaluationStatusRETRACTED string = "RETRACTED"
)

// prop value enum
func (m *Evaluation) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Evaluation) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this evaluation based on the context it is used
func (m *Evaluation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnswers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssignee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCalibration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvaluationForm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvaluationSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvaluator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
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

func (m *Evaluation) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Evaluation) contextValidateAnswers(ctx context.Context, formats strfmt.Registry) error {

	if m.Answers != nil {
		if err := m.Answers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("answers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("answers")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) contextValidateAssignee(ctx context.Context, formats strfmt.Registry) error {

	if m.Assignee != nil {
		if err := m.Assignee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) contextValidateCalibration(ctx context.Context, formats strfmt.Registry) error {

	if m.Calibration != nil {
		if err := m.Calibration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calibration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("calibration")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) contextValidateConversation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Evaluation) contextValidateEvaluationForm(ctx context.Context, formats strfmt.Registry) error {

	if m.EvaluationForm != nil {
		if err := m.EvaluationForm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluationForm")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) contextValidateEvaluationSource(ctx context.Context, formats strfmt.Registry) error {

	if m.EvaluationSource != nil {
		if err := m.EvaluationSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluationSource")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) contextValidateEvaluator(ctx context.Context, formats strfmt.Registry) error {

	if m.Evaluator != nil {
		if err := m.Evaluator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("evaluator")
			}
			return err
		}
	}

	return nil
}

func (m *Evaluation) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Evaluation) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Evaluation) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Evaluation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Evaluation) UnmarshalBinary(b []byte) error {
	var res Evaluation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
