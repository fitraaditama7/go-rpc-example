package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	blogpb "grpc_boilerplate/proto"
	"io"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all blog posts",
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &blogpb.ListBlogsReq{}
		stream, err := client.ListBlogs(context.Background(), req)
		if err != nil {
			return err
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			fmt.Println(res.GetBlog())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}