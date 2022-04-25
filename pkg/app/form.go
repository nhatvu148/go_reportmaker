package app

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/nhatvu148/go_reportmaker/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	// jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	// if err1 != nil {
	// 	// Handle error
	// }
	// fmt.Println(string(jsonData))

	err := c.Bind(form)
	if err != nil {
		fmt.Println(c.ContentType())
		fmt.Println(c.Request.Method)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
