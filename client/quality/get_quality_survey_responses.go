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

// GetQualitySurveyReader is a Reader for the GetQualitySurvey structure.
type GetQualitySurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualitySurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualitySurveyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualitySurveyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualitySurveyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualitySurveyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualitySurveyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualitySurveyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualitySurveyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualitySurveyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualitySurveyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualitySurveyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualitySurveyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualitySurveyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualitySurveyOK creates a GetQualitySurveyOK with default headers values
func NewGetQualitySurveyOK() *GetQualitySurveyOK {
	return &GetQualitySurveyOK{}
}

/*GetQualitySurveyOK handles this case with default header values.

successful operation
*/
type GetQualitySurveyOK struct {
	Payload *models.Survey
}

func (o *GetQualitySurveyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyOK  %+v", 200, o.Payload)
}

func (o *GetQualitySurveyOK) GetPayload() *models.Survey {
	return o.Payload
}

func (o *GetQualitySurveyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Survey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyBadRequest creates a GetQualitySurveyBadRequest with default headers values
func NewGetQualitySurveyBadRequest() *GetQualitySurveyBadRequest {
	return &GetQualitySurveyBadRequest{}
}

/*GetQualitySurveyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualitySurveyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualitySurveyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyUnauthorized creates a GetQualitySurveyUnauthorized with default headers values
func NewGetQualitySurveyUnauthorized() *GetQualitySurveyUnauthorized {
	return &GetQualitySurveyUnauthorized{}
}

/*GetQualitySurveyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualitySurveyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualitySurveyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyForbidden creates a GetQualitySurveyForbidden with default headers values
func NewGetQualitySurveyForbidden() *GetQualitySurveyForbidden {
	return &GetQualitySurveyForbidden{}
}

/*GetQualitySurveyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualitySurveyForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyForbidden  %+v", 403, o.Payload)
}

func (o *GetQualitySurveyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyNotFound creates a GetQualitySurveyNotFound with default headers values
func NewGetQualitySurveyNotFound() *GetQualitySurveyNotFound {
	return &GetQualitySurveyNotFound{}
}

/*GetQualitySurveyNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualitySurveyNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyNotFound  %+v", 404, o.Payload)
}

func (o *GetQualitySurveyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyRequestTimeout creates a GetQualitySurveyRequestTimeout with default headers values
func NewGetQualitySurveyRequestTimeout() *GetQualitySurveyRequestTimeout {
	return &GetQualitySurveyRequestTimeout{}
}

/*GetQualitySurveyRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualitySurveyRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualitySurveyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyRequestEntityTooLarge creates a GetQualitySurveyRequestEntityTooLarge with default headers values
func NewGetQualitySurveyRequestEntityTooLarge() *GetQualitySurveyRequestEntityTooLarge {
	return &GetQualitySurveyRequestEntityTooLarge{}
}

/*GetQualitySurveyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetQualitySurveyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualitySurveyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyUnsupportedMediaType creates a GetQualitySurveyUnsupportedMediaType with default headers values
func NewGetQualitySurveyUnsupportedMediaType() *GetQualitySurveyUnsupportedMediaType {
	return &GetQualitySurveyUnsupportedMediaType{}
}

/*GetQualitySurveyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualitySurveyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualitySurveyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyTooManyRequests creates a GetQualitySurveyTooManyRequests with default headers values
func NewGetQualitySurveyTooManyRequests() *GetQualitySurveyTooManyRequests {
	return &GetQualitySurveyTooManyRequests{}
}

/*GetQualitySurveyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualitySurveyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualitySurveyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyInternalServerError creates a GetQualitySurveyInternalServerError with default headers values
func NewGetQualitySurveyInternalServerError() *GetQualitySurveyInternalServerError {
	return &GetQualitySurveyInternalServerError{}
}

/*GetQualitySurveyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualitySurveyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualitySurveyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyServiceUnavailable creates a GetQualitySurveyServiceUnavailable with default headers values
func NewGetQualitySurveyServiceUnavailable() *GetQualitySurveyServiceUnavailable {
	return &GetQualitySurveyServiceUnavailable{}
}

/*GetQualitySurveyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualitySurveyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualitySurveyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualitySurveyGatewayTimeout creates a GetQualitySurveyGatewayTimeout with default headers values
func NewGetQualitySurveyGatewayTimeout() *GetQualitySurveyGatewayTimeout {
	return &GetQualitySurveyGatewayTimeout{}
}

/*GetQualitySurveyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualitySurveyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualitySurveyGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/surveys/{surveyId}][%d] getQualitySurveyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualitySurveyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualitySurveyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
