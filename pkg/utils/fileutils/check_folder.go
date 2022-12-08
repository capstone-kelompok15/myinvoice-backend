package fileutils

import "os"

var isUploadFolderExist bool

func CheckUploadFolder() bool {
	if !isUploadFolderExist {
		os.Mkdir("uploads", os.ModePerm)
		isUploadFolderExist = true
	}
	return true
}
