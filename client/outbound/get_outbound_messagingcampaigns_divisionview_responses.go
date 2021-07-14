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

// GetOutboundMessagingcampaignsDivisionviewReader is a Reader for the GetOutboundMessagingcampaignsDivisionview structure.
type GetOutboundMessagingcampaignsDivisionviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundMessagingcampaignsDivisionviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundMessagingcampaignsDivisionviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundMessagingcampaignsDivisionviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundMessagingcampaignsDivisionviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundMessagingcampaignsDivisionviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundMessagingcampaignsDivisionviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundMessagingcampaignsDivisionviewRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundMessagingcampaignsDivisionviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundMessagingcampaignsDivisionviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundMessagingcampaignsDivisionviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundMessagingcampaignsDivisionviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewOK creates a GetOutboundMessagingcampaignsDivisionviewOK with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewOK() *GetOutboundMessagingcampaignsDivisionviewOK {
	return &GetOutboundMessagingcampaignsDivisionviewOK{}
}

/*GetOutboundMessagingcampaignsDivisionviewOK handles this case with default header values.

successful operation
*/
type GetOutboundMessagingcampaignsDivisionviewOK struct {
	Payload *models.MessagingCampaignDivisionView
}

func (o *GetOutboundMessagingcampaignsDivisionviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewOK) GetPayload() *models.MessagingCampaignDivisionView {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaignDivisionView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewBadRequest creates a GetOutboundMessagingcampaignsDivisionviewBadRequest with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewBadRequest() *GetOutboundMessagingcampaignsDivisionviewBadRequest {
	return &GetOutboundMessagingcampaignsDivisionviewBadRequest{}
}

/*GetOutboundMessagingcampaignsDivisionviewBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundMessagingcampaignsDivisionviewBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewUnauthorized creates a GetOutboundMessagingcampaignsDivisionviewUnauthorized with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewUnauthorized() *GetOutboundMessagingcampaignsDivisionviewUnauthorized {
	return &GetOutboundMessagingcampaignsDivisionviewUnauthorized{}
}

/*GetOutboundMessagingcampaignsDivisionviewUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundMessagingcampaignsDivisionviewUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewForbidden creates a GetOutboundMessagingcampaignsDivisionviewForbidden with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewForbidden() *GetOutboundMessagingcampaignsDivisionviewForbidden {
	return &GetOutboundMessagingcampaignsDivisionviewForbidden{}
}

/*GetOutboundMessagingcampaignsDivisionviewForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundMessagingcampaignsDivisionviewForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewNotFound creates a GetOutboundMessagingcampaignsDivisionviewNotFound with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewNotFound() *GetOutboundMessagingcampaignsDivisionviewNotFound {
	return &GetOutboundMessagingcampaignsDivisionviewNotFound{}
}

/*GetOutboundMessagingcampaignsDivisionviewNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundMessagingcampaignsDivisionviewNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewRequestTimeout creates a GetOutboundMessagingcampaignsDivisionviewRequestTimeout with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewRequestTimeout() *GetOutboundMessagingcampaignsDivisionviewRequestTimeout {
	return &GetOutboundMessagingcampaignsDivisionviewRequestTimeout{}
}

/*GetOutboundMessagingcampaignsDivisionviewRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundMessagingcampaignsDivisionviewRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge creates a GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge() *GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge {
	return &GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge{}
}

/*GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType creates a GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType() *GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType {
	return &GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType{}
}

/*GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewTooManyRequests creates a GetOutboundMessagingcampaignsDivisionviewTooManyRequests with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewTooManyRequests() *GetOutboundMessagingcampaignsDivisionviewTooManyRequests {
	return &GetOutboundMessagingcampaignsDivisionviewTooManyRequests{}
}

/*GetOutboundMessagingcampaignsDivisionviewTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundMessagingcampaignsDivisionviewTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewInternalServerError creates a GetOutboundMessagingcampaignsDivisionviewInternalServerError with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewInternalServerError() *GetOutboundMessagingcampaignsDivisionviewInternalServerError {
	return &GetOutboundMessagingcampaignsDivisionviewInternalServerError{}
}

/*GetOutboundMessagingcampaignsDivisionviewInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundMessagingcampaignsDivisionviewInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewServiceUnavailable creates a GetOutboundMessagingcampaignsDivisionviewServiceUnavailable with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewServiceUnavailable() *GetOutboundMessagingcampaignsDivisionviewServiceUnavailable {
	return &GetOutboundMessagingcampaignsDivisionviewServiceUnavailable{}
}

/*GetOutboundMessagingcampaignsDivisionviewServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundMessagingcampaignsDivisionviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewGatewayTimeout creates a GetOutboundMessagingcampaignsDivisionviewGatewayTimeout with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewGatewayTimeout() *GetOutboundMessagingcampaignsDivisionviewGatewayTimeout {
	return &GetOutboundMessagingcampaignsDivisionviewGatewayTimeout{}
}

/*GetOutboundMessagingcampaignsDivisionviewGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundMessagingcampaignsDivisionviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignsDivisionviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}][%d] getOutboundMessagingcampaignsDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
