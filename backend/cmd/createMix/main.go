package main

import (
	"cueclub/data"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var URL string
	flag.StringVar(&URL, "url", "", "url of video to add to database")
	flag.Parse()

	newMix, err := getYouTubeInfo(URL)
	if err != nil {
		log.Fatal("error fetching video data:", err)
	}

	newMix.YTID, err = extractIdFromURL(URL)
	if err != nil {
		log.Fatal(err)
	}
	newMix.URL = URL

	err = addMixToDB(newMix)
	if err != nil {
		log.Fatal("Error adding mix to db:", err)
	}

}

func getYouTubeInfo(videoURL string) (data.Mix, error) {
	oembedURL := "https://www.youtube.com/oembed?url=" + videoURL + "&format=json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", oembedURL, nil)
	if err != nil {
		return data.Mix{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; CueClubBot/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return data.Mix{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return data.Mix{}, fmt.Errorf("oEmbed request failed with status: %d", resp.StatusCode)
	}

	var result data.Mix
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return data.Mix{}, fmt.Errorf("could not decode oEmbed response: %w", err)
	}

	return result, nil
}

func extractIdFromURL(URL string) (string, error) {
	r := regexp.MustCompile(`[?&]v=([^&]+)`)
	match := r.FindStringSubmatch(URL)

	if len(match) > 1 {
		return match[1], nil
	}

	return "", fmt.Errorf("no ID found in URL")
}

func addMixToDB(m data.Mix) error {
	db, err := sql.Open("sqlite3", "./data/cueclub.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if mixExists(db, m) {
		println("Mix already in database")
		os.Exit(0)
	}

	stmt, err := db.Prepare(`
	INSERT INTO mixes(
		ytid,
		url,
		title,
		creator,
		creator_url,
		version
		) VALUES (?,?,?,?,?,1	
		)	
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.YTID, m.URL, m.Title, m.Creator, m.CreatorURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mix created sucessfully!")
	return nil
}

func mixExists(db *sql.DB, m data.Mix) bool {
	var count int
	query := `SELECT COUNT(*)
		FROM mixes
		WHERE ytid = ?`

	err := db.QueryRow(query, m.YTID).Scan(&count)
	if err != nil {
		log.Fatal("error checking for duplicates:", err)
	}
	return count > 0
}
