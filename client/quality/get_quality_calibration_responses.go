// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityCalibrationReader is a Reader for the GetQualityCalibration structure.
type GetQualityCalibrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityCalibrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityCalibrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityCalibrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityCalibrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityCalibrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityCalibrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityCalibrationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityCalibrationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityCalibrationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityCalibrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityCalibrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityCalibrationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityCalibrationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityCalibrationOK creates a GetQualityCalibrationOK with default headers values
func NewGetQualityCalibrationOK() *GetQualityCalibrationOK {
	return &GetQualityCalibrationOK{}
}

/*
GetQualityCalibrationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityCalibrationOK struct {
	Payload *models.Calibration
}

// IsSuccess returns true when this get quality calibration o k response has a 2xx status code
func (o *GetQualityCalibrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality calibration o k response has a 3xx status code
func (o *GetQualityCalibrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration o k response has a 4xx status code
func (o *GetQualityCalibrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality calibration o k response has a 5xx status code
func (o *GetQualityCalibrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration o k response a status code equal to that given
func (o *GetQualityCalibrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityCalibrationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationOK  %+v", 200, o.Payload)
}

func (o *GetQualityCalibrationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationOK  %+v", 200, o.Payload)
}

func (o *GetQualityCalibrationOK) GetPayload() *models.Calibration {
	return o.Payload
}

func (o *GetQualityCalibrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Calibration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationBadRequest creates a GetQualityCalibrationBadRequest with default headers values
func NewGetQualityCalibrationBadRequest() *GetQualityCalibrationBadRequest {
	return &GetQualityCalibrationBadRequest{}
}

/*
GetQualityCalibrationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityCalibrationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration bad request response has a 2xx status code
func (o *GetQualityCalibrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration bad request response has a 3xx status code
func (o *GetQualityCalibrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration bad request response has a 4xx status code
func (o *GetQualityCalibrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration bad request response has a 5xx status code
func (o *GetQualityCalibrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration bad request response a status code equal to that given
func (o *GetQualityCalibrationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityCalibrationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityCalibrationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityCalibrationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationUnauthorized creates a GetQualityCalibrationUnauthorized with default headers values
func NewGetQualityCalibrationUnauthorized() *GetQualityCalibrationUnauthorized {
	return &GetQualityCalibrationUnauthorized{}
}

/*
GetQualityCalibrationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityCalibrationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration unauthorized response has a 2xx status code
func (o *GetQualityCalibrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration unauthorized response has a 3xx status code
func (o *GetQualityCalibrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration unauthorized response has a 4xx status code
func (o *GetQualityCalibrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration unauthorized response has a 5xx status code
func (o *GetQualityCalibrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration unauthorized response a status code equal to that given
func (o *GetQualityCalibrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityCalibrationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityCalibrationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityCalibrationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationForbidden creates a GetQualityCalibrationForbidden with default headers values
func NewGetQualityCalibrationForbidden() *GetQualityCalibrationForbidden {
	return &GetQualityCalibrationForbidden{}
}

/*
GetQualityCalibrationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityCalibrationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration forbidden response has a 2xx status code
func (o *GetQualityCalibrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration forbidden response has a 3xx status code
func (o *GetQualityCalibrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration forbidden response has a 4xx status code
func (o *GetQualityCalibrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration forbidden response has a 5xx status code
func (o *GetQualityCalibrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration forbidden response a status code equal to that given
func (o *GetQualityCalibrationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityCalibrationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityCalibrationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityCalibrationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationNotFound creates a GetQualityCalibrationNotFound with default headers values
func NewGetQualityCalibrationNotFound() *GetQualityCalibrationNotFound {
	return &GetQualityCalibrationNotFound{}
}

/*
GetQualityCalibrationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityCalibrationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration not found response has a 2xx status code
func (o *GetQualityCalibrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration not found response has a 3xx status code
func (o *GetQualityCalibrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration not found response has a 4xx status code
func (o *GetQualityCalibrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration not found response has a 5xx status code
func (o *GetQualityCalibrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration not found response a status code equal to that given
func (o *GetQualityCalibrationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityCalibrationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityCalibrationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityCalibrationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationRequestTimeout creates a GetQualityCalibrationRequestTimeout with default headers values
func NewGetQualityCalibrationRequestTimeout() *GetQualityCalibrationRequestTimeout {
	return &GetQualityCalibrationRequestTimeout{}
}

/*
GetQualityCalibrationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityCalibrationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration request timeout response has a 2xx status code
func (o *GetQualityCalibrationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration request timeout response has a 3xx status code
func (o *GetQualityCalibrationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration request timeout response has a 4xx status code
func (o *GetQualityCalibrationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration request timeout response has a 5xx status code
func (o *GetQualityCalibrationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration request timeout response a status code equal to that given
func (o *GetQualityCalibrationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityCalibrationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityCalibrationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityCalibrationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationRequestEntityTooLarge creates a GetQualityCalibrationRequestEntityTooLarge with default headers values
func NewGetQualityCalibrationRequestEntityTooLarge() *GetQualityCalibrationRequestEntityTooLarge {
	return &GetQualityCalibrationRequestEntityTooLarge{}
}

/*
GetQualityCalibrationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetQualityCalibrationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration request entity too large response has a 2xx status code
func (o *GetQualityCalibrationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration request entity too large response has a 3xx status code
func (o *GetQualityCalibrationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration request entity too large response has a 4xx status code
func (o *GetQualityCalibrationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration request entity too large response has a 5xx status code
func (o *GetQualityCalibrationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration request entity too large response a status code equal to that given
func (o *GetQualityCalibrationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityCalibrationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityCalibrationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityCalibrationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationUnsupportedMediaType creates a GetQualityCalibrationUnsupportedMediaType with default headers values
func NewGetQualityCalibrationUnsupportedMediaType() *GetQualityCalibrationUnsupportedMediaType {
	return &GetQualityCalibrationUnsupportedMediaType{}
}

/*
GetQualityCalibrationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityCalibrationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration unsupported media type response has a 2xx status code
func (o *GetQualityCalibrationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration unsupported media type response has a 3xx status code
func (o *GetQualityCalibrationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration unsupported media type response has a 4xx status code
func (o *GetQualityCalibrationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration unsupported media type response has a 5xx status code
func (o *GetQualityCalibrationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration unsupported media type response a status code equal to that given
func (o *GetQualityCalibrationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityCalibrationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityCalibrationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityCalibrationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationTooManyRequests creates a GetQualityCalibrationTooManyRequests with default headers values
func NewGetQualityCalibrationTooManyRequests() *GetQualityCalibrationTooManyRequests {
	return &GetQualityCalibrationTooManyRequests{}
}

/*
GetQualityCalibrationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityCalibrationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration too many requests response has a 2xx status code
func (o *GetQualityCalibrationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration too many requests response has a 3xx status code
func (o *GetQualityCalibrationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration too many requests response has a 4xx status code
func (o *GetQualityCalibrationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality calibration too many requests response has a 5xx status code
func (o *GetQualityCalibrationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality calibration too many requests response a status code equal to that given
func (o *GetQualityCalibrationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityCalibrationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityCalibrationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityCalibrationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationInternalServerError creates a GetQualityCalibrationInternalServerError with default headers values
func NewGetQualityCalibrationInternalServerError() *GetQualityCalibrationInternalServerError {
	return &GetQualityCalibrationInternalServerError{}
}

/*
GetQualityCalibrationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityCalibrationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration internal server error response has a 2xx status code
func (o *GetQualityCalibrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration internal server error response has a 3xx status code
func (o *GetQualityCalibrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration internal server error response has a 4xx status code
func (o *GetQualityCalibrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality calibration internal server error response has a 5xx status code
func (o *GetQualityCalibrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality calibration internal server error response a status code equal to that given
func (o *GetQualityCalibrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityCalibrationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityCalibrationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityCalibrationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationServiceUnavailable creates a GetQualityCalibrationServiceUnavailable with default headers values
func NewGetQualityCalibrationServiceUnavailable() *GetQualityCalibrationServiceUnavailable {
	return &GetQualityCalibrationServiceUnavailable{}
}

/*
GetQualityCalibrationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityCalibrationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration service unavailable response has a 2xx status code
func (o *GetQualityCalibrationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration service unavailable response has a 3xx status code
func (o *GetQualityCalibrationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration service unavailable response has a 4xx status code
func (o *GetQualityCalibrationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality calibration service unavailable response has a 5xx status code
func (o *GetQualityCalibrationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality calibration service unavailable response a status code equal to that given
func (o *GetQualityCalibrationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityCalibrationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityCalibrationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityCalibrationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityCalibrationGatewayTimeout creates a GetQualityCalibrationGatewayTimeout with default headers values
func NewGetQualityCalibrationGatewayTimeout() *GetQualityCalibrationGatewayTimeout {
	return &GetQualityCalibrationGatewayTimeout{}
}

/*
GetQualityCalibrationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityCalibrationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality calibration gateway timeout response has a 2xx status code
func (o *GetQualityCalibrationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality calibration gateway timeout response has a 3xx status code
func (o *GetQualityCalibrationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality calibration gateway timeout response has a 4xx status code
func (o *GetQualityCalibrationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality calibration gateway timeout response has a 5xx status code
func (o *GetQualityCalibrationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality calibration gateway timeout response a status code equal to that given
func (o *GetQualityCalibrationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityCalibrationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityCalibrationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/calibrations/{calibrationId}][%d] getQualityCalibrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityCalibrationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityCalibrationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
