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

// GetRoutingPredictorReader is a Reader for the GetRoutingPredictor structure.
type GetRoutingPredictorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingPredictorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingPredictorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingPredictorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingPredictorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingPredictorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingPredictorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingPredictorRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingPredictorRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingPredictorUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingPredictorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingPredictorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingPredictorServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingPredictorGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingPredictorOK creates a GetRoutingPredictorOK with default headers values
func NewGetRoutingPredictorOK() *GetRoutingPredictorOK {
	return &GetRoutingPredictorOK{}
}

/*
GetRoutingPredictorOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingPredictorOK struct {
	Payload *models.Predictor
}

// IsSuccess returns true when this get routing predictor o k response has a 2xx status code
func (o *GetRoutingPredictorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing predictor o k response has a 3xx status code
func (o *GetRoutingPredictorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor o k response has a 4xx status code
func (o *GetRoutingPredictorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor o k response has a 5xx status code
func (o *GetRoutingPredictorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor o k response a status code equal to that given
func (o *GetRoutingPredictorOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingPredictorOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorOK  %+v", 200, o.Payload)
}

func (o *GetRoutingPredictorOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorOK  %+v", 200, o.Payload)
}

func (o *GetRoutingPredictorOK) GetPayload() *models.Predictor {
	return o.Payload
}

func (o *GetRoutingPredictorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Predictor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorBadRequest creates a GetRoutingPredictorBadRequest with default headers values
func NewGetRoutingPredictorBadRequest() *GetRoutingPredictorBadRequest {
	return &GetRoutingPredictorBadRequest{}
}

/*
GetRoutingPredictorBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingPredictorBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor bad request response has a 2xx status code
func (o *GetRoutingPredictorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor bad request response has a 3xx status code
func (o *GetRoutingPredictorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor bad request response has a 4xx status code
func (o *GetRoutingPredictorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor bad request response has a 5xx status code
func (o *GetRoutingPredictorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor bad request response a status code equal to that given
func (o *GetRoutingPredictorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingPredictorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingPredictorBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingPredictorBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorUnauthorized creates a GetRoutingPredictorUnauthorized with default headers values
func NewGetRoutingPredictorUnauthorized() *GetRoutingPredictorUnauthorized {
	return &GetRoutingPredictorUnauthorized{}
}

/*
GetRoutingPredictorUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingPredictorUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor unauthorized response has a 2xx status code
func (o *GetRoutingPredictorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor unauthorized response has a 3xx status code
func (o *GetRoutingPredictorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor unauthorized response has a 4xx status code
func (o *GetRoutingPredictorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor unauthorized response has a 5xx status code
func (o *GetRoutingPredictorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor unauthorized response a status code equal to that given
func (o *GetRoutingPredictorUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingPredictorUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingPredictorUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingPredictorUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorForbidden creates a GetRoutingPredictorForbidden with default headers values
func NewGetRoutingPredictorForbidden() *GetRoutingPredictorForbidden {
	return &GetRoutingPredictorForbidden{}
}

/*
GetRoutingPredictorForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingPredictorForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor forbidden response has a 2xx status code
func (o *GetRoutingPredictorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor forbidden response has a 3xx status code
func (o *GetRoutingPredictorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor forbidden response has a 4xx status code
func (o *GetRoutingPredictorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor forbidden response has a 5xx status code
func (o *GetRoutingPredictorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor forbidden response a status code equal to that given
func (o *GetRoutingPredictorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingPredictorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingPredictorForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingPredictorForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorNotFound creates a GetRoutingPredictorNotFound with default headers values
func NewGetRoutingPredictorNotFound() *GetRoutingPredictorNotFound {
	return &GetRoutingPredictorNotFound{}
}

/*
GetRoutingPredictorNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingPredictorNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor not found response has a 2xx status code
func (o *GetRoutingPredictorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor not found response has a 3xx status code
func (o *GetRoutingPredictorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor not found response has a 4xx status code
func (o *GetRoutingPredictorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor not found response has a 5xx status code
func (o *GetRoutingPredictorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor not found response a status code equal to that given
func (o *GetRoutingPredictorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingPredictorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingPredictorNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingPredictorNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorRequestTimeout creates a GetRoutingPredictorRequestTimeout with default headers values
func NewGetRoutingPredictorRequestTimeout() *GetRoutingPredictorRequestTimeout {
	return &GetRoutingPredictorRequestTimeout{}
}

/*
GetRoutingPredictorRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingPredictorRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor request timeout response has a 2xx status code
func (o *GetRoutingPredictorRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor request timeout response has a 3xx status code
func (o *GetRoutingPredictorRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor request timeout response has a 4xx status code
func (o *GetRoutingPredictorRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor request timeout response has a 5xx status code
func (o *GetRoutingPredictorRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor request timeout response a status code equal to that given
func (o *GetRoutingPredictorRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingPredictorRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingPredictorRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingPredictorRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorRequestEntityTooLarge creates a GetRoutingPredictorRequestEntityTooLarge with default headers values
func NewGetRoutingPredictorRequestEntityTooLarge() *GetRoutingPredictorRequestEntityTooLarge {
	return &GetRoutingPredictorRequestEntityTooLarge{}
}

/*
GetRoutingPredictorRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingPredictorRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor request entity too large response has a 2xx status code
func (o *GetRoutingPredictorRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor request entity too large response has a 3xx status code
func (o *GetRoutingPredictorRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor request entity too large response has a 4xx status code
func (o *GetRoutingPredictorRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor request entity too large response has a 5xx status code
func (o *GetRoutingPredictorRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor request entity too large response a status code equal to that given
func (o *GetRoutingPredictorRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingPredictorRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingPredictorRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingPredictorRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorUnsupportedMediaType creates a GetRoutingPredictorUnsupportedMediaType with default headers values
func NewGetRoutingPredictorUnsupportedMediaType() *GetRoutingPredictorUnsupportedMediaType {
	return &GetRoutingPredictorUnsupportedMediaType{}
}

/*
GetRoutingPredictorUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingPredictorUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor unsupported media type response has a 2xx status code
func (o *GetRoutingPredictorUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor unsupported media type response has a 3xx status code
func (o *GetRoutingPredictorUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor unsupported media type response has a 4xx status code
func (o *GetRoutingPredictorUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor unsupported media type response has a 5xx status code
func (o *GetRoutingPredictorUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor unsupported media type response a status code equal to that given
func (o *GetRoutingPredictorUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingPredictorUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingPredictorUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingPredictorUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorTooManyRequests creates a GetRoutingPredictorTooManyRequests with default headers values
func NewGetRoutingPredictorTooManyRequests() *GetRoutingPredictorTooManyRequests {
	return &GetRoutingPredictorTooManyRequests{}
}

/*
GetRoutingPredictorTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingPredictorTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor too many requests response has a 2xx status code
func (o *GetRoutingPredictorTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor too many requests response has a 3xx status code
func (o *GetRoutingPredictorTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor too many requests response has a 4xx status code
func (o *GetRoutingPredictorTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor too many requests response has a 5xx status code
func (o *GetRoutingPredictorTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor too many requests response a status code equal to that given
func (o *GetRoutingPredictorTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingPredictorTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingPredictorTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingPredictorTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorInternalServerError creates a GetRoutingPredictorInternalServerError with default headers values
func NewGetRoutingPredictorInternalServerError() *GetRoutingPredictorInternalServerError {
	return &GetRoutingPredictorInternalServerError{}
}

/*
GetRoutingPredictorInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingPredictorInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor internal server error response has a 2xx status code
func (o *GetRoutingPredictorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor internal server error response has a 3xx status code
func (o *GetRoutingPredictorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor internal server error response has a 4xx status code
func (o *GetRoutingPredictorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor internal server error response has a 5xx status code
func (o *GetRoutingPredictorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing predictor internal server error response a status code equal to that given
func (o *GetRoutingPredictorInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingPredictorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingPredictorInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingPredictorInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorServiceUnavailable creates a GetRoutingPredictorServiceUnavailable with default headers values
func NewGetRoutingPredictorServiceUnavailable() *GetRoutingPredictorServiceUnavailable {
	return &GetRoutingPredictorServiceUnavailable{}
}

/*
GetRoutingPredictorServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingPredictorServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor service unavailable response has a 2xx status code
func (o *GetRoutingPredictorServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor service unavailable response has a 3xx status code
func (o *GetRoutingPredictorServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor service unavailable response has a 4xx status code
func (o *GetRoutingPredictorServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor service unavailable response has a 5xx status code
func (o *GetRoutingPredictorServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing predictor service unavailable response a status code equal to that given
func (o *GetRoutingPredictorServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingPredictorServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingPredictorServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingPredictorServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorGatewayTimeout creates a GetRoutingPredictorGatewayTimeout with default headers values
func NewGetRoutingPredictorGatewayTimeout() *GetRoutingPredictorGatewayTimeout {
	return &GetRoutingPredictorGatewayTimeout{}
}

/*
GetRoutingPredictorGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingPredictorGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor gateway timeout response has a 2xx status code
func (o *GetRoutingPredictorGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor gateway timeout response has a 3xx status code
func (o *GetRoutingPredictorGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor gateway timeout response has a 4xx status code
func (o *GetRoutingPredictorGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor gateway timeout response has a 5xx status code
func (o *GetRoutingPredictorGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing predictor gateway timeout response a status code equal to that given
func (o *GetRoutingPredictorGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingPredictorGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingPredictorGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}][%d] getRoutingPredictorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingPredictorGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
