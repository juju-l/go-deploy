package rt

import (
	// "fmt"
	"go-deploy/util/ext"
	"go-deploy/util/getarchives"
	"os"
	"strconv"
	"time"
)

//

func CronJob() {
	for range time.Tick(
			time.Duration(ext.Must(strconv.Atoi(os.Getenv("cycle"))).(int)) * time.Second,
		) {
	getarchives.Zip()
	}
}

func init() {
	//
}