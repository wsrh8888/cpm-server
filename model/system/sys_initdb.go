package system

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

func MysqlDataInitialize(inits ...InitData) error {
	for i := 0; i < len(inits); i++ {
		if inits[i].CheckDataExist() {
			continue
		}

		if err := inits[i].Initialize(); err != nil {
			return err
		} else {
			continue
		}
	}
	return nil
}
