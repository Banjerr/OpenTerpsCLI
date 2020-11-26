package strain

import (
	"openterpscli/utilities"
)

func GetStrains(BaseURL string) string {
	data := utilities.GetData(BaseURL, "strains")
	return data
}

func GetStrain() string {
	return ""
}

func AddStrain(BaseURL string) string {

	return ""
}

func UpdateStrain() string {
	return ""
}

func RemoveStrain() string {
	return ""
}
