package fighter

type Fighter interface {
	// Name returns the fighter's name
	Name() string

	// Health returns fighter's health at the moment when game starts
	Health() int32

	// CurrentHealth returns fighter's current health calculated after some damage
	CurrentHealth() int32

	// BodyParts returns list of fighter's body parts which can be attacked
	BodyParts() []DamageGetter

	// Attack makes a damage and returns number of damage
	// and returns value telling if the damage is critical or not
	Attack() (int32, bool)

	DamageCollector
}

type DamageCollector interface {
	// CollectGottenDamage gets damage and collect it
	CollectGottenDamage(int32)
}

const (
	defaultFighterHealth = 50
)

func New(cfg *Config, bp []DamageGetter, p Params) Fighter {

	g := &unit{
		name:       cfg.Name,
		health:     cfg.Health,
		baseDamage: cfg.BaseDamage,
		gotDamage:  make([]int32, 0),
	}

	for _, p := range bp {
		p.SetOwner(g)
		g.bodyParts = append(g.bodyParts, p)
	}

	return g
}

type DamageGetter interface {
	// CatchDamage gets damage which make an enemy
	// and returns result calculated damage
	CatchDamage(int32, bool) int32

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
	name       string
	health     int32
	baseDamage int32
	gotDamage  []int32
	bodyParts  []DamageGetter
}

func (u *unit) CollectGottenDamage(d int32) {
	u.gotDamage = append(u.gotDamage, d)
}

func (u *unit) Attack() (int32, bool) {
	return u.baseDamage, false
}

func (u *unit) Health() int32 {
	return u.health
}

func (u *unit) Name() string {
	return u.name
}

func (u *unit) CurrentHealth() int32 {
	return u.health - u.totalGottenDamage()
}

func (u *unit) BodyParts() []DamageGetter {
	return u.bodyParts
}

func (u *unit) totalGottenDamage() (total int32) {
	for _, d := range u.gotDamage {
		total += d
	}
	return
}
