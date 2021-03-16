package v1

// MapRoutes emails REST API routes
func (h *emailHandlers) MapRoutes() {
	h.group.POST("", h.Create())
	h.group.GET("/:email_id", h.GetByID())
	h.group.GET("/search", h.Search())
}
