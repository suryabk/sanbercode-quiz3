package utils

func ThicknessStatus(total_pages int) string {

	if total_pages <= 100 {
		return "tipis"
	} else if total_pages <= 200 {
		return "sedang"
	} else {
		return "tebal"
	}

}
