package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Third party pkgs
import (
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	serviceAddress string
	initRange      string
	endRange       string
	group          int
	// tls
	tlsCertificate string
	tlsKey         string
	// storage data
	stack *processingStack
)

var RunCmd = &cobra.Command{
	Use:     "server",
	Short:   "Run server",
	Long:    "Run server for parallel network scanning.",
	Example: "server -a 0.0.0.0:8080",
}

func init() {
	RunCmd.PersistentFlags().StringVarP(&serviceAddress, "address", "a", "0.0.0.0:443", "the service address")
	RunCmd.PersistentFlags().StringVarP(&initRange, "init", "i", "", "the intial range")
	RunCmd.PersistentFlags().StringVarP(&endRange, "end", "e", "", "the end")
	RunCmd.PersistentFlags().IntVarP(&group, "group", "g", 1000, "group size")
	RunCmd.PersistentFlags().StringVarP(&tlsCertificate, "certificate", "", "", "TLS certificate")
	RunCmd.PersistentFlags().StringVarP(&tlsKey, "key", "", "", "TLS key")
	RunCmd.RunE = run
}

type StandardResponse struct {
	Status string `json:"status"` // status string;
	Error  string `json:"error"`  // error description.
}

// Response status messages
const (
	AckResponse = "ACK" // Acknowledge message;
	NakResponse = "NAK" // Not acknowledged message.
)

// Ping function to verify if the service is on
// or not.
func getPing(w http.ResponseWriter, r *http.Request) {
	/* return value */
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&StandardResponse{
		Status: AckResponse,
	})
}

type Payload struct {
	Payload string `json:"payload"`
}

// riseError rises an error returning a standard error
// message.
func riseError(status int, msg string, w http.ResponseWriter, ipa string) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(
		StandardResponse{
			NakResponse,
			msg,
		})
	if err != nil {
		panic(err)
	}
}

func getPayload(w http.ResponseWriter, r *http.Request) {
	if stack == nil {
		riseError(http.StatusInternalServerError, "nil pointer", w, r.RemoteAddr)
		return
	}
	// get actual payload
	payload, err := stack.GetNext()
	if err != nil {
		riseError(http.StatusNotFound, err.Error(), w, r.RemoteAddr)
		return
	}

	// create result structure
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Payload{
		Payload: payload,
	})
}

func run(cmd *cobra.Command, args []string) error {
	// generate subnets
	var err error
	subnets, err := GenerateSubnet(initRange, group, endRange)
	if err != nil {
		return err
	}
	stack = newProcessingStack(subnets)

	// create router
	route := mux.NewRouter()
	// get payload
	route.HandleFunc("/v1/payload", getPayload).Methods("GET")
	// utility routes
	route.HandleFunc("/v1/ping", getPing).Methods("GET")
	// root routes
	http.Handle("/", route)

	// init http or https connection
	if tlsCertificate != "" &&
		tlsKey != "" {
		log.Printf("Starting listening with TLS on address %s.\n", serviceAddress)
		// set up SSL/TLS
		err = http.ListenAndServeTLS(
			serviceAddress,
			tlsCertificate,
			tlsKey,
			nil)
		if err != nil {
			return fmt.Errorf("https unable to listen and serve on address: %s cause error: %s", serviceAddress, err.Error())
		}
	} else {
		log.Printf("Starting listening on address %s (no SSL/TLS this can produce security risks).\n", serviceAddress)
		// plain https
		err = http.ListenAndServe(serviceAddress, nil)
		if err != nil {
			return fmt.Errorf("http unable to listen and serve on address: %s cause error: %s", serviceAddress, err.Error())
		}
	}
	return nil
}

func Execute() error {
	_, err := RunCmd.ExecuteC()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := Execute()
	if err != nil {
		log.Printf("%s.\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
