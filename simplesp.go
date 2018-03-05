package main
import(
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "regexp"
    "strings"
    "flag"
)
var (
    deep = 1
)
func main(){
    url := flag.String("url","","http url for request")
    exp := flag.String("regexp","[a-zA-Z0-9]{32,}","regexp for match")
    flag.IntVar(&deep,"deep",1,"the deep dor search")
    flag.Parse()
    Querycode(*url,*exp,0)
}

func Querycode(url string,exp string, currentDeep int){
    if url=="" || currentDeep > deep {
        return
    }
    fmt.Printf("%v ",url)
    doc, err := goquery.NewDocument(url)
    if err != nil{
        fmt.Println("open url error")
        return
    }
    //match result
    re := regexp.MustCompile(exp)
    fmt.Println(RemoveDuplicatesAndEmpty(re.FindAllString(doc.Text(),-1)))
    urlMap := map[string]int{}
    //match a tag
    doc.Find("a[href]").Each(func(index int,ele *goquery.Selection){
        href,_ := ele.Attr("href")
        // fmt.Println(href)
        if strings.HasPrefix(href,"http"){
            if urlMap[href] != 1 {
                urlMap[href] = 1
                Querycode(href,exp,currentDeep+1)
            }
        }
    })
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string){
    a_len := len(a)
    for i:=0; i < a_len; i++{
        if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
            continue;
        }
        ret = append(ret, a[i])
    }
    return
}
