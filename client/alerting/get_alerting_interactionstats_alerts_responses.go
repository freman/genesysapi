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

// GetAlertingInteractionstatsAlertsReader is a Reader for the GetAlertingInteractionstatsAlerts structure.
type GetAlertingInteractionstatsAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingInteractionstatsAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingInteractionstatsAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertingInteractionstatsAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlertingInteractionstatsAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertingInteractionstatsAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertingInteractionstatsAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAlertingInteractionstatsAlertsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAlertingInteractionstatsAlertsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAlertingInteractionstatsAlertsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAlertingInteractionstatsAlertsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertingInteractionstatsAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlertingInteractionstatsAlertsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlertingInteractionstatsAlertsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertingInteractionstatsAlertsOK creates a GetAlertingInteractionstatsAlertsOK with default headers values
func NewGetAlertingInteractionstatsAlertsOK() *GetAlertingInteractionstatsAlertsOK {
	return &GetAlertingInteractionstatsAlertsOK{}
}

/*GetAlertingInteractionstatsAlertsOK handles this case with default header values.

successful operation
*/
type GetAlertingInteractionstatsAlertsOK struct {
	Payload *models.InteractionStatsAlertContainer
}

func (o *GetAlertingInteractionstatsAlertsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsOK) GetPayload() *models.InteractionStatsAlertContainer {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InteractionStatsAlertContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsBadRequest creates a GetAlertingInteractionstatsAlertsBadRequest with default headers values
func NewGetAlertingInteractionstatsAlertsBadRequest() *GetAlertingInteractionstatsAlertsBadRequest {
	return &GetAlertingInteractionstatsAlertsBadRequest{}
}

/*GetAlertingInteractionstatsAlertsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingInteractionstatsAlertsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnauthorized creates a GetAlertingInteractionstatsAlertsUnauthorized with default headers values
func NewGetAlertingInteractionstatsAlertsUnauthorized() *GetAlertingInteractionstatsAlertsUnauthorized {
	return &GetAlertingInteractionstatsAlertsUnauthorized{}
}

/*GetAlertingInteractionstatsAlertsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingInteractionstatsAlertsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsForbidden creates a GetAlertingInteractionstatsAlertsForbidden with default headers values
func NewGetAlertingInteractionstatsAlertsForbidden() *GetAlertingInteractionstatsAlertsForbidden {
	return &GetAlertingInteractionstatsAlertsForbidden{}
}

/*GetAlertingInteractionstatsAlertsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingInteractionstatsAlertsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsNotFound creates a GetAlertingInteractionstatsAlertsNotFound with default headers values
func NewGetAlertingInteractionstatsAlertsNotFound() *GetAlertingInteractionstatsAlertsNotFound {
	return &GetAlertingInteractionstatsAlertsNotFound{}
}

/*GetAlertingInteractionstatsAlertsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAlertingInteractionstatsAlertsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsRequestTimeout creates a GetAlertingInteractionstatsAlertsRequestTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsRequestTimeout() *GetAlertingInteractionstatsAlertsRequestTimeout {
	return &GetAlertingInteractionstatsAlertsRequestTimeout{}
}

/*GetAlertingInteractionstatsAlertsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAlertingInteractionstatsAlertsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsRequestEntityTooLarge creates a GetAlertingInteractionstatsAlertsRequestEntityTooLarge with default headers values
func NewGetAlertingInteractionstatsAlertsRequestEntityTooLarge() *GetAlertingInteractionstatsAlertsRequestEntityTooLarge {
	return &GetAlertingInteractionstatsAlertsRequestEntityTooLarge{}
}

/*GetAlertingInteractionstatsAlertsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAlertingInteractionstatsAlertsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnsupportedMediaType creates a GetAlertingInteractionstatsAlertsUnsupportedMediaType with default headers values
func NewGetAlertingInteractionstatsAlertsUnsupportedMediaType() *GetAlertingInteractionstatsAlertsUnsupportedMediaType {
	return &GetAlertingInteractionstatsAlertsUnsupportedMediaType{}
}

/*GetAlertingInteractionstatsAlertsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingInteractionstatsAlertsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsTooManyRequests creates a GetAlertingInteractionstatsAlertsTooManyRequests with default headers values
func NewGetAlertingInteractionstatsAlertsTooManyRequests() *GetAlertingInteractionstatsAlertsTooManyRequests {
	return &GetAlertingInteractionstatsAlertsTooManyRequests{}
}

/*GetAlertingInteractionstatsAlertsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAlertingInteractionstatsAlertsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsInternalServerError creates a GetAlertingInteractionstatsAlertsInternalServerError with default headers values
func NewGetAlertingInteractionstatsAlertsInternalServerError() *GetAlertingInteractionstatsAlertsInternalServerError {
	return &GetAlertingInteractionstatsAlertsInternalServerError{}
}

/*GetAlertingInteractionstatsAlertsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingInteractionstatsAlertsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsServiceUnavailable creates a GetAlertingInteractionstatsAlertsServiceUnavailable with default headers values
func NewGetAlertingInteractionstatsAlertsServiceUnavailable() *GetAlertingInteractionstatsAlertsServiceUnavailable {
	return &GetAlertingInteractionstatsAlertsServiceUnavailable{}
}

/*GetAlertingInteractionstatsAlertsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingInteractionstatsAlertsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsGatewayTimeout creates a GetAlertingInteractionstatsAlertsGatewayTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsGatewayTimeout() *GetAlertingInteractionstatsAlertsGatewayTimeout {
	return &GetAlertingInteractionstatsAlertsGatewayTimeout{}
}

/*GetAlertingInteractionstatsAlertsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAlertingInteractionstatsAlertsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
