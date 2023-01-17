// Code generated by go-swagger; DO NOT EDIT.

package alerting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteAlertingInteractionstatsAlertReader is a Reader for the DeleteAlertingInteractionstatsAlert structure.
type DeleteAlertingInteractionstatsAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertingInteractionstatsAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAlertingInteractionstatsAlertNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAlertingInteractionstatsAlertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAlertingInteractionstatsAlertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAlertingInteractionstatsAlertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAlertingInteractionstatsAlertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteAlertingInteractionstatsAlertRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteAlertingInteractionstatsAlertRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteAlertingInteractionstatsAlertUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAlertingInteractionstatsAlertTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAlertingInteractionstatsAlertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAlertingInteractionstatsAlertServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteAlertingInteractionstatsAlertGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAlertingInteractionstatsAlertNoContent creates a DeleteAlertingInteractionstatsAlertNoContent with default headers values
func NewDeleteAlertingInteractionstatsAlertNoContent() *DeleteAlertingInteractionstatsAlertNoContent {
	return &DeleteAlertingInteractionstatsAlertNoContent{}
}

/*
DeleteAlertingInteractionstatsAlertNoContent describes a response with status code 204, with default header values.

Interaction stats alert deleted
*/
type DeleteAlertingInteractionstatsAlertNoContent struct {
}

// IsSuccess returns true when this delete alerting interactionstats alert no content response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete alerting interactionstats alert no content response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert no content response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alerting interactionstats alert no content response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert no content response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteAlertingInteractionstatsAlertNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertNoContent ", 204)
}

func (o *DeleteAlertingInteractionstatsAlertNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertNoContent ", 204)
}

func (o *DeleteAlertingInteractionstatsAlertNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAlertingInteractionstatsAlertBadRequest creates a DeleteAlertingInteractionstatsAlertBadRequest with default headers values
func NewDeleteAlertingInteractionstatsAlertBadRequest() *DeleteAlertingInteractionstatsAlertBadRequest {
	return &DeleteAlertingInteractionstatsAlertBadRequest{}
}

/*
DeleteAlertingInteractionstatsAlertBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteAlertingInteractionstatsAlertBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert bad request response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert bad request response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert bad request response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert bad request response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert bad request response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteAlertingInteractionstatsAlertBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertUnauthorized creates a DeleteAlertingInteractionstatsAlertUnauthorized with default headers values
func NewDeleteAlertingInteractionstatsAlertUnauthorized() *DeleteAlertingInteractionstatsAlertUnauthorized {
	return &DeleteAlertingInteractionstatsAlertUnauthorized{}
}

/*
DeleteAlertingInteractionstatsAlertUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteAlertingInteractionstatsAlertUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert unauthorized response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert unauthorized response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert unauthorized response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert unauthorized response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert unauthorized response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteAlertingInteractionstatsAlertUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertForbidden creates a DeleteAlertingInteractionstatsAlertForbidden with default headers values
func NewDeleteAlertingInteractionstatsAlertForbidden() *DeleteAlertingInteractionstatsAlertForbidden {
	return &DeleteAlertingInteractionstatsAlertForbidden{}
}

/*
DeleteAlertingInteractionstatsAlertForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteAlertingInteractionstatsAlertForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert forbidden response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert forbidden response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert forbidden response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert forbidden response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert forbidden response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteAlertingInteractionstatsAlertForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertNotFound creates a DeleteAlertingInteractionstatsAlertNotFound with default headers values
func NewDeleteAlertingInteractionstatsAlertNotFound() *DeleteAlertingInteractionstatsAlertNotFound {
	return &DeleteAlertingInteractionstatsAlertNotFound{}
}

/*
DeleteAlertingInteractionstatsAlertNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteAlertingInteractionstatsAlertNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert not found response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert not found response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert not found response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert not found response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert not found response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteAlertingInteractionstatsAlertNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertRequestTimeout creates a DeleteAlertingInteractionstatsAlertRequestTimeout with default headers values
func NewDeleteAlertingInteractionstatsAlertRequestTimeout() *DeleteAlertingInteractionstatsAlertRequestTimeout {
	return &DeleteAlertingInteractionstatsAlertRequestTimeout{}
}

/*
DeleteAlertingInteractionstatsAlertRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteAlertingInteractionstatsAlertRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert request timeout response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert request timeout response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert request timeout response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert request timeout response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert request timeout response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertRequestEntityTooLarge creates a DeleteAlertingInteractionstatsAlertRequestEntityTooLarge with default headers values
func NewDeleteAlertingInteractionstatsAlertRequestEntityTooLarge() *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge {
	return &DeleteAlertingInteractionstatsAlertRequestEntityTooLarge{}
}

/*
DeleteAlertingInteractionstatsAlertRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteAlertingInteractionstatsAlertRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert request entity too large response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert request entity too large response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert request entity too large response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert request entity too large response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert request entity too large response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertUnsupportedMediaType creates a DeleteAlertingInteractionstatsAlertUnsupportedMediaType with default headers values
func NewDeleteAlertingInteractionstatsAlertUnsupportedMediaType() *DeleteAlertingInteractionstatsAlertUnsupportedMediaType {
	return &DeleteAlertingInteractionstatsAlertUnsupportedMediaType{}
}

/*
DeleteAlertingInteractionstatsAlertUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteAlertingInteractionstatsAlertUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert unsupported media type response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert unsupported media type response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert unsupported media type response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert unsupported media type response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert unsupported media type response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertTooManyRequests creates a DeleteAlertingInteractionstatsAlertTooManyRequests with default headers values
func NewDeleteAlertingInteractionstatsAlertTooManyRequests() *DeleteAlertingInteractionstatsAlertTooManyRequests {
	return &DeleteAlertingInteractionstatsAlertTooManyRequests{}
}

/*
DeleteAlertingInteractionstatsAlertTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteAlertingInteractionstatsAlertTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert too many requests response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert too many requests response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert too many requests response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alerting interactionstats alert too many requests response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alerting interactionstats alert too many requests response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertInternalServerError creates a DeleteAlertingInteractionstatsAlertInternalServerError with default headers values
func NewDeleteAlertingInteractionstatsAlertInternalServerError() *DeleteAlertingInteractionstatsAlertInternalServerError {
	return &DeleteAlertingInteractionstatsAlertInternalServerError{}
}

/*
DeleteAlertingInteractionstatsAlertInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteAlertingInteractionstatsAlertInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert internal server error response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert internal server error response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert internal server error response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alerting interactionstats alert internal server error response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete alerting interactionstats alert internal server error response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteAlertingInteractionstatsAlertInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertServiceUnavailable creates a DeleteAlertingInteractionstatsAlertServiceUnavailable with default headers values
func NewDeleteAlertingInteractionstatsAlertServiceUnavailable() *DeleteAlertingInteractionstatsAlertServiceUnavailable {
	return &DeleteAlertingInteractionstatsAlertServiceUnavailable{}
}

/*
DeleteAlertingInteractionstatsAlertServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteAlertingInteractionstatsAlertServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert service unavailable response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert service unavailable response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert service unavailable response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alerting interactionstats alert service unavailable response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete alerting interactionstats alert service unavailable response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsAlertGatewayTimeout creates a DeleteAlertingInteractionstatsAlertGatewayTimeout with default headers values
func NewDeleteAlertingInteractionstatsAlertGatewayTimeout() *DeleteAlertingInteractionstatsAlertGatewayTimeout {
	return &DeleteAlertingInteractionstatsAlertGatewayTimeout{}
}

/*
DeleteAlertingInteractionstatsAlertGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteAlertingInteractionstatsAlertGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete alerting interactionstats alert gateway timeout response has a 2xx status code
func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alerting interactionstats alert gateway timeout response has a 3xx status code
func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alerting interactionstats alert gateway timeout response has a 4xx status code
func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alerting interactionstats alert gateway timeout response has a 5xx status code
func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete alerting interactionstats alert gateway timeout response a status code equal to that given
func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
