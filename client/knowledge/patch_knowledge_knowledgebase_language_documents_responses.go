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

// PatchKnowledgeKnowledgebaseLanguageDocumentsReader is a Reader for the PatchKnowledgeKnowledgebaseLanguageDocuments structure.
type PatchKnowledgeKnowledgebaseLanguageDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsOK creates a PatchKnowledgeKnowledgebaseLanguageDocumentsOK with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsOK() *PatchKnowledgeKnowledgebaseLanguageDocumentsOK {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsOK{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsOK struct {
	Payload *models.DocumentListing
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents o k response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents o k response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents o k response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language documents o k response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents o k response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) GetPayload() *models.DocumentListing {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest creates a PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest() *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents bad request response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents bad request response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents bad request response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents bad request response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents bad request response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized creates a PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized() *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents unauthorized response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents unauthorized response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents unauthorized response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents unauthorized response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents unauthorized response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsForbidden creates a PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsForbidden() *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents forbidden response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents forbidden response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents forbidden response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents forbidden response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents forbidden response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsNotFound creates a PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsNotFound() *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents not found response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents not found response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents not found response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents not found response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents not found response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout creates a PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout() *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents request timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents request timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents request timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents request timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents request timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents request entity too large response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents request entity too large response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents request entity too large response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents request entity too large response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents request entity too large response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType() *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents unsupported media type response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents unsupported media type response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents unsupported media type response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents unsupported media type response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents unsupported media type response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests creates a PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests() *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents too many requests response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents too many requests response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents too many requests response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language documents too many requests response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language documents too many requests response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError creates a PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError() *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents internal server error response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents internal server error response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents internal server error response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language documents internal server error response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language documents internal server error response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable creates a PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable() *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents service unavailable response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents service unavailable response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents service unavailable response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language documents service unavailable response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language documents service unavailable response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout creates a PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout() *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language documents gateway timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language documents gateway timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language documents gateway timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language documents gateway timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language documents gateway timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
