package main

import (
	//"os"
	//"fmt"
	//"strings"
	//"os/exec"
	log "github.com/kdar/factorlog"
	"flag"
	"os"
	"time"
	"math/rand"
	"net/http"
	"fmt"
	"io/ioutil"
	"sync"
	"runtime"
)

var (
	logFlag        = flag.String("log", "", "set log path")
	logger *log.FactorLog
	maxGoroutines int = 16
	report string = "http://api.1sapp.com/readtimer/report"
	intPointReward string = "http://api.1sapp.com/mission/intPointReward"
)

func main () {
	logger = SetGlobalLogger(*logFlag)

	logger.Info("CPU: ", runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU())

	goroutineFunc()

	time.Sleep(time.Duration(30)*time.Second)

	goroutineFunc()

}

func goroutineFunc(){
	var wg sync.WaitGroup
	limiter := make(chan struct{}, maxGoroutines)
	times := 0

	for i := 0; i < maxGoroutines; i++ {
		limiter <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
				<-limiter
			}()

			requestReport()

			times ++
			logger.Info("testing########: ", times)
		}()
	}
	wg.Wait()

	times = 0
}

func requestReport(){
	paramsList := "ODVBNjU3RUY5Q0FBMDg2QjgyOEE5MzFDM0EwQkQwMDEuY0dGeVlXMGZOR1l3WlRjMk5UY3RZalExTWkwME1HTmpMV0UxTVRjdE1HSm1aVFk0T0dWbE5UUXpIblpsY25OcGIyNGZNaDV3YkdGMFptOXliUjloYm1SeWIybGsuWV1dVYUgODTfJ9HLeW7usd0yadVzOnepJ6lfw5kILTLruj8fe8CrDo0xit3IVYJnVHMorA358cY781W8ne%2FUXAreJpqVnuaL2HLyZraFXR6HeeAEtkBUrQqivgM3kem8cOwgnvdBuDSTX0oHO0o5tdkqyQgKfXIqa1gbNWo8EkNe4bwtfHSJrF6k%2Ftbw%2F0ogxg51EnHiewYInrR66XC32DjPA66C1yN4tUXMYi73LI5B2wng%2B9AIhw9EDVYsVoL0MmFl9Kp5nGY%2F2kG%2BYMr1CaHijWIINB79WRnHdKe7igerthelTFxjmUz5r1OggZKSfxscIu2hsSn%2FWS9ZP6Iop5id3BepKtd%2FRK6r90gQWimgSFVk2Xna2NSmP2m8x5QaKIWKULRL%2BvdRFYctsTJFNOJKKXL0iIC0qX05AkdROjZ8YkjP4Yh6Q4lCtSaW6D5Tvk0sJ83Cu2uH6qVrQCMg5LVvzsQ2I8lpmmU67GhXOlptf6yI7CWh0poZgtzh4Fde03IZwrJtpdQ%2BIfzkkvXGZdF1XN84%2FqAA6NWzUYQZo6cCmznuhtv8q%2FZLPQJ7Uoj88gZpR%2F6M6RYmaXwALc6XnYSIqWMk17HsC%2BoMN1pngW4xkLqMzUgJRcP7OE%2F4u9pWhp6qp%2BFKc06f8S%2FjDeSS4Kmyl2zb%2F%2F8yWVQfT2OZSzVRkOthesj50H%2BWzQSNylAwrWUSDWmxjJvPGPziHrhosBtk%2BxnUVrtRtXpGFGwAJtfez67Ln2BHnu6rh66cW2GtHQ%3D%3D"

	reqUrl := report + "?qdata=" + paramsList
//logger.Info(reqUrl, "======")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Host", "api.1sapp.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "br, gzip, deflate")
	req.Header.Set("Accept-Language", "zh-Hans;q=1")

	cookies := "aliyungf_tc=AQAAAF9f3j0RUwEAd8hbLwNUIuTvGGLu; qkV2=724966bf13638e681a2d60ac983338df"
	req.Header.Set("Cookie", cookies)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(body)

	fmt.Println(len(str))

	return
}


func randomUserAgent() string {
	var userAgent = []string{
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	}
	rand.Seed(time.Now().UnixNano())
	return userAgent[rand.Intn(len(userAgent))]
}

func SetGlobalLogger(logPath string) *log.FactorLog {
	sfmt := `%{Color "red:white" "CRITICAL"}%{Color "red" "ERROR"}%{Color "yellow" "WARN"}%{Color "green" "INFO"}%{Color "cyan" "DEBUG"}%{Color "blue" "TRACE"}[%{Date} %{Time}] [%{SEVERITY}:%{ShortFile}:%{Line}] %{Message}%{Color "reset"}`
	logger := log.New(os.Stdout, log.NewStdFormatter(sfmt))
	if len(logPath) > 0 {
		logf, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
		if err != nil {
			return logger
		}
		logger = log.New(logf, log.NewStdFormatter(sfmt))
	}
	logger.SetSeverities(log.INFO | log.WARN | log.ERROR | log.FATAL | log.CRITICAL)
	return logger
}