package fighter

const (
	defaultFighterHealth = 50
)

type Config struct {
	Name       string
	Health     int
	BaseDamage int
}

type PartConfig struct {
	Name string
}
