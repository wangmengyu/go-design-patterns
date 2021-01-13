package factory_method

import "fmt"

func GetGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong type")
}
