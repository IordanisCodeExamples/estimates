package service

import (
	"context"
	"fmt"

	client "github.com/junkd0g/go-client-cintworks"
)

// GetEstimates calculates the feasibility of completing a survey project within a predefined set of time periods.
// It checks the feasibility for each period defined in FieldPeriods and determines whether the project can be completed
// within those timeframes based on a set of default and provided parameters.
//
// Parameters:
// - ctx: Context for the request, used for cancellation and deadlines.
// - limit: The maximum number of respondents for the survey.
// - lengthOfInterview: The estimated duration of the survey in minutes.
//
// Returns:
// - isFeasible: A boolean indicating whether the project is feasible within the checked time periods.
// - message: A string message providing details on the feasibility and the estimated completion time if feasible.
// - error: An error object that will be non-nil if an error occurs during the feasibility check process.
//
// The method iterates through a list of predefined field periods (FieldPeriods) and queries the CintAPI for each
// to get feasibility estimates. If the feasibility score for a period is less than 0.75, it is considered feasible,
// and the method returns indicating the project can be completed within that period. The method returns early
// if a feasible period is found. If no feasible period is found, it indicates that the project cannot be completed
// within 1 month by default.
func (s *Service) GetEstimates(ctx context.Context, limit, lengthOfInterview int32) (bool, string, error) {
	isFeasible := false
	message := "The project cannot be completed within 1 month."

	periodMapping := map[int32]string{
		1:  OneDayString,
		7:  OneWeekString,
		30: OneMonthString,
	}

	for _, period := range FieldPeriods {
		resp, err := s.CintAPI.GetFeasibilityEstimates(
			client.SurveySettingsRequest{
				Limit:             limit,
				LengthOfInterview: lengthOfInterview,
				FieldPeriod:       period,
				IncidenceRate:     DefaultIncidenceRate,
				CountryId:         DefaultCountryID,
				QuotaGroups: []client.QuotaGroupRequest{
					{
						Quotas: []client.QuotaRequest{
							{
								Limit: limit,
								TargetGroup: client.TargetGroupRequest{
									MaxAge: DefaultMaxAge,
									MinAge: DefaultMinAge,
								},
							},
						},
						Limit: limit,
					},
				},
			},
		)
		if err != nil {
			return false, "", fmt.Errorf("failed to get feasibility estimates: %w", err)
		}
		fmt.Println("-----------------------------------")
		fmt.Println(resp)
		fmt.Println("-----------------------------------")

		if resp.Feasibility < 0.75 {
			isFeasible = true
			// Update the message with the appropriate delivery time based on the period
			if _, exists := periodMapping[period]; exists {
				message = periodMapping[period]
			}
			break // Exit the loop if a feasible period is found
		}
	}

	return isFeasible, message, nil
}
