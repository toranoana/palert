package main

import (
    "flag"
    "fmt"
    "github.com/mitchellh/go-ps"
    "net/http"
    "net/url"
    "os"
    "strconv"
)


// 引数1 探す プロセス名
// 引数2 なかった場合のメッセージ
// 引数3 alert slack URL

// Go言語 プロセスの情報出力
// https://blog.y-yuki.net/entry/2018/08/02/004042


func main() {
    flag.Parse()

    searchProcessName := flag.Arg(0)
    message := flag.Arg(1)
    slackUrl := flag.Arg(2)

    processes, err := ps.Processes()

    if err != nil {
        os.Exit(1)
    }

    for _, p := range processes {
        // fmt.Printf("%d : %s\n", i, p.Executable())

        if searchProcessName == p.Executable() {
            fmt.Println("HIT PID=" + strconv.Itoa(p.Pid()))
            os.Exit(0)
        }

    }

    fmt.Println(message)
    PostSlack(slackUrl, message)
    os.Exit(1)
}

func PostSlack(slackUrl string, message string) {

    response, err := http.PostForm(slackUrl,
        url.Values{
            //"payload": {`{"text:"` + message + `"}`},
            "payload": { `{"text": "` + message + `"}` },
        })

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("status:", response.Status)
}

