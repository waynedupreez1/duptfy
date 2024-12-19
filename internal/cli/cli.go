package cli

import (
	"fmt"
    "strings"
	"os/exec"
    "net/http"
	"github.com/waynedupreez1/duptfy/internal/flags"
	"github.com/waynedupreez1/duptfy/internal/logger"
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

func (t *CLI) RunCmd() (string, error) {

    t.logger.Info(fmt.Sprintf("Running cmd: %s", t.flags.Cmd))

    out, err := exec.Command("bash", "-c", t.flags.Cmd).Output()
    if err != nil {
        
        errorStr := fmt.Errorf("command failed")
        t.logger.Error(errorStr.Error())
        return "", errorStr
    }

    return string(out), nil
}

func (t *CLI) SendNtfy(priority string) {

    t.logger.Info(fmt.Sprintf("Sent to Server: %s", t.flags.Server.String()))

    req, _ := http.NewRequest("POST", t.flags.Server.String(), strings.NewReader("Remote access to phils-laptop detected. Act right away."))
    req.Header.Set("Title", "Unauthorized access detected")
    req.Header.Set("Priority", "urgent")
    req.Header.Set("Tags", "warning,skull")
    http.DefaultClient.Do(req)
}
