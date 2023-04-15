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

// DeleteTelephonyProvidersEdgesPhoneReader is a Reader for the DeleteTelephonyProvidersEdgesPhone structure.
type DeleteTelephonyProvidersEdgesPhoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgesPhoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgesPhoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgesPhoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgesPhoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgesPhoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgesPhoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTelephonyProvidersEdgesPhoneRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgesPhoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgesPhoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgesPhoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgesPhoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgesPhoneOK creates a DeleteTelephonyProvidersEdgesPhoneOK with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneOK() *DeleteTelephonyProvidersEdgesPhoneOK {
	return &DeleteTelephonyProvidersEdgesPhoneOK{}
}

/*
DeleteTelephonyProvidersEdgesPhoneOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgesPhoneOK struct {
}

// IsSuccess returns true when this delete telephony providers edges phone o k response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete telephony providers edges phone o k response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone o k response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges phone o k response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone o k response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteTelephonyProvidersEdgesPhoneOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesPhoneOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesPhoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneBadRequest creates a DeleteTelephonyProvidersEdgesPhoneBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneBadRequest() *DeleteTelephonyProvidersEdgesPhoneBadRequest {
	return &DeleteTelephonyProvidersEdgesPhoneBadRequest{}
}

/*
DeleteTelephonyProvidersEdgesPhoneBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgesPhoneBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone bad request response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone bad request response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone bad request response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone bad request response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone bad request response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneUnauthorized creates a DeleteTelephonyProvidersEdgesPhoneUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneUnauthorized() *DeleteTelephonyProvidersEdgesPhoneUnauthorized {
	return &DeleteTelephonyProvidersEdgesPhoneUnauthorized{}
}

/*
DeleteTelephonyProvidersEdgesPhoneUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgesPhoneUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone unauthorized response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone unauthorized response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone unauthorized response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone unauthorized response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone unauthorized response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneForbidden creates a DeleteTelephonyProvidersEdgesPhoneForbidden with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneForbidden() *DeleteTelephonyProvidersEdgesPhoneForbidden {
	return &DeleteTelephonyProvidersEdgesPhoneForbidden{}
}

/*
DeleteTelephonyProvidersEdgesPhoneForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgesPhoneForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone forbidden response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone forbidden response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone forbidden response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone forbidden response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone forbidden response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneNotFound creates a DeleteTelephonyProvidersEdgesPhoneNotFound with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneNotFound() *DeleteTelephonyProvidersEdgesPhoneNotFound {
	return &DeleteTelephonyProvidersEdgesPhoneNotFound{}
}

/*
DeleteTelephonyProvidersEdgesPhoneNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgesPhoneNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone not found response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone not found response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone not found response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone not found response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone not found response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneRequestTimeout creates a DeleteTelephonyProvidersEdgesPhoneRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneRequestTimeout() *DeleteTelephonyProvidersEdgesPhoneRequestTimeout {
	return &DeleteTelephonyProvidersEdgesPhoneRequestTimeout{}
}

/*
DeleteTelephonyProvidersEdgesPhoneRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgesPhoneRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone request timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone request timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone request timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone request timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone request timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge() *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge{}
}

/*
DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone request entity too large response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone request entity too large response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone request entity too large response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone request entity too large response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone request entity too large response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType creates a DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType() *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType{}
}

/*
DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone unsupported media type response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone unsupported media type response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone unsupported media type response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone unsupported media type response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone unsupported media type response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneTooManyRequests creates a DeleteTelephonyProvidersEdgesPhoneTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneTooManyRequests() *DeleteTelephonyProvidersEdgesPhoneTooManyRequests {
	return &DeleteTelephonyProvidersEdgesPhoneTooManyRequests{}
}

/*
DeleteTelephonyProvidersEdgesPhoneTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgesPhoneTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone too many requests response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone too many requests response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone too many requests response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges phone too many requests response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges phone too many requests response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneInternalServerError creates a DeleteTelephonyProvidersEdgesPhoneInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneInternalServerError() *DeleteTelephonyProvidersEdgesPhoneInternalServerError {
	return &DeleteTelephonyProvidersEdgesPhoneInternalServerError{}
}

/*
DeleteTelephonyProvidersEdgesPhoneInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgesPhoneInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone internal server error response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone internal server error response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone internal server error response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges phone internal server error response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges phone internal server error response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneServiceUnavailable creates a DeleteTelephonyProvidersEdgesPhoneServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneServiceUnavailable() *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable {
	return &DeleteTelephonyProvidersEdgesPhoneServiceUnavailable{}
}

/*
DeleteTelephonyProvidersEdgesPhoneServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgesPhoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone service unavailable response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone service unavailable response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone service unavailable response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges phone service unavailable response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges phone service unavailable response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesPhoneGatewayTimeout creates a DeleteTelephonyProvidersEdgesPhoneGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesPhoneGatewayTimeout() *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout {
	return &DeleteTelephonyProvidersEdgesPhoneGatewayTimeout{}
}

/*
DeleteTelephonyProvidersEdgesPhoneGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgesPhoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges phone gateway timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges phone gateway timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges phone gateway timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges phone gateway timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges phone gateway timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/phones/{phoneId}][%d] deleteTelephonyProvidersEdgesPhoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesPhoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
