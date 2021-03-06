package app

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

var (
	AppVersion string
	BuildTime  string
	DB         *sql.DB
)

func InitDB() {
	file := revel.Config.StringDefault("stream.dbfile", "")
	var err error
	DB, err = sql.Open("sqlite3", file)
	if err != nil {
		revel.INFO.Println("DB error", err)
	}
	revel.INFO.Println("DB connected")
}

func init() {
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(InitDB)
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
