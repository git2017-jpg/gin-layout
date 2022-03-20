package main

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// @title gin-layout
// @version 1.0
// @description gin web 开发模板
// @termsOfService http://swagger.io/terms/

// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1：9080
// @BasePath /v1/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp().Run()
}
