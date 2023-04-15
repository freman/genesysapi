// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostIntegrationsActionDraftPublishReader is a Reader for the PostIntegrationsActionDraftPublish structure.
type PostIntegrationsActionDraftPublishReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionDraftPublishReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionDraftPublishOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionDraftPublishBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionDraftPublishUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionDraftPublishForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionDraftPublishNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsActionDraftPublishRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionDraftPublishRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionDraftPublishUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionDraftPublishTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionDraftPublishInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionDraftPublishServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionDraftPublishGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionDraftPublishOK creates a PostIntegrationsActionDraftPublishOK with default headers values
func NewPostIntegrationsActionDraftPublishOK() *PostIntegrationsActionDraftPublishOK {
	return &PostIntegrationsActionDraftPublishOK{}
}

/*
PostIntegrationsActionDraftPublishOK describes a response with status code 200, with default header values.

successful operation
*/
type PostIntegrationsActionDraftPublishOK struct {
	Payload *models.Action
}

// IsSuccess returns true when this post integrations action draft publish o k response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post integrations action draft publish o k response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish o k response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations action draft publish o k response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish o k response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostIntegrationsActionDraftPublishOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishOK) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishBadRequest creates a PostIntegrationsActionDraftPublishBadRequest with default headers values
func NewPostIntegrationsActionDraftPublishBadRequest() *PostIntegrationsActionDraftPublishBadRequest {
	return &PostIntegrationsActionDraftPublishBadRequest{}
}

/*
PostIntegrationsActionDraftPublishBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionDraftPublishBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish bad request response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish bad request response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish bad request response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish bad request response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish bad request response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostIntegrationsActionDraftPublishBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishUnauthorized creates a PostIntegrationsActionDraftPublishUnauthorized with default headers values
func NewPostIntegrationsActionDraftPublishUnauthorized() *PostIntegrationsActionDraftPublishUnauthorized {
	return &PostIntegrationsActionDraftPublishUnauthorized{}
}

/*
PostIntegrationsActionDraftPublishUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionDraftPublishUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish unauthorized response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish unauthorized response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish unauthorized response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish unauthorized response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish unauthorized response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostIntegrationsActionDraftPublishUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishForbidden creates a PostIntegrationsActionDraftPublishForbidden with default headers values
func NewPostIntegrationsActionDraftPublishForbidden() *PostIntegrationsActionDraftPublishForbidden {
	return &PostIntegrationsActionDraftPublishForbidden{}
}

/*
PostIntegrationsActionDraftPublishForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionDraftPublishForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish forbidden response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish forbidden response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish forbidden response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish forbidden response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish forbidden response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostIntegrationsActionDraftPublishForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishNotFound creates a PostIntegrationsActionDraftPublishNotFound with default headers values
func NewPostIntegrationsActionDraftPublishNotFound() *PostIntegrationsActionDraftPublishNotFound {
	return &PostIntegrationsActionDraftPublishNotFound{}
}

/*
PostIntegrationsActionDraftPublishNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionDraftPublishNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish not found response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish not found response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish not found response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish not found response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish not found response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostIntegrationsActionDraftPublishNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishRequestTimeout creates a PostIntegrationsActionDraftPublishRequestTimeout with default headers values
func NewPostIntegrationsActionDraftPublishRequestTimeout() *PostIntegrationsActionDraftPublishRequestTimeout {
	return &PostIntegrationsActionDraftPublishRequestTimeout{}
}

/*
PostIntegrationsActionDraftPublishRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsActionDraftPublishRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish request timeout response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish request timeout response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish request timeout response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish request timeout response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish request timeout response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostIntegrationsActionDraftPublishRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishRequestEntityTooLarge creates a PostIntegrationsActionDraftPublishRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionDraftPublishRequestEntityTooLarge() *PostIntegrationsActionDraftPublishRequestEntityTooLarge {
	return &PostIntegrationsActionDraftPublishRequestEntityTooLarge{}
}

/*
PostIntegrationsActionDraftPublishRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostIntegrationsActionDraftPublishRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish request entity too large response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish request entity too large response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish request entity too large response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish request entity too large response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish request entity too large response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishUnsupportedMediaType creates a PostIntegrationsActionDraftPublishUnsupportedMediaType with default headers values
func NewPostIntegrationsActionDraftPublishUnsupportedMediaType() *PostIntegrationsActionDraftPublishUnsupportedMediaType {
	return &PostIntegrationsActionDraftPublishUnsupportedMediaType{}
}

/*
PostIntegrationsActionDraftPublishUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionDraftPublishUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish unsupported media type response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish unsupported media type response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish unsupported media type response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish unsupported media type response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish unsupported media type response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishTooManyRequests creates a PostIntegrationsActionDraftPublishTooManyRequests with default headers values
func NewPostIntegrationsActionDraftPublishTooManyRequests() *PostIntegrationsActionDraftPublishTooManyRequests {
	return &PostIntegrationsActionDraftPublishTooManyRequests{}
}

/*
PostIntegrationsActionDraftPublishTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsActionDraftPublishTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish too many requests response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish too many requests response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish too many requests response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations action draft publish too many requests response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations action draft publish too many requests response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostIntegrationsActionDraftPublishTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishInternalServerError creates a PostIntegrationsActionDraftPublishInternalServerError with default headers values
func NewPostIntegrationsActionDraftPublishInternalServerError() *PostIntegrationsActionDraftPublishInternalServerError {
	return &PostIntegrationsActionDraftPublishInternalServerError{}
}

/*
PostIntegrationsActionDraftPublishInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionDraftPublishInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish internal server error response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish internal server error response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish internal server error response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations action draft publish internal server error response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations action draft publish internal server error response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostIntegrationsActionDraftPublishInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishServiceUnavailable creates a PostIntegrationsActionDraftPublishServiceUnavailable with default headers values
func NewPostIntegrationsActionDraftPublishServiceUnavailable() *PostIntegrationsActionDraftPublishServiceUnavailable {
	return &PostIntegrationsActionDraftPublishServiceUnavailable{}
}

/*
PostIntegrationsActionDraftPublishServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionDraftPublishServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish service unavailable response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish service unavailable response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish service unavailable response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations action draft publish service unavailable response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations action draft publish service unavailable response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostIntegrationsActionDraftPublishServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftPublishGatewayTimeout creates a PostIntegrationsActionDraftPublishGatewayTimeout with default headers values
func NewPostIntegrationsActionDraftPublishGatewayTimeout() *PostIntegrationsActionDraftPublishGatewayTimeout {
	return &PostIntegrationsActionDraftPublishGatewayTimeout{}
}

/*
PostIntegrationsActionDraftPublishGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostIntegrationsActionDraftPublishGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations action draft publish gateway timeout response has a 2xx status code
func (o *PostIntegrationsActionDraftPublishGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations action draft publish gateway timeout response has a 3xx status code
func (o *PostIntegrationsActionDraftPublishGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations action draft publish gateway timeout response has a 4xx status code
func (o *PostIntegrationsActionDraftPublishGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations action draft publish gateway timeout response has a 5xx status code
func (o *PostIntegrationsActionDraftPublishGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations action draft publish gateway timeout response a status code equal to that given
func (o *PostIntegrationsActionDraftPublishGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostIntegrationsActionDraftPublishGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft/publish][%d] postIntegrationsActionDraftPublishGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionDraftPublishGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftPublishGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
