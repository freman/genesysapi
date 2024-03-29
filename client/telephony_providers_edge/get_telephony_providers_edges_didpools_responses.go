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

// GetTelephonyProvidersEdgesDidpoolsReader is a Reader for the GetTelephonyProvidersEdgesDidpools structure.
type GetTelephonyProvidersEdgesDidpoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesDidpoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesDidpoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesDidpoolsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesDidpoolsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesDidpoolsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesDidpoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesDidpoolsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesDidpoolsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesDidpoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesDidpoolsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesDidpoolsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesDidpoolsOK creates a GetTelephonyProvidersEdgesDidpoolsOK with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsOK() *GetTelephonyProvidersEdgesDidpoolsOK {
	return &GetTelephonyProvidersEdgesDidpoolsOK{}
}

/*
GetTelephonyProvidersEdgesDidpoolsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesDidpoolsOK struct {
	Payload *models.DIDPoolEntityListing
}

// IsSuccess returns true when this get telephony providers edges didpools o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edges didpools o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgesDidpoolsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsOK) GetPayload() *models.DIDPoolEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DIDPoolEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsBadRequest creates a GetTelephonyProvidersEdgesDidpoolsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsBadRequest() *GetTelephonyProvidersEdgesDidpoolsBadRequest {
	return &GetTelephonyProvidersEdgesDidpoolsBadRequest{}
}

/*
GetTelephonyProvidersEdgesDidpoolsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesDidpoolsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsUnauthorized creates a GetTelephonyProvidersEdgesDidpoolsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsUnauthorized() *GetTelephonyProvidersEdgesDidpoolsUnauthorized {
	return &GetTelephonyProvidersEdgesDidpoolsUnauthorized{}
}

/*
GetTelephonyProvidersEdgesDidpoolsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesDidpoolsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsForbidden creates a GetTelephonyProvidersEdgesDidpoolsForbidden with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsForbidden() *GetTelephonyProvidersEdgesDidpoolsForbidden {
	return &GetTelephonyProvidersEdgesDidpoolsForbidden{}
}

/*
GetTelephonyProvidersEdgesDidpoolsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesDidpoolsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsNotFound creates a GetTelephonyProvidersEdgesDidpoolsNotFound with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsNotFound() *GetTelephonyProvidersEdgesDidpoolsNotFound {
	return &GetTelephonyProvidersEdgesDidpoolsNotFound{}
}

/*
GetTelephonyProvidersEdgesDidpoolsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesDidpoolsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsRequestTimeout creates a GetTelephonyProvidersEdgesDidpoolsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsRequestTimeout() *GetTelephonyProvidersEdgesDidpoolsRequestTimeout {
	return &GetTelephonyProvidersEdgesDidpoolsRequestTimeout{}
}

/*
GetTelephonyProvidersEdgesDidpoolsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesDidpoolsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge() *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType creates a GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType() *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsTooManyRequests creates a GetTelephonyProvidersEdgesDidpoolsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsTooManyRequests() *GetTelephonyProvidersEdgesDidpoolsTooManyRequests {
	return &GetTelephonyProvidersEdgesDidpoolsTooManyRequests{}
}

/*
GetTelephonyProvidersEdgesDidpoolsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesDidpoolsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edges didpools too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edges didpools too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsInternalServerError creates a GetTelephonyProvidersEdgesDidpoolsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsInternalServerError() *GetTelephonyProvidersEdgesDidpoolsInternalServerError {
	return &GetTelephonyProvidersEdgesDidpoolsInternalServerError{}
}

/*
GetTelephonyProvidersEdgesDidpoolsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesDidpoolsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges didpools internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsServiceUnavailable creates a GetTelephonyProvidersEdgesDidpoolsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsServiceUnavailable() *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable {
	return &GetTelephonyProvidersEdgesDidpoolsServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgesDidpoolsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesDidpoolsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges didpools service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidpoolsGatewayTimeout creates a GetTelephonyProvidersEdgesDidpoolsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidpoolsGatewayTimeout() *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout {
	return &GetTelephonyProvidersEdgesDidpoolsGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgesDidpoolsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesDidpoolsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edges didpools gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edges didpools gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edges didpools gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edges didpools gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edges didpools gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/didpools][%d] getTelephonyProvidersEdgesDidpoolsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidpoolsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
