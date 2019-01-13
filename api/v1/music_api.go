package v1

const (
	// MinaGroup music group api
	MusicGroup = Version

	// ApiGetMainPage indicate main page
	ApiGetMainPage ApiId = "/"

	// indicate head main page
	ApiHeadMainPage ApiId = "/"

	// ApiGetStaticAssets get static assets
	ApiGetStaticAssets ApiId = "/static/*filepath"

	// ApiHeadStaticAssets head static assets
	ApiHeadStaticAssets ApiId = "/static/*filepath" // Head static assets

	// ApiGetAppInfo get application information
	ApiGetAppInfo ApiId = "/appinfo/"

	// ApiGetAlbums get albums
	ApiGetAlbums ApiId = "/albums/"

	// ApiCreateAlbums create an albums
	ApiCreateAlbums ApiId = "/albums/"

	// ApiUpdateAlbums update an albums
	ApiUpdateAlbums ApiId = "/albums/"

	// ApiGetAlbumsById get album by Id
	ApiGetAlbumsById ApiId = "/albums/:albumId/"

	// ApiDeleteAlbumsById delete an albums by Id
	ApiDeleteAlbumsById ApiId = "/albums/:albumId/"
)
