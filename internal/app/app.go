package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/nurlan42/todo/cfg"
	"github.com/nurlan42/todo/internal/repository"
	"github.com/nurlan42/todo/internal/usecase"
	"github.com/nurlan42/todo/pkg/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	delivery "github.com/nurlan42/todo/internal/delivery/http"
	log "github.com/sirupsen/logrus"
)

//github.com/nurlan42/todo1

func Run(cfg *cfg.Config) error {
	sqlDB, err := db.Connect(cfg.DB.TODO)
	if err != nil {
		return err
	}

	todoRepo := repository.New(sqlDB, cfg)

	todoUsecase := usecase.New(todoRepo)

	todoHandler := delivery.New(todoUsecase)

	handler := todoHandler.Init()

	httpSrv := http.Server{
		Addr:         ":" + cfg.HTTP.Port,
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.ReadTimeout * time.Second,
		WriteTimeout: cfg.HTTP.WriteTimeout * time.Second,
	}

	ch := make(chan os.Signal, 1)

	go func() {
		err := httpSrv.ListenAndServe()

		fmt.Printf("\nserver is being shutdown\n")

		if !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("failed to stop server: %s\n", err.Error())
			ch <- os.Kill
		}
	}()

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop srv.Shutdown(): %s\n", err.Error())
	}

	//if err := sqlDB.Close(); err != nil {
	//	return fmt.Errorf("failed to stop sqlDB: %s\n", err.Error())
	//}

	return nil
}
