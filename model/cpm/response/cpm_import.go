package response

type AutoListUrl struct {
	List string `json:"List" gorm:"column:table_name"`
	Name string `json:"name" gorm:"comment:组件名字"`
}
