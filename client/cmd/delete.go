package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	blogpb "grpc_boilerplate/proto"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete blog post by it's id",
	Long: `Delete a blog post by it's MongoDB Identifier.
If no blog post is found for the ID it will return 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req := &blogpb.DeleteBlogReq{
			Id: id,
		}
		_, err = client.DeleteBlog(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Println("Succesfully deleted the blog with id ", id)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringP("id", "i", "", "id of the blog")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}