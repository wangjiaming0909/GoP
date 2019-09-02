package userservice

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//UserRegisterCallBack callback
func UserRegisterCallBack(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	fmt.Println(string(x))
	c.Writer.WriteString("hello")
}

func registerUser(ui *UserInfo) {

}
