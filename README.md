# soup
Get an element with Id from html string
```
import "github.com/amanchik/soup"
div := soup.GetDivById("<html><div id='some_id'>some content</div></html>", "some_id")
```