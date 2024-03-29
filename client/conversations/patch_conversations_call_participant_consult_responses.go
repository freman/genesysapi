// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchConversationsCallParticipantConsultReader is a Reader for the PatchConversationsCallParticipantConsult structure.
type PatchConversationsCallParticipantConsultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCallParticipantConsultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsCallParticipantConsultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCallParticipantConsultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCallParticipantConsultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCallParticipantConsultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCallParticipantConsultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCallParticipantConsultRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCallParticipantConsultRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCallParticipantConsultUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCallParticipantConsultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCallParticipantConsultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCallParticipantConsultServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCallParticipantConsultGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCallParticipantConsultOK creates a PatchConversationsCallParticipantConsultOK with default headers values
func NewPatchConversationsCallParticipantConsultOK() *PatchConversationsCallParticipantConsultOK {
	return &PatchConversationsCallParticipantConsultOK{}
}

/*
PatchConversationsCallParticipantConsultOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchConversationsCallParticipantConsultOK struct {
	Payload *models.ConsultTransferResponse
}

// IsSuccess returns true when this patch conversations call participant consult o k response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations call participant consult o k response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult o k response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant consult o k response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult o k response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchConversationsCallParticipantConsultOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultOK) GetPayload() *models.ConsultTransferResponse {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsultTransferResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultBadRequest creates a PatchConversationsCallParticipantConsultBadRequest with default headers values
func NewPatchConversationsCallParticipantConsultBadRequest() *PatchConversationsCallParticipantConsultBadRequest {
	return &PatchConversationsCallParticipantConsultBadRequest{}
}

/*
PatchConversationsCallParticipantConsultBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCallParticipantConsultBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult bad request response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult bad request response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult bad request response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult bad request response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult bad request response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsCallParticipantConsultBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultUnauthorized creates a PatchConversationsCallParticipantConsultUnauthorized with default headers values
func NewPatchConversationsCallParticipantConsultUnauthorized() *PatchConversationsCallParticipantConsultUnauthorized {
	return &PatchConversationsCallParticipantConsultUnauthorized{}
}

/*
PatchConversationsCallParticipantConsultUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCallParticipantConsultUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult unauthorized response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult unauthorized response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult unauthorized response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult unauthorized response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult unauthorized response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsCallParticipantConsultUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultForbidden creates a PatchConversationsCallParticipantConsultForbidden with default headers values
func NewPatchConversationsCallParticipantConsultForbidden() *PatchConversationsCallParticipantConsultForbidden {
	return &PatchConversationsCallParticipantConsultForbidden{}
}

/*
PatchConversationsCallParticipantConsultForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCallParticipantConsultForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult forbidden response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult forbidden response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult forbidden response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult forbidden response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult forbidden response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsCallParticipantConsultForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultNotFound creates a PatchConversationsCallParticipantConsultNotFound with default headers values
func NewPatchConversationsCallParticipantConsultNotFound() *PatchConversationsCallParticipantConsultNotFound {
	return &PatchConversationsCallParticipantConsultNotFound{}
}

/*
PatchConversationsCallParticipantConsultNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsCallParticipantConsultNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult not found response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult not found response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult not found response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult not found response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult not found response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsCallParticipantConsultNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultRequestTimeout creates a PatchConversationsCallParticipantConsultRequestTimeout with default headers values
func NewPatchConversationsCallParticipantConsultRequestTimeout() *PatchConversationsCallParticipantConsultRequestTimeout {
	return &PatchConversationsCallParticipantConsultRequestTimeout{}
}

/*
PatchConversationsCallParticipantConsultRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCallParticipantConsultRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult request timeout response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult request timeout response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult request timeout response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult request timeout response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult request timeout response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsCallParticipantConsultRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultRequestEntityTooLarge creates a PatchConversationsCallParticipantConsultRequestEntityTooLarge with default headers values
func NewPatchConversationsCallParticipantConsultRequestEntityTooLarge() *PatchConversationsCallParticipantConsultRequestEntityTooLarge {
	return &PatchConversationsCallParticipantConsultRequestEntityTooLarge{}
}

/*
PatchConversationsCallParticipantConsultRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsCallParticipantConsultRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult request entity too large response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult request entity too large response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult request entity too large response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult request entity too large response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult request entity too large response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultUnsupportedMediaType creates a PatchConversationsCallParticipantConsultUnsupportedMediaType with default headers values
func NewPatchConversationsCallParticipantConsultUnsupportedMediaType() *PatchConversationsCallParticipantConsultUnsupportedMediaType {
	return &PatchConversationsCallParticipantConsultUnsupportedMediaType{}
}

/*
PatchConversationsCallParticipantConsultUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCallParticipantConsultUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult unsupported media type response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult unsupported media type response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult unsupported media type response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult unsupported media type response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult unsupported media type response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultTooManyRequests creates a PatchConversationsCallParticipantConsultTooManyRequests with default headers values
func NewPatchConversationsCallParticipantConsultTooManyRequests() *PatchConversationsCallParticipantConsultTooManyRequests {
	return &PatchConversationsCallParticipantConsultTooManyRequests{}
}

/*
PatchConversationsCallParticipantConsultTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCallParticipantConsultTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult too many requests response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult too many requests response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult too many requests response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant consult too many requests response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant consult too many requests response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsCallParticipantConsultTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultInternalServerError creates a PatchConversationsCallParticipantConsultInternalServerError with default headers values
func NewPatchConversationsCallParticipantConsultInternalServerError() *PatchConversationsCallParticipantConsultInternalServerError {
	return &PatchConversationsCallParticipantConsultInternalServerError{}
}

/*
PatchConversationsCallParticipantConsultInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCallParticipantConsultInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult internal server error response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult internal server error response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult internal server error response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant consult internal server error response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations call participant consult internal server error response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsCallParticipantConsultInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultServiceUnavailable creates a PatchConversationsCallParticipantConsultServiceUnavailable with default headers values
func NewPatchConversationsCallParticipantConsultServiceUnavailable() *PatchConversationsCallParticipantConsultServiceUnavailable {
	return &PatchConversationsCallParticipantConsultServiceUnavailable{}
}

/*
PatchConversationsCallParticipantConsultServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCallParticipantConsultServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult service unavailable response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult service unavailable response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult service unavailable response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant consult service unavailable response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations call participant consult service unavailable response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsCallParticipantConsultServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantConsultGatewayTimeout creates a PatchConversationsCallParticipantConsultGatewayTimeout with default headers values
func NewPatchConversationsCallParticipantConsultGatewayTimeout() *PatchConversationsCallParticipantConsultGatewayTimeout {
	return &PatchConversationsCallParticipantConsultGatewayTimeout{}
}

/*
PatchConversationsCallParticipantConsultGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsCallParticipantConsultGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant consult gateway timeout response has a 2xx status code
func (o *PatchConversationsCallParticipantConsultGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant consult gateway timeout response has a 3xx status code
func (o *PatchConversationsCallParticipantConsultGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant consult gateway timeout response has a 4xx status code
func (o *PatchConversationsCallParticipantConsultGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant consult gateway timeout response has a 5xx status code
func (o *PatchConversationsCallParticipantConsultGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations call participant consult gateway timeout response a status code equal to that given
func (o *PatchConversationsCallParticipantConsultGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsCallParticipantConsultGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] patchConversationsCallParticipantConsultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallParticipantConsultGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantConsultGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
