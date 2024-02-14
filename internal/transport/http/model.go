package transporthttp

// Request represents the parameters for a feasibility check request.
// It includes details such as the limit on the number of respondents and the expected length of the interview.
type Request struct {
	// Limit is the maximum number of participants for the survey.
	Limit int32 `json:"limit"`
	// LengthOfInterview is the estimated duration of the survey in minutes.
	LengthOfInterview int32 `json:"lengthOfInterview"`
}

// Response encapsulates the outcome of a feasibility check.
// It provides information on whether a survey is feasible and the expected delivery time if it is.
type Response struct {
	// Feasible indicates whether the survey can be conducted with the specified parameters.
	Feasible bool `json:"feasible"`
	// DeliveryTime estimates when the survey results can be delivered, given in a specific time format.
	DeliveryTime string `json:"deliveryTime"`
}
