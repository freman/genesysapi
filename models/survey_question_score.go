// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SurveyQuestionScore survey question score
//
// swagger:model SurveyQuestionScore
type SurveyQuestionScore struct {

	// answer Id
	AnswerID string `json:"answerId,omitempty"`

	// free text answer
	FreeTextAnswer string `json:"freeTextAnswer,omitempty"`

	// marked n a
	MarkedNA bool `json:"markedNA,omitempty"`

	// nps score
	NpsScore int32 `json:"npsScore,omitempty"`

	// nps text answer
	NpsTextAnswer string `json:"npsTextAnswer,omitempty"`

	// question Id
	QuestionID string `json:"questionId,omitempty"`

	// score
	Score int32 `json:"score,omitempty"`
}

// Validate validates this survey question score
func (m *SurveyQuestionScore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SurveyQuestionScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SurveyQuestionScore) UnmarshalBinary(b []byte) error {
	var res SurveyQuestionScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
