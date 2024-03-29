// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingSkillsReader is a Reader for the GetRoutingSkills structure.
type GetRoutingSkillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSkillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSkillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSkillsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSkillsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSkillsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSkillsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSkillsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSkillsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSkillsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSkillsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSkillsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSkillsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSkillsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSkillsOK creates a GetRoutingSkillsOK with default headers values
func NewGetRoutingSkillsOK() *GetRoutingSkillsOK {
	return &GetRoutingSkillsOK{}
}

/*
GetRoutingSkillsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingSkillsOK struct {
	Payload *models.SkillEntityListing
}

// IsSuccess returns true when this get routing skills o k response has a 2xx status code
func (o *GetRoutingSkillsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing skills o k response has a 3xx status code
func (o *GetRoutingSkillsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills o k response has a 4xx status code
func (o *GetRoutingSkillsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skills o k response has a 5xx status code
func (o *GetRoutingSkillsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills o k response a status code equal to that given
func (o *GetRoutingSkillsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingSkillsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSkillsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSkillsOK) GetPayload() *models.SkillEntityListing {
	return o.Payload
}

func (o *GetRoutingSkillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SkillEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsBadRequest creates a GetRoutingSkillsBadRequest with default headers values
func NewGetRoutingSkillsBadRequest() *GetRoutingSkillsBadRequest {
	return &GetRoutingSkillsBadRequest{}
}

/*
GetRoutingSkillsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSkillsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills bad request response has a 2xx status code
func (o *GetRoutingSkillsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills bad request response has a 3xx status code
func (o *GetRoutingSkillsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills bad request response has a 4xx status code
func (o *GetRoutingSkillsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills bad request response has a 5xx status code
func (o *GetRoutingSkillsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills bad request response a status code equal to that given
func (o *GetRoutingSkillsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingSkillsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSkillsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSkillsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsUnauthorized creates a GetRoutingSkillsUnauthorized with default headers values
func NewGetRoutingSkillsUnauthorized() *GetRoutingSkillsUnauthorized {
	return &GetRoutingSkillsUnauthorized{}
}

/*
GetRoutingSkillsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSkillsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills unauthorized response has a 2xx status code
func (o *GetRoutingSkillsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills unauthorized response has a 3xx status code
func (o *GetRoutingSkillsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills unauthorized response has a 4xx status code
func (o *GetRoutingSkillsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills unauthorized response has a 5xx status code
func (o *GetRoutingSkillsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills unauthorized response a status code equal to that given
func (o *GetRoutingSkillsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingSkillsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSkillsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSkillsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsForbidden creates a GetRoutingSkillsForbidden with default headers values
func NewGetRoutingSkillsForbidden() *GetRoutingSkillsForbidden {
	return &GetRoutingSkillsForbidden{}
}

/*
GetRoutingSkillsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSkillsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills forbidden response has a 2xx status code
func (o *GetRoutingSkillsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills forbidden response has a 3xx status code
func (o *GetRoutingSkillsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills forbidden response has a 4xx status code
func (o *GetRoutingSkillsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills forbidden response has a 5xx status code
func (o *GetRoutingSkillsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills forbidden response a status code equal to that given
func (o *GetRoutingSkillsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingSkillsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSkillsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSkillsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsNotFound creates a GetRoutingSkillsNotFound with default headers values
func NewGetRoutingSkillsNotFound() *GetRoutingSkillsNotFound {
	return &GetRoutingSkillsNotFound{}
}

/*
GetRoutingSkillsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingSkillsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills not found response has a 2xx status code
func (o *GetRoutingSkillsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills not found response has a 3xx status code
func (o *GetRoutingSkillsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills not found response has a 4xx status code
func (o *GetRoutingSkillsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills not found response has a 5xx status code
func (o *GetRoutingSkillsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills not found response a status code equal to that given
func (o *GetRoutingSkillsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingSkillsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSkillsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSkillsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsRequestTimeout creates a GetRoutingSkillsRequestTimeout with default headers values
func NewGetRoutingSkillsRequestTimeout() *GetRoutingSkillsRequestTimeout {
	return &GetRoutingSkillsRequestTimeout{}
}

/*
GetRoutingSkillsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSkillsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills request timeout response has a 2xx status code
func (o *GetRoutingSkillsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills request timeout response has a 3xx status code
func (o *GetRoutingSkillsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills request timeout response has a 4xx status code
func (o *GetRoutingSkillsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills request timeout response has a 5xx status code
func (o *GetRoutingSkillsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills request timeout response a status code equal to that given
func (o *GetRoutingSkillsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingSkillsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSkillsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSkillsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsRequestEntityTooLarge creates a GetRoutingSkillsRequestEntityTooLarge with default headers values
func NewGetRoutingSkillsRequestEntityTooLarge() *GetRoutingSkillsRequestEntityTooLarge {
	return &GetRoutingSkillsRequestEntityTooLarge{}
}

/*
GetRoutingSkillsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingSkillsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills request entity too large response has a 2xx status code
func (o *GetRoutingSkillsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills request entity too large response has a 3xx status code
func (o *GetRoutingSkillsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills request entity too large response has a 4xx status code
func (o *GetRoutingSkillsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills request entity too large response has a 5xx status code
func (o *GetRoutingSkillsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills request entity too large response a status code equal to that given
func (o *GetRoutingSkillsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingSkillsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSkillsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSkillsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsUnsupportedMediaType creates a GetRoutingSkillsUnsupportedMediaType with default headers values
func NewGetRoutingSkillsUnsupportedMediaType() *GetRoutingSkillsUnsupportedMediaType {
	return &GetRoutingSkillsUnsupportedMediaType{}
}

/*
GetRoutingSkillsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSkillsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills unsupported media type response has a 2xx status code
func (o *GetRoutingSkillsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills unsupported media type response has a 3xx status code
func (o *GetRoutingSkillsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills unsupported media type response has a 4xx status code
func (o *GetRoutingSkillsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills unsupported media type response has a 5xx status code
func (o *GetRoutingSkillsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills unsupported media type response a status code equal to that given
func (o *GetRoutingSkillsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingSkillsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSkillsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSkillsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsTooManyRequests creates a GetRoutingSkillsTooManyRequests with default headers values
func NewGetRoutingSkillsTooManyRequests() *GetRoutingSkillsTooManyRequests {
	return &GetRoutingSkillsTooManyRequests{}
}

/*
GetRoutingSkillsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSkillsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills too many requests response has a 2xx status code
func (o *GetRoutingSkillsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills too many requests response has a 3xx status code
func (o *GetRoutingSkillsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills too many requests response has a 4xx status code
func (o *GetRoutingSkillsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing skills too many requests response has a 5xx status code
func (o *GetRoutingSkillsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing skills too many requests response a status code equal to that given
func (o *GetRoutingSkillsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingSkillsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSkillsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSkillsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsInternalServerError creates a GetRoutingSkillsInternalServerError with default headers values
func NewGetRoutingSkillsInternalServerError() *GetRoutingSkillsInternalServerError {
	return &GetRoutingSkillsInternalServerError{}
}

/*
GetRoutingSkillsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSkillsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills internal server error response has a 2xx status code
func (o *GetRoutingSkillsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills internal server error response has a 3xx status code
func (o *GetRoutingSkillsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills internal server error response has a 4xx status code
func (o *GetRoutingSkillsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skills internal server error response has a 5xx status code
func (o *GetRoutingSkillsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skills internal server error response a status code equal to that given
func (o *GetRoutingSkillsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingSkillsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSkillsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSkillsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsServiceUnavailable creates a GetRoutingSkillsServiceUnavailable with default headers values
func NewGetRoutingSkillsServiceUnavailable() *GetRoutingSkillsServiceUnavailable {
	return &GetRoutingSkillsServiceUnavailable{}
}

/*
GetRoutingSkillsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSkillsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills service unavailable response has a 2xx status code
func (o *GetRoutingSkillsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills service unavailable response has a 3xx status code
func (o *GetRoutingSkillsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills service unavailable response has a 4xx status code
func (o *GetRoutingSkillsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skills service unavailable response has a 5xx status code
func (o *GetRoutingSkillsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skills service unavailable response a status code equal to that given
func (o *GetRoutingSkillsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingSkillsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSkillsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSkillsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSkillsGatewayTimeout creates a GetRoutingSkillsGatewayTimeout with default headers values
func NewGetRoutingSkillsGatewayTimeout() *GetRoutingSkillsGatewayTimeout {
	return &GetRoutingSkillsGatewayTimeout{}
}

/*
GetRoutingSkillsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingSkillsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing skills gateway timeout response has a 2xx status code
func (o *GetRoutingSkillsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing skills gateway timeout response has a 3xx status code
func (o *GetRoutingSkillsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing skills gateway timeout response has a 4xx status code
func (o *GetRoutingSkillsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing skills gateway timeout response has a 5xx status code
func (o *GetRoutingSkillsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing skills gateway timeout response a status code equal to that given
func (o *GetRoutingSkillsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingSkillsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSkillsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/skills][%d] getRoutingSkillsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSkillsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSkillsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
