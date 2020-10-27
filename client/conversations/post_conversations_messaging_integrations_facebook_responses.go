// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostConversationsMessagingIntegrationsFacebookReader is a Reader for the PostConversationsMessagingIntegrationsFacebook structure.
type PostConversationsMessagingIntegrationsFacebookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessagingIntegrationsFacebookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsMessagingIntegrationsFacebookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostConversationsMessagingIntegrationsFacebookAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessagingIntegrationsFacebookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessagingIntegrationsFacebookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessagingIntegrationsFacebookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessagingIntegrationsFacebookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessagingIntegrationsFacebookUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessagingIntegrationsFacebookTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessagingIntegrationsFacebookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessagingIntegrationsFacebookServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessagingIntegrationsFacebookGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessagingIntegrationsFacebookOK creates a PostConversationsMessagingIntegrationsFacebookOK with default headers values
func NewPostConversationsMessagingIntegrationsFacebookOK() *PostConversationsMessagingIntegrationsFacebookOK {
	return &PostConversationsMessagingIntegrationsFacebookOK{}
}

/*PostConversationsMessagingIntegrationsFacebookOK handles this case with default header values.

successful operation
*/
type PostConversationsMessagingIntegrationsFacebookOK struct {
	Payload *models.FacebookIntegration
}

func (o *PostConversationsMessagingIntegrationsFacebookOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookOK  %+v", 200, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookOK) GetPayload() *models.FacebookIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FacebookIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookAccepted creates a PostConversationsMessagingIntegrationsFacebookAccepted with default headers values
func NewPostConversationsMessagingIntegrationsFacebookAccepted() *PostConversationsMessagingIntegrationsFacebookAccepted {
	return &PostConversationsMessagingIntegrationsFacebookAccepted{}
}

/*PostConversationsMessagingIntegrationsFacebookAccepted handles this case with default header values.

Accepted - If async is true, the integration creation in progress.
*/
type PostConversationsMessagingIntegrationsFacebookAccepted struct {
	Payload *models.FacebookIntegration
}

func (o *PostConversationsMessagingIntegrationsFacebookAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookAccepted  %+v", 202, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookAccepted) GetPayload() *models.FacebookIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FacebookIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookBadRequest creates a PostConversationsMessagingIntegrationsFacebookBadRequest with default headers values
func NewPostConversationsMessagingIntegrationsFacebookBadRequest() *PostConversationsMessagingIntegrationsFacebookBadRequest {
	return &PostConversationsMessagingIntegrationsFacebookBadRequest{}
}

/*PostConversationsMessagingIntegrationsFacebookBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessagingIntegrationsFacebookBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookUnauthorized creates a PostConversationsMessagingIntegrationsFacebookUnauthorized with default headers values
func NewPostConversationsMessagingIntegrationsFacebookUnauthorized() *PostConversationsMessagingIntegrationsFacebookUnauthorized {
	return &PostConversationsMessagingIntegrationsFacebookUnauthorized{}
}

/*PostConversationsMessagingIntegrationsFacebookUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessagingIntegrationsFacebookUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookForbidden creates a PostConversationsMessagingIntegrationsFacebookForbidden with default headers values
func NewPostConversationsMessagingIntegrationsFacebookForbidden() *PostConversationsMessagingIntegrationsFacebookForbidden {
	return &PostConversationsMessagingIntegrationsFacebookForbidden{}
}

/*PostConversationsMessagingIntegrationsFacebookForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessagingIntegrationsFacebookForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookNotFound creates a PostConversationsMessagingIntegrationsFacebookNotFound with default headers values
func NewPostConversationsMessagingIntegrationsFacebookNotFound() *PostConversationsMessagingIntegrationsFacebookNotFound {
	return &PostConversationsMessagingIntegrationsFacebookNotFound{}
}

/*PostConversationsMessagingIntegrationsFacebookNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessagingIntegrationsFacebookNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge creates a PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge with default headers values
func NewPostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge() *PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge {
	return &PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge{}
}

/*PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookUnsupportedMediaType creates a PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType with default headers values
func NewPostConversationsMessagingIntegrationsFacebookUnsupportedMediaType() *PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType {
	return &PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType{}
}

/*PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookTooManyRequests creates a PostConversationsMessagingIntegrationsFacebookTooManyRequests with default headers values
func NewPostConversationsMessagingIntegrationsFacebookTooManyRequests() *PostConversationsMessagingIntegrationsFacebookTooManyRequests {
	return &PostConversationsMessagingIntegrationsFacebookTooManyRequests{}
}

/*PostConversationsMessagingIntegrationsFacebookTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsMessagingIntegrationsFacebookTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookInternalServerError creates a PostConversationsMessagingIntegrationsFacebookInternalServerError with default headers values
func NewPostConversationsMessagingIntegrationsFacebookInternalServerError() *PostConversationsMessagingIntegrationsFacebookInternalServerError {
	return &PostConversationsMessagingIntegrationsFacebookInternalServerError{}
}

/*PostConversationsMessagingIntegrationsFacebookInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessagingIntegrationsFacebookInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookServiceUnavailable creates a PostConversationsMessagingIntegrationsFacebookServiceUnavailable with default headers values
func NewPostConversationsMessagingIntegrationsFacebookServiceUnavailable() *PostConversationsMessagingIntegrationsFacebookServiceUnavailable {
	return &PostConversationsMessagingIntegrationsFacebookServiceUnavailable{}
}

/*PostConversationsMessagingIntegrationsFacebookServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessagingIntegrationsFacebookServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsFacebookGatewayTimeout creates a PostConversationsMessagingIntegrationsFacebookGatewayTimeout with default headers values
func NewPostConversationsMessagingIntegrationsFacebookGatewayTimeout() *PostConversationsMessagingIntegrationsFacebookGatewayTimeout {
	return &PostConversationsMessagingIntegrationsFacebookGatewayTimeout{}
}

/*PostConversationsMessagingIntegrationsFacebookGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessagingIntegrationsFacebookGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsFacebookGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/facebook][%d] postConversationsMessagingIntegrationsFacebookGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsFacebookGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsFacebookGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
