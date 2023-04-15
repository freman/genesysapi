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

// GetTelephonyProvidersEdgeDiagnosticRouteReader is a Reader for the GetTelephonyProvidersEdgeDiagnosticRoute structure.
type GetTelephonyProvidersEdgeDiagnosticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeDiagnosticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteOK creates a GetTelephonyProvidersEdgeDiagnosticRouteOK with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteOK() *GetTelephonyProvidersEdgeDiagnosticRouteOK {
	return &GetTelephonyProvidersEdgeDiagnosticRouteOK{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteOK describes a response with status code 200, with default header values.

Request to get network diagnostic was successful.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteOK struct {
	Payload *models.EdgeNetworkDiagnosticResponse
}

// IsSuccess returns true when this get telephony providers edge diagnostic route o k response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get telephony providers edge diagnostic route o k response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route o k response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge diagnostic route o k response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route o k response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) GetPayload() *models.EdgeNetworkDiagnosticResponse {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnosticResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteBadRequest creates a GetTelephonyProvidersEdgeDiagnosticRouteBadRequest with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteBadRequest() *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest {
	return &GetTelephonyProvidersEdgeDiagnosticRouteBadRequest{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route bad request response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route bad request response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route bad request response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route bad request response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route bad request response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteUnauthorized creates a GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteUnauthorized() *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized {
	return &GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route unauthorized response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route unauthorized response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route unauthorized response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route unauthorized response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route unauthorized response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteForbidden creates a GetTelephonyProvidersEdgeDiagnosticRouteForbidden with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteForbidden() *GetTelephonyProvidersEdgeDiagnosticRouteForbidden {
	return &GetTelephonyProvidersEdgeDiagnosticRouteForbidden{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route forbidden response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route forbidden response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route forbidden response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route forbidden response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route forbidden response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteNotFound creates a GetTelephonyProvidersEdgeDiagnosticRouteNotFound with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteNotFound() *GetTelephonyProvidersEdgeDiagnosticRouteNotFound {
	return &GetTelephonyProvidersEdgeDiagnosticRouteNotFound{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route not found response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route not found response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route not found response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route not found response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route not found response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout creates a GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout() *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout {
	return &GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route request timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route request timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route request timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route request timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route request timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge creates a GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge() *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route request entity too large response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route request entity too large response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route request entity too large response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route request entity too large response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route request entity too large response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType creates a GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType() *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route unsupported media type response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route unsupported media type response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route unsupported media type response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route unsupported media type response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route unsupported media type response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests creates a GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests() *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests {
	return &GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route too many requests response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route too many requests response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route too many requests response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get telephony providers edge diagnostic route too many requests response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get telephony providers edge diagnostic route too many requests response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteInternalServerError creates a GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteInternalServerError() *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError {
	return &GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route internal server error response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route internal server error response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route internal server error response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge diagnostic route internal server error response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge diagnostic route internal server error response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable creates a GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable() *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable {
	return &GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route service unavailable response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route service unavailable response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route service unavailable response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge diagnostic route service unavailable response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge diagnostic route service unavailable response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout creates a GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout() *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout {
	return &GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout{}
}

/*
GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get telephony providers edge diagnostic route gateway timeout response has a 2xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get telephony providers edge diagnostic route gateway timeout response has a 3xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get telephony providers edge diagnostic route gateway timeout response has a 4xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get telephony providers edge diagnostic route gateway timeout response has a 5xx status code
func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get telephony providers edge diagnostic route gateway timeout response a status code equal to that given
func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] getTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
