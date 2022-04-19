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

// DeleteOutboundMessagingcampaignReader is a Reader for the DeleteOutboundMessagingcampaign structure.
type DeleteOutboundMessagingcampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundMessagingcampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundMessagingcampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteOutboundMessagingcampaignNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundMessagingcampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundMessagingcampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundMessagingcampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundMessagingcampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundMessagingcampaignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundMessagingcampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundMessagingcampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundMessagingcampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundMessagingcampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundMessagingcampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundMessagingcampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundMessagingcampaignOK creates a DeleteOutboundMessagingcampaignOK with default headers values
func NewDeleteOutboundMessagingcampaignOK() *DeleteOutboundMessagingcampaignOK {
	return &DeleteOutboundMessagingcampaignOK{}
}

/*DeleteOutboundMessagingcampaignOK handles this case with default header values.

successful operation
*/
type DeleteOutboundMessagingcampaignOK struct {
	Payload *models.MessagingCampaign
}

func (o *DeleteOutboundMessagingcampaignOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignOK  %+v", 200, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignOK) GetPayload() *models.MessagingCampaign {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignNoContent creates a DeleteOutboundMessagingcampaignNoContent with default headers values
func NewDeleteOutboundMessagingcampaignNoContent() *DeleteOutboundMessagingcampaignNoContent {
	return &DeleteOutboundMessagingcampaignNoContent{}
}

/*DeleteOutboundMessagingcampaignNoContent handles this case with default header values.

Messaging Campaign Deleted
*/
type DeleteOutboundMessagingcampaignNoContent struct {
}

func (o *DeleteOutboundMessagingcampaignNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignNoContent ", 204)
}

func (o *DeleteOutboundMessagingcampaignNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundMessagingcampaignBadRequest creates a DeleteOutboundMessagingcampaignBadRequest with default headers values
func NewDeleteOutboundMessagingcampaignBadRequest() *DeleteOutboundMessagingcampaignBadRequest {
	return &DeleteOutboundMessagingcampaignBadRequest{}
}

/*DeleteOutboundMessagingcampaignBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundMessagingcampaignBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignUnauthorized creates a DeleteOutboundMessagingcampaignUnauthorized with default headers values
func NewDeleteOutboundMessagingcampaignUnauthorized() *DeleteOutboundMessagingcampaignUnauthorized {
	return &DeleteOutboundMessagingcampaignUnauthorized{}
}

/*DeleteOutboundMessagingcampaignUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundMessagingcampaignUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignForbidden creates a DeleteOutboundMessagingcampaignForbidden with default headers values
func NewDeleteOutboundMessagingcampaignForbidden() *DeleteOutboundMessagingcampaignForbidden {
	return &DeleteOutboundMessagingcampaignForbidden{}
}

/*DeleteOutboundMessagingcampaignForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundMessagingcampaignForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignNotFound creates a DeleteOutboundMessagingcampaignNotFound with default headers values
func NewDeleteOutboundMessagingcampaignNotFound() *DeleteOutboundMessagingcampaignNotFound {
	return &DeleteOutboundMessagingcampaignNotFound{}
}

/*DeleteOutboundMessagingcampaignNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundMessagingcampaignNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignRequestTimeout creates a DeleteOutboundMessagingcampaignRequestTimeout with default headers values
func NewDeleteOutboundMessagingcampaignRequestTimeout() *DeleteOutboundMessagingcampaignRequestTimeout {
	return &DeleteOutboundMessagingcampaignRequestTimeout{}
}

/*DeleteOutboundMessagingcampaignRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundMessagingcampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignRequestEntityTooLarge creates a DeleteOutboundMessagingcampaignRequestEntityTooLarge with default headers values
func NewDeleteOutboundMessagingcampaignRequestEntityTooLarge() *DeleteOutboundMessagingcampaignRequestEntityTooLarge {
	return &DeleteOutboundMessagingcampaignRequestEntityTooLarge{}
}

/*DeleteOutboundMessagingcampaignRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteOutboundMessagingcampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignUnsupportedMediaType creates a DeleteOutboundMessagingcampaignUnsupportedMediaType with default headers values
func NewDeleteOutboundMessagingcampaignUnsupportedMediaType() *DeleteOutboundMessagingcampaignUnsupportedMediaType {
	return &DeleteOutboundMessagingcampaignUnsupportedMediaType{}
}

/*DeleteOutboundMessagingcampaignUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundMessagingcampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignTooManyRequests creates a DeleteOutboundMessagingcampaignTooManyRequests with default headers values
func NewDeleteOutboundMessagingcampaignTooManyRequests() *DeleteOutboundMessagingcampaignTooManyRequests {
	return &DeleteOutboundMessagingcampaignTooManyRequests{}
}

/*DeleteOutboundMessagingcampaignTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundMessagingcampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignInternalServerError creates a DeleteOutboundMessagingcampaignInternalServerError with default headers values
func NewDeleteOutboundMessagingcampaignInternalServerError() *DeleteOutboundMessagingcampaignInternalServerError {
	return &DeleteOutboundMessagingcampaignInternalServerError{}
}

/*DeleteOutboundMessagingcampaignInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundMessagingcampaignInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignServiceUnavailable creates a DeleteOutboundMessagingcampaignServiceUnavailable with default headers values
func NewDeleteOutboundMessagingcampaignServiceUnavailable() *DeleteOutboundMessagingcampaignServiceUnavailable {
	return &DeleteOutboundMessagingcampaignServiceUnavailable{}
}

/*DeleteOutboundMessagingcampaignServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundMessagingcampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundMessagingcampaignGatewayTimeout creates a DeleteOutboundMessagingcampaignGatewayTimeout with default headers values
func NewDeleteOutboundMessagingcampaignGatewayTimeout() *DeleteOutboundMessagingcampaignGatewayTimeout {
	return &DeleteOutboundMessagingcampaignGatewayTimeout{}
}

/*DeleteOutboundMessagingcampaignGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundMessagingcampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundMessagingcampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundMessagingcampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
