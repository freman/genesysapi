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

/*
DeleteOutboundMessagingcampaignOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteOutboundMessagingcampaignOK struct {
	Payload *models.MessagingCampaign
}

// IsSuccess returns true when this delete outbound messagingcampaign o k response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete outbound messagingcampaign o k response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign o k response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound messagingcampaign o k response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign o k response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOutboundMessagingcampaignOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignOK  %+v", 200, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignOK) String() string {
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

/*
DeleteOutboundMessagingcampaignNoContent describes a response with status code 204, with default header values.

Messaging Campaign Deleted
*/
type DeleteOutboundMessagingcampaignNoContent struct {
}

// IsSuccess returns true when this delete outbound messagingcampaign no content response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete outbound messagingcampaign no content response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign no content response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound messagingcampaign no content response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign no content response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteOutboundMessagingcampaignNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignNoContent ", 204)
}

func (o *DeleteOutboundMessagingcampaignNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignNoContent ", 204)
}

func (o *DeleteOutboundMessagingcampaignNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundMessagingcampaignBadRequest creates a DeleteOutboundMessagingcampaignBadRequest with default headers values
func NewDeleteOutboundMessagingcampaignBadRequest() *DeleteOutboundMessagingcampaignBadRequest {
	return &DeleteOutboundMessagingcampaignBadRequest{}
}

/*
DeleteOutboundMessagingcampaignBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundMessagingcampaignBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign bad request response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign bad request response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign bad request response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign bad request response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign bad request response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteOutboundMessagingcampaignBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignBadRequest) String() string {
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

/*
DeleteOutboundMessagingcampaignUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundMessagingcampaignUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign unauthorized response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign unauthorized response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign unauthorized response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign unauthorized response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign unauthorized response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteOutboundMessagingcampaignUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignUnauthorized) String() string {
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

/*
DeleteOutboundMessagingcampaignForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundMessagingcampaignForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign forbidden response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign forbidden response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign forbidden response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign forbidden response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign forbidden response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteOutboundMessagingcampaignForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignForbidden) String() string {
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

/*
DeleteOutboundMessagingcampaignNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteOutboundMessagingcampaignNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign not found response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign not found response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign not found response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign not found response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign not found response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteOutboundMessagingcampaignNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignNotFound) String() string {
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

/*
DeleteOutboundMessagingcampaignRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundMessagingcampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign request timeout response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign request timeout response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign request timeout response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign request timeout response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign request timeout response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteOutboundMessagingcampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignRequestTimeout) String() string {
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

/*
DeleteOutboundMessagingcampaignRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteOutboundMessagingcampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign request entity too large response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign request entity too large response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign request entity too large response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign request entity too large response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign request entity too large response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignRequestEntityTooLarge) String() string {
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

/*
DeleteOutboundMessagingcampaignUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundMessagingcampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign unsupported media type response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign unsupported media type response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign unsupported media type response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign unsupported media type response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign unsupported media type response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignUnsupportedMediaType) String() string {
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

/*
DeleteOutboundMessagingcampaignTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundMessagingcampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign too many requests response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign too many requests response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign too many requests response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound messagingcampaign too many requests response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound messagingcampaign too many requests response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteOutboundMessagingcampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignTooManyRequests) String() string {
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

/*
DeleteOutboundMessagingcampaignInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundMessagingcampaignInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign internal server error response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign internal server error response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign internal server error response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound messagingcampaign internal server error response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound messagingcampaign internal server error response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteOutboundMessagingcampaignInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignInternalServerError) String() string {
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

/*
DeleteOutboundMessagingcampaignServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundMessagingcampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign service unavailable response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign service unavailable response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign service unavailable response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound messagingcampaign service unavailable response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound messagingcampaign service unavailable response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteOutboundMessagingcampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignServiceUnavailable) String() string {
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

/*
DeleteOutboundMessagingcampaignGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteOutboundMessagingcampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound messagingcampaign gateway timeout response has a 2xx status code
func (o *DeleteOutboundMessagingcampaignGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound messagingcampaign gateway timeout response has a 3xx status code
func (o *DeleteOutboundMessagingcampaignGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound messagingcampaign gateway timeout response has a 4xx status code
func (o *DeleteOutboundMessagingcampaignGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound messagingcampaign gateway timeout response has a 5xx status code
func (o *DeleteOutboundMessagingcampaignGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound messagingcampaign gateway timeout response a status code equal to that given
func (o *DeleteOutboundMessagingcampaignGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteOutboundMessagingcampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] deleteOutboundMessagingcampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundMessagingcampaignGatewayTimeout) String() string {
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
