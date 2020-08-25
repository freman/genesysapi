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

// GetOutboundCampaignsReader is a Reader for the GetOutboundCampaigns structure.
type GetOutboundCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsOK creates a GetOutboundCampaignsOK with default headers values
func NewGetOutboundCampaignsOK() *GetOutboundCampaignsOK {
	return &GetOutboundCampaignsOK{}
}

/*GetOutboundCampaignsOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignsOK struct {
	Payload *models.CampaignEntityListing
}

func (o *GetOutboundCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsOK) GetPayload() *models.CampaignEntityListing {
	return o.Payload
}

func (o *GetOutboundCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsBadRequest creates a GetOutboundCampaignsBadRequest with default headers values
func NewGetOutboundCampaignsBadRequest() *GetOutboundCampaignsBadRequest {
	return &GetOutboundCampaignsBadRequest{}
}

/*GetOutboundCampaignsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsUnauthorized creates a GetOutboundCampaignsUnauthorized with default headers values
func NewGetOutboundCampaignsUnauthorized() *GetOutboundCampaignsUnauthorized {
	return &GetOutboundCampaignsUnauthorized{}
}

/*GetOutboundCampaignsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsForbidden creates a GetOutboundCampaignsForbidden with default headers values
func NewGetOutboundCampaignsForbidden() *GetOutboundCampaignsForbidden {
	return &GetOutboundCampaignsForbidden{}
}

/*GetOutboundCampaignsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsNotFound creates a GetOutboundCampaignsNotFound with default headers values
func NewGetOutboundCampaignsNotFound() *GetOutboundCampaignsNotFound {
	return &GetOutboundCampaignsNotFound{}
}

/*GetOutboundCampaignsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsRequestEntityTooLarge creates a GetOutboundCampaignsRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsRequestEntityTooLarge() *GetOutboundCampaignsRequestEntityTooLarge {
	return &GetOutboundCampaignsRequestEntityTooLarge{}
}

/*GetOutboundCampaignsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundCampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsUnsupportedMediaType creates a GetOutboundCampaignsUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsUnsupportedMediaType() *GetOutboundCampaignsUnsupportedMediaType {
	return &GetOutboundCampaignsUnsupportedMediaType{}
}

/*GetOutboundCampaignsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsTooManyRequests creates a GetOutboundCampaignsTooManyRequests with default headers values
func NewGetOutboundCampaignsTooManyRequests() *GetOutboundCampaignsTooManyRequests {
	return &GetOutboundCampaignsTooManyRequests{}
}

/*GetOutboundCampaignsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundCampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsInternalServerError creates a GetOutboundCampaignsInternalServerError with default headers values
func NewGetOutboundCampaignsInternalServerError() *GetOutboundCampaignsInternalServerError {
	return &GetOutboundCampaignsInternalServerError{}
}

/*GetOutboundCampaignsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsServiceUnavailable creates a GetOutboundCampaignsServiceUnavailable with default headers values
func NewGetOutboundCampaignsServiceUnavailable() *GetOutboundCampaignsServiceUnavailable {
	return &GetOutboundCampaignsServiceUnavailable{}
}

/*GetOutboundCampaignsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsGatewayTimeout creates a GetOutboundCampaignsGatewayTimeout with default headers values
func NewGetOutboundCampaignsGatewayTimeout() *GetOutboundCampaignsGatewayTimeout {
	return &GetOutboundCampaignsGatewayTimeout{}
}

/*GetOutboundCampaignsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns][%d] getOutboundCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}