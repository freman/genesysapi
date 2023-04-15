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

// GetIntegrationsActionsDraftsReader is a Reader for the GetIntegrationsActionsDrafts structure.
type GetIntegrationsActionsDraftsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionsDraftsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionsDraftsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionsDraftsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionsDraftsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionsDraftsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionsDraftsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsActionsDraftsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionsDraftsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionsDraftsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionsDraftsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionsDraftsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionsDraftsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionsDraftsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionsDraftsOK creates a GetIntegrationsActionsDraftsOK with default headers values
func NewGetIntegrationsActionsDraftsOK() *GetIntegrationsActionsDraftsOK {
	return &GetIntegrationsActionsDraftsOK{}
}

/*
GetIntegrationsActionsDraftsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsActionsDraftsOK struct {
	Payload *models.ActionEntityListing
}

// IsSuccess returns true when this get integrations actions drafts o k response has a 2xx status code
func (o *GetIntegrationsActionsDraftsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations actions drafts o k response has a 3xx status code
func (o *GetIntegrationsActionsDraftsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts o k response has a 4xx status code
func (o *GetIntegrationsActionsDraftsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions drafts o k response has a 5xx status code
func (o *GetIntegrationsActionsDraftsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts o k response a status code equal to that given
func (o *GetIntegrationsActionsDraftsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsActionsDraftsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionsDraftsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionsDraftsOK) GetPayload() *models.ActionEntityListing {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsBadRequest creates a GetIntegrationsActionsDraftsBadRequest with default headers values
func NewGetIntegrationsActionsDraftsBadRequest() *GetIntegrationsActionsDraftsBadRequest {
	return &GetIntegrationsActionsDraftsBadRequest{}
}

/*
GetIntegrationsActionsDraftsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionsDraftsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts bad request response has a 2xx status code
func (o *GetIntegrationsActionsDraftsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts bad request response has a 3xx status code
func (o *GetIntegrationsActionsDraftsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts bad request response has a 4xx status code
func (o *GetIntegrationsActionsDraftsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts bad request response has a 5xx status code
func (o *GetIntegrationsActionsDraftsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts bad request response a status code equal to that given
func (o *GetIntegrationsActionsDraftsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsActionsDraftsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionsDraftsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionsDraftsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsUnauthorized creates a GetIntegrationsActionsDraftsUnauthorized with default headers values
func NewGetIntegrationsActionsDraftsUnauthorized() *GetIntegrationsActionsDraftsUnauthorized {
	return &GetIntegrationsActionsDraftsUnauthorized{}
}

/*
GetIntegrationsActionsDraftsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionsDraftsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts unauthorized response has a 2xx status code
func (o *GetIntegrationsActionsDraftsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts unauthorized response has a 3xx status code
func (o *GetIntegrationsActionsDraftsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts unauthorized response has a 4xx status code
func (o *GetIntegrationsActionsDraftsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts unauthorized response has a 5xx status code
func (o *GetIntegrationsActionsDraftsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts unauthorized response a status code equal to that given
func (o *GetIntegrationsActionsDraftsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsActionsDraftsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionsDraftsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionsDraftsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsForbidden creates a GetIntegrationsActionsDraftsForbidden with default headers values
func NewGetIntegrationsActionsDraftsForbidden() *GetIntegrationsActionsDraftsForbidden {
	return &GetIntegrationsActionsDraftsForbidden{}
}

/*
GetIntegrationsActionsDraftsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionsDraftsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts forbidden response has a 2xx status code
func (o *GetIntegrationsActionsDraftsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts forbidden response has a 3xx status code
func (o *GetIntegrationsActionsDraftsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts forbidden response has a 4xx status code
func (o *GetIntegrationsActionsDraftsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts forbidden response has a 5xx status code
func (o *GetIntegrationsActionsDraftsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts forbidden response a status code equal to that given
func (o *GetIntegrationsActionsDraftsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsActionsDraftsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionsDraftsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionsDraftsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsNotFound creates a GetIntegrationsActionsDraftsNotFound with default headers values
func NewGetIntegrationsActionsDraftsNotFound() *GetIntegrationsActionsDraftsNotFound {
	return &GetIntegrationsActionsDraftsNotFound{}
}

/*
GetIntegrationsActionsDraftsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionsDraftsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts not found response has a 2xx status code
func (o *GetIntegrationsActionsDraftsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts not found response has a 3xx status code
func (o *GetIntegrationsActionsDraftsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts not found response has a 4xx status code
func (o *GetIntegrationsActionsDraftsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts not found response has a 5xx status code
func (o *GetIntegrationsActionsDraftsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts not found response a status code equal to that given
func (o *GetIntegrationsActionsDraftsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsActionsDraftsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionsDraftsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionsDraftsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsRequestTimeout creates a GetIntegrationsActionsDraftsRequestTimeout with default headers values
func NewGetIntegrationsActionsDraftsRequestTimeout() *GetIntegrationsActionsDraftsRequestTimeout {
	return &GetIntegrationsActionsDraftsRequestTimeout{}
}

/*
GetIntegrationsActionsDraftsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsActionsDraftsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts request timeout response has a 2xx status code
func (o *GetIntegrationsActionsDraftsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts request timeout response has a 3xx status code
func (o *GetIntegrationsActionsDraftsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts request timeout response has a 4xx status code
func (o *GetIntegrationsActionsDraftsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts request timeout response has a 5xx status code
func (o *GetIntegrationsActionsDraftsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts request timeout response a status code equal to that given
func (o *GetIntegrationsActionsDraftsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsActionsDraftsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionsDraftsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionsDraftsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsRequestEntityTooLarge creates a GetIntegrationsActionsDraftsRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionsDraftsRequestEntityTooLarge() *GetIntegrationsActionsDraftsRequestEntityTooLarge {
	return &GetIntegrationsActionsDraftsRequestEntityTooLarge{}
}

/*
GetIntegrationsActionsDraftsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsActionsDraftsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts request entity too large response has a 2xx status code
func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts request entity too large response has a 3xx status code
func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts request entity too large response has a 4xx status code
func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts request entity too large response has a 5xx status code
func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts request entity too large response a status code equal to that given
func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsUnsupportedMediaType creates a GetIntegrationsActionsDraftsUnsupportedMediaType with default headers values
func NewGetIntegrationsActionsDraftsUnsupportedMediaType() *GetIntegrationsActionsDraftsUnsupportedMediaType {
	return &GetIntegrationsActionsDraftsUnsupportedMediaType{}
}

/*
GetIntegrationsActionsDraftsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionsDraftsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts unsupported media type response has a 2xx status code
func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts unsupported media type response has a 3xx status code
func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts unsupported media type response has a 4xx status code
func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts unsupported media type response has a 5xx status code
func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts unsupported media type response a status code equal to that given
func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsTooManyRequests creates a GetIntegrationsActionsDraftsTooManyRequests with default headers values
func NewGetIntegrationsActionsDraftsTooManyRequests() *GetIntegrationsActionsDraftsTooManyRequests {
	return &GetIntegrationsActionsDraftsTooManyRequests{}
}

/*
GetIntegrationsActionsDraftsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsActionsDraftsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts too many requests response has a 2xx status code
func (o *GetIntegrationsActionsDraftsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts too many requests response has a 3xx status code
func (o *GetIntegrationsActionsDraftsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts too many requests response has a 4xx status code
func (o *GetIntegrationsActionsDraftsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions drafts too many requests response has a 5xx status code
func (o *GetIntegrationsActionsDraftsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions drafts too many requests response a status code equal to that given
func (o *GetIntegrationsActionsDraftsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsActionsDraftsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionsDraftsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionsDraftsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsInternalServerError creates a GetIntegrationsActionsDraftsInternalServerError with default headers values
func NewGetIntegrationsActionsDraftsInternalServerError() *GetIntegrationsActionsDraftsInternalServerError {
	return &GetIntegrationsActionsDraftsInternalServerError{}
}

/*
GetIntegrationsActionsDraftsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionsDraftsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts internal server error response has a 2xx status code
func (o *GetIntegrationsActionsDraftsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts internal server error response has a 3xx status code
func (o *GetIntegrationsActionsDraftsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts internal server error response has a 4xx status code
func (o *GetIntegrationsActionsDraftsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions drafts internal server error response has a 5xx status code
func (o *GetIntegrationsActionsDraftsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations actions drafts internal server error response a status code equal to that given
func (o *GetIntegrationsActionsDraftsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsActionsDraftsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionsDraftsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionsDraftsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsServiceUnavailable creates a GetIntegrationsActionsDraftsServiceUnavailable with default headers values
func NewGetIntegrationsActionsDraftsServiceUnavailable() *GetIntegrationsActionsDraftsServiceUnavailable {
	return &GetIntegrationsActionsDraftsServiceUnavailable{}
}

/*
GetIntegrationsActionsDraftsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionsDraftsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts service unavailable response has a 2xx status code
func (o *GetIntegrationsActionsDraftsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts service unavailable response has a 3xx status code
func (o *GetIntegrationsActionsDraftsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts service unavailable response has a 4xx status code
func (o *GetIntegrationsActionsDraftsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions drafts service unavailable response has a 5xx status code
func (o *GetIntegrationsActionsDraftsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations actions drafts service unavailable response a status code equal to that given
func (o *GetIntegrationsActionsDraftsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsActionsDraftsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionsDraftsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionsDraftsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsDraftsGatewayTimeout creates a GetIntegrationsActionsDraftsGatewayTimeout with default headers values
func NewGetIntegrationsActionsDraftsGatewayTimeout() *GetIntegrationsActionsDraftsGatewayTimeout {
	return &GetIntegrationsActionsDraftsGatewayTimeout{}
}

/*
GetIntegrationsActionsDraftsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsActionsDraftsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions drafts gateway timeout response has a 2xx status code
func (o *GetIntegrationsActionsDraftsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions drafts gateway timeout response has a 3xx status code
func (o *GetIntegrationsActionsDraftsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions drafts gateway timeout response has a 4xx status code
func (o *GetIntegrationsActionsDraftsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions drafts gateway timeout response has a 5xx status code
func (o *GetIntegrationsActionsDraftsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations actions drafts gateway timeout response a status code equal to that given
func (o *GetIntegrationsActionsDraftsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsActionsDraftsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionsDraftsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/drafts][%d] getIntegrationsActionsDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionsDraftsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsDraftsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
