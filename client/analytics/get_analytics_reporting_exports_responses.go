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

// GetAnalyticsReportingExportsReader is a Reader for the GetAnalyticsReportingExports structure.
type GetAnalyticsReportingExportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingExportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingExportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingExportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingExportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingExportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingExportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsReportingExportsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingExportsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingExportsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingExportsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingExportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingExportsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingExportsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingExportsOK creates a GetAnalyticsReportingExportsOK with default headers values
func NewGetAnalyticsReportingExportsOK() *GetAnalyticsReportingExportsOK {
	return &GetAnalyticsReportingExportsOK{}
}

/*
GetAnalyticsReportingExportsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAnalyticsReportingExportsOK struct {
	Payload *models.ReportingExportJobListing
}

// IsSuccess returns true when this get analytics reporting exports o k response has a 2xx status code
func (o *GetAnalyticsReportingExportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get analytics reporting exports o k response has a 3xx status code
func (o *GetAnalyticsReportingExportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports o k response has a 4xx status code
func (o *GetAnalyticsReportingExportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting exports o k response has a 5xx status code
func (o *GetAnalyticsReportingExportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports o k response a status code equal to that given
func (o *GetAnalyticsReportingExportsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAnalyticsReportingExportsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingExportsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingExportsOK) GetPayload() *models.ReportingExportJobListing {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportingExportJobListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsBadRequest creates a GetAnalyticsReportingExportsBadRequest with default headers values
func NewGetAnalyticsReportingExportsBadRequest() *GetAnalyticsReportingExportsBadRequest {
	return &GetAnalyticsReportingExportsBadRequest{}
}

/*
GetAnalyticsReportingExportsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingExportsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports bad request response has a 2xx status code
func (o *GetAnalyticsReportingExportsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports bad request response has a 3xx status code
func (o *GetAnalyticsReportingExportsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports bad request response has a 4xx status code
func (o *GetAnalyticsReportingExportsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports bad request response has a 5xx status code
func (o *GetAnalyticsReportingExportsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports bad request response a status code equal to that given
func (o *GetAnalyticsReportingExportsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAnalyticsReportingExportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingExportsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingExportsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsUnauthorized creates a GetAnalyticsReportingExportsUnauthorized with default headers values
func NewGetAnalyticsReportingExportsUnauthorized() *GetAnalyticsReportingExportsUnauthorized {
	return &GetAnalyticsReportingExportsUnauthorized{}
}

/*
GetAnalyticsReportingExportsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingExportsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports unauthorized response has a 2xx status code
func (o *GetAnalyticsReportingExportsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports unauthorized response has a 3xx status code
func (o *GetAnalyticsReportingExportsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports unauthorized response has a 4xx status code
func (o *GetAnalyticsReportingExportsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports unauthorized response has a 5xx status code
func (o *GetAnalyticsReportingExportsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports unauthorized response a status code equal to that given
func (o *GetAnalyticsReportingExportsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAnalyticsReportingExportsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingExportsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingExportsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsForbidden creates a GetAnalyticsReportingExportsForbidden with default headers values
func NewGetAnalyticsReportingExportsForbidden() *GetAnalyticsReportingExportsForbidden {
	return &GetAnalyticsReportingExportsForbidden{}
}

/*
GetAnalyticsReportingExportsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingExportsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports forbidden response has a 2xx status code
func (o *GetAnalyticsReportingExportsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports forbidden response has a 3xx status code
func (o *GetAnalyticsReportingExportsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports forbidden response has a 4xx status code
func (o *GetAnalyticsReportingExportsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports forbidden response has a 5xx status code
func (o *GetAnalyticsReportingExportsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports forbidden response a status code equal to that given
func (o *GetAnalyticsReportingExportsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAnalyticsReportingExportsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingExportsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingExportsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsNotFound creates a GetAnalyticsReportingExportsNotFound with default headers values
func NewGetAnalyticsReportingExportsNotFound() *GetAnalyticsReportingExportsNotFound {
	return &GetAnalyticsReportingExportsNotFound{}
}

/*
GetAnalyticsReportingExportsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingExportsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports not found response has a 2xx status code
func (o *GetAnalyticsReportingExportsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports not found response has a 3xx status code
func (o *GetAnalyticsReportingExportsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports not found response has a 4xx status code
func (o *GetAnalyticsReportingExportsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports not found response has a 5xx status code
func (o *GetAnalyticsReportingExportsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports not found response a status code equal to that given
func (o *GetAnalyticsReportingExportsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAnalyticsReportingExportsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingExportsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingExportsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsRequestTimeout creates a GetAnalyticsReportingExportsRequestTimeout with default headers values
func NewGetAnalyticsReportingExportsRequestTimeout() *GetAnalyticsReportingExportsRequestTimeout {
	return &GetAnalyticsReportingExportsRequestTimeout{}
}

/*
GetAnalyticsReportingExportsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsReportingExportsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports request timeout response has a 2xx status code
func (o *GetAnalyticsReportingExportsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports request timeout response has a 3xx status code
func (o *GetAnalyticsReportingExportsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports request timeout response has a 4xx status code
func (o *GetAnalyticsReportingExportsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports request timeout response has a 5xx status code
func (o *GetAnalyticsReportingExportsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports request timeout response a status code equal to that given
func (o *GetAnalyticsReportingExportsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAnalyticsReportingExportsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingExportsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingExportsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsRequestEntityTooLarge creates a GetAnalyticsReportingExportsRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingExportsRequestEntityTooLarge() *GetAnalyticsReportingExportsRequestEntityTooLarge {
	return &GetAnalyticsReportingExportsRequestEntityTooLarge{}
}

/*
GetAnalyticsReportingExportsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAnalyticsReportingExportsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports request entity too large response has a 2xx status code
func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports request entity too large response has a 3xx status code
func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports request entity too large response has a 4xx status code
func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports request entity too large response has a 5xx status code
func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports request entity too large response a status code equal to that given
func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsUnsupportedMediaType creates a GetAnalyticsReportingExportsUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingExportsUnsupportedMediaType() *GetAnalyticsReportingExportsUnsupportedMediaType {
	return &GetAnalyticsReportingExportsUnsupportedMediaType{}
}

/*
GetAnalyticsReportingExportsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingExportsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports unsupported media type response has a 2xx status code
func (o *GetAnalyticsReportingExportsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports unsupported media type response has a 3xx status code
func (o *GetAnalyticsReportingExportsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports unsupported media type response has a 4xx status code
func (o *GetAnalyticsReportingExportsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports unsupported media type response has a 5xx status code
func (o *GetAnalyticsReportingExportsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports unsupported media type response a status code equal to that given
func (o *GetAnalyticsReportingExportsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAnalyticsReportingExportsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingExportsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingExportsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsTooManyRequests creates a GetAnalyticsReportingExportsTooManyRequests with default headers values
func NewGetAnalyticsReportingExportsTooManyRequests() *GetAnalyticsReportingExportsTooManyRequests {
	return &GetAnalyticsReportingExportsTooManyRequests{}
}

/*
GetAnalyticsReportingExportsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsReportingExportsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports too many requests response has a 2xx status code
func (o *GetAnalyticsReportingExportsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports too many requests response has a 3xx status code
func (o *GetAnalyticsReportingExportsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports too many requests response has a 4xx status code
func (o *GetAnalyticsReportingExportsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting exports too many requests response has a 5xx status code
func (o *GetAnalyticsReportingExportsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting exports too many requests response a status code equal to that given
func (o *GetAnalyticsReportingExportsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAnalyticsReportingExportsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingExportsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingExportsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsInternalServerError creates a GetAnalyticsReportingExportsInternalServerError with default headers values
func NewGetAnalyticsReportingExportsInternalServerError() *GetAnalyticsReportingExportsInternalServerError {
	return &GetAnalyticsReportingExportsInternalServerError{}
}

/*
GetAnalyticsReportingExportsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingExportsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports internal server error response has a 2xx status code
func (o *GetAnalyticsReportingExportsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports internal server error response has a 3xx status code
func (o *GetAnalyticsReportingExportsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports internal server error response has a 4xx status code
func (o *GetAnalyticsReportingExportsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting exports internal server error response has a 5xx status code
func (o *GetAnalyticsReportingExportsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting exports internal server error response a status code equal to that given
func (o *GetAnalyticsReportingExportsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAnalyticsReportingExportsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingExportsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingExportsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsServiceUnavailable creates a GetAnalyticsReportingExportsServiceUnavailable with default headers values
func NewGetAnalyticsReportingExportsServiceUnavailable() *GetAnalyticsReportingExportsServiceUnavailable {
	return &GetAnalyticsReportingExportsServiceUnavailable{}
}

/*
GetAnalyticsReportingExportsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingExportsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports service unavailable response has a 2xx status code
func (o *GetAnalyticsReportingExportsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports service unavailable response has a 3xx status code
func (o *GetAnalyticsReportingExportsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports service unavailable response has a 4xx status code
func (o *GetAnalyticsReportingExportsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting exports service unavailable response has a 5xx status code
func (o *GetAnalyticsReportingExportsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting exports service unavailable response a status code equal to that given
func (o *GetAnalyticsReportingExportsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAnalyticsReportingExportsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingExportsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingExportsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingExportsGatewayTimeout creates a GetAnalyticsReportingExportsGatewayTimeout with default headers values
func NewGetAnalyticsReportingExportsGatewayTimeout() *GetAnalyticsReportingExportsGatewayTimeout {
	return &GetAnalyticsReportingExportsGatewayTimeout{}
}

/*
GetAnalyticsReportingExportsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAnalyticsReportingExportsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting exports gateway timeout response has a 2xx status code
func (o *GetAnalyticsReportingExportsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting exports gateway timeout response has a 3xx status code
func (o *GetAnalyticsReportingExportsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting exports gateway timeout response has a 4xx status code
func (o *GetAnalyticsReportingExportsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting exports gateway timeout response has a 5xx status code
func (o *GetAnalyticsReportingExportsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting exports gateway timeout response a status code equal to that given
func (o *GetAnalyticsReportingExportsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAnalyticsReportingExportsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingExportsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/exports][%d] getAnalyticsReportingExportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingExportsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingExportsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
