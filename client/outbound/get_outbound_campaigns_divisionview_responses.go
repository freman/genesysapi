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

// GetOutboundCampaignsDivisionviewReader is a Reader for the GetOutboundCampaignsDivisionview structure.
type GetOutboundCampaignsDivisionviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsDivisionviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsDivisionviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsDivisionviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsDivisionviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsDivisionviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsDivisionviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignsDivisionviewRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsDivisionviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsDivisionviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsDivisionviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsDivisionviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsDivisionviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsDivisionviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsDivisionviewOK creates a GetOutboundCampaignsDivisionviewOK with default headers values
func NewGetOutboundCampaignsDivisionviewOK() *GetOutboundCampaignsDivisionviewOK {
	return &GetOutboundCampaignsDivisionviewOK{}
}

/*GetOutboundCampaignsDivisionviewOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignsDivisionviewOK struct {
	Payload *models.CampaignDivisionView
}

func (o *GetOutboundCampaignsDivisionviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewOK) GetPayload() *models.CampaignDivisionView {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDivisionView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewBadRequest creates a GetOutboundCampaignsDivisionviewBadRequest with default headers values
func NewGetOutboundCampaignsDivisionviewBadRequest() *GetOutboundCampaignsDivisionviewBadRequest {
	return &GetOutboundCampaignsDivisionviewBadRequest{}
}

/*GetOutboundCampaignsDivisionviewBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsDivisionviewBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewUnauthorized creates a GetOutboundCampaignsDivisionviewUnauthorized with default headers values
func NewGetOutboundCampaignsDivisionviewUnauthorized() *GetOutboundCampaignsDivisionviewUnauthorized {
	return &GetOutboundCampaignsDivisionviewUnauthorized{}
}

/*GetOutboundCampaignsDivisionviewUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsDivisionviewUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewForbidden creates a GetOutboundCampaignsDivisionviewForbidden with default headers values
func NewGetOutboundCampaignsDivisionviewForbidden() *GetOutboundCampaignsDivisionviewForbidden {
	return &GetOutboundCampaignsDivisionviewForbidden{}
}

/*GetOutboundCampaignsDivisionviewForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsDivisionviewForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewNotFound creates a GetOutboundCampaignsDivisionviewNotFound with default headers values
func NewGetOutboundCampaignsDivisionviewNotFound() *GetOutboundCampaignsDivisionviewNotFound {
	return &GetOutboundCampaignsDivisionviewNotFound{}
}

/*GetOutboundCampaignsDivisionviewNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsDivisionviewNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewRequestTimeout creates a GetOutboundCampaignsDivisionviewRequestTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewRequestTimeout() *GetOutboundCampaignsDivisionviewRequestTimeout {
	return &GetOutboundCampaignsDivisionviewRequestTimeout{}
}

/*GetOutboundCampaignsDivisionviewRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignsDivisionviewRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewRequestEntityTooLarge creates a GetOutboundCampaignsDivisionviewRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsDivisionviewRequestEntityTooLarge() *GetOutboundCampaignsDivisionviewRequestEntityTooLarge {
	return &GetOutboundCampaignsDivisionviewRequestEntityTooLarge{}
}

/*GetOutboundCampaignsDivisionviewRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCampaignsDivisionviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewUnsupportedMediaType creates a GetOutboundCampaignsDivisionviewUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsDivisionviewUnsupportedMediaType() *GetOutboundCampaignsDivisionviewUnsupportedMediaType {
	return &GetOutboundCampaignsDivisionviewUnsupportedMediaType{}
}

/*GetOutboundCampaignsDivisionviewUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsDivisionviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewTooManyRequests creates a GetOutboundCampaignsDivisionviewTooManyRequests with default headers values
func NewGetOutboundCampaignsDivisionviewTooManyRequests() *GetOutboundCampaignsDivisionviewTooManyRequests {
	return &GetOutboundCampaignsDivisionviewTooManyRequests{}
}

/*GetOutboundCampaignsDivisionviewTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignsDivisionviewTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewInternalServerError creates a GetOutboundCampaignsDivisionviewInternalServerError with default headers values
func NewGetOutboundCampaignsDivisionviewInternalServerError() *GetOutboundCampaignsDivisionviewInternalServerError {
	return &GetOutboundCampaignsDivisionviewInternalServerError{}
}

/*GetOutboundCampaignsDivisionviewInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsDivisionviewInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewServiceUnavailable creates a GetOutboundCampaignsDivisionviewServiceUnavailable with default headers values
func NewGetOutboundCampaignsDivisionviewServiceUnavailable() *GetOutboundCampaignsDivisionviewServiceUnavailable {
	return &GetOutboundCampaignsDivisionviewServiceUnavailable{}
}

/*GetOutboundCampaignsDivisionviewServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsDivisionviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewGatewayTimeout creates a GetOutboundCampaignsDivisionviewGatewayTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewGatewayTimeout() *GetOutboundCampaignsDivisionviewGatewayTimeout {
	return &GetOutboundCampaignsDivisionviewGatewayTimeout{}
}

/*GetOutboundCampaignsDivisionviewGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignsDivisionviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
