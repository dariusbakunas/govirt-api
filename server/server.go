package server

import (
	"fmt"
)

func Init(port int) {
	r := NewRouter()
	r.Run(fmt.Sprintf(":%d", port))
}