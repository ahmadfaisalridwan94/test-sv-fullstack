package controllers

import "backend/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")	

	//Posts routes
	s.Router.HandleFunc("/article/", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	
	s.Router.HandleFunc("/article/{limit}/{offset}", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.Queries("status", "{status}")

	s.Router.HandleFunc("/article/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/article/{id}", middlewares.SetMiddlewareJSON(s.UpdatePost)).Methods("POST")
	s.Router.HandleFunc("/article/{id}/", middlewares.SetMiddlewareJSON(s.DeletePost)).Methods("POST")
	s.Router.HandleFunc("/count_article/{status}", middlewares.SetMiddlewareJSON(s.GetCountPosts)).Methods("GET")
}
