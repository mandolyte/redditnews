
==============================================================
Comments/Feedback on book
--------------------------------------------------------------
==> first:
- on second iteration, you have this statement, but no check is made for any error:

	err = json.NewDecoder(resp.Body).Decode(r)
	
I left the check from the first iteration in place:

	if err != nil {
		log.Fatal(err)
	}

===============================================================
godoc
---------------------------------------------------------------

$ godoc $GOPATH/src/github.com/mandolyte/redditnews
PACKAGE DOCUMENTATION

package redditnews
    import "/home/cecil/Workspace/src/github.com/mandolyte/redditnews"



FUNCTIONS

func Email() string

func Get(reddit string) ([]Item, error)
    Get fetches the most recent Items posted to the specified subreddit.


TYPES

type Item struct {
    Author string `json:"author"`
    Score  int    `json:"score"`
    URL    string `json:"url"`
    Title  string `json:"title"`
}
    Item describes a RedditNews item.


func (i Item) String() string



SUBDIRECTORIES

	archive
	reddit
	redditmail
