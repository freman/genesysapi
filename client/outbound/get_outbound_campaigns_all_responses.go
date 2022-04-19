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

// GetOutboundCampaignsAllReader is a Reader for the GetOutboundCampaignsAll structure.
type GetOutboundCampaignsAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignsAllRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsAllRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsAllUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsAllTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsAllServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsAllGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsAllOK creates a GetOutboundCampaignsAllOK with default headers values
func NewGetOutboundCampaignsAllOK() *GetOutboundCampaignsAllOK {
	return &GetOutboundCampaignsAllOK{}
}

/*GetOutboundCampaignsAllOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignsAllOK struct {
	Payload *models.CommonCampaignEntityListing
}

func (o *GetOutboundCampaignsAllOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsAllOK) GetPayload() *models.CommonCampaignEntityListing {
	return o.Payload
}

func (o *GetOutboundCampaignsAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonCampaignEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllBadRequest creates a GetOutboundCampaignsAllBadRequest with default headers values
func NewGetOutboundCampaignsAllBadRequest() *GetOutboundCampaignsAllBadRequest {
	return &GetOutboundCampaignsAllBadRequest{}
}

/*GetOutboundCampaignsAllBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsAllBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsAllBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllUnauthorized creates a GetOutboundCampaignsAllUnauthorized with default headers values
func NewGetOutboundCampaignsAllUnauthorized() *GetOutboundCampaignsAllUnauthorized {
	return &GetOutboundCampaignsAllUnauthorized{}
}

/*GetOutboundCampaignsAllUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsAllUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsAllUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllForbidden creates a GetOutboundCampaignsAllForbidden with default headers values
func NewGetOutboundCampaignsAllForbidden() *GetOutboundCampaignsAllForbidden {
	return &GetOutboundCampaignsAllForbidden{}
}

/*GetOutboundCampaignsAllForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsAllForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsAllForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllNotFound creates a GetOutboundCampaignsAllNotFound with default headers values
func NewGetOutboundCampaignsAllNotFound() *GetOutboundCampaignsAllNotFound {
	return &GetOutboundCampaignsAllNotFound{}
}

/*GetOutboundCampaignsAllNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsAllNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsAllNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllRequestTimeout creates a GetOutboundCampaignsAllRequestTimeout with default headers values
func NewGetOutboundCampaignsAllRequestTimeout() *GetOutboundCampaignsAllRequestTimeout {
	return &GetOutboundCampaignsAllRequestTimeout{}
}

/*GetOutboundCampaignsAllRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignsAllRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsAllRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllRequestEntityTooLarge creates a GetOutboundCampaignsAllRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsAllRequestEntityTooLarge() *GetOutboundCampaignsAllRequestEntityTooLarge {
	return &GetOutboundCampaignsAllRequestEntityTooLarge{}
}

/*GetOutboundCampaignsAllRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCampaignsAllRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsAllRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllUnsupportedMediaType creates a GetOutboundCampaignsAllUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsAllUnsupportedMediaType() *GetOutboundCampaignsAllUnsupportedMediaType {
	return &GetOutboundCampaignsAllUnsupportedMediaType{}
}

/*GetOutboundCampaignsAllUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsAllUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsAllUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllTooManyRequests creates a GetOutboundCampaignsAllTooManyRequests with default headers values
func NewGetOutboundCampaignsAllTooManyRequests() *GetOutboundCampaignsAllTooManyRequests {
	return &GetOutboundCampaignsAllTooManyRequests{}
}

/*GetOutboundCampaignsAllTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignsAllTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsAllTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllInternalServerError creates a GetOutboundCampaignsAllInternalServerError with default headers values
func NewGetOutboundCampaignsAllInternalServerError() *GetOutboundCampaignsAllInternalServerError {
	return &GetOutboundCampaignsAllInternalServerError{}
}

/*GetOutboundCampaignsAllInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsAllInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsAllInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllServiceUnavailable creates a GetOutboundCampaignsAllServiceUnavailable with default headers values
func NewGetOutboundCampaignsAllServiceUnavailable() *GetOutboundCampaignsAllServiceUnavailable {
	return &GetOutboundCampaignsAllServiceUnavailable{}
}

/*GetOutboundCampaignsAllServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsAllServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsAllServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsAllGatewayTimeout creates a GetOutboundCampaignsAllGatewayTimeout with default headers values
func NewGetOutboundCampaignsAllGatewayTimeout() *GetOutboundCampaignsAllGatewayTimeout {
	return &GetOutboundCampaignsAllGatewayTimeout{}
}

/*GetOutboundCampaignsAllGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignsAllGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsAllGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/all][%d] getOutboundCampaignsAllGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsAllGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsAllGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
