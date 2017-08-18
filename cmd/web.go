package cmd

import (
	"fmt"
	"net/http"

	"github.com/daniellowtw/xavier/api"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

func NewWebCmd() *cobra.Command {
	var Port int
	cmd := &cobra.Command{
		Use: "web",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := newServiceFromCmd(cmd)
			if err != nil {
				return err
			}
			r := mux.NewRouter()
			subRouter := r.PathPrefix("/api").Subrouter()
			r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(assetFS())))
			//r.Handle("/", http.FileServer(assetFS()))
			api.Register(s, subRouter)
			r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(r.URL.Path)
				w.WriteHeader(404)
			})
			fmt.Printf("Server running on port %d\n", Port)
			return http.ListenAndServe(fmt.Sprintf(":%d", Port), r)
		},
	}
	cmd.Flags().IntVarP(&Port, "port", "p", 9090, "port to run the API server on")
	return cmd
}
