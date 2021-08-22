package fighter

const (
	defaultFighterHealth = 50
	defaultMinBaseDamage = 5
	defaultMaxBaseDamage = 10
)

type Config struct {
	Name       string
	BaseHealth int
	BaseDamage *BaseDamage
}

type BaseDamage struct {
	Min, Max int
}

type PartConfig struct {
	Name string
}
