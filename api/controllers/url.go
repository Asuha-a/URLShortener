package controllers

import (
	"context"
	"log"
	"time"

	pbURL "github.com/Asuha-a/URLShortener/api/pb/url"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	urlAddress = "url:50052"
)

// GetAllURL get all of URLs
func GetAllURL(c *gin.Context) {
	conn, err := grpc.Dial(urlAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbURL.NewURLClient(conn)

	token := c.Query("token")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.GetAllURL(ctx, &pbURL.GetAllURLRequest{
		Token: token,
	})
	/*
		type node struct {
			Condition     string `json:"condition"`
			MatchedURL    string `json:"matched_url"`
			NotMatchedURL string `json:"not_matched_url"`
			Matched       *node  `json:"matched"`
			NotMatched    *node  `json:"not_matched"`
		}

		type urlJSON struct {
			UUID       uuid.UUID `json:"uuid"`
			UserID     uuid.UUID `json:"user_id"`
			Title      string    `json:"title"`
			RedirectTo node      `json:"redirect_to"`
			CreatedAt  time.Time `json:"created_at"`
		}
	*/

	var responces []gin.H
	succeed := true
	for {
		r, err := stream.Recv()
		if err != nil {
			succeed = false
			panic(err)
			c.AbortWithStatus(400)
		} else {
			/*
				url := &urlJSON{
					UUID:       r.GetUuid(),
					UserID:     r.GetUserId(),
					Title:      r.GetTitle(),
					RedirectTo: r.GetRedirectTo(),
					CreatedAt:  r.GetCreatedAt(),
				}
			*/
			url := gin.H{
				"uuid":        r.GetUuid(),
				"user_id":     r.GetUserId(),
				"title":       r.GetTitle(),
				"redirect_to": r.GetRedirectTo(),
				"created_at":  r.GetCreatedAt(),
			}
			// responce, _ := json.Marshal(url)
			responces = append(responces, url)
		}
	}
	if succeed {
		c.JSON(200, gin.H{
			"urls": responces,
		})
	}
}

// PostURL create the shorten URL
func PostURL(c *gin.Context) {
	conn, err := grpc.Dial(urlAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbURL.NewURLClient(conn)

	type request struct {
		token      string
		title      string
		redirectTo *pbURL.Node
	}
	var s request
	c.Bind(&s)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.PostURL(ctx, &pbURL.PostURLRequest{
		Token:      s.token,
		Title:      s.title,
		RedirectTo: s.redirectTo,
	})
	if err != nil {
		panic(err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(201, gin.H{
			"uuid":        r.GetUuid(),
			"user_id":     r.GetUserId(),
			"tilte":       r.GetTitle(),
			"url":         r.GetUrl(),
			"redirect_to": r.GetRedirectTo(),
			"created_at":  r.GetCreatedAt(),
		})
	}
}
