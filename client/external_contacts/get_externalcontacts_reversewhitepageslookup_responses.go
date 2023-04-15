// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetExternalcontactsReversewhitepageslookupReader is a Reader for the GetExternalcontactsReversewhitepageslookup structure.
type GetExternalcontactsReversewhitepageslookupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsReversewhitepageslookupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsReversewhitepageslookupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsReversewhitepageslookupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsReversewhitepageslookupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsReversewhitepageslookupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsReversewhitepageslookupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsReversewhitepageslookupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsReversewhitepageslookupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsReversewhitepageslookupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsReversewhitepageslookupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsReversewhitepageslookupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsReversewhitepageslookupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsReversewhitepageslookupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsReversewhitepageslookupOK creates a GetExternalcontactsReversewhitepageslookupOK with default headers values
func NewGetExternalcontactsReversewhitepageslookupOK() *GetExternalcontactsReversewhitepageslookupOK {
	return &GetExternalcontactsReversewhitepageslookupOK{}
}

/*
GetExternalcontactsReversewhitepageslookupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsReversewhitepageslookupOK struct {
	Payload *models.ReverseWhitepagesLookupResult
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup o k response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup o k response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup o k response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup o k response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup o k response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsReversewhitepageslookupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupOK) GetPayload() *models.ReverseWhitepagesLookupResult {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReverseWhitepagesLookupResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupBadRequest creates a GetExternalcontactsReversewhitepageslookupBadRequest with default headers values
func NewGetExternalcontactsReversewhitepageslookupBadRequest() *GetExternalcontactsReversewhitepageslookupBadRequest {
	return &GetExternalcontactsReversewhitepageslookupBadRequest{}
}

/*
GetExternalcontactsReversewhitepageslookupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsReversewhitepageslookupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup bad request response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup bad request response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup bad request response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup bad request response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup bad request response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupUnauthorized creates a GetExternalcontactsReversewhitepageslookupUnauthorized with default headers values
func NewGetExternalcontactsReversewhitepageslookupUnauthorized() *GetExternalcontactsReversewhitepageslookupUnauthorized {
	return &GetExternalcontactsReversewhitepageslookupUnauthorized{}
}

/*
GetExternalcontactsReversewhitepageslookupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsReversewhitepageslookupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup unauthorized response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup unauthorized response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup unauthorized response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup unauthorized response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup unauthorized response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupForbidden creates a GetExternalcontactsReversewhitepageslookupForbidden with default headers values
func NewGetExternalcontactsReversewhitepageslookupForbidden() *GetExternalcontactsReversewhitepageslookupForbidden {
	return &GetExternalcontactsReversewhitepageslookupForbidden{}
}

/*
GetExternalcontactsReversewhitepageslookupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsReversewhitepageslookupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup forbidden response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup forbidden response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup forbidden response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup forbidden response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup forbidden response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupNotFound creates a GetExternalcontactsReversewhitepageslookupNotFound with default headers values
func NewGetExternalcontactsReversewhitepageslookupNotFound() *GetExternalcontactsReversewhitepageslookupNotFound {
	return &GetExternalcontactsReversewhitepageslookupNotFound{}
}

/*
GetExternalcontactsReversewhitepageslookupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsReversewhitepageslookupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup not found response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup not found response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup not found response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup not found response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup not found response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupRequestTimeout creates a GetExternalcontactsReversewhitepageslookupRequestTimeout with default headers values
func NewGetExternalcontactsReversewhitepageslookupRequestTimeout() *GetExternalcontactsReversewhitepageslookupRequestTimeout {
	return &GetExternalcontactsReversewhitepageslookupRequestTimeout{}
}

/*
GetExternalcontactsReversewhitepageslookupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsReversewhitepageslookupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup request timeout response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup request timeout response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup request timeout response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup request timeout response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup request timeout response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupRequestEntityTooLarge creates a GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge with default headers values
func NewGetExternalcontactsReversewhitepageslookupRequestEntityTooLarge() *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge {
	return &GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge{}
}

/*
GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup request entity too large response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup request entity too large response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup request entity too large response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup request entity too large response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup request entity too large response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupUnsupportedMediaType creates a GetExternalcontactsReversewhitepageslookupUnsupportedMediaType with default headers values
func NewGetExternalcontactsReversewhitepageslookupUnsupportedMediaType() *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType {
	return &GetExternalcontactsReversewhitepageslookupUnsupportedMediaType{}
}

/*
GetExternalcontactsReversewhitepageslookupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsReversewhitepageslookupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup unsupported media type response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup unsupported media type response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup unsupported media type response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup unsupported media type response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup unsupported media type response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupTooManyRequests creates a GetExternalcontactsReversewhitepageslookupTooManyRequests with default headers values
func NewGetExternalcontactsReversewhitepageslookupTooManyRequests() *GetExternalcontactsReversewhitepageslookupTooManyRequests {
	return &GetExternalcontactsReversewhitepageslookupTooManyRequests{}
}

/*
GetExternalcontactsReversewhitepageslookupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsReversewhitepageslookupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup too many requests response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup too many requests response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup too many requests response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup too many requests response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup too many requests response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupInternalServerError creates a GetExternalcontactsReversewhitepageslookupInternalServerError with default headers values
func NewGetExternalcontactsReversewhitepageslookupInternalServerError() *GetExternalcontactsReversewhitepageslookupInternalServerError {
	return &GetExternalcontactsReversewhitepageslookupInternalServerError{}
}

/*
GetExternalcontactsReversewhitepageslookupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsReversewhitepageslookupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup internal server error response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup internal server error response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup internal server error response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup internal server error response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup internal server error response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupServiceUnavailable creates a GetExternalcontactsReversewhitepageslookupServiceUnavailable with default headers values
func NewGetExternalcontactsReversewhitepageslookupServiceUnavailable() *GetExternalcontactsReversewhitepageslookupServiceUnavailable {
	return &GetExternalcontactsReversewhitepageslookupServiceUnavailable{}
}

/*
GetExternalcontactsReversewhitepageslookupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsReversewhitepageslookupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup service unavailable response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup service unavailable response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup service unavailable response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup service unavailable response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup service unavailable response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupGatewayTimeout creates a GetExternalcontactsReversewhitepageslookupGatewayTimeout with default headers values
func NewGetExternalcontactsReversewhitepageslookupGatewayTimeout() *GetExternalcontactsReversewhitepageslookupGatewayTimeout {
	return &GetExternalcontactsReversewhitepageslookupGatewayTimeout{}
}

/*
GetExternalcontactsReversewhitepageslookupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsReversewhitepageslookupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts reversewhitepageslookup gateway timeout response has a 2xx status code
func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts reversewhitepageslookup gateway timeout response has a 3xx status code
func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts reversewhitepageslookup gateway timeout response has a 4xx status code
func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts reversewhitepageslookup gateway timeout response has a 5xx status code
func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts reversewhitepageslookup gateway timeout response a status code equal to that given
func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
