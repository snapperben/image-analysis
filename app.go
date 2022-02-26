package image_analysis

import "time"

// GalleryOfImages represents a group of images and their metadata
type GalleryOfImages struct {
	imagesMataData map[uint64]ImageMetadata
	mobileRegion   string
	galleryRegion  string
	timeOfYear     string
	firstTime      time.Time
	lastTime       time.Time
	locations      []ImageLocation
}

// ImageLocation represents the location of an image
type ImageLocation struct {
	country    string
	state      string
	county     string
	townOrCity string
	a
}

// ImageMetadata stores all the information about when and where an image was taken
// and all the collected metadata that can be inferred from that data
type ImageMetadata struct {
	id           int
	galleryIndex int
	lat          float32
	long         float32
	imageTime    time.Time // This is the actual Golang time of the image
	imageDate    uint64    // This id the UNIX timestamp at midday on the data of the image
	prevImage    *ImageMetaData
	nextImage    *ImageMetaData
}

// DateMetadata is designed to describe the significance of the date of an image so that the
// image can be labeled appropriately for the date
type DateMetadata struct {
	dateMetaName      string
	isHoliday         bool
	isReligious       bool
	daysAroundActual  int
	eventSignificance int
}

// CalendarManager holds all calendars for all regions
type CalendarManager struct {
	regionName string
}

// Calendar represents a region's calendar of significant dates
type Calendar struct {
	regionName    string
	calendarYear  int
	calendarDates map[uint64]CalendarDate
}

// CalendarDate is designed to represent any significant date in a region's calendar.
// The date of the calendar id UTC midday on the actual date
type CalendarDate struct {
	dateMidday       uint64
	dateName         string
	dateSignificance int // A percentage representing the significance of a date for the region
}

/////////////////////////////////   Member functions

// addImageToGallery adds an image to a gallery
func (g *GalleryOfImages) addImageToGallery() (_dateMetadata []DateMetadata, _err error) {
	_err = nil

	// TODO Iterate through the calendars that represent the gallery and determine if the gallery covers
	//		any significant dates

	return
}

// getTimeOfDay returns the type of gallery for the title
func (g GalleryOfImages) getGalleryType() string {
	// TODO Analyse the gallery and determine if it covers a trip, holiday or a home event
}

// overSignificantDates return metaData about significant dates the gallery covers
func (g *GalleryOfImages) getSignificantDates() (_dateMetadata []DateMetadata, _err error) {
	_err = nil

	// TODO Iterate through the calendars that represent the gallery and determine if the gallery covers
	//		any significant dates

	return
}

// readImageMetadata unmarshalls the image metadata from CSV files
func readImageMetadata() {

}

// processImageLocations Goes through the images in a file/gallery
func processImageLocations() {

}

func main() {

}
