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

/*DeleteAlertingInteractionstatsAlertNoContent handles this case with default header values.

Interaction stats alert deleted
*/
type DeleteAlertingInteractionstatsAlertNoContent struct {
}

func (o *DeleteAlertingInteractionstatsAlertNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/alerts/{alertId}][%d] deleteAlertingInteractionstatsAlertNoContent ", 204)
}

func (o *DeleteAlertingInteractionstatsAlertNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAlertingInteractionstatsAlertBadRequest creates a DeleteAlertingInteractionstatsAlertBadRequest with default headers values
func NewDeleteAlertingInteractionstatsAlertBadRequest() *DeleteAlertingInteractionstatsAlertBadRequest {
	return &DeleteAlertingInteractionstatsAlertBadRequest{}
}

/*DeleteAlertingInteractionstatsAlertBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteAlertingInteractionstatsAlertBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertBadRequest) Error() string {
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

/*DeleteAlertingInteractionstatsAlertUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteAlertingInteractionstatsAlertUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertUnauthorized) Error() string {
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

/*DeleteAlertingInteractionstatsAlertForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteAlertingInteractionstatsAlertForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertForbidden) Error() string {
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

/*DeleteAlertingInteractionstatsAlertNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteAlertingInteractionstatsAlertNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertNotFound) Error() string {
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

/*DeleteAlertingInteractionstatsAlertRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteAlertingInteractionstatsAlertRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertRequestTimeout) Error() string {
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

/*DeleteAlertingInteractionstatsAlertRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteAlertingInteractionstatsAlertRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertRequestEntityTooLarge) Error() string {
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

/*DeleteAlertingInteractionstatsAlertUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteAlertingInteractionstatsAlertUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertUnsupportedMediaType) Error() string {
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

/*DeleteAlertingInteractionstatsAlertTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteAlertingInteractionstatsAlertTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertTooManyRequests) Error() string {
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

/*DeleteAlertingInteractionstatsAlertInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteAlertingInteractionstatsAlertInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertInternalServerError) Error() string {
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

/*DeleteAlertingInteractionstatsAlertServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteAlertingInteractionstatsAlertServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertServiceUnavailable) Error() string {
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

/*DeleteAlertingInteractionstatsAlertGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteAlertingInteractionstatsAlertGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsAlertGatewayTimeout) Error() string {
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
