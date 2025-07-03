package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err:= os.Open("example.txt") // opening the file
	if err!= nil{
		// log the err.
		// Or PANIC
		panic(err)
	}
	
	fileInfo, err := f.Stat()
	if err!= nil{
		// log the err.
		// Or PANIC
		panic(err)
	}

	defer f.Close() // closing the file

	//!💡 INFOS about the file
	fmt.Println("file-name:",fileInfo.Name())
	fmt.Println("folder or not:",fileInfo.IsDir())
	fmt.Println("file-size:",fileInfo.Size())
	fmt.Println("file permission:",fileInfo.Mode())
	fmt.Println("file modified at:",fileInfo.ModTime())

	//! 💡 READING file-content

	buff :=make([]byte, fileInfo.Size()) // buffer

	d, err := f.Read(buff)
	if err!= nil{
		// log the err.
		// Or PANIC
		panic(err)
	}

	for i:=0; i<len(buff);i++{
		println("data",d, string(buff))
	}
	// Another way easy way (no need to open the file)
	// Good for small files
	// Loads all of the file-content at once, unlike BUFFERs
	f1, err:= os.ReadFile("example.txt")
	if err!= nil{
		panic(err)
	}

	fmt.Println("Using os.ReadFile: ",string(f1))

	//! 💡 Reading FOLDERS
	dir,err := os.Open("../")
	if err!= nil{
		panic(err)
	}

	defer dir.Close();

	dirInfo, err := dir.ReadDir(-1)

	for _,di := range dirInfo{
		fmt.Println(di.Name(), di.IsDir())
	}

	//! 💡 Writing to file(s)
	// first, creation
	f2,err:=os.Create("example2.txt")
	if err!= nil{
		panic(err)
	}
	defer f2.Close()

	f2.WriteString("Writing to file, example.. ")
	f2.WriteString(" Appending more stuff.. ")

	// Appending in another way
	bytes:= []byte("Hey there.. ")
	f2.Write(bytes)

	//! 💡 Read and write to another file - Streaming Fashion 🎏
	srcFile, err := os.Open("example.txt")
	if err!= nil{
		panic(err)
	}
	defer srcFile.Close()
	destFile, err := os.Create("example3.txt")
	if err!= nil{
		panic(err)
	}
	defer destFile.Close()

	reader:=bufio.NewReader(srcFile)
	writer:= bufio.NewWriter(destFile)

	for {
		b, err:=reader.ReadByte()
		if err!= nil{
			if err.Error() != "EOF"{
				panic(err)
			}
			break
		}

		err1:=writer.WriteByte(b)
		if err1!= nil{
			panic(err1)
		}

	}


	writer.Flush()
	fmt.Println("Written to file successfully! ✅")

	//! 💡 Deleting a file
	err2:=os.Remove("example4.txt")
	if err2!= nil{
		panic(err2)
	}

	fmt.Println("File deleted successfully ☑️")
}