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

// PostTelephonyProvidersEdgeDiagnosticRouteReader is a Reader for the PostTelephonyProvidersEdgeDiagnosticRoute structure.
type PostTelephonyProvidersEdgeDiagnosticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeDiagnosticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteAccepted creates a PostTelephonyProvidersEdgeDiagnosticRouteAccepted with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteAccepted() *PostTelephonyProvidersEdgeDiagnosticRouteAccepted {
	return &PostTelephonyProvidersEdgeDiagnosticRouteAccepted{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteAccepted describes a response with status code 202, with default header values.

Request to get network diagnostic has been accepted
*/
type PostTelephonyProvidersEdgeDiagnosticRouteAccepted struct {
	Payload *models.EdgeNetworkDiagnostic
}

// IsSuccess returns true when this post telephony providers edge diagnostic route accepted response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post telephony providers edge diagnostic route accepted response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route accepted response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic route accepted response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route accepted response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteAccepted  %+v", 202, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteAccepted  %+v", 202, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) GetPayload() *models.EdgeNetworkDiagnostic {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnostic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteBadRequest creates a PostTelephonyProvidersEdgeDiagnosticRouteBadRequest with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteBadRequest() *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest {
	return &PostTelephonyProvidersEdgeDiagnosticRouteBadRequest{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route bad request response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route bad request response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route bad request response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route bad request response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route bad request response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteUnauthorized creates a PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteUnauthorized() *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized {
	return &PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route unauthorized response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route unauthorized response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route unauthorized response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route unauthorized response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route unauthorized response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteForbidden creates a PostTelephonyProvidersEdgeDiagnosticRouteForbidden with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteForbidden() *PostTelephonyProvidersEdgeDiagnosticRouteForbidden {
	return &PostTelephonyProvidersEdgeDiagnosticRouteForbidden{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route forbidden response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route forbidden response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route forbidden response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route forbidden response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route forbidden response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteNotFound creates a PostTelephonyProvidersEdgeDiagnosticRouteNotFound with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteNotFound() *PostTelephonyProvidersEdgeDiagnosticRouteNotFound {
	return &PostTelephonyProvidersEdgeDiagnosticRouteNotFound{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route not found response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route not found response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route not found response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route not found response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route not found response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout creates a PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout() *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout {
	return &PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route request timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route request timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route request timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route request timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route request timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge creates a PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge() *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route request entity too large response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route request entity too large response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route request entity too large response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route request entity too large response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route request entity too large response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType creates a PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType() *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route unsupported media type response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route unsupported media type response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route unsupported media type response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route unsupported media type response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route unsupported media type response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests creates a PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests() *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests {
	return &PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route too many requests response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route too many requests response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route too many requests response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic route too many requests response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic route too many requests response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteInternalServerError creates a PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteInternalServerError() *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError {
	return &PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route internal server error response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route internal server error response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route internal server error response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic route internal server error response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge diagnostic route internal server error response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable creates a PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable() *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable {
	return &PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route service unavailable response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route service unavailable response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route service unavailable response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic route service unavailable response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge diagnostic route service unavailable response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout creates a PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout() *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout {
	return &PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout{}
}

/*
PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic route gateway timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic route gateway timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic route gateway timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic route gateway timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge diagnostic route gateway timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route][%d] postTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticRouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
