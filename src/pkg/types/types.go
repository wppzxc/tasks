package types

const (
	AdminUser = "administrator"
)

type Config struct {
	BindIP string
	BindPort string
	MysqlUrl string
	Version  bool
}
