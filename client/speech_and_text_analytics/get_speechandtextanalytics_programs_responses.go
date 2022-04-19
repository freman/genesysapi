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

// GetSpeechandtextanalyticsProgramsReader is a Reader for the GetSpeechandtextanalyticsPrograms structure.
type GetSpeechandtextanalyticsProgramsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsProgramsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsProgramsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsProgramsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsProgramsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsProgramsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsProgramsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsProgramsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsProgramsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsProgramsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsProgramsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsProgramsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsProgramsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsProgramsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsProgramsOK creates a GetSpeechandtextanalyticsProgramsOK with default headers values
func NewGetSpeechandtextanalyticsProgramsOK() *GetSpeechandtextanalyticsProgramsOK {
	return &GetSpeechandtextanalyticsProgramsOK{}
}

/*GetSpeechandtextanalyticsProgramsOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsProgramsOK struct {
	Payload *models.ProgramsEntityListing
}

func (o *GetSpeechandtextanalyticsProgramsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsOK) GetPayload() *models.ProgramsEntityListing {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProgramsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsBadRequest creates a GetSpeechandtextanalyticsProgramsBadRequest with default headers values
func NewGetSpeechandtextanalyticsProgramsBadRequest() *GetSpeechandtextanalyticsProgramsBadRequest {
	return &GetSpeechandtextanalyticsProgramsBadRequest{}
}

/*GetSpeechandtextanalyticsProgramsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsProgramsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsUnauthorized creates a GetSpeechandtextanalyticsProgramsUnauthorized with default headers values
func NewGetSpeechandtextanalyticsProgramsUnauthorized() *GetSpeechandtextanalyticsProgramsUnauthorized {
	return &GetSpeechandtextanalyticsProgramsUnauthorized{}
}

/*GetSpeechandtextanalyticsProgramsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsProgramsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsForbidden creates a GetSpeechandtextanalyticsProgramsForbidden with default headers values
func NewGetSpeechandtextanalyticsProgramsForbidden() *GetSpeechandtextanalyticsProgramsForbidden {
	return &GetSpeechandtextanalyticsProgramsForbidden{}
}

/*GetSpeechandtextanalyticsProgramsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsProgramsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsNotFound creates a GetSpeechandtextanalyticsProgramsNotFound with default headers values
func NewGetSpeechandtextanalyticsProgramsNotFound() *GetSpeechandtextanalyticsProgramsNotFound {
	return &GetSpeechandtextanalyticsProgramsNotFound{}
}

/*GetSpeechandtextanalyticsProgramsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsProgramsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsRequestTimeout creates a GetSpeechandtextanalyticsProgramsRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramsRequestTimeout() *GetSpeechandtextanalyticsProgramsRequestTimeout {
	return &GetSpeechandtextanalyticsProgramsRequestTimeout{}
}

/*GetSpeechandtextanalyticsProgramsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsProgramsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsRequestEntityTooLarge creates a GetSpeechandtextanalyticsProgramsRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsProgramsRequestEntityTooLarge() *GetSpeechandtextanalyticsProgramsRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsProgramsRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsProgramsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetSpeechandtextanalyticsProgramsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsUnsupportedMediaType creates a GetSpeechandtextanalyticsProgramsUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsProgramsUnsupportedMediaType() *GetSpeechandtextanalyticsProgramsUnsupportedMediaType {
	return &GetSpeechandtextanalyticsProgramsUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsProgramsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsProgramsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsTooManyRequests creates a GetSpeechandtextanalyticsProgramsTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsProgramsTooManyRequests() *GetSpeechandtextanalyticsProgramsTooManyRequests {
	return &GetSpeechandtextanalyticsProgramsTooManyRequests{}
}

/*GetSpeechandtextanalyticsProgramsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsProgramsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsInternalServerError creates a GetSpeechandtextanalyticsProgramsInternalServerError with default headers values
func NewGetSpeechandtextanalyticsProgramsInternalServerError() *GetSpeechandtextanalyticsProgramsInternalServerError {
	return &GetSpeechandtextanalyticsProgramsInternalServerError{}
}

/*GetSpeechandtextanalyticsProgramsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsProgramsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsServiceUnavailable creates a GetSpeechandtextanalyticsProgramsServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsProgramsServiceUnavailable() *GetSpeechandtextanalyticsProgramsServiceUnavailable {
	return &GetSpeechandtextanalyticsProgramsServiceUnavailable{}
}

/*GetSpeechandtextanalyticsProgramsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsProgramsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramsGatewayTimeout creates a GetSpeechandtextanalyticsProgramsGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramsGatewayTimeout() *GetSpeechandtextanalyticsProgramsGatewayTimeout {
	return &GetSpeechandtextanalyticsProgramsGatewayTimeout{}
}

/*GetSpeechandtextanalyticsProgramsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsProgramsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsProgramsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs][%d] getSpeechandtextanalyticsProgramsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
