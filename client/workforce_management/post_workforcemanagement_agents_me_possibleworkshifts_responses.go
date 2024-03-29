// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementAgentsMePossibleworkshiftsReader is a Reader for the PostWorkforcemanagementAgentsMePossibleworkshifts structure.
type PostWorkforcemanagementAgentsMePossibleworkshiftsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsOK creates a PostWorkforcemanagementAgentsMePossibleworkshiftsOK with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsOK() *PostWorkforcemanagementAgentsMePossibleworkshiftsOK {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsOK{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsOK struct {
	Payload *models.AgentPossibleWorkShiftsResponse
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts o k response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts o k response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts o k response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts o k response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts o k response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) GetPayload() *models.AgentPossibleWorkShiftsResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentPossibleWorkShiftsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest creates a PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest() *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts bad request response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts bad request response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts bad request response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts bad request response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts bad request response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized creates a PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized() *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsForbidden creates a PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsForbidden() *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts forbidden response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts forbidden response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts forbidden response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts forbidden response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts forbidden response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsNotFound creates a PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsNotFound() *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts not found response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts not found response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts not found response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts not found response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts not found response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout creates a PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout() *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts request timeout response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts request timeout response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts request timeout response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts request timeout response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts request timeout response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge creates a PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge() *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType creates a PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType() *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests creates a PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests() *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts too many requests response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts too many requests response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts too many requests response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts too many requests response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts too many requests response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError creates a PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError() *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts internal server error response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts internal server error response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts internal server error response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts internal server error response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts internal server error response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable creates a PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable() *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout creates a PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout with default headers values
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout() *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout{}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement agents me possibleworkshifts gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement agents me possibleworkshifts gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement agents me possibleworkshifts gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement agents me possibleworkshifts gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement agents me possibleworkshifts gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/agents/me/possibleworkshifts][%d] postWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
