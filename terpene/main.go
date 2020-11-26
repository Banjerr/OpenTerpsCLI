package terpene

import "openterpscli/utilities"

func GetTerpenes(BaseURL string) string {
	data := utilities.GetData(BaseURL, "terpenes")
	return data
}

func GetTerpene() string {
	return ""
}

func AddTerpene() string {
	return ""
}

func UpdateTerpene() string {
	return ""
}

func RemoveTerpene() string {
	return ""
}
