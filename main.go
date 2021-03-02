package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//	헬스 체크 수행하기. 
	r.GET("/health", helathCheck)

	//	:greet 에 해당하는 path parameter 를 획득하여 인사를 반환한다. 
	r.GET("/hello/:greet", greeting)

	r.Run()
}

func helathCheck(c *gin.Context) {

	line := readFile("health")
	if strings.Contains(string(line), "OK") {
		c.String(http.StatusOK, string(line))
	} else {
		c.String(http.StatusInternalServerError, string(line))
	}
}

func readFile(fileName string) string {

	fo, err := os.Open("health")
	if err != nil {
		return ""
	}

	reader := bufio.NewReader(fo)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break;
		} else {
			return string(line)
		}
	}

	return ""
}

func greeting(c *gin.Context) {
	greet := c.Param("greet")
	c.String(http.StatusOK, "Hello " + greet)
}