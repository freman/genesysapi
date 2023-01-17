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

// GetConversationsChatParticipantWrapupcodesReader is a Reader for the GetConversationsChatParticipantWrapupcodes structure.
type GetConversationsChatParticipantWrapupcodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatParticipantWrapupcodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatParticipantWrapupcodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatParticipantWrapupcodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatParticipantWrapupcodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatParticipantWrapupcodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatParticipantWrapupcodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsChatParticipantWrapupcodesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatParticipantWrapupcodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatParticipantWrapupcodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatParticipantWrapupcodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatParticipantWrapupcodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatParticipantWrapupcodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatParticipantWrapupcodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatParticipantWrapupcodesOK creates a GetConversationsChatParticipantWrapupcodesOK with default headers values
func NewGetConversationsChatParticipantWrapupcodesOK() *GetConversationsChatParticipantWrapupcodesOK {
	return &GetConversationsChatParticipantWrapupcodesOK{}
}

/*
GetConversationsChatParticipantWrapupcodesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsChatParticipantWrapupcodesOK struct {
	Payload []*models.WrapupCode
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes o k response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes o k response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes o k response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapupcodes o k response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes o k response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsChatParticipantWrapupcodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesOK) GetPayload() []*models.WrapupCode {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesBadRequest creates a GetConversationsChatParticipantWrapupcodesBadRequest with default headers values
func NewGetConversationsChatParticipantWrapupcodesBadRequest() *GetConversationsChatParticipantWrapupcodesBadRequest {
	return &GetConversationsChatParticipantWrapupcodesBadRequest{}
}

/*
GetConversationsChatParticipantWrapupcodesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatParticipantWrapupcodesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes bad request response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes bad request response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes bad request response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes bad request response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes bad request response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsChatParticipantWrapupcodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesUnauthorized creates a GetConversationsChatParticipantWrapupcodesUnauthorized with default headers values
func NewGetConversationsChatParticipantWrapupcodesUnauthorized() *GetConversationsChatParticipantWrapupcodesUnauthorized {
	return &GetConversationsChatParticipantWrapupcodesUnauthorized{}
}

/*
GetConversationsChatParticipantWrapupcodesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatParticipantWrapupcodesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes unauthorized response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes unauthorized response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes unauthorized response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes unauthorized response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes unauthorized response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesForbidden creates a GetConversationsChatParticipantWrapupcodesForbidden with default headers values
func NewGetConversationsChatParticipantWrapupcodesForbidden() *GetConversationsChatParticipantWrapupcodesForbidden {
	return &GetConversationsChatParticipantWrapupcodesForbidden{}
}

/*
GetConversationsChatParticipantWrapupcodesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatParticipantWrapupcodesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes forbidden response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes forbidden response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes forbidden response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes forbidden response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes forbidden response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsChatParticipantWrapupcodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesNotFound creates a GetConversationsChatParticipantWrapupcodesNotFound with default headers values
func NewGetConversationsChatParticipantWrapupcodesNotFound() *GetConversationsChatParticipantWrapupcodesNotFound {
	return &GetConversationsChatParticipantWrapupcodesNotFound{}
}

/*
GetConversationsChatParticipantWrapupcodesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsChatParticipantWrapupcodesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes not found response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes not found response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes not found response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes not found response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes not found response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsChatParticipantWrapupcodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesRequestTimeout creates a GetConversationsChatParticipantWrapupcodesRequestTimeout with default headers values
func NewGetConversationsChatParticipantWrapupcodesRequestTimeout() *GetConversationsChatParticipantWrapupcodesRequestTimeout {
	return &GetConversationsChatParticipantWrapupcodesRequestTimeout{}
}

/*
GetConversationsChatParticipantWrapupcodesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsChatParticipantWrapupcodesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes request timeout response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes request timeout response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes request timeout response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes request timeout response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes request timeout response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesRequestEntityTooLarge creates a GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge with default headers values
func NewGetConversationsChatParticipantWrapupcodesRequestEntityTooLarge() *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge {
	return &GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge{}
}

/*
GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes request entity too large response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes request entity too large response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes request entity too large response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes request entity too large response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes request entity too large response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesUnsupportedMediaType creates a GetConversationsChatParticipantWrapupcodesUnsupportedMediaType with default headers values
func NewGetConversationsChatParticipantWrapupcodesUnsupportedMediaType() *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType {
	return &GetConversationsChatParticipantWrapupcodesUnsupportedMediaType{}
}

/*
GetConversationsChatParticipantWrapupcodesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatParticipantWrapupcodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes unsupported media type response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes unsupported media type response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes unsupported media type response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes unsupported media type response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes unsupported media type response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesTooManyRequests creates a GetConversationsChatParticipantWrapupcodesTooManyRequests with default headers values
func NewGetConversationsChatParticipantWrapupcodesTooManyRequests() *GetConversationsChatParticipantWrapupcodesTooManyRequests {
	return &GetConversationsChatParticipantWrapupcodesTooManyRequests{}
}

/*
GetConversationsChatParticipantWrapupcodesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsChatParticipantWrapupcodesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes too many requests response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes too many requests response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes too many requests response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapupcodes too many requests response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapupcodes too many requests response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesInternalServerError creates a GetConversationsChatParticipantWrapupcodesInternalServerError with default headers values
func NewGetConversationsChatParticipantWrapupcodesInternalServerError() *GetConversationsChatParticipantWrapupcodesInternalServerError {
	return &GetConversationsChatParticipantWrapupcodesInternalServerError{}
}

/*
GetConversationsChatParticipantWrapupcodesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatParticipantWrapupcodesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes internal server error response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes internal server error response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes internal server error response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapupcodes internal server error response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat participant wrapupcodes internal server error response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesServiceUnavailable creates a GetConversationsChatParticipantWrapupcodesServiceUnavailable with default headers values
func NewGetConversationsChatParticipantWrapupcodesServiceUnavailable() *GetConversationsChatParticipantWrapupcodesServiceUnavailable {
	return &GetConversationsChatParticipantWrapupcodesServiceUnavailable{}
}

/*
GetConversationsChatParticipantWrapupcodesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatParticipantWrapupcodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes service unavailable response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes service unavailable response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes service unavailable response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapupcodes service unavailable response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat participant wrapupcodes service unavailable response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupcodesGatewayTimeout creates a GetConversationsChatParticipantWrapupcodesGatewayTimeout with default headers values
func NewGetConversationsChatParticipantWrapupcodesGatewayTimeout() *GetConversationsChatParticipantWrapupcodesGatewayTimeout {
	return &GetConversationsChatParticipantWrapupcodesGatewayTimeout{}
}

/*
GetConversationsChatParticipantWrapupcodesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsChatParticipantWrapupcodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapupcodes gateway timeout response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapupcodes gateway timeout response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapupcodes gateway timeout response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapupcodes gateway timeout response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat participant wrapupcodes gateway timeout response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsChatParticipantWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupcodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
