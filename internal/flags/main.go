package flags

import (
    "os"
    "fmt"
    "flag"
    "net/url"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

type Flags struct {
    Server *url.URL
    Cmd string
}

func Get(logger logger.ILogger) *Flags {

    logger.Info("Get Flags")

    var urlString,cmd string
    flag.StringVar(&urlString, "s", "", "ntfy endpoint URL ie. http://example.com/backup")
    flag.StringVar(&cmd, "c", "", "Bash command to run ie. 'ls -als | grep blah'")

    flag.Parse()

    if urlString == "" || cmd == "" {
        flag.Usage()
        os.Exit(1)
    }

    serverURL, err := url.Parse(urlString)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

    if !serverURL.IsAbs(){
        logger.Error(fmt.Sprintf("Invalid URL: %s", serverURL))
        os.Exit(1)
    }

    logger.Info(fmt.Sprintf("ntfy Server URL: %s", serverURL))
    logger.Info(fmt.Sprintf("Command to run: %s", cmd))
    
    f := Flags{
        Server: serverURL, 
        Cmd: cmd,
    }

    return &f
}
