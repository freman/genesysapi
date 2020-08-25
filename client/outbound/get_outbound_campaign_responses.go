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

// GetOutboundCampaignReader is a Reader for the GetOutboundCampaign structure.
type GetOutboundCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignOK creates a GetOutboundCampaignOK with default headers values
func NewGetOutboundCampaignOK() *GetOutboundCampaignOK {
	return &GetOutboundCampaignOK{}
}

/*GetOutboundCampaignOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignOK struct {
	Payload *models.Campaign
}

func (o *GetOutboundCampaignOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignOK) GetPayload() *models.Campaign {
	return o.Payload
}

func (o *GetOutboundCampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Campaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignBadRequest creates a GetOutboundCampaignBadRequest with default headers values
func NewGetOutboundCampaignBadRequest() *GetOutboundCampaignBadRequest {
	return &GetOutboundCampaignBadRequest{}
}

/*GetOutboundCampaignBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignUnauthorized creates a GetOutboundCampaignUnauthorized with default headers values
func NewGetOutboundCampaignUnauthorized() *GetOutboundCampaignUnauthorized {
	return &GetOutboundCampaignUnauthorized{}
}

/*GetOutboundCampaignUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignForbidden creates a GetOutboundCampaignForbidden with default headers values
func NewGetOutboundCampaignForbidden() *GetOutboundCampaignForbidden {
	return &GetOutboundCampaignForbidden{}
}

/*GetOutboundCampaignForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignNotFound creates a GetOutboundCampaignNotFound with default headers values
func NewGetOutboundCampaignNotFound() *GetOutboundCampaignNotFound {
	return &GetOutboundCampaignNotFound{}
}

/*GetOutboundCampaignNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignRequestEntityTooLarge creates a GetOutboundCampaignRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignRequestEntityTooLarge() *GetOutboundCampaignRequestEntityTooLarge {
	return &GetOutboundCampaignRequestEntityTooLarge{}
}

/*GetOutboundCampaignRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundCampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignUnsupportedMediaType creates a GetOutboundCampaignUnsupportedMediaType with default headers values
func NewGetOutboundCampaignUnsupportedMediaType() *GetOutboundCampaignUnsupportedMediaType {
	return &GetOutboundCampaignUnsupportedMediaType{}
}

/*GetOutboundCampaignUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignTooManyRequests creates a GetOutboundCampaignTooManyRequests with default headers values
func NewGetOutboundCampaignTooManyRequests() *GetOutboundCampaignTooManyRequests {
	return &GetOutboundCampaignTooManyRequests{}
}

/*GetOutboundCampaignTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundCampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInternalServerError creates a GetOutboundCampaignInternalServerError with default headers values
func NewGetOutboundCampaignInternalServerError() *GetOutboundCampaignInternalServerError {
	return &GetOutboundCampaignInternalServerError{}
}

/*GetOutboundCampaignInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignServiceUnavailable creates a GetOutboundCampaignServiceUnavailable with default headers values
func NewGetOutboundCampaignServiceUnavailable() *GetOutboundCampaignServiceUnavailable {
	return &GetOutboundCampaignServiceUnavailable{}
}

/*GetOutboundCampaignServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignGatewayTimeout creates a GetOutboundCampaignGatewayTimeout with default headers values
func NewGetOutboundCampaignGatewayTimeout() *GetOutboundCampaignGatewayTimeout {
	return &GetOutboundCampaignGatewayTimeout{}
}

/*GetOutboundCampaignGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}