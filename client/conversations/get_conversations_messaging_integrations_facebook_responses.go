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

// GetConversationsMessagingIntegrationsFacebookReader is a Reader for the GetConversationsMessagingIntegrationsFacebook structure.
type GetConversationsMessagingIntegrationsFacebookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsFacebookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsFacebookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsFacebookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsFacebookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsFacebookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsFacebookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingIntegrationsFacebookRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsFacebookUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsFacebookTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsFacebookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsFacebookServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsFacebookGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsFacebookOK creates a GetConversationsMessagingIntegrationsFacebookOK with default headers values
func NewGetConversationsMessagingIntegrationsFacebookOK() *GetConversationsMessagingIntegrationsFacebookOK {
	return &GetConversationsMessagingIntegrationsFacebookOK{}
}

/*
GetConversationsMessagingIntegrationsFacebookOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsFacebookOK struct {
	Payload *models.FacebookIntegrationEntityListing
}

// IsSuccess returns true when this get conversations messaging integrations facebook o k response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations messaging integrations facebook o k response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook o k response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations facebook o k response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook o k response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessagingIntegrationsFacebookOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookOK) GetPayload() *models.FacebookIntegrationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FacebookIntegrationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookBadRequest creates a GetConversationsMessagingIntegrationsFacebookBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsFacebookBadRequest() *GetConversationsMessagingIntegrationsFacebookBadRequest {
	return &GetConversationsMessagingIntegrationsFacebookBadRequest{}
}

/*
GetConversationsMessagingIntegrationsFacebookBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsFacebookBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook bad request response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook bad request response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook bad request response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook bad request response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook bad request response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookUnauthorized creates a GetConversationsMessagingIntegrationsFacebookUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsFacebookUnauthorized() *GetConversationsMessagingIntegrationsFacebookUnauthorized {
	return &GetConversationsMessagingIntegrationsFacebookUnauthorized{}
}

/*
GetConversationsMessagingIntegrationsFacebookUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsFacebookUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook unauthorized response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook unauthorized response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook unauthorized response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook unauthorized response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook unauthorized response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookForbidden creates a GetConversationsMessagingIntegrationsFacebookForbidden with default headers values
func NewGetConversationsMessagingIntegrationsFacebookForbidden() *GetConversationsMessagingIntegrationsFacebookForbidden {
	return &GetConversationsMessagingIntegrationsFacebookForbidden{}
}

/*
GetConversationsMessagingIntegrationsFacebookForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsFacebookForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook forbidden response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook forbidden response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook forbidden response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook forbidden response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook forbidden response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessagingIntegrationsFacebookForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookNotFound creates a GetConversationsMessagingIntegrationsFacebookNotFound with default headers values
func NewGetConversationsMessagingIntegrationsFacebookNotFound() *GetConversationsMessagingIntegrationsFacebookNotFound {
	return &GetConversationsMessagingIntegrationsFacebookNotFound{}
}

/*
GetConversationsMessagingIntegrationsFacebookNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsFacebookNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook not found response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook not found response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook not found response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook not found response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook not found response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessagingIntegrationsFacebookNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookRequestTimeout creates a GetConversationsMessagingIntegrationsFacebookRequestTimeout with default headers values
func NewGetConversationsMessagingIntegrationsFacebookRequestTimeout() *GetConversationsMessagingIntegrationsFacebookRequestTimeout {
	return &GetConversationsMessagingIntegrationsFacebookRequestTimeout{}
}

/*
GetConversationsMessagingIntegrationsFacebookRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingIntegrationsFacebookRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook request timeout response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook request timeout response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook request timeout response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook request timeout response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook request timeout response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge() *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge{}
}

/*
GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook request entity too large response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook request entity too large response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook request entity too large response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook request entity too large response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook request entity too large response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookUnsupportedMediaType creates a GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsFacebookUnsupportedMediaType() *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType{}
}

/*
GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook unsupported media type response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook unsupported media type response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook unsupported media type response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook unsupported media type response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook unsupported media type response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookTooManyRequests creates a GetConversationsMessagingIntegrationsFacebookTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsFacebookTooManyRequests() *GetConversationsMessagingIntegrationsFacebookTooManyRequests {
	return &GetConversationsMessagingIntegrationsFacebookTooManyRequests{}
}

/*
GetConversationsMessagingIntegrationsFacebookTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingIntegrationsFacebookTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook too many requests response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook too many requests response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook too many requests response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations facebook too many requests response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations facebook too many requests response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookInternalServerError creates a GetConversationsMessagingIntegrationsFacebookInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsFacebookInternalServerError() *GetConversationsMessagingIntegrationsFacebookInternalServerError {
	return &GetConversationsMessagingIntegrationsFacebookInternalServerError{}
}

/*
GetConversationsMessagingIntegrationsFacebookInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsFacebookInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook internal server error response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook internal server error response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook internal server error response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations facebook internal server error response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations facebook internal server error response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookServiceUnavailable creates a GetConversationsMessagingIntegrationsFacebookServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsFacebookServiceUnavailable() *GetConversationsMessagingIntegrationsFacebookServiceUnavailable {
	return &GetConversationsMessagingIntegrationsFacebookServiceUnavailable{}
}

/*
GetConversationsMessagingIntegrationsFacebookServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsFacebookServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook service unavailable response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook service unavailable response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook service unavailable response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations facebook service unavailable response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations facebook service unavailable response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsFacebookGatewayTimeout creates a GetConversationsMessagingIntegrationsFacebookGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsFacebookGatewayTimeout() *GetConversationsMessagingIntegrationsFacebookGatewayTimeout {
	return &GetConversationsMessagingIntegrationsFacebookGatewayTimeout{}
}

/*
GetConversationsMessagingIntegrationsFacebookGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsFacebookGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations facebook gateway timeout response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations facebook gateway timeout response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations facebook gateway timeout response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations facebook gateway timeout response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations facebook gateway timeout response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/facebook][%d] getConversationsMessagingIntegrationsFacebookGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsFacebookGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
