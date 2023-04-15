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

// DialerAction dialer action
//
// swagger:model DialerAction
type DialerAction struct {

	// Additional type specification for this DialerAction.
	// Required: true
	// Enum: [DO_NOT_DIAL MODIFY_CONTACT_ATTRIBUTE SWITCH_TO_PREVIEW APPEND_NUMBER_TO_DNC_LIST APPEND_CUSTOM_ENTRY_TO_DNC_LIST SCHEDULE_CALLBACK CONTACT_UNCALLABLE NUMBER_UNCALLABLE SET_CALLER_ID SET_SKILLS DATA_ACTION]
	ActionTypeName *string `json:"actionTypeName"`

	// The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionBehavior.
	AgentWrapupField string `json:"agentWrapupField,omitempty"`

	// The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionBehavior.
	CallAnalysisResultField string `json:"callAnalysisResultField,omitempty"`

	// A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionBehavior.
	ContactColumnToDataActionFieldMappings []*ContactColumnToDataActionFieldMapping `json:"contactColumnToDataActionFieldMappings"`

	// The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionBehavior.
	ContactIDField string `json:"contactIdField,omitempty"`

	// The Data Action to use for this action. Required for a dataActionBehavior.
	DataAction *DomainEntityRef `json:"dataAction,omitempty"`

	// A map of key-value pairs pertinent to the DialerAction. Different types of DialerActions require different properties. MODIFY_CONTACT_ATTRIBUTE with an updateOption of SET takes a contact column as the key and accepts any value. SCHEDULE_CALLBACK takes a key 'callbackOffset' that specifies how far in the future the callback should be scheduled, in minutes. SET_CALLER_ID takes two keys: 'callerAddress', which should be the caller id phone number, and 'callerName'. For either key, you can also specify a column on the contact to get the value from. To do this, specify 'contact.Column', where 'Column' is the name of the contact column from which to get the value. SET_SKILLS takes a key 'skills' with an array of skill ids wrapped into a string (Example: {'skills': '['skillIdHere']'} ).
	Properties map[string]string `json:"properties,omitempty"`

	// The type of this DialerAction.
	// Required: true
	// Enum: [Action modifyContactAttribute dataActionBehavior]
	Type *string `json:"type"`

	// Specifies how a contact attribute should be updated. Required for MODIFY_CONTACT_ATTRIBUTE.
	// Enum: [SET INCREMENT DECREMENT CURRENT_TIME]
	UpdateOption string `json:"updateOption,omitempty"`
}

// Validate validates this dialer action
func (m *DialerAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionTypeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactColumnToDataActionFieldMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dialerActionTypeActionTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DO_NOT_DIAL","MODIFY_CONTACT_ATTRIBUTE","SWITCH_TO_PREVIEW","APPEND_NUMBER_TO_DNC_LIST","APPEND_CUSTOM_ENTRY_TO_DNC_LIST","SCHEDULE_CALLBACK","CONTACT_UNCALLABLE","NUMBER_UNCALLABLE","SET_CALLER_ID","SET_SKILLS","DATA_ACTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dialerActionTypeActionTypeNamePropEnum = append(dialerActionTypeActionTypeNamePropEnum, v)
	}
}

const (

	// DialerActionActionTypeNameDONOTDIAL captures enum value "DO_NOT_DIAL"
	DialerActionActionTypeNameDONOTDIAL string = "DO_NOT_DIAL"

	// DialerActionActionTypeNameMODIFYCONTACTATTRIBUTE captures enum value "MODIFY_CONTACT_ATTRIBUTE"
	DialerActionActionTypeNameMODIFYCONTACTATTRIBUTE string = "MODIFY_CONTACT_ATTRIBUTE"

	// DialerActionActionTypeNameSWITCHTOPREVIEW captures enum value "SWITCH_TO_PREVIEW"
	DialerActionActionTypeNameSWITCHTOPREVIEW string = "SWITCH_TO_PREVIEW"

	// DialerActionActionTypeNameAPPENDNUMBERTODNCLIST captures enum value "APPEND_NUMBER_TO_DNC_LIST"
	DialerActionActionTypeNameAPPENDNUMBERTODNCLIST string = "APPEND_NUMBER_TO_DNC_LIST"

	// DialerActionActionTypeNameAPPENDCUSTOMENTRYTODNCLIST captures enum value "APPEND_CUSTOM_ENTRY_TO_DNC_LIST"
	DialerActionActionTypeNameAPPENDCUSTOMENTRYTODNCLIST string = "APPEND_CUSTOM_ENTRY_TO_DNC_LIST"

	// DialerActionActionTypeNameSCHEDULECALLBACK captures enum value "SCHEDULE_CALLBACK"
	DialerActionActionTypeNameSCHEDULECALLBACK string = "SCHEDULE_CALLBACK"

	// DialerActionActionTypeNameCONTACTUNCALLABLE captures enum value "CONTACT_UNCALLABLE"
	DialerActionActionTypeNameCONTACTUNCALLABLE string = "CONTACT_UNCALLABLE"

	// DialerActionActionTypeNameNUMBERUNCALLABLE captures enum value "NUMBER_UNCALLABLE"
	DialerActionActionTypeNameNUMBERUNCALLABLE string = "NUMBER_UNCALLABLE"

	// DialerActionActionTypeNameSETCALLERID captures enum value "SET_CALLER_ID"
	DialerActionActionTypeNameSETCALLERID string = "SET_CALLER_ID"

	// DialerActionActionTypeNameSETSKILLS captures enum value "SET_SKILLS"
	DialerActionActionTypeNameSETSKILLS string = "SET_SKILLS"

	// DialerActionActionTypeNameDATAACTION captures enum value "DATA_ACTION"
	DialerActionActionTypeNameDATAACTION string = "DATA_ACTION"
)

// prop value enum
func (m *DialerAction) validateActionTypeNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dialerActionTypeActionTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DialerAction) validateActionTypeName(formats strfmt.Registry) error {

	if err := validate.Required("actionTypeName", "body", m.ActionTypeName); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionTypeNameEnum("actionTypeName", "body", *m.ActionTypeName); err != nil {
		return err
	}

	return nil
}

func (m *DialerAction) validateContactColumnToDataActionFieldMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactColumnToDataActionFieldMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactColumnToDataActionFieldMappings); i++ {
		if swag.IsZero(m.ContactColumnToDataActionFieldMappings[i]) { // not required
			continue
		}

		if m.ContactColumnToDataActionFieldMappings[i] != nil {
			if err := m.ContactColumnToDataActionFieldMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactColumnToDataActionFieldMappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contactColumnToDataActionFieldMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DialerAction) validateDataAction(formats strfmt.Registry) error {
	if swag.IsZero(m.DataAction) { // not required
		return nil
	}

	if m.DataAction != nil {
		if err := m.DataAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataAction")
			}
			return err
		}
	}

	return nil
}

var dialerActionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Action","modifyContactAttribute","dataActionBehavior"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dialerActionTypeTypePropEnum = append(dialerActionTypeTypePropEnum, v)
	}
}

const (

	// DialerActionTypeAction captures enum value "Action"
	DialerActionTypeAction string = "Action"

	// DialerActionTypeModifyContactAttribute captures enum value "modifyContactAttribute"
	DialerActionTypeModifyContactAttribute string = "modifyContactAttribute"

	// DialerActionTypeDataActionBehavior captures enum value "dataActionBehavior"
	DialerActionTypeDataActionBehavior string = "dataActionBehavior"
)

// prop value enum
func (m *DialerAction) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dialerActionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DialerAction) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var dialerActionTypeUpdateOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SET","INCREMENT","DECREMENT","CURRENT_TIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dialerActionTypeUpdateOptionPropEnum = append(dialerActionTypeUpdateOptionPropEnum, v)
	}
}

const (

	// DialerActionUpdateOptionSET captures enum value "SET"
	DialerActionUpdateOptionSET string = "SET"

	// DialerActionUpdateOptionINCREMENT captures enum value "INCREMENT"
	DialerActionUpdateOptionINCREMENT string = "INCREMENT"

	// DialerActionUpdateOptionDECREMENT captures enum value "DECREMENT"
	DialerActionUpdateOptionDECREMENT string = "DECREMENT"

	// DialerActionUpdateOptionCURRENTTIME captures enum value "CURRENT_TIME"
	DialerActionUpdateOptionCURRENTTIME string = "CURRENT_TIME"
)

// prop value enum
func (m *DialerAction) validateUpdateOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dialerActionTypeUpdateOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DialerAction) validateUpdateOption(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateOption) { // not required
		return nil
	}

	// value enum
	if err := m.validateUpdateOptionEnum("updateOption", "body", m.UpdateOption); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dialer action based on the context it is used
func (m *DialerAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContactColumnToDataActionFieldMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DialerAction) contextValidateContactColumnToDataActionFieldMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContactColumnToDataActionFieldMappings); i++ {

		if m.ContactColumnToDataActionFieldMappings[i] != nil {
			if err := m.ContactColumnToDataActionFieldMappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactColumnToDataActionFieldMappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contactColumnToDataActionFieldMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DialerAction) contextValidateDataAction(ctx context.Context, formats strfmt.Registry) error {

	if m.DataAction != nil {
		if err := m.DataAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataAction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DialerAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DialerAction) UnmarshalBinary(b []byte) error {
	var res DialerAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
