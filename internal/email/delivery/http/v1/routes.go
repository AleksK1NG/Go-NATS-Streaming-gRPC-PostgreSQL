package v1

// MapRoutes products routes
func (h *emailHandlers) MapRoutes() {
	h.group.POST("", h.Create())
	// h.group.GET("/:email_id", h.GetByID())
	// h.group.POST("/search", h.Search())
}
