# soup
Get an element with Id from html string
```
import "github.com/amanchik/soup"
div := soup.GetDivById("<html><div id='some_id'>some content</div></html>", "some_id")
fmt.Println(div)
```
This prints.
```
<div id='some_id'>some content</div>
```