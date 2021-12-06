package models

type Stand struct {
	StandID		string  `json:"standId"`
	Name       string  `json:"stand"`
	Power       string  `json:"power"`
	Speed       string  `json:"speed"`
	Range       string  `json:"range"`
	Persistence string  `json:"persistence"`
	Precision   string  `json:"precision"`
	Development string  `json:"development"`
}

type StandsInformation struct {
	Stands []Stand
}


func NewStand(id string, name string, power string, speed string, standRange string, persistence string, precision string, development string) *Stand{
	stand := new(Stand)
	stand.StandID = id
	stand.Name = name
	stand.Power = power
	stand.Speed = speed
	stand.Range = standRange
	stand.Persistence = persistence
	stand.Precision = precision
	stand.Development = development

	return stand
}

func (standsList *StandsInformation) AddStand(stand Stand) []Stand {
	standsList.Stands = append(standsList.Stands, stand)

	return standsList.Stands
}

func (stand Stand) GetId() string {
	return stand.StandID
}

func (stand *Stand) SetId(id string) {
	stand.StandID = id
}

func (stand Stand) GetName() string {
	return stand.Name
}

func (stand *Stand) SetName(name string) {
	stand.Name = name
}

func (stand Stand) GetPower() string {
	return stand.Power
}

func (stand *Stand) SetPower(power string) {
	stand.Power = power
}

func (stand Stand) GetSpeed() string {
	return stand.Speed
}

func (stand *Stand) SetSpeed(speed string) {
	stand.Speed = speed
}

func (stand Stand) GetRange() string {
	return stand.Range
}

func (stand *Stand) SetRange(standRange string) {
	stand.Range = standRange
}

func (stand Stand) GetPersistence() string {
	return stand.Persistence
}

func (stand *Stand) SetPersistence(persistence string) {
	stand.Persistence = persistence
}

func (stand Stand) GetPrecision() string {
	return stand.Precision
}

func (stand *Stand) SetPrecision(precision string) {
	stand.Precision = precision
}

func (stand Stand) GetDevelopment() string {
	return stand.Development
}

func (stand *Stand) SetDevelopment(development string) {
	stand.Development = development
}