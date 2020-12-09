package genfunc

// GetGenTableNameTemp get gen table template str
func GetGenTableNameTemp() string {
	return genTnf
}

// GetGenColumnNameTemp get gen column template str
func GetGenColumnNameTemp() string {
	return genColumn
}

// GetGenBaseTemp get gen base template str
func GetGenBaseTemp() string {
	return genBase
}

// GetGenLogicTemp get gen logic template str
func GetGenLogicTemp() string {
	return genlogic
}

// GetGenPreloadTemp get gen preload template str
func GetGenPreloadTemp(multi bool) string {
	if multi {
		return genPreloadMulti
	}
	return genPreload
}
