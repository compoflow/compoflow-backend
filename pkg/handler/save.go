package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/minio/minio-go/v6"
)

type SaveRequest struct {
	OrgName     string `json:"org_name"`
	ProjectName string `json:"project_name"`
	XmlStr      string `json:"xml"`
}

type SaveResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// SaveHandler save xml to minio
// minio bucket name is organization name
// xml file dir is '{projectName}/bpmn.xml'
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	logger := HandlerLogger.WithName("SaveHandler")

	if r.Method == "POST" {
		logger.Info("Handle post request of save xml")

		var saveRequest SaveRequest
		data, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(err, "Fail to read request body")
			return
		}
		err = json.Unmarshal(data, &saveRequest)
		if err != nil {
			logger.Error(err, "Fail to parse xml str")
			return
		}

		// Get xml data
		xmlstr := saveRequest.XmlStr

		// init minio client
		minioClient, err := minio.New(
			os.Getenv("ENDPOINT"),
			os.Getenv("ACCESS_KEY_ID"),
			os.Getenv("SECRET_ACCESS_KEY"),
			false)
		if err != nil {
			logger.Error(err, "Fail to connect minio")
			return
		}

		// check project bucket exist or not
		bucketName := saveRequest.OrgName
		found, err := minioClient.BucketExists(bucketName)
		if err != nil {
			logger.Error(err, "Fail to check bucket {}", bucketName)
			return
		}
		if !found {
			// if not found, create bucket
			err = minioClient.MakeBucket(bucketName, "us-east-1")
			if err != nil {
				logger.Error(err, "Fail to create bucket {}", bucketName)
				return
			}
			logger.Info("Successfully created {}.", bucketName)
		}

		filename := saveRequest.ProjectName + "/bpmn.xml"
		file := bytes.NewReader([]byte(xmlstr))
		n, err := minioClient.PutObject(bucketName, filename, file,
			file.Size(), minio.PutObjectOptions{ContentType: "application/xml"})
		if err != nil {
			logger.Error(err, "Fail to put xml to minio {}/{}", bucketName, filename)
			return
		}
		logger.Info("Successfully uploaded bytes: {}", n)
		resp := SaveResponse{
			Code: 200,
			Msg:  "success",
		}
		respStr, _ := json.Marshal(resp)
		w.Write(respStr)
	}
}
