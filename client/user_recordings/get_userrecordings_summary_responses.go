// Code generated by go-swagger; DO NOT EDIT.

package user_recordings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserrecordingsSummaryReader is a Reader for the GetUserrecordingsSummary structure.
type GetUserrecordingsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserrecordingsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserrecordingsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserrecordingsSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserrecordingsSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserrecordingsSummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserrecordingsSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserrecordingsSummaryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserrecordingsSummaryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserrecordingsSummaryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserrecordingsSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserrecordingsSummaryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserrecordingsSummaryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserrecordingsSummaryOK creates a GetUserrecordingsSummaryOK with default headers values
func NewGetUserrecordingsSummaryOK() *GetUserrecordingsSummaryOK {
	return &GetUserrecordingsSummaryOK{}
}

/*GetUserrecordingsSummaryOK handles this case with default header values.

successful operation
*/
type GetUserrecordingsSummaryOK struct {
	Payload *models.FaxSummary
}

func (o *GetUserrecordingsSummaryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetUserrecordingsSummaryOK) GetPayload() *models.FaxSummary {
	return o.Payload
}

func (o *GetUserrecordingsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryBadRequest creates a GetUserrecordingsSummaryBadRequest with default headers values
func NewGetUserrecordingsSummaryBadRequest() *GetUserrecordingsSummaryBadRequest {
	return &GetUserrecordingsSummaryBadRequest{}
}

/*GetUserrecordingsSummaryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserrecordingsSummaryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserrecordingsSummaryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryUnauthorized creates a GetUserrecordingsSummaryUnauthorized with default headers values
func NewGetUserrecordingsSummaryUnauthorized() *GetUserrecordingsSummaryUnauthorized {
	return &GetUserrecordingsSummaryUnauthorized{}
}

/*GetUserrecordingsSummaryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserrecordingsSummaryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserrecordingsSummaryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryForbidden creates a GetUserrecordingsSummaryForbidden with default headers values
func NewGetUserrecordingsSummaryForbidden() *GetUserrecordingsSummaryForbidden {
	return &GetUserrecordingsSummaryForbidden{}
}

/*GetUserrecordingsSummaryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserrecordingsSummaryForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryForbidden  %+v", 403, o.Payload)
}

func (o *GetUserrecordingsSummaryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryNotFound creates a GetUserrecordingsSummaryNotFound with default headers values
func NewGetUserrecordingsSummaryNotFound() *GetUserrecordingsSummaryNotFound {
	return &GetUserrecordingsSummaryNotFound{}
}

/*GetUserrecordingsSummaryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserrecordingsSummaryNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryNotFound  %+v", 404, o.Payload)
}

func (o *GetUserrecordingsSummaryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryRequestEntityTooLarge creates a GetUserrecordingsSummaryRequestEntityTooLarge with default headers values
func NewGetUserrecordingsSummaryRequestEntityTooLarge() *GetUserrecordingsSummaryRequestEntityTooLarge {
	return &GetUserrecordingsSummaryRequestEntityTooLarge{}
}

/*GetUserrecordingsSummaryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserrecordingsSummaryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserrecordingsSummaryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryUnsupportedMediaType creates a GetUserrecordingsSummaryUnsupportedMediaType with default headers values
func NewGetUserrecordingsSummaryUnsupportedMediaType() *GetUserrecordingsSummaryUnsupportedMediaType {
	return &GetUserrecordingsSummaryUnsupportedMediaType{}
}

/*GetUserrecordingsSummaryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserrecordingsSummaryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserrecordingsSummaryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryTooManyRequests creates a GetUserrecordingsSummaryTooManyRequests with default headers values
func NewGetUserrecordingsSummaryTooManyRequests() *GetUserrecordingsSummaryTooManyRequests {
	return &GetUserrecordingsSummaryTooManyRequests{}
}

/*GetUserrecordingsSummaryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUserrecordingsSummaryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserrecordingsSummaryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryInternalServerError creates a GetUserrecordingsSummaryInternalServerError with default headers values
func NewGetUserrecordingsSummaryInternalServerError() *GetUserrecordingsSummaryInternalServerError {
	return &GetUserrecordingsSummaryInternalServerError{}
}

/*GetUserrecordingsSummaryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserrecordingsSummaryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserrecordingsSummaryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryServiceUnavailable creates a GetUserrecordingsSummaryServiceUnavailable with default headers values
func NewGetUserrecordingsSummaryServiceUnavailable() *GetUserrecordingsSummaryServiceUnavailable {
	return &GetUserrecordingsSummaryServiceUnavailable{}
}

/*GetUserrecordingsSummaryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserrecordingsSummaryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserrecordingsSummaryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsSummaryGatewayTimeout creates a GetUserrecordingsSummaryGatewayTimeout with default headers values
func NewGetUserrecordingsSummaryGatewayTimeout() *GetUserrecordingsSummaryGatewayTimeout {
	return &GetUserrecordingsSummaryGatewayTimeout{}
}

/*GetUserrecordingsSummaryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserrecordingsSummaryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingsSummaryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/summary][%d] getUserrecordingsSummaryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserrecordingsSummaryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsSummaryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
