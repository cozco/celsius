package celsius

//ConvertToCelsius takes a fathrenheit temperature and coverts it to celsius.
func ConvertToCelsius(fahrenheit *float64) float64 {
	//Logic for Celsius conversion
	celsius := (*fahrenheit - 32) * 5 / 9
	return celsius
}
