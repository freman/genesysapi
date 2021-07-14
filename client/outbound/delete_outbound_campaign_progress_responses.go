// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteOutboundCampaignProgressReader is a Reader for the DeleteOutboundCampaignProgress structure.
type DeleteOutboundCampaignProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundCampaignProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteOutboundCampaignProgressAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundCampaignProgressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundCampaignProgressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundCampaignProgressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundCampaignProgressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundCampaignProgressRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundCampaignProgressRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundCampaignProgressUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundCampaignProgressTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundCampaignProgressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundCampaignProgressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundCampaignProgressGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundCampaignProgressAccepted creates a DeleteOutboundCampaignProgressAccepted with default headers values
func NewDeleteOutboundCampaignProgressAccepted() *DeleteOutboundCampaignProgressAccepted {
	return &DeleteOutboundCampaignProgressAccepted{}
}

/*DeleteOutboundCampaignProgressAccepted handles this case with default header values.

Accepted - the campaign will be recycled momentarily
*/
type DeleteOutboundCampaignProgressAccepted struct {
}

func (o *DeleteOutboundCampaignProgressAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressAccepted ", 202)
}

func (o *DeleteOutboundCampaignProgressAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundCampaignProgressBadRequest creates a DeleteOutboundCampaignProgressBadRequest with default headers values
func NewDeleteOutboundCampaignProgressBadRequest() *DeleteOutboundCampaignProgressBadRequest {
	return &DeleteOutboundCampaignProgressBadRequest{}
}

/*DeleteOutboundCampaignProgressBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundCampaignProgressBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundCampaignProgressBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressUnauthorized creates a DeleteOutboundCampaignProgressUnauthorized with default headers values
func NewDeleteOutboundCampaignProgressUnauthorized() *DeleteOutboundCampaignProgressUnauthorized {
	return &DeleteOutboundCampaignProgressUnauthorized{}
}

/*DeleteOutboundCampaignProgressUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundCampaignProgressUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundCampaignProgressUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressForbidden creates a DeleteOutboundCampaignProgressForbidden with default headers values
func NewDeleteOutboundCampaignProgressForbidden() *DeleteOutboundCampaignProgressForbidden {
	return &DeleteOutboundCampaignProgressForbidden{}
}

/*DeleteOutboundCampaignProgressForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundCampaignProgressForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundCampaignProgressForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressNotFound creates a DeleteOutboundCampaignProgressNotFound with default headers values
func NewDeleteOutboundCampaignProgressNotFound() *DeleteOutboundCampaignProgressNotFound {
	return &DeleteOutboundCampaignProgressNotFound{}
}

/*DeleteOutboundCampaignProgressNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundCampaignProgressNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundCampaignProgressNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressRequestTimeout creates a DeleteOutboundCampaignProgressRequestTimeout with default headers values
func NewDeleteOutboundCampaignProgressRequestTimeout() *DeleteOutboundCampaignProgressRequestTimeout {
	return &DeleteOutboundCampaignProgressRequestTimeout{}
}

/*DeleteOutboundCampaignProgressRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundCampaignProgressRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundCampaignProgressRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressRequestEntityTooLarge creates a DeleteOutboundCampaignProgressRequestEntityTooLarge with default headers values
func NewDeleteOutboundCampaignProgressRequestEntityTooLarge() *DeleteOutboundCampaignProgressRequestEntityTooLarge {
	return &DeleteOutboundCampaignProgressRequestEntityTooLarge{}
}

/*DeleteOutboundCampaignProgressRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundCampaignProgressRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundCampaignProgressRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressUnsupportedMediaType creates a DeleteOutboundCampaignProgressUnsupportedMediaType with default headers values
func NewDeleteOutboundCampaignProgressUnsupportedMediaType() *DeleteOutboundCampaignProgressUnsupportedMediaType {
	return &DeleteOutboundCampaignProgressUnsupportedMediaType{}
}

/*DeleteOutboundCampaignProgressUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundCampaignProgressUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundCampaignProgressUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressTooManyRequests creates a DeleteOutboundCampaignProgressTooManyRequests with default headers values
func NewDeleteOutboundCampaignProgressTooManyRequests() *DeleteOutboundCampaignProgressTooManyRequests {
	return &DeleteOutboundCampaignProgressTooManyRequests{}
}

/*DeleteOutboundCampaignProgressTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundCampaignProgressTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundCampaignProgressTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressInternalServerError creates a DeleteOutboundCampaignProgressInternalServerError with default headers values
func NewDeleteOutboundCampaignProgressInternalServerError() *DeleteOutboundCampaignProgressInternalServerError {
	return &DeleteOutboundCampaignProgressInternalServerError{}
}

/*DeleteOutboundCampaignProgressInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundCampaignProgressInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundCampaignProgressInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressServiceUnavailable creates a DeleteOutboundCampaignProgressServiceUnavailable with default headers values
func NewDeleteOutboundCampaignProgressServiceUnavailable() *DeleteOutboundCampaignProgressServiceUnavailable {
	return &DeleteOutboundCampaignProgressServiceUnavailable{}
}

/*DeleteOutboundCampaignProgressServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundCampaignProgressServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundCampaignProgressServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignProgressGatewayTimeout creates a DeleteOutboundCampaignProgressGatewayTimeout with default headers values
func NewDeleteOutboundCampaignProgressGatewayTimeout() *DeleteOutboundCampaignProgressGatewayTimeout {
	return &DeleteOutboundCampaignProgressGatewayTimeout{}
}

/*DeleteOutboundCampaignProgressGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundCampaignProgressGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundCampaignProgressGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}/progress][%d] deleteOutboundCampaignProgressGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundCampaignProgressGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignProgressGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
