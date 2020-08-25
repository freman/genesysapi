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

// PostQualitySurveysScoringReader is a Reader for the PostQualitySurveysScoring structure.
type PostQualitySurveysScoringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualitySurveysScoringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualitySurveysScoringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualitySurveysScoringBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualitySurveysScoringUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualitySurveysScoringForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualitySurveysScoringNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualitySurveysScoringRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualitySurveysScoringUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualitySurveysScoringTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualitySurveysScoringInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualitySurveysScoringServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualitySurveysScoringGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualitySurveysScoringOK creates a PostQualitySurveysScoringOK with default headers values
func NewPostQualitySurveysScoringOK() *PostQualitySurveysScoringOK {
	return &PostQualitySurveysScoringOK{}
}

/*PostQualitySurveysScoringOK handles this case with default header values.

successful operation
*/
type PostQualitySurveysScoringOK struct {
	Payload *models.SurveyScoringSet
}

func (o *PostQualitySurveysScoringOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringOK  %+v", 200, o.Payload)
}

func (o *PostQualitySurveysScoringOK) GetPayload() *models.SurveyScoringSet {
	return o.Payload
}

func (o *PostQualitySurveysScoringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SurveyScoringSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringBadRequest creates a PostQualitySurveysScoringBadRequest with default headers values
func NewPostQualitySurveysScoringBadRequest() *PostQualitySurveysScoringBadRequest {
	return &PostQualitySurveysScoringBadRequest{}
}

/*PostQualitySurveysScoringBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualitySurveysScoringBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualitySurveysScoringBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringUnauthorized creates a PostQualitySurveysScoringUnauthorized with default headers values
func NewPostQualitySurveysScoringUnauthorized() *PostQualitySurveysScoringUnauthorized {
	return &PostQualitySurveysScoringUnauthorized{}
}

/*PostQualitySurveysScoringUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualitySurveysScoringUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualitySurveysScoringUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringForbidden creates a PostQualitySurveysScoringForbidden with default headers values
func NewPostQualitySurveysScoringForbidden() *PostQualitySurveysScoringForbidden {
	return &PostQualitySurveysScoringForbidden{}
}

/*PostQualitySurveysScoringForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualitySurveysScoringForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringForbidden  %+v", 403, o.Payload)
}

func (o *PostQualitySurveysScoringForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringNotFound creates a PostQualitySurveysScoringNotFound with default headers values
func NewPostQualitySurveysScoringNotFound() *PostQualitySurveysScoringNotFound {
	return &PostQualitySurveysScoringNotFound{}
}

/*PostQualitySurveysScoringNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualitySurveysScoringNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringNotFound  %+v", 404, o.Payload)
}

func (o *PostQualitySurveysScoringNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringRequestEntityTooLarge creates a PostQualitySurveysScoringRequestEntityTooLarge with default headers values
func NewPostQualitySurveysScoringRequestEntityTooLarge() *PostQualitySurveysScoringRequestEntityTooLarge {
	return &PostQualitySurveysScoringRequestEntityTooLarge{}
}

/*PostQualitySurveysScoringRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostQualitySurveysScoringRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualitySurveysScoringRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringUnsupportedMediaType creates a PostQualitySurveysScoringUnsupportedMediaType with default headers values
func NewPostQualitySurveysScoringUnsupportedMediaType() *PostQualitySurveysScoringUnsupportedMediaType {
	return &PostQualitySurveysScoringUnsupportedMediaType{}
}

/*PostQualitySurveysScoringUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualitySurveysScoringUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualitySurveysScoringUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringTooManyRequests creates a PostQualitySurveysScoringTooManyRequests with default headers values
func NewPostQualitySurveysScoringTooManyRequests() *PostQualitySurveysScoringTooManyRequests {
	return &PostQualitySurveysScoringTooManyRequests{}
}

/*PostQualitySurveysScoringTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostQualitySurveysScoringTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualitySurveysScoringTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringInternalServerError creates a PostQualitySurveysScoringInternalServerError with default headers values
func NewPostQualitySurveysScoringInternalServerError() *PostQualitySurveysScoringInternalServerError {
	return &PostQualitySurveysScoringInternalServerError{}
}

/*PostQualitySurveysScoringInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualitySurveysScoringInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualitySurveysScoringInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringServiceUnavailable creates a PostQualitySurveysScoringServiceUnavailable with default headers values
func NewPostQualitySurveysScoringServiceUnavailable() *PostQualitySurveysScoringServiceUnavailable {
	return &PostQualitySurveysScoringServiceUnavailable{}
}

/*PostQualitySurveysScoringServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualitySurveysScoringServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualitySurveysScoringServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySurveysScoringGatewayTimeout creates a PostQualitySurveysScoringGatewayTimeout with default headers values
func NewPostQualitySurveysScoringGatewayTimeout() *PostQualitySurveysScoringGatewayTimeout {
	return &PostQualitySurveysScoringGatewayTimeout{}
}

/*PostQualitySurveysScoringGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualitySurveysScoringGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySurveysScoringGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/surveys/scoring][%d] postQualitySurveysScoringGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualitySurveysScoringGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySurveysScoringGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}