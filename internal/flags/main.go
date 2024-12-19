// Handles all of the flags returned when the cli is executed
// Author: Wayne du Preez

package flags

import (
    "os"
    "fmt"
    "flag"
    "net/url"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

type Flags struct {
    Server *url.URL //Required
    Cmd string  //Required
    Message string  //Required
}

func Get(logger logger.ILogger) *Flags {

    logger.Info("Get Flags")

    var urlString,cmd, message string
    flag.StringVar(&urlString, "s", "", "Required. ntfy endpoint URL ie. http://example.com/backup.")
    flag.StringVar(&cmd, "c", "", "Required. Bash command to run ie. 'ls -als | grep blah'.")
    flag.StringVar(&message, "m", "", "Required. Message when sent to ntfy ie. 'Local Rsnapshot backup'.")

    flag.Parse()

    if urlString == "" || cmd == "" || message == "" {
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
    logger.Info(fmt.Sprintf("Command that will run: %s", cmd))
    
    flags := Flags{
        Server: serverURL, 
        Cmd: cmd,
        Message: message,
    }

    return &flags
}
