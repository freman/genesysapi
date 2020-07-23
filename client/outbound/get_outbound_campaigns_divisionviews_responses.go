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

// GetOutboundCampaignsDivisionviewsReader is a Reader for the GetOutboundCampaignsDivisionviews structure.
type GetOutboundCampaignsDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsDivisionviewsOK creates a GetOutboundCampaignsDivisionviewsOK with default headers values
func NewGetOutboundCampaignsDivisionviewsOK() *GetOutboundCampaignsDivisionviewsOK {
	return &GetOutboundCampaignsDivisionviewsOK{}
}

/*GetOutboundCampaignsDivisionviewsOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignsDivisionviewsOK struct {
	Payload *models.CampaignDivisionViewListing
}

func (o *GetOutboundCampaignsDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsOK) GetPayload() *models.CampaignDivisionViewListing {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDivisionViewListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsBadRequest creates a GetOutboundCampaignsDivisionviewsBadRequest with default headers values
func NewGetOutboundCampaignsDivisionviewsBadRequest() *GetOutboundCampaignsDivisionviewsBadRequest {
	return &GetOutboundCampaignsDivisionviewsBadRequest{}
}

/*GetOutboundCampaignsDivisionviewsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsUnauthorized creates a GetOutboundCampaignsDivisionviewsUnauthorized with default headers values
func NewGetOutboundCampaignsDivisionviewsUnauthorized() *GetOutboundCampaignsDivisionviewsUnauthorized {
	return &GetOutboundCampaignsDivisionviewsUnauthorized{}
}

/*GetOutboundCampaignsDivisionviewsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsForbidden creates a GetOutboundCampaignsDivisionviewsForbidden with default headers values
func NewGetOutboundCampaignsDivisionviewsForbidden() *GetOutboundCampaignsDivisionviewsForbidden {
	return &GetOutboundCampaignsDivisionviewsForbidden{}
}

/*GetOutboundCampaignsDivisionviewsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsNotFound creates a GetOutboundCampaignsDivisionviewsNotFound with default headers values
func NewGetOutboundCampaignsDivisionviewsNotFound() *GetOutboundCampaignsDivisionviewsNotFound {
	return &GetOutboundCampaignsDivisionviewsNotFound{}
}

/*GetOutboundCampaignsDivisionviewsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsRequestEntityTooLarge creates a GetOutboundCampaignsDivisionviewsRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsDivisionviewsRequestEntityTooLarge() *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge {
	return &GetOutboundCampaignsDivisionviewsRequestEntityTooLarge{}
}

/*GetOutboundCampaignsDivisionviewsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundCampaignsDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsUnsupportedMediaType creates a GetOutboundCampaignsDivisionviewsUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsDivisionviewsUnsupportedMediaType() *GetOutboundCampaignsDivisionviewsUnsupportedMediaType {
	return &GetOutboundCampaignsDivisionviewsUnsupportedMediaType{}
}

/*GetOutboundCampaignsDivisionviewsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsTooManyRequests creates a GetOutboundCampaignsDivisionviewsTooManyRequests with default headers values
func NewGetOutboundCampaignsDivisionviewsTooManyRequests() *GetOutboundCampaignsDivisionviewsTooManyRequests {
	return &GetOutboundCampaignsDivisionviewsTooManyRequests{}
}

/*GetOutboundCampaignsDivisionviewsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundCampaignsDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsInternalServerError creates a GetOutboundCampaignsDivisionviewsInternalServerError with default headers values
func NewGetOutboundCampaignsDivisionviewsInternalServerError() *GetOutboundCampaignsDivisionviewsInternalServerError {
	return &GetOutboundCampaignsDivisionviewsInternalServerError{}
}

/*GetOutboundCampaignsDivisionviewsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsServiceUnavailable creates a GetOutboundCampaignsDivisionviewsServiceUnavailable with default headers values
func NewGetOutboundCampaignsDivisionviewsServiceUnavailable() *GetOutboundCampaignsDivisionviewsServiceUnavailable {
	return &GetOutboundCampaignsDivisionviewsServiceUnavailable{}
}

/*GetOutboundCampaignsDivisionviewsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsGatewayTimeout creates a GetOutboundCampaignsDivisionviewsGatewayTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewsGatewayTimeout() *GetOutboundCampaignsDivisionviewsGatewayTimeout {
	return &GetOutboundCampaignsDivisionviewsGatewayTimeout{}
}

/*GetOutboundCampaignsDivisionviewsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignsDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
