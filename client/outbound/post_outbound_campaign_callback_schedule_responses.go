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

// PostOutboundCampaignCallbackScheduleReader is a Reader for the PostOutboundCampaignCallbackSchedule structure.
type PostOutboundCampaignCallbackScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundCampaignCallbackScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundCampaignCallbackScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundCampaignCallbackScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundCampaignCallbackScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundCampaignCallbackScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundCampaignCallbackScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundCampaignCallbackScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundCampaignCallbackScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundCampaignCallbackScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundCampaignCallbackScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundCampaignCallbackScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundCampaignCallbackScheduleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundCampaignCallbackScheduleOK creates a PostOutboundCampaignCallbackScheduleOK with default headers values
func NewPostOutboundCampaignCallbackScheduleOK() *PostOutboundCampaignCallbackScheduleOK {
	return &PostOutboundCampaignCallbackScheduleOK{}
}

/*PostOutboundCampaignCallbackScheduleOK handles this case with default header values.

successful operation
*/
type PostOutboundCampaignCallbackScheduleOK struct {
	Payload *models.ContactCallbackRequest
}

func (o *PostOutboundCampaignCallbackScheduleOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleOK  %+v", 200, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleOK) GetPayload() *models.ContactCallbackRequest {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactCallbackRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleBadRequest creates a PostOutboundCampaignCallbackScheduleBadRequest with default headers values
func NewPostOutboundCampaignCallbackScheduleBadRequest() *PostOutboundCampaignCallbackScheduleBadRequest {
	return &PostOutboundCampaignCallbackScheduleBadRequest{}
}

/*PostOutboundCampaignCallbackScheduleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundCampaignCallbackScheduleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleUnauthorized creates a PostOutboundCampaignCallbackScheduleUnauthorized with default headers values
func NewPostOutboundCampaignCallbackScheduleUnauthorized() *PostOutboundCampaignCallbackScheduleUnauthorized {
	return &PostOutboundCampaignCallbackScheduleUnauthorized{}
}

/*PostOutboundCampaignCallbackScheduleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundCampaignCallbackScheduleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleForbidden creates a PostOutboundCampaignCallbackScheduleForbidden with default headers values
func NewPostOutboundCampaignCallbackScheduleForbidden() *PostOutboundCampaignCallbackScheduleForbidden {
	return &PostOutboundCampaignCallbackScheduleForbidden{}
}

/*PostOutboundCampaignCallbackScheduleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundCampaignCallbackScheduleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleNotFound creates a PostOutboundCampaignCallbackScheduleNotFound with default headers values
func NewPostOutboundCampaignCallbackScheduleNotFound() *PostOutboundCampaignCallbackScheduleNotFound {
	return &PostOutboundCampaignCallbackScheduleNotFound{}
}

/*PostOutboundCampaignCallbackScheduleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundCampaignCallbackScheduleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleRequestEntityTooLarge creates a PostOutboundCampaignCallbackScheduleRequestEntityTooLarge with default headers values
func NewPostOutboundCampaignCallbackScheduleRequestEntityTooLarge() *PostOutboundCampaignCallbackScheduleRequestEntityTooLarge {
	return &PostOutboundCampaignCallbackScheduleRequestEntityTooLarge{}
}

/*PostOutboundCampaignCallbackScheduleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundCampaignCallbackScheduleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleUnsupportedMediaType creates a PostOutboundCampaignCallbackScheduleUnsupportedMediaType with default headers values
func NewPostOutboundCampaignCallbackScheduleUnsupportedMediaType() *PostOutboundCampaignCallbackScheduleUnsupportedMediaType {
	return &PostOutboundCampaignCallbackScheduleUnsupportedMediaType{}
}

/*PostOutboundCampaignCallbackScheduleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundCampaignCallbackScheduleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleTooManyRequests creates a PostOutboundCampaignCallbackScheduleTooManyRequests with default headers values
func NewPostOutboundCampaignCallbackScheduleTooManyRequests() *PostOutboundCampaignCallbackScheduleTooManyRequests {
	return &PostOutboundCampaignCallbackScheduleTooManyRequests{}
}

/*PostOutboundCampaignCallbackScheduleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOutboundCampaignCallbackScheduleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleInternalServerError creates a PostOutboundCampaignCallbackScheduleInternalServerError with default headers values
func NewPostOutboundCampaignCallbackScheduleInternalServerError() *PostOutboundCampaignCallbackScheduleInternalServerError {
	return &PostOutboundCampaignCallbackScheduleInternalServerError{}
}

/*PostOutboundCampaignCallbackScheduleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundCampaignCallbackScheduleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleServiceUnavailable creates a PostOutboundCampaignCallbackScheduleServiceUnavailable with default headers values
func NewPostOutboundCampaignCallbackScheduleServiceUnavailable() *PostOutboundCampaignCallbackScheduleServiceUnavailable {
	return &PostOutboundCampaignCallbackScheduleServiceUnavailable{}
}

/*PostOutboundCampaignCallbackScheduleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundCampaignCallbackScheduleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignCallbackScheduleGatewayTimeout creates a PostOutboundCampaignCallbackScheduleGatewayTimeout with default headers values
func NewPostOutboundCampaignCallbackScheduleGatewayTimeout() *PostOutboundCampaignCallbackScheduleGatewayTimeout {
	return &PostOutboundCampaignCallbackScheduleGatewayTimeout{}
}

/*PostOutboundCampaignCallbackScheduleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundCampaignCallbackScheduleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCampaignCallbackScheduleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns/{campaignId}/callback/schedule][%d] postOutboundCampaignCallbackScheduleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundCampaignCallbackScheduleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignCallbackScheduleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
