# go-mp4tag
MP4 tagger written in Go.

## Setup
```
go get github.com/Sorrow446/go-mp4tag
```
```go
import "github.com/Sorrow446/go-mp4tag"
```

## Usage Examples
```go
	tags := &mp4tag.Tags{
		Album: "album",
		AlbumArtist: "album artist",
		Title:       "title",
		TrackNumber: 1,
		TrackTotal:  20,
		Genre: "genre",
		DiskNumber:  3,
		DiskTotal:   10,
		Comment:     "comment",
	}
	err := mp4tag.Write("1.m4a", tags)
	if err != nil {
		panic(err)
	}
```
 Write album, album artist, title, track number, track total, genre, disk number, disk total, and comment tags.
 
 
 ```go
	tags := &mp4tag.Tags{
		Custom: map[string]string{
			"CUSTOMFIELD": "custom field",
			"CUSTOMFIELD2": "custom field 2",
		},
		Delete: []string{"genre", "cover"},
	}
	err := mp4tag.Write("1.m4a", tags)
	if err != nil {
		panic(err)
	}
```
Write two custom fields named `CUSTOMFIELD` and `CUSTOMFIELD2`, delete genre tag, and remove cover.


```go
	coverBytes, err := ioutil.ReadFile("cover.jpg")
	if err != nil {
		panic(err)
	}
	tags := &mp4tag.Tags{
		Cover: coverBytes,
	}
	err = mp4tag.Write("1.m4a", tags)
	if err != nil {
		panic(err)
	}
```
Write cover from `cover.jpg`.

## Misc
```go
type Tags struct {
	Album       string
	AlbumArtist string
	Artist      string
	Comment     string
	Composer    string
	Cover       []byte
	Custom      map[string]string
	Delete      []string
	DiskNumber  int
	DiskTotal   int
	Genre       string
	Label       string
	Title       string
	TrackNumber int
	TrackTotal  int
	Year        string
}
```
iTunes-style metadata only.       
Delete strings: album, albumartist, artist, comment, composer, cover, disk, genre, label, title, track, year.    
Custom tag deletion is not implemented yet.

## Thank you
 go-mp4tag relies heavily on abema's go-mp4 library.

## Disclaimer
Although go-mp4tag has been thoroughly tested, I will not be responsible for the very tiny chance of any corruption to your MP4 files.
