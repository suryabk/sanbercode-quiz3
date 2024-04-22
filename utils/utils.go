package utils

import "regexp"

func ThicknessStatus(total_pages int) string {
	if total_pages <= 100 {
		return "tipis"
	} else if total_pages <= 200 {
		return "sedang"
	} else {
		return "tebal"
	}
}

func ValidateImageURL(imageURL string) bool {
	// Regular expression pattern for URL validation
	pattern := `^(http[s]?):\/\/?([^\s\/$.?#].[^\s]*)+$`
	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)
	// Check if the imageURL matches the regex pattern
	return regex.MatchString(imageURL)

}
