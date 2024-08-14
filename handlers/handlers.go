package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello World!\n")
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article ...\n")
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article Np.%d", articleID)
	if req.Method == http.MethodGet {
		io.WriteString(w, resString)
	} else {
		http.Error(w, "invalid error", http.StatusMethodNotAllowed)
	}
}
func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice ...\n")
	} else {
		http.Error(w, "invalid error", http.StatusMethodNotAllowed)
	}
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment ...\n")
	} else {
		http.Error(w, "invalid error", http.StatusMethodNotAllowed)
	}

}
