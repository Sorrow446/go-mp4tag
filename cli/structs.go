package main

type Args struct {
	Album       string            `arg:"--album" help:"Write album tag."`
	AlbumArtist string            `arg:"--albumArtist" help:"Write album artist tag."`
	Artist      string            `arg:"--artist" help:"Write artist tag."`
	Comment     string            `arg:"--comment" help:"Write comment tag."`
	Composer    string            `arg:"--composer" help:"Write composer tag."`
	Copyright   string            `arg:"--copyright" help:"Write copyright tag."`
	Cover       string            `arg:"--cover" help:"Path of cover to write. JPEG is recommended."`
	Custom      map[string]string `arg:"--custom" help:"Write custom tags. Multiple tags with the same field name can be written.\n\t\t\t Example: \"--custom MYCUSTOMFIELD1=value1 MYCUSTOMFIELD2=value2\""`
	Delete      []string          `arg:"-d, --delete" help:"Tags to delete.\n\t\t\t Options: album, albumartist, artist, comment, composer, cover, disk, genre, label, title, track, year.\n\t\t\t Example: \"-d album albumartist\""`
	DiskNumber  int               `arg:"--diskNumber" help:"Write disk number tag."`
	DiskTotal   int               `arg:"--diskTotal" help:"Write disk total tag. Can't be written without disk number tag."`
	FilePath    string            `arg:"positional, required" help:"Path of file to write to."`
	Genre       string            `arg:"--genre" help:"Write genre tag."`
	Label       string            `arg:"--label" help:"Write label tag."`
	Title       string            `arg:"--title" help:"Write title tag."`
	TrackNumber int               `arg:"--trackNumber" help:"Write track number tag."`
	TrackTotal  int               `arg:"--trackTotal" help:"Write track total tag. Can't be written without track number tag."`
	Year        string            `arg:"--year" help:"Write year tag."`
}

type Tags struct {
	Album       string
	AlbumArtist string
	Artist      string
	Comment     string
	Composer    string
	Copyright   string
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
