package models

import (
	"github.com/users/LENOVO/Downloads/Stumble/pkg/config"
	"github.com/jinzhu/gorm"
)

var db * gorm.DB

type User struct{
	gorm.Model
	id int `gorm:"" json: "id"`
	name string `json: "name"`
	location float32 `json: "location"`
	gender string `json: "gender"`
	email string `json: "email"`
}

type Like struct{
	gorm.Model
	id int `gorm:""json:"id"`
	who_likes int `json:"who_likes"`
	who_is_liked int `json:"who_is_liked"`
}

type Match struct{
	who_likes int
	who_is_liked int
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{},&Like{})
}

func GetMatches() (*[]Match, *gorm.DB){
	var likes_pair []Like
	db := db.Table("likes").Select("who_likes","who_is_liked").Find(&likes_pair)
	i := 0
	var matches []Match
	for i<len(likes_pair) {
		m := likes_pair[i].who_likes
		n := likes_pair[i].who_is_liked
		o := likes_pair[i].id
		j := 0
		var ind Match
		for j < len(likes_pair){
			w := likes_pair[j].id
			e := likes_pair[j].who_likes
			r := likes_pair[j].who_is_liked
			if((m==r) && (n == e)){
				ind.who_likes = m
				ind.who_is_liked = n
				RemoveIndex(likes_pair, o)
				RemoveIndex(likes_pair, w)
			}
		}
		matches = append(matches, ind)
		i = i+1
	}

	return &matches, db
}

func RemoveIndex(s []Like, index int) []Like {
    ret := make([]Like, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}


func GetAllWithinK(name string, k int) ([]string, *gorm.DB){
	var usr_by_name User
	db := db.Table("users").Where("name = ?",name).Find(&usr_by_name)

	loc := usr_by_name.location

	low := loc - float32(k)
	high := loc + float32(k)

	var usr_within_k []User
	db1 := db.Table("users").Where("location BETWEEN ? AND ?", low, high).Find(&usr_within_k)

	n := len(usr_within_k)

	var name_usr_within_k []string

	i := 0
	for i<n{
		name_usr_within_k[i] = usr_within_k[i].name
	}

	return name_usr_within_k, db1
}

func GetUsersByQuery(query string) ([]string, *gorm.DB){
	var usr_by_query []string

	db := db.Table("users").Select("name").Where("name LIKE ?","%"+query+"%").Find(&usr_by_query)

	return usr_by_query, db
}
