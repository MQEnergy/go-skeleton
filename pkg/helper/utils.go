package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand/v2"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hashicorp/go-uuid"
	"github.com/spf13/cast"
	"gorm.io/gorm/schema"
)

// InAnySlice 判断某个字符串是否在字符串切片中
func InAnySlice[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// InAnyMap 判断某个map的值是否存在
func InAnyMap[T comparable](haystack map[string]T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// GetKeyByMap 根据map中的值获取键
func GetKeyByMap[T comparable](m map[string]T, value T) string {
	for key, val := range m {
		if val == value {
			return key
		}
	}
	return ""
}

// GenerateBaseSnowId 生成雪花算法ID
func GenerateBaseSnowId(num int, n *snowflake.Node) string {
	if n == nil {
		localIp, err := GetLocalIpToInt()
		if err != nil {
			return ""
		}
		node, err := snowflake.NewNode(int64(localIp) % 1023)
		n = node
	}
	id := n.Generate()
	switch num {
	case 2:
		return id.Base2()
	case 32:
		return id.Base32()
	case 36:
		return id.Base36()
	case 58:
		return id.Base58()
	case 64:
		return id.Base64()
	default:
		return cast.ToString(id.Int64())
	}
}

// GenerateUuid 生成随机字符串
func GenerateUuid(size int) string {
	str, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return gstr.SubStr(str, 0, size)
}

// RandString 随机字符串
func RandString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.IntN(len(charset))]
	}
	return string(b)
}

// GeneratePasswordHash 生成密码hash值
func GeneratePasswordHash(password string, salt string) string {
	s := sha256.New()
	io.WriteString(s, password+salt)
	str := fmt.Sprintf("%x", s.Sum(nil))
	return str
}

// GenerateHash 生成md5 hash值
func GenerateHash(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// IsPathExist 判断所给路径文件/文件夹是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// MakeMultiDir 调用os.MkdirAll递归创建文件夹
func MakeMultiDir(filePath string) error {
	if !IsPathExist(filePath) {
		return os.MkdirAll(filePath, os.ModePerm)
	}
	return nil
}

// MakeFileOrPath 创建文件/文件夹
func MakeFileOrPath(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// WriteContentToFile
// @Description: 写文件
// @param filePath
// @return error
func WriteContentToFile(file *multipart.FileHeader, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	open, err := file.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	fileBytes, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	if _, err := f.Write(fileBytes); err != nil {
		return err
	}
	return nil
}

// MakeTimeFormatDir
// @Description: 创建时间格式的目录 如：upload/{path}/2023-01-07/
// @param rootPath 根目录
// @param pathName 子目录名称
// @param timeFormat 时间格式 如：2006-01-02、20060102
// @return string
// @return error
func MakeTimeFormatDir(rootPath, pathName, timeFormat string) (string, error) {
	filePath := "upload/"
	if pathName != "" {
		filePath += pathName + "/"
	}
	filePath += time.Now().Format(timeFormat) + "/"
	if err := MakeMultiDir(rootPath + filePath); err != nil {
		return "", err
	}
	return filePath, nil
}

// String2Int 将数组的string转int
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

// GetStructColumnName 获取结构体中的字段名称 _type: 1: 获取tag字段值 2：获取结构体字段值
func GetStructColumnName(s interface{}, _type int) ([]string, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return []string{}, fmt.Errorf("interface is not a struct")
	}
	t := v.Type()
	var fields []string
	for i := 0; i < v.NumField(); i++ {
		var field string
		if _type == 1 {
			field = t.Field(i).Tag.Get("json")
			if field == "" {
				tagSetting := schema.ParseTagSetting(t.Field(i).Tag.Get("gorm"), ";")
				field = tagSetting["COLUMN"]
			}
		} else {
			field = t.Field(i).Name
		}
		fields = append(fields, field)
	}
	return fields, nil
}

// GetProjectModuleName 获取当前项目的module名称
func GetProjectModuleName() string {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(output), "\n")
}

// GetFileNamesByDirPath 获取当前文件夹下的所有文件和文件夹名称（包括子文件夹和文件）
func GetFileNamesByDirPath(root string) ([]map[string]interface{}, error) {
	paths := make([]map[string]interface{}, 0)
	dirs, err := GetAllDirs(root)
	if err != nil {
		return paths, err
	}
	// 获取每个文件夹的第一级文件列表
	for _, dir := range dirs {
		var pathItem map[string]interface{}
		fileNames := make([]string, 0)
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}
		pathItem = map[string]interface{}{
			"path":  strings.Replace(dir, root, "", 1),
			"files": []string{},
		}
		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".go") {
				continue
			}
			if strings.HasSuffix(file.Name(), "_test.go") {
				continue
			}
			fileNames = append(fileNames, strings.Replace(file.Name(), ".go", "", -1))
		}
		pathItem["files"] = fileNames
		paths = append(paths, pathItem)
	}
	return paths, nil
}

// getFilesInDir 获取指定目录下的所有文件名称
func getFilesInDir(dir string) ([]string, error) {
	files := make([]string, 0)
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否为文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, strings.Replace(info.Name(), ".go", "", -1))
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return files, nil
}

// GetAllDirs 获取指定文件夹中所有文件夹路径
func GetAllDirs(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	})
	return dirs, err
}

// ToCamelCase 将字符串转换成驼峰写法
func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '_' || r == '-' || r == '.' || r == '~' || r == ':'
	})
	var result string
	for i, word := range words {
		if i == 0 {
			result += strings.ToLower(word)
		} else {
			result += strings.Title(word)
		}
	}
	return result
}

// ArrayChunk 数组分组
func ArrayChunk[T any](arr []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}
