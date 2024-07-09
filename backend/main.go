package main

import (
	"log"
	"net/http"
	"task-manager-backend/routers"
	"task-manager-backend/utils"

	"github.com/joho/godotenv"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK int = iota
	exitError
)

func main() {
	// Create separate main instead of doing the actual code here
	// since os.Exit can not handle `defer`. DON'T call `os.Exit` in the any other place.
	// os.Exit(realMain(os.Args))
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitDB()
	router := routers.InitRouter()
	//log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":443", router))
}

// func realMain(args []string) int {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	utils.InitDB()
// 	// router := routers.InitRouter()
// 	// log.Println("Listening on Port 8080")
// 	// log.Fatal(http.ListenAndServe(":8080", router))
// 	userHandler := handlers.NewUserProfileHandler()
// 	httpServer := http.NewServer(
// 		userHandler,
// 	)
// 	log.Println("service initialization successful")

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	wg, ctx := errgroup.WithContext(ctx)
// 	addr := fmt.Sprintf(":%s", "8080")
// 	ln, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	async := func(name string, f func() error) {
// 		wg.Go(func() error {
// 			err := f()
// 			log.Println("process stopped", zap.String("name", name), zap.Error(err))
// 			return err
// 		})
// 	}

// 	async("http server", func() error { return httpServer.Serve(ln) })

// 	// Waiting for SIGTERM or Interrupt signal. If server receives them,
// 	// http server and any other services running will shutdown gracefully.
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
// 	select {
// 	case <-sigCh:
// 		log.Println("received SIGTERM, exiting server gracefully")
// 	case <-ctx.Done():
// 	}

// 	// gracefully shutdown http server
// 	if err := httpServer.GracefulStop(context.Background()); err != nil {
// 		log.Println("failed to close http server", zap.Error(err))
// 	}

// 	cancel()
// 	if err := wg.Wait(); err != nil {
// 		log.Println("unhandled error received", zap.Error(err))
// 		return exitError
// 	}

// 	log.Println("exiting the service")
// 	return exitOK

// }
