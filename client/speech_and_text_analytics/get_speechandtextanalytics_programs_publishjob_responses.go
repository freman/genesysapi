// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetSpeechandtextanalyticsProgramsPublishjobReader is a Reader for the GetSpeechandtextanalyticsProgramsPublishjob structure.
type GetSpeechandtextanalyticsProgramsPublishjobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsProgramsPublishjobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsProgramsPublishjobOK creates a GetSpeechandtextanalyticsProgramsPublishjobOK with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobOK() *GetSpeechandtextanalyticsProgramsPublishjobOK {
	return &GetSpeechandtextanalyticsProgramsPublishjobOK{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsProgramsPublishjobOK struct {
	Payload *models.ProgramJob
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobOK) GetPayload() *models.ProgramJob {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProgramJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobBadRequest creates a GetSpeechandtextanalyticsProgramsPublishjobBadRequest with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobBadRequest() *GetSpeechandtextanalyticsProgramsPublishjobBadRequest {
	return &GetSpeechandtextanalyticsProgramsPublishjobBadRequest{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsProgramsPublishjobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobUnauthorized creates a GetSpeechandtextanalyticsProgramsPublishjobUnauthorized with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobUnauthorized() *GetSpeechandtextanalyticsProgramsPublishjobUnauthorized {
	return &GetSpeechandtextanalyticsProgramsPublishjobUnauthorized{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsProgramsPublishjobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobForbidden creates a GetSpeechandtextanalyticsProgramsPublishjobForbidden with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobForbidden() *GetSpeechandtextanalyticsProgramsPublishjobForbidden {
	return &GetSpeechandtextanalyticsProgramsPublishjobForbidden{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsProgramsPublishjobForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobNotFound creates a GetSpeechandtextanalyticsProgramsPublishjobNotFound with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobNotFound() *GetSpeechandtextanalyticsProgramsPublishjobNotFound {
	return &GetSpeechandtextanalyticsProgramsPublishjobNotFound{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsProgramsPublishjobNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobRequestTimeout creates a GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobRequestTimeout() *GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout {
	return &GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge creates a GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge() *GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType creates a GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType() *GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType {
	return &GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobTooManyRequests creates a GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobTooManyRequests() *GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests {
	return &GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobInternalServerError creates a GetSpeechandtextanalyticsProgramsPublishjobInternalServerError with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobInternalServerError() *GetSpeechandtextanalyticsProgramsPublishjobInternalServerError {
	return &GetSpeechandtextanalyticsProgramsPublishjobInternalServerError{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsProgramsPublishjobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable creates a GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable() *GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable {
	return &GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout creates a GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout() *GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout {
	return &GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout{}
}

/*GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}][%d] getSpeechandtextanalyticsProgramsPublishjobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsPublishjobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
