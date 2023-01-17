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

// PatchKnowledgeKnowledgebaseLanguageCategoryReader is a Reader for the PatchKnowledgeKnowledgebaseLanguageCategory structure.
type PatchKnowledgeKnowledgebaseLanguageCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryOK creates a PatchKnowledgeKnowledgebaseLanguageCategoryOK with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryOK() *PatchKnowledgeKnowledgebaseLanguageCategoryOK {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryOK{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryOK struct {
	Payload *models.KnowledgeExtendedCategory
}

// IsSuccess returns true when this patch knowledge knowledgebase language category o k response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase language category o k response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category o k response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language category o k response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category o k response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) GetPayload() *models.KnowledgeExtendedCategory {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeExtendedCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryBadRequest creates a PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryBadRequest() *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category bad request response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category bad request response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category bad request response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category bad request response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category bad request response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized creates a PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized() *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category unauthorized response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category unauthorized response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category unauthorized response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category unauthorized response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category unauthorized response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryForbidden creates a PatchKnowledgeKnowledgebaseLanguageCategoryForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryForbidden() *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryForbidden{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category forbidden response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category forbidden response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category forbidden response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category forbidden response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category forbidden response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryNotFound creates a PatchKnowledgeKnowledgebaseLanguageCategoryNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryNotFound() *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryNotFound{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category not found response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category not found response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category not found response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category not found response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category not found response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout creates a PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout() *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category request timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category request timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category request timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category request timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category request timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category request entity too large response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category request entity too large response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category request entity too large response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category request entity too large response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category request entity too large response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType() *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category unsupported media type response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category unsupported media type response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category unsupported media type response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category unsupported media type response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category unsupported media type response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests creates a PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests() *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category too many requests response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category too many requests response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category too many requests response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase language category too many requests response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase language category too many requests response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError creates a PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError() *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category internal server error response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category internal server error response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category internal server error response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language category internal server error response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language category internal server error response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable creates a PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable() *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category service unavailable response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category service unavailable response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category service unavailable response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language category service unavailable response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language category service unavailable response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout creates a PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout() *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase language category gateway timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase language category gateway timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase language category gateway timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase language category gateway timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase language category gateway timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
