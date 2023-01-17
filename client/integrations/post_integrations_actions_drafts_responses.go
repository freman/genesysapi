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

// PostIntegrationsActionsDraftsReader is a Reader for the PostIntegrationsActionsDrafts structure.
type PostIntegrationsActionsDraftsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionsDraftsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionsDraftsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionsDraftsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionsDraftsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionsDraftsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionsDraftsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsActionsDraftsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionsDraftsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionsDraftsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionsDraftsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionsDraftsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionsDraftsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionsDraftsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionsDraftsOK creates a PostIntegrationsActionsDraftsOK with default headers values
func NewPostIntegrationsActionsDraftsOK() *PostIntegrationsActionsDraftsOK {
	return &PostIntegrationsActionsDraftsOK{}
}

/*
PostIntegrationsActionsDraftsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostIntegrationsActionsDraftsOK struct {
	Payload *models.Action
}

// IsSuccess returns true when this post integrations actions drafts o k response has a 2xx status code
func (o *PostIntegrationsActionsDraftsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post integrations actions drafts o k response has a 3xx status code
func (o *PostIntegrationsActionsDraftsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts o k response has a 4xx status code
func (o *PostIntegrationsActionsDraftsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations actions drafts o k response has a 5xx status code
func (o *PostIntegrationsActionsDraftsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts o k response a status code equal to that given
func (o *PostIntegrationsActionsDraftsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostIntegrationsActionsDraftsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionsDraftsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionsDraftsOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsBadRequest creates a PostIntegrationsActionsDraftsBadRequest with default headers values
func NewPostIntegrationsActionsDraftsBadRequest() *PostIntegrationsActionsDraftsBadRequest {
	return &PostIntegrationsActionsDraftsBadRequest{}
}

/*
PostIntegrationsActionsDraftsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionsDraftsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts bad request response has a 2xx status code
func (o *PostIntegrationsActionsDraftsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts bad request response has a 3xx status code
func (o *PostIntegrationsActionsDraftsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts bad request response has a 4xx status code
func (o *PostIntegrationsActionsDraftsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts bad request response has a 5xx status code
func (o *PostIntegrationsActionsDraftsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts bad request response a status code equal to that given
func (o *PostIntegrationsActionsDraftsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostIntegrationsActionsDraftsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionsDraftsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionsDraftsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsUnauthorized creates a PostIntegrationsActionsDraftsUnauthorized with default headers values
func NewPostIntegrationsActionsDraftsUnauthorized() *PostIntegrationsActionsDraftsUnauthorized {
	return &PostIntegrationsActionsDraftsUnauthorized{}
}

/*
PostIntegrationsActionsDraftsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionsDraftsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts unauthorized response has a 2xx status code
func (o *PostIntegrationsActionsDraftsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts unauthorized response has a 3xx status code
func (o *PostIntegrationsActionsDraftsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts unauthorized response has a 4xx status code
func (o *PostIntegrationsActionsDraftsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts unauthorized response has a 5xx status code
func (o *PostIntegrationsActionsDraftsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts unauthorized response a status code equal to that given
func (o *PostIntegrationsActionsDraftsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostIntegrationsActionsDraftsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionsDraftsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionsDraftsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsForbidden creates a PostIntegrationsActionsDraftsForbidden with default headers values
func NewPostIntegrationsActionsDraftsForbidden() *PostIntegrationsActionsDraftsForbidden {
	return &PostIntegrationsActionsDraftsForbidden{}
}

/*
PostIntegrationsActionsDraftsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionsDraftsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts forbidden response has a 2xx status code
func (o *PostIntegrationsActionsDraftsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts forbidden response has a 3xx status code
func (o *PostIntegrationsActionsDraftsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts forbidden response has a 4xx status code
func (o *PostIntegrationsActionsDraftsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts forbidden response has a 5xx status code
func (o *PostIntegrationsActionsDraftsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts forbidden response a status code equal to that given
func (o *PostIntegrationsActionsDraftsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostIntegrationsActionsDraftsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionsDraftsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionsDraftsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsNotFound creates a PostIntegrationsActionsDraftsNotFound with default headers values
func NewPostIntegrationsActionsDraftsNotFound() *PostIntegrationsActionsDraftsNotFound {
	return &PostIntegrationsActionsDraftsNotFound{}
}

/*
PostIntegrationsActionsDraftsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionsDraftsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts not found response has a 2xx status code
func (o *PostIntegrationsActionsDraftsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts not found response has a 3xx status code
func (o *PostIntegrationsActionsDraftsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts not found response has a 4xx status code
func (o *PostIntegrationsActionsDraftsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts not found response has a 5xx status code
func (o *PostIntegrationsActionsDraftsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts not found response a status code equal to that given
func (o *PostIntegrationsActionsDraftsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostIntegrationsActionsDraftsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionsDraftsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionsDraftsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsRequestTimeout creates a PostIntegrationsActionsDraftsRequestTimeout with default headers values
func NewPostIntegrationsActionsDraftsRequestTimeout() *PostIntegrationsActionsDraftsRequestTimeout {
	return &PostIntegrationsActionsDraftsRequestTimeout{}
}

/*
PostIntegrationsActionsDraftsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsActionsDraftsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts request timeout response has a 2xx status code
func (o *PostIntegrationsActionsDraftsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts request timeout response has a 3xx status code
func (o *PostIntegrationsActionsDraftsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts request timeout response has a 4xx status code
func (o *PostIntegrationsActionsDraftsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts request timeout response has a 5xx status code
func (o *PostIntegrationsActionsDraftsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts request timeout response a status code equal to that given
func (o *PostIntegrationsActionsDraftsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostIntegrationsActionsDraftsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionsDraftsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionsDraftsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsRequestEntityTooLarge creates a PostIntegrationsActionsDraftsRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionsDraftsRequestEntityTooLarge() *PostIntegrationsActionsDraftsRequestEntityTooLarge {
	return &PostIntegrationsActionsDraftsRequestEntityTooLarge{}
}

/*
PostIntegrationsActionsDraftsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostIntegrationsActionsDraftsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts request entity too large response has a 2xx status code
func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts request entity too large response has a 3xx status code
func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts request entity too large response has a 4xx status code
func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts request entity too large response has a 5xx status code
func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts request entity too large response a status code equal to that given
func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsUnsupportedMediaType creates a PostIntegrationsActionsDraftsUnsupportedMediaType with default headers values
func NewPostIntegrationsActionsDraftsUnsupportedMediaType() *PostIntegrationsActionsDraftsUnsupportedMediaType {
	return &PostIntegrationsActionsDraftsUnsupportedMediaType{}
}

/*
PostIntegrationsActionsDraftsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionsDraftsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts unsupported media type response has a 2xx status code
func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts unsupported media type response has a 3xx status code
func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts unsupported media type response has a 4xx status code
func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts unsupported media type response has a 5xx status code
func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts unsupported media type response a status code equal to that given
func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsTooManyRequests creates a PostIntegrationsActionsDraftsTooManyRequests with default headers values
func NewPostIntegrationsActionsDraftsTooManyRequests() *PostIntegrationsActionsDraftsTooManyRequests {
	return &PostIntegrationsActionsDraftsTooManyRequests{}
}

/*
PostIntegrationsActionsDraftsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsActionsDraftsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts too many requests response has a 2xx status code
func (o *PostIntegrationsActionsDraftsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts too many requests response has a 3xx status code
func (o *PostIntegrationsActionsDraftsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts too many requests response has a 4xx status code
func (o *PostIntegrationsActionsDraftsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations actions drafts too many requests response has a 5xx status code
func (o *PostIntegrationsActionsDraftsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations actions drafts too many requests response a status code equal to that given
func (o *PostIntegrationsActionsDraftsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsInternalServerError creates a PostIntegrationsActionsDraftsInternalServerError with default headers values
func NewPostIntegrationsActionsDraftsInternalServerError() *PostIntegrationsActionsDraftsInternalServerError {
	return &PostIntegrationsActionsDraftsInternalServerError{}
}

/*
PostIntegrationsActionsDraftsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionsDraftsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts internal server error response has a 2xx status code
func (o *PostIntegrationsActionsDraftsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts internal server error response has a 3xx status code
func (o *PostIntegrationsActionsDraftsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts internal server error response has a 4xx status code
func (o *PostIntegrationsActionsDraftsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations actions drafts internal server error response has a 5xx status code
func (o *PostIntegrationsActionsDraftsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations actions drafts internal server error response a status code equal to that given
func (o *PostIntegrationsActionsDraftsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostIntegrationsActionsDraftsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionsDraftsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionsDraftsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsServiceUnavailable creates a PostIntegrationsActionsDraftsServiceUnavailable with default headers values
func NewPostIntegrationsActionsDraftsServiceUnavailable() *PostIntegrationsActionsDraftsServiceUnavailable {
	return &PostIntegrationsActionsDraftsServiceUnavailable{}
}

/*
PostIntegrationsActionsDraftsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionsDraftsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts service unavailable response has a 2xx status code
func (o *PostIntegrationsActionsDraftsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts service unavailable response has a 3xx status code
func (o *PostIntegrationsActionsDraftsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts service unavailable response has a 4xx status code
func (o *PostIntegrationsActionsDraftsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations actions drafts service unavailable response has a 5xx status code
func (o *PostIntegrationsActionsDraftsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations actions drafts service unavailable response a status code equal to that given
func (o *PostIntegrationsActionsDraftsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsGatewayTimeout creates a PostIntegrationsActionsDraftsGatewayTimeout with default headers values
func NewPostIntegrationsActionsDraftsGatewayTimeout() *PostIntegrationsActionsDraftsGatewayTimeout {
	return &PostIntegrationsActionsDraftsGatewayTimeout{}
}

/*
PostIntegrationsActionsDraftsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostIntegrationsActionsDraftsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations actions drafts gateway timeout response has a 2xx status code
func (o *PostIntegrationsActionsDraftsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations actions drafts gateway timeout response has a 3xx status code
func (o *PostIntegrationsActionsDraftsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations actions drafts gateway timeout response has a 4xx status code
func (o *PostIntegrationsActionsDraftsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations actions drafts gateway timeout response has a 5xx status code
func (o *PostIntegrationsActionsDraftsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations actions drafts gateway timeout response a status code equal to that given
func (o *PostIntegrationsActionsDraftsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
