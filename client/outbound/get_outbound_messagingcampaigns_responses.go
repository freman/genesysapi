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

// GetOutboundMessagingcampaignsReader is a Reader for the GetOutboundMessagingcampaigns structure.
type GetOutboundMessagingcampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundMessagingcampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundMessagingcampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundMessagingcampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundMessagingcampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundMessagingcampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundMessagingcampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundMessagingcampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundMessagingcampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundMessagingcampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundMessagingcampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundMessagingcampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundMessagingcampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundMessagingcampaignsOK creates a GetOutboundMessagingcampaignsOK with default headers values
func NewGetOutboundMessagingcampaignsOK() *GetOutboundMessagingcampaignsOK {
	return &GetOutboundMessagingcampaignsOK{}
}

/*GetOutboundMessagingcampaignsOK handles this case with default header values.

successful operation
*/
type GetOutboundMessagingcampaignsOK struct {
	Payload *models.MessagingCampaignEntityListing
}

func (o *GetOutboundMessagingcampaignsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundMessagingcampaignsOK) GetPayload() *models.MessagingCampaignEntityListing {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaignEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsBadRequest creates a GetOutboundMessagingcampaignsBadRequest with default headers values
func NewGetOutboundMessagingcampaignsBadRequest() *GetOutboundMessagingcampaignsBadRequest {
	return &GetOutboundMessagingcampaignsBadRequest{}
}

/*GetOutboundMessagingcampaignsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundMessagingcampaignsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundMessagingcampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsUnauthorized creates a GetOutboundMessagingcampaignsUnauthorized with default headers values
func NewGetOutboundMessagingcampaignsUnauthorized() *GetOutboundMessagingcampaignsUnauthorized {
	return &GetOutboundMessagingcampaignsUnauthorized{}
}

/*GetOutboundMessagingcampaignsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundMessagingcampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundMessagingcampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsForbidden creates a GetOutboundMessagingcampaignsForbidden with default headers values
func NewGetOutboundMessagingcampaignsForbidden() *GetOutboundMessagingcampaignsForbidden {
	return &GetOutboundMessagingcampaignsForbidden{}
}

/*GetOutboundMessagingcampaignsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundMessagingcampaignsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundMessagingcampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsNotFound creates a GetOutboundMessagingcampaignsNotFound with default headers values
func NewGetOutboundMessagingcampaignsNotFound() *GetOutboundMessagingcampaignsNotFound {
	return &GetOutboundMessagingcampaignsNotFound{}
}

/*GetOutboundMessagingcampaignsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundMessagingcampaignsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundMessagingcampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsRequestEntityTooLarge creates a GetOutboundMessagingcampaignsRequestEntityTooLarge with default headers values
func NewGetOutboundMessagingcampaignsRequestEntityTooLarge() *GetOutboundMessagingcampaignsRequestEntityTooLarge {
	return &GetOutboundMessagingcampaignsRequestEntityTooLarge{}
}

/*GetOutboundMessagingcampaignsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundMessagingcampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundMessagingcampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsUnsupportedMediaType creates a GetOutboundMessagingcampaignsUnsupportedMediaType with default headers values
func NewGetOutboundMessagingcampaignsUnsupportedMediaType() *GetOutboundMessagingcampaignsUnsupportedMediaType {
	return &GetOutboundMessagingcampaignsUnsupportedMediaType{}
}

/*GetOutboundMessagingcampaignsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundMessagingcampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundMessagingcampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsTooManyRequests creates a GetOutboundMessagingcampaignsTooManyRequests with default headers values
func NewGetOutboundMessagingcampaignsTooManyRequests() *GetOutboundMessagingcampaignsTooManyRequests {
	return &GetOutboundMessagingcampaignsTooManyRequests{}
}

/*GetOutboundMessagingcampaignsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundMessagingcampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundMessagingcampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsInternalServerError creates a GetOutboundMessagingcampaignsInternalServerError with default headers values
func NewGetOutboundMessagingcampaignsInternalServerError() *GetOutboundMessagingcampaignsInternalServerError {
	return &GetOutboundMessagingcampaignsInternalServerError{}
}

/*GetOutboundMessagingcampaignsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundMessagingcampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundMessagingcampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsServiceUnavailable creates a GetOutboundMessagingcampaignsServiceUnavailable with default headers values
func NewGetOutboundMessagingcampaignsServiceUnavailable() *GetOutboundMessagingcampaignsServiceUnavailable {
	return &GetOutboundMessagingcampaignsServiceUnavailable{}
}

/*GetOutboundMessagingcampaignsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundMessagingcampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundMessagingcampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsGatewayTimeout creates a GetOutboundMessagingcampaignsGatewayTimeout with default headers values
func NewGetOutboundMessagingcampaignsGatewayTimeout() *GetOutboundMessagingcampaignsGatewayTimeout {
	return &GetOutboundMessagingcampaignsGatewayTimeout{}
}

/*GetOutboundMessagingcampaignsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundMessagingcampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns][%d] getOutboundMessagingcampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundMessagingcampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
