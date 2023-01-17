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

// Document document
//
// swagger:model Document
type Document struct {

	// A list of permitted action rights for the user making the request
	ACL []string `json:"acl"`

	// attributes
	Attributes []*DocumentAttribute `json:"attributes"`

	// caller address
	CallerAddress string `json:"callerAddress,omitempty"`

	// change number
	ChangeNumber int32 `json:"changeNumber,omitempty"`

	// content length
	ContentLength int64 `json:"contentLength,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// content Uri
	// Format: uri
	ContentURI strfmt.URI `json:"contentUri,omitempty"`

	// created by
	CreatedBy *DomainEntityRef `json:"createdBy,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateUploaded strfmt.DateTime `json:"dateUploaded,omitempty"`

	// download sharing Uri
	// Format: uri
	DownloadSharingURI strfmt.URI `json:"downloadSharingUri,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// lock info
	LockInfo *LockInfo `json:"lockInfo,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// page count
	PageCount int64 `json:"pageCount,omitempty"`

	// read
	Read bool `json:"read"`

	// receiver address
	ReceiverAddress string `json:"receiverAddress,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// sharing status
	// Enum: [NONE LIMITED PUBLIC]
	SharingStatus string `json:"sharingStatus,omitempty"`

	// sharing Uri
	// Format: uri
	SharingURI strfmt.URI `json:"sharingUri,omitempty"`

	// system type
	// Enum: [DOCUMENT FAX RECORDING]
	SystemType string `json:"systemType,omitempty"`

	// tag values
	TagValues []*TagValue `json:"tagValues"`

	// tags
	Tags []string `json:"tags"`

	// thumbnails
	Thumbnails []*DocumentThumbnail `json:"thumbnails"`

	// upload destination Uri
	// Format: uri
	UploadDestinationURI strfmt.URI `json:"uploadDestinationUri,omitempty"`

	// upload method
	// Enum: [SINGLE_PUT MULTIPART_POST]
	UploadMethod string `json:"uploadMethod,omitempty"`

	// upload status
	UploadStatus *DomainEntityRef `json:"uploadStatus,omitempty"`

	// uploaded by
	UploadedBy *DomainEntityRef `json:"uploadedBy,omitempty"`

	// workspace
	Workspace *DomainEntityRef `json:"workspace,omitempty"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateUploaded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadSharingURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharingURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadDestinationURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) validateContentURI(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentURI) { // not required
		return nil
	}

	if err := validate.FormatOf("contentUri", "body", "uri", m.ContentURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *Document) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateDateUploaded(formats strfmt.Registry) error {
	if swag.IsZero(m.DateUploaded) { // not required
		return nil
	}

	if err := validate.FormatOf("dateUploaded", "body", "date-time", m.DateUploaded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateDownloadSharingURI(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadSharingURI) { // not required
		return nil
	}

	if err := validate.FormatOf("downloadSharingUri", "body", "uri", m.DownloadSharingURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateLockInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.LockInfo) { // not required
		return nil
	}

	if m.LockInfo != nil {
		if err := m.LockInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lockInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lockInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Document) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var documentTypeSharingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","LIMITED","PUBLIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeSharingStatusPropEnum = append(documentTypeSharingStatusPropEnum, v)
	}
}

const (

	// DocumentSharingStatusNONE captures enum value "NONE"
	DocumentSharingStatusNONE string = "NONE"

	// DocumentSharingStatusLIMITED captures enum value "LIMITED"
	DocumentSharingStatusLIMITED string = "LIMITED"

	// DocumentSharingStatusPUBLIC captures enum value "PUBLIC"
	DocumentSharingStatusPUBLIC string = "PUBLIC"
)

// prop value enum
func (m *Document) validateSharingStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentTypeSharingStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateSharingStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SharingStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSharingStatusEnum("sharingStatus", "body", m.SharingStatus); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateSharingURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SharingURI) { // not required
		return nil
	}

	if err := validate.FormatOf("sharingUri", "body", "uri", m.SharingURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var documentTypeSystemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DOCUMENT","FAX","RECORDING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeSystemTypePropEnum = append(documentTypeSystemTypePropEnum, v)
	}
}

const (

	// DocumentSystemTypeDOCUMENT captures enum value "DOCUMENT"
	DocumentSystemTypeDOCUMENT string = "DOCUMENT"

	// DocumentSystemTypeFAX captures enum value "FAX"
	DocumentSystemTypeFAX string = "FAX"

	// DocumentSystemTypeRECORDING captures enum value "RECORDING"
	DocumentSystemTypeRECORDING string = "RECORDING"
)

// prop value enum
func (m *Document) validateSystemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentTypeSystemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateSystemType(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSystemTypeEnum("systemType", "body", m.SystemType); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateTagValues(formats strfmt.Registry) error {
	if swag.IsZero(m.TagValues) { // not required
		return nil
	}

	for i := 0; i < len(m.TagValues); i++ {
		if swag.IsZero(m.TagValues[i]) { // not required
			continue
		}

		if m.TagValues[i] != nil {
			if err := m.TagValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) validateThumbnails(formats strfmt.Registry) error {
	if swag.IsZero(m.Thumbnails) { // not required
		return nil
	}

	for i := 0; i < len(m.Thumbnails); i++ {
		if swag.IsZero(m.Thumbnails[i]) { // not required
			continue
		}

		if m.Thumbnails[i] != nil {
			if err := m.Thumbnails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("thumbnails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("thumbnails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) validateUploadDestinationURI(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadDestinationURI) { // not required
		return nil
	}

	if err := validate.FormatOf("uploadDestinationUri", "body", "uri", m.UploadDestinationURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var documentTypeUploadMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE_PUT","MULTIPART_POST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentTypeUploadMethodPropEnum = append(documentTypeUploadMethodPropEnum, v)
	}
}

const (

	// DocumentUploadMethodSINGLEPUT captures enum value "SINGLE_PUT"
	DocumentUploadMethodSINGLEPUT string = "SINGLE_PUT"

	// DocumentUploadMethodMULTIPARTPOST captures enum value "MULTIPART_POST"
	DocumentUploadMethodMULTIPARTPOST string = "MULTIPART_POST"
)

// prop value enum
func (m *Document) validateUploadMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentTypeUploadMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Document) validateUploadMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateUploadMethodEnum("uploadMethod", "body", m.UploadMethod); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateUploadStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadStatus) { // not required
		return nil
	}

	if m.UploadStatus != nil {
		if err := m.UploadStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Document) validateUploadedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadedBy) { // not required
		return nil
	}

	if m.UploadedBy != nil {
		if err := m.UploadedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Document) validateWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this document based on the context it is used
func (m *Document) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLockInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThumbnails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {
		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *Document) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Document) contextValidateLockInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.LockInfo != nil {
		if err := m.LockInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lockInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lockInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Document) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Document) contextValidateTagValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagValues); i++ {

		if m.TagValues[i] != nil {
			if err := m.TagValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) contextValidateThumbnails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Thumbnails); i++ {

		if m.Thumbnails[i] != nil {
			if err := m.Thumbnails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("thumbnails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("thumbnails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) contextValidateUploadStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadStatus != nil {
		if err := m.UploadStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadStatus")
			}
			return err
		}
	}

	return nil
}

func (m *Document) contextValidateUploadedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadedBy != nil {
		if err := m.UploadedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Document) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspace != nil {
		if err := m.Workspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
