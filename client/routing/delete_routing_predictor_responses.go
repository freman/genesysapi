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

// DeleteRoutingPredictorReader is a Reader for the DeleteRoutingPredictor structure.
type DeleteRoutingPredictorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingPredictorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingPredictorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingPredictorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingPredictorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingPredictorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingPredictorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingPredictorRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingPredictorRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingPredictorUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingPredictorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingPredictorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingPredictorServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingPredictorGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingPredictorNoContent creates a DeleteRoutingPredictorNoContent with default headers values
func NewDeleteRoutingPredictorNoContent() *DeleteRoutingPredictorNoContent {
	return &DeleteRoutingPredictorNoContent{}
}

/*
DeleteRoutingPredictorNoContent describes a response with status code 204, with default header values.

Deleted successfully
*/
type DeleteRoutingPredictorNoContent struct {
}

// IsSuccess returns true when this delete routing predictor no content response has a 2xx status code
func (o *DeleteRoutingPredictorNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing predictor no content response has a 3xx status code
func (o *DeleteRoutingPredictorNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor no content response has a 4xx status code
func (o *DeleteRoutingPredictorNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing predictor no content response has a 5xx status code
func (o *DeleteRoutingPredictorNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor no content response a status code equal to that given
func (o *DeleteRoutingPredictorNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteRoutingPredictorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorNoContent ", 204)
}

func (o *DeleteRoutingPredictorNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorNoContent ", 204)
}

func (o *DeleteRoutingPredictorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingPredictorBadRequest creates a DeleteRoutingPredictorBadRequest with default headers values
func NewDeleteRoutingPredictorBadRequest() *DeleteRoutingPredictorBadRequest {
	return &DeleteRoutingPredictorBadRequest{}
}

/*
DeleteRoutingPredictorBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingPredictorBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor bad request response has a 2xx status code
func (o *DeleteRoutingPredictorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor bad request response has a 3xx status code
func (o *DeleteRoutingPredictorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor bad request response has a 4xx status code
func (o *DeleteRoutingPredictorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor bad request response has a 5xx status code
func (o *DeleteRoutingPredictorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor bad request response a status code equal to that given
func (o *DeleteRoutingPredictorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingPredictorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingPredictorBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingPredictorBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorUnauthorized creates a DeleteRoutingPredictorUnauthorized with default headers values
func NewDeleteRoutingPredictorUnauthorized() *DeleteRoutingPredictorUnauthorized {
	return &DeleteRoutingPredictorUnauthorized{}
}

/*
DeleteRoutingPredictorUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingPredictorUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor unauthorized response has a 2xx status code
func (o *DeleteRoutingPredictorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor unauthorized response has a 3xx status code
func (o *DeleteRoutingPredictorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor unauthorized response has a 4xx status code
func (o *DeleteRoutingPredictorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor unauthorized response has a 5xx status code
func (o *DeleteRoutingPredictorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor unauthorized response a status code equal to that given
func (o *DeleteRoutingPredictorUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingPredictorUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingPredictorUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingPredictorUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorForbidden creates a DeleteRoutingPredictorForbidden with default headers values
func NewDeleteRoutingPredictorForbidden() *DeleteRoutingPredictorForbidden {
	return &DeleteRoutingPredictorForbidden{}
}

/*
DeleteRoutingPredictorForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingPredictorForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor forbidden response has a 2xx status code
func (o *DeleteRoutingPredictorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor forbidden response has a 3xx status code
func (o *DeleteRoutingPredictorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor forbidden response has a 4xx status code
func (o *DeleteRoutingPredictorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor forbidden response has a 5xx status code
func (o *DeleteRoutingPredictorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor forbidden response a status code equal to that given
func (o *DeleteRoutingPredictorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingPredictorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingPredictorForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingPredictorForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorNotFound creates a DeleteRoutingPredictorNotFound with default headers values
func NewDeleteRoutingPredictorNotFound() *DeleteRoutingPredictorNotFound {
	return &DeleteRoutingPredictorNotFound{}
}

/*
DeleteRoutingPredictorNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingPredictorNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor not found response has a 2xx status code
func (o *DeleteRoutingPredictorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor not found response has a 3xx status code
func (o *DeleteRoutingPredictorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor not found response has a 4xx status code
func (o *DeleteRoutingPredictorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor not found response has a 5xx status code
func (o *DeleteRoutingPredictorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor not found response a status code equal to that given
func (o *DeleteRoutingPredictorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingPredictorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingPredictorNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingPredictorNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorRequestTimeout creates a DeleteRoutingPredictorRequestTimeout with default headers values
func NewDeleteRoutingPredictorRequestTimeout() *DeleteRoutingPredictorRequestTimeout {
	return &DeleteRoutingPredictorRequestTimeout{}
}

/*
DeleteRoutingPredictorRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingPredictorRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor request timeout response has a 2xx status code
func (o *DeleteRoutingPredictorRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor request timeout response has a 3xx status code
func (o *DeleteRoutingPredictorRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor request timeout response has a 4xx status code
func (o *DeleteRoutingPredictorRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor request timeout response has a 5xx status code
func (o *DeleteRoutingPredictorRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor request timeout response a status code equal to that given
func (o *DeleteRoutingPredictorRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingPredictorRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingPredictorRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingPredictorRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorRequestEntityTooLarge creates a DeleteRoutingPredictorRequestEntityTooLarge with default headers values
func NewDeleteRoutingPredictorRequestEntityTooLarge() *DeleteRoutingPredictorRequestEntityTooLarge {
	return &DeleteRoutingPredictorRequestEntityTooLarge{}
}

/*
DeleteRoutingPredictorRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteRoutingPredictorRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor request entity too large response has a 2xx status code
func (o *DeleteRoutingPredictorRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor request entity too large response has a 3xx status code
func (o *DeleteRoutingPredictorRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor request entity too large response has a 4xx status code
func (o *DeleteRoutingPredictorRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor request entity too large response has a 5xx status code
func (o *DeleteRoutingPredictorRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor request entity too large response a status code equal to that given
func (o *DeleteRoutingPredictorRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingPredictorRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingPredictorRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingPredictorRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorUnsupportedMediaType creates a DeleteRoutingPredictorUnsupportedMediaType with default headers values
func NewDeleteRoutingPredictorUnsupportedMediaType() *DeleteRoutingPredictorUnsupportedMediaType {
	return &DeleteRoutingPredictorUnsupportedMediaType{}
}

/*
DeleteRoutingPredictorUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingPredictorUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor unsupported media type response has a 2xx status code
func (o *DeleteRoutingPredictorUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor unsupported media type response has a 3xx status code
func (o *DeleteRoutingPredictorUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor unsupported media type response has a 4xx status code
func (o *DeleteRoutingPredictorUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor unsupported media type response has a 5xx status code
func (o *DeleteRoutingPredictorUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor unsupported media type response a status code equal to that given
func (o *DeleteRoutingPredictorUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingPredictorUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingPredictorUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingPredictorUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorTooManyRequests creates a DeleteRoutingPredictorTooManyRequests with default headers values
func NewDeleteRoutingPredictorTooManyRequests() *DeleteRoutingPredictorTooManyRequests {
	return &DeleteRoutingPredictorTooManyRequests{}
}

/*
DeleteRoutingPredictorTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingPredictorTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor too many requests response has a 2xx status code
func (o *DeleteRoutingPredictorTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor too many requests response has a 3xx status code
func (o *DeleteRoutingPredictorTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor too many requests response has a 4xx status code
func (o *DeleteRoutingPredictorTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing predictor too many requests response has a 5xx status code
func (o *DeleteRoutingPredictorTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing predictor too many requests response a status code equal to that given
func (o *DeleteRoutingPredictorTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingPredictorTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingPredictorTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingPredictorTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorInternalServerError creates a DeleteRoutingPredictorInternalServerError with default headers values
func NewDeleteRoutingPredictorInternalServerError() *DeleteRoutingPredictorInternalServerError {
	return &DeleteRoutingPredictorInternalServerError{}
}

/*
DeleteRoutingPredictorInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingPredictorInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor internal server error response has a 2xx status code
func (o *DeleteRoutingPredictorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor internal server error response has a 3xx status code
func (o *DeleteRoutingPredictorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor internal server error response has a 4xx status code
func (o *DeleteRoutingPredictorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing predictor internal server error response has a 5xx status code
func (o *DeleteRoutingPredictorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing predictor internal server error response a status code equal to that given
func (o *DeleteRoutingPredictorInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingPredictorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingPredictorInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingPredictorInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorServiceUnavailable creates a DeleteRoutingPredictorServiceUnavailable with default headers values
func NewDeleteRoutingPredictorServiceUnavailable() *DeleteRoutingPredictorServiceUnavailable {
	return &DeleteRoutingPredictorServiceUnavailable{}
}

/*
DeleteRoutingPredictorServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingPredictorServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor service unavailable response has a 2xx status code
func (o *DeleteRoutingPredictorServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor service unavailable response has a 3xx status code
func (o *DeleteRoutingPredictorServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor service unavailable response has a 4xx status code
func (o *DeleteRoutingPredictorServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing predictor service unavailable response has a 5xx status code
func (o *DeleteRoutingPredictorServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing predictor service unavailable response a status code equal to that given
func (o *DeleteRoutingPredictorServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingPredictorServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingPredictorServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingPredictorServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingPredictorGatewayTimeout creates a DeleteRoutingPredictorGatewayTimeout with default headers values
func NewDeleteRoutingPredictorGatewayTimeout() *DeleteRoutingPredictorGatewayTimeout {
	return &DeleteRoutingPredictorGatewayTimeout{}
}

/*
DeleteRoutingPredictorGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingPredictorGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing predictor gateway timeout response has a 2xx status code
func (o *DeleteRoutingPredictorGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing predictor gateway timeout response has a 3xx status code
func (o *DeleteRoutingPredictorGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing predictor gateway timeout response has a 4xx status code
func (o *DeleteRoutingPredictorGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing predictor gateway timeout response has a 5xx status code
func (o *DeleteRoutingPredictorGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing predictor gateway timeout response a status code equal to that given
func (o *DeleteRoutingPredictorGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingPredictorGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingPredictorGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/predictors/{predictorId}][%d] deleteRoutingPredictorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingPredictorGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingPredictorGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
