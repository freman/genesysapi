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

// PostEmployeeperformanceExternalmetricsDataReader is a Reader for the PostEmployeeperformanceExternalmetricsData structure.
type PostEmployeeperformanceExternalmetricsDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEmployeeperformanceExternalmetricsDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEmployeeperformanceExternalmetricsDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostEmployeeperformanceExternalmetricsDataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostEmployeeperformanceExternalmetricsDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostEmployeeperformanceExternalmetricsDataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostEmployeeperformanceExternalmetricsDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostEmployeeperformanceExternalmetricsDataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostEmployeeperformanceExternalmetricsDataUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostEmployeeperformanceExternalmetricsDataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostEmployeeperformanceExternalmetricsDataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostEmployeeperformanceExternalmetricsDataServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostEmployeeperformanceExternalmetricsDataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostEmployeeperformanceExternalmetricsDataOK creates a PostEmployeeperformanceExternalmetricsDataOK with default headers values
func NewPostEmployeeperformanceExternalmetricsDataOK() *PostEmployeeperformanceExternalmetricsDataOK {
	return &PostEmployeeperformanceExternalmetricsDataOK{}
}

/*
PostEmployeeperformanceExternalmetricsDataOK describes a response with status code 200, with default header values.

successful operation
*/
type PostEmployeeperformanceExternalmetricsDataOK struct {
	Payload *models.ExternalMetricDataWriteResponse
}

// IsSuccess returns true when this post employeeperformance externalmetrics data o k response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post employeeperformance externalmetrics data o k response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data o k response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post employeeperformance externalmetrics data o k response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data o k response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostEmployeeperformanceExternalmetricsDataOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataOK  %+v", 200, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataOK) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataOK  %+v", 200, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataOK) GetPayload() *models.ExternalMetricDataWriteResponse {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalMetricDataWriteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataBadRequest creates a PostEmployeeperformanceExternalmetricsDataBadRequest with default headers values
func NewPostEmployeeperformanceExternalmetricsDataBadRequest() *PostEmployeeperformanceExternalmetricsDataBadRequest {
	return &PostEmployeeperformanceExternalmetricsDataBadRequest{}
}

/*
PostEmployeeperformanceExternalmetricsDataBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostEmployeeperformanceExternalmetricsDataBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data bad request response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data bad request response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data bad request response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data bad request response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data bad request response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataBadRequest  %+v", 400, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataBadRequest  %+v", 400, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataUnauthorized creates a PostEmployeeperformanceExternalmetricsDataUnauthorized with default headers values
func NewPostEmployeeperformanceExternalmetricsDataUnauthorized() *PostEmployeeperformanceExternalmetricsDataUnauthorized {
	return &PostEmployeeperformanceExternalmetricsDataUnauthorized{}
}

/*
PostEmployeeperformanceExternalmetricsDataUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostEmployeeperformanceExternalmetricsDataUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data unauthorized response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data unauthorized response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data unauthorized response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data unauthorized response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data unauthorized response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataUnauthorized  %+v", 401, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataUnauthorized  %+v", 401, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataForbidden creates a PostEmployeeperformanceExternalmetricsDataForbidden with default headers values
func NewPostEmployeeperformanceExternalmetricsDataForbidden() *PostEmployeeperformanceExternalmetricsDataForbidden {
	return &PostEmployeeperformanceExternalmetricsDataForbidden{}
}

/*
PostEmployeeperformanceExternalmetricsDataForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostEmployeeperformanceExternalmetricsDataForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data forbidden response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data forbidden response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data forbidden response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data forbidden response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data forbidden response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostEmployeeperformanceExternalmetricsDataForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataForbidden  %+v", 403, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataForbidden  %+v", 403, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataNotFound creates a PostEmployeeperformanceExternalmetricsDataNotFound with default headers values
func NewPostEmployeeperformanceExternalmetricsDataNotFound() *PostEmployeeperformanceExternalmetricsDataNotFound {
	return &PostEmployeeperformanceExternalmetricsDataNotFound{}
}

/*
PostEmployeeperformanceExternalmetricsDataNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostEmployeeperformanceExternalmetricsDataNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data not found response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data not found response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data not found response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data not found response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data not found response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostEmployeeperformanceExternalmetricsDataNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataNotFound  %+v", 404, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataNotFound  %+v", 404, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataRequestTimeout creates a PostEmployeeperformanceExternalmetricsDataRequestTimeout with default headers values
func NewPostEmployeeperformanceExternalmetricsDataRequestTimeout() *PostEmployeeperformanceExternalmetricsDataRequestTimeout {
	return &PostEmployeeperformanceExternalmetricsDataRequestTimeout{}
}

/*
PostEmployeeperformanceExternalmetricsDataRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostEmployeeperformanceExternalmetricsDataRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data request timeout response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data request timeout response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data request timeout response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data request timeout response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data request timeout response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge creates a PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge with default headers values
func NewPostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge() *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge {
	return &PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge{}
}

/*
PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data request entity too large response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data request entity too large response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data request entity too large response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data request entity too large response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data request entity too large response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataUnsupportedMediaType creates a PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType with default headers values
func NewPostEmployeeperformanceExternalmetricsDataUnsupportedMediaType() *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType {
	return &PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType{}
}

/*
PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data unsupported media type response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data unsupported media type response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data unsupported media type response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data unsupported media type response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data unsupported media type response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataTooManyRequests creates a PostEmployeeperformanceExternalmetricsDataTooManyRequests with default headers values
func NewPostEmployeeperformanceExternalmetricsDataTooManyRequests() *PostEmployeeperformanceExternalmetricsDataTooManyRequests {
	return &PostEmployeeperformanceExternalmetricsDataTooManyRequests{}
}

/*
PostEmployeeperformanceExternalmetricsDataTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostEmployeeperformanceExternalmetricsDataTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data too many requests response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data too many requests response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data too many requests response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post employeeperformance externalmetrics data too many requests response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post employeeperformance externalmetrics data too many requests response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataInternalServerError creates a PostEmployeeperformanceExternalmetricsDataInternalServerError with default headers values
func NewPostEmployeeperformanceExternalmetricsDataInternalServerError() *PostEmployeeperformanceExternalmetricsDataInternalServerError {
	return &PostEmployeeperformanceExternalmetricsDataInternalServerError{}
}

/*
PostEmployeeperformanceExternalmetricsDataInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostEmployeeperformanceExternalmetricsDataInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data internal server error response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data internal server error response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data internal server error response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post employeeperformance externalmetrics data internal server error response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post employeeperformance externalmetrics data internal server error response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataInternalServerError  %+v", 500, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataInternalServerError  %+v", 500, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataServiceUnavailable creates a PostEmployeeperformanceExternalmetricsDataServiceUnavailable with default headers values
func NewPostEmployeeperformanceExternalmetricsDataServiceUnavailable() *PostEmployeeperformanceExternalmetricsDataServiceUnavailable {
	return &PostEmployeeperformanceExternalmetricsDataServiceUnavailable{}
}

/*
PostEmployeeperformanceExternalmetricsDataServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostEmployeeperformanceExternalmetricsDataServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data service unavailable response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data service unavailable response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data service unavailable response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post employeeperformance externalmetrics data service unavailable response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post employeeperformance externalmetrics data service unavailable response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEmployeeperformanceExternalmetricsDataGatewayTimeout creates a PostEmployeeperformanceExternalmetricsDataGatewayTimeout with default headers values
func NewPostEmployeeperformanceExternalmetricsDataGatewayTimeout() *PostEmployeeperformanceExternalmetricsDataGatewayTimeout {
	return &PostEmployeeperformanceExternalmetricsDataGatewayTimeout{}
}

/*
PostEmployeeperformanceExternalmetricsDataGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostEmployeeperformanceExternalmetricsDataGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post employeeperformance externalmetrics data gateway timeout response has a 2xx status code
func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post employeeperformance externalmetrics data gateway timeout response has a 3xx status code
func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post employeeperformance externalmetrics data gateway timeout response has a 4xx status code
func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post employeeperformance externalmetrics data gateway timeout response has a 5xx status code
func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post employeeperformance externalmetrics data gateway timeout response a status code equal to that given
func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/employeeperformance/externalmetrics/data][%d] postEmployeeperformanceExternalmetricsDataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEmployeeperformanceExternalmetricsDataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
