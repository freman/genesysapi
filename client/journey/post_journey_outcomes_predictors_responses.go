// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostJourneyOutcomesPredictorsReader is a Reader for the PostJourneyOutcomesPredictors structure.
type PostJourneyOutcomesPredictorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostJourneyOutcomesPredictorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostJourneyOutcomesPredictorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostJourneyOutcomesPredictorsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostJourneyOutcomesPredictorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostJourneyOutcomesPredictorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostJourneyOutcomesPredictorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostJourneyOutcomesPredictorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostJourneyOutcomesPredictorsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostJourneyOutcomesPredictorsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostJourneyOutcomesPredictorsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostJourneyOutcomesPredictorsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostJourneyOutcomesPredictorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostJourneyOutcomesPredictorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostJourneyOutcomesPredictorsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostJourneyOutcomesPredictorsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostJourneyOutcomesPredictorsOK creates a PostJourneyOutcomesPredictorsOK with default headers values
func NewPostJourneyOutcomesPredictorsOK() *PostJourneyOutcomesPredictorsOK {
	return &PostJourneyOutcomesPredictorsOK{}
}

/*
PostJourneyOutcomesPredictorsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostJourneyOutcomesPredictorsOK struct {
	Payload *models.OutcomePredictor
}

// IsSuccess returns true when this post journey outcomes predictors o k response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post journey outcomes predictors o k response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors o k response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post journey outcomes predictors o k response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors o k response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostJourneyOutcomesPredictorsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsOK  %+v", 200, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsOK  %+v", 200, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsOK) GetPayload() *models.OutcomePredictor {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutcomePredictor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsCreated creates a PostJourneyOutcomesPredictorsCreated with default headers values
func NewPostJourneyOutcomesPredictorsCreated() *PostJourneyOutcomesPredictorsCreated {
	return &PostJourneyOutcomesPredictorsCreated{}
}

/*
PostJourneyOutcomesPredictorsCreated describes a response with status code 201, with default header values.

Predictor created.
*/
type PostJourneyOutcomesPredictorsCreated struct {
	Payload *models.OutcomePredictor
}

// IsSuccess returns true when this post journey outcomes predictors created response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post journey outcomes predictors created response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors created response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post journey outcomes predictors created response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors created response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostJourneyOutcomesPredictorsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsCreated  %+v", 201, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsCreated  %+v", 201, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsCreated) GetPayload() *models.OutcomePredictor {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutcomePredictor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsBadRequest creates a PostJourneyOutcomesPredictorsBadRequest with default headers values
func NewPostJourneyOutcomesPredictorsBadRequest() *PostJourneyOutcomesPredictorsBadRequest {
	return &PostJourneyOutcomesPredictorsBadRequest{}
}

/*
PostJourneyOutcomesPredictorsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostJourneyOutcomesPredictorsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors bad request response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors bad request response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors bad request response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors bad request response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors bad request response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostJourneyOutcomesPredictorsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsBadRequest  %+v", 400, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsBadRequest  %+v", 400, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsUnauthorized creates a PostJourneyOutcomesPredictorsUnauthorized with default headers values
func NewPostJourneyOutcomesPredictorsUnauthorized() *PostJourneyOutcomesPredictorsUnauthorized {
	return &PostJourneyOutcomesPredictorsUnauthorized{}
}

/*
PostJourneyOutcomesPredictorsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostJourneyOutcomesPredictorsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors unauthorized response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors unauthorized response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors unauthorized response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors unauthorized response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors unauthorized response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostJourneyOutcomesPredictorsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsForbidden creates a PostJourneyOutcomesPredictorsForbidden with default headers values
func NewPostJourneyOutcomesPredictorsForbidden() *PostJourneyOutcomesPredictorsForbidden {
	return &PostJourneyOutcomesPredictorsForbidden{}
}

/*
PostJourneyOutcomesPredictorsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostJourneyOutcomesPredictorsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors forbidden response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors forbidden response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors forbidden response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors forbidden response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors forbidden response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostJourneyOutcomesPredictorsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsForbidden  %+v", 403, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsForbidden  %+v", 403, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsNotFound creates a PostJourneyOutcomesPredictorsNotFound with default headers values
func NewPostJourneyOutcomesPredictorsNotFound() *PostJourneyOutcomesPredictorsNotFound {
	return &PostJourneyOutcomesPredictorsNotFound{}
}

/*
PostJourneyOutcomesPredictorsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostJourneyOutcomesPredictorsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors not found response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors not found response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors not found response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors not found response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors not found response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostJourneyOutcomesPredictorsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsNotFound  %+v", 404, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsNotFound  %+v", 404, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsRequestTimeout creates a PostJourneyOutcomesPredictorsRequestTimeout with default headers values
func NewPostJourneyOutcomesPredictorsRequestTimeout() *PostJourneyOutcomesPredictorsRequestTimeout {
	return &PostJourneyOutcomesPredictorsRequestTimeout{}
}

/*
PostJourneyOutcomesPredictorsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostJourneyOutcomesPredictorsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors request timeout response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors request timeout response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors request timeout response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors request timeout response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors request timeout response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostJourneyOutcomesPredictorsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsConflict creates a PostJourneyOutcomesPredictorsConflict with default headers values
func NewPostJourneyOutcomesPredictorsConflict() *PostJourneyOutcomesPredictorsConflict {
	return &PostJourneyOutcomesPredictorsConflict{}
}

/*
PostJourneyOutcomesPredictorsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostJourneyOutcomesPredictorsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors conflict response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors conflict response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors conflict response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors conflict response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors conflict response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostJourneyOutcomesPredictorsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsConflict  %+v", 409, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsConflict  %+v", 409, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsRequestEntityTooLarge creates a PostJourneyOutcomesPredictorsRequestEntityTooLarge with default headers values
func NewPostJourneyOutcomesPredictorsRequestEntityTooLarge() *PostJourneyOutcomesPredictorsRequestEntityTooLarge {
	return &PostJourneyOutcomesPredictorsRequestEntityTooLarge{}
}

/*
PostJourneyOutcomesPredictorsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostJourneyOutcomesPredictorsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors request entity too large response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors request entity too large response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors request entity too large response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors request entity too large response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors request entity too large response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsUnsupportedMediaType creates a PostJourneyOutcomesPredictorsUnsupportedMediaType with default headers values
func NewPostJourneyOutcomesPredictorsUnsupportedMediaType() *PostJourneyOutcomesPredictorsUnsupportedMediaType {
	return &PostJourneyOutcomesPredictorsUnsupportedMediaType{}
}

/*
PostJourneyOutcomesPredictorsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostJourneyOutcomesPredictorsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors unsupported media type response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors unsupported media type response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors unsupported media type response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors unsupported media type response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors unsupported media type response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsTooManyRequests creates a PostJourneyOutcomesPredictorsTooManyRequests with default headers values
func NewPostJourneyOutcomesPredictorsTooManyRequests() *PostJourneyOutcomesPredictorsTooManyRequests {
	return &PostJourneyOutcomesPredictorsTooManyRequests{}
}

/*
PostJourneyOutcomesPredictorsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostJourneyOutcomesPredictorsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors too many requests response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors too many requests response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors too many requests response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post journey outcomes predictors too many requests response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post journey outcomes predictors too many requests response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostJourneyOutcomesPredictorsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsInternalServerError creates a PostJourneyOutcomesPredictorsInternalServerError with default headers values
func NewPostJourneyOutcomesPredictorsInternalServerError() *PostJourneyOutcomesPredictorsInternalServerError {
	return &PostJourneyOutcomesPredictorsInternalServerError{}
}

/*
PostJourneyOutcomesPredictorsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostJourneyOutcomesPredictorsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors internal server error response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors internal server error response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors internal server error response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post journey outcomes predictors internal server error response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post journey outcomes predictors internal server error response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostJourneyOutcomesPredictorsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsServiceUnavailable creates a PostJourneyOutcomesPredictorsServiceUnavailable with default headers values
func NewPostJourneyOutcomesPredictorsServiceUnavailable() *PostJourneyOutcomesPredictorsServiceUnavailable {
	return &PostJourneyOutcomesPredictorsServiceUnavailable{}
}

/*
PostJourneyOutcomesPredictorsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostJourneyOutcomesPredictorsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors service unavailable response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors service unavailable response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors service unavailable response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post journey outcomes predictors service unavailable response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post journey outcomes predictors service unavailable response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostJourneyOutcomesPredictorsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesPredictorsGatewayTimeout creates a PostJourneyOutcomesPredictorsGatewayTimeout with default headers values
func NewPostJourneyOutcomesPredictorsGatewayTimeout() *PostJourneyOutcomesPredictorsGatewayTimeout {
	return &PostJourneyOutcomesPredictorsGatewayTimeout{}
}

/*
PostJourneyOutcomesPredictorsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostJourneyOutcomesPredictorsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post journey outcomes predictors gateway timeout response has a 2xx status code
func (o *PostJourneyOutcomesPredictorsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post journey outcomes predictors gateway timeout response has a 3xx status code
func (o *PostJourneyOutcomesPredictorsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post journey outcomes predictors gateway timeout response has a 4xx status code
func (o *PostJourneyOutcomesPredictorsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post journey outcomes predictors gateway timeout response has a 5xx status code
func (o *PostJourneyOutcomesPredictorsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post journey outcomes predictors gateway timeout response a status code equal to that given
func (o *PostJourneyOutcomesPredictorsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostJourneyOutcomesPredictorsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes/predictors][%d] postJourneyOutcomesPredictorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostJourneyOutcomesPredictorsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesPredictorsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
