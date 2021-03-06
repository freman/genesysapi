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

// GetUserrecordingMediaReader is a Reader for the GetUserrecordingMedia structure.
type GetUserrecordingMediaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserrecordingMediaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserrecordingMediaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserrecordingMediaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserrecordingMediaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserrecordingMediaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserrecordingMediaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserrecordingMediaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserrecordingMediaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserrecordingMediaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserrecordingMediaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserrecordingMediaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserrecordingMediaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserrecordingMediaOK creates a GetUserrecordingMediaOK with default headers values
func NewGetUserrecordingMediaOK() *GetUserrecordingMediaOK {
	return &GetUserrecordingMediaOK{}
}

/*GetUserrecordingMediaOK handles this case with default header values.

successful operation
*/
type GetUserrecordingMediaOK struct {
	Payload *models.DownloadResponse
}

func (o *GetUserrecordingMediaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaOK  %+v", 200, o.Payload)
}

func (o *GetUserrecordingMediaOK) GetPayload() *models.DownloadResponse {
	return o.Payload
}

func (o *GetUserrecordingMediaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DownloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaBadRequest creates a GetUserrecordingMediaBadRequest with default headers values
func NewGetUserrecordingMediaBadRequest() *GetUserrecordingMediaBadRequest {
	return &GetUserrecordingMediaBadRequest{}
}

/*GetUserrecordingMediaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserrecordingMediaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserrecordingMediaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaUnauthorized creates a GetUserrecordingMediaUnauthorized with default headers values
func NewGetUserrecordingMediaUnauthorized() *GetUserrecordingMediaUnauthorized {
	return &GetUserrecordingMediaUnauthorized{}
}

/*GetUserrecordingMediaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserrecordingMediaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserrecordingMediaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaForbidden creates a GetUserrecordingMediaForbidden with default headers values
func NewGetUserrecordingMediaForbidden() *GetUserrecordingMediaForbidden {
	return &GetUserrecordingMediaForbidden{}
}

/*GetUserrecordingMediaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserrecordingMediaForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaForbidden  %+v", 403, o.Payload)
}

func (o *GetUserrecordingMediaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaNotFound creates a GetUserrecordingMediaNotFound with default headers values
func NewGetUserrecordingMediaNotFound() *GetUserrecordingMediaNotFound {
	return &GetUserrecordingMediaNotFound{}
}

/*GetUserrecordingMediaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserrecordingMediaNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaNotFound  %+v", 404, o.Payload)
}

func (o *GetUserrecordingMediaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaRequestEntityTooLarge creates a GetUserrecordingMediaRequestEntityTooLarge with default headers values
func NewGetUserrecordingMediaRequestEntityTooLarge() *GetUserrecordingMediaRequestEntityTooLarge {
	return &GetUserrecordingMediaRequestEntityTooLarge{}
}

/*GetUserrecordingMediaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserrecordingMediaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserrecordingMediaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaUnsupportedMediaType creates a GetUserrecordingMediaUnsupportedMediaType with default headers values
func NewGetUserrecordingMediaUnsupportedMediaType() *GetUserrecordingMediaUnsupportedMediaType {
	return &GetUserrecordingMediaUnsupportedMediaType{}
}

/*GetUserrecordingMediaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserrecordingMediaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserrecordingMediaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaTooManyRequests creates a GetUserrecordingMediaTooManyRequests with default headers values
func NewGetUserrecordingMediaTooManyRequests() *GetUserrecordingMediaTooManyRequests {
	return &GetUserrecordingMediaTooManyRequests{}
}

/*GetUserrecordingMediaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUserrecordingMediaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserrecordingMediaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaInternalServerError creates a GetUserrecordingMediaInternalServerError with default headers values
func NewGetUserrecordingMediaInternalServerError() *GetUserrecordingMediaInternalServerError {
	return &GetUserrecordingMediaInternalServerError{}
}

/*GetUserrecordingMediaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserrecordingMediaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserrecordingMediaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaServiceUnavailable creates a GetUserrecordingMediaServiceUnavailable with default headers values
func NewGetUserrecordingMediaServiceUnavailable() *GetUserrecordingMediaServiceUnavailable {
	return &GetUserrecordingMediaServiceUnavailable{}
}

/*GetUserrecordingMediaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserrecordingMediaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserrecordingMediaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingMediaGatewayTimeout creates a GetUserrecordingMediaGatewayTimeout with default headers values
func NewGetUserrecordingMediaGatewayTimeout() *GetUserrecordingMediaGatewayTimeout {
	return &GetUserrecordingMediaGatewayTimeout{}
}

/*GetUserrecordingMediaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserrecordingMediaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserrecordingMediaGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings/{recordingId}/media][%d] getUserrecordingMediaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserrecordingMediaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingMediaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
