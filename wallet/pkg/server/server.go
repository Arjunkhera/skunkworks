package server

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
	root "wallet/pkg"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server is used to access the mux router
type Server struct {
	router *mux.Router
	config *root.ServerConfig
}

// NewServer creates an instance of a new server
func NewServer(config *root.Config) *Server {
	s := Server{
		router: mux.NewRouter(),
		config: config.Server}

	return &s
}

// Start initiates the server and listens for calls
func (s *Server) Start(flag bool) {

	// poll the server to check if it has started
	go func(flag bool) {
		for {
			time.Sleep(time.Second)

			resp, err := http.Get("http://localhost:8000/login")
			if err != nil {
				log.Println("Failed:", err)
				continue
			}

			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Println("Not OK:", resp.StatusCode)
				continue
			}

			// to check whether the boot file is present or not
			if flag {
				exec.Command("xdg-open", "http://localhost:8000/login").Start()
			} else {
				exec.Command("xdg-open", "http://localhost:8000/signup").Start()
			}
			// Reached this point: server is up and running!
			break
		}
		log.Println("SERVER UP AND RUNNING!")
	}(flag)

	// start the server
	log.Println("Listening on port " + s.config.Port)

	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}

// CreateBootRouter creates BootRouter for handling user related functions
func (s *Server) CreateBootRouter(u root.UserService, d root.DeviceService) {
	NewBootRouter(u, d, s.getSubrouter("/boot"))
}

// CreatePairIdentityRouter creates PairIdentityRouter for handling user related functions
func (s *Server) CreatePairIdentityRouter(pId root.PairIdentityService) {
	NewPairIdentityRouter(pId, s.getSubrouter("/pairId"), s.config.SkynetIP, s.config.ChainIP)
}

// CreateRoutes registers the independent handler functions
func (s *Server) CreateRoutes() {

	s.router.HandleFunc("/signup", displaySignUpHandler).Methods("GET")
	s.router.HandleFunc("/login", displayLoginHandler).Methods("GET")
	s.router.HandleFunc("/view", displayViewHandler)
}
