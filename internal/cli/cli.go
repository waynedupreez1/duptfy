// This is the meat and potatoes which will run commands and send it 
// to ntfy
// Author: Wayne du Preez

package cli

import (
	"fmt"
    "strings"
	"os/exec"
    "net/http"
	"github.com/waynedupreez1/duptfy/internal/flags"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

type priority int

const (
    low priority = iota
    high
)

type CLI struct {
    logger logger.ILogger
    flags *flags.Flags
}

func New(logger logger.ILogger, flags *flags.Flags) *CLI {

    logger.Info("Initializing CLI")

    cli := CLI{
        logger: logger,
        flags: flags,
    }

    return &cli
}

func (t *CLI) Main() {

    t.logger.Info("Running Main")

    _, err := t.runCmd()

    if err != nil {
        t.sendNtfy(high, err.Error())
    } else {
        t.sendNtfy(low, "")
    }
}

func (t *CLI) runCmd() (string, error) {

    t.logger.Info(fmt.Sprintf("Running cmd: %s", t.flags.Cmd))

    out, err := exec.Command("bash", "-c", t.flags.Cmd).CombinedOutput()
    if err != nil {

        var errorStr error
        
        // Some commands do not produce an error output via std or err out
        if len(string(out)) != 0 {
            errorStr = fmt.Errorf("command failed: %s", string(out))
            t.logger.Error(errorStr.Error())
        } else {
            errorStr = fmt.Errorf("")
            t.logger.Error("Command Failed")
        }
        return "", errorStr
    }

    return string(out), nil
}

func (t *CLI) sendNtfy(priority priority, errorMessage string) {

    var comment *strings.Reader    
    pri := ""
    emoji := ""
    
    switch priority {
        case low:
            pri = "low"
            comment = strings.NewReader("Succeeded")
            emoji = "+1"

        case high:
            pri = "high"
            
            if len(errorMessage) != 0 {
                comment = strings.NewReader(errorMessage)
            } else {
                comment = strings.NewReader("Command Failed")
            }
            emoji = "warning"
    }
    
    t.logger.Info(fmt.Sprintf("Sent to Server: %s", t.flags.Server.String()))

    req, err := http.NewRequest("POST", t.flags.Server.String(), comment)
    if err != nil {
        t.logger.Error(fmt.Sprintf("http req failed: %s", err.Error()))
    } else {
        req.Header.Set("Title", t.flags.Message)
        req.Header.Set("Priority", pri)
        req.Header.Set("Tags", emoji)
        http.DefaultClient.Do(req)  
    }
}
