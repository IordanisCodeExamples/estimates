package service

const (
	DefaultIncidenceRate int32  = 100       // DefaultIncidenceRate represents the default incidence rate percentage.
	DefaultCountryID     int32  = 1         // DefaultCountryID represents the default country identifier.
	DefaultMaxAge        int32  = 99        // DefaultMaxAge represents the default maximum age for a target group.
	DefaultMinAge        int32  = 18        // DefaultMinAge represents the default minimum age for a target group.
	OneDayString         string = "1 day"   // OneDayString represents the string representation of 1 day.
	OneWeekString        string = "1 week"  // OneWeekString represents the string representation of 1 week.
	OneMonthString       string = "1 month" // OneMonthString represents the string representation of 1 month.
)

var (
	FieldPeriods [3]int32 = [3]int32{1, 7, 30}
)
