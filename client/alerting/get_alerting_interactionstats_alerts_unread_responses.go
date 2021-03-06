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

// GetAlertingInteractionstatsAlertsUnreadReader is a Reader for the GetAlertingInteractionstatsAlertsUnread structure.
type GetAlertingInteractionstatsAlertsUnreadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingInteractionstatsAlertsUnreadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingInteractionstatsAlertsUnreadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertingInteractionstatsAlertsUnreadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlertingInteractionstatsAlertsUnreadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertingInteractionstatsAlertsUnreadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertingInteractionstatsAlertsUnreadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAlertingInteractionstatsAlertsUnreadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertingInteractionstatsAlertsUnreadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlertingInteractionstatsAlertsUnreadServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlertingInteractionstatsAlertsUnreadGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertingInteractionstatsAlertsUnreadOK creates a GetAlertingInteractionstatsAlertsUnreadOK with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadOK() *GetAlertingInteractionstatsAlertsUnreadOK {
	return &GetAlertingInteractionstatsAlertsUnreadOK{}
}

/*GetAlertingInteractionstatsAlertsUnreadOK handles this case with default header values.

successful operation
*/
type GetAlertingInteractionstatsAlertsUnreadOK struct {
	Payload *models.UnreadMetric
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) GetPayload() *models.UnreadMetric {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnreadMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadBadRequest creates a GetAlertingInteractionstatsAlertsUnreadBadRequest with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadBadRequest() *GetAlertingInteractionstatsAlertsUnreadBadRequest {
	return &GetAlertingInteractionstatsAlertsUnreadBadRequest{}
}

/*GetAlertingInteractionstatsAlertsUnreadBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingInteractionstatsAlertsUnreadBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadUnauthorized creates a GetAlertingInteractionstatsAlertsUnreadUnauthorized with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadUnauthorized() *GetAlertingInteractionstatsAlertsUnreadUnauthorized {
	return &GetAlertingInteractionstatsAlertsUnreadUnauthorized{}
}

/*GetAlertingInteractionstatsAlertsUnreadUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingInteractionstatsAlertsUnreadUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadForbidden creates a GetAlertingInteractionstatsAlertsUnreadForbidden with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadForbidden() *GetAlertingInteractionstatsAlertsUnreadForbidden {
	return &GetAlertingInteractionstatsAlertsUnreadForbidden{}
}

/*GetAlertingInteractionstatsAlertsUnreadForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingInteractionstatsAlertsUnreadForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadNotFound creates a GetAlertingInteractionstatsAlertsUnreadNotFound with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadNotFound() *GetAlertingInteractionstatsAlertsUnreadNotFound {
	return &GetAlertingInteractionstatsAlertsUnreadNotFound{}
}

/*GetAlertingInteractionstatsAlertsUnreadNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAlertingInteractionstatsAlertsUnreadNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge creates a GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge() *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge {
	return &GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge{}
}

/*GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType creates a GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType() *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType {
	return &GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType{}
}

/*GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadTooManyRequests creates a GetAlertingInteractionstatsAlertsUnreadTooManyRequests with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadTooManyRequests() *GetAlertingInteractionstatsAlertsUnreadTooManyRequests {
	return &GetAlertingInteractionstatsAlertsUnreadTooManyRequests{}
}

/*GetAlertingInteractionstatsAlertsUnreadTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAlertingInteractionstatsAlertsUnreadTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadInternalServerError creates a GetAlertingInteractionstatsAlertsUnreadInternalServerError with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadInternalServerError() *GetAlertingInteractionstatsAlertsUnreadInternalServerError {
	return &GetAlertingInteractionstatsAlertsUnreadInternalServerError{}
}

/*GetAlertingInteractionstatsAlertsUnreadInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingInteractionstatsAlertsUnreadInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadServiceUnavailable creates a GetAlertingInteractionstatsAlertsUnreadServiceUnavailable with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadServiceUnavailable() *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable {
	return &GetAlertingInteractionstatsAlertsUnreadServiceUnavailable{}
}

/*GetAlertingInteractionstatsAlertsUnreadServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingInteractionstatsAlertsUnreadServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadGatewayTimeout creates a GetAlertingInteractionstatsAlertsUnreadGatewayTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadGatewayTimeout() *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout {
	return &GetAlertingInteractionstatsAlertsUnreadGatewayTimeout{}
}

/*GetAlertingInteractionstatsAlertsUnreadGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAlertingInteractionstatsAlertsUnreadGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
