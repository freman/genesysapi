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

// Dependency dependency
//
// swagger:model Dependency
type Dependency struct {

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

// Validate validates this dependency
func (m *Dependency) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *Dependency) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var dependencyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACDLANGUAGE","ACDSKILL","ACDWRAPUPCODE","BOTCONNECTORBOT","BOTCONNECTORINTEGRATION","BOTFLOW","BRIDGEACTION","COMMONMODULEFLOW","COMPOSERSCRIPT","CONTACTLIST","DATAACTION","DATATABLE","DIALOGENGINEBOT","DIALOGENGINEBOTVERSION","DIALOGFLOWAGENT","DIALOGFLOWCXAGENT","DIGITALBOTFLOW","EMAILROUTE","EMERGENCYGROUP","FLOWACTION","FLOWDATATYPE","FLOWMILESTONE","FLOWOUTCOME","GRAMMAR","GROUP","IMAGE","INBOUNDCALLFLOW","INBOUNDCHATFLOW","INBOUNDEMAILFLOW","INBOUNDSHORTMESSAGEFLOW","INQUEUECALLFLOW","INQUEUEEMAILFLOW","INQUEUESHORTMESSAGEFLOW","IVRCONFIGURATION","KNOWLEDGEBASE","KNOWLEDGEBASEDOCUMENT","LANGUAGE","LEXBOT","LEXBOTALIAS","LEXV2BOT","LEXV2BOTALIAS","NLUDOMAIN","NUANCEMIXBOT","NUANCEMIXINTEGRATION","OAUTHCLIENT","OUTBOUNDCALLFLOW","QUEUE","RECORDINGPOLICY","RESPONSE","SCHEDULE","SCHEDULEGROUP","SECUREACTION","SECURECALLFLOW","STTENGINE","SURVEYINVITEFLOW","SYSTEMPROMPT","TTSENGINE","TTSVOICE","USER","USERPROMPT","VOICEFLOW","VOICEMAILFLOW","WIDGET","WORKFLOW","WORKITEMFLOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dependencyTypeTypePropEnum = append(dependencyTypeTypePropEnum, v)
	}
}

const (

	// DependencyTypeACDLANGUAGE captures enum value "ACDLANGUAGE"
	DependencyTypeACDLANGUAGE string = "ACDLANGUAGE"

	// DependencyTypeACDSKILL captures enum value "ACDSKILL"
	DependencyTypeACDSKILL string = "ACDSKILL"

	// DependencyTypeACDWRAPUPCODE captures enum value "ACDWRAPUPCODE"
	DependencyTypeACDWRAPUPCODE string = "ACDWRAPUPCODE"

	// DependencyTypeBOTCONNECTORBOT captures enum value "BOTCONNECTORBOT"
	DependencyTypeBOTCONNECTORBOT string = "BOTCONNECTORBOT"

	// DependencyTypeBOTCONNECTORINTEGRATION captures enum value "BOTCONNECTORINTEGRATION"
	DependencyTypeBOTCONNECTORINTEGRATION string = "BOTCONNECTORINTEGRATION"

	// DependencyTypeBOTFLOW captures enum value "BOTFLOW"
	DependencyTypeBOTFLOW string = "BOTFLOW"

	// DependencyTypeBRIDGEACTION captures enum value "BRIDGEACTION"
	DependencyTypeBRIDGEACTION string = "BRIDGEACTION"

	// DependencyTypeCOMMONMODULEFLOW captures enum value "COMMONMODULEFLOW"
	DependencyTypeCOMMONMODULEFLOW string = "COMMONMODULEFLOW"

	// DependencyTypeCOMPOSERSCRIPT captures enum value "COMPOSERSCRIPT"
	DependencyTypeCOMPOSERSCRIPT string = "COMPOSERSCRIPT"

	// DependencyTypeCONTACTLIST captures enum value "CONTACTLIST"
	DependencyTypeCONTACTLIST string = "CONTACTLIST"

	// DependencyTypeDATAACTION captures enum value "DATAACTION"
	DependencyTypeDATAACTION string = "DATAACTION"

	// DependencyTypeDATATABLE captures enum value "DATATABLE"
	DependencyTypeDATATABLE string = "DATATABLE"

	// DependencyTypeDIALOGENGINEBOT captures enum value "DIALOGENGINEBOT"
	DependencyTypeDIALOGENGINEBOT string = "DIALOGENGINEBOT"

	// DependencyTypeDIALOGENGINEBOTVERSION captures enum value "DIALOGENGINEBOTVERSION"
	DependencyTypeDIALOGENGINEBOTVERSION string = "DIALOGENGINEBOTVERSION"

	// DependencyTypeDIALOGFLOWAGENT captures enum value "DIALOGFLOWAGENT"
	DependencyTypeDIALOGFLOWAGENT string = "DIALOGFLOWAGENT"

	// DependencyTypeDIALOGFLOWCXAGENT captures enum value "DIALOGFLOWCXAGENT"
	DependencyTypeDIALOGFLOWCXAGENT string = "DIALOGFLOWCXAGENT"

	// DependencyTypeDIGITALBOTFLOW captures enum value "DIGITALBOTFLOW"
	DependencyTypeDIGITALBOTFLOW string = "DIGITALBOTFLOW"

	// DependencyTypeEMAILROUTE captures enum value "EMAILROUTE"
	DependencyTypeEMAILROUTE string = "EMAILROUTE"

	// DependencyTypeEMERGENCYGROUP captures enum value "EMERGENCYGROUP"
	DependencyTypeEMERGENCYGROUP string = "EMERGENCYGROUP"

	// DependencyTypeFLOWACTION captures enum value "FLOWACTION"
	DependencyTypeFLOWACTION string = "FLOWACTION"

	// DependencyTypeFLOWDATATYPE captures enum value "FLOWDATATYPE"
	DependencyTypeFLOWDATATYPE string = "FLOWDATATYPE"

	// DependencyTypeFLOWMILESTONE captures enum value "FLOWMILESTONE"
	DependencyTypeFLOWMILESTONE string = "FLOWMILESTONE"

	// DependencyTypeFLOWOUTCOME captures enum value "FLOWOUTCOME"
	DependencyTypeFLOWOUTCOME string = "FLOWOUTCOME"

	// DependencyTypeGRAMMAR captures enum value "GRAMMAR"
	DependencyTypeGRAMMAR string = "GRAMMAR"

	// DependencyTypeGROUP captures enum value "GROUP"
	DependencyTypeGROUP string = "GROUP"

	// DependencyTypeIMAGE captures enum value "IMAGE"
	DependencyTypeIMAGE string = "IMAGE"

	// DependencyTypeINBOUNDCALLFLOW captures enum value "INBOUNDCALLFLOW"
	DependencyTypeINBOUNDCALLFLOW string = "INBOUNDCALLFLOW"

	// DependencyTypeINBOUNDCHATFLOW captures enum value "INBOUNDCHATFLOW"
	DependencyTypeINBOUNDCHATFLOW string = "INBOUNDCHATFLOW"

	// DependencyTypeINBOUNDEMAILFLOW captures enum value "INBOUNDEMAILFLOW"
	DependencyTypeINBOUNDEMAILFLOW string = "INBOUNDEMAILFLOW"

	// DependencyTypeINBOUNDSHORTMESSAGEFLOW captures enum value "INBOUNDSHORTMESSAGEFLOW"
	DependencyTypeINBOUNDSHORTMESSAGEFLOW string = "INBOUNDSHORTMESSAGEFLOW"

	// DependencyTypeINQUEUECALLFLOW captures enum value "INQUEUECALLFLOW"
	DependencyTypeINQUEUECALLFLOW string = "INQUEUECALLFLOW"

	// DependencyTypeINQUEUEEMAILFLOW captures enum value "INQUEUEEMAILFLOW"
	DependencyTypeINQUEUEEMAILFLOW string = "INQUEUEEMAILFLOW"

	// DependencyTypeINQUEUESHORTMESSAGEFLOW captures enum value "INQUEUESHORTMESSAGEFLOW"
	DependencyTypeINQUEUESHORTMESSAGEFLOW string = "INQUEUESHORTMESSAGEFLOW"

	// DependencyTypeIVRCONFIGURATION captures enum value "IVRCONFIGURATION"
	DependencyTypeIVRCONFIGURATION string = "IVRCONFIGURATION"

	// DependencyTypeKNOWLEDGEBASE captures enum value "KNOWLEDGEBASE"
	DependencyTypeKNOWLEDGEBASE string = "KNOWLEDGEBASE"

	// DependencyTypeKNOWLEDGEBASEDOCUMENT captures enum value "KNOWLEDGEBASEDOCUMENT"
	DependencyTypeKNOWLEDGEBASEDOCUMENT string = "KNOWLEDGEBASEDOCUMENT"

	// DependencyTypeLANGUAGE captures enum value "LANGUAGE"
	DependencyTypeLANGUAGE string = "LANGUAGE"

	// DependencyTypeLEXBOT captures enum value "LEXBOT"
	DependencyTypeLEXBOT string = "LEXBOT"

	// DependencyTypeLEXBOTALIAS captures enum value "LEXBOTALIAS"
	DependencyTypeLEXBOTALIAS string = "LEXBOTALIAS"

	// DependencyTypeLEXV2BOT captures enum value "LEXV2BOT"
	DependencyTypeLEXV2BOT string = "LEXV2BOT"

	// DependencyTypeLEXV2BOTALIAS captures enum value "LEXV2BOTALIAS"
	DependencyTypeLEXV2BOTALIAS string = "LEXV2BOTALIAS"

	// DependencyTypeNLUDOMAIN captures enum value "NLUDOMAIN"
	DependencyTypeNLUDOMAIN string = "NLUDOMAIN"

	// DependencyTypeNUANCEMIXBOT captures enum value "NUANCEMIXBOT"
	DependencyTypeNUANCEMIXBOT string = "NUANCEMIXBOT"

	// DependencyTypeNUANCEMIXINTEGRATION captures enum value "NUANCEMIXINTEGRATION"
	DependencyTypeNUANCEMIXINTEGRATION string = "NUANCEMIXINTEGRATION"

	// DependencyTypeOAUTHCLIENT captures enum value "OAUTHCLIENT"
	DependencyTypeOAUTHCLIENT string = "OAUTHCLIENT"

	// DependencyTypeOUTBOUNDCALLFLOW captures enum value "OUTBOUNDCALLFLOW"
	DependencyTypeOUTBOUNDCALLFLOW string = "OUTBOUNDCALLFLOW"

	// DependencyTypeQUEUE captures enum value "QUEUE"
	DependencyTypeQUEUE string = "QUEUE"

	// DependencyTypeRECORDINGPOLICY captures enum value "RECORDINGPOLICY"
	DependencyTypeRECORDINGPOLICY string = "RECORDINGPOLICY"

	// DependencyTypeRESPONSE captures enum value "RESPONSE"
	DependencyTypeRESPONSE string = "RESPONSE"

	// DependencyTypeSCHEDULE captures enum value "SCHEDULE"
	DependencyTypeSCHEDULE string = "SCHEDULE"

	// DependencyTypeSCHEDULEGROUP captures enum value "SCHEDULEGROUP"
	DependencyTypeSCHEDULEGROUP string = "SCHEDULEGROUP"

	// DependencyTypeSECUREACTION captures enum value "SECUREACTION"
	DependencyTypeSECUREACTION string = "SECUREACTION"

	// DependencyTypeSECURECALLFLOW captures enum value "SECURECALLFLOW"
	DependencyTypeSECURECALLFLOW string = "SECURECALLFLOW"

	// DependencyTypeSTTENGINE captures enum value "STTENGINE"
	DependencyTypeSTTENGINE string = "STTENGINE"

	// DependencyTypeSURVEYINVITEFLOW captures enum value "SURVEYINVITEFLOW"
	DependencyTypeSURVEYINVITEFLOW string = "SURVEYINVITEFLOW"

	// DependencyTypeSYSTEMPROMPT captures enum value "SYSTEMPROMPT"
	DependencyTypeSYSTEMPROMPT string = "SYSTEMPROMPT"

	// DependencyTypeTTSENGINE captures enum value "TTSENGINE"
	DependencyTypeTTSENGINE string = "TTSENGINE"

	// DependencyTypeTTSVOICE captures enum value "TTSVOICE"
	DependencyTypeTTSVOICE string = "TTSVOICE"

	// DependencyTypeUSER captures enum value "USER"
	DependencyTypeUSER string = "USER"

	// DependencyTypeUSERPROMPT captures enum value "USERPROMPT"
	DependencyTypeUSERPROMPT string = "USERPROMPT"

	// DependencyTypeVOICEFLOW captures enum value "VOICEFLOW"
	DependencyTypeVOICEFLOW string = "VOICEFLOW"

	// DependencyTypeVOICEMAILFLOW captures enum value "VOICEMAILFLOW"
	DependencyTypeVOICEMAILFLOW string = "VOICEMAILFLOW"

	// DependencyTypeWIDGET captures enum value "WIDGET"
	DependencyTypeWIDGET string = "WIDGET"

	// DependencyTypeWORKFLOW captures enum value "WORKFLOW"
	DependencyTypeWORKFLOW string = "WORKFLOW"

	// DependencyTypeWORKITEMFLOW captures enum value "WORKITEMFLOW"
	DependencyTypeWORKITEMFLOW string = "WORKITEMFLOW"
)

// prop value enum
func (m *Dependency) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dependencyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Dependency) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dependency based on the context it is used
func (m *Dependency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dependency) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Dependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dependency) UnmarshalBinary(b []byte) error {
	var res Dependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
