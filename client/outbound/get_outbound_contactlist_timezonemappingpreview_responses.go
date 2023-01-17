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

// GetOutboundContactlistTimezonemappingpreviewReader is a Reader for the GetOutboundContactlistTimezonemappingpreview structure.
type GetOutboundContactlistTimezonemappingpreviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundContactlistTimezonemappingpreviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundContactlistTimezonemappingpreviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundContactlistTimezonemappingpreviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundContactlistTimezonemappingpreviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundContactlistTimezonemappingpreviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundContactlistTimezonemappingpreviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundContactlistTimezonemappingpreviewRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundContactlistTimezonemappingpreviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundContactlistTimezonemappingpreviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundContactlistTimezonemappingpreviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundContactlistTimezonemappingpreviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundContactlistTimezonemappingpreviewOK creates a GetOutboundContactlistTimezonemappingpreviewOK with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewOK() *GetOutboundContactlistTimezonemappingpreviewOK {
	return &GetOutboundContactlistTimezonemappingpreviewOK{}
}

/*
GetOutboundContactlistTimezonemappingpreviewOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundContactlistTimezonemappingpreviewOK struct {
	Payload *models.TimeZoneMappingPreview
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview o k response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview o k response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview o k response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview o k response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview o k response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) GetPayload() *models.TimeZoneMappingPreview {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeZoneMappingPreview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewBadRequest creates a GetOutboundContactlistTimezonemappingpreviewBadRequest with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewBadRequest() *GetOutboundContactlistTimezonemappingpreviewBadRequest {
	return &GetOutboundContactlistTimezonemappingpreviewBadRequest{}
}

/*
GetOutboundContactlistTimezonemappingpreviewBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundContactlistTimezonemappingpreviewBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview bad request response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview bad request response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview bad request response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview bad request response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview bad request response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewUnauthorized creates a GetOutboundContactlistTimezonemappingpreviewUnauthorized with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewUnauthorized() *GetOutboundContactlistTimezonemappingpreviewUnauthorized {
	return &GetOutboundContactlistTimezonemappingpreviewUnauthorized{}
}

/*
GetOutboundContactlistTimezonemappingpreviewUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundContactlistTimezonemappingpreviewUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview unauthorized response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview unauthorized response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview unauthorized response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview unauthorized response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview unauthorized response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewForbidden creates a GetOutboundContactlistTimezonemappingpreviewForbidden with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewForbidden() *GetOutboundContactlistTimezonemappingpreviewForbidden {
	return &GetOutboundContactlistTimezonemappingpreviewForbidden{}
}

/*
GetOutboundContactlistTimezonemappingpreviewForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundContactlistTimezonemappingpreviewForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview forbidden response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview forbidden response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview forbidden response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview forbidden response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview forbidden response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewNotFound creates a GetOutboundContactlistTimezonemappingpreviewNotFound with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewNotFound() *GetOutboundContactlistTimezonemappingpreviewNotFound {
	return &GetOutboundContactlistTimezonemappingpreviewNotFound{}
}

/*
GetOutboundContactlistTimezonemappingpreviewNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundContactlistTimezonemappingpreviewNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview not found response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview not found response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview not found response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview not found response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview not found response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewRequestTimeout creates a GetOutboundContactlistTimezonemappingpreviewRequestTimeout with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewRequestTimeout() *GetOutboundContactlistTimezonemappingpreviewRequestTimeout {
	return &GetOutboundContactlistTimezonemappingpreviewRequestTimeout{}
}

/*
GetOutboundContactlistTimezonemappingpreviewRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundContactlistTimezonemappingpreviewRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview request timeout response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview request timeout response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview request timeout response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview request timeout response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview request timeout response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge creates a GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge() *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge {
	return &GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge{}
}

/*
GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview request entity too large response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview request entity too large response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview request entity too large response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview request entity too large response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview request entity too large response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType creates a GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType() *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType {
	return &GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType{}
}

/*
GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview unsupported media type response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview unsupported media type response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview unsupported media type response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview unsupported media type response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview unsupported media type response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewTooManyRequests creates a GetOutboundContactlistTimezonemappingpreviewTooManyRequests with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewTooManyRequests() *GetOutboundContactlistTimezonemappingpreviewTooManyRequests {
	return &GetOutboundContactlistTimezonemappingpreviewTooManyRequests{}
}

/*
GetOutboundContactlistTimezonemappingpreviewTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundContactlistTimezonemappingpreviewTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview too many requests response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview too many requests response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview too many requests response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview too many requests response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview too many requests response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewInternalServerError creates a GetOutboundContactlistTimezonemappingpreviewInternalServerError with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewInternalServerError() *GetOutboundContactlistTimezonemappingpreviewInternalServerError {
	return &GetOutboundContactlistTimezonemappingpreviewInternalServerError{}
}

/*
GetOutboundContactlistTimezonemappingpreviewInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundContactlistTimezonemappingpreviewInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview internal server error response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview internal server error response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview internal server error response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview internal server error response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview internal server error response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewServiceUnavailable creates a GetOutboundContactlistTimezonemappingpreviewServiceUnavailable with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewServiceUnavailable() *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable {
	return &GetOutboundContactlistTimezonemappingpreviewServiceUnavailable{}
}

/*
GetOutboundContactlistTimezonemappingpreviewServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundContactlistTimezonemappingpreviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview service unavailable response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview service unavailable response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview service unavailable response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview service unavailable response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview service unavailable response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewGatewayTimeout creates a GetOutboundContactlistTimezonemappingpreviewGatewayTimeout with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewGatewayTimeout() *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout {
	return &GetOutboundContactlistTimezonemappingpreviewGatewayTimeout{}
}

/*
GetOutboundContactlistTimezonemappingpreviewGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundContactlistTimezonemappingpreviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist timezonemappingpreview gateway timeout response has a 2xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist timezonemappingpreview gateway timeout response has a 3xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist timezonemappingpreview gateway timeout response has a 4xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist timezonemappingpreview gateway timeout response has a 5xx status code
func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound contactlist timezonemappingpreview gateway timeout response a status code equal to that given
func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
