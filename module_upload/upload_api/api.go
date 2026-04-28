package upload_api

import (
	"compress/gzip"
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"

	"github.com/team-ide/framework/util"
	"github.com/team-ide/framework/web"
)

var defaultFilesSaveDir = "./data/files"
var defaultFilesSavePathPrefix = "uploads"
var defaultUploadPath = "upload"
var defaultFilesPathPrefix = "files"
var defaultDownloadPath = "download"

type Config struct {
	// FilesSaveDir 文件保存目录
	FilesSaveDir string
	// FilesSavePathPrefix 文件保存路径前缀
	FilesSavePathPrefix string
	// UploadPath 上传请求路径
	UploadPath string
	// FilesPathSuffix 文件请求路径前缀
	FilesPathPrefix string
	// DownloadPath 下载请求路径
	DownloadPath string
}

func NewUploadService(cfg *Config) *UploadService {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.FilesSaveDir == "" {
		cfg.FilesSaveDir = defaultFilesSaveDir
	}
	if cfg.FilesSavePathPrefix == "" {
		cfg.FilesSavePathPrefix = defaultFilesSavePathPrefix
	}
	if cfg.UploadPath == "" {
		cfg.UploadPath = defaultUploadPath
	}
	if cfg.FilesPathPrefix == "" {
		cfg.FilesPathPrefix = defaultFilesPathPrefix
	}
	if cfg.DownloadPath == "" {
		cfg.DownloadPath = defaultDownloadPath
	}
	res := &UploadService{}
	res.Config = cfg

	return res
}

type UploadService struct {
	*Config
}

func (this_ *UploadService) GetWebApi() (webApi *web.WebApi) {
	webApi = web.NewWebApi("/")
	webApi.Add(this_.UploadPath, this_.Upload).SetNotLogin().SetNotLog().SetComment("文件上传")
	webApi.Add(this_.FilesPathPrefix+"/*", this_.Files).SetGet().SetNotLogin().SetNotLog().SetComment("文件访问")
	webApi.Add(this_.DownloadPath, this_.Download).SetNotLogin().SetNotLog().SetComment("文件下载")
	return
}

type UploadResponse struct {
	Files []*FileInfo `json:"files"`
}

func (this_ *UploadService) Upload(request *web.WebRequest) (res any, err error) {
	var quality int
	qualityV := request.GetParam("quality")
	if qualityV != "" {
		quality = util.StringToInt(qualityV)
	}
	files, err := request.GetFiles("file")
	if err != nil {
		return
	}
	defer func() {
		for _, f := range files {
			_ = f.Close()
		}
	}()

	r := &UploadResponse{}
	var f *FileInfo
	for _, file := range files {
		f, err = this_.Save(file, quality)
		if err != nil {
			return
		}
		r.Files = append(r.Files, f)
	}

	res = r
	return
}

func (this_ *UploadService) GetFileSavePath(filePath string) string {
	savePath := util.PathJoin(this_.FilesSaveDir, filePath)
	savePath = util.FormatPath(savePath)
	return savePath
}
func (this_ *UploadService) Save(uploadFile *web.UploadFile, quality int) (res *FileInfo, err error) {

	nowTime := time.Now()
	filePath := fmt.Sprintf("%s/%d/%d/%d/%s", this_.FilesSavePathPrefix, nowTime.Year(), nowTime.Month(), nowTime.Day(), util.GetUuid())

	defer func() { _ = uploadFile.Close() }()
	fileName := uploadFile.Filename
	//获取文件名称带后缀
	fileNameWithSuffix := path.Base(uploadFile.Filename)
	//获取文件的后缀(文件类型)
	fileType := path.Ext(fileNameWithSuffix)
	filePath = util.PathJoin(filePath, fileName)

	savePath := this_.GetFileSavePath(filePath)
	saveDir := filepath.Dir(savePath)
	if exist, _ := util.PathExists(saveDir); !exist {
		err = os.MkdirAll(saveDir, os.ModePerm)
		if err != nil {
			return
		}
	}

	createFile, err := os.Create(savePath)
	if err != nil {
		return
	}
	defer func() { _ = createFile.Close() }()

	written, err := io.Copy(createFile, uploadFile)
	if err != nil {
		return
	}

	var QualityPath string

	if quality > 0 {
		QualityPath, _ = this_.qualityCompression(filePath, fileType, quality)
	}

	res = &FileInfo{}
	res.Name = fileName
	res.Type = fileType
	res.Path = filePath
	res.QualityPath = QualityPath
	res.Size = written
	return
}

type FileInfo struct {
	Name        string `json:"name,omitempty"`
	Size        int64  `json:"size,omitempty"`
	Type        string `json:"type,omitempty"`
	Path        string `json:"path,omitempty"`
	QualityPath string `json:"qualityPath,omitempty"`
}

func (this_ *UploadService) qualityCompression(filePath string, fileType string, quality int) (res string, err error) {

	f, err := os.Open(this_.GetFileSavePath(filePath))
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()
	img, err := jpeg.Decode(f)
	if err != nil {
		return
	}
	savePath := fmt.Sprintf("%s-%d%s", filePath, quality, fileType)

	err = imaging.Save(img, this_.GetFileSavePath(savePath), imaging.JPEGQuality(quality))
	if err != nil {
		return
	}
	res = savePath
	return
}

func (this_ *UploadService) Files(request *web.WebRequest) (res any, err error) {
	index := strings.Index(request.Path, this_.FilesPathPrefix)
	if index == -1 {
		err = web.NotFoundError
		return
	}
	filePath := request.Path[index+len(this_.FilesPathPrefix):]

	fileSavePath := this_.GetFileSavePath(filePath)

	if isSub, _ := util.IsSubPath(this_.FilesSaveDir, fileSavePath); !isSub {
		err = web.NotFoundError
		return
	}
	if exist, _ := util.PathExists(fileSavePath); !exist {
		err = web.NotFoundError
		return
	}
	f, err := os.Open(fileSavePath)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()
	request.SetStatus(http.StatusOK)
	res = web.WebNotResponse
	_, err = io.Copy(request.GetWriter(), f)
	if err != nil {
		return
	}
	return
}

func (this_ *UploadService) Download(request *web.WebRequest) (res any, err error) {
	filePath := request.GetParam("path")

	fileSavePath := this_.GetFileSavePath(filePath)
	if isSub, _ := util.IsSubPath(this_.FilesSaveDir, fileSavePath); !isSub {
		err = web.NotFoundError
		return
	}
	if exist, _ := util.PathExists(fileSavePath); !exist {
		err = web.NotFoundError
		return
	}
	f, err := os.Open(fileSavePath)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()

	request.SetStatus(http.StatusOK)
	res = web.WebNotResponse
	request.SetHeader("Content-Type", "application/octet-stream")
	request.SetHeader("Content-Transfer-Encoding", "binary")

	fileName := request.GetParam("fileName")
	if fileName == "" {
		fileName = path.Base(filePath)
	}
	isZip := request.GetParam("zip") == "true"
	if fileName == "" {
		fileName = path.Base(filePath)
	}
	if isZip {
		fileName += ".gz"
	}

	request.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", url.QueryEscape(fileName)))
	request.SetHeader("download-file-name", fileName)
	if isZip {
		gz := gzip.NewWriter(request.GetWriter())
		defer func() { _ = gz.Close() }()
		_, err = io.Copy(gz, f)
	} else {
		_, err = io.Copy(request.GetWriter(), f)
	}
	if err != nil {
		return
	}
	return
}
