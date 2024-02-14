# Delivery Estimate Time
Delivery Time Estimator Service

The Delivery Time Estimator Service is a backend service designed to estimate the delivery time for a survey project based on the total number of respondents required (limit) and the length of the interview (lengthOfInterview). The service interacts with the Cint Feasibility API to determine the feasibility of the project and to calculate the shortest feasible delivery time.

Features:

Estimation of Delivery Time: The service provides an endpoint that accepts parameters such as limit and lengthOfInterview and returns an estimation of the delivery time for the survey project.

Feasibility Check: The service evaluates the feasibility of the project by making calls to the Cint Feasibility API. A project is considered feasible if the feasibility score returned by the Cint API is greater than 0.75.

Shortest Feasible Delivery Time: The service determines the shortest feasible delivery time by considering the fieldPeriod parameter provided by the Cint API, which represents the delivery time in days.

## Getting started

### Run

`export CINT_API_KEY=your_api_key`
 `go mod tidy`
 `go mod vendor`
 `go run cmd/service/app.go`

### curl example on default setting

`curl --location 'localhost:8080/estimate/deliverytime' \
--header 'Content-Type: application/json' \
--data '{

    "limit" : 100,
    "lengthOfInterview" : 2
}'
`
