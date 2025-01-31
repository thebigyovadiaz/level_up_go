// Challenge: Make a Playlist
// Task: given N sorted slices of K songs implement a
//		function that outputs the merged slice of sorted songs.

package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
)

const FILE = "./step-01/09/songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	PlayCount int    `json:"play_count"`
	Album     string `json:"album"`

	AlbumCount, SongCount int
}

// An PlaylistHeap is a max-heap of PlaylistEntries.
type PlaylistHeap []Song

func (h PlaylistHeap) Len() int {
	return len(h)
}
func (h PlaylistHeap) Less(i, j int) bool {
	// We want Pop to return the highest play count.
	return h[i].PlayCount > h[j].PlayCount
}
func (h PlaylistHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PlaylistHeap) Push(x any) {
	*h = append(*h, x.(Song))
}

func (h *PlaylistHeap) Pop() any {
	original := *h
	n := len(original)
	x := original[n-1]
	*h = original[0 : n-1]
	return x
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	var playlist []Song
	pHeap := &PlaylistHeap{}
	if len(albums) == 0 {
		return playlist
	}

	// initialize the heap and add first of each album, since they are the max
	heap.Init(pHeap)
	for i, f := range albums {
		firstSong := f[0]
		firstSong.AlbumCount, firstSong.SongCount = i, 0
		heap.Push(pHeap, firstSong)
	}

	for pHeap.Len() != 0 {
		// take max elem from the list
		p := heap.Pop(pHeap)
		song := p.(Song)
		playlist = append(playlist, song)
		// the next song after the max is a good candidate to look at
		if song.SongCount < len(albums[song.AlbumCount])-1 {
			nextSong := albums[song.AlbumCount][song.SongCount+1]
			nextSong.AlbumCount, nextSong.SongCount =
				song.AlbumCount, song.SongCount+1
			heap.Push(pHeap, nextSong)
		}
	}

	return playlist
}

func importData() [][]Song {
	file, err := ioutil.ReadFile(FILE)
	if err != nil {
		panic(err)
	}

	var data [][]Song
	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatal(err)
	}

	return data
}

func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}
