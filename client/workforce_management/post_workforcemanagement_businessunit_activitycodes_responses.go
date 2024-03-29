// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementBusinessunitActivitycodesReader is a Reader for the PostWorkforcemanagementBusinessunitActivitycodes structure.
type PostWorkforcemanagementBusinessunitActivitycodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitActivitycodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitActivitycodesOK creates a PostWorkforcemanagementBusinessunitActivitycodesOK with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesOK() *PostWorkforcemanagementBusinessunitActivitycodesOK {
	return &PostWorkforcemanagementBusinessunitActivitycodesOK{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitActivitycodesOK struct {
	Payload *models.BusinessUnitActivityCode
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes o k response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes o k response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes o k response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes o k response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes o k response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) GetPayload() *models.BusinessUnitActivityCode {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessUnitActivityCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesBadRequest creates a PostWorkforcemanagementBusinessunitActivitycodesBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesBadRequest() *PostWorkforcemanagementBusinessunitActivitycodesBadRequest {
	return &PostWorkforcemanagementBusinessunitActivitycodesBadRequest{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitActivitycodesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes bad request response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes bad request response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes bad request response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes bad request response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes bad request response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesUnauthorized creates a PostWorkforcemanagementBusinessunitActivitycodesUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesUnauthorized() *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized {
	return &PostWorkforcemanagementBusinessunitActivitycodesUnauthorized{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitActivitycodesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesForbidden creates a PostWorkforcemanagementBusinessunitActivitycodesForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesForbidden() *PostWorkforcemanagementBusinessunitActivitycodesForbidden {
	return &PostWorkforcemanagementBusinessunitActivitycodesForbidden{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitActivitycodesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes forbidden response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes forbidden response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes forbidden response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes forbidden response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes forbidden response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesNotFound creates a PostWorkforcemanagementBusinessunitActivitycodesNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesNotFound() *PostWorkforcemanagementBusinessunitActivitycodesNotFound {
	return &PostWorkforcemanagementBusinessunitActivitycodesNotFound{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitActivitycodesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes not found response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes not found response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes not found response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes not found response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes not found response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesRequestTimeout creates a PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesRequestTimeout() *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout {
	return &PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes request timeout response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes request timeout response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes request timeout response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes request timeout response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes request timeout response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType() *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesTooManyRequests creates a PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesTooManyRequests() *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests {
	return &PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes too many requests response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes too many requests response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes too many requests response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes too many requests response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes too many requests response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesInternalServerError creates a PostWorkforcemanagementBusinessunitActivitycodesInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesInternalServerError() *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError {
	return &PostWorkforcemanagementBusinessunitActivitycodesInternalServerError{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitActivitycodesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes internal server error response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes internal server error response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes internal server error response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes internal server error response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes internal server error response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable creates a PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable() *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout creates a PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout() *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout{}
}

/*
PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement businessunit activitycodes gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement businessunit activitycodes gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement businessunit activitycodes gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement businessunit activitycodes gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement businessunit activitycodes gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes][%d] postWorkforcemanagementBusinessunitActivitycodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitActivitycodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
