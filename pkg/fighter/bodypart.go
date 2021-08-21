package fighter

type Part interface {
	// CatchDamage gets damage which make the fighter
	// and returns calculated damage
	CatchDamage(int, bool) int

	// Name return the name a body part which gets damage
	Name() string

	// Block sets block on the body part
	Block()

	// Unblock removes block from the body part
	Unblock()

	// SetArmor sets armor to the body part
	SetArmor(Armor) bool

	// SetOwner sets it's owner which wil get the damage
	SetOwner(DamageCollector)
}

func NewPart(cfg *PartConfig) Part {

	p := &part{
		name: cfg.Name,
	}

	return p
}

// part an any part of body
type part struct {
	name      string
	gotDamage int
	armor     Armor
	blocked   bool
	owner     DamageCollector
}

func (p *part) SetOwner(dc DamageCollector) {
	p.owner = dc
}

func (p *part) CatchDamage(damage int, critical bool) int {
	calculatedDamage := p.calculatedReceivedDamage(damage, critical)

	p.owner.CollectGottenDamage(calculatedDamage)

	return calculatedDamage
}

func (p *part) calculatedReceivedDamage(damage int, critical bool) int {
	if p.blocked != critical {
		return 0
	}

	return damage
}

func (p *part) Name() string {
	return p.name
}

func (p *part) SetArmor(a Armor) bool {
	p.armor = a

	return true
}

func (p *part) Block() {
	p.blocked = true
}

func (p *part) Unblock() {
	p.blocked = false
}
