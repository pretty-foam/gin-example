package models

import "github.com/jinzhu/gorm"

//Info 用户信息
type Info struct{
	Name string 
	Age int 
	Sex string
	*gorm.Model
}

//CreateInfo 创建信息
func (i *Info) CreateInfo()  {
 	
}
