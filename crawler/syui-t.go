package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

/*

http://syui.github.io/

    <div id="main" class="container">
      <div id="content">
        <div class="row">
          <div class="page-content col-md-9">
            <div class="blog-index">
              <p class="meta text-muted text-uppercase"></p>

              <h1 class="entry-title"><a href=
              "https://syui.github.io/blog/2016/05/25/spritestudio-sample" class=
              "permalink">spritestudio-sample</a></h1>

*/

func main() {
	doc, _ := goquery.NewDocument("http://syui.github.io/")
	doc.Find("#content .blog-index h1 a").Each(func(_ int, s *goquery.Selection) {
		if articleUrl, ok := s.Attr("href"); ok {
			fmt.Println(articleUrl)
		}
	})
}

/*

https://syui.github.io/blog/2016/05/25/spritestudio-sample
https://syui.github.io/blog/2016/05/24/spritestudio-undroid
https://syui.github.io/blog/2016/05/22/mac-imovie
https://syui.github.io/blog/2016/05/21/ue4-mmd-miku-motion
https://syui.github.io/blog/2016/05/20/music-make-vocaloid
...

*/
