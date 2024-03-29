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

// PutTelephonyProvidersEdgesOutboundrouteReader is a Reader for the PutTelephonyProvidersEdgesOutboundroute structure.
type PutTelephonyProvidersEdgesOutboundrouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgesOutboundrouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgesOutboundrouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgesOutboundrouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgesOutboundrouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgesOutboundrouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgesOutboundrouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutTelephonyProvidersEdgesOutboundrouteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgesOutboundrouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgesOutboundrouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgesOutboundrouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgesOutboundrouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgesOutboundrouteOK creates a PutTelephonyProvidersEdgesOutboundrouteOK with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteOK() *PutTelephonyProvidersEdgesOutboundrouteOK {
	return &PutTelephonyProvidersEdgesOutboundrouteOK{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteOK describes a response with status code 200, with default header values.

successful operation
*/
type PutTelephonyProvidersEdgesOutboundrouteOK struct {
	Payload *models.OutboundRoute
}

// IsSuccess returns true when this put telephony providers edges outboundroute o k response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put telephony providers edges outboundroute o k response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute o k response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edges outboundroute o k response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute o k response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutTelephonyProvidersEdgesOutboundrouteOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteOK) GetPayload() *models.OutboundRoute {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteBadRequest creates a PutTelephonyProvidersEdgesOutboundrouteBadRequest with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteBadRequest() *PutTelephonyProvidersEdgesOutboundrouteBadRequest {
	return &PutTelephonyProvidersEdgesOutboundrouteBadRequest{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgesOutboundrouteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute bad request response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute bad request response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute bad request response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute bad request response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute bad request response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteUnauthorized creates a PutTelephonyProvidersEdgesOutboundrouteUnauthorized with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteUnauthorized() *PutTelephonyProvidersEdgesOutboundrouteUnauthorized {
	return &PutTelephonyProvidersEdgesOutboundrouteUnauthorized{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgesOutboundrouteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute unauthorized response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute unauthorized response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute unauthorized response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute unauthorized response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute unauthorized response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteForbidden creates a PutTelephonyProvidersEdgesOutboundrouteForbidden with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteForbidden() *PutTelephonyProvidersEdgesOutboundrouteForbidden {
	return &PutTelephonyProvidersEdgesOutboundrouteForbidden{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgesOutboundrouteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute forbidden response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute forbidden response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute forbidden response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute forbidden response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute forbidden response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteNotFound creates a PutTelephonyProvidersEdgesOutboundrouteNotFound with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteNotFound() *PutTelephonyProvidersEdgesOutboundrouteNotFound {
	return &PutTelephonyProvidersEdgesOutboundrouteNotFound{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgesOutboundrouteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute not found response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute not found response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute not found response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute not found response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute not found response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteRequestTimeout creates a PutTelephonyProvidersEdgesOutboundrouteRequestTimeout with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteRequestTimeout() *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout {
	return &PutTelephonyProvidersEdgesOutboundrouteRequestTimeout{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutTelephonyProvidersEdgesOutboundrouteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute request timeout response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute request timeout response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute request timeout response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute request timeout response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute request timeout response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge creates a PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge() *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute request entity too large response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute request entity too large response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute request entity too large response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute request entity too large response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute request entity too large response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType creates a PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType() *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType {
	return &PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute unsupported media type response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute unsupported media type response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute unsupported media type response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute unsupported media type response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute unsupported media type response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteTooManyRequests creates a PutTelephonyProvidersEdgesOutboundrouteTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteTooManyRequests() *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests {
	return &PutTelephonyProvidersEdgesOutboundrouteTooManyRequests{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutTelephonyProvidersEdgesOutboundrouteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute too many requests response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute too many requests response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute too many requests response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put telephony providers edges outboundroute too many requests response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put telephony providers edges outboundroute too many requests response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteInternalServerError creates a PutTelephonyProvidersEdgesOutboundrouteInternalServerError with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteInternalServerError() *PutTelephonyProvidersEdgesOutboundrouteInternalServerError {
	return &PutTelephonyProvidersEdgesOutboundrouteInternalServerError{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgesOutboundrouteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute internal server error response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute internal server error response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute internal server error response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edges outboundroute internal server error response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put telephony providers edges outboundroute internal server error response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteServiceUnavailable creates a PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteServiceUnavailable() *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable {
	return &PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute service unavailable response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute service unavailable response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute service unavailable response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edges outboundroute service unavailable response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put telephony providers edges outboundroute service unavailable response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesOutboundrouteGatewayTimeout creates a PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgesOutboundrouteGatewayTimeout() *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout {
	return &PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout{}
}

/*
PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put telephony providers edges outboundroute gateway timeout response has a 2xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put telephony providers edges outboundroute gateway timeout response has a 3xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put telephony providers edges outboundroute gateway timeout response has a 4xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put telephony providers edges outboundroute gateway timeout response has a 5xx status code
func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put telephony providers edges outboundroute gateway timeout response a status code equal to that given
func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}][%d] putTelephonyProvidersEdgesOutboundrouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesOutboundrouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
