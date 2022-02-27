package image_analysis

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// GalleryOfImages represents a group of images and their metadata
type GalleryOfImages struct {
	imagesMataData   map[uint64]ImageMetadata
	mobileHomeRegion string
	galleryRegion    string
	timeOfYear       string
	headImage        *ImageMetadata
	tailImage        *ImageMetadata
	locations        []ImageLocation
}

func (g *GalleryOfImages) insertImageMetadata(_im *ImageMetadata) {
	if g.headImage == nil {
		g.headImage = _im
	} else {
		if _im.ImageTime.Before(g.headImage.ImageTime) {
			_im.nextImage = g.headImage
			g.headImage = _im
		} else {
			g.headImage.InsertImageMetadata(_im)
		}
	}
}

// ReadCSVGallery reads in a CSV fuile of image metadata into the gallery
func (g *GalleryOfImages) ReadCSVGallery(_csvData string) error {
	var galleryImageCounter int = 0
	r := csv.NewReader(strings.NewReader(_csvData))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		im, err := NewImageMetadata(galleryImageCounter, record)
		if err != nil {
			return err
		}
		g.insertImageMetadata(im)
	}
	return nil
}

func (g *GalleryOfImages) ProcessImages(_csvData string) {

}

// ImageLocation represents the location of an image
type ImageLocation struct {
	country    string
	state      string
	county     string
	townOrCity string
}

// ImageMetadata stores all the information about when and where an image was taken
// and all the collected metadata that can be inferred from that data
type ImageMetadata struct {
	Id            int
	GalleryIndex  int
	Lat           float64
	Long          float64
	ImageTime     time.Time // This is the actual Golang time of the image
	prevImage     *ImageMetadata
	nextImage     *ImageMetadata
	imageLocation *ImageLocation
}

func (im *ImageMetadata) InsertImageMetadata(_im *ImageMetadata) {
	if _im.ImageTime.Before(im.ImageTime) {
		im.prevImage = _im
		im.nextImage = im
	} else {
		if im.nextImage == nil {
			im.nextImage = _im
		} else {
			im.nextImage.InsertImageMetadata(_im)
		}
	}
}

// NewImageMetadata is a utility function to construct an ImageMetadata object from []string{Date Lat Long}
func NewImageMetadata(_imageId int, _metaData []string) (*ImageMetadata, error) {
	if len(_metaData) != 3 {
		return nil, fmt.Errorf("input data to NewImageMetaData invalid, not enough data : %v", _metaData)
	}
	im := new(ImageMetadata)
	im.Id = _imageId
	imageTime, err := time.Parse(`2019-12-01T03:45:000Z`, _metaData[0])
	if err != nil {
		return nil, fmt.Errorf("time invalid for image #%d: %s", _imageId, _metaData[0])
	}
	im.ImageTime = imageTime

	coordinate, err := strconv.ParseFloat(_metaData[1], 32)
	if err != nil {
		return nil, fmt.Errorf("latitude invalid for image #%d: %s", _imageId, _metaData[1])
	}
	im.Lat = coordinate
	coordinate, err = strconv.ParseFloat(_metaData[2], 32)
	if err != nil {
		return nil, fmt.Errorf("longitude invalid for image #%d: %s", _imageId, _metaData[1])
	}
	im.Long = coordinate

	return im, nil
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
	return ""
}

// getSignificantDates return metaData about significant dates the gallery covers
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
