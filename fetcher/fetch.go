package fetcher

import (
	"net/http"
	"fmt"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
)
//原生的方式
func BaseFetch(url string )([]byte,error){


	resp,err:= http.Get("https://book.douban.com/")

	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error status code:%d",resp.StatusCode)
	}

	bodyReader:= bufio.NewReader(resp.Body)
	e:= DeterminEncoding(bodyReader)

	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())


	return ioutil.ReadAll(utf8Reader)


}





 //模拟浏览器访问
func Fetch(url string )([]byte,error){




	client:=&http.Client{}
	req,err:= http.NewRequest("GET",url,nil)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",url)
	}


	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")

	resp,err:= client.Do(req)


	bodyReader:= bufio.NewReader(resp.Body)
	e:= DeterminEncoding(bodyReader)

	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())


	return ioutil.ReadAll(utf8Reader)


}


func DeterminEncoding(r * bufio.Reader) encoding.Encoding{

	bytes,err:= r.Peek(1024)

	if err!=nil{
		log.Printf("fetch error:%v",err)
		return unicode.UTF8
	}

	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}


