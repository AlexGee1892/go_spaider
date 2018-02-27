package main
import "fmt"
import "strconv"
import "net/http"
import "io/ioutil"
import "log"
import "os"

func httpGet(url string)(content string, statusCode int){
    resp, err := http.Get(url)
    if err != nil{
        log.Println(err)
        content = ""
        statusCode = -100
        return
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        log.Println(err)
        content = ""
        statusCode = resp.StatusCode
        return
    }
    content = string(data)
    statusCode = resp.StatusCode
    return
}
func Tieba(begin int, end int){
    var pn int
    var url string
    for page := begin; page <= end; page++{
        fmt.Println("正在爬取第", page , "页")
        pn = (page - 1) * 50
        url = "https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn=" + strconv.Itoa(pn)
        content, rcode := httpGet(url)
        if rcode != 200{
            fmt.Println("statusCode = ", rcode)
            continue
        }

        filename := strconv.Itoa(page) + ".html"
        fout,err := os.Create(filename)
        if err != nil{
            log.Println(err)
            return

        }
        defer fout.Close()
        fout.WriteString(content)
        
    }
    
}

func main(){
    var begin string
    var end string

    fmt.Println("请输入要爬取的起始页码：")
    fmt.Scanf("%s\n", &begin)
    fmt.Println("请输入要爬取的结束页码：")
    fmt.Scanf("%s\n", &end)
    b, _ := strconv.Atoi(begin)
    e, _ := strconv.Atoi(end)

    Tieba(b,e)
}
