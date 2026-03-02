package main

func main() {

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	//error
	// 	panic(err)
	// }

	// files, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name:", files.Name())
	// fmt.Println("file or folder:", files.IsDir())
	// fmt.Println("file permission:", files.Mode())
	// fmt.Println("file mode time:", files.ModTime())
	// fmt.Println("file size:", files.Size()) //size in bytes

	//read file

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	//error
	// 	panic(err)
	// }

	// defer f.Close()

	// buffer := make([]byte, 12)

	// d, err := f.Read(buffer)

	// if err != nil {
	// 	panic(err)
	// }

	// // for i := 0; i < len(buffer); i++ {

	// // 	fmt.Println("data", d, string(buffer[i]))
	// // }

	// println("data read:", string(buffer))
	// println("bytes read:", d)

	//direct file reading
	// f, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// println("file content:", string(f))

	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// FileInfo, err := dir.ReadDir(1)

	// for _, file := range FileInfo {
	// 	fmt.Println("file name:", file.Name())
	// }

	//create file

	// bytes := []byte("hii krish\nNice to meet you")
	// f.write(bytes)

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("hii krish\n")
	// f.WriteString("Nice to meet you")

	// // Read file
	// data, err := os.ReadFile("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// content := string(data)

	// updated := strings.ReplaceAll(content, "krish", "krishan")

	// err = os.WriteFile("example2.txt", []byte(updated), 0644)
	// if err != nil {
	// 	panic(err)
	// }

	//read and write  to another file(streaming fashion)

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destfile, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer destfile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destfile)

	// for {
	// 	b, err := reader.ReadByte()

	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			break
	// 		}
	// 		panic(err)
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("file copied successfully")

	//delete file

	// sourceFile, err := os.Open("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// err = sourceFile.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.Remove("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file deleted successfully")

}

// Disk file
//    ↓
// os.ReadFile()
//    ↓
// []byte
//    ↓
// string conversion
//    ↓
// ReplaceAll()
//    ↓
// new string
//    ↓
// os.WriteFile()
//    ↓
// Disk updated
