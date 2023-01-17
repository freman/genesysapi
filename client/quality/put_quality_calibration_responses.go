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

// PutQualityCalibrationReader is a Reader for the PutQualityCalibration structure.
type PutQualityCalibrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutQualityCalibrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutQualityCalibrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutQualityCalibrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutQualityCalibrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutQualityCalibrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutQualityCalibrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutQualityCalibrationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutQualityCalibrationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutQualityCalibrationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutQualityCalibrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutQualityCalibrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutQualityCalibrationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutQualityCalibrationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutQualityCalibrationOK creates a PutQualityCalibrationOK with default headers values
func NewPutQualityCalibrationOK() *PutQualityCalibrationOK {
	return &PutQualityCalibrationOK{}
}

/*
PutQualityCalibrationOK describes a response with status code 200, with default header values.

successful operation
*/
type PutQualityCalibrationOK struct {
	Payload *models.Calibration
}

// IsSuccess returns true when this put quality calibration o k response has a 2xx status code
func (o *PutQualityCalibrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put quality calibration o k response has a 3xx status code
func (o *PutQualityCalibrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration o k response has a 4xx status code
func (o *PutQualityCalibrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put quality calibration o k response has a 5xx status code
func (o *PutQualityCalibrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration o k response a status code equal to that given
func (o *PutQualityCalibrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutQualityCalibrationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationOK  %+v", 200, o.Payload)
}

func (o *PutQualityCalibrationOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationOK  %+v", 200, o.Payload)
}

func (o *PutQualityCalibrationOK) GetPayload() *models.Calibration {
	return o.Payload
}

func (o *PutQualityCalibrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Calibration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationBadRequest creates a PutQualityCalibrationBadRequest with default headers values
func NewPutQualityCalibrationBadRequest() *PutQualityCalibrationBadRequest {
	return &PutQualityCalibrationBadRequest{}
}

/*
PutQualityCalibrationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutQualityCalibrationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration bad request response has a 2xx status code
func (o *PutQualityCalibrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration bad request response has a 3xx status code
func (o *PutQualityCalibrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration bad request response has a 4xx status code
func (o *PutQualityCalibrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration bad request response has a 5xx status code
func (o *PutQualityCalibrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration bad request response a status code equal to that given
func (o *PutQualityCalibrationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutQualityCalibrationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationBadRequest  %+v", 400, o.Payload)
}

func (o *PutQualityCalibrationBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationBadRequest  %+v", 400, o.Payload)
}

func (o *PutQualityCalibrationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationUnauthorized creates a PutQualityCalibrationUnauthorized with default headers values
func NewPutQualityCalibrationUnauthorized() *PutQualityCalibrationUnauthorized {
	return &PutQualityCalibrationUnauthorized{}
}

/*
PutQualityCalibrationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutQualityCalibrationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration unauthorized response has a 2xx status code
func (o *PutQualityCalibrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration unauthorized response has a 3xx status code
func (o *PutQualityCalibrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration unauthorized response has a 4xx status code
func (o *PutQualityCalibrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration unauthorized response has a 5xx status code
func (o *PutQualityCalibrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration unauthorized response a status code equal to that given
func (o *PutQualityCalibrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutQualityCalibrationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutQualityCalibrationUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutQualityCalibrationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationForbidden creates a PutQualityCalibrationForbidden with default headers values
func NewPutQualityCalibrationForbidden() *PutQualityCalibrationForbidden {
	return &PutQualityCalibrationForbidden{}
}

/*
PutQualityCalibrationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutQualityCalibrationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration forbidden response has a 2xx status code
func (o *PutQualityCalibrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration forbidden response has a 3xx status code
func (o *PutQualityCalibrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration forbidden response has a 4xx status code
func (o *PutQualityCalibrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration forbidden response has a 5xx status code
func (o *PutQualityCalibrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration forbidden response a status code equal to that given
func (o *PutQualityCalibrationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutQualityCalibrationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationForbidden  %+v", 403, o.Payload)
}

func (o *PutQualityCalibrationForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationForbidden  %+v", 403, o.Payload)
}

func (o *PutQualityCalibrationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationNotFound creates a PutQualityCalibrationNotFound with default headers values
func NewPutQualityCalibrationNotFound() *PutQualityCalibrationNotFound {
	return &PutQualityCalibrationNotFound{}
}

/*
PutQualityCalibrationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutQualityCalibrationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration not found response has a 2xx status code
func (o *PutQualityCalibrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration not found response has a 3xx status code
func (o *PutQualityCalibrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration not found response has a 4xx status code
func (o *PutQualityCalibrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration not found response has a 5xx status code
func (o *PutQualityCalibrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration not found response a status code equal to that given
func (o *PutQualityCalibrationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutQualityCalibrationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationNotFound  %+v", 404, o.Payload)
}

func (o *PutQualityCalibrationNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationNotFound  %+v", 404, o.Payload)
}

func (o *PutQualityCalibrationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationRequestTimeout creates a PutQualityCalibrationRequestTimeout with default headers values
func NewPutQualityCalibrationRequestTimeout() *PutQualityCalibrationRequestTimeout {
	return &PutQualityCalibrationRequestTimeout{}
}

/*
PutQualityCalibrationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutQualityCalibrationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration request timeout response has a 2xx status code
func (o *PutQualityCalibrationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration request timeout response has a 3xx status code
func (o *PutQualityCalibrationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration request timeout response has a 4xx status code
func (o *PutQualityCalibrationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration request timeout response has a 5xx status code
func (o *PutQualityCalibrationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration request timeout response a status code equal to that given
func (o *PutQualityCalibrationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutQualityCalibrationRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutQualityCalibrationRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutQualityCalibrationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationRequestEntityTooLarge creates a PutQualityCalibrationRequestEntityTooLarge with default headers values
func NewPutQualityCalibrationRequestEntityTooLarge() *PutQualityCalibrationRequestEntityTooLarge {
	return &PutQualityCalibrationRequestEntityTooLarge{}
}

/*
PutQualityCalibrationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutQualityCalibrationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration request entity too large response has a 2xx status code
func (o *PutQualityCalibrationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration request entity too large response has a 3xx status code
func (o *PutQualityCalibrationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration request entity too large response has a 4xx status code
func (o *PutQualityCalibrationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration request entity too large response has a 5xx status code
func (o *PutQualityCalibrationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration request entity too large response a status code equal to that given
func (o *PutQualityCalibrationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutQualityCalibrationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutQualityCalibrationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutQualityCalibrationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationUnsupportedMediaType creates a PutQualityCalibrationUnsupportedMediaType with default headers values
func NewPutQualityCalibrationUnsupportedMediaType() *PutQualityCalibrationUnsupportedMediaType {
	return &PutQualityCalibrationUnsupportedMediaType{}
}

/*
PutQualityCalibrationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutQualityCalibrationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration unsupported media type response has a 2xx status code
func (o *PutQualityCalibrationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration unsupported media type response has a 3xx status code
func (o *PutQualityCalibrationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration unsupported media type response has a 4xx status code
func (o *PutQualityCalibrationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration unsupported media type response has a 5xx status code
func (o *PutQualityCalibrationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration unsupported media type response a status code equal to that given
func (o *PutQualityCalibrationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutQualityCalibrationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutQualityCalibrationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutQualityCalibrationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationTooManyRequests creates a PutQualityCalibrationTooManyRequests with default headers values
func NewPutQualityCalibrationTooManyRequests() *PutQualityCalibrationTooManyRequests {
	return &PutQualityCalibrationTooManyRequests{}
}

/*
PutQualityCalibrationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutQualityCalibrationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration too many requests response has a 2xx status code
func (o *PutQualityCalibrationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration too many requests response has a 3xx status code
func (o *PutQualityCalibrationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration too many requests response has a 4xx status code
func (o *PutQualityCalibrationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put quality calibration too many requests response has a 5xx status code
func (o *PutQualityCalibrationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put quality calibration too many requests response a status code equal to that given
func (o *PutQualityCalibrationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutQualityCalibrationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutQualityCalibrationTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutQualityCalibrationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationInternalServerError creates a PutQualityCalibrationInternalServerError with default headers values
func NewPutQualityCalibrationInternalServerError() *PutQualityCalibrationInternalServerError {
	return &PutQualityCalibrationInternalServerError{}
}

/*
PutQualityCalibrationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutQualityCalibrationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration internal server error response has a 2xx status code
func (o *PutQualityCalibrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration internal server error response has a 3xx status code
func (o *PutQualityCalibrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration internal server error response has a 4xx status code
func (o *PutQualityCalibrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put quality calibration internal server error response has a 5xx status code
func (o *PutQualityCalibrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put quality calibration internal server error response a status code equal to that given
func (o *PutQualityCalibrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutQualityCalibrationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutQualityCalibrationInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutQualityCalibrationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationServiceUnavailable creates a PutQualityCalibrationServiceUnavailable with default headers values
func NewPutQualityCalibrationServiceUnavailable() *PutQualityCalibrationServiceUnavailable {
	return &PutQualityCalibrationServiceUnavailable{}
}

/*
PutQualityCalibrationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutQualityCalibrationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration service unavailable response has a 2xx status code
func (o *PutQualityCalibrationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration service unavailable response has a 3xx status code
func (o *PutQualityCalibrationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration service unavailable response has a 4xx status code
func (o *PutQualityCalibrationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put quality calibration service unavailable response has a 5xx status code
func (o *PutQualityCalibrationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put quality calibration service unavailable response a status code equal to that given
func (o *PutQualityCalibrationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutQualityCalibrationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutQualityCalibrationServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutQualityCalibrationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityCalibrationGatewayTimeout creates a PutQualityCalibrationGatewayTimeout with default headers values
func NewPutQualityCalibrationGatewayTimeout() *PutQualityCalibrationGatewayTimeout {
	return &PutQualityCalibrationGatewayTimeout{}
}

/*
PutQualityCalibrationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutQualityCalibrationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put quality calibration gateway timeout response has a 2xx status code
func (o *PutQualityCalibrationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put quality calibration gateway timeout response has a 3xx status code
func (o *PutQualityCalibrationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put quality calibration gateway timeout response has a 4xx status code
func (o *PutQualityCalibrationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put quality calibration gateway timeout response has a 5xx status code
func (o *PutQualityCalibrationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put quality calibration gateway timeout response a status code equal to that given
func (o *PutQualityCalibrationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutQualityCalibrationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutQualityCalibrationGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/quality/calibrations/{calibrationId}][%d] putQualityCalibrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutQualityCalibrationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityCalibrationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
