package x2gobroker

import (
	"context"
	"sync"
)

type App struct {
	wg *sync.WaitGroup
}

func (app *App) NewApp() *App {
	return nil
}

type HTTPServer interface {
	Start()
	Stop()
}

func (app *App) NewHTTPServer(ctx context.Context, wg *sync.WaitGroup, srv *HTTPServer) {

}
