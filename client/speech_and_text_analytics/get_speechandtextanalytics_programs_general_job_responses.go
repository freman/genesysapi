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

// GetSpeechandtextanalyticsProgramsGeneralJobReader is a Reader for the GetSpeechandtextanalyticsProgramsGeneralJob structure.
type GetSpeechandtextanalyticsProgramsGeneralJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsProgramsGeneralJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobOK creates a GetSpeechandtextanalyticsProgramsGeneralJobOK with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobOK() *GetSpeechandtextanalyticsProgramsGeneralJobOK {
	return &GetSpeechandtextanalyticsProgramsGeneralJobOK{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsProgramsGeneralJobOK struct {
	Payload *models.GeneralProgramJob
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobOK) GetPayload() *models.GeneralProgramJob {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralProgramJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobBadRequest creates a GetSpeechandtextanalyticsProgramsGeneralJobBadRequest with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobBadRequest() *GetSpeechandtextanalyticsProgramsGeneralJobBadRequest {
	return &GetSpeechandtextanalyticsProgramsGeneralJobBadRequest{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobUnauthorized creates a GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobUnauthorized() *GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized {
	return &GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobForbidden creates a GetSpeechandtextanalyticsProgramsGeneralJobForbidden with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobForbidden() *GetSpeechandtextanalyticsProgramsGeneralJobForbidden {
	return &GetSpeechandtextanalyticsProgramsGeneralJobForbidden{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobNotFound creates a GetSpeechandtextanalyticsProgramsGeneralJobNotFound with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobNotFound() *GetSpeechandtextanalyticsProgramsGeneralJobNotFound {
	return &GetSpeechandtextanalyticsProgramsGeneralJobNotFound{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge creates a GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge() *GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType creates a GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType() *GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType {
	return &GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests creates a GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests() *GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests {
	return &GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobInternalServerError creates a GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobInternalServerError() *GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError {
	return &GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable creates a GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable() *GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable {
	return &GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout creates a GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout() *GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout {
	return &GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout{}
}

/*GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}][%d] getSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGeneralJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
