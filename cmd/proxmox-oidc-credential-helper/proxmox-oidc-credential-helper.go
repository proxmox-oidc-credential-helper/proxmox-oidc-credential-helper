package main

import (
	"context"
	"flag"
	"github.com/camaeel/proxmox-oidc-credential-helper/pkg/browser"
	"github.com/camaeel/proxmox-oidc-credential-helper/pkg/callback"
	"github.com/camaeel/proxmox-oidc-credential-helper/pkg/config"
	"github.com/camaeel/proxmox-oidc-credential-helper/pkg/proxmox"
	"log/slog"
	"os"
	"time"
)

func main() {
	cfg := config.Config{}

	flag.IntVar(&cfg.CallbackPort, "callback-port", 8996, "callback port. Default is 8996")
	flag.StringVar(&cfg.CallbackPath, "callback-path", "/oidc/callback", "callback port. Default is 8996")
	flag.StringVar(&cfg.ProxmoxURL, "proxmox-url", "", "Url to proxmox server with protocol and port, i.e. https://proxmox.example.com:8006")
	flag.IntVar(&cfg.TimeoutSeconds, "timeout-url", 180, "Timeout in seconds for whole authentication. By default 180 seconds.")
	flag.BoolVar(&cfg.VerboseLog, "verbose", false, "Verbose logging")

	flag.Parse()

	if cfg.ProxmoxURL == "" {
		slog.Error("Flag proxmox-url is required")
		os.Exit(1)
	}

	if cfg.VerboseLog {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.TimeoutSeconds)*time.Minute)
	defer cancel()

	cancelHttp := callback.StartHttpServer(ctx, cancel, cfg.CallbackPort, cfg.CallbackPath)
	defer cancelHttp()

	redirectUrl, err := proxmox.GetOidcURL(cfg.ProxmoxURL)
	if err != nil {
		slog.Error("Unable to obtain oidc URL from proxmox server", slog.String("error", err.Error()))
		os.Exit(1)
	}
	err = browser.OpenURL(redirectUrl)
	if err != nil {
		slog.Error("Unable to open browser", slog.String("error", err.Error()))
		os.Exit(1)
	}
	<-ctx.Done()
}
