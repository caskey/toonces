package units

func ConvertGtoLb(gs float64) float64 {
	return gs * FromGstoLbs
}

func ConvertLbtoG(lbs float64) float64 {
	return lbs * FromLbstoGs
}
