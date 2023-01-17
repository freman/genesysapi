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

// GetConversationsMessagingStickerReader is a Reader for the GetConversationsMessagingSticker structure.
type GetConversationsMessagingStickerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingStickerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingStickerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingStickerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingStickerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingStickerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingStickerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingStickerRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingStickerRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingStickerUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingStickerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingStickerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingStickerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingStickerGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingStickerOK creates a GetConversationsMessagingStickerOK with default headers values
func NewGetConversationsMessagingStickerOK() *GetConversationsMessagingStickerOK {
	return &GetConversationsMessagingStickerOK{}
}

/*
GetConversationsMessagingStickerOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessagingStickerOK struct {
	Payload *models.MessagingStickerEntityListing
}

// IsSuccess returns true when this get conversations messaging sticker o k response has a 2xx status code
func (o *GetConversationsMessagingStickerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations messaging sticker o k response has a 3xx status code
func (o *GetConversationsMessagingStickerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker o k response has a 4xx status code
func (o *GetConversationsMessagingStickerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging sticker o k response has a 5xx status code
func (o *GetConversationsMessagingStickerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker o k response a status code equal to that given
func (o *GetConversationsMessagingStickerOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessagingStickerOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingStickerOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingStickerOK) GetPayload() *models.MessagingStickerEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingStickerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingStickerEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerBadRequest creates a GetConversationsMessagingStickerBadRequest with default headers values
func NewGetConversationsMessagingStickerBadRequest() *GetConversationsMessagingStickerBadRequest {
	return &GetConversationsMessagingStickerBadRequest{}
}

/*
GetConversationsMessagingStickerBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingStickerBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker bad request response has a 2xx status code
func (o *GetConversationsMessagingStickerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker bad request response has a 3xx status code
func (o *GetConversationsMessagingStickerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker bad request response has a 4xx status code
func (o *GetConversationsMessagingStickerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker bad request response has a 5xx status code
func (o *GetConversationsMessagingStickerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker bad request response a status code equal to that given
func (o *GetConversationsMessagingStickerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessagingStickerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingStickerBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingStickerBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerUnauthorized creates a GetConversationsMessagingStickerUnauthorized with default headers values
func NewGetConversationsMessagingStickerUnauthorized() *GetConversationsMessagingStickerUnauthorized {
	return &GetConversationsMessagingStickerUnauthorized{}
}

/*
GetConversationsMessagingStickerUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingStickerUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker unauthorized response has a 2xx status code
func (o *GetConversationsMessagingStickerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker unauthorized response has a 3xx status code
func (o *GetConversationsMessagingStickerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker unauthorized response has a 4xx status code
func (o *GetConversationsMessagingStickerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker unauthorized response has a 5xx status code
func (o *GetConversationsMessagingStickerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker unauthorized response a status code equal to that given
func (o *GetConversationsMessagingStickerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessagingStickerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingStickerUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingStickerUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerForbidden creates a GetConversationsMessagingStickerForbidden with default headers values
func NewGetConversationsMessagingStickerForbidden() *GetConversationsMessagingStickerForbidden {
	return &GetConversationsMessagingStickerForbidden{}
}

/*
GetConversationsMessagingStickerForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingStickerForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker forbidden response has a 2xx status code
func (o *GetConversationsMessagingStickerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker forbidden response has a 3xx status code
func (o *GetConversationsMessagingStickerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker forbidden response has a 4xx status code
func (o *GetConversationsMessagingStickerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker forbidden response has a 5xx status code
func (o *GetConversationsMessagingStickerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker forbidden response a status code equal to that given
func (o *GetConversationsMessagingStickerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessagingStickerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingStickerForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingStickerForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerNotFound creates a GetConversationsMessagingStickerNotFound with default headers values
func NewGetConversationsMessagingStickerNotFound() *GetConversationsMessagingStickerNotFound {
	return &GetConversationsMessagingStickerNotFound{}
}

/*
GetConversationsMessagingStickerNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingStickerNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker not found response has a 2xx status code
func (o *GetConversationsMessagingStickerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker not found response has a 3xx status code
func (o *GetConversationsMessagingStickerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker not found response has a 4xx status code
func (o *GetConversationsMessagingStickerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker not found response has a 5xx status code
func (o *GetConversationsMessagingStickerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker not found response a status code equal to that given
func (o *GetConversationsMessagingStickerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessagingStickerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingStickerNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingStickerNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerRequestTimeout creates a GetConversationsMessagingStickerRequestTimeout with default headers values
func NewGetConversationsMessagingStickerRequestTimeout() *GetConversationsMessagingStickerRequestTimeout {
	return &GetConversationsMessagingStickerRequestTimeout{}
}

/*
GetConversationsMessagingStickerRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingStickerRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker request timeout response has a 2xx status code
func (o *GetConversationsMessagingStickerRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker request timeout response has a 3xx status code
func (o *GetConversationsMessagingStickerRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker request timeout response has a 4xx status code
func (o *GetConversationsMessagingStickerRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker request timeout response has a 5xx status code
func (o *GetConversationsMessagingStickerRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker request timeout response a status code equal to that given
func (o *GetConversationsMessagingStickerRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessagingStickerRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingStickerRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingStickerRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerRequestEntityTooLarge creates a GetConversationsMessagingStickerRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingStickerRequestEntityTooLarge() *GetConversationsMessagingStickerRequestEntityTooLarge {
	return &GetConversationsMessagingStickerRequestEntityTooLarge{}
}

/*
GetConversationsMessagingStickerRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsMessagingStickerRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker request entity too large response has a 2xx status code
func (o *GetConversationsMessagingStickerRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker request entity too large response has a 3xx status code
func (o *GetConversationsMessagingStickerRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker request entity too large response has a 4xx status code
func (o *GetConversationsMessagingStickerRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker request entity too large response has a 5xx status code
func (o *GetConversationsMessagingStickerRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker request entity too large response a status code equal to that given
func (o *GetConversationsMessagingStickerRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerUnsupportedMediaType creates a GetConversationsMessagingStickerUnsupportedMediaType with default headers values
func NewGetConversationsMessagingStickerUnsupportedMediaType() *GetConversationsMessagingStickerUnsupportedMediaType {
	return &GetConversationsMessagingStickerUnsupportedMediaType{}
}

/*
GetConversationsMessagingStickerUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingStickerUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker unsupported media type response has a 2xx status code
func (o *GetConversationsMessagingStickerUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker unsupported media type response has a 3xx status code
func (o *GetConversationsMessagingStickerUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker unsupported media type response has a 4xx status code
func (o *GetConversationsMessagingStickerUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker unsupported media type response has a 5xx status code
func (o *GetConversationsMessagingStickerUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker unsupported media type response a status code equal to that given
func (o *GetConversationsMessagingStickerUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerTooManyRequests creates a GetConversationsMessagingStickerTooManyRequests with default headers values
func NewGetConversationsMessagingStickerTooManyRequests() *GetConversationsMessagingStickerTooManyRequests {
	return &GetConversationsMessagingStickerTooManyRequests{}
}

/*
GetConversationsMessagingStickerTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingStickerTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker too many requests response has a 2xx status code
func (o *GetConversationsMessagingStickerTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker too many requests response has a 3xx status code
func (o *GetConversationsMessagingStickerTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker too many requests response has a 4xx status code
func (o *GetConversationsMessagingStickerTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging sticker too many requests response has a 5xx status code
func (o *GetConversationsMessagingStickerTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging sticker too many requests response a status code equal to that given
func (o *GetConversationsMessagingStickerTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessagingStickerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingStickerTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingStickerTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerInternalServerError creates a GetConversationsMessagingStickerInternalServerError with default headers values
func NewGetConversationsMessagingStickerInternalServerError() *GetConversationsMessagingStickerInternalServerError {
	return &GetConversationsMessagingStickerInternalServerError{}
}

/*
GetConversationsMessagingStickerInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingStickerInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker internal server error response has a 2xx status code
func (o *GetConversationsMessagingStickerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker internal server error response has a 3xx status code
func (o *GetConversationsMessagingStickerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker internal server error response has a 4xx status code
func (o *GetConversationsMessagingStickerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging sticker internal server error response has a 5xx status code
func (o *GetConversationsMessagingStickerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging sticker internal server error response a status code equal to that given
func (o *GetConversationsMessagingStickerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessagingStickerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingStickerInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingStickerInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerServiceUnavailable creates a GetConversationsMessagingStickerServiceUnavailable with default headers values
func NewGetConversationsMessagingStickerServiceUnavailable() *GetConversationsMessagingStickerServiceUnavailable {
	return &GetConversationsMessagingStickerServiceUnavailable{}
}

/*
GetConversationsMessagingStickerServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingStickerServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker service unavailable response has a 2xx status code
func (o *GetConversationsMessagingStickerServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker service unavailable response has a 3xx status code
func (o *GetConversationsMessagingStickerServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker service unavailable response has a 4xx status code
func (o *GetConversationsMessagingStickerServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging sticker service unavailable response has a 5xx status code
func (o *GetConversationsMessagingStickerServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging sticker service unavailable response a status code equal to that given
func (o *GetConversationsMessagingStickerServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessagingStickerServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingStickerServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingStickerServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingStickerGatewayTimeout creates a GetConversationsMessagingStickerGatewayTimeout with default headers values
func NewGetConversationsMessagingStickerGatewayTimeout() *GetConversationsMessagingStickerGatewayTimeout {
	return &GetConversationsMessagingStickerGatewayTimeout{}
}

/*
GetConversationsMessagingStickerGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessagingStickerGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging sticker gateway timeout response has a 2xx status code
func (o *GetConversationsMessagingStickerGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging sticker gateway timeout response has a 3xx status code
func (o *GetConversationsMessagingStickerGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging sticker gateway timeout response has a 4xx status code
func (o *GetConversationsMessagingStickerGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging sticker gateway timeout response has a 5xx status code
func (o *GetConversationsMessagingStickerGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging sticker gateway timeout response a status code equal to that given
func (o *GetConversationsMessagingStickerGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessagingStickerGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingStickerGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/stickers/{messengerType}][%d] getConversationsMessagingStickerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingStickerGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingStickerGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
