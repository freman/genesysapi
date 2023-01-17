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

// PatchKnowledgeKnowledgebaseLabelReader is a Reader for the PatchKnowledgeKnowledgebaseLabel structure.
type PatchKnowledgeKnowledgebaseLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseLabelRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchKnowledgeKnowledgebaseLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseLabelUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseLabelTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseLabelServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseLabelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseLabelOK creates a PatchKnowledgeKnowledgebaseLabelOK with default headers values
func NewPatchKnowledgeKnowledgebaseLabelOK() *PatchKnowledgeKnowledgebaseLabelOK {
	return &PatchKnowledgeKnowledgebaseLabelOK{}
}

/*
PatchKnowledgeKnowledgebaseLabelOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseLabelOK struct {
	Payload *models.LabelResponse
}

// IsSuccess returns true when this patch knowledge knowledgebase label o k response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase label o k response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label o k response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase label o k response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label o k response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchKnowledgeKnowledgebaseLabelOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelOK) GetPayload() *models.LabelResponse {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelBadRequest creates a PatchKnowledgeKnowledgebaseLabelBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseLabelBadRequest() *PatchKnowledgeKnowledgebaseLabelBadRequest {
	return &PatchKnowledgeKnowledgebaseLabelBadRequest{}
}

/*
PatchKnowledgeKnowledgebaseLabelBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseLabelBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label bad request response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label bad request response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label bad request response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label bad request response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label bad request response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelUnauthorized creates a PatchKnowledgeKnowledgebaseLabelUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseLabelUnauthorized() *PatchKnowledgeKnowledgebaseLabelUnauthorized {
	return &PatchKnowledgeKnowledgebaseLabelUnauthorized{}
}

/*
PatchKnowledgeKnowledgebaseLabelUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseLabelUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label unauthorized response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label unauthorized response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label unauthorized response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label unauthorized response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label unauthorized response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelForbidden creates a PatchKnowledgeKnowledgebaseLabelForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseLabelForbidden() *PatchKnowledgeKnowledgebaseLabelForbidden {
	return &PatchKnowledgeKnowledgebaseLabelForbidden{}
}

/*
PatchKnowledgeKnowledgebaseLabelForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseLabelForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label forbidden response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label forbidden response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label forbidden response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label forbidden response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label forbidden response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeKnowledgebaseLabelForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelNotFound creates a PatchKnowledgeKnowledgebaseLabelNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseLabelNotFound() *PatchKnowledgeKnowledgebaseLabelNotFound {
	return &PatchKnowledgeKnowledgebaseLabelNotFound{}
}

/*
PatchKnowledgeKnowledgebaseLabelNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseLabelNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label not found response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label not found response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label not found response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label not found response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label not found response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeKnowledgebaseLabelNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelRequestTimeout creates a PatchKnowledgeKnowledgebaseLabelRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLabelRequestTimeout() *PatchKnowledgeKnowledgebaseLabelRequestTimeout {
	return &PatchKnowledgeKnowledgebaseLabelRequestTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLabelRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseLabelRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label request timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label request timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label request timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label request timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label request timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelConflict creates a PatchKnowledgeKnowledgebaseLabelConflict with default headers values
func NewPatchKnowledgeKnowledgebaseLabelConflict() *PatchKnowledgeKnowledgebaseLabelConflict {
	return &PatchKnowledgeKnowledgebaseLabelConflict{}
}

/*
PatchKnowledgeKnowledgebaseLabelConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchKnowledgeKnowledgebaseLabelConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label conflict response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label conflict response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label conflict response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label conflict response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label conflict response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PatchKnowledgeKnowledgebaseLabelConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelConflict  %+v", 409, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelConflict) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelConflict  %+v", 409, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge{}
}

/*
PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label request entity too large response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label request entity too large response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label request entity too large response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label request entity too large response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label request entity too large response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseLabelUnsupportedMediaType() *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType{}
}

/*
PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label unsupported media type response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label unsupported media type response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label unsupported media type response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label unsupported media type response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label unsupported media type response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelTooManyRequests creates a PatchKnowledgeKnowledgebaseLabelTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseLabelTooManyRequests() *PatchKnowledgeKnowledgebaseLabelTooManyRequests {
	return &PatchKnowledgeKnowledgebaseLabelTooManyRequests{}
}

/*
PatchKnowledgeKnowledgebaseLabelTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseLabelTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label too many requests response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label too many requests response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label too many requests response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase label too many requests response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase label too many requests response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelInternalServerError creates a PatchKnowledgeKnowledgebaseLabelInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseLabelInternalServerError() *PatchKnowledgeKnowledgebaseLabelInternalServerError {
	return &PatchKnowledgeKnowledgebaseLabelInternalServerError{}
}

/*
PatchKnowledgeKnowledgebaseLabelInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseLabelInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label internal server error response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label internal server error response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label internal server error response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase label internal server error response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase label internal server error response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelServiceUnavailable creates a PatchKnowledgeKnowledgebaseLabelServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseLabelServiceUnavailable() *PatchKnowledgeKnowledgebaseLabelServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseLabelServiceUnavailable{}
}

/*
PatchKnowledgeKnowledgebaseLabelServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseLabelServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label service unavailable response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label service unavailable response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label service unavailable response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase label service unavailable response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase label service unavailable response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLabelGatewayTimeout creates a PatchKnowledgeKnowledgebaseLabelGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLabelGatewayTimeout() *PatchKnowledgeKnowledgebaseLabelGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseLabelGatewayTimeout{}
}

/*
PatchKnowledgeKnowledgebaseLabelGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseLabelGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase label gateway timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase label gateway timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase label gateway timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase label gateway timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase label gateway timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] patchKnowledgeKnowledgebaseLabelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLabelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
