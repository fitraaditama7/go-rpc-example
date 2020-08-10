package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	blogpb "grpc_boilerplate/proto"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Blog Post",
	Long: `Create a new BlogPost on the server thought gRPC.
		
A Blogpost requires an AuthorId, Title and Content`,
	RunE: func(cmd *cobra.Command, args []string) error {
		author, err := cmd.Flags().GetString("author")
		title, err := cmd.Flags().GetString("title")
		content, err := cmd.Flags().GetString("content")

		if err != nil {
			return err
		}

		blog := &blogpb.Blog{
			AuthorId: author,
			Title:    title,
			Content:  content,
		}
		res, err := client.CreateBlog(
			context.TODO(),
			&blogpb.CreateBlogReq{Blog: blog})

		if err != nil {
			return err
		}
		fmt.Printf("Blog Created: %s\n", res.Blog.Id)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("author", "a", "", "Add an Author")
	createCmd.Flags().StringP("title", "t", "", "Add a title for the blog")
	createCmd.Flags().StringP("content", "c", "", "The content for the blog")
	createCmd.MarkFlagRequired("author")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(createCmd)
}
