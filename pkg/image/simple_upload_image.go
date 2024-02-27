package Image

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SimpleUploadImage(c *gin.Context) {
	productId, _ := c.GetPostForm("product_id")
	{
		if len(productId) == 0 {
			productId = "default"
		}
		err := os.MkdirAll(fmt.Sprintf("%s/product/%s", IMAGE_DIR, productId), os.ModePerm)
		if err != nil {
			httpError(c, http.StatusBadRequest, fmt.Sprintf("uploadImage os.MkdirAll error: %s", err))
			return
		}
	}
	path := fmt.Sprintf("%s/product/%s", IMAGE_DIR, productId)

	form, _ := c.MultipartForm()
	var fileName string
	imgExt := "jpeg"
	for key := range form.File {
		fileName = key
		arr := strings.Split(fileName, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr)-1]
		}
		continue
	}
	file, _, err := c.Request.FormFile(fileName)
	if err != nil {
		httpError(c, http.StatusBadRequest, fmt.Sprintf("UploadXml c.Request.FormFile error: %s", err.Error()))
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}

	fullFileName := fmt.Sprintf("%s.%s", RandomFilename(), imgExt)
	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	if err != nil {
		httpError(c, http.StatusBadRequest, fmt.Sprintf("uploadImage os.Create err: %s", err))
		return
	}
	defer fileOnDisk.Close()

	_, err = fileOnDisk.Write(fileBytes)
	if err != nil {
		httpError(c, http.StatusBadRequest, err.Error())
		return
	}

	httpSuccess(c, map[string]string{"file": fmt.Sprintf("%s/%s", strings.Replace(path, IMAGE_DIR, STAT_IMAGE_PATH, 1), fullFileName)})

}
