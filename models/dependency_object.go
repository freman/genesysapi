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

// DependencyObject dependency object
//
// swagger:model DependencyObject
type DependencyObject struct {

	// consumed resources
	ConsumedResources []*Dependency `json:"consumedResources"`

	// consuming resources
	ConsumingResources []*Dependency `json:"consumingResources"`

	// deleted
	Deleted bool `json:"deleted"`

	// The dependency identifier
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// state unknown
	StateUnknown bool `json:"stateUnknown"`

	// type
	// Enum: [ACDLANGUAGE ACDSKILL ACDWRAPUPCODE BOTCONNECTORBOT BOTCONNECTORINTEGRATION BOTFLOW BRIDGEACTION COMMONMODULEFLOW COMPOSERSCRIPT CONTACTLIST DATAACTION DATATABLE DIALOGENGINEBOT DIALOGENGINEBOTVERSION DIALOGFLOWAGENT DIALOGFLOWCXAGENT DIGITALBOTFLOW EMAILROUTE EMERGENCYGROUP FLOWACTION FLOWDATATYPE FLOWMILESTONE FLOWOUTCOME GRAMMAR GROUP IMAGE INBOUNDCALLFLOW INBOUNDCHATFLOW INBOUNDEMAILFLOW INBOUNDSHORTMESSAGEFLOW INQUEUECALLFLOW INQUEUEEMAILFLOW INQUEUESHORTMESSAGEFLOW IVRCONFIGURATION KNOWLEDGEBASE KNOWLEDGEBASEDOCUMENT LANGUAGE LEXBOT LEXBOTALIAS LEXV2BOT LEXV2BOTALIAS NLUDOMAIN NUANCEMIXBOT NUANCEMIXINTEGRATION OAUTHCLIENT OUTBOUNDCALLFLOW QUEUE RECORDINGPOLICY RESPONSE SCHEDULE SCHEDULEGROUP SECUREACTION SECURECALLFLOW STTENGINE SURVEYINVITEFLOW SYSTEMPROMPT TTSENGINE TTSVOICE USER USERPROMPT VOICEFLOW VOICEMAILFLOW WIDGET WORKFLOW WORKITEMFLOW]
	Type string `json:"type,omitempty"`

	// updated
	Updated bool `json:"updated"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this dependency object
func (m *DependencyObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumedResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumingResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *DependencyObject) validateConsumedResources(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumedResources) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsumedResources); i++ {
		if swag.IsZero(m.ConsumedResources[i]) { // not required
			continue
		}

		if m.ConsumedResources[i] != nil {
			if err := m.ConsumedResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumedResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumedResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DependencyObject) validateConsumingResources(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumingResources) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsumingResources); i++ {
		if swag.IsZero(m.ConsumingResources[i]) { // not required
			continue
		}

		if m.ConsumingResources[i] != nil {
			if err := m.ConsumingResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumingResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumingResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DependencyObject) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var dependencyObjectTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACDLANGUAGE","ACDSKILL","ACDWRAPUPCODE","BOTCONNECTORBOT","BOTCONNECTORINTEGRATION","BOTFLOW","BRIDGEACTION","COMMONMODULEFLOW","COMPOSERSCRIPT","CONTACTLIST","DATAACTION","DATATABLE","DIALOGENGINEBOT","DIALOGENGINEBOTVERSION","DIALOGFLOWAGENT","DIALOGFLOWCXAGENT","DIGITALBOTFLOW","EMAILROUTE","EMERGENCYGROUP","FLOWACTION","FLOWDATATYPE","FLOWMILESTONE","FLOWOUTCOME","GRAMMAR","GROUP","IMAGE","INBOUNDCALLFLOW","INBOUNDCHATFLOW","INBOUNDEMAILFLOW","INBOUNDSHORTMESSAGEFLOW","INQUEUECALLFLOW","INQUEUEEMAILFLOW","INQUEUESHORTMESSAGEFLOW","IVRCONFIGURATION","KNOWLEDGEBASE","KNOWLEDGEBASEDOCUMENT","LANGUAGE","LEXBOT","LEXBOTALIAS","LEXV2BOT","LEXV2BOTALIAS","NLUDOMAIN","NUANCEMIXBOT","NUANCEMIXINTEGRATION","OAUTHCLIENT","OUTBOUNDCALLFLOW","QUEUE","RECORDINGPOLICY","RESPONSE","SCHEDULE","SCHEDULEGROUP","SECUREACTION","SECURECALLFLOW","STTENGINE","SURVEYINVITEFLOW","SYSTEMPROMPT","TTSENGINE","TTSVOICE","USER","USERPROMPT","VOICEFLOW","VOICEMAILFLOW","WIDGET","WORKFLOW","WORKITEMFLOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dependencyObjectTypeTypePropEnum = append(dependencyObjectTypeTypePropEnum, v)
	}
}

const (

	// DependencyObjectTypeACDLANGUAGE captures enum value "ACDLANGUAGE"
	DependencyObjectTypeACDLANGUAGE string = "ACDLANGUAGE"

	// DependencyObjectTypeACDSKILL captures enum value "ACDSKILL"
	DependencyObjectTypeACDSKILL string = "ACDSKILL"

	// DependencyObjectTypeACDWRAPUPCODE captures enum value "ACDWRAPUPCODE"
	DependencyObjectTypeACDWRAPUPCODE string = "ACDWRAPUPCODE"

	// DependencyObjectTypeBOTCONNECTORBOT captures enum value "BOTCONNECTORBOT"
	DependencyObjectTypeBOTCONNECTORBOT string = "BOTCONNECTORBOT"

	// DependencyObjectTypeBOTCONNECTORINTEGRATION captures enum value "BOTCONNECTORINTEGRATION"
	DependencyObjectTypeBOTCONNECTORINTEGRATION string = "BOTCONNECTORINTEGRATION"

	// DependencyObjectTypeBOTFLOW captures enum value "BOTFLOW"
	DependencyObjectTypeBOTFLOW string = "BOTFLOW"

	// DependencyObjectTypeBRIDGEACTION captures enum value "BRIDGEACTION"
	DependencyObjectTypeBRIDGEACTION string = "BRIDGEACTION"

	// DependencyObjectTypeCOMMONMODULEFLOW captures enum value "COMMONMODULEFLOW"
	DependencyObjectTypeCOMMONMODULEFLOW string = "COMMONMODULEFLOW"

	// DependencyObjectTypeCOMPOSERSCRIPT captures enum value "COMPOSERSCRIPT"
	DependencyObjectTypeCOMPOSERSCRIPT string = "COMPOSERSCRIPT"

	// DependencyObjectTypeCONTACTLIST captures enum value "CONTACTLIST"
	DependencyObjectTypeCONTACTLIST string = "CONTACTLIST"

	// DependencyObjectTypeDATAACTION captures enum value "DATAACTION"
	DependencyObjectTypeDATAACTION string = "DATAACTION"

	// DependencyObjectTypeDATATABLE captures enum value "DATATABLE"
	DependencyObjectTypeDATATABLE string = "DATATABLE"

	// DependencyObjectTypeDIALOGENGINEBOT captures enum value "DIALOGENGINEBOT"
	DependencyObjectTypeDIALOGENGINEBOT string = "DIALOGENGINEBOT"

	// DependencyObjectTypeDIALOGENGINEBOTVERSION captures enum value "DIALOGENGINEBOTVERSION"
	DependencyObjectTypeDIALOGENGINEBOTVERSION string = "DIALOGENGINEBOTVERSION"

	// DependencyObjectTypeDIALOGFLOWAGENT captures enum value "DIALOGFLOWAGENT"
	DependencyObjectTypeDIALOGFLOWAGENT string = "DIALOGFLOWAGENT"

	// DependencyObjectTypeDIALOGFLOWCXAGENT captures enum value "DIALOGFLOWCXAGENT"
	DependencyObjectTypeDIALOGFLOWCXAGENT string = "DIALOGFLOWCXAGENT"

	// DependencyObjectTypeDIGITALBOTFLOW captures enum value "DIGITALBOTFLOW"
	DependencyObjectTypeDIGITALBOTFLOW string = "DIGITALBOTFLOW"

	// DependencyObjectTypeEMAILROUTE captures enum value "EMAILROUTE"
	DependencyObjectTypeEMAILROUTE string = "EMAILROUTE"

	// DependencyObjectTypeEMERGENCYGROUP captures enum value "EMERGENCYGROUP"
	DependencyObjectTypeEMERGENCYGROUP string = "EMERGENCYGROUP"

	// DependencyObjectTypeFLOWACTION captures enum value "FLOWACTION"
	DependencyObjectTypeFLOWACTION string = "FLOWACTION"

	// DependencyObjectTypeFLOWDATATYPE captures enum value "FLOWDATATYPE"
	DependencyObjectTypeFLOWDATATYPE string = "FLOWDATATYPE"

	// DependencyObjectTypeFLOWMILESTONE captures enum value "FLOWMILESTONE"
	DependencyObjectTypeFLOWMILESTONE string = "FLOWMILESTONE"

	// DependencyObjectTypeFLOWOUTCOME captures enum value "FLOWOUTCOME"
	DependencyObjectTypeFLOWOUTCOME string = "FLOWOUTCOME"

	// DependencyObjectTypeGRAMMAR captures enum value "GRAMMAR"
	DependencyObjectTypeGRAMMAR string = "GRAMMAR"

	// DependencyObjectTypeGROUP captures enum value "GROUP"
	DependencyObjectTypeGROUP string = "GROUP"

	// DependencyObjectTypeIMAGE captures enum value "IMAGE"
	DependencyObjectTypeIMAGE string = "IMAGE"

	// DependencyObjectTypeINBOUNDCALLFLOW captures enum value "INBOUNDCALLFLOW"
	DependencyObjectTypeINBOUNDCALLFLOW string = "INBOUNDCALLFLOW"

	// DependencyObjectTypeINBOUNDCHATFLOW captures enum value "INBOUNDCHATFLOW"
	DependencyObjectTypeINBOUNDCHATFLOW string = "INBOUNDCHATFLOW"

	// DependencyObjectTypeINBOUNDEMAILFLOW captures enum value "INBOUNDEMAILFLOW"
	DependencyObjectTypeINBOUNDEMAILFLOW string = "INBOUNDEMAILFLOW"

	// DependencyObjectTypeINBOUNDSHORTMESSAGEFLOW captures enum value "INBOUNDSHORTMESSAGEFLOW"
	DependencyObjectTypeINBOUNDSHORTMESSAGEFLOW string = "INBOUNDSHORTMESSAGEFLOW"

	// DependencyObjectTypeINQUEUECALLFLOW captures enum value "INQUEUECALLFLOW"
	DependencyObjectTypeINQUEUECALLFLOW string = "INQUEUECALLFLOW"

	// DependencyObjectTypeINQUEUEEMAILFLOW captures enum value "INQUEUEEMAILFLOW"
	DependencyObjectTypeINQUEUEEMAILFLOW string = "INQUEUEEMAILFLOW"

	// DependencyObjectTypeINQUEUESHORTMESSAGEFLOW captures enum value "INQUEUESHORTMESSAGEFLOW"
	DependencyObjectTypeINQUEUESHORTMESSAGEFLOW string = "INQUEUESHORTMESSAGEFLOW"

	// DependencyObjectTypeIVRCONFIGURATION captures enum value "IVRCONFIGURATION"
	DependencyObjectTypeIVRCONFIGURATION string = "IVRCONFIGURATION"

	// DependencyObjectTypeKNOWLEDGEBASE captures enum value "KNOWLEDGEBASE"
	DependencyObjectTypeKNOWLEDGEBASE string = "KNOWLEDGEBASE"

	// DependencyObjectTypeKNOWLEDGEBASEDOCUMENT captures enum value "KNOWLEDGEBASEDOCUMENT"
	DependencyObjectTypeKNOWLEDGEBASEDOCUMENT string = "KNOWLEDGEBASEDOCUMENT"

	// DependencyObjectTypeLANGUAGE captures enum value "LANGUAGE"
	DependencyObjectTypeLANGUAGE string = "LANGUAGE"

	// DependencyObjectTypeLEXBOT captures enum value "LEXBOT"
	DependencyObjectTypeLEXBOT string = "LEXBOT"

	// DependencyObjectTypeLEXBOTALIAS captures enum value "LEXBOTALIAS"
	DependencyObjectTypeLEXBOTALIAS string = "LEXBOTALIAS"

	// DependencyObjectTypeLEXV2BOT captures enum value "LEXV2BOT"
	DependencyObjectTypeLEXV2BOT string = "LEXV2BOT"

	// DependencyObjectTypeLEXV2BOTALIAS captures enum value "LEXV2BOTALIAS"
	DependencyObjectTypeLEXV2BOTALIAS string = "LEXV2BOTALIAS"

	// DependencyObjectTypeNLUDOMAIN captures enum value "NLUDOMAIN"
	DependencyObjectTypeNLUDOMAIN string = "NLUDOMAIN"

	// DependencyObjectTypeNUANCEMIXBOT captures enum value "NUANCEMIXBOT"
	DependencyObjectTypeNUANCEMIXBOT string = "NUANCEMIXBOT"

	// DependencyObjectTypeNUANCEMIXINTEGRATION captures enum value "NUANCEMIXINTEGRATION"
	DependencyObjectTypeNUANCEMIXINTEGRATION string = "NUANCEMIXINTEGRATION"

	// DependencyObjectTypeOAUTHCLIENT captures enum value "OAUTHCLIENT"
	DependencyObjectTypeOAUTHCLIENT string = "OAUTHCLIENT"

	// DependencyObjectTypeOUTBOUNDCALLFLOW captures enum value "OUTBOUNDCALLFLOW"
	DependencyObjectTypeOUTBOUNDCALLFLOW string = "OUTBOUNDCALLFLOW"

	// DependencyObjectTypeQUEUE captures enum value "QUEUE"
	DependencyObjectTypeQUEUE string = "QUEUE"

	// DependencyObjectTypeRECORDINGPOLICY captures enum value "RECORDINGPOLICY"
	DependencyObjectTypeRECORDINGPOLICY string = "RECORDINGPOLICY"

	// DependencyObjectTypeRESPONSE captures enum value "RESPONSE"
	DependencyObjectTypeRESPONSE string = "RESPONSE"

	// DependencyObjectTypeSCHEDULE captures enum value "SCHEDULE"
	DependencyObjectTypeSCHEDULE string = "SCHEDULE"

	// DependencyObjectTypeSCHEDULEGROUP captures enum value "SCHEDULEGROUP"
	DependencyObjectTypeSCHEDULEGROUP string = "SCHEDULEGROUP"

	// DependencyObjectTypeSECUREACTION captures enum value "SECUREACTION"
	DependencyObjectTypeSECUREACTION string = "SECUREACTION"

	// DependencyObjectTypeSECURECALLFLOW captures enum value "SECURECALLFLOW"
	DependencyObjectTypeSECURECALLFLOW string = "SECURECALLFLOW"

	// DependencyObjectTypeSTTENGINE captures enum value "STTENGINE"
	DependencyObjectTypeSTTENGINE string = "STTENGINE"

	// DependencyObjectTypeSURVEYINVITEFLOW captures enum value "SURVEYINVITEFLOW"
	DependencyObjectTypeSURVEYINVITEFLOW string = "SURVEYINVITEFLOW"

	// DependencyObjectTypeSYSTEMPROMPT captures enum value "SYSTEMPROMPT"
	DependencyObjectTypeSYSTEMPROMPT string = "SYSTEMPROMPT"

	// DependencyObjectTypeTTSENGINE captures enum value "TTSENGINE"
	DependencyObjectTypeTTSENGINE string = "TTSENGINE"

	// DependencyObjectTypeTTSVOICE captures enum value "TTSVOICE"
	DependencyObjectTypeTTSVOICE string = "TTSVOICE"

	// DependencyObjectTypeUSER captures enum value "USER"
	DependencyObjectTypeUSER string = "USER"

	// DependencyObjectTypeUSERPROMPT captures enum value "USERPROMPT"
	DependencyObjectTypeUSERPROMPT string = "USERPROMPT"

	// DependencyObjectTypeVOICEFLOW captures enum value "VOICEFLOW"
	DependencyObjectTypeVOICEFLOW string = "VOICEFLOW"

	// DependencyObjectTypeVOICEMAILFLOW captures enum value "VOICEMAILFLOW"
	DependencyObjectTypeVOICEMAILFLOW string = "VOICEMAILFLOW"

	// DependencyObjectTypeWIDGET captures enum value "WIDGET"
	DependencyObjectTypeWIDGET string = "WIDGET"

	// DependencyObjectTypeWORKFLOW captures enum value "WORKFLOW"
	DependencyObjectTypeWORKFLOW string = "WORKFLOW"

	// DependencyObjectTypeWORKITEMFLOW captures enum value "WORKITEMFLOW"
	DependencyObjectTypeWORKITEMFLOW string = "WORKITEMFLOW"
)

// prop value enum
func (m *DependencyObject) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dependencyObjectTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DependencyObject) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dependency object based on the context it is used
func (m *DependencyObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsumedResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsumingResources(ctx, formats); err != nil {
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

func (m *DependencyObject) contextValidateConsumedResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsumedResources); i++ {

		if m.ConsumedResources[i] != nil {
			if err := m.ConsumedResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumedResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumedResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DependencyObject) contextValidateConsumingResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsumingResources); i++ {

		if m.ConsumingResources[i] != nil {
			if err := m.ConsumingResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consumingResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("consumingResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DependencyObject) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DependencyObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DependencyObject) UnmarshalBinary(b []byte) error {
	var res DependencyObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
