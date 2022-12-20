package handler

// func DBConnPool(pool *pgxpool.Pool) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("db", pool)
// 		c.Next()
// 	}
// }

// func IsBoardExists() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db_pool := c.MustGet("db").(*pgxpool.Pool)
// 		board := c.Param("board")
// 		res, err := db.SelectBoardExists(db_pool, board)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 		} else if !res {
// 			c.AbortWithStatus(http.StatusNotFound)
// 		} else {
// 			c.Next()
// 		}

// 	}
// }

// func SetupMiddlewares(app *gin.Engine) {
// 	app.Use(gin.Logger())
// 	app.Use(gin.Recovery())
// }

// func PassConf(cfg *HandlerConfig) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("cfg", cfg)
// 		c.Next()
// 	}
// }
