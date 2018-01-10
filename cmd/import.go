package cmd

import (
	"fmt"

	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/gilliek/go-opml/opml"
	"github.com/spf13/cobra"
)

var (
	ImportFeedSourceFromFileCmd = &cobra.Command{
		Use: "import",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("first arg is file name")
			}
			in, err := opml.NewOPMLFromFile(args[0])
			if err != nil {
				return err
			}
			importCount := 0
			logErr := func(err error) {
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				importCount++
			}
			s, err := service.NewServiceFromCmd(cmd)
			if err != nil {
				return err
			}
			for _, out := range flattenOutlines(in.Outlines()) {
				switch {
				case out.XMLURL != "":
					fmt.Println("XML", out.XMLURL)
					logErr(s.AddFeed(out.XMLURL))
				case out.URL != "":
					fmt.Println("URL", out.URL)
					logErr(s.AddFeed(out.URL))
				case out.HTMLURL != "":
					fmt.Println("HTML", out.HTMLURL)
					logErr(s.AddFeed(out.HTMLURL))
				}
			}
			fmt.Printf("Imported %d feeds\n", importCount)
			return nil
		},
	}
)

func flattenOutlines(outline []opml.Outline) []opml.Outline {
	var res []opml.Outline
	for _, o := range outline {
		res = append(res, o)
		res = append(res, flattenOutlines(o.Outlines)...)
	}
	return res
}
