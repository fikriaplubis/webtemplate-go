package file

import (
	"bufio"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type Service interface {
	Save(header multipart.FileHeader, path string, filename string) (bool, string)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) Save(header multipart.FileHeader, path string, filename string) (bool, string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return true, err.Error()
	}

	dst, err := os.Create(path + filename)
	if err != nil {
		return true, err.Error()
	}

	defer dst.Close()

	file, _ := header.Open()
	_, err = io.Copy(dst, file)
	if err != nil {
		return true, err.Error()
	}

	return false, ""
}

func (s *service) Decrypt(header multipart.FileHeader) []string {
	var result []string
	strPwd := "SINARMAS"

	file, _ := header.Open()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		strBuff := ""
		mod := 0
		a := 0
		x := 1
		txt := fileScanner.Text()
		for i := 0; i < len(txt)/2; i++ {
			c := txt[a : a+2]
			dec, _ := strconv.ParseInt(c, 16, 64)
			mod = x % len(strPwd)
			isimod := strPwd[mod : mod+1]
			isiASC := int64(isimod[0])
			c = string(rune(dec - isiASC))
			strBuff = strBuff + c
			x += 1
			a += 2
		}
		result = append(result, strBuff)
	}

	return result
}
