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

// PatchConversationsCallParticipantCommunicationReader is a Reader for the PatchConversationsCallParticipantCommunication structure.
type PatchConversationsCallParticipantCommunicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCallParticipantCommunicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsCallParticipantCommunicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCallParticipantCommunicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCallParticipantCommunicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCallParticipantCommunicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCallParticipantCommunicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCallParticipantCommunicationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCallParticipantCommunicationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCallParticipantCommunicationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCallParticipantCommunicationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCallParticipantCommunicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCallParticipantCommunicationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCallParticipantCommunicationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCallParticipantCommunicationOK creates a PatchConversationsCallParticipantCommunicationOK with default headers values
func NewPatchConversationsCallParticipantCommunicationOK() *PatchConversationsCallParticipantCommunicationOK {
	return &PatchConversationsCallParticipantCommunicationOK{}
}

/*
PatchConversationsCallParticipantCommunicationOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchConversationsCallParticipantCommunicationOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this patch conversations call participant communication o k response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations call participant communication o k response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication o k response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant communication o k response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication o k response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchConversationsCallParticipantCommunicationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationBadRequest creates a PatchConversationsCallParticipantCommunicationBadRequest with default headers values
func NewPatchConversationsCallParticipantCommunicationBadRequest() *PatchConversationsCallParticipantCommunicationBadRequest {
	return &PatchConversationsCallParticipantCommunicationBadRequest{}
}

/*
PatchConversationsCallParticipantCommunicationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCallParticipantCommunicationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication bad request response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication bad request response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication bad request response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication bad request response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication bad request response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsCallParticipantCommunicationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationUnauthorized creates a PatchConversationsCallParticipantCommunicationUnauthorized with default headers values
func NewPatchConversationsCallParticipantCommunicationUnauthorized() *PatchConversationsCallParticipantCommunicationUnauthorized {
	return &PatchConversationsCallParticipantCommunicationUnauthorized{}
}

/*
PatchConversationsCallParticipantCommunicationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCallParticipantCommunicationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication unauthorized response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication unauthorized response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication unauthorized response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication unauthorized response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication unauthorized response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsCallParticipantCommunicationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationForbidden creates a PatchConversationsCallParticipantCommunicationForbidden with default headers values
func NewPatchConversationsCallParticipantCommunicationForbidden() *PatchConversationsCallParticipantCommunicationForbidden {
	return &PatchConversationsCallParticipantCommunicationForbidden{}
}

/*
PatchConversationsCallParticipantCommunicationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCallParticipantCommunicationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication forbidden response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication forbidden response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication forbidden response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication forbidden response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication forbidden response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsCallParticipantCommunicationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationNotFound creates a PatchConversationsCallParticipantCommunicationNotFound with default headers values
func NewPatchConversationsCallParticipantCommunicationNotFound() *PatchConversationsCallParticipantCommunicationNotFound {
	return &PatchConversationsCallParticipantCommunicationNotFound{}
}

/*
PatchConversationsCallParticipantCommunicationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsCallParticipantCommunicationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication not found response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication not found response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication not found response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication not found response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication not found response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsCallParticipantCommunicationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationRequestTimeout creates a PatchConversationsCallParticipantCommunicationRequestTimeout with default headers values
func NewPatchConversationsCallParticipantCommunicationRequestTimeout() *PatchConversationsCallParticipantCommunicationRequestTimeout {
	return &PatchConversationsCallParticipantCommunicationRequestTimeout{}
}

/*
PatchConversationsCallParticipantCommunicationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCallParticipantCommunicationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication request timeout response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication request timeout response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication request timeout response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication request timeout response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication request timeout response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationRequestEntityTooLarge creates a PatchConversationsCallParticipantCommunicationRequestEntityTooLarge with default headers values
func NewPatchConversationsCallParticipantCommunicationRequestEntityTooLarge() *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge {
	return &PatchConversationsCallParticipantCommunicationRequestEntityTooLarge{}
}

/*
PatchConversationsCallParticipantCommunicationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsCallParticipantCommunicationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication request entity too large response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication request entity too large response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication request entity too large response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication request entity too large response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication request entity too large response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationUnsupportedMediaType creates a PatchConversationsCallParticipantCommunicationUnsupportedMediaType with default headers values
func NewPatchConversationsCallParticipantCommunicationUnsupportedMediaType() *PatchConversationsCallParticipantCommunicationUnsupportedMediaType {
	return &PatchConversationsCallParticipantCommunicationUnsupportedMediaType{}
}

/*
PatchConversationsCallParticipantCommunicationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCallParticipantCommunicationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication unsupported media type response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication unsupported media type response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication unsupported media type response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication unsupported media type response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication unsupported media type response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationTooManyRequests creates a PatchConversationsCallParticipantCommunicationTooManyRequests with default headers values
func NewPatchConversationsCallParticipantCommunicationTooManyRequests() *PatchConversationsCallParticipantCommunicationTooManyRequests {
	return &PatchConversationsCallParticipantCommunicationTooManyRequests{}
}

/*
PatchConversationsCallParticipantCommunicationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCallParticipantCommunicationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication too many requests response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication too many requests response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication too many requests response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations call participant communication too many requests response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations call participant communication too many requests response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationInternalServerError creates a PatchConversationsCallParticipantCommunicationInternalServerError with default headers values
func NewPatchConversationsCallParticipantCommunicationInternalServerError() *PatchConversationsCallParticipantCommunicationInternalServerError {
	return &PatchConversationsCallParticipantCommunicationInternalServerError{}
}

/*
PatchConversationsCallParticipantCommunicationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCallParticipantCommunicationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication internal server error response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication internal server error response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication internal server error response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant communication internal server error response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations call participant communication internal server error response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsCallParticipantCommunicationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationServiceUnavailable creates a PatchConversationsCallParticipantCommunicationServiceUnavailable with default headers values
func NewPatchConversationsCallParticipantCommunicationServiceUnavailable() *PatchConversationsCallParticipantCommunicationServiceUnavailable {
	return &PatchConversationsCallParticipantCommunicationServiceUnavailable{}
}

/*
PatchConversationsCallParticipantCommunicationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCallParticipantCommunicationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication service unavailable response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication service unavailable response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication service unavailable response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant communication service unavailable response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations call participant communication service unavailable response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallParticipantCommunicationGatewayTimeout creates a PatchConversationsCallParticipantCommunicationGatewayTimeout with default headers values
func NewPatchConversationsCallParticipantCommunicationGatewayTimeout() *PatchConversationsCallParticipantCommunicationGatewayTimeout {
	return &PatchConversationsCallParticipantCommunicationGatewayTimeout{}
}

/*
PatchConversationsCallParticipantCommunicationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsCallParticipantCommunicationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations call participant communication gateway timeout response has a 2xx status code
func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations call participant communication gateway timeout response has a 3xx status code
func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations call participant communication gateway timeout response has a 4xx status code
func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations call participant communication gateway timeout response has a 5xx status code
func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations call participant communication gateway timeout response a status code equal to that given
func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/calls/{conversationId}/participants/{participantId}/communications/{communicationId}][%d] patchConversationsCallParticipantCommunicationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallParticipantCommunicationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
