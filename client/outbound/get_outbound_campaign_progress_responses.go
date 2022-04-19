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

// GetOutboundCampaignProgressReader is a Reader for the GetOutboundCampaignProgress structure.
type GetOutboundCampaignProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignProgressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignProgressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignProgressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignProgressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignProgressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignProgressRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignProgressRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignProgressUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignProgressTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignProgressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignProgressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignProgressGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignProgressOK creates a GetOutboundCampaignProgressOK with default headers values
func NewGetOutboundCampaignProgressOK() *GetOutboundCampaignProgressOK {
	return &GetOutboundCampaignProgressOK{}
}

/*GetOutboundCampaignProgressOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignProgressOK struct {
	Payload *models.CampaignProgress
}

func (o *GetOutboundCampaignProgressOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignProgressOK) GetPayload() *models.CampaignProgress {
	return o.Payload
}

func (o *GetOutboundCampaignProgressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignProgress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressBadRequest creates a GetOutboundCampaignProgressBadRequest with default headers values
func NewGetOutboundCampaignProgressBadRequest() *GetOutboundCampaignProgressBadRequest {
	return &GetOutboundCampaignProgressBadRequest{}
}

/*GetOutboundCampaignProgressBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignProgressBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignProgressBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressUnauthorized creates a GetOutboundCampaignProgressUnauthorized with default headers values
func NewGetOutboundCampaignProgressUnauthorized() *GetOutboundCampaignProgressUnauthorized {
	return &GetOutboundCampaignProgressUnauthorized{}
}

/*GetOutboundCampaignProgressUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignProgressUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignProgressUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressForbidden creates a GetOutboundCampaignProgressForbidden with default headers values
func NewGetOutboundCampaignProgressForbidden() *GetOutboundCampaignProgressForbidden {
	return &GetOutboundCampaignProgressForbidden{}
}

/*GetOutboundCampaignProgressForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignProgressForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignProgressForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressNotFound creates a GetOutboundCampaignProgressNotFound with default headers values
func NewGetOutboundCampaignProgressNotFound() *GetOutboundCampaignProgressNotFound {
	return &GetOutboundCampaignProgressNotFound{}
}

/*GetOutboundCampaignProgressNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignProgressNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignProgressNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressRequestTimeout creates a GetOutboundCampaignProgressRequestTimeout with default headers values
func NewGetOutboundCampaignProgressRequestTimeout() *GetOutboundCampaignProgressRequestTimeout {
	return &GetOutboundCampaignProgressRequestTimeout{}
}

/*GetOutboundCampaignProgressRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignProgressRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignProgressRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressRequestEntityTooLarge creates a GetOutboundCampaignProgressRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignProgressRequestEntityTooLarge() *GetOutboundCampaignProgressRequestEntityTooLarge {
	return &GetOutboundCampaignProgressRequestEntityTooLarge{}
}

/*GetOutboundCampaignProgressRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCampaignProgressRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignProgressRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressUnsupportedMediaType creates a GetOutboundCampaignProgressUnsupportedMediaType with default headers values
func NewGetOutboundCampaignProgressUnsupportedMediaType() *GetOutboundCampaignProgressUnsupportedMediaType {
	return &GetOutboundCampaignProgressUnsupportedMediaType{}
}

/*GetOutboundCampaignProgressUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignProgressUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignProgressUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressTooManyRequests creates a GetOutboundCampaignProgressTooManyRequests with default headers values
func NewGetOutboundCampaignProgressTooManyRequests() *GetOutboundCampaignProgressTooManyRequests {
	return &GetOutboundCampaignProgressTooManyRequests{}
}

/*GetOutboundCampaignProgressTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignProgressTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignProgressTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressInternalServerError creates a GetOutboundCampaignProgressInternalServerError with default headers values
func NewGetOutboundCampaignProgressInternalServerError() *GetOutboundCampaignProgressInternalServerError {
	return &GetOutboundCampaignProgressInternalServerError{}
}

/*GetOutboundCampaignProgressInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignProgressInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignProgressInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressServiceUnavailable creates a GetOutboundCampaignProgressServiceUnavailable with default headers values
func NewGetOutboundCampaignProgressServiceUnavailable() *GetOutboundCampaignProgressServiceUnavailable {
	return &GetOutboundCampaignProgressServiceUnavailable{}
}

/*GetOutboundCampaignProgressServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignProgressServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignProgressServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignProgressGatewayTimeout creates a GetOutboundCampaignProgressGatewayTimeout with default headers values
func NewGetOutboundCampaignProgressGatewayTimeout() *GetOutboundCampaignProgressGatewayTimeout {
	return &GetOutboundCampaignProgressGatewayTimeout{}
}

/*GetOutboundCampaignProgressGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignProgressGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignProgressGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/progress][%d] getOutboundCampaignProgressGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignProgressGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignProgressGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
