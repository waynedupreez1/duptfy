package flags

import (
    "os"
    "fmt"
    "flag"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

type Flags struct {
    server string
    cmd string
}

func New(logger logger.ILogger) Flags {
    
    var server,cmd string
    flag.StringVar(&server, "s", "", "ntfy endpoint URL ie. http://example.com/backup")
    flag.StringVar(&cmd, "c", "", "Bash command to run ie. 'ls -als | grep blah'")

    flag.Parse()

    if server == "" || cmd == "" {
        logger.Error("Missing command line variables")
        flag.Usage()
        os.Exit(1)
    }
    
    logger.Info(fmt.Sprintf("ntfy Server: %s", server))
    logger.Info(fmt.Sprintf("Command to run: %s", cmd))
    
    f := Flags {server: server, cmd: cmd}

    return f
}
