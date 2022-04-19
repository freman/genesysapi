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

// PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateReader is a Reader for the PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate structure.
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK struct {
	Payload *models.AsyncForecastOperationResult
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK) GetPayload() *models.AsyncForecastOperationResult {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncForecastOperationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated handles this case with default header values.

The forecast was successfully generated
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated struct {
	Payload *models.AsyncForecastOperationResult
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated) GetPayload() *models.AsyncForecastOperationResult {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncForecastOperationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted handles this case with default header values.

The request was accepted and the result will be sent asynchronously via notification
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted struct {
	Payload *models.AsyncForecastOperationResult
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted  %+v", 202, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted) GetPayload() *models.AsyncForecastOperationResult {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncForecastOperationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway handles this case with default header values.

Bad Gateway
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway  %+v", 502, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout creates a PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate][%d] postWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
