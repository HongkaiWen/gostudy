package util
import "intf"

func readFile(file intf.IFile) {
	file.Read()
}
