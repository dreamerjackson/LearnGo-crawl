package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)
//原生的方式
func BaseFetch(url string )([]byte,error){


	resp,err:= http.Get(url)

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




var ratelimit = time.Tick(1*time.Millisecond)

func FetchVideo(url string )(io.ReadCloser,error){

	<-ratelimit

	client:=&http.Client{}

	req,err:= http.NewRequest("GET",url,nil)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",url)
	}


	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	resp,err:= client.Do(req)

	if err !=nil || resp==nil{
		fmt.Println("err:",err)
		return  nil,err
	}
	if resp.ContentLength >  12 *1024*1024{
		return nil,errors.New("too big")
	}
	return resp.Body,nil
}

//模拟浏览器访问
func Fetch(url string )([]byte,error){

	<-ratelimit


	client:=&http.Client{}

	req,err:= http.NewRequest("GET",url,nil)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",url)
	}


	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	resp,err:= client.Do(req)
	if err !=nil {
		fmt.Println("err:",err)
		return  nil,err
	}

	bodyReader:= bufio.NewReader(resp.Body)
	e:= DeterminEncoding(bodyReader)

	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())


	return ioutil.ReadAll(utf8Reader)
}

//http代理访问访问
func ProxyFetch(weburl string )([]byte,error){

	<-ratelimit


	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:51816")//根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	req,err:= http.NewRequest("POST",weburl,nil)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",weburl)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	resp,err:= client.Do(req)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",weburl)
	}
	bodyReader:= bufio.NewReader(resp.Body)
	e:= DeterminEncoding(bodyReader)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",weburl)
	}
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


