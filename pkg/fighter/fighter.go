package fighter

type Fighter interface {
	// Name returns the fighter's name
	Name() string

	// Health returns fighter's baseHealth at the moment when game starts
	Health() int

	// CurrentHealth returns fighter's current baseHealth calculated after some damage
	CurrentHealth() int

	// BodyParts returns list of fighter's body parts which can be attacked
	BodyParts() []DamageGetter

	// OrderedBodyPartsNumbers return list of body parts numbers
	OrderedBodyPartsNumbers() []int

	// Attack makes a damage and returns number of damage
	// and returns value telling if the damage is critical or not
	Attack() (int, bool)

	DamageCollector
}

type DamageCollector interface {
	// CollectGottenDamage gets damage and collect it
	CollectGottenDamage(int)
}

func New(cfg *Config, bp []DamageGetter, p Params) Fighter {

	if cfg.BaseHealth <= 0 {
		cfg.BaseHealth = defaultFighterHealth
	}

	if cfg.BaseDamage == nil {
		cfg.BaseDamage = &BaseDamage{defaultMinBaseDamage, defaultMaxBaseDamage}
	}

	g := &unit{
		name:                    cfg.Name,
		baseHealth:              cfg.BaseHealth,
		baseDamage:              cfg.BaseDamage,
		gotDamage:               make([]int, 0),
		bodyPartsOrderedNumbers: make([]int, 0),
	}

	for idx, p := range bp {
		p.SetOwner(g)
		g.bodyParts = append(g.bodyParts, p)
		g.bodyPartsOrderedNumbers = append(g.bodyPartsOrderedNumbers, idx)
	}

	return g
}

type DamageGetter interface {
	// CatchDamage gets damage which make an enemy
	// and returns result calculated damage
	CatchDamage(int, bool) int

	// Name return the name of a body part which gets damage
	Name() string

	// Block sets block on the body part
	Block()

	// Unblock removes block from the body part
	Unblock()

	OwnerSetter
}

type OwnerSetter interface {
	SetOwner(DamageCollector)
}

type Params map[string]string

type unit struct {
	name                    string
	baseHealth              int
	baseDamage              *BaseDamage
	gotDamage               []int
	bodyParts               []DamageGetter
	bodyPartsOrderedNumbers []int
}

func (u *unit) OrderedBodyPartsNumbers() []int {
	return u.bodyPartsOrderedNumbers
}

func (u *unit) CollectGottenDamage(d int) {
	u.gotDamage = append(u.gotDamage, d)
}

func (u *unit) Attack() (int, bool) {
	return u.baseDamage.Min, false
}

func (u *unit) Health() int {
	return u.baseHealth
}

func (u *unit) Name() string {
	return u.name
}

func (u *unit) CurrentHealth() int {
	return u.baseHealth - u.totalGottenDamage()
}

func (u *unit) BodyParts() []DamageGetter {
	return u.bodyParts
}

func (u *unit) totalGottenDamage() (total int) {
	for _, d := range u.gotDamage {
		total += d
	}
	return
}
