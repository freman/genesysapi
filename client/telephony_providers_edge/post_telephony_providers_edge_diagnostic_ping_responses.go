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

// PostTelephonyProvidersEdgeDiagnosticPingReader is a Reader for the PostTelephonyProvidersEdgeDiagnosticPing structure.
type PostTelephonyProvidersEdgeDiagnosticPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeDiagnosticPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticPingAccepted creates a PostTelephonyProvidersEdgeDiagnosticPingAccepted with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingAccepted() *PostTelephonyProvidersEdgeDiagnosticPingAccepted {
	return &PostTelephonyProvidersEdgeDiagnosticPingAccepted{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingAccepted describes a response with status code 202, with default header values.

Request to get network diagnostic has been accepted
*/
type PostTelephonyProvidersEdgeDiagnosticPingAccepted struct {
	Payload *models.EdgeNetworkDiagnostic
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping accepted response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping accepted response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping accepted response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic ping accepted response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping accepted response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingAccepted  %+v", 202, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingAccepted  %+v", 202, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) GetPayload() *models.EdgeNetworkDiagnostic {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnostic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingBadRequest creates a PostTelephonyProvidersEdgeDiagnosticPingBadRequest with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingBadRequest() *PostTelephonyProvidersEdgeDiagnosticPingBadRequest {
	return &PostTelephonyProvidersEdgeDiagnosticPingBadRequest{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeDiagnosticPingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping bad request response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping bad request response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping bad request response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping bad request response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping bad request response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingUnauthorized creates a PostTelephonyProvidersEdgeDiagnosticPingUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingUnauthorized() *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized {
	return &PostTelephonyProvidersEdgeDiagnosticPingUnauthorized{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeDiagnosticPingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping unauthorized response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping unauthorized response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping unauthorized response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping unauthorized response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping unauthorized response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingForbidden creates a PostTelephonyProvidersEdgeDiagnosticPingForbidden with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingForbidden() *PostTelephonyProvidersEdgeDiagnosticPingForbidden {
	return &PostTelephonyProvidersEdgeDiagnosticPingForbidden{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeDiagnosticPingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping forbidden response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping forbidden response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping forbidden response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping forbidden response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping forbidden response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingNotFound creates a PostTelephonyProvidersEdgeDiagnosticPingNotFound with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingNotFound() *PostTelephonyProvidersEdgeDiagnosticPingNotFound {
	return &PostTelephonyProvidersEdgeDiagnosticPingNotFound{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeDiagnosticPingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping not found response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping not found response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping not found response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping not found response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping not found response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingRequestTimeout creates a PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingRequestTimeout() *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout {
	return &PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping request timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping request timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping request timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping request timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping request timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge creates a PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge() *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping request entity too large response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping request entity too large response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping request entity too large response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping request entity too large response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping request entity too large response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType creates a PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType() *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping unsupported media type response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping unsupported media type response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping unsupported media type response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping unsupported media type response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping unsupported media type response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingTooManyRequests creates a PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingTooManyRequests() *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests {
	return &PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping too many requests response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping too many requests response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping too many requests response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge diagnostic ping too many requests response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge diagnostic ping too many requests response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingInternalServerError creates a PostTelephonyProvidersEdgeDiagnosticPingInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingInternalServerError() *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError {
	return &PostTelephonyProvidersEdgeDiagnosticPingInternalServerError{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeDiagnosticPingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping internal server error response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping internal server error response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping internal server error response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic ping internal server error response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge diagnostic ping internal server error response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable creates a PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable() *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable {
	return &PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping service unavailable response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping service unavailable response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping service unavailable response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic ping service unavailable response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge diagnostic ping service unavailable response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout creates a PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout() *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout {
	return &PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout{}
}

/*
PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge diagnostic ping gateway timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge diagnostic ping gateway timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge diagnostic ping gateway timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge diagnostic ping gateway timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge diagnostic ping gateway timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping][%d] postTelephonyProvidersEdgeDiagnosticPingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeDiagnosticPingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
