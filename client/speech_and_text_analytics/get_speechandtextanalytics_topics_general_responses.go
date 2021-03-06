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

// GetSpeechandtextanalyticsTopicsGeneralReader is a Reader for the GetSpeechandtextanalyticsTopicsGeneral structure.
type GetSpeechandtextanalyticsTopicsGeneralReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsTopicsGeneralReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsTopicsGeneralOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsTopicsGeneralBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsTopicsGeneralUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsTopicsGeneralForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsTopicsGeneralNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsTopicsGeneralTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsTopicsGeneralInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsTopicsGeneralServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsTopicsGeneralGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsTopicsGeneralOK creates a GetSpeechandtextanalyticsTopicsGeneralOK with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralOK() *GetSpeechandtextanalyticsTopicsGeneralOK {
	return &GetSpeechandtextanalyticsTopicsGeneralOK{}
}

/*GetSpeechandtextanalyticsTopicsGeneralOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsTopicsGeneralOK struct {
	Payload *models.GeneralTopicsEntityListing
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) GetPayload() *models.GeneralTopicsEntityListing {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralTopicsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralBadRequest creates a GetSpeechandtextanalyticsTopicsGeneralBadRequest with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralBadRequest() *GetSpeechandtextanalyticsTopicsGeneralBadRequest {
	return &GetSpeechandtextanalyticsTopicsGeneralBadRequest{}
}

/*GetSpeechandtextanalyticsTopicsGeneralBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsTopicsGeneralBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralUnauthorized creates a GetSpeechandtextanalyticsTopicsGeneralUnauthorized with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralUnauthorized() *GetSpeechandtextanalyticsTopicsGeneralUnauthorized {
	return &GetSpeechandtextanalyticsTopicsGeneralUnauthorized{}
}

/*GetSpeechandtextanalyticsTopicsGeneralUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsTopicsGeneralUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralForbidden creates a GetSpeechandtextanalyticsTopicsGeneralForbidden with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralForbidden() *GetSpeechandtextanalyticsTopicsGeneralForbidden {
	return &GetSpeechandtextanalyticsTopicsGeneralForbidden{}
}

/*GetSpeechandtextanalyticsTopicsGeneralForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsTopicsGeneralForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralNotFound creates a GetSpeechandtextanalyticsTopicsGeneralNotFound with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralNotFound() *GetSpeechandtextanalyticsTopicsGeneralNotFound {
	return &GetSpeechandtextanalyticsTopicsGeneralNotFound{}
}

/*GetSpeechandtextanalyticsTopicsGeneralNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsTopicsGeneralNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge creates a GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge() *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType creates a GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType() *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType {
	return &GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralTooManyRequests creates a GetSpeechandtextanalyticsTopicsGeneralTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralTooManyRequests() *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests {
	return &GetSpeechandtextanalyticsTopicsGeneralTooManyRequests{}
}

/*GetSpeechandtextanalyticsTopicsGeneralTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetSpeechandtextanalyticsTopicsGeneralTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralInternalServerError creates a GetSpeechandtextanalyticsTopicsGeneralInternalServerError with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralInternalServerError() *GetSpeechandtextanalyticsTopicsGeneralInternalServerError {
	return &GetSpeechandtextanalyticsTopicsGeneralInternalServerError{}
}

/*GetSpeechandtextanalyticsTopicsGeneralInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsTopicsGeneralInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralServiceUnavailable creates a GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralServiceUnavailable() *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable {
	return &GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable{}
}

/*GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralGatewayTimeout creates a GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralGatewayTimeout() *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout {
	return &GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout{}
}

/*GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
