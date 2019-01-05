package api

const (
	// OperationIds for api handler
	ApiGetMainPage      = iota // Get main page
	ApiHeadMainPage            // Head main page
	ApiGetStaticAssets         // Get static assets
	ApiHeadStaticAssets        // Head static assets
	ApiGetAppInfo              // Get application information
	ApiGetAlbums               // Get albums
	ApiCreateAlbums            // Create an albums
	ApiUpdateAlbums            // Update an albums
	ApiGetAlbumsById           // Get album by Id
	ApiDeleteAlbumsById        // Delete an albums by Id

	groupStatic = "static"   // Static assets group
	groupApi    = ApiVersion // OpenAPI group
)

// OperationIds contains id-operation map used for install and register handler info
var OperationIds = map[int]*Operation{
	ApiHeadMainPage:     apiHead("/"),
	ApiGetMainPage:      apiGet("/"),
	ApiGetStaticAssets:  apiGet("/*filepath", groupStatic),
	ApiHeadStaticAssets: apiHead("/*filepath", groupStatic),
	ApiGetAppInfo:       apiGet("/appinfo", groupApi),
	ApiGetAlbums:        apiGet("/albums/", groupApi),
	ApiCreateAlbums:     apiPut("/albums/", groupApi),
	ApiUpdateAlbums:     apiPost("/albums/", groupApi),
	ApiGetAlbumsById:    apiGet("/albums/:albumId", groupApi),
	ApiDeleteAlbumsById: apiDelete("/albums/:albumId", groupApi),
}
