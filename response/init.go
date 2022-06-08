package response

import (
	"BasicOA/serve/regular"
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

// PostForm 验证模块
func PostForm(c *gin.Context, pattern string, key string) (string, int, error) {
	var (
		err  error
		ints int
	)
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(c.PostForm(key)) {
		err = nil
	} else {
		err = errors.New(key + "正则验证不通过")
	}
	if isInt(pattern) {
		if ints, err = strconv.Atoi(c.PostForm(key)); err != nil {
			err = errors.New(key + "数据int验证不通过")
		}
	}
	return c.PostForm(key), ints, err
}

// GetCookie 验证模块
func GetCookie(c *gin.Context, pattern string, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", errors.New(key + "未获取到cookie")
	}
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(cookie) {
		err = nil
	} else {
		err = errors.New(key + "正则验证不通过")
	}
	return cookie, err
}

// GetParam 验证模块
func GetParam(c *gin.Context, pattern string, key string) (string, int, error) {
	var (
		err  error
		ints int
	)
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(c.Param(key)) {
		err = nil
	} else {
		err = errors.New(key + "正则验证不通过")
	}
	if isInt(pattern) {
		if ints, err = strconv.Atoi(c.Param(key)); err != nil {
			err = errors.New(key + "数据int验证不通过")
		}
	}
	return c.Param(key), ints, err
}

func GetQuery(c *gin.Context, pattern string, key string) (string, int, error) {
	var (
		err  error
		ints int
	)
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(c.Query(key)) {
		err = nil
	} else {
		err = errors.New(key + "正则验证不通过")
	}
	if isInt(pattern) {
		if ints, err = strconv.Atoi(c.Query(key)); err != nil {
			err = errors.New(key + "数据int验证不通过")
		}
	}
	return c.Query(key), ints, err
}

func isInt(item string) bool {
	for _, eachItem := range regular.IsInt {
		if eachItem == item {
			return true
		}
	}
	return false
}

func GetHeader(c *gin.Context, pattern string, key string) (string, error) {
	var err error
	reg := regexp.MustCompile(pattern)
	if reg.MatchString(c.Request.Header.Get(key)) {
		err = nil
	} else {
		err = errors.New(key + "正则验证不通过")
	}
	return c.Request.Header.Get(key), err
}
