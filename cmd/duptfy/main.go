package main

import (
	"fmt"

	"github.com/waynedupreez1/duptfy/internal/flags"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

func main() {

    log := logger.New(logger.Info)

    flags := flags.New(log)

    log.Info(fmt.Sprintf("blah: %+v", flags))
}
