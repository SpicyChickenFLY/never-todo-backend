package controller

import (
	"github.com/gin-gonic/gin"
)

// GetUserInfoByID get UserInfo
func GetUserInfoByID(c *gin.Context) {
	// var userInfos model.UserInfos = model.UserInfos{}
	// tx := mysql.GormDB.Begin()
	// err := service.GetUserInfoByID(tx, &userInfos)
	// result := gin.H{
	// 	"userInfos": userInfos,
	// }
	// err = mysql.CheckTransaction(tx, err)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError,
	// 		gin.H{"Status": -1, "Result": result})
	// 	log.Print(err)
	// } else {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": 0, "Result": result})
	// }
}

// AddUserInfo add UserInfo
func AddUserInfo(c *gin.Context) {
	// data := &struct {
	// 	UserInfoContent string `json:"UserInfoContent"`
	// 	UserInfoDesc    string `json:"UserInfoDesc"`
	// }{}
	// c.BindJSON(&data)
	// log.Printf("receive post request: %v", data)
	// userInfo := model.UserInfo{Content: data.UserInfoContent, Desc: data.UserInfoDesc}
	// tx := mysql.GormDB.Begin()
	// err := service.AddUserInfo(tx, userInfo)
	// err = mysql.CheckTransaction(tx, err)
	// if err != nil {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": -1, "Result": userInfo})
	// 	log.Print(err)
	// } else {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": 0, "Result": userInfo})
	// }
}

// DelUserInfo add UserInfo
func DelUserInfo(c *gin.Context) {
	// userInfoID, err := strconv.Atoi(c.Param("userInfo_id"))
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusInternalServerError,
	// 		gin.H{"Status": -1})
	// 	return
	// }
	// tx := mysql.GormDB.Begin()
	// err = service.DelUserInfo(tx, userInfoID)
	// err = mysql.CheckTransaction(tx, err)
	// if err != nil {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": -1})
	// 	log.Print(err)
	// } else {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": 0})
	// }
}

// UpdUserInfo update UserInfo
func UpdUserInfo(c *gin.Context) {
	// data := &struct {
	// 	UserInfoID      int    `json:"UserInfoID"`
	// 	UserInfoContent string `json:"UserInfoContent"`
	// 	UserInfoDesc    string `json:"UserInfoDesc"`
	// }{}
	// c.BindJSON(&data)
	// userInfo := model.UserInfo{
	// 	ID:      data.UserInfoID,
	// 	Content: data.UserInfoContent,
	// 	Desc:    data.UserInfoDesc,
	// }
	// log.Printf("receive post request: %v", data)
	// tx := mysql.GormDB.Begin()
	// err := service.UpdUserInfo(tx, &userInfo)
	// err = mysql.CheckTransaction(tx, err)
	// if err != nil {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": -1, "Result": userInfo})
	// 	log.Print(err)
	// } else {
	// 	c.JSON(http.StatusOK,
	// 		gin.H{"Status": 0, "Result": userInfo})
	// }
}
