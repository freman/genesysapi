// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectPromptsReader is a Reader for the PostArchitectPrompts structure.
type PostArchitectPromptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectPromptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectPromptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectPromptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectPromptsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectPromptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectPromptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostArchitectPromptsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostArchitectPromptsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectPromptsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectPromptsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectPromptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectPromptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectPromptsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectPromptsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectPromptsOK creates a PostArchitectPromptsOK with default headers values
func NewPostArchitectPromptsOK() *PostArchitectPromptsOK {
	return &PostArchitectPromptsOK{}
}

/*
PostArchitectPromptsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostArchitectPromptsOK struct {
	Payload *models.Prompt
}

// IsSuccess returns true when this post architect prompts o k response has a 2xx status code
func (o *PostArchitectPromptsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post architect prompts o k response has a 3xx status code
func (o *PostArchitectPromptsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts o k response has a 4xx status code
func (o *PostArchitectPromptsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect prompts o k response has a 5xx status code
func (o *PostArchitectPromptsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts o k response a status code equal to that given
func (o *PostArchitectPromptsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostArchitectPromptsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectPromptsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectPromptsOK) GetPayload() *models.Prompt {
	return o.Payload
}

func (o *PostArchitectPromptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prompt)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsBadRequest creates a PostArchitectPromptsBadRequest with default headers values
func NewPostArchitectPromptsBadRequest() *PostArchitectPromptsBadRequest {
	return &PostArchitectPromptsBadRequest{}
}

/*
PostArchitectPromptsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectPromptsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts bad request response has a 2xx status code
func (o *PostArchitectPromptsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts bad request response has a 3xx status code
func (o *PostArchitectPromptsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts bad request response has a 4xx status code
func (o *PostArchitectPromptsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts bad request response has a 5xx status code
func (o *PostArchitectPromptsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts bad request response a status code equal to that given
func (o *PostArchitectPromptsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostArchitectPromptsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectPromptsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectPromptsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsUnauthorized creates a PostArchitectPromptsUnauthorized with default headers values
func NewPostArchitectPromptsUnauthorized() *PostArchitectPromptsUnauthorized {
	return &PostArchitectPromptsUnauthorized{}
}

/*
PostArchitectPromptsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectPromptsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts unauthorized response has a 2xx status code
func (o *PostArchitectPromptsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts unauthorized response has a 3xx status code
func (o *PostArchitectPromptsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts unauthorized response has a 4xx status code
func (o *PostArchitectPromptsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts unauthorized response has a 5xx status code
func (o *PostArchitectPromptsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts unauthorized response a status code equal to that given
func (o *PostArchitectPromptsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostArchitectPromptsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectPromptsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectPromptsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsForbidden creates a PostArchitectPromptsForbidden with default headers values
func NewPostArchitectPromptsForbidden() *PostArchitectPromptsForbidden {
	return &PostArchitectPromptsForbidden{}
}

/*
PostArchitectPromptsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectPromptsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts forbidden response has a 2xx status code
func (o *PostArchitectPromptsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts forbidden response has a 3xx status code
func (o *PostArchitectPromptsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts forbidden response has a 4xx status code
func (o *PostArchitectPromptsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts forbidden response has a 5xx status code
func (o *PostArchitectPromptsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts forbidden response a status code equal to that given
func (o *PostArchitectPromptsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostArchitectPromptsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectPromptsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectPromptsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsNotFound creates a PostArchitectPromptsNotFound with default headers values
func NewPostArchitectPromptsNotFound() *PostArchitectPromptsNotFound {
	return &PostArchitectPromptsNotFound{}
}

/*
PostArchitectPromptsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostArchitectPromptsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts not found response has a 2xx status code
func (o *PostArchitectPromptsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts not found response has a 3xx status code
func (o *PostArchitectPromptsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts not found response has a 4xx status code
func (o *PostArchitectPromptsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts not found response has a 5xx status code
func (o *PostArchitectPromptsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts not found response a status code equal to that given
func (o *PostArchitectPromptsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostArchitectPromptsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectPromptsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectPromptsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsRequestTimeout creates a PostArchitectPromptsRequestTimeout with default headers values
func NewPostArchitectPromptsRequestTimeout() *PostArchitectPromptsRequestTimeout {
	return &PostArchitectPromptsRequestTimeout{}
}

/*
PostArchitectPromptsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostArchitectPromptsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts request timeout response has a 2xx status code
func (o *PostArchitectPromptsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts request timeout response has a 3xx status code
func (o *PostArchitectPromptsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts request timeout response has a 4xx status code
func (o *PostArchitectPromptsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts request timeout response has a 5xx status code
func (o *PostArchitectPromptsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts request timeout response a status code equal to that given
func (o *PostArchitectPromptsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostArchitectPromptsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectPromptsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectPromptsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsConflict creates a PostArchitectPromptsConflict with default headers values
func NewPostArchitectPromptsConflict() *PostArchitectPromptsConflict {
	return &PostArchitectPromptsConflict{}
}

/*
PostArchitectPromptsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostArchitectPromptsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts conflict response has a 2xx status code
func (o *PostArchitectPromptsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts conflict response has a 3xx status code
func (o *PostArchitectPromptsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts conflict response has a 4xx status code
func (o *PostArchitectPromptsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts conflict response has a 5xx status code
func (o *PostArchitectPromptsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts conflict response a status code equal to that given
func (o *PostArchitectPromptsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostArchitectPromptsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsConflict  %+v", 409, o.Payload)
}

func (o *PostArchitectPromptsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsConflict  %+v", 409, o.Payload)
}

func (o *PostArchitectPromptsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsRequestEntityTooLarge creates a PostArchitectPromptsRequestEntityTooLarge with default headers values
func NewPostArchitectPromptsRequestEntityTooLarge() *PostArchitectPromptsRequestEntityTooLarge {
	return &PostArchitectPromptsRequestEntityTooLarge{}
}

/*
PostArchitectPromptsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostArchitectPromptsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts request entity too large response has a 2xx status code
func (o *PostArchitectPromptsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts request entity too large response has a 3xx status code
func (o *PostArchitectPromptsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts request entity too large response has a 4xx status code
func (o *PostArchitectPromptsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts request entity too large response has a 5xx status code
func (o *PostArchitectPromptsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts request entity too large response a status code equal to that given
func (o *PostArchitectPromptsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostArchitectPromptsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectPromptsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectPromptsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsUnsupportedMediaType creates a PostArchitectPromptsUnsupportedMediaType with default headers values
func NewPostArchitectPromptsUnsupportedMediaType() *PostArchitectPromptsUnsupportedMediaType {
	return &PostArchitectPromptsUnsupportedMediaType{}
}

/*
PostArchitectPromptsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectPromptsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts unsupported media type response has a 2xx status code
func (o *PostArchitectPromptsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts unsupported media type response has a 3xx status code
func (o *PostArchitectPromptsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts unsupported media type response has a 4xx status code
func (o *PostArchitectPromptsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts unsupported media type response has a 5xx status code
func (o *PostArchitectPromptsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts unsupported media type response a status code equal to that given
func (o *PostArchitectPromptsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostArchitectPromptsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectPromptsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectPromptsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsTooManyRequests creates a PostArchitectPromptsTooManyRequests with default headers values
func NewPostArchitectPromptsTooManyRequests() *PostArchitectPromptsTooManyRequests {
	return &PostArchitectPromptsTooManyRequests{}
}

/*
PostArchitectPromptsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostArchitectPromptsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts too many requests response has a 2xx status code
func (o *PostArchitectPromptsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts too many requests response has a 3xx status code
func (o *PostArchitectPromptsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts too many requests response has a 4xx status code
func (o *PostArchitectPromptsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect prompts too many requests response has a 5xx status code
func (o *PostArchitectPromptsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect prompts too many requests response a status code equal to that given
func (o *PostArchitectPromptsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostArchitectPromptsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectPromptsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectPromptsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsInternalServerError creates a PostArchitectPromptsInternalServerError with default headers values
func NewPostArchitectPromptsInternalServerError() *PostArchitectPromptsInternalServerError {
	return &PostArchitectPromptsInternalServerError{}
}

/*
PostArchitectPromptsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectPromptsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts internal server error response has a 2xx status code
func (o *PostArchitectPromptsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts internal server error response has a 3xx status code
func (o *PostArchitectPromptsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts internal server error response has a 4xx status code
func (o *PostArchitectPromptsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect prompts internal server error response has a 5xx status code
func (o *PostArchitectPromptsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect prompts internal server error response a status code equal to that given
func (o *PostArchitectPromptsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostArchitectPromptsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectPromptsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectPromptsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsServiceUnavailable creates a PostArchitectPromptsServiceUnavailable with default headers values
func NewPostArchitectPromptsServiceUnavailable() *PostArchitectPromptsServiceUnavailable {
	return &PostArchitectPromptsServiceUnavailable{}
}

/*
PostArchitectPromptsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectPromptsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts service unavailable response has a 2xx status code
func (o *PostArchitectPromptsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts service unavailable response has a 3xx status code
func (o *PostArchitectPromptsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts service unavailable response has a 4xx status code
func (o *PostArchitectPromptsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect prompts service unavailable response has a 5xx status code
func (o *PostArchitectPromptsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect prompts service unavailable response a status code equal to that given
func (o *PostArchitectPromptsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostArchitectPromptsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectPromptsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectPromptsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectPromptsGatewayTimeout creates a PostArchitectPromptsGatewayTimeout with default headers values
func NewPostArchitectPromptsGatewayTimeout() *PostArchitectPromptsGatewayTimeout {
	return &PostArchitectPromptsGatewayTimeout{}
}

/*
PostArchitectPromptsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostArchitectPromptsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect prompts gateway timeout response has a 2xx status code
func (o *PostArchitectPromptsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect prompts gateway timeout response has a 3xx status code
func (o *PostArchitectPromptsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect prompts gateway timeout response has a 4xx status code
func (o *PostArchitectPromptsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect prompts gateway timeout response has a 5xx status code
func (o *PostArchitectPromptsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect prompts gateway timeout response a status code equal to that given
func (o *PostArchitectPromptsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostArchitectPromptsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectPromptsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/prompts][%d] postArchitectPromptsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectPromptsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectPromptsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
