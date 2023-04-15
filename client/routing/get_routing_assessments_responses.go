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

// GetRoutingAssessmentsReader is a Reader for the GetRoutingAssessments structure.
type GetRoutingAssessmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingAssessmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingAssessmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingAssessmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingAssessmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingAssessmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingAssessmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingAssessmentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingAssessmentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingAssessmentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingAssessmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingAssessmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingAssessmentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingAssessmentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingAssessmentsOK creates a GetRoutingAssessmentsOK with default headers values
func NewGetRoutingAssessmentsOK() *GetRoutingAssessmentsOK {
	return &GetRoutingAssessmentsOK{}
}

/*
GetRoutingAssessmentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingAssessmentsOK struct {
	Payload *models.AssessmentListing
}

// IsSuccess returns true when this get routing assessments o k response has a 2xx status code
func (o *GetRoutingAssessmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing assessments o k response has a 3xx status code
func (o *GetRoutingAssessmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments o k response has a 4xx status code
func (o *GetRoutingAssessmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing assessments o k response has a 5xx status code
func (o *GetRoutingAssessmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments o k response a status code equal to that given
func (o *GetRoutingAssessmentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingAssessmentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingAssessmentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingAssessmentsOK) GetPayload() *models.AssessmentListing {
	return o.Payload
}

func (o *GetRoutingAssessmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssessmentListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsBadRequest creates a GetRoutingAssessmentsBadRequest with default headers values
func NewGetRoutingAssessmentsBadRequest() *GetRoutingAssessmentsBadRequest {
	return &GetRoutingAssessmentsBadRequest{}
}

/*
GetRoutingAssessmentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingAssessmentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments bad request response has a 2xx status code
func (o *GetRoutingAssessmentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments bad request response has a 3xx status code
func (o *GetRoutingAssessmentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments bad request response has a 4xx status code
func (o *GetRoutingAssessmentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments bad request response has a 5xx status code
func (o *GetRoutingAssessmentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments bad request response a status code equal to that given
func (o *GetRoutingAssessmentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingAssessmentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingAssessmentsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingAssessmentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsUnauthorized creates a GetRoutingAssessmentsUnauthorized with default headers values
func NewGetRoutingAssessmentsUnauthorized() *GetRoutingAssessmentsUnauthorized {
	return &GetRoutingAssessmentsUnauthorized{}
}

/*
GetRoutingAssessmentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingAssessmentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments unauthorized response has a 2xx status code
func (o *GetRoutingAssessmentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments unauthorized response has a 3xx status code
func (o *GetRoutingAssessmentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments unauthorized response has a 4xx status code
func (o *GetRoutingAssessmentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments unauthorized response has a 5xx status code
func (o *GetRoutingAssessmentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments unauthorized response a status code equal to that given
func (o *GetRoutingAssessmentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingAssessmentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingAssessmentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingAssessmentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsForbidden creates a GetRoutingAssessmentsForbidden with default headers values
func NewGetRoutingAssessmentsForbidden() *GetRoutingAssessmentsForbidden {
	return &GetRoutingAssessmentsForbidden{}
}

/*
GetRoutingAssessmentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingAssessmentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments forbidden response has a 2xx status code
func (o *GetRoutingAssessmentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments forbidden response has a 3xx status code
func (o *GetRoutingAssessmentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments forbidden response has a 4xx status code
func (o *GetRoutingAssessmentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments forbidden response has a 5xx status code
func (o *GetRoutingAssessmentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments forbidden response a status code equal to that given
func (o *GetRoutingAssessmentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingAssessmentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingAssessmentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingAssessmentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsNotFound creates a GetRoutingAssessmentsNotFound with default headers values
func NewGetRoutingAssessmentsNotFound() *GetRoutingAssessmentsNotFound {
	return &GetRoutingAssessmentsNotFound{}
}

/*
GetRoutingAssessmentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingAssessmentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments not found response has a 2xx status code
func (o *GetRoutingAssessmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments not found response has a 3xx status code
func (o *GetRoutingAssessmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments not found response has a 4xx status code
func (o *GetRoutingAssessmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments not found response has a 5xx status code
func (o *GetRoutingAssessmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments not found response a status code equal to that given
func (o *GetRoutingAssessmentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingAssessmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingAssessmentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingAssessmentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsRequestTimeout creates a GetRoutingAssessmentsRequestTimeout with default headers values
func NewGetRoutingAssessmentsRequestTimeout() *GetRoutingAssessmentsRequestTimeout {
	return &GetRoutingAssessmentsRequestTimeout{}
}

/*
GetRoutingAssessmentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingAssessmentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments request timeout response has a 2xx status code
func (o *GetRoutingAssessmentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments request timeout response has a 3xx status code
func (o *GetRoutingAssessmentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments request timeout response has a 4xx status code
func (o *GetRoutingAssessmentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments request timeout response has a 5xx status code
func (o *GetRoutingAssessmentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments request timeout response a status code equal to that given
func (o *GetRoutingAssessmentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingAssessmentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingAssessmentsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingAssessmentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsRequestEntityTooLarge creates a GetRoutingAssessmentsRequestEntityTooLarge with default headers values
func NewGetRoutingAssessmentsRequestEntityTooLarge() *GetRoutingAssessmentsRequestEntityTooLarge {
	return &GetRoutingAssessmentsRequestEntityTooLarge{}
}

/*
GetRoutingAssessmentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingAssessmentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments request entity too large response has a 2xx status code
func (o *GetRoutingAssessmentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments request entity too large response has a 3xx status code
func (o *GetRoutingAssessmentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments request entity too large response has a 4xx status code
func (o *GetRoutingAssessmentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments request entity too large response has a 5xx status code
func (o *GetRoutingAssessmentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments request entity too large response a status code equal to that given
func (o *GetRoutingAssessmentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingAssessmentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingAssessmentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingAssessmentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsUnsupportedMediaType creates a GetRoutingAssessmentsUnsupportedMediaType with default headers values
func NewGetRoutingAssessmentsUnsupportedMediaType() *GetRoutingAssessmentsUnsupportedMediaType {
	return &GetRoutingAssessmentsUnsupportedMediaType{}
}

/*
GetRoutingAssessmentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingAssessmentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments unsupported media type response has a 2xx status code
func (o *GetRoutingAssessmentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments unsupported media type response has a 3xx status code
func (o *GetRoutingAssessmentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments unsupported media type response has a 4xx status code
func (o *GetRoutingAssessmentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments unsupported media type response has a 5xx status code
func (o *GetRoutingAssessmentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments unsupported media type response a status code equal to that given
func (o *GetRoutingAssessmentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingAssessmentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingAssessmentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingAssessmentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsTooManyRequests creates a GetRoutingAssessmentsTooManyRequests with default headers values
func NewGetRoutingAssessmentsTooManyRequests() *GetRoutingAssessmentsTooManyRequests {
	return &GetRoutingAssessmentsTooManyRequests{}
}

/*
GetRoutingAssessmentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingAssessmentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments too many requests response has a 2xx status code
func (o *GetRoutingAssessmentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments too many requests response has a 3xx status code
func (o *GetRoutingAssessmentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments too many requests response has a 4xx status code
func (o *GetRoutingAssessmentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing assessments too many requests response has a 5xx status code
func (o *GetRoutingAssessmentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing assessments too many requests response a status code equal to that given
func (o *GetRoutingAssessmentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingAssessmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingAssessmentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingAssessmentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsInternalServerError creates a GetRoutingAssessmentsInternalServerError with default headers values
func NewGetRoutingAssessmentsInternalServerError() *GetRoutingAssessmentsInternalServerError {
	return &GetRoutingAssessmentsInternalServerError{}
}

/*
GetRoutingAssessmentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingAssessmentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments internal server error response has a 2xx status code
func (o *GetRoutingAssessmentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments internal server error response has a 3xx status code
func (o *GetRoutingAssessmentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments internal server error response has a 4xx status code
func (o *GetRoutingAssessmentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing assessments internal server error response has a 5xx status code
func (o *GetRoutingAssessmentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing assessments internal server error response a status code equal to that given
func (o *GetRoutingAssessmentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingAssessmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingAssessmentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingAssessmentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsServiceUnavailable creates a GetRoutingAssessmentsServiceUnavailable with default headers values
func NewGetRoutingAssessmentsServiceUnavailable() *GetRoutingAssessmentsServiceUnavailable {
	return &GetRoutingAssessmentsServiceUnavailable{}
}

/*
GetRoutingAssessmentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingAssessmentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments service unavailable response has a 2xx status code
func (o *GetRoutingAssessmentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments service unavailable response has a 3xx status code
func (o *GetRoutingAssessmentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments service unavailable response has a 4xx status code
func (o *GetRoutingAssessmentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing assessments service unavailable response has a 5xx status code
func (o *GetRoutingAssessmentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing assessments service unavailable response a status code equal to that given
func (o *GetRoutingAssessmentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingAssessmentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingAssessmentsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingAssessmentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingAssessmentsGatewayTimeout creates a GetRoutingAssessmentsGatewayTimeout with default headers values
func NewGetRoutingAssessmentsGatewayTimeout() *GetRoutingAssessmentsGatewayTimeout {
	return &GetRoutingAssessmentsGatewayTimeout{}
}

/*
GetRoutingAssessmentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingAssessmentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing assessments gateway timeout response has a 2xx status code
func (o *GetRoutingAssessmentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing assessments gateway timeout response has a 3xx status code
func (o *GetRoutingAssessmentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing assessments gateway timeout response has a 4xx status code
func (o *GetRoutingAssessmentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing assessments gateway timeout response has a 5xx status code
func (o *GetRoutingAssessmentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing assessments gateway timeout response a status code equal to that given
func (o *GetRoutingAssessmentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingAssessmentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingAssessmentsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/assessments][%d] getRoutingAssessmentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingAssessmentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingAssessmentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
