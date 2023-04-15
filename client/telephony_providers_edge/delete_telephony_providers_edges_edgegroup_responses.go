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

// DeleteTelephonyProvidersEdgesEdgegroupReader is a Reader for the DeleteTelephonyProvidersEdgesEdgegroup structure.
type DeleteTelephonyProvidersEdgesEdgegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgesEdgegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgesEdgegroupOK creates a DeleteTelephonyProvidersEdgesEdgegroupOK with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupOK() *DeleteTelephonyProvidersEdgesEdgegroupOK {
	return &DeleteTelephonyProvidersEdgesEdgegroupOK{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgesEdgegroupOK struct {
}

// IsSuccess returns true when this delete telephony providers edges edgegroup o k response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete telephony providers edges edgegroup o k response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup o k response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges edgegroup o k response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup o k response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupBadRequest creates a DeleteTelephonyProvidersEdgesEdgegroupBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupBadRequest() *DeleteTelephonyProvidersEdgesEdgegroupBadRequest {
	return &DeleteTelephonyProvidersEdgesEdgegroupBadRequest{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgesEdgegroupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup bad request response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup bad request response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup bad request response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup bad request response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup bad request response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupUnauthorized creates a DeleteTelephonyProvidersEdgesEdgegroupUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupUnauthorized() *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized {
	return &DeleteTelephonyProvidersEdgesEdgegroupUnauthorized{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgesEdgegroupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup unauthorized response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup unauthorized response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup unauthorized response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup unauthorized response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup unauthorized response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupForbidden creates a DeleteTelephonyProvidersEdgesEdgegroupForbidden with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupForbidden() *DeleteTelephonyProvidersEdgesEdgegroupForbidden {
	return &DeleteTelephonyProvidersEdgesEdgegroupForbidden{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgesEdgegroupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup forbidden response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup forbidden response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup forbidden response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup forbidden response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup forbidden response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupNotFound creates a DeleteTelephonyProvidersEdgesEdgegroupNotFound with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupNotFound() *DeleteTelephonyProvidersEdgesEdgegroupNotFound {
	return &DeleteTelephonyProvidersEdgesEdgegroupNotFound{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgesEdgegroupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup not found response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup not found response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup not found response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup not found response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup not found response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupRequestTimeout creates a DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupRequestTimeout() *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout {
	return &DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup request timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup request timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup request timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup request timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup request timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge() *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup request entity too large response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup request entity too large response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup request entity too large response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup request entity too large response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup request entity too large response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType creates a DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType() *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup unsupported media type response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup unsupported media type response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup unsupported media type response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup unsupported media type response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup unsupported media type response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupTooManyRequests creates a DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupTooManyRequests() *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests {
	return &DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup too many requests response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup too many requests response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup too many requests response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges edgegroup too many requests response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges edgegroup too many requests response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupInternalServerError creates a DeleteTelephonyProvidersEdgesEdgegroupInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupInternalServerError() *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError {
	return &DeleteTelephonyProvidersEdgesEdgegroupInternalServerError{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgesEdgegroupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup internal server error response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup internal server error response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup internal server error response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges edgegroup internal server error response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges edgegroup internal server error response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable creates a DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable() *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable {
	return &DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup service unavailable response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup service unavailable response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup service unavailable response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges edgegroup service unavailable response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges edgegroup service unavailable response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout creates a DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout() *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout {
	return &DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout{}
}

/*
DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges edgegroup gateway timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges edgegroup gateway timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges edgegroup gateway timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges edgegroup gateway timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges edgegroup gateway timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] deleteTelephonyProvidersEdgesEdgegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesEdgegroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
