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
)

// OperationIds contains id-operation map used for install and register handler info
var OperationIds = map[int]*Operation{
	ApiHeadMainPage:     apiHead("/", "/"),
	ApiGetMainPage:      apiGet("/", "/"),
	ApiGetStaticAssets:  apiGet("/", "/static/*filepath"),
	ApiHeadStaticAssets: apiHead("/", "/static/*filepath"),
	ApiGetAppInfo:       apiGet(ApiVersion, "/appinfo"),
	ApiGetAlbums:        apiGet(ApiVersion, "/albums/"),
	ApiCreateAlbums:     apiPut(ApiVersion, "/albums/"),
	ApiUpdateAlbums:     apiPost(ApiVersion, "/albums/"),
	ApiGetAlbumsById:    apiGet(ApiVersion, "/albums/:albumId"),
	ApiDeleteAlbumsById: apiDelete(ApiVersion, "/albums/:albumId"),
}
