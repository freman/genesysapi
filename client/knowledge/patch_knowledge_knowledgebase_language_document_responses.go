// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchKnowledgeKnowledgebaseLanguageDocumentReader is a Reader for the PatchKnowledgeKnowledgebaseLanguageDocument structure.
type PatchKnowledgeKnowledgebaseLanguageDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentOK creates a PatchKnowledgeKnowledgebaseLanguageDocumentOK with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentOK() *PatchKnowledgeKnowledgebaseLanguageDocumentOK {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentOK{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentOK struct {
	Payload *models.KnowledgeDocument
}

// IsSuccess returns true when this patch knowledge knowledgebase language document o k response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase language document o k response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document o k response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language document o k response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document o k response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) GetPayload() *models.KnowledgeDocument {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentBadRequest creates a PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentBadRequest() *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document bad request response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document bad request response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document bad request response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document bad request response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document bad request response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized creates a PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized() *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document unauthorized response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document unauthorized response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document unauthorized response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document unauthorized response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document unauthorized response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentForbidden creates a PatchKnowledgeKnowledgebaseLanguageDocumentForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentForbidden() *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentForbidden{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document forbidden response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document forbidden response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document forbidden response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document forbidden response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document forbidden response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentNotFound creates a PatchKnowledgeKnowledgebaseLanguageDocumentNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentNotFound() *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentNotFound{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document not found response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document not found response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document not found response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document not found response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document not found response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout creates a PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout() *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document request timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document request timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document request timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document request timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document request timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document request entity too large response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document request entity too large response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document request entity too large response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document request entity too large response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document request entity too large response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType() *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document unsupported media type response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document unsupported media type response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document unsupported media type response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document unsupported media type response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document unsupported media type response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests creates a PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests() *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document too many requests response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document too many requests response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document too many requests response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language document too many requests response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language document too many requests response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError creates a PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError() *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document internal server error response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document internal server error response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document internal server error response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language document internal server error response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language document internal server error response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable creates a PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable() *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document service unavailable response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document service unavailable response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document service unavailable response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language document service unavailable response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language document service unavailable response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout creates a PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout() *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language document gateway timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language document gateway timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language document gateway timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language document gateway timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language document gateway timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] patchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
