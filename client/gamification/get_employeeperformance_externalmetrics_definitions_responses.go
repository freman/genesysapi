// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetEmployeeperformanceExternalmetricsDefinitionsReader is a Reader for the GetEmployeeperformanceExternalmetricsDefinitions structure.
type GetEmployeeperformanceExternalmetricsDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeeperformanceExternalmetricsDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsOK creates a GetEmployeeperformanceExternalmetricsDefinitionsOK with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsOK() *GetEmployeeperformanceExternalmetricsDefinitionsOK {
	return &GetEmployeeperformanceExternalmetricsDefinitionsOK{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetEmployeeperformanceExternalmetricsDefinitionsOK struct {
	Payload *models.ExternalMetricDefinitionListing
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions o k response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions o k response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions o k response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions o k response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions o k response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) GetPayload() *models.ExternalMetricDefinitionListing {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalMetricDefinitionListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsBadRequest creates a GetEmployeeperformanceExternalmetricsDefinitionsBadRequest with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsBadRequest() *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest {
	return &GetEmployeeperformanceExternalmetricsDefinitionsBadRequest{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions bad request response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions bad request response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions bad request response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions bad request response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions bad request response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsUnauthorized creates a GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsUnauthorized() *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized {
	return &GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions unauthorized response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions unauthorized response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions unauthorized response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions unauthorized response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions unauthorized response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsForbidden creates a GetEmployeeperformanceExternalmetricsDefinitionsForbidden with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsForbidden() *GetEmployeeperformanceExternalmetricsDefinitionsForbidden {
	return &GetEmployeeperformanceExternalmetricsDefinitionsForbidden{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions forbidden response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions forbidden response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions forbidden response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions forbidden response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions forbidden response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsNotFound creates a GetEmployeeperformanceExternalmetricsDefinitionsNotFound with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsNotFound() *GetEmployeeperformanceExternalmetricsDefinitionsNotFound {
	return &GetEmployeeperformanceExternalmetricsDefinitionsNotFound{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions not found response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions not found response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions not found response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions not found response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions not found response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout creates a GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout() *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout {
	return &GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions request timeout response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions request timeout response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions request timeout response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions request timeout response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions request timeout response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge creates a GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge() *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge {
	return &GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions request entity too large response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions request entity too large response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions request entity too large response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions request entity too large response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions request entity too large response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType creates a GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType() *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType {
	return &GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions unsupported media type response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions unsupported media type response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions unsupported media type response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions unsupported media type response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions unsupported media type response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests creates a GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests() *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests {
	return &GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions too many requests response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions too many requests response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions too many requests response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions too many requests response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get employeeperformance externalmetrics definitions too many requests response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsInternalServerError creates a GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsInternalServerError() *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError {
	return &GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions internal server error response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions internal server error response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions internal server error response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions internal server error response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get employeeperformance externalmetrics definitions internal server error response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable creates a GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable() *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable {
	return &GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions service unavailable response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions service unavailable response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions service unavailable response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions service unavailable response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get employeeperformance externalmetrics definitions service unavailable response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout creates a GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout with default headers values
func NewGetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout() *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout {
	return &GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout{}
}

/*
GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get employeeperformance externalmetrics definitions gateway timeout response has a 2xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get employeeperformance externalmetrics definitions gateway timeout response has a 3xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get employeeperformance externalmetrics definitions gateway timeout response has a 4xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get employeeperformance externalmetrics definitions gateway timeout response has a 5xx status code
func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get employeeperformance externalmetrics definitions gateway timeout response a status code equal to that given
func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/employeeperformance/externalmetrics/definitions][%d] getEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEmployeeperformanceExternalmetricsDefinitionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
