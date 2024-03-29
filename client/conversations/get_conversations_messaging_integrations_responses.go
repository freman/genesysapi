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

// GetConversationsMessagingIntegrationsReader is a Reader for the GetConversationsMessagingIntegrations structure.
type GetConversationsMessagingIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingIntegrationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsOK creates a GetConversationsMessagingIntegrationsOK with default headers values
func NewGetConversationsMessagingIntegrationsOK() *GetConversationsMessagingIntegrationsOK {
	return &GetConversationsMessagingIntegrationsOK{}
}

/*
GetConversationsMessagingIntegrationsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsOK struct {
	Payload *models.MessagingIntegrationEntityListing
}

// IsSuccess returns true when this get conversations messaging integrations o k response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations messaging integrations o k response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations o k response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations o k response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations o k response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessagingIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOK) GetPayload() *models.MessagingIntegrationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingIntegrationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsBadRequest creates a GetConversationsMessagingIntegrationsBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsBadRequest() *GetConversationsMessagingIntegrationsBadRequest {
	return &GetConversationsMessagingIntegrationsBadRequest{}
}

/*
GetConversationsMessagingIntegrationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations bad request response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations bad request response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations bad request response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations bad request response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations bad request response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessagingIntegrationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsUnauthorized creates a GetConversationsMessagingIntegrationsUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsUnauthorized() *GetConversationsMessagingIntegrationsUnauthorized {
	return &GetConversationsMessagingIntegrationsUnauthorized{}
}

/*
GetConversationsMessagingIntegrationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations unauthorized response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations unauthorized response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations unauthorized response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations unauthorized response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations unauthorized response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessagingIntegrationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsForbidden creates a GetConversationsMessagingIntegrationsForbidden with default headers values
func NewGetConversationsMessagingIntegrationsForbidden() *GetConversationsMessagingIntegrationsForbidden {
	return &GetConversationsMessagingIntegrationsForbidden{}
}

/*
GetConversationsMessagingIntegrationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations forbidden response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations forbidden response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations forbidden response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations forbidden response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations forbidden response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessagingIntegrationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsNotFound creates a GetConversationsMessagingIntegrationsNotFound with default headers values
func NewGetConversationsMessagingIntegrationsNotFound() *GetConversationsMessagingIntegrationsNotFound {
	return &GetConversationsMessagingIntegrationsNotFound{}
}

/*
GetConversationsMessagingIntegrationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations not found response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations not found response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations not found response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations not found response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations not found response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessagingIntegrationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsRequestTimeout creates a GetConversationsMessagingIntegrationsRequestTimeout with default headers values
func NewGetConversationsMessagingIntegrationsRequestTimeout() *GetConversationsMessagingIntegrationsRequestTimeout {
	return &GetConversationsMessagingIntegrationsRequestTimeout{}
}

/*
GetConversationsMessagingIntegrationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingIntegrationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations request timeout response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations request timeout response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations request timeout response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations request timeout response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations request timeout response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessagingIntegrationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsRequestEntityTooLarge() *GetConversationsMessagingIntegrationsRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsRequestEntityTooLarge{}
}

/*
GetConversationsMessagingIntegrationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsMessagingIntegrationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations request entity too large response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations request entity too large response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations request entity too large response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations request entity too large response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations request entity too large response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsUnsupportedMediaType creates a GetConversationsMessagingIntegrationsUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsUnsupportedMediaType() *GetConversationsMessagingIntegrationsUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsUnsupportedMediaType{}
}

/*
GetConversationsMessagingIntegrationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations unsupported media type response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations unsupported media type response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations unsupported media type response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations unsupported media type response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations unsupported media type response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTooManyRequests creates a GetConversationsMessagingIntegrationsTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsTooManyRequests() *GetConversationsMessagingIntegrationsTooManyRequests {
	return &GetConversationsMessagingIntegrationsTooManyRequests{}
}

/*
GetConversationsMessagingIntegrationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingIntegrationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations too many requests response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations too many requests response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations too many requests response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations too many requests response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations too many requests response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessagingIntegrationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsInternalServerError creates a GetConversationsMessagingIntegrationsInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsInternalServerError() *GetConversationsMessagingIntegrationsInternalServerError {
	return &GetConversationsMessagingIntegrationsInternalServerError{}
}

/*
GetConversationsMessagingIntegrationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations internal server error response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations internal server error response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations internal server error response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations internal server error response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations internal server error response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessagingIntegrationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsServiceUnavailable creates a GetConversationsMessagingIntegrationsServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsServiceUnavailable() *GetConversationsMessagingIntegrationsServiceUnavailable {
	return &GetConversationsMessagingIntegrationsServiceUnavailable{}
}

/*
GetConversationsMessagingIntegrationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations service unavailable response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations service unavailable response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations service unavailable response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations service unavailable response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations service unavailable response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessagingIntegrationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsGatewayTimeout creates a GetConversationsMessagingIntegrationsGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsGatewayTimeout() *GetConversationsMessagingIntegrationsGatewayTimeout {
	return &GetConversationsMessagingIntegrationsGatewayTimeout{}
}

/*
GetConversationsMessagingIntegrationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations gateway timeout response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations gateway timeout response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations gateway timeout response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations gateway timeout response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations gateway timeout response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessagingIntegrationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations][%d] getConversationsMessagingIntegrationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
