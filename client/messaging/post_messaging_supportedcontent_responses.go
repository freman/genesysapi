// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostMessagingSupportedcontentReader is a Reader for the PostMessagingSupportedcontent structure.
type PostMessagingSupportedcontentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMessagingSupportedcontentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMessagingSupportedcontentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostMessagingSupportedcontentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMessagingSupportedcontentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostMessagingSupportedcontentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostMessagingSupportedcontentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostMessagingSupportedcontentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostMessagingSupportedcontentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostMessagingSupportedcontentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostMessagingSupportedcontentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostMessagingSupportedcontentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMessagingSupportedcontentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostMessagingSupportedcontentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostMessagingSupportedcontentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostMessagingSupportedcontentOK creates a PostMessagingSupportedcontentOK with default headers values
func NewPostMessagingSupportedcontentOK() *PostMessagingSupportedcontentOK {
	return &PostMessagingSupportedcontentOK{}
}

/*
PostMessagingSupportedcontentOK describes a response with status code 200, with default header values.

successful operation
*/
type PostMessagingSupportedcontentOK struct {
	Payload *models.SupportedContent
}

// IsSuccess returns true when this post messaging supportedcontent o k response has a 2xx status code
func (o *PostMessagingSupportedcontentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post messaging supportedcontent o k response has a 3xx status code
func (o *PostMessagingSupportedcontentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent o k response has a 4xx status code
func (o *PostMessagingSupportedcontentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post messaging supportedcontent o k response has a 5xx status code
func (o *PostMessagingSupportedcontentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent o k response a status code equal to that given
func (o *PostMessagingSupportedcontentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostMessagingSupportedcontentOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentOK  %+v", 200, o.Payload)
}

func (o *PostMessagingSupportedcontentOK) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentOK  %+v", 200, o.Payload)
}

func (o *PostMessagingSupportedcontentOK) GetPayload() *models.SupportedContent {
	return o.Payload
}

func (o *PostMessagingSupportedcontentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentCreated creates a PostMessagingSupportedcontentCreated with default headers values
func NewPostMessagingSupportedcontentCreated() *PostMessagingSupportedcontentCreated {
	return &PostMessagingSupportedcontentCreated{}
}

/*
PostMessagingSupportedcontentCreated describes a response with status code 201, with default header values.

Operation was successful
*/
type PostMessagingSupportedcontentCreated struct {
	Payload *models.SupportedContent
}

// IsSuccess returns true when this post messaging supportedcontent created response has a 2xx status code
func (o *PostMessagingSupportedcontentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post messaging supportedcontent created response has a 3xx status code
func (o *PostMessagingSupportedcontentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent created response has a 4xx status code
func (o *PostMessagingSupportedcontentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post messaging supportedcontent created response has a 5xx status code
func (o *PostMessagingSupportedcontentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent created response a status code equal to that given
func (o *PostMessagingSupportedcontentCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostMessagingSupportedcontentCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentCreated  %+v", 201, o.Payload)
}

func (o *PostMessagingSupportedcontentCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentCreated  %+v", 201, o.Payload)
}

func (o *PostMessagingSupportedcontentCreated) GetPayload() *models.SupportedContent {
	return o.Payload
}

func (o *PostMessagingSupportedcontentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentBadRequest creates a PostMessagingSupportedcontentBadRequest with default headers values
func NewPostMessagingSupportedcontentBadRequest() *PostMessagingSupportedcontentBadRequest {
	return &PostMessagingSupportedcontentBadRequest{}
}

/*
PostMessagingSupportedcontentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostMessagingSupportedcontentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent bad request response has a 2xx status code
func (o *PostMessagingSupportedcontentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent bad request response has a 3xx status code
func (o *PostMessagingSupportedcontentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent bad request response has a 4xx status code
func (o *PostMessagingSupportedcontentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent bad request response has a 5xx status code
func (o *PostMessagingSupportedcontentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent bad request response a status code equal to that given
func (o *PostMessagingSupportedcontentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostMessagingSupportedcontentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentBadRequest  %+v", 400, o.Payload)
}

func (o *PostMessagingSupportedcontentBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentBadRequest  %+v", 400, o.Payload)
}

func (o *PostMessagingSupportedcontentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentUnauthorized creates a PostMessagingSupportedcontentUnauthorized with default headers values
func NewPostMessagingSupportedcontentUnauthorized() *PostMessagingSupportedcontentUnauthorized {
	return &PostMessagingSupportedcontentUnauthorized{}
}

/*
PostMessagingSupportedcontentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostMessagingSupportedcontentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent unauthorized response has a 2xx status code
func (o *PostMessagingSupportedcontentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent unauthorized response has a 3xx status code
func (o *PostMessagingSupportedcontentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent unauthorized response has a 4xx status code
func (o *PostMessagingSupportedcontentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent unauthorized response has a 5xx status code
func (o *PostMessagingSupportedcontentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent unauthorized response a status code equal to that given
func (o *PostMessagingSupportedcontentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostMessagingSupportedcontentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentUnauthorized  %+v", 401, o.Payload)
}

func (o *PostMessagingSupportedcontentUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentUnauthorized  %+v", 401, o.Payload)
}

func (o *PostMessagingSupportedcontentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentForbidden creates a PostMessagingSupportedcontentForbidden with default headers values
func NewPostMessagingSupportedcontentForbidden() *PostMessagingSupportedcontentForbidden {
	return &PostMessagingSupportedcontentForbidden{}
}

/*
PostMessagingSupportedcontentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostMessagingSupportedcontentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent forbidden response has a 2xx status code
func (o *PostMessagingSupportedcontentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent forbidden response has a 3xx status code
func (o *PostMessagingSupportedcontentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent forbidden response has a 4xx status code
func (o *PostMessagingSupportedcontentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent forbidden response has a 5xx status code
func (o *PostMessagingSupportedcontentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent forbidden response a status code equal to that given
func (o *PostMessagingSupportedcontentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostMessagingSupportedcontentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentForbidden  %+v", 403, o.Payload)
}

func (o *PostMessagingSupportedcontentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentForbidden  %+v", 403, o.Payload)
}

func (o *PostMessagingSupportedcontentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentNotFound creates a PostMessagingSupportedcontentNotFound with default headers values
func NewPostMessagingSupportedcontentNotFound() *PostMessagingSupportedcontentNotFound {
	return &PostMessagingSupportedcontentNotFound{}
}

/*
PostMessagingSupportedcontentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostMessagingSupportedcontentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent not found response has a 2xx status code
func (o *PostMessagingSupportedcontentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent not found response has a 3xx status code
func (o *PostMessagingSupportedcontentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent not found response has a 4xx status code
func (o *PostMessagingSupportedcontentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent not found response has a 5xx status code
func (o *PostMessagingSupportedcontentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent not found response a status code equal to that given
func (o *PostMessagingSupportedcontentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostMessagingSupportedcontentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentNotFound  %+v", 404, o.Payload)
}

func (o *PostMessagingSupportedcontentNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentNotFound  %+v", 404, o.Payload)
}

func (o *PostMessagingSupportedcontentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentRequestTimeout creates a PostMessagingSupportedcontentRequestTimeout with default headers values
func NewPostMessagingSupportedcontentRequestTimeout() *PostMessagingSupportedcontentRequestTimeout {
	return &PostMessagingSupportedcontentRequestTimeout{}
}

/*
PostMessagingSupportedcontentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostMessagingSupportedcontentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent request timeout response has a 2xx status code
func (o *PostMessagingSupportedcontentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent request timeout response has a 3xx status code
func (o *PostMessagingSupportedcontentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent request timeout response has a 4xx status code
func (o *PostMessagingSupportedcontentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent request timeout response has a 5xx status code
func (o *PostMessagingSupportedcontentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent request timeout response a status code equal to that given
func (o *PostMessagingSupportedcontentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostMessagingSupportedcontentRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostMessagingSupportedcontentRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostMessagingSupportedcontentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentRequestEntityTooLarge creates a PostMessagingSupportedcontentRequestEntityTooLarge with default headers values
func NewPostMessagingSupportedcontentRequestEntityTooLarge() *PostMessagingSupportedcontentRequestEntityTooLarge {
	return &PostMessagingSupportedcontentRequestEntityTooLarge{}
}

/*
PostMessagingSupportedcontentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostMessagingSupportedcontentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent request entity too large response has a 2xx status code
func (o *PostMessagingSupportedcontentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent request entity too large response has a 3xx status code
func (o *PostMessagingSupportedcontentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent request entity too large response has a 4xx status code
func (o *PostMessagingSupportedcontentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent request entity too large response has a 5xx status code
func (o *PostMessagingSupportedcontentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent request entity too large response a status code equal to that given
func (o *PostMessagingSupportedcontentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostMessagingSupportedcontentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostMessagingSupportedcontentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostMessagingSupportedcontentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentUnsupportedMediaType creates a PostMessagingSupportedcontentUnsupportedMediaType with default headers values
func NewPostMessagingSupportedcontentUnsupportedMediaType() *PostMessagingSupportedcontentUnsupportedMediaType {
	return &PostMessagingSupportedcontentUnsupportedMediaType{}
}

/*
PostMessagingSupportedcontentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostMessagingSupportedcontentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent unsupported media type response has a 2xx status code
func (o *PostMessagingSupportedcontentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent unsupported media type response has a 3xx status code
func (o *PostMessagingSupportedcontentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent unsupported media type response has a 4xx status code
func (o *PostMessagingSupportedcontentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent unsupported media type response has a 5xx status code
func (o *PostMessagingSupportedcontentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent unsupported media type response a status code equal to that given
func (o *PostMessagingSupportedcontentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostMessagingSupportedcontentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostMessagingSupportedcontentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostMessagingSupportedcontentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentTooManyRequests creates a PostMessagingSupportedcontentTooManyRequests with default headers values
func NewPostMessagingSupportedcontentTooManyRequests() *PostMessagingSupportedcontentTooManyRequests {
	return &PostMessagingSupportedcontentTooManyRequests{}
}

/*
PostMessagingSupportedcontentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostMessagingSupportedcontentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent too many requests response has a 2xx status code
func (o *PostMessagingSupportedcontentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent too many requests response has a 3xx status code
func (o *PostMessagingSupportedcontentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent too many requests response has a 4xx status code
func (o *PostMessagingSupportedcontentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post messaging supportedcontent too many requests response has a 5xx status code
func (o *PostMessagingSupportedcontentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post messaging supportedcontent too many requests response a status code equal to that given
func (o *PostMessagingSupportedcontentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostMessagingSupportedcontentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMessagingSupportedcontentTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMessagingSupportedcontentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentInternalServerError creates a PostMessagingSupportedcontentInternalServerError with default headers values
func NewPostMessagingSupportedcontentInternalServerError() *PostMessagingSupportedcontentInternalServerError {
	return &PostMessagingSupportedcontentInternalServerError{}
}

/*
PostMessagingSupportedcontentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostMessagingSupportedcontentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent internal server error response has a 2xx status code
func (o *PostMessagingSupportedcontentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent internal server error response has a 3xx status code
func (o *PostMessagingSupportedcontentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent internal server error response has a 4xx status code
func (o *PostMessagingSupportedcontentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post messaging supportedcontent internal server error response has a 5xx status code
func (o *PostMessagingSupportedcontentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post messaging supportedcontent internal server error response a status code equal to that given
func (o *PostMessagingSupportedcontentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostMessagingSupportedcontentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentInternalServerError  %+v", 500, o.Payload)
}

func (o *PostMessagingSupportedcontentInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentInternalServerError  %+v", 500, o.Payload)
}

func (o *PostMessagingSupportedcontentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentServiceUnavailable creates a PostMessagingSupportedcontentServiceUnavailable with default headers values
func NewPostMessagingSupportedcontentServiceUnavailable() *PostMessagingSupportedcontentServiceUnavailable {
	return &PostMessagingSupportedcontentServiceUnavailable{}
}

/*
PostMessagingSupportedcontentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostMessagingSupportedcontentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent service unavailable response has a 2xx status code
func (o *PostMessagingSupportedcontentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent service unavailable response has a 3xx status code
func (o *PostMessagingSupportedcontentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent service unavailable response has a 4xx status code
func (o *PostMessagingSupportedcontentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post messaging supportedcontent service unavailable response has a 5xx status code
func (o *PostMessagingSupportedcontentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post messaging supportedcontent service unavailable response a status code equal to that given
func (o *PostMessagingSupportedcontentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostMessagingSupportedcontentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostMessagingSupportedcontentServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostMessagingSupportedcontentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMessagingSupportedcontentGatewayTimeout creates a PostMessagingSupportedcontentGatewayTimeout with default headers values
func NewPostMessagingSupportedcontentGatewayTimeout() *PostMessagingSupportedcontentGatewayTimeout {
	return &PostMessagingSupportedcontentGatewayTimeout{}
}

/*
PostMessagingSupportedcontentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostMessagingSupportedcontentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post messaging supportedcontent gateway timeout response has a 2xx status code
func (o *PostMessagingSupportedcontentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post messaging supportedcontent gateway timeout response has a 3xx status code
func (o *PostMessagingSupportedcontentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post messaging supportedcontent gateway timeout response has a 4xx status code
func (o *PostMessagingSupportedcontentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post messaging supportedcontent gateway timeout response has a 5xx status code
func (o *PostMessagingSupportedcontentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post messaging supportedcontent gateway timeout response a status code equal to that given
func (o *PostMessagingSupportedcontentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostMessagingSupportedcontentGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostMessagingSupportedcontentGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/messaging/supportedcontent][%d] postMessagingSupportedcontentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostMessagingSupportedcontentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMessagingSupportedcontentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
