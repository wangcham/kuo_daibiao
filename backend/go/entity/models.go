package entity

//定义用户表名

func (User) TableName() string{
	return "users"
} 

type User struct {
    ID           uint   `gorm:"primaryKey;autoIncrement"`
    Username     string `gorm:"unique"`
    Password     string
    Name         string
    ProfileText  string
    PhotoPath    string
    Identity     string
}

//定义老师代课表名

func (Course) TableName() string{
	return "courses"
}

type Course struct {
    Type         int    `gorm:"foreignKey:AllCourses.Type"`
    CourseName   string `gorm:"foreignKey:AllCourses.CourseName"`
    TeacherName  string `gorm:"foreignKey:User.Name"`
    TeacherUsername  string `gorm:"foreignKey:User.Username"`
    TeacherID  uint `gorm:"foreignKey:User.ID"`
}
//定义全部课程

func (AllCourse) TableName() string{
	return "allcourses"
}
type AllCourse struct{
	Type int //1为初中，2为高中，3为大学
	CourseName string
}

