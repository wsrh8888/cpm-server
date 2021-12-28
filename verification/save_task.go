package verification

type SaveTask struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Name    string `form:"name" json:"name"  binding:"required"`
	Version string `form:"version" json:"version"`
	Main    string `form:"main" json:"main" binding:"required"`
	Author  string `form:"author" json:"author" binding:"required"`
}
