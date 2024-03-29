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

// DeleteTelephonyProvidersEdgesTrunkbasesettingReader is a Reader for the DeleteTelephonyProvidersEdgesTrunkbasesetting structure.
type DeleteTelephonyProvidersEdgesTrunkbasesettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingOK creates a DeleteTelephonyProvidersEdgesTrunkbasesettingOK with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingOK() *DeleteTelephonyProvidersEdgesTrunkbasesettingOK {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingOK{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingOK struct {
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting o k response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting o k response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting o k response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting o k response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting o k response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest creates a DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest() *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting bad request response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting bad request response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting bad request response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting bad request response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting bad request response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized creates a DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized() *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting unauthorized response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting unauthorized response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting unauthorized response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting unauthorized response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting unauthorized response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingForbidden creates a DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingForbidden() *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting forbidden response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting forbidden response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting forbidden response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting forbidden response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting forbidden response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingNotFound creates a DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingNotFound() *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting not found response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting not found response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting not found response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting not found response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting not found response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout creates a DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout() *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting request timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting request timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting request timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting request timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting request timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge() *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting request entity too large response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting request entity too large response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting request entity too large response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting request entity too large response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting request entity too large response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType creates a DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType() *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting unsupported media type response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting unsupported media type response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting unsupported media type response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting unsupported media type response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting unsupported media type response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests creates a DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests() *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting too many requests response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting too many requests response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting too many requests response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting too many requests response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting too many requests response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError creates a DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError() *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting internal server error response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting internal server error response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting internal server error response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting internal server error response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting internal server error response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable creates a DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable() *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting service unavailable response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting service unavailable response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting service unavailable response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting service unavailable response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting service unavailable response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout creates a DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout() *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout {
	return &DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout{}
}

/*
DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete telephony providers edges trunkbasesetting gateway timeout response has a 2xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete telephony providers edges trunkbasesetting gateway timeout response has a 3xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete telephony providers edges trunkbasesetting gateway timeout response has a 4xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete telephony providers edges trunkbasesetting gateway timeout response has a 5xx status code
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete telephony providers edges trunkbasesetting gateway timeout response a status code equal to that given
func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}][%d] deleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesTrunkbasesettingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
