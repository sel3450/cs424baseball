package baseballplayer

type BaseballPlayer struct {
	FirstName       string
	LastName        string
	PlateAppearance int
	AtBats          int
	Singles         int
	Doubles         int
	Triples         int
	HomeRuns        int
	Walks           int
	HitByPitch      int
}

func (p BaseballPlayer) GetFirstName() string {
	return p.FirstName
}

func (p *BaseballPlayer) SetFirstName(firstname string) {
	p.FirstName = firstname
}

func (p BaseballPlayer) GetLastName() string {
	return p.LastName
}

func (p *BaseballPlayer) SetLastName(lastname string) {
	p.LastName = lastname
}

func (p BaseballPlayer) GetPlateAppearance() int {
	return p.PlateAppearance
}

func (p *BaseballPlayer) SetPlateAppearance(plateappearance int) {
	p.PlateAppearance = plateappearance
}

func (p BaseballPlayer) GetAtBats() int {
	return p.AtBats
}

func (p *BaseballPlayer) SetAtBats(atbats int) {
	p.AtBats = atbats
}

func (p BaseballPlayer) GetSingles() int {
	return p.Singles
}

func (p *BaseballPlayer) SetSingles(singles int) {
	p.Singles = singles
}

func (p BaseballPlayer) GetDoubles() int {
	return p.Doubles
}

func (p *BaseballPlayer) SetDoubles(doubles int) {
	p.Doubles = doubles
}

func (p BaseballPlayer) GetTriples() int {
	return p.Triples
}

func (p *BaseballPlayer) SetTriples(triples int) {
	p.Triples = triples
}

func (p BaseballPlayer) GetHomeruns() int {
	return p.HomeRuns
}

func (p *BaseballPlayer) SetHomeruns(homeruns int) {
	p.HomeRuns = homeruns
}

func (p BaseballPlayer) GetWalks() int {
	return p.Walks
}

func (p *BaseballPlayer) SetWalks(walks int) {
	p.Walks = walks
}

func (p BaseballPlayer) GetHitByPitch() int {
	return p.HitByPitch
}

func (p *BaseballPlayer) SetHitByPitch(hbp int) {
	p.HitByPitch = hbp
}
