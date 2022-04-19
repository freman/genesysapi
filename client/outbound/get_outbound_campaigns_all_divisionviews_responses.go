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

// GetOutboundCampaignsAllDivisionviewsReader is a Reader for the GetOutboundCampaignsAllDivisionviews structure.
type GetOutboundCampaignsAllDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsAllDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsAllDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsAllDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsAllDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsAllDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsAllDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignsAllDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsAllDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsAllDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsAllDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsAllDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsAllDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsAllDivisionviewsOK creates a GetOutboundCampaignsAllDivisionviewsOK with default headers values
func NewGetOutboundCampaignsAllDivisionviewsOK() *GetOutboundCampaignsAllDivisionviewsOK {
	return &GetOutboundCampaignsAllDivisionviewsOK{}
}

/*GetOutboundCampaignsAllDivisionviewsOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignsAllDivisionviewsOK struct {
	Payload *models.CommonCampaignDivisionViewEntityListing
}

func (o *GetOutboundCampaignsAllDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsOK) GetPayload() *models.CommonCampaignDivisionViewEntityListing {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonCampaignDivisionViewEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsBadRequest creates a GetOutboundCampaignsAllDivisionviewsBadRequest with default headers values
func NewGetOutboundCampaignsAllDivisionviewsBadRequest() *GetOutboundCampaignsAllDivisionviewsBadRequest {
	return &GetOutboundCampaignsAllDivisionviewsBadRequest{}
}

/*GetOutboundCampaignsAllDivisionviewsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsAllDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsUnauthorized creates a GetOutboundCampaignsAllDivisionviewsUnauthorized with default headers values
func NewGetOutboundCampaignsAllDivisionviewsUnauthorized() *GetOutboundCampaignsAllDivisionviewsUnauthorized {
	return &GetOutboundCampaignsAllDivisionviewsUnauthorized{}
}

/*GetOutboundCampaignsAllDivisionviewsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsAllDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsForbidden creates a GetOutboundCampaignsAllDivisionviewsForbidden with default headers values
func NewGetOutboundCampaignsAllDivisionviewsForbidden() *GetOutboundCampaignsAllDivisionviewsForbidden {
	return &GetOutboundCampaignsAllDivisionviewsForbidden{}
}

/*GetOutboundCampaignsAllDivisionviewsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsAllDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsNotFound creates a GetOutboundCampaignsAllDivisionviewsNotFound with default headers values
func NewGetOutboundCampaignsAllDivisionviewsNotFound() *GetOutboundCampaignsAllDivisionviewsNotFound {
	return &GetOutboundCampaignsAllDivisionviewsNotFound{}
}

/*GetOutboundCampaignsAllDivisionviewsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsAllDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsRequestTimeout creates a GetOutboundCampaignsAllDivisionviewsRequestTimeout with default headers values
func NewGetOutboundCampaignsAllDivisionviewsRequestTimeout() *GetOutboundCampaignsAllDivisionviewsRequestTimeout {
	return &GetOutboundCampaignsAllDivisionviewsRequestTimeout{}
}

/*GetOutboundCampaignsAllDivisionviewsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignsAllDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge creates a GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge() *GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge {
	return &GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge{}
}

/*GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsUnsupportedMediaType creates a GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsAllDivisionviewsUnsupportedMediaType() *GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType {
	return &GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType{}
}

/*GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsTooManyRequests creates a GetOutboundCampaignsAllDivisionviewsTooManyRequests with default headers values
func NewGetOutboundCampaignsAllDivisionviewsTooManyRequests() *GetOutboundCampaignsAllDivisionviewsTooManyRequests {
	return &GetOutboundCampaignsAllDivisionviewsTooManyRequests{}
}

/*GetOutboundCampaignsAllDivisionviewsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignsAllDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsInternalServerError creates a GetOutboundCampaignsAllDivisionviewsInternalServerError with default headers values
func NewGetOutboundCampaignsAllDivisionviewsInternalServerError() *GetOutboundCampaignsAllDivisionviewsInternalServerError {
	return &GetOutboundCampaignsAllDivisionviewsInternalServerError{}
}

/*GetOutboundCampaignsAllDivisionviewsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsAllDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsServiceUnavailable creates a GetOutboundCampaignsAllDivisionviewsServiceUnavailable with default headers values
func NewGetOutboundCampaignsAllDivisionviewsServiceUnavailable() *GetOutboundCampaignsAllDivisionviewsServiceUnavailable {
	return &GetOutboundCampaignsAllDivisionviewsServiceUnavailable{}
}

/*GetOutboundCampaignsAllDivisionviewsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsAllDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllDivisionviewsGatewayTimeout creates a GetOutboundCampaignsAllDivisionviewsGatewayTimeout with default headers values
func NewGetOutboundCampaignsAllDivisionviewsGatewayTimeout() *GetOutboundCampaignsAllDivisionviewsGatewayTimeout {
	return &GetOutboundCampaignsAllDivisionviewsGatewayTimeout{}
}

/*GetOutboundCampaignsAllDivisionviewsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignsAllDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all/divisionviews][%d] getOutboundCampaignsAllDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsAllDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
