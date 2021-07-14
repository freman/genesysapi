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

// GetSpeechandtextanalyticsProgramReader is a Reader for the GetSpeechandtextanalyticsProgram structure.
type GetSpeechandtextanalyticsProgramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsProgramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsProgramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsProgramBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsProgramUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsProgramForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsProgramNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsProgramRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsProgramRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsProgramUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsProgramTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsProgramInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsProgramServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsProgramGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsProgramOK creates a GetSpeechandtextanalyticsProgramOK with default headers values
func NewGetSpeechandtextanalyticsProgramOK() *GetSpeechandtextanalyticsProgramOK {
	return &GetSpeechandtextanalyticsProgramOK{}
}

/*GetSpeechandtextanalyticsProgramOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsProgramOK struct {
	Payload *models.Program
}

func (o *GetSpeechandtextanalyticsProgramOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramOK) GetPayload() *models.Program {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Program)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramBadRequest creates a GetSpeechandtextanalyticsProgramBadRequest with default headers values
func NewGetSpeechandtextanalyticsProgramBadRequest() *GetSpeechandtextanalyticsProgramBadRequest {
	return &GetSpeechandtextanalyticsProgramBadRequest{}
}

/*GetSpeechandtextanalyticsProgramBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsProgramBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramUnauthorized creates a GetSpeechandtextanalyticsProgramUnauthorized with default headers values
func NewGetSpeechandtextanalyticsProgramUnauthorized() *GetSpeechandtextanalyticsProgramUnauthorized {
	return &GetSpeechandtextanalyticsProgramUnauthorized{}
}

/*GetSpeechandtextanalyticsProgramUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsProgramUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramForbidden creates a GetSpeechandtextanalyticsProgramForbidden with default headers values
func NewGetSpeechandtextanalyticsProgramForbidden() *GetSpeechandtextanalyticsProgramForbidden {
	return &GetSpeechandtextanalyticsProgramForbidden{}
}

/*GetSpeechandtextanalyticsProgramForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsProgramForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramNotFound creates a GetSpeechandtextanalyticsProgramNotFound with default headers values
func NewGetSpeechandtextanalyticsProgramNotFound() *GetSpeechandtextanalyticsProgramNotFound {
	return &GetSpeechandtextanalyticsProgramNotFound{}
}

/*GetSpeechandtextanalyticsProgramNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsProgramNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramRequestTimeout creates a GetSpeechandtextanalyticsProgramRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramRequestTimeout() *GetSpeechandtextanalyticsProgramRequestTimeout {
	return &GetSpeechandtextanalyticsProgramRequestTimeout{}
}

/*GetSpeechandtextanalyticsProgramRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsProgramRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramRequestEntityTooLarge creates a GetSpeechandtextanalyticsProgramRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsProgramRequestEntityTooLarge() *GetSpeechandtextanalyticsProgramRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsProgramRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsProgramRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsProgramRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramUnsupportedMediaType creates a GetSpeechandtextanalyticsProgramUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsProgramUnsupportedMediaType() *GetSpeechandtextanalyticsProgramUnsupportedMediaType {
	return &GetSpeechandtextanalyticsProgramUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsProgramUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsProgramUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramTooManyRequests creates a GetSpeechandtextanalyticsProgramTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsProgramTooManyRequests() *GetSpeechandtextanalyticsProgramTooManyRequests {
	return &GetSpeechandtextanalyticsProgramTooManyRequests{}
}

/*GetSpeechandtextanalyticsProgramTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsProgramTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramInternalServerError creates a GetSpeechandtextanalyticsProgramInternalServerError with default headers values
func NewGetSpeechandtextanalyticsProgramInternalServerError() *GetSpeechandtextanalyticsProgramInternalServerError {
	return &GetSpeechandtextanalyticsProgramInternalServerError{}
}

/*GetSpeechandtextanalyticsProgramInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsProgramInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramServiceUnavailable creates a GetSpeechandtextanalyticsProgramServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsProgramServiceUnavailable() *GetSpeechandtextanalyticsProgramServiceUnavailable {
	return &GetSpeechandtextanalyticsProgramServiceUnavailable{}
}

/*GetSpeechandtextanalyticsProgramServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsProgramServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramGatewayTimeout creates a GetSpeechandtextanalyticsProgramGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramGatewayTimeout() *GetSpeechandtextanalyticsProgramGatewayTimeout {
	return &GetSpeechandtextanalyticsProgramGatewayTimeout{}
}

/*GetSpeechandtextanalyticsProgramGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsProgramGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}][%d] getSpeechandtextanalyticsProgramGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
