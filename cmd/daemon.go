package cmd

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/grid/contracts/eth"
	"github.com/gridprotocol/dumper/database"
	"github.com/gridprotocol/dumper/dumper"
	comm "github.com/gridprotocol/platform-v2/common"
	"github.com/gridprotocol/platform-v2/lib/config"
	"github.com/gridprotocol/platform-v2/logs"
	"github.com/gridprotocol/platform-v2/server"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

var logger = logs.Logger("daemon")

var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "platform daemon",
	Subcommands: []*cli.Command{
		runCmd,
		stopCmd,
	},
}

// run daemon
var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "chain",
			Usage: "input chain name, e.g.(dev)",
			Value: "dev",
		},
	},
	Action: func(ctx *cli.Context) error {
		chain := ctx.String("chain")

		// parse config file
		err := config.InitConfig()
		if err != nil {
			logger.DPanicf("failed to init the config: %v", err)
		}
		ep := config.GetConfig().Http.Listen
		logger.Info("server endpoint:", ep)

		var chain_ep string

		// select contracts addresses for each chain
		switch chain {
		case "local":
			chain_ep = eth.Ganache
			comm.Contracts = comm.LocalContracts.Contracts
		case "dev":
			chain_ep = eth.DevChain
			comm.Contracts = comm.DevContracts.Contracts
		case "sepo":
			chain_ep = eth.Sepolia
			comm.Contracts = comm.SepoContracts.Contracts
		}

		logger.Infof("chain selected:%s, chain endpoint:%s\n", chain, chain_ep)
		logger.Infof("contract addresses:", comm.Contracts)

		// listen endpoint and chain endpoint
		opts := server.ServerOption{
			Endpoint:       ep,
			Chain_Endpoint: chain_ep,
		}

		// init db
		logger.Info("init db..")
		err = database.InitDatabase("./grid")
		if err != nil {
			return err
		}

		// contract address
		registryAddress := common.HexToAddress(comm.Contracts.Registry)
		marketAddress := common.HexToAddress(comm.Contracts.Market)

		logger.Info("init dumper..")
		logger.Info(chain, registryAddress, marketAddress)

		// init dumper
		dumper, err := dumper.NewGRIDDumper(chain_ep, registryAddress, marketAddress)
		if err != nil {
			return err
		}

		logger.Info("first dump..")
		err = dumper.DumpGRID()
		if err != nil {
			return err
		}

		// sync chain for db
		logger.Info("sync db with block chain..")
		go dumper.SubscribeGRID(ctx.Context)

		// create http server with routes
		srv := server.NewServer(opts)

		// start server
		go func() {
			logger.Info("start listen..")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		}()

		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Info("Shutting down server...")

		cctx, cancel := context.WithTimeout(ctx.Context, 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(cctx); err != nil {
			logger.Fatal("Server forced to shutdown: ", err)
		}

		logger.Info("Server exiting")
		return nil
	},
}

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "stop server",
	Action: func(_ *cli.Context) error {
		pidpath, err := homedir.Expand("./")
		if err != nil {
			return nil
		}

		pd, _ := os.ReadFile(path.Join(pidpath, "pid"))

		err = kill(string(pd))
		if err != nil {
			return err
		}
		logger.Info("gateway gracefully exit...")

		return nil
	},
}

func kill(pid string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("kill", "-15", pid).Run()
	case "windows":
		return exec.Command("taskkill", "/F", "/T", "/PID", pid).Run()
	default:
		return fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}
}

// func getEndpointByChain(chain string) string {
// 	switch chain {
// 	case "local":
// 		return eth.Ganache
// 	case "dev":
// 		return "https://devchain.metamemo.one:8501"
// 		//return "http://10.10.100.82:8201"

// 		// case "test":
// 		// 	return "https://testchain.metamemo.one:24180"
// 		// case "product":
// 		// 	return "https://chain.metamemo.one:8501"
// 		// case "goerli":
// 		// 	return "https://eth-goerli.g.alchemy.com/v2/Bn3AbuwyuTWanFLJiflS-dc23r1Re_Af"
// 	}
// 	return "https://devchain.metamemo.one:8501"
// }
