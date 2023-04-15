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

// PostTelephonyProvidersEdgeRebootReader is a Reader for the PostTelephonyProvidersEdgeReboot structure.
type PostTelephonyProvidersEdgeRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgeRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeRebootBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeRebootRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeRebootRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeRebootUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeRebootTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeRebootServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeRebootGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeRebootOK creates a PostTelephonyProvidersEdgeRebootOK with default headers values
func NewPostTelephonyProvidersEdgeRebootOK() *PostTelephonyProvidersEdgeRebootOK {
	return &PostTelephonyProvidersEdgeRebootOK{}
}

/*
PostTelephonyProvidersEdgeRebootOK describes a response with status code 200, with default header values.

successful operation
*/
type PostTelephonyProvidersEdgeRebootOK struct {
	Payload string
}

// IsSuccess returns true when this post telephony providers edge reboot o k response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post telephony providers edge reboot o k response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot o k response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge reboot o k response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot o k response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostTelephonyProvidersEdgeRebootOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootOK) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootOK) GetPayload() string {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootBadRequest creates a PostTelephonyProvidersEdgeRebootBadRequest with default headers values
func NewPostTelephonyProvidersEdgeRebootBadRequest() *PostTelephonyProvidersEdgeRebootBadRequest {
	return &PostTelephonyProvidersEdgeRebootBadRequest{}
}

/*
PostTelephonyProvidersEdgeRebootBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeRebootBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot bad request response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot bad request response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot bad request response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot bad request response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot bad request response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootUnauthorized creates a PostTelephonyProvidersEdgeRebootUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeRebootUnauthorized() *PostTelephonyProvidersEdgeRebootUnauthorized {
	return &PostTelephonyProvidersEdgeRebootUnauthorized{}
}

/*
PostTelephonyProvidersEdgeRebootUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeRebootUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot unauthorized response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot unauthorized response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot unauthorized response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot unauthorized response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot unauthorized response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootForbidden creates a PostTelephonyProvidersEdgeRebootForbidden with default headers values
func NewPostTelephonyProvidersEdgeRebootForbidden() *PostTelephonyProvidersEdgeRebootForbidden {
	return &PostTelephonyProvidersEdgeRebootForbidden{}
}

/*
PostTelephonyProvidersEdgeRebootForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeRebootForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot forbidden response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot forbidden response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot forbidden response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot forbidden response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot forbidden response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootNotFound creates a PostTelephonyProvidersEdgeRebootNotFound with default headers values
func NewPostTelephonyProvidersEdgeRebootNotFound() *PostTelephonyProvidersEdgeRebootNotFound {
	return &PostTelephonyProvidersEdgeRebootNotFound{}
}

/*
PostTelephonyProvidersEdgeRebootNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeRebootNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot not found response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot not found response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot not found response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot not found response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot not found response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootRequestTimeout creates a PostTelephonyProvidersEdgeRebootRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeRebootRequestTimeout() *PostTelephonyProvidersEdgeRebootRequestTimeout {
	return &PostTelephonyProvidersEdgeRebootRequestTimeout{}
}

/*
PostTelephonyProvidersEdgeRebootRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeRebootRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot request timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot request timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot request timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot request timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot request timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootRequestEntityTooLarge creates a PostTelephonyProvidersEdgeRebootRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeRebootRequestEntityTooLarge() *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeRebootRequestEntityTooLarge{}
}

/*
PostTelephonyProvidersEdgeRebootRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostTelephonyProvidersEdgeRebootRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot request entity too large response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot request entity too large response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot request entity too large response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot request entity too large response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot request entity too large response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootUnsupportedMediaType creates a PostTelephonyProvidersEdgeRebootUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeRebootUnsupportedMediaType() *PostTelephonyProvidersEdgeRebootUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeRebootUnsupportedMediaType{}
}

/*
PostTelephonyProvidersEdgeRebootUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeRebootUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot unsupported media type response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot unsupported media type response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot unsupported media type response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot unsupported media type response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot unsupported media type response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootTooManyRequests creates a PostTelephonyProvidersEdgeRebootTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeRebootTooManyRequests() *PostTelephonyProvidersEdgeRebootTooManyRequests {
	return &PostTelephonyProvidersEdgeRebootTooManyRequests{}
}

/*
PostTelephonyProvidersEdgeRebootTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeRebootTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot too many requests response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot too many requests response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot too many requests response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge reboot too many requests response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge reboot too many requests response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootInternalServerError creates a PostTelephonyProvidersEdgeRebootInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeRebootInternalServerError() *PostTelephonyProvidersEdgeRebootInternalServerError {
	return &PostTelephonyProvidersEdgeRebootInternalServerError{}
}

/*
PostTelephonyProvidersEdgeRebootInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeRebootInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot internal server error response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot internal server error response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot internal server error response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge reboot internal server error response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge reboot internal server error response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootServiceUnavailable creates a PostTelephonyProvidersEdgeRebootServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeRebootServiceUnavailable() *PostTelephonyProvidersEdgeRebootServiceUnavailable {
	return &PostTelephonyProvidersEdgeRebootServiceUnavailable{}
}

/*
PostTelephonyProvidersEdgeRebootServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeRebootServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot service unavailable response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot service unavailable response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot service unavailable response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge reboot service unavailable response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge reboot service unavailable response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeRebootGatewayTimeout creates a PostTelephonyProvidersEdgeRebootGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeRebootGatewayTimeout() *PostTelephonyProvidersEdgeRebootGatewayTimeout {
	return &PostTelephonyProvidersEdgeRebootGatewayTimeout{}
}

/*
PostTelephonyProvidersEdgeRebootGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeRebootGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge reboot gateway timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge reboot gateway timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge reboot gateway timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge reboot gateway timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge reboot gateway timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/reboot][%d] postTelephonyProvidersEdgeRebootGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeRebootGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
