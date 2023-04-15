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

// DeleteTelephonyProvidersEdgeSoftwareupdateReader is a Reader for the DeleteTelephonyProvidersEdgeSoftwareupdate structure.
type DeleteTelephonyProvidersEdgeSoftwareupdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateOK creates a DeleteTelephonyProvidersEdgeSoftwareupdateOK with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateOK() *DeleteTelephonyProvidersEdgeSoftwareupdateOK {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateOK{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateOK struct {
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate o k response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate o k response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate o k response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge softwareupdate o k response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate o k response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateBadRequest creates a DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateBadRequest() *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate bad request response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate bad request response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate bad request response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate bad request response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate bad request response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized creates a DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized() *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate unauthorized response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate unauthorized response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate unauthorized response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate unauthorized response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate unauthorized response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateForbidden creates a DeleteTelephonyProvidersEdgeSoftwareupdateForbidden with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateForbidden() *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateForbidden{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate forbidden response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate forbidden response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate forbidden response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate forbidden response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate forbidden response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateNotFound creates a DeleteTelephonyProvidersEdgeSoftwareupdateNotFound with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateNotFound() *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateNotFound{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate not found response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate not found response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate not found response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate not found response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate not found response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout creates a DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout() *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate request timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate request timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate request timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate request timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate request timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge() *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate request entity too large response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate request entity too large response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate request entity too large response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate request entity too large response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate request entity too large response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType creates a DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType() *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate unsupported media type response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate unsupported media type response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate unsupported media type response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate unsupported media type response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate unsupported media type response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests creates a DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests() *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate too many requests response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate too many requests response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate too many requests response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edge softwareupdate too many requests response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edge softwareupdate too many requests response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError creates a DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError() *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate internal server error response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate internal server error response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate internal server error response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge softwareupdate internal server error response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edge softwareupdate internal server error response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable creates a DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable() *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate service unavailable response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate service unavailable response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate service unavailable response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge softwareupdate service unavailable response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edge softwareupdate service unavailable response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout creates a DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout() *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout{}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edge softwareupdate gateway timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edge softwareupdate gateway timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edge softwareupdate gateway timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edge softwareupdate gateway timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edge softwareupdate gateway timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate][%d] deleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgeSoftwareupdateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
