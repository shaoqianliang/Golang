package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//启动入口必须是package main,并包main函数

//包在$GOPATH/pkg目录下

//func main() {
//s1 := "100"
//i1, err := strconv.Atoi(s1)
//if err != nil {
//	fmt.Println("can't covert to int format")
//} else {
//	fmt.Printf("change successful, and the value is %d", i1)
//}
//file, err := os.Open("./a.txt")
//if err != nil {
//	fmt.Println("open file failed")
//	return
//}
//defer file.Close()
////循环读取
//var content []byte
//var tmp = make([]byte, 12)
//for {
//	n, err := file.Read(tmp)
//	if err == io.EOF {
//		fmt.Println("read entire file ", err)
//		break
//	}
//	if err != nil {
//		fmt.Println("read file failed")
//		return
//	}
//	content = append(content, tmp[:n]...)
//}
//
////fmt.Printf("读取了%d字节数据\n数据内容", n)
//fmt.Println(string(content))
//fmt.Println(time.Now()) //本地时间戳

//时区偏移
//secondsEastUTC := int((8*time.Hour).Seconds())
//time.FixedZone("BeiJing Time", secondsEastUTC)
//time.LoadLocation("Ameraca/New_York")
//定时器
//	ticker := time.Tick(time.Second)
//	for i := range ticker {
//		fmt.Println(i)
//	}
//}

//	func funcB() {
//		defer func() {
//			err := recover()
//			//如果程序出出现了panic错误,可以通过recover恢复过来,panic中的内容不执行
//			if err != nil {
//				fmt.Println("recover in B")
//			}
//		}()
//		panic("panic in B")
//	}
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found album"})
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/album", postAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.Run("192.168.0.105:8080")
}
