package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"createDate"`
	UpdateDate time.Time `json:"updateDate"`
}

var (
	posts     = make(map[int]Post)
	idCounter = 1
	mutex     = &sync.Mutex{}
)

func initPosts() {
	post1 := Post{
		ID:         idCounter,
		Title:      "첫 번째 샘플 포스트",
		Content:    "이것은 첫 번째 포스트의 내용입니다.",
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}
	posts[idCounter] = post1
	idCounter++
	post2 := Post{
		ID:         idCounter,
		Title:      "두 번째 포스트",
		Content:    "두 번째 포스트의 내용입니다.",
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}
	posts[idCounter] = post2
	idCounter++
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	postList := make([]Post, 0, len(posts))
	for _, post := range posts {
		postList = append(postList, post)
	}

	json.NewEncoder(w).Encode(postList)
}

func getPostByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	post, found := posts[id]
	if !found {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	response := struct {
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		CreateDate time.Time `json:"createDate"`
		UpdateDate time.Time `json:"updateDate"`
	}{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		CreateDate: post.CreateDate,
		UpdateDate: post.UpdateDate,
	}

	json.NewEncoder(w).Encode(response)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPost.ID = idCounter
	newPost.CreateDate = time.Now()
	newPost.UpdateDate = time.Now()
	posts[idCounter] = newPost
	idCounter++

	json.NewEncoder(w).Encode(newPost)
}

func main() {
	initPosts()

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Get("id") != "" {
				getPostByID(w, r)
			} else {
				getAllPosts(w, r)
			}
		case http.MethodPost:
			addPost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
