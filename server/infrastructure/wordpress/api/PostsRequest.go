package api

import (
	"server/infrastructure/wordpress/types"
	"time"
)

func GetWPPosts() (*[]types.WPPost, error) {
	var CacheKey = "wp_posts_cache"
	var APIURL = "https://www.heiwadai-hotel.co.jp/wp-json/wp/v2/posts/?_embed"
	var CacheExpiry = 60 * time.Minute

	posts, err := Request[[]types.WPPost](APIURL, CacheKey, CacheExpiry)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetWPPost(id uint) (*types.WPPost, error) {
	var CacheKey = "wp_post_cache_" + string(rune(id))
	var APIURL = "https://www.heiwadai-hotel.co.jp/wp-json/wp/v2/posts/" + string(rune(id)) + "?_embed"
	var CacheExpiry = 60 * time.Minute

	post, err := Request[types.WPPost](APIURL, CacheKey, CacheExpiry)
	if err != nil {
		return nil, err
	}
	return post, nil
}
