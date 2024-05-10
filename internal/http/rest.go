package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"server/internal/http/router"
	"time"
)

func StartServer() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rest-Service is active."))
	})

	/*	// RESTy routes for "articles" resource
		r.Route("/articles", func(r chi.Router) {
			r.With(paginate).Get("/", listArticles)                           // GET /articles
			r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

			r.Post("/", createArticle)       // POST /articles
			r.Get("/search", searchArticles) // GET /articles/search

			// Regexp url parameters:
			r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug) // GET /articles/home-is-toronto

			// Subrouters:
			r.Route("/{articleID}", func(r chi.Router) {
				r.Use(ArticleCtx)
				r.Get("/", getArticle)       // GET /articles/123
				r.Put("/", updateArticle)    // PUT /articles/123
				r.Delete("/", deleteArticle) // DELETE /articles/123
			})
		})
	*/
	// Mount the admin sub-router

	r.Route("/api/v2", func(r chi.Router) {
		r.Mount("/package", router.PackageRouter())
	})

	http.ListenAndServe(":8080", r)
}
