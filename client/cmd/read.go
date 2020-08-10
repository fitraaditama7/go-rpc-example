package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	blogpb "grpc_boilerplate/proto"
)

var readCmd = &cobra.Command{
	Use:  "read",
	Short: "Find Blog Post By Id",
	Long: `Find a Blog Post By MongoDB Unique ID.

If no blogpost is found for it's ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req := &blogpb.ReadBlogReq{Id: id}
		res, err := client.ReadBlog(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	readCmd.Flags().StringP("id", "i", "", "The id of the blog")
	readCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(readCmd)
}
