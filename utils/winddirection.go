package utils

// GetWindDirection ...
func GetWindDirection(deg int) string {

	switch {
	case 11 < deg && deg <= 33:
		return "северо северо восточный"
	case 33 < deg && deg <= 56:
		return "северо восточный"
	case 56 < deg && deg <= 76:
		return "восточно северо восточный"
	case 76 < deg && deg <= 101:
		return "Восточный"
	case 101 < deg && deg <= 123:
		return "восточно юго восточный"
	case 123 < deg && deg <= 146:
		return "юго восточный"
	case 146 < deg && deg <= 168:
		return "юго юго восточный"
	case 168 < deg && deg <= 191:
		return "южный"
	case 191 < deg && deg <= 213:
		return "юго юго западный"
	case 213 < deg && deg <= 236:
		return "югозападный"
	case 236 < deg && deg <= 258:
		return "западно юго западный"
	case 258 < deg && deg <= 281:
		return "западный"
	case 281 < deg && deg <= 303:
		return "западно северо западный"
	case 303 < deg && deg <= 326:
		return "северо западный"
	case 326 < deg && deg <= 348:
		return "северо сверо западный"
	case 348 < deg || deg <= 11:
		return "северный"
	}

	return ""
}
