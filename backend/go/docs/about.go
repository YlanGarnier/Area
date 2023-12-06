package docs

import (
	"time"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	c.JSON(200, gin.H{
		"client": gin.H{
			"host": c.ClientIP(),
		},
		"server": gin.H{
			"current_time": time.Now().Unix(),
			"services": []gin.H{
				{
					"name": "github",
					"reactions": []gin.H{
						{
							"name":        "create_issue",
							"description": "Create an issue on a repository",
						},
						{
							"name":        "create_repo",
							"description": "Create a repository",
						},
						{
							"name":        "create_gists",
							"description": "Create a gist",
						},
						{
							"name":        "add_topics",
							"description": "Add topics to a repository",
						},
						{
							"name":        "sync_branch",
							"description": "Sync a branch with the remote upstream",
						},
					},
				},
				{
					"name": "discord",
					"actions": []gin.H{
						{
							"name":        "new_guild",
							"description": "The user join a new guild",
						},
					},
				},
				{
					"name": "ethereum",
					"actions": []gin.H{
						{
							"name":        "watch_transaction",
							"description": "The provided address is the sender or the receiver of a transaction",
						},
						{
							"name":        "watch_events",
							"description": "The provided address is the sender or the receiver of a transaction",
						},
					},
				},
				{
					"name": "http",
					"actions": []gin.H{
						{
							"name":        "watcher",
							"description": "The provided url is called and compared to the expected value",
						},
					},
				},
				{
					"name": "google_gmail",
					"actions": []gin.H{
						{
							"name":        "new_email",
							"description": "A new email is received in the inbox. ",
						},
						{
							"name":        "new_email_from",
							"description": "A new email is received in the inbox from the provided email address. ",
						},
						{
							"name":        "new_email_at_date",
							"description": "A new email is received in the inbox at the provided date. ",
						},
						{
							"name":        "new_email_from_and_at_date",
							"description": "A new email is received in the inbox from the provided email address and at the provided date. ",
						},
						{
							"name":        "new_draft",
							"description": "A new draft is created in the inbox. ",
						},
						{
							"name":        "new_draft_to",
							"description": "A new draft is created in the inbox to the provided email address. ",
						},
						{
							"name":        "new_draft_at_date",
							"description": "A new draft is created in the inbox at the provided date. ",
						},
						{
							"name":        "new_draft_to_and_at_date",
							"description": "A new draft is created in the inbox to the provided email address and at the provided date. ",
						},
						{
							"name":        "new_label",
							"description": "A new label is created in the inbox. ",
						},
						{
							"name":        "new_label_with_name",
							"description": "A new label is created in the inbox with the provided name. ",
						},
						{
							"name":        "new_email_in_label",
							"description": "A new email is received in the inbox and is labeled with the provided label name. ",
						},
					},
				},
				{
					"name": "spotify",
					"actions": []gin.H{
						{
							"name":        "watch_artist",
							"description": "The user is listening to the provided artist",
						},
						{
							"name":        "watch_song",
							"description": "The user is listening to the provided song",
						},
					},
				},
				{
					"name": "twitter",
					"reactions": []gin.H{
						{
							"name":        "post_tweet",
							"description": "Post a tweet",
						},
						{
							"name":        "post_tweet_with_content",
							"description": "Post a tweet with the provided content",
						},
						{
							"name":        "post_tweet_with_poll",
							"description": "Post a tweet with a poll",
						},
						{
							"name":        "post_tweet_with_content_and_poll",
							"description": "Post a tweet with the provided content and a poll",
						},
					},
				},
				{
					"name": "twitch",
					"reactions": []gin.H{
						{
							"name":        "send_default_message",
							"description": "Send a default message.",
						},
						{
							"name":        "send_message",
							"description": "Send a message.",
						},
					},
				},
				{
					"name": "notion",
					"reactions": []gin.H{
						{
							"name":        "create_default_comment",
							"description": "Create a default comment",
						},
						{
							"name":        "create_block",
							"description": "Create a block",
						},
						{
							"name":        "create_page",
							"description": "Create a page",
						},
						{
							"name":        "create_comment",
							"description": "Create a comment",
						},
						{
							"name":        "create_default_page",
							"description": "Create a default page",
						},
					},
				},
				{
					"name": "linkedin",
					"reactions": []gin.H{
						{
							"name":        "create_default_post",
							"description": "Create a default post",
						},
						{
							"name":        "create_post",
							"description": "Create a post",
						},
					},
				},
				{
					"name": "dropbox",
					"reactions": []gin.H{
						{
							"name":        "create_folder",
							"description": "Create a folder",
						},
						{
							"name":        "create_file",
							"description": "Create a file",
						},
						{
							"name":        "create_tags",
							"description": "Create tags",
						},
					},
				},
				{
					"name": "miro",
					"reactions": []gin.H{
						{
							"name":        "create_board",
							"description": "Create a board",
						},
						{
							"name":        "rename_board",
							"description": "Rename a board",
						},
						{
							"name":        "create_tags",
							"description": "Create tags",
						},
						{
							"name":        "create_sticky_card",
							"description": "Create a sticky card",
						},
					},
				},
			},
		},
	})
}
