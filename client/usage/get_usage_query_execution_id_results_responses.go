// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUsageQueryExecutionIDResultsReader is a Reader for the GetUsageQueryExecutionIDResults structure.
type GetUsageQueryExecutionIDResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsageQueryExecutionIDResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsageQueryExecutionIDResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsageQueryExecutionIDResultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsageQueryExecutionIDResultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsageQueryExecutionIDResultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsageQueryExecutionIDResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUsageQueryExecutionIDResultsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUsageQueryExecutionIDResultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUsageQueryExecutionIDResultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsageQueryExecutionIDResultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsageQueryExecutionIDResultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUsageQueryExecutionIDResultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUsageQueryExecutionIDResultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsageQueryExecutionIDResultsOK creates a GetUsageQueryExecutionIDResultsOK with default headers values
func NewGetUsageQueryExecutionIDResultsOK() *GetUsageQueryExecutionIDResultsOK {
	return &GetUsageQueryExecutionIDResultsOK{}
}

/*
GetUsageQueryExecutionIDResultsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUsageQueryExecutionIDResultsOK struct {
	Payload *models.APIUsageQueryResult
}

// IsSuccess returns true when this get usage query execution Id results o k response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get usage query execution Id results o k response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results o k response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get usage query execution Id results o k response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results o k response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUsageQueryExecutionIDResultsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsOK  %+v", 200, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsOK  %+v", 200, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsOK) GetPayload() *models.APIUsageQueryResult {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUsageQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsBadRequest creates a GetUsageQueryExecutionIDResultsBadRequest with default headers values
func NewGetUsageQueryExecutionIDResultsBadRequest() *GetUsageQueryExecutionIDResultsBadRequest {
	return &GetUsageQueryExecutionIDResultsBadRequest{}
}

/*
GetUsageQueryExecutionIDResultsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUsageQueryExecutionIDResultsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results bad request response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results bad request response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results bad request response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results bad request response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results bad request response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsUnauthorized creates a GetUsageQueryExecutionIDResultsUnauthorized with default headers values
func NewGetUsageQueryExecutionIDResultsUnauthorized() *GetUsageQueryExecutionIDResultsUnauthorized {
	return &GetUsageQueryExecutionIDResultsUnauthorized{}
}

/*
GetUsageQueryExecutionIDResultsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUsageQueryExecutionIDResultsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results unauthorized response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results unauthorized response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results unauthorized response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results unauthorized response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results unauthorized response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsForbidden creates a GetUsageQueryExecutionIDResultsForbidden with default headers values
func NewGetUsageQueryExecutionIDResultsForbidden() *GetUsageQueryExecutionIDResultsForbidden {
	return &GetUsageQueryExecutionIDResultsForbidden{}
}

/*
GetUsageQueryExecutionIDResultsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUsageQueryExecutionIDResultsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results forbidden response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results forbidden response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results forbidden response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results forbidden response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results forbidden response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUsageQueryExecutionIDResultsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsForbidden  %+v", 403, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsForbidden  %+v", 403, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsNotFound creates a GetUsageQueryExecutionIDResultsNotFound with default headers values
func NewGetUsageQueryExecutionIDResultsNotFound() *GetUsageQueryExecutionIDResultsNotFound {
	return &GetUsageQueryExecutionIDResultsNotFound{}
}

/*
GetUsageQueryExecutionIDResultsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUsageQueryExecutionIDResultsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results not found response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results not found response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results not found response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results not found response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results not found response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUsageQueryExecutionIDResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsNotFound  %+v", 404, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsNotFound  %+v", 404, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsRequestTimeout creates a GetUsageQueryExecutionIDResultsRequestTimeout with default headers values
func NewGetUsageQueryExecutionIDResultsRequestTimeout() *GetUsageQueryExecutionIDResultsRequestTimeout {
	return &GetUsageQueryExecutionIDResultsRequestTimeout{}
}

/*
GetUsageQueryExecutionIDResultsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUsageQueryExecutionIDResultsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results request timeout response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results request timeout response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results request timeout response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results request timeout response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results request timeout response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUsageQueryExecutionIDResultsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsRequestEntityTooLarge creates a GetUsageQueryExecutionIDResultsRequestEntityTooLarge with default headers values
func NewGetUsageQueryExecutionIDResultsRequestEntityTooLarge() *GetUsageQueryExecutionIDResultsRequestEntityTooLarge {
	return &GetUsageQueryExecutionIDResultsRequestEntityTooLarge{}
}

/*
GetUsageQueryExecutionIDResultsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUsageQueryExecutionIDResultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results request entity too large response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results request entity too large response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results request entity too large response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results request entity too large response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results request entity too large response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsUnsupportedMediaType creates a GetUsageQueryExecutionIDResultsUnsupportedMediaType with default headers values
func NewGetUsageQueryExecutionIDResultsUnsupportedMediaType() *GetUsageQueryExecutionIDResultsUnsupportedMediaType {
	return &GetUsageQueryExecutionIDResultsUnsupportedMediaType{}
}

/*
GetUsageQueryExecutionIDResultsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUsageQueryExecutionIDResultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results unsupported media type response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results unsupported media type response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results unsupported media type response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results unsupported media type response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results unsupported media type response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsTooManyRequests creates a GetUsageQueryExecutionIDResultsTooManyRequests with default headers values
func NewGetUsageQueryExecutionIDResultsTooManyRequests() *GetUsageQueryExecutionIDResultsTooManyRequests {
	return &GetUsageQueryExecutionIDResultsTooManyRequests{}
}

/*
GetUsageQueryExecutionIDResultsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUsageQueryExecutionIDResultsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results too many requests response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results too many requests response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results too many requests response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get usage query execution Id results too many requests response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get usage query execution Id results too many requests response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsInternalServerError creates a GetUsageQueryExecutionIDResultsInternalServerError with default headers values
func NewGetUsageQueryExecutionIDResultsInternalServerError() *GetUsageQueryExecutionIDResultsInternalServerError {
	return &GetUsageQueryExecutionIDResultsInternalServerError{}
}

/*
GetUsageQueryExecutionIDResultsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUsageQueryExecutionIDResultsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results internal server error response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results internal server error response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results internal server error response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get usage query execution Id results internal server error response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get usage query execution Id results internal server error response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsServiceUnavailable creates a GetUsageQueryExecutionIDResultsServiceUnavailable with default headers values
func NewGetUsageQueryExecutionIDResultsServiceUnavailable() *GetUsageQueryExecutionIDResultsServiceUnavailable {
	return &GetUsageQueryExecutionIDResultsServiceUnavailable{}
}

/*
GetUsageQueryExecutionIDResultsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUsageQueryExecutionIDResultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results service unavailable response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results service unavailable response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results service unavailable response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get usage query execution Id results service unavailable response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get usage query execution Id results service unavailable response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsGatewayTimeout creates a GetUsageQueryExecutionIDResultsGatewayTimeout with default headers values
func NewGetUsageQueryExecutionIDResultsGatewayTimeout() *GetUsageQueryExecutionIDResultsGatewayTimeout {
	return &GetUsageQueryExecutionIDResultsGatewayTimeout{}
}

/*
GetUsageQueryExecutionIDResultsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUsageQueryExecutionIDResultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get usage query execution Id results gateway timeout response has a 2xx status code
func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get usage query execution Id results gateway timeout response has a 3xx status code
func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get usage query execution Id results gateway timeout response has a 4xx status code
func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get usage query execution Id results gateway timeout response has a 5xx status code
func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get usage query execution Id results gateway timeout response a status code equal to that given
func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
