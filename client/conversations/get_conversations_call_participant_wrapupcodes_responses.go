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

// GetConversationsCallParticipantWrapupcodesReader is a Reader for the GetConversationsCallParticipantWrapupcodes structure.
type GetConversationsCallParticipantWrapupcodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallParticipantWrapupcodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallParticipantWrapupcodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallParticipantWrapupcodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallParticipantWrapupcodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallParticipantWrapupcodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallParticipantWrapupcodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallParticipantWrapupcodesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallParticipantWrapupcodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallParticipantWrapupcodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallParticipantWrapupcodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallParticipantWrapupcodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallParticipantWrapupcodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallParticipantWrapupcodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallParticipantWrapupcodesOK creates a GetConversationsCallParticipantWrapupcodesOK with default headers values
func NewGetConversationsCallParticipantWrapupcodesOK() *GetConversationsCallParticipantWrapupcodesOK {
	return &GetConversationsCallParticipantWrapupcodesOK{}
}

/*
GetConversationsCallParticipantWrapupcodesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsCallParticipantWrapupcodesOK struct {
	Payload []*models.WrapupCode
}

// IsSuccess returns true when this get conversations call participant wrapupcodes o k response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations call participant wrapupcodes o k response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes o k response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapupcodes o k response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes o k response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsCallParticipantWrapupcodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesOK) GetPayload() []*models.WrapupCode {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesBadRequest creates a GetConversationsCallParticipantWrapupcodesBadRequest with default headers values
func NewGetConversationsCallParticipantWrapupcodesBadRequest() *GetConversationsCallParticipantWrapupcodesBadRequest {
	return &GetConversationsCallParticipantWrapupcodesBadRequest{}
}

/*
GetConversationsCallParticipantWrapupcodesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallParticipantWrapupcodesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes bad request response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes bad request response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes bad request response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes bad request response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes bad request response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsCallParticipantWrapupcodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesUnauthorized creates a GetConversationsCallParticipantWrapupcodesUnauthorized with default headers values
func NewGetConversationsCallParticipantWrapupcodesUnauthorized() *GetConversationsCallParticipantWrapupcodesUnauthorized {
	return &GetConversationsCallParticipantWrapupcodesUnauthorized{}
}

/*
GetConversationsCallParticipantWrapupcodesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallParticipantWrapupcodesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes unauthorized response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes unauthorized response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes unauthorized response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes unauthorized response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes unauthorized response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesForbidden creates a GetConversationsCallParticipantWrapupcodesForbidden with default headers values
func NewGetConversationsCallParticipantWrapupcodesForbidden() *GetConversationsCallParticipantWrapupcodesForbidden {
	return &GetConversationsCallParticipantWrapupcodesForbidden{}
}

/*
GetConversationsCallParticipantWrapupcodesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallParticipantWrapupcodesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes forbidden response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes forbidden response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes forbidden response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes forbidden response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes forbidden response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsCallParticipantWrapupcodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesNotFound creates a GetConversationsCallParticipantWrapupcodesNotFound with default headers values
func NewGetConversationsCallParticipantWrapupcodesNotFound() *GetConversationsCallParticipantWrapupcodesNotFound {
	return &GetConversationsCallParticipantWrapupcodesNotFound{}
}

/*
GetConversationsCallParticipantWrapupcodesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsCallParticipantWrapupcodesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes not found response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes not found response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes not found response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes not found response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes not found response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsCallParticipantWrapupcodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesRequestTimeout creates a GetConversationsCallParticipantWrapupcodesRequestTimeout with default headers values
func NewGetConversationsCallParticipantWrapupcodesRequestTimeout() *GetConversationsCallParticipantWrapupcodesRequestTimeout {
	return &GetConversationsCallParticipantWrapupcodesRequestTimeout{}
}

/*
GetConversationsCallParticipantWrapupcodesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallParticipantWrapupcodesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes request timeout response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes request timeout response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes request timeout response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes request timeout response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes request timeout response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesRequestEntityTooLarge creates a GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge with default headers values
func NewGetConversationsCallParticipantWrapupcodesRequestEntityTooLarge() *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge {
	return &GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge{}
}

/*
GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes request entity too large response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes request entity too large response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes request entity too large response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes request entity too large response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes request entity too large response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesUnsupportedMediaType creates a GetConversationsCallParticipantWrapupcodesUnsupportedMediaType with default headers values
func NewGetConversationsCallParticipantWrapupcodesUnsupportedMediaType() *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType {
	return &GetConversationsCallParticipantWrapupcodesUnsupportedMediaType{}
}

/*
GetConversationsCallParticipantWrapupcodesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallParticipantWrapupcodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes unsupported media type response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes unsupported media type response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes unsupported media type response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes unsupported media type response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes unsupported media type response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesTooManyRequests creates a GetConversationsCallParticipantWrapupcodesTooManyRequests with default headers values
func NewGetConversationsCallParticipantWrapupcodesTooManyRequests() *GetConversationsCallParticipantWrapupcodesTooManyRequests {
	return &GetConversationsCallParticipantWrapupcodesTooManyRequests{}
}

/*
GetConversationsCallParticipantWrapupcodesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallParticipantWrapupcodesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes too many requests response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes too many requests response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes too many requests response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapupcodes too many requests response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapupcodes too many requests response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesInternalServerError creates a GetConversationsCallParticipantWrapupcodesInternalServerError with default headers values
func NewGetConversationsCallParticipantWrapupcodesInternalServerError() *GetConversationsCallParticipantWrapupcodesInternalServerError {
	return &GetConversationsCallParticipantWrapupcodesInternalServerError{}
}

/*
GetConversationsCallParticipantWrapupcodesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallParticipantWrapupcodesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes internal server error response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes internal server error response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes internal server error response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapupcodes internal server error response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations call participant wrapupcodes internal server error response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesServiceUnavailable creates a GetConversationsCallParticipantWrapupcodesServiceUnavailable with default headers values
func NewGetConversationsCallParticipantWrapupcodesServiceUnavailable() *GetConversationsCallParticipantWrapupcodesServiceUnavailable {
	return &GetConversationsCallParticipantWrapupcodesServiceUnavailable{}
}

/*
GetConversationsCallParticipantWrapupcodesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallParticipantWrapupcodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes service unavailable response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes service unavailable response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes service unavailable response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapupcodes service unavailable response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations call participant wrapupcodes service unavailable response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupcodesGatewayTimeout creates a GetConversationsCallParticipantWrapupcodesGatewayTimeout with default headers values
func NewGetConversationsCallParticipantWrapupcodesGatewayTimeout() *GetConversationsCallParticipantWrapupcodesGatewayTimeout {
	return &GetConversationsCallParticipantWrapupcodesGatewayTimeout{}
}

/*
GetConversationsCallParticipantWrapupcodesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsCallParticipantWrapupcodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapupcodes gateway timeout response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapupcodes gateway timeout response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapupcodes gateway timeout response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapupcodes gateway timeout response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations call participant wrapupcodes gateway timeout response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallParticipantWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupcodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
