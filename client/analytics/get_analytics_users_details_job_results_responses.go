// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsUsersDetailsJobResultsReader is a Reader for the GetAnalyticsUsersDetailsJobResults structure.
type GetAnalyticsUsersDetailsJobResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsUsersDetailsJobResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsUsersDetailsJobResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsUsersDetailsJobResultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsUsersDetailsJobResultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsUsersDetailsJobResultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsUsersDetailsJobResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsUsersDetailsJobResultsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsUsersDetailsJobResultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsUsersDetailsJobResultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsUsersDetailsJobResultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsUsersDetailsJobResultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsUsersDetailsJobResultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsUsersDetailsJobResultsOK creates a GetAnalyticsUsersDetailsJobResultsOK with default headers values
func NewGetAnalyticsUsersDetailsJobResultsOK() *GetAnalyticsUsersDetailsJobResultsOK {
	return &GetAnalyticsUsersDetailsJobResultsOK{}
}

/*
GetAnalyticsUsersDetailsJobResultsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAnalyticsUsersDetailsJobResultsOK struct {
	Payload *models.AnalyticsUserDetailsAsyncQueryResponse
}

// IsSuccess returns true when this get analytics users details job results o k response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get analytics users details job results o k response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results o k response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics users details job results o k response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results o k response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAnalyticsUsersDetailsJobResultsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsOK) GetPayload() *models.AnalyticsUserDetailsAsyncQueryResponse {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsUserDetailsAsyncQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsBadRequest creates a GetAnalyticsUsersDetailsJobResultsBadRequest with default headers values
func NewGetAnalyticsUsersDetailsJobResultsBadRequest() *GetAnalyticsUsersDetailsJobResultsBadRequest {
	return &GetAnalyticsUsersDetailsJobResultsBadRequest{}
}

/*
GetAnalyticsUsersDetailsJobResultsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsUsersDetailsJobResultsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results bad request response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results bad request response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results bad request response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results bad request response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results bad request response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsUnauthorized creates a GetAnalyticsUsersDetailsJobResultsUnauthorized with default headers values
func NewGetAnalyticsUsersDetailsJobResultsUnauthorized() *GetAnalyticsUsersDetailsJobResultsUnauthorized {
	return &GetAnalyticsUsersDetailsJobResultsUnauthorized{}
}

/*
GetAnalyticsUsersDetailsJobResultsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsUsersDetailsJobResultsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results unauthorized response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results unauthorized response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results unauthorized response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results unauthorized response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results unauthorized response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsForbidden creates a GetAnalyticsUsersDetailsJobResultsForbidden with default headers values
func NewGetAnalyticsUsersDetailsJobResultsForbidden() *GetAnalyticsUsersDetailsJobResultsForbidden {
	return &GetAnalyticsUsersDetailsJobResultsForbidden{}
}

/*
GetAnalyticsUsersDetailsJobResultsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsUsersDetailsJobResultsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results forbidden response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results forbidden response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results forbidden response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results forbidden response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results forbidden response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAnalyticsUsersDetailsJobResultsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsNotFound creates a GetAnalyticsUsersDetailsJobResultsNotFound with default headers values
func NewGetAnalyticsUsersDetailsJobResultsNotFound() *GetAnalyticsUsersDetailsJobResultsNotFound {
	return &GetAnalyticsUsersDetailsJobResultsNotFound{}
}

/*
GetAnalyticsUsersDetailsJobResultsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAnalyticsUsersDetailsJobResultsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results not found response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results not found response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results not found response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results not found response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results not found response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAnalyticsUsersDetailsJobResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsRequestTimeout creates a GetAnalyticsUsersDetailsJobResultsRequestTimeout with default headers values
func NewGetAnalyticsUsersDetailsJobResultsRequestTimeout() *GetAnalyticsUsersDetailsJobResultsRequestTimeout {
	return &GetAnalyticsUsersDetailsJobResultsRequestTimeout{}
}

/*
GetAnalyticsUsersDetailsJobResultsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsUsersDetailsJobResultsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results request timeout response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results request timeout response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results request timeout response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results request timeout response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results request timeout response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge creates a GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge with default headers values
func NewGetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge() *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge {
	return &GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge{}
}

/*
GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results request entity too large response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results request entity too large response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results request entity too large response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results request entity too large response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results request entity too large response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsUnsupportedMediaType creates a GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType with default headers values
func NewGetAnalyticsUsersDetailsJobResultsUnsupportedMediaType() *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType {
	return &GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType{}
}

/*
GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results unsupported media type response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results unsupported media type response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results unsupported media type response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results unsupported media type response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results unsupported media type response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsTooManyRequests creates a GetAnalyticsUsersDetailsJobResultsTooManyRequests with default headers values
func NewGetAnalyticsUsersDetailsJobResultsTooManyRequests() *GetAnalyticsUsersDetailsJobResultsTooManyRequests {
	return &GetAnalyticsUsersDetailsJobResultsTooManyRequests{}
}

/*
GetAnalyticsUsersDetailsJobResultsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsUsersDetailsJobResultsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results too many requests response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results too many requests response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results too many requests response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics users details job results too many requests response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics users details job results too many requests response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsInternalServerError creates a GetAnalyticsUsersDetailsJobResultsInternalServerError with default headers values
func NewGetAnalyticsUsersDetailsJobResultsInternalServerError() *GetAnalyticsUsersDetailsJobResultsInternalServerError {
	return &GetAnalyticsUsersDetailsJobResultsInternalServerError{}
}

/*
GetAnalyticsUsersDetailsJobResultsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsUsersDetailsJobResultsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results internal server error response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results internal server error response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results internal server error response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics users details job results internal server error response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics users details job results internal server error response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsServiceUnavailable creates a GetAnalyticsUsersDetailsJobResultsServiceUnavailable with default headers values
func NewGetAnalyticsUsersDetailsJobResultsServiceUnavailable() *GetAnalyticsUsersDetailsJobResultsServiceUnavailable {
	return &GetAnalyticsUsersDetailsJobResultsServiceUnavailable{}
}

/*
GetAnalyticsUsersDetailsJobResultsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsUsersDetailsJobResultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results service unavailable response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results service unavailable response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results service unavailable response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics users details job results service unavailable response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics users details job results service unavailable response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobResultsGatewayTimeout creates a GetAnalyticsUsersDetailsJobResultsGatewayTimeout with default headers values
func NewGetAnalyticsUsersDetailsJobResultsGatewayTimeout() *GetAnalyticsUsersDetailsJobResultsGatewayTimeout {
	return &GetAnalyticsUsersDetailsJobResultsGatewayTimeout{}
}

/*
GetAnalyticsUsersDetailsJobResultsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAnalyticsUsersDetailsJobResultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics users details job results gateway timeout response has a 2xx status code
func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics users details job results gateway timeout response has a 3xx status code
func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics users details job results gateway timeout response has a 4xx status code
func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics users details job results gateway timeout response has a 5xx status code
func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics users details job results gateway timeout response a status code equal to that given
func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}/results][%d] getAnalyticsUsersDetailsJobResultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobResultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
