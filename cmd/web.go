package cmd

import (
	"fmt"
	"github.com/daniellowtw/xavier/api"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	Port   = 9090
	WebCmd = &cobra.Command{
		Use: "web",
		RunE: func(cmd *cobra.Command, args []string) error {
			r := mux.NewRouter()
			subRouter := r.PathPrefix("/_api").Subrouter()
			s := api.Service{
				StoreEngine: e,
			}
			s.Register(subRouter)
			r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(r.URL.Path)
			})
			fmt.Printf("Server running on port %d\n", Port)
			return http.ListenAndServe(fmt.Sprintf(":%d", Port), r)
		},
	}
)
