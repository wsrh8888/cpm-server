package system

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

func MysqlDataInitialize(inits ...InitData) error {
	return nil
}
