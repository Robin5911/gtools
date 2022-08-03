package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CheckPathExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func ReadFileAll(filePath string) ([]byte,error){
	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data,nil
}
func WriteFileAll(filePath string,data []byte,perm os.FileMode) error{
	err := ioutil.WriteFile(filePath,data,perm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
