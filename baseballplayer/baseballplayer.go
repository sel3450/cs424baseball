package baseballplayer

type BaseballPlayer struct {
	FirstName        string
	LastName         string
	PlateAppearance  int
	AtBats           int
	Singles          int
	Doubles          int
	Triples          int
	HomeRuns         int
	Walks            int
	HitByPitch       int
	BattingAverage   float64
	Slugging         float64
	OnbasePercentage float64
	OPS              float64
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

func (p *BaseballPlayer) ComputeBattingAverage() {
	var battingaverage float64 = (float64(p.Singles) + float64(p.Doubles) + float64(p.Triples) + float64(p.HomeRuns)) / float64(p.AtBats)
	p.BattingAverage = battingaverage
}

func (p *BaseballPlayer) ComputeSlugging() {
	var slugging float64 = (float64(p.Singles) + (2 * float64(p.Doubles)) + (3 * float64(p.Triples)) + (4 * float64(p.HomeRuns))) / float64(p.AtBats)
	p.Slugging = slugging
}

func (p *BaseballPlayer) ComputeOnBase() {
	var onbasepercentage float64 = (float64(p.Singles) + float64(p.Doubles) + float64(p.Triples) + float64(p.HomeRuns) + float64(p.Walks) + float64(p.HitByPitch)) / float64(p.PlateAppearance)
	p.OnbasePercentage = onbasepercentage
}

func (p *BaseballPlayer) ComputeOPS() {
	var ops float64 = p.Slugging + p.OnbasePercentage
	p.OPS = ops
}
