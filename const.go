package insert_it

var Region = Regions["Mannheim"]
var svcUrl = baseurl + Region + service
var imgUrl = baseurl + Region + "img/"

const (
	baseurl = "https://www.insert-it.de/"
	service = "Webservice/"

	getAllStreets             = "GetAllStreets"
	getHourseNumbersForStreet = "GetHouseNumbersByStreetName?streetName=%s"
	getNextEmptyings          = "GetNextEmptyingsFromLocation?"
	getEmptyings              = "GetEmptyingsByStreetNameAndNumber?"
	searchParams              = "streetName=%s&houseNumberStart=%d&houseNumberStartExtra=%s&houseNumberEnd=%s&houseNumberEndExtra=%s"

	servicePointTypes = "GetAllPointObjectTypes"
	servicePoints     = "GetAllPointObjects"
)

var Regions = map[string]string{
	"Hattingen": "BmsAbfallkalenderHattingen/",
	"Herne":     "BmsAbfallkalenderHerne/",
	"Kassel":    "BmsAbfallkalenderKassel/",
	"Krefeld":   "BmsAbfallkalenderKrefeld/",
	"Luebeck":   "BmsAbfallkalenderLuebeck/",
	"Mannheim":  "BmsAbfallkalenderMannheim/",
	"Offenbach": "BmsAbfallkalenderOffenbach/",
}
