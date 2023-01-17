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

// PatchKnowledgeKnowledgebaseImportJobReader is a Reader for the PatchKnowledgeKnowledgebaseImportJob structure.
type PatchKnowledgeKnowledgebaseImportJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseImportJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseImportJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPatchKnowledgeKnowledgebaseImportJobAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseImportJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseImportJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseImportJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseImportJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseImportJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseImportJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseImportJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseImportJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseImportJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseImportJobOK creates a PatchKnowledgeKnowledgebaseImportJobOK with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobOK() *PatchKnowledgeKnowledgebaseImportJobOK {
	return &PatchKnowledgeKnowledgebaseImportJobOK{}
}

/*
PatchKnowledgeKnowledgebaseImportJobOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseImportJobOK struct {
	Payload *models.KnowledgeImportJobResponse
}

// IsSuccess returns true when this patch knowledge knowledgebase import job o k response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase import job o k response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job o k response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase import job o k response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job o k response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchKnowledgeKnowledgebaseImportJobOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobOK) GetPayload() *models.KnowledgeImportJobResponse {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeImportJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobAccepted creates a PatchKnowledgeKnowledgebaseImportJobAccepted with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobAccepted() *PatchKnowledgeKnowledgebaseImportJobAccepted {
	return &PatchKnowledgeKnowledgebaseImportJobAccepted{}
}

/*
PatchKnowledgeKnowledgebaseImportJobAccepted describes a response with status code 202, with default header values.

Import job in progress
*/
type PatchKnowledgeKnowledgebaseImportJobAccepted struct {
	Payload *models.KnowledgeImportJobResponse
}

// IsSuccess returns true when this patch knowledge knowledgebase import job accepted response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase import job accepted response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job accepted response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase import job accepted response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job accepted response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobAccepted  %+v", 202, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobAccepted  %+v", 202, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) GetPayload() *models.KnowledgeImportJobResponse {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeImportJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobBadRequest creates a PatchKnowledgeKnowledgebaseImportJobBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobBadRequest() *PatchKnowledgeKnowledgebaseImportJobBadRequest {
	return &PatchKnowledgeKnowledgebaseImportJobBadRequest{}
}

/*
PatchKnowledgeKnowledgebaseImportJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseImportJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job bad request response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job bad request response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job bad request response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job bad request response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job bad request response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobUnauthorized creates a PatchKnowledgeKnowledgebaseImportJobUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobUnauthorized() *PatchKnowledgeKnowledgebaseImportJobUnauthorized {
	return &PatchKnowledgeKnowledgebaseImportJobUnauthorized{}
}

/*
PatchKnowledgeKnowledgebaseImportJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseImportJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job unauthorized response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job unauthorized response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job unauthorized response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job unauthorized response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job unauthorized response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobForbidden creates a PatchKnowledgeKnowledgebaseImportJobForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobForbidden() *PatchKnowledgeKnowledgebaseImportJobForbidden {
	return &PatchKnowledgeKnowledgebaseImportJobForbidden{}
}

/*
PatchKnowledgeKnowledgebaseImportJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseImportJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job forbidden response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job forbidden response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job forbidden response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job forbidden response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job forbidden response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobNotFound creates a PatchKnowledgeKnowledgebaseImportJobNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobNotFound() *PatchKnowledgeKnowledgebaseImportJobNotFound {
	return &PatchKnowledgeKnowledgebaseImportJobNotFound{}
}

/*
PatchKnowledgeKnowledgebaseImportJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseImportJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job not found response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job not found response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job not found response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job not found response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job not found response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobRequestTimeout creates a PatchKnowledgeKnowledgebaseImportJobRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobRequestTimeout() *PatchKnowledgeKnowledgebaseImportJobRequestTimeout {
	return &PatchKnowledgeKnowledgebaseImportJobRequestTimeout{}
}

/*
PatchKnowledgeKnowledgebaseImportJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseImportJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job request timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job request timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job request timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job request timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job request timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge{}
}

/*
PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job request entity too large response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job request entity too large response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job request entity too large response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job request entity too large response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job request entity too large response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType() *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType{}
}

/*
PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job unsupported media type response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job unsupported media type response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job unsupported media type response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job unsupported media type response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job unsupported media type response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobTooManyRequests creates a PatchKnowledgeKnowledgebaseImportJobTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobTooManyRequests() *PatchKnowledgeKnowledgebaseImportJobTooManyRequests {
	return &PatchKnowledgeKnowledgebaseImportJobTooManyRequests{}
}

/*
PatchKnowledgeKnowledgebaseImportJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseImportJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job too many requests response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job too many requests response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job too many requests response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase import job too many requests response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase import job too many requests response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobInternalServerError creates a PatchKnowledgeKnowledgebaseImportJobInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobInternalServerError() *PatchKnowledgeKnowledgebaseImportJobInternalServerError {
	return &PatchKnowledgeKnowledgebaseImportJobInternalServerError{}
}

/*
PatchKnowledgeKnowledgebaseImportJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseImportJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job internal server error response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job internal server error response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job internal server error response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase import job internal server error response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase import job internal server error response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobServiceUnavailable creates a PatchKnowledgeKnowledgebaseImportJobServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobServiceUnavailable() *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseImportJobServiceUnavailable{}
}

/*
PatchKnowledgeKnowledgebaseImportJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseImportJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job service unavailable response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job service unavailable response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job service unavailable response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase import job service unavailable response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase import job service unavailable response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseImportJobGatewayTimeout creates a PatchKnowledgeKnowledgebaseImportJobGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseImportJobGatewayTimeout() *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseImportJobGatewayTimeout{}
}

/*
PatchKnowledgeKnowledgebaseImportJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseImportJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase import job gateway timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase import job gateway timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase import job gateway timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase import job gateway timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase import job gateway timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] patchKnowledgeKnowledgebaseImportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseImportJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
