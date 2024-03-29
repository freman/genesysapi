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

// PostTelephonyProvidersEdgeLogicalinterfacesReader is a Reader for the PostTelephonyProvidersEdgeLogicalinterfaces structure.
type PostTelephonyProvidersEdgeLogicalinterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeLogicalinterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesOK creates a PostTelephonyProvidersEdgeLogicalinterfacesOK with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesOK() *PostTelephonyProvidersEdgeLogicalinterfacesOK {
	return &PostTelephonyProvidersEdgeLogicalinterfacesOK{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostTelephonyProvidersEdgeLogicalinterfacesOK struct {
	Payload *models.DomainLogicalInterface
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces o k response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces o k response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces o k response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces o k response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces o k response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) GetPayload() *models.DomainLogicalInterface {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainLogicalInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesBadRequest creates a PostTelephonyProvidersEdgeLogicalinterfacesBadRequest with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesBadRequest() *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest {
	return &PostTelephonyProvidersEdgeLogicalinterfacesBadRequest{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces bad request response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces bad request response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces bad request response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces bad request response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces bad request response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesUnauthorized creates a PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesUnauthorized() *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized {
	return &PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces unauthorized response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces unauthorized response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces unauthorized response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces unauthorized response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces unauthorized response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesForbidden creates a PostTelephonyProvidersEdgeLogicalinterfacesForbidden with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesForbidden() *PostTelephonyProvidersEdgeLogicalinterfacesForbidden {
	return &PostTelephonyProvidersEdgeLogicalinterfacesForbidden{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces forbidden response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces forbidden response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces forbidden response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces forbidden response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces forbidden response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesNotFound creates a PostTelephonyProvidersEdgeLogicalinterfacesNotFound with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesNotFound() *PostTelephonyProvidersEdgeLogicalinterfacesNotFound {
	return &PostTelephonyProvidersEdgeLogicalinterfacesNotFound{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces not found response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces not found response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces not found response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces not found response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces not found response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout creates a PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout() *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout {
	return &PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces request timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces request timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces request timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces request timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces request timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge creates a PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge() *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces request entity too large response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces request entity too large response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces request entity too large response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces request entity too large response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces request entity too large response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType creates a PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType() *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces unsupported media type response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces unsupported media type response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces unsupported media type response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces unsupported media type response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces unsupported media type response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests creates a PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests() *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests {
	return &PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces too many requests response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces too many requests response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces too many requests response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces too many requests response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logicalinterfaces too many requests response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesInternalServerError creates a PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesInternalServerError() *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError {
	return &PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces internal server error response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces internal server error response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces internal server error response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces internal server error response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge logicalinterfaces internal server error response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable creates a PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable() *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable {
	return &PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces service unavailable response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces service unavailable response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces service unavailable response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces service unavailable response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge logicalinterfaces service unavailable response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout creates a PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout() *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout {
	return &PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout{}
}

/*
PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logicalinterfaces gateway timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logicalinterfaces gateway timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logicalinterfaces gateway timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logicalinterfaces gateway timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge logicalinterfaces gateway timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
