package merchant

import (
	"bufio"
	"compay/common/constants"
	"compay/src/models/user"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type requestAdminLogReaderTail struct {
	FileName string `json:"file_name" binding:"required,max=255"`
	Page     int    `json:"page" binding:"omitempty,gt=0"`
	PerPage  int    `json:"per_page" binding:"omitempty,gt=0,max=500"`
}

type logReaderInfo struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func AdminLogReaderList(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	entries, err := os.ReadDir("resources/logs")
	if err != nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	files := []logReaderInfo{}
	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		files = append(files, logReaderInfo{
			Name: entry.Name(),
			Size: info.Size(),
		})
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": files})
}

func AdminLogReaderTail(context *gin.Context) {
	adminUserModel := context.MustGet(constants.GinUser).(*user.Model)
	if !isUserRootAdmin(adminUserModel.GetId()) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	var input requestAdminLogReaderTail
	if err := context.ShouldBindJSON(&input); err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	if strings.Contains(input.FileName, "/") || strings.Contains(input.FileName, "\\") || strings.Contains(input.FileName, "..") {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	perPage := 100
	page := input.Page
	if page <= 0 {
		page = 1
	}

	filePath := filepath.Join("resources/logs", input.FileName)

	file, err := os.Open(filePath)
	if err != nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}
	defer file.Close()

	totalLines := countLines(file)

	totalPages := totalLines / perPage
	if totalLines%perPage != 0 {
		totalPages++
	}
	if totalPages == 0 {
		totalPages = 1
	}

	skipFromEnd := (page - 1) * perPage
	startLine := totalLines - skipFromEnd - perPage
	if startLine < 0 {
		perPage += startLine
		startLine = 0
	}

	if perPage <= 0 {
		efwginhelpers.ResponseSuccess(context, gin.H{
			"data":        []string{},
			"total_lines": totalLines,
			"total_pages": totalPages,
			"page":        page,
			"per_page":    input.PerPage,
		})
		return
	}

	lines := readLinesRange(file, startLine, perPage)

	efwginhelpers.ResponseSuccess(context, gin.H{
		"data":        lines,
		"total_lines": totalLines,
		"total_pages": totalPages,
		"page":        page,
		"per_page":    perPage,
	})
}

func countLines(file *os.File) int {
	file.Seek(0, io.SeekStart)

	count := 0
	buf := make([]byte, 32*1024)
	for {
		n, err := file.Read(buf)
		for i := range n {
			if buf[i] == '\n' {
				count++
			}
		}
		if err != nil {
			break
		}
	}

	return count
}

func readLinesRange(file *os.File, start int, count int) []string {
	file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	lines := make([]string, 0, count)
	lineNum := 0

	for scanner.Scan() {
		if lineNum >= start+count {
			break
		}
		if lineNum >= start {
			lines = append(lines, scanner.Text())
		}
		lineNum++
	}

	return lines
}
