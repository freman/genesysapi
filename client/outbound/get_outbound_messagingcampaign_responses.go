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

// GetOutboundMessagingcampaignReader is a Reader for the GetOutboundMessagingcampaign structure.
type GetOutboundMessagingcampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundMessagingcampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundMessagingcampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundMessagingcampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundMessagingcampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundMessagingcampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundMessagingcampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundMessagingcampaignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundMessagingcampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundMessagingcampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundMessagingcampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundMessagingcampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundMessagingcampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundMessagingcampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundMessagingcampaignOK creates a GetOutboundMessagingcampaignOK with default headers values
func NewGetOutboundMessagingcampaignOK() *GetOutboundMessagingcampaignOK {
	return &GetOutboundMessagingcampaignOK{}
}

/*GetOutboundMessagingcampaignOK handles this case with default header values.

successful operation
*/
type GetOutboundMessagingcampaignOK struct {
	Payload *models.MessagingCampaign
}

func (o *GetOutboundMessagingcampaignOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignOK  %+v", 200, o.Payload)
}

func (o *GetOutboundMessagingcampaignOK) GetPayload() *models.MessagingCampaign {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignBadRequest creates a GetOutboundMessagingcampaignBadRequest with default headers values
func NewGetOutboundMessagingcampaignBadRequest() *GetOutboundMessagingcampaignBadRequest {
	return &GetOutboundMessagingcampaignBadRequest{}
}

/*GetOutboundMessagingcampaignBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundMessagingcampaignBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundMessagingcampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignUnauthorized creates a GetOutboundMessagingcampaignUnauthorized with default headers values
func NewGetOutboundMessagingcampaignUnauthorized() *GetOutboundMessagingcampaignUnauthorized {
	return &GetOutboundMessagingcampaignUnauthorized{}
}

/*GetOutboundMessagingcampaignUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundMessagingcampaignUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundMessagingcampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignForbidden creates a GetOutboundMessagingcampaignForbidden with default headers values
func NewGetOutboundMessagingcampaignForbidden() *GetOutboundMessagingcampaignForbidden {
	return &GetOutboundMessagingcampaignForbidden{}
}

/*GetOutboundMessagingcampaignForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundMessagingcampaignForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundMessagingcampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignNotFound creates a GetOutboundMessagingcampaignNotFound with default headers values
func NewGetOutboundMessagingcampaignNotFound() *GetOutboundMessagingcampaignNotFound {
	return &GetOutboundMessagingcampaignNotFound{}
}

/*GetOutboundMessagingcampaignNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundMessagingcampaignNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundMessagingcampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignRequestTimeout creates a GetOutboundMessagingcampaignRequestTimeout with default headers values
func NewGetOutboundMessagingcampaignRequestTimeout() *GetOutboundMessagingcampaignRequestTimeout {
	return &GetOutboundMessagingcampaignRequestTimeout{}
}

/*GetOutboundMessagingcampaignRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundMessagingcampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundMessagingcampaignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignRequestEntityTooLarge creates a GetOutboundMessagingcampaignRequestEntityTooLarge with default headers values
func NewGetOutboundMessagingcampaignRequestEntityTooLarge() *GetOutboundMessagingcampaignRequestEntityTooLarge {
	return &GetOutboundMessagingcampaignRequestEntityTooLarge{}
}

/*GetOutboundMessagingcampaignRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundMessagingcampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundMessagingcampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignUnsupportedMediaType creates a GetOutboundMessagingcampaignUnsupportedMediaType with default headers values
func NewGetOutboundMessagingcampaignUnsupportedMediaType() *GetOutboundMessagingcampaignUnsupportedMediaType {
	return &GetOutboundMessagingcampaignUnsupportedMediaType{}
}

/*GetOutboundMessagingcampaignUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundMessagingcampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundMessagingcampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignTooManyRequests creates a GetOutboundMessagingcampaignTooManyRequests with default headers values
func NewGetOutboundMessagingcampaignTooManyRequests() *GetOutboundMessagingcampaignTooManyRequests {
	return &GetOutboundMessagingcampaignTooManyRequests{}
}

/*GetOutboundMessagingcampaignTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundMessagingcampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundMessagingcampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignInternalServerError creates a GetOutboundMessagingcampaignInternalServerError with default headers values
func NewGetOutboundMessagingcampaignInternalServerError() *GetOutboundMessagingcampaignInternalServerError {
	return &GetOutboundMessagingcampaignInternalServerError{}
}

/*GetOutboundMessagingcampaignInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundMessagingcampaignInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundMessagingcampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignServiceUnavailable creates a GetOutboundMessagingcampaignServiceUnavailable with default headers values
func NewGetOutboundMessagingcampaignServiceUnavailable() *GetOutboundMessagingcampaignServiceUnavailable {
	return &GetOutboundMessagingcampaignServiceUnavailable{}
}

/*GetOutboundMessagingcampaignServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundMessagingcampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundMessagingcampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignGatewayTimeout creates a GetOutboundMessagingcampaignGatewayTimeout with default headers values
func NewGetOutboundMessagingcampaignGatewayTimeout() *GetOutboundMessagingcampaignGatewayTimeout {
	return &GetOutboundMessagingcampaignGatewayTimeout{}
}

/*GetOutboundMessagingcampaignGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundMessagingcampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundMessagingcampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] getOutboundMessagingcampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundMessagingcampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
