package main
import "fmt"
import "io"
import "os"

func main(){
    w,err := CopyFile("e:/temp/1.jpg","e:/temp/2.jpg")
    if err!=nil{
        fmt.Println(err.Error())
    }

    fmt.Println(w)
}

func CopyFile(src,dst string)(w int64,err error){
    srcFile,err := os.Open(src)
    if err!=nil{
        fmt.Println(err.Error())
        return
    }
    defer srcFile.Close()

    dstFile,err := os.Create(dst)

    if err!=nil{
        fmt.Println(err.Error())
        return
    }

    defer dstFile.Close()

    return io.Copy(dstFile,srcFile)
}