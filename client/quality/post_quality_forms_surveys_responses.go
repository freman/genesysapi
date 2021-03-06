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

// PostQualityFormsSurveysReader is a Reader for the PostQualityFormsSurveys structure.
type PostQualityFormsSurveysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityFormsSurveysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityFormsSurveysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityFormsSurveysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityFormsSurveysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityFormsSurveysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityFormsSurveysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityFormsSurveysRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityFormsSurveysUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityFormsSurveysTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityFormsSurveysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityFormsSurveysServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityFormsSurveysGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityFormsSurveysOK creates a PostQualityFormsSurveysOK with default headers values
func NewPostQualityFormsSurveysOK() *PostQualityFormsSurveysOK {
	return &PostQualityFormsSurveysOK{}
}

/*PostQualityFormsSurveysOK handles this case with default header values.

successful operation
*/
type PostQualityFormsSurveysOK struct {
	Payload *models.SurveyForm
}

func (o *PostQualityFormsSurveysOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsSurveysOK) GetPayload() *models.SurveyForm {
	return o.Payload
}

func (o *PostQualityFormsSurveysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SurveyForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysBadRequest creates a PostQualityFormsSurveysBadRequest with default headers values
func NewPostQualityFormsSurveysBadRequest() *PostQualityFormsSurveysBadRequest {
	return &PostQualityFormsSurveysBadRequest{}
}

/*PostQualityFormsSurveysBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityFormsSurveysBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsSurveysBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysUnauthorized creates a PostQualityFormsSurveysUnauthorized with default headers values
func NewPostQualityFormsSurveysUnauthorized() *PostQualityFormsSurveysUnauthorized {
	return &PostQualityFormsSurveysUnauthorized{}
}

/*PostQualityFormsSurveysUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityFormsSurveysUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsSurveysUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysForbidden creates a PostQualityFormsSurveysForbidden with default headers values
func NewPostQualityFormsSurveysForbidden() *PostQualityFormsSurveysForbidden {
	return &PostQualityFormsSurveysForbidden{}
}

/*PostQualityFormsSurveysForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityFormsSurveysForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsSurveysForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysNotFound creates a PostQualityFormsSurveysNotFound with default headers values
func NewPostQualityFormsSurveysNotFound() *PostQualityFormsSurveysNotFound {
	return &PostQualityFormsSurveysNotFound{}
}

/*PostQualityFormsSurveysNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualityFormsSurveysNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsSurveysNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysRequestEntityTooLarge creates a PostQualityFormsSurveysRequestEntityTooLarge with default headers values
func NewPostQualityFormsSurveysRequestEntityTooLarge() *PostQualityFormsSurveysRequestEntityTooLarge {
	return &PostQualityFormsSurveysRequestEntityTooLarge{}
}

/*PostQualityFormsSurveysRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostQualityFormsSurveysRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysUnsupportedMediaType creates a PostQualityFormsSurveysUnsupportedMediaType with default headers values
func NewPostQualityFormsSurveysUnsupportedMediaType() *PostQualityFormsSurveysUnsupportedMediaType {
	return &PostQualityFormsSurveysUnsupportedMediaType{}
}

/*PostQualityFormsSurveysUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityFormsSurveysUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysTooManyRequests creates a PostQualityFormsSurveysTooManyRequests with default headers values
func NewPostQualityFormsSurveysTooManyRequests() *PostQualityFormsSurveysTooManyRequests {
	return &PostQualityFormsSurveysTooManyRequests{}
}

/*PostQualityFormsSurveysTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostQualityFormsSurveysTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsSurveysTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysInternalServerError creates a PostQualityFormsSurveysInternalServerError with default headers values
func NewPostQualityFormsSurveysInternalServerError() *PostQualityFormsSurveysInternalServerError {
	return &PostQualityFormsSurveysInternalServerError{}
}

/*PostQualityFormsSurveysInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityFormsSurveysInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsSurveysInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysServiceUnavailable creates a PostQualityFormsSurveysServiceUnavailable with default headers values
func NewPostQualityFormsSurveysServiceUnavailable() *PostQualityFormsSurveysServiceUnavailable {
	return &PostQualityFormsSurveysServiceUnavailable{}
}

/*PostQualityFormsSurveysServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityFormsSurveysServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsSurveysServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsSurveysGatewayTimeout creates a PostQualityFormsSurveysGatewayTimeout with default headers values
func NewPostQualityFormsSurveysGatewayTimeout() *PostQualityFormsSurveysGatewayTimeout {
	return &PostQualityFormsSurveysGatewayTimeout{}
}

/*PostQualityFormsSurveysGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualityFormsSurveysGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsSurveysGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/surveys][%d] postQualityFormsSurveysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsSurveysGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsSurveysGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
