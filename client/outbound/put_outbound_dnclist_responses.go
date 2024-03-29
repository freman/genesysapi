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

// PutOutboundDnclistReader is a Reader for the PutOutboundDnclist structure.
type PutOutboundDnclistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundDnclistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundDnclistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundDnclistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundDnclistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundDnclistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundDnclistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundDnclistRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundDnclistConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundDnclistRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundDnclistUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundDnclistTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundDnclistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPutOutboundDnclistNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundDnclistServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundDnclistGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundDnclistOK creates a PutOutboundDnclistOK with default headers values
func NewPutOutboundDnclistOK() *PutOutboundDnclistOK {
	return &PutOutboundDnclistOK{}
}

/*
PutOutboundDnclistOK describes a response with status code 200, with default header values.

successful operation
*/
type PutOutboundDnclistOK struct {
	Payload *models.DncList
}

// IsSuccess returns true when this put outbound dnclist o k response has a 2xx status code
func (o *PutOutboundDnclistOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put outbound dnclist o k response has a 3xx status code
func (o *PutOutboundDnclistOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist o k response has a 4xx status code
func (o *PutOutboundDnclistOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound dnclist o k response has a 5xx status code
func (o *PutOutboundDnclistOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist o k response a status code equal to that given
func (o *PutOutboundDnclistOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutOutboundDnclistOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistOK  %+v", 200, o.Payload)
}

func (o *PutOutboundDnclistOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistOK  %+v", 200, o.Payload)
}

func (o *PutOutboundDnclistOK) GetPayload() *models.DncList {
	return o.Payload
}

func (o *PutOutboundDnclistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DncList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistBadRequest creates a PutOutboundDnclistBadRequest with default headers values
func NewPutOutboundDnclistBadRequest() *PutOutboundDnclistBadRequest {
	return &PutOutboundDnclistBadRequest{}
}

/*
PutOutboundDnclistBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundDnclistBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist bad request response has a 2xx status code
func (o *PutOutboundDnclistBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist bad request response has a 3xx status code
func (o *PutOutboundDnclistBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist bad request response has a 4xx status code
func (o *PutOutboundDnclistBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist bad request response has a 5xx status code
func (o *PutOutboundDnclistBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist bad request response a status code equal to that given
func (o *PutOutboundDnclistBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutOutboundDnclistBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundDnclistBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundDnclistBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistUnauthorized creates a PutOutboundDnclistUnauthorized with default headers values
func NewPutOutboundDnclistUnauthorized() *PutOutboundDnclistUnauthorized {
	return &PutOutboundDnclistUnauthorized{}
}

/*
PutOutboundDnclistUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundDnclistUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist unauthorized response has a 2xx status code
func (o *PutOutboundDnclistUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist unauthorized response has a 3xx status code
func (o *PutOutboundDnclistUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist unauthorized response has a 4xx status code
func (o *PutOutboundDnclistUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist unauthorized response has a 5xx status code
func (o *PutOutboundDnclistUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist unauthorized response a status code equal to that given
func (o *PutOutboundDnclistUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutOutboundDnclistUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundDnclistUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundDnclistUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistForbidden creates a PutOutboundDnclistForbidden with default headers values
func NewPutOutboundDnclistForbidden() *PutOutboundDnclistForbidden {
	return &PutOutboundDnclistForbidden{}
}

/*
PutOutboundDnclistForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundDnclistForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist forbidden response has a 2xx status code
func (o *PutOutboundDnclistForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist forbidden response has a 3xx status code
func (o *PutOutboundDnclistForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist forbidden response has a 4xx status code
func (o *PutOutboundDnclistForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist forbidden response has a 5xx status code
func (o *PutOutboundDnclistForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist forbidden response a status code equal to that given
func (o *PutOutboundDnclistForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutOutboundDnclistForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundDnclistForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundDnclistForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistNotFound creates a PutOutboundDnclistNotFound with default headers values
func NewPutOutboundDnclistNotFound() *PutOutboundDnclistNotFound {
	return &PutOutboundDnclistNotFound{}
}

/*
PutOutboundDnclistNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutOutboundDnclistNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist not found response has a 2xx status code
func (o *PutOutboundDnclistNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist not found response has a 3xx status code
func (o *PutOutboundDnclistNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist not found response has a 4xx status code
func (o *PutOutboundDnclistNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist not found response has a 5xx status code
func (o *PutOutboundDnclistNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist not found response a status code equal to that given
func (o *PutOutboundDnclistNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutOutboundDnclistNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundDnclistNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundDnclistNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistRequestTimeout creates a PutOutboundDnclistRequestTimeout with default headers values
func NewPutOutboundDnclistRequestTimeout() *PutOutboundDnclistRequestTimeout {
	return &PutOutboundDnclistRequestTimeout{}
}

/*
PutOutboundDnclistRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundDnclistRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist request timeout response has a 2xx status code
func (o *PutOutboundDnclistRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist request timeout response has a 3xx status code
func (o *PutOutboundDnclistRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist request timeout response has a 4xx status code
func (o *PutOutboundDnclistRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist request timeout response has a 5xx status code
func (o *PutOutboundDnclistRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist request timeout response a status code equal to that given
func (o *PutOutboundDnclistRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutOutboundDnclistRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundDnclistRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundDnclistRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistConflict creates a PutOutboundDnclistConflict with default headers values
func NewPutOutboundDnclistConflict() *PutOutboundDnclistConflict {
	return &PutOutboundDnclistConflict{}
}

/*
PutOutboundDnclistConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutOutboundDnclistConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist conflict response has a 2xx status code
func (o *PutOutboundDnclistConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist conflict response has a 3xx status code
func (o *PutOutboundDnclistConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist conflict response has a 4xx status code
func (o *PutOutboundDnclistConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist conflict response has a 5xx status code
func (o *PutOutboundDnclistConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist conflict response a status code equal to that given
func (o *PutOutboundDnclistConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutOutboundDnclistConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundDnclistConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundDnclistConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistRequestEntityTooLarge creates a PutOutboundDnclistRequestEntityTooLarge with default headers values
func NewPutOutboundDnclistRequestEntityTooLarge() *PutOutboundDnclistRequestEntityTooLarge {
	return &PutOutboundDnclistRequestEntityTooLarge{}
}

/*
PutOutboundDnclistRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutOutboundDnclistRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist request entity too large response has a 2xx status code
func (o *PutOutboundDnclistRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist request entity too large response has a 3xx status code
func (o *PutOutboundDnclistRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist request entity too large response has a 4xx status code
func (o *PutOutboundDnclistRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist request entity too large response has a 5xx status code
func (o *PutOutboundDnclistRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist request entity too large response a status code equal to that given
func (o *PutOutboundDnclistRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutOutboundDnclistRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundDnclistRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundDnclistRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistUnsupportedMediaType creates a PutOutboundDnclistUnsupportedMediaType with default headers values
func NewPutOutboundDnclistUnsupportedMediaType() *PutOutboundDnclistUnsupportedMediaType {
	return &PutOutboundDnclistUnsupportedMediaType{}
}

/*
PutOutboundDnclistUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundDnclistUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist unsupported media type response has a 2xx status code
func (o *PutOutboundDnclistUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist unsupported media type response has a 3xx status code
func (o *PutOutboundDnclistUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist unsupported media type response has a 4xx status code
func (o *PutOutboundDnclistUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist unsupported media type response has a 5xx status code
func (o *PutOutboundDnclistUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist unsupported media type response a status code equal to that given
func (o *PutOutboundDnclistUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutOutboundDnclistUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundDnclistUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundDnclistUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistTooManyRequests creates a PutOutboundDnclistTooManyRequests with default headers values
func NewPutOutboundDnclistTooManyRequests() *PutOutboundDnclistTooManyRequests {
	return &PutOutboundDnclistTooManyRequests{}
}

/*
PutOutboundDnclistTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundDnclistTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist too many requests response has a 2xx status code
func (o *PutOutboundDnclistTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist too many requests response has a 3xx status code
func (o *PutOutboundDnclistTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist too many requests response has a 4xx status code
func (o *PutOutboundDnclistTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound dnclist too many requests response has a 5xx status code
func (o *PutOutboundDnclistTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound dnclist too many requests response a status code equal to that given
func (o *PutOutboundDnclistTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutOutboundDnclistTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundDnclistTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundDnclistTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistInternalServerError creates a PutOutboundDnclistInternalServerError with default headers values
func NewPutOutboundDnclistInternalServerError() *PutOutboundDnclistInternalServerError {
	return &PutOutboundDnclistInternalServerError{}
}

/*
PutOutboundDnclistInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundDnclistInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist internal server error response has a 2xx status code
func (o *PutOutboundDnclistInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist internal server error response has a 3xx status code
func (o *PutOutboundDnclistInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist internal server error response has a 4xx status code
func (o *PutOutboundDnclistInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound dnclist internal server error response has a 5xx status code
func (o *PutOutboundDnclistInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound dnclist internal server error response a status code equal to that given
func (o *PutOutboundDnclistInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutOutboundDnclistInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundDnclistInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundDnclistInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistNotImplemented creates a PutOutboundDnclistNotImplemented with default headers values
func NewPutOutboundDnclistNotImplemented() *PutOutboundDnclistNotImplemented {
	return &PutOutboundDnclistNotImplemented{}
}

/*
PutOutboundDnclistNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type PutOutboundDnclistNotImplemented struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist not implemented response has a 2xx status code
func (o *PutOutboundDnclistNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist not implemented response has a 3xx status code
func (o *PutOutboundDnclistNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist not implemented response has a 4xx status code
func (o *PutOutboundDnclistNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound dnclist not implemented response has a 5xx status code
func (o *PutOutboundDnclistNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound dnclist not implemented response a status code equal to that given
func (o *PutOutboundDnclistNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *PutOutboundDnclistNotImplemented) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistNotImplemented  %+v", 501, o.Payload)
}

func (o *PutOutboundDnclistNotImplemented) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistNotImplemented  %+v", 501, o.Payload)
}

func (o *PutOutboundDnclistNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistServiceUnavailable creates a PutOutboundDnclistServiceUnavailable with default headers values
func NewPutOutboundDnclistServiceUnavailable() *PutOutboundDnclistServiceUnavailable {
	return &PutOutboundDnclistServiceUnavailable{}
}

/*
PutOutboundDnclistServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundDnclistServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist service unavailable response has a 2xx status code
func (o *PutOutboundDnclistServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist service unavailable response has a 3xx status code
func (o *PutOutboundDnclistServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist service unavailable response has a 4xx status code
func (o *PutOutboundDnclistServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound dnclist service unavailable response has a 5xx status code
func (o *PutOutboundDnclistServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound dnclist service unavailable response a status code equal to that given
func (o *PutOutboundDnclistServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutOutboundDnclistServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundDnclistServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundDnclistServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundDnclistGatewayTimeout creates a PutOutboundDnclistGatewayTimeout with default headers values
func NewPutOutboundDnclistGatewayTimeout() *PutOutboundDnclistGatewayTimeout {
	return &PutOutboundDnclistGatewayTimeout{}
}

/*
PutOutboundDnclistGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutOutboundDnclistGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound dnclist gateway timeout response has a 2xx status code
func (o *PutOutboundDnclistGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound dnclist gateway timeout response has a 3xx status code
func (o *PutOutboundDnclistGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound dnclist gateway timeout response has a 4xx status code
func (o *PutOutboundDnclistGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound dnclist gateway timeout response has a 5xx status code
func (o *PutOutboundDnclistGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound dnclist gateway timeout response a status code equal to that given
func (o *PutOutboundDnclistGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutOutboundDnclistGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundDnclistGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/dnclists/{dncListId}][%d] putOutboundDnclistGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundDnclistGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundDnclistGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
