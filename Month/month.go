package Month_Package

type Month struct {
	Name     string          `json:"name"`
	AllTemps [32][24]float64 `json:"AllTemps"`
	AvgTemp  float64         `json:"AvgTemp"`
	Days     int             `json:"Days"`
}

func (m *Month) TimeStamp(temp float64, day int, hour int) {
	m.AllTemps[day][hour] = temp
}
func (m *Month) CalcMonthlyAvg() {
	sum := 0.0
	var amount int
	for i := 0; i < len(m.AllTemps); i++ {
		for e := 0; e < len(m.AllTemps[i]); e++ {
			if m.AllTemps[i][e] == 0 {
				break
			}
			sum += m.AllTemps[i][e]
			amount++
		}
	}
	m.AvgTemp = sum / float64(amount)
}
func (m *Month) CalcDays() {
	for day := 0; day < len(m.AllTemps); day++ {
		dayHasTemperature := false
		for hour := 0; hour < len(m.AllTemps[day]); hour++ {
			if m.AllTemps[day][hour] != 0 {
				dayHasTemperature = true
				break
			}
		}
		if dayHasTemperature {
			m.Days++
		}
	}
}

func NewMonth(name string) *Month {
	return &Month{
		Name: name,
	}
}
