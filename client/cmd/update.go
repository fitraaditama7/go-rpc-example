package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	blogpb "grpc_boilerplate/proto"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Find Blog Post By Id",
	Long: `Find a Blog Post By MongoDB Unique ID.

If no blogpost is found for it's ID it will return a 'Not Found' error'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		author, err := cmd.Flags().GetString("author")
		title, err := cmd.Flags().GetString("title")
		content, err := cmd.Flags().GetString("content")

		req := &blogpb.UpdateBlogReq{
			Blog: &blogpb.Blog{
				Id:       id,
				AuthorId: author,
				Title:    title,
				Content:  content,
			},
		}

		res, err := client.UpdateBlog(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "The id of the blog")
	updateCmd.Flags().StringP("author", "a", "", "The Author of the blog post")
	updateCmd.Flags().StringP("title", "t", "", "A title for the blog post")
	updateCmd.Flags().StringP("content", "c", "", "The content for the blog post")
	updateCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(updateCmd)
}
