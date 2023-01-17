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

// GetConversationsMessagingIntegrationsLineReader is a Reader for the GetConversationsMessagingIntegrationsLine structure.
type GetConversationsMessagingIntegrationsLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsLineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsLineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsLineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsLineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingIntegrationsLineRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsLineRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsLineUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsLineTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsLineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsLineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsLineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsLineOK creates a GetConversationsMessagingIntegrationsLineOK with default headers values
func NewGetConversationsMessagingIntegrationsLineOK() *GetConversationsMessagingIntegrationsLineOK {
	return &GetConversationsMessagingIntegrationsLineOK{}
}

/*
GetConversationsMessagingIntegrationsLineOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsLineOK struct {
	Payload *models.LineIntegrationEntityListing
}

// IsSuccess returns true when this get conversations messaging integrations line o k response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations messaging integrations line o k response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line o k response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations line o k response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line o k response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessagingIntegrationsLineOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineOK) GetPayload() *models.LineIntegrationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineIntegrationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineBadRequest creates a GetConversationsMessagingIntegrationsLineBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsLineBadRequest() *GetConversationsMessagingIntegrationsLineBadRequest {
	return &GetConversationsMessagingIntegrationsLineBadRequest{}
}

/*
GetConversationsMessagingIntegrationsLineBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsLineBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line bad request response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line bad request response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line bad request response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line bad request response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line bad request response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineUnauthorized creates a GetConversationsMessagingIntegrationsLineUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsLineUnauthorized() *GetConversationsMessagingIntegrationsLineUnauthorized {
	return &GetConversationsMessagingIntegrationsLineUnauthorized{}
}

/*
GetConversationsMessagingIntegrationsLineUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsLineUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line unauthorized response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line unauthorized response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line unauthorized response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line unauthorized response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line unauthorized response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineForbidden creates a GetConversationsMessagingIntegrationsLineForbidden with default headers values
func NewGetConversationsMessagingIntegrationsLineForbidden() *GetConversationsMessagingIntegrationsLineForbidden {
	return &GetConversationsMessagingIntegrationsLineForbidden{}
}

/*
GetConversationsMessagingIntegrationsLineForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsLineForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line forbidden response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line forbidden response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line forbidden response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line forbidden response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line forbidden response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineNotFound creates a GetConversationsMessagingIntegrationsLineNotFound with default headers values
func NewGetConversationsMessagingIntegrationsLineNotFound() *GetConversationsMessagingIntegrationsLineNotFound {
	return &GetConversationsMessagingIntegrationsLineNotFound{}
}

/*
GetConversationsMessagingIntegrationsLineNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsLineNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line not found response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line not found response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line not found response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line not found response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line not found response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineRequestTimeout creates a GetConversationsMessagingIntegrationsLineRequestTimeout with default headers values
func NewGetConversationsMessagingIntegrationsLineRequestTimeout() *GetConversationsMessagingIntegrationsLineRequestTimeout {
	return &GetConversationsMessagingIntegrationsLineRequestTimeout{}
}

/*
GetConversationsMessagingIntegrationsLineRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingIntegrationsLineRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line request timeout response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line request timeout response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line request timeout response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line request timeout response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line request timeout response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsLineRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsLineRequestEntityTooLarge() *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsLineRequestEntityTooLarge{}
}

/*
GetConversationsMessagingIntegrationsLineRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsMessagingIntegrationsLineRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line request entity too large response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line request entity too large response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line request entity too large response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line request entity too large response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line request entity too large response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineUnsupportedMediaType creates a GetConversationsMessagingIntegrationsLineUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsLineUnsupportedMediaType() *GetConversationsMessagingIntegrationsLineUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsLineUnsupportedMediaType{}
}

/*
GetConversationsMessagingIntegrationsLineUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsLineUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line unsupported media type response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line unsupported media type response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line unsupported media type response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line unsupported media type response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line unsupported media type response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineTooManyRequests creates a GetConversationsMessagingIntegrationsLineTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsLineTooManyRequests() *GetConversationsMessagingIntegrationsLineTooManyRequests {
	return &GetConversationsMessagingIntegrationsLineTooManyRequests{}
}

/*
GetConversationsMessagingIntegrationsLineTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingIntegrationsLineTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line too many requests response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line too many requests response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line too many requests response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging integrations line too many requests response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging integrations line too many requests response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineInternalServerError creates a GetConversationsMessagingIntegrationsLineInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsLineInternalServerError() *GetConversationsMessagingIntegrationsLineInternalServerError {
	return &GetConversationsMessagingIntegrationsLineInternalServerError{}
}

/*
GetConversationsMessagingIntegrationsLineInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsLineInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line internal server error response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line internal server error response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line internal server error response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations line internal server error response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations line internal server error response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineServiceUnavailable creates a GetConversationsMessagingIntegrationsLineServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsLineServiceUnavailable() *GetConversationsMessagingIntegrationsLineServiceUnavailable {
	return &GetConversationsMessagingIntegrationsLineServiceUnavailable{}
}

/*
GetConversationsMessagingIntegrationsLineServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsLineServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line service unavailable response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line service unavailable response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line service unavailable response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations line service unavailable response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations line service unavailable response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineGatewayTimeout creates a GetConversationsMessagingIntegrationsLineGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsLineGatewayTimeout() *GetConversationsMessagingIntegrationsLineGatewayTimeout {
	return &GetConversationsMessagingIntegrationsLineGatewayTimeout{}
}

/*
GetConversationsMessagingIntegrationsLineGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsLineGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging integrations line gateway timeout response has a 2xx status code
func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging integrations line gateway timeout response has a 3xx status code
func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging integrations line gateway timeout response has a 4xx status code
func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging integrations line gateway timeout response has a 5xx status code
func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging integrations line gateway timeout response a status code equal to that given
func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
