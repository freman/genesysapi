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

// GetArchitectPromptResourcesReader is a Reader for the GetArchitectPromptResources structure.
type GetArchitectPromptResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectPromptResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectPromptResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectPromptResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectPromptResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectPromptResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectPromptResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectPromptResourcesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectPromptResourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectPromptResourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectPromptResourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectPromptResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectPromptResourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectPromptResourcesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectPromptResourcesOK creates a GetArchitectPromptResourcesOK with default headers values
func NewGetArchitectPromptResourcesOK() *GetArchitectPromptResourcesOK {
	return &GetArchitectPromptResourcesOK{}
}

/*
GetArchitectPromptResourcesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectPromptResourcesOK struct {
	Payload *models.PromptAssetEntityListing
}

// IsSuccess returns true when this get architect prompt resources o k response has a 2xx status code
func (o *GetArchitectPromptResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect prompt resources o k response has a 3xx status code
func (o *GetArchitectPromptResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources o k response has a 4xx status code
func (o *GetArchitectPromptResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resources o k response has a 5xx status code
func (o *GetArchitectPromptResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources o k response a status code equal to that given
func (o *GetArchitectPromptResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectPromptResourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesOK  %+v", 200, o.Payload)
}

func (o *GetArchitectPromptResourcesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesOK  %+v", 200, o.Payload)
}

func (o *GetArchitectPromptResourcesOK) GetPayload() *models.PromptAssetEntityListing {
	return o.Payload
}

func (o *GetArchitectPromptResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromptAssetEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesBadRequest creates a GetArchitectPromptResourcesBadRequest with default headers values
func NewGetArchitectPromptResourcesBadRequest() *GetArchitectPromptResourcesBadRequest {
	return &GetArchitectPromptResourcesBadRequest{}
}

/*
GetArchitectPromptResourcesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectPromptResourcesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources bad request response has a 2xx status code
func (o *GetArchitectPromptResourcesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources bad request response has a 3xx status code
func (o *GetArchitectPromptResourcesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources bad request response has a 4xx status code
func (o *GetArchitectPromptResourcesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources bad request response has a 5xx status code
func (o *GetArchitectPromptResourcesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources bad request response a status code equal to that given
func (o *GetArchitectPromptResourcesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectPromptResourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectPromptResourcesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectPromptResourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesUnauthorized creates a GetArchitectPromptResourcesUnauthorized with default headers values
func NewGetArchitectPromptResourcesUnauthorized() *GetArchitectPromptResourcesUnauthorized {
	return &GetArchitectPromptResourcesUnauthorized{}
}

/*
GetArchitectPromptResourcesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectPromptResourcesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources unauthorized response has a 2xx status code
func (o *GetArchitectPromptResourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources unauthorized response has a 3xx status code
func (o *GetArchitectPromptResourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources unauthorized response has a 4xx status code
func (o *GetArchitectPromptResourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources unauthorized response has a 5xx status code
func (o *GetArchitectPromptResourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources unauthorized response a status code equal to that given
func (o *GetArchitectPromptResourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectPromptResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectPromptResourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectPromptResourcesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesForbidden creates a GetArchitectPromptResourcesForbidden with default headers values
func NewGetArchitectPromptResourcesForbidden() *GetArchitectPromptResourcesForbidden {
	return &GetArchitectPromptResourcesForbidden{}
}

/*
GetArchitectPromptResourcesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectPromptResourcesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources forbidden response has a 2xx status code
func (o *GetArchitectPromptResourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources forbidden response has a 3xx status code
func (o *GetArchitectPromptResourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources forbidden response has a 4xx status code
func (o *GetArchitectPromptResourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources forbidden response has a 5xx status code
func (o *GetArchitectPromptResourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources forbidden response a status code equal to that given
func (o *GetArchitectPromptResourcesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectPromptResourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectPromptResourcesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectPromptResourcesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesNotFound creates a GetArchitectPromptResourcesNotFound with default headers values
func NewGetArchitectPromptResourcesNotFound() *GetArchitectPromptResourcesNotFound {
	return &GetArchitectPromptResourcesNotFound{}
}

/*
GetArchitectPromptResourcesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectPromptResourcesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources not found response has a 2xx status code
func (o *GetArchitectPromptResourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources not found response has a 3xx status code
func (o *GetArchitectPromptResourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources not found response has a 4xx status code
func (o *GetArchitectPromptResourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources not found response has a 5xx status code
func (o *GetArchitectPromptResourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources not found response a status code equal to that given
func (o *GetArchitectPromptResourcesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectPromptResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectPromptResourcesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectPromptResourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesRequestTimeout creates a GetArchitectPromptResourcesRequestTimeout with default headers values
func NewGetArchitectPromptResourcesRequestTimeout() *GetArchitectPromptResourcesRequestTimeout {
	return &GetArchitectPromptResourcesRequestTimeout{}
}

/*
GetArchitectPromptResourcesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectPromptResourcesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources request timeout response has a 2xx status code
func (o *GetArchitectPromptResourcesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources request timeout response has a 3xx status code
func (o *GetArchitectPromptResourcesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources request timeout response has a 4xx status code
func (o *GetArchitectPromptResourcesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources request timeout response has a 5xx status code
func (o *GetArchitectPromptResourcesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources request timeout response a status code equal to that given
func (o *GetArchitectPromptResourcesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectPromptResourcesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectPromptResourcesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectPromptResourcesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesRequestEntityTooLarge creates a GetArchitectPromptResourcesRequestEntityTooLarge with default headers values
func NewGetArchitectPromptResourcesRequestEntityTooLarge() *GetArchitectPromptResourcesRequestEntityTooLarge {
	return &GetArchitectPromptResourcesRequestEntityTooLarge{}
}

/*
GetArchitectPromptResourcesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetArchitectPromptResourcesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources request entity too large response has a 2xx status code
func (o *GetArchitectPromptResourcesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources request entity too large response has a 3xx status code
func (o *GetArchitectPromptResourcesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources request entity too large response has a 4xx status code
func (o *GetArchitectPromptResourcesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources request entity too large response has a 5xx status code
func (o *GetArchitectPromptResourcesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources request entity too large response a status code equal to that given
func (o *GetArchitectPromptResourcesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectPromptResourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectPromptResourcesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectPromptResourcesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesUnsupportedMediaType creates a GetArchitectPromptResourcesUnsupportedMediaType with default headers values
func NewGetArchitectPromptResourcesUnsupportedMediaType() *GetArchitectPromptResourcesUnsupportedMediaType {
	return &GetArchitectPromptResourcesUnsupportedMediaType{}
}

/*
GetArchitectPromptResourcesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectPromptResourcesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources unsupported media type response has a 2xx status code
func (o *GetArchitectPromptResourcesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources unsupported media type response has a 3xx status code
func (o *GetArchitectPromptResourcesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources unsupported media type response has a 4xx status code
func (o *GetArchitectPromptResourcesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources unsupported media type response has a 5xx status code
func (o *GetArchitectPromptResourcesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources unsupported media type response a status code equal to that given
func (o *GetArchitectPromptResourcesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectPromptResourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectPromptResourcesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectPromptResourcesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesTooManyRequests creates a GetArchitectPromptResourcesTooManyRequests with default headers values
func NewGetArchitectPromptResourcesTooManyRequests() *GetArchitectPromptResourcesTooManyRequests {
	return &GetArchitectPromptResourcesTooManyRequests{}
}

/*
GetArchitectPromptResourcesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectPromptResourcesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources too many requests response has a 2xx status code
func (o *GetArchitectPromptResourcesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources too many requests response has a 3xx status code
func (o *GetArchitectPromptResourcesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources too many requests response has a 4xx status code
func (o *GetArchitectPromptResourcesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resources too many requests response has a 5xx status code
func (o *GetArchitectPromptResourcesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resources too many requests response a status code equal to that given
func (o *GetArchitectPromptResourcesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectPromptResourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectPromptResourcesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectPromptResourcesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesInternalServerError creates a GetArchitectPromptResourcesInternalServerError with default headers values
func NewGetArchitectPromptResourcesInternalServerError() *GetArchitectPromptResourcesInternalServerError {
	return &GetArchitectPromptResourcesInternalServerError{}
}

/*
GetArchitectPromptResourcesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectPromptResourcesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources internal server error response has a 2xx status code
func (o *GetArchitectPromptResourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources internal server error response has a 3xx status code
func (o *GetArchitectPromptResourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources internal server error response has a 4xx status code
func (o *GetArchitectPromptResourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resources internal server error response has a 5xx status code
func (o *GetArchitectPromptResourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect prompt resources internal server error response a status code equal to that given
func (o *GetArchitectPromptResourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectPromptResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectPromptResourcesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectPromptResourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesServiceUnavailable creates a GetArchitectPromptResourcesServiceUnavailable with default headers values
func NewGetArchitectPromptResourcesServiceUnavailable() *GetArchitectPromptResourcesServiceUnavailable {
	return &GetArchitectPromptResourcesServiceUnavailable{}
}

/*
GetArchitectPromptResourcesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectPromptResourcesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources service unavailable response has a 2xx status code
func (o *GetArchitectPromptResourcesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources service unavailable response has a 3xx status code
func (o *GetArchitectPromptResourcesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources service unavailable response has a 4xx status code
func (o *GetArchitectPromptResourcesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resources service unavailable response has a 5xx status code
func (o *GetArchitectPromptResourcesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect prompt resources service unavailable response a status code equal to that given
func (o *GetArchitectPromptResourcesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectPromptResourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectPromptResourcesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectPromptResourcesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourcesGatewayTimeout creates a GetArchitectPromptResourcesGatewayTimeout with default headers values
func NewGetArchitectPromptResourcesGatewayTimeout() *GetArchitectPromptResourcesGatewayTimeout {
	return &GetArchitectPromptResourcesGatewayTimeout{}
}

/*
GetArchitectPromptResourcesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectPromptResourcesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resources gateway timeout response has a 2xx status code
func (o *GetArchitectPromptResourcesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resources gateway timeout response has a 3xx status code
func (o *GetArchitectPromptResourcesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resources gateway timeout response has a 4xx status code
func (o *GetArchitectPromptResourcesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resources gateway timeout response has a 5xx status code
func (o *GetArchitectPromptResourcesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect prompt resources gateway timeout response a status code equal to that given
func (o *GetArchitectPromptResourcesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectPromptResourcesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectPromptResourcesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources][%d] getArchitectPromptResourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectPromptResourcesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourcesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
