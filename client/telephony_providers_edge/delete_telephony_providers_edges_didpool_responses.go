// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteTelephonyProvidersEdgesDidpoolReader is a Reader for the DeleteTelephonyProvidersEdgesDidpool structure.
type DeleteTelephonyProvidersEdgesDidpoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgesDidpoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgesDidpoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgesDidpoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgesDidpoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgesDidpoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgesDidpoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTelephonyProvidersEdgesDidpoolRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteTelephonyProvidersEdgesDidpoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgesDidpoolTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgesDidpoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgesDidpoolServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgesDidpoolGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgesDidpoolOK creates a DeleteTelephonyProvidersEdgesDidpoolOK with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolOK() *DeleteTelephonyProvidersEdgesDidpoolOK {
	return &DeleteTelephonyProvidersEdgesDidpoolOK{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgesDidpoolOK struct {
}

// IsSuccess returns true when this delete telephony providers edges didpool o k response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete telephony providers edges didpool o k response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool o k response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges didpool o k response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool o k response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteTelephonyProvidersEdgesDidpoolOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolBadRequest creates a DeleteTelephonyProvidersEdgesDidpoolBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolBadRequest() *DeleteTelephonyProvidersEdgesDidpoolBadRequest {
	return &DeleteTelephonyProvidersEdgesDidpoolBadRequest{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgesDidpoolBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool bad request response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool bad request response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool bad request response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool bad request response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool bad request response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolUnauthorized creates a DeleteTelephonyProvidersEdgesDidpoolUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolUnauthorized() *DeleteTelephonyProvidersEdgesDidpoolUnauthorized {
	return &DeleteTelephonyProvidersEdgesDidpoolUnauthorized{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgesDidpoolUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool unauthorized response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool unauthorized response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool unauthorized response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool unauthorized response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool unauthorized response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolForbidden creates a DeleteTelephonyProvidersEdgesDidpoolForbidden with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolForbidden() *DeleteTelephonyProvidersEdgesDidpoolForbidden {
	return &DeleteTelephonyProvidersEdgesDidpoolForbidden{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgesDidpoolForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool forbidden response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool forbidden response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool forbidden response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool forbidden response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool forbidden response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolNotFound creates a DeleteTelephonyProvidersEdgesDidpoolNotFound with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolNotFound() *DeleteTelephonyProvidersEdgesDidpoolNotFound {
	return &DeleteTelephonyProvidersEdgesDidpoolNotFound{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgesDidpoolNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool not found response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool not found response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool not found response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool not found response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool not found response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolRequestTimeout creates a DeleteTelephonyProvidersEdgesDidpoolRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolRequestTimeout() *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout {
	return &DeleteTelephonyProvidersEdgesDidpoolRequestTimeout{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgesDidpoolRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool request timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool request timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool request timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool request timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool request timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolConflict creates a DeleteTelephonyProvidersEdgesDidpoolConflict with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolConflict() *DeleteTelephonyProvidersEdgesDidpoolConflict {
	return &DeleteTelephonyProvidersEdgesDidpoolConflict{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteTelephonyProvidersEdgesDidpoolConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool conflict response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool conflict response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool conflict response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool conflict response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool conflict response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolConflict  %+v", 409, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolConflict  %+v", 409, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge() *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool request entity too large response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool request entity too large response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool request entity too large response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool request entity too large response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool request entity too large response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType creates a DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType() *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool unsupported media type response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool unsupported media type response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool unsupported media type response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool unsupported media type response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool unsupported media type response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolTooManyRequests creates a DeleteTelephonyProvidersEdgesDidpoolTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolTooManyRequests() *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests {
	return &DeleteTelephonyProvidersEdgesDidpoolTooManyRequests{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgesDidpoolTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool too many requests response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool too many requests response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool too many requests response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges didpool too many requests response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges didpool too many requests response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolInternalServerError creates a DeleteTelephonyProvidersEdgesDidpoolInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolInternalServerError() *DeleteTelephonyProvidersEdgesDidpoolInternalServerError {
	return &DeleteTelephonyProvidersEdgesDidpoolInternalServerError{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgesDidpoolInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool internal server error response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool internal server error response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool internal server error response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges didpool internal server error response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges didpool internal server error response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolServiceUnavailable creates a DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolServiceUnavailable() *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable {
	return &DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool service unavailable response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool service unavailable response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool service unavailable response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges didpool service unavailable response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges didpool service unavailable response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesDidpoolGatewayTimeout creates a DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesDidpoolGatewayTimeout() *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout {
	return &DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout{}
}

/*
DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges didpool gateway timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges didpool gateway timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges didpool gateway timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges didpool gateway timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges didpool gateway timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] deleteTelephonyProvidersEdgesDidpoolGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesDidpoolGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
