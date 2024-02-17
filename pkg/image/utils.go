package Image

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const IMAGE_DIR = "./image"
const STAT_IMAGE_PATH = "/stat-img"

// func randomFilename() string {
// 	t := time.Now()
// 	entropy := rand.New(rand.NewSource(t.UnixNano()))
// 	return strings.ToLower(fmt.Sprintf("%v", ulid.MustNew(ulid.Timestamp(t), entropy)))
// }

func httpError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"ok":      false,
		"message": message,
	})
	c.Abort()
}

func httpSuccess(c *gin.Context, res interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"result": res,
	})
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
