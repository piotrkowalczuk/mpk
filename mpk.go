package mpk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	basePath                = "http://pasazer.mpk.wroc.pl/position.php"
	TransportationTypeBus   = TransportationType("bus")
	TransportationTypeTram  = TransportationType("tram")
	TransportationTypeTrain = TransportationType("train")
)

// TransportationType ...
type TransportationType string

// String...
func (t TransportationType) String() string {
	return string(t)
}

// Position represents vehicle position in single moment of time.
type Position struct {
	Driver int64     `json:"k"`
	Line   string    `json:"name"`
	Type   string    `json:"type"`
	X      float64   `json:"x"`
	Y      float64   `json:"y"`
	Moment time.Time `json:"moment"`
}

// Service ...
type Service struct {
	BasePath string
	GPS      *GPSService

	client *http.Client
}

// New ...
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("mpk: client is nil")
	}

	s := &Service{client: client, BasePath: basePath}
	s.GPS = NewGPSService(s)

	return s, nil
}

// GPSService ...
type GPSService struct {
	service *Service
}

// NewGPSService ...
func NewGPSService(service *Service) *GPSService {
	return &GPSService{service: service}
}

// Fetch ...
func (gs *GPSService) Fetch(transits map[TransportationType][]string) ([]*Position, error) {
	values := url.Values{}
	now := time.Now()

	for transportationType := range transits {
		for _, name := range transits[transportationType] {
			values.Add("busList["+transportationType.String()+"][]", name)
		}
	}

	req, _ := http.NewRequest(
		"POST",
		gs.service.BasePath,
		bytes.NewBufferString(values.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	resp, err := gs.service.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var positions []*Position
	if err := json.NewDecoder(resp.Body).Decode(&positions); err != nil {
		return nil, err
	}

	for i := range positions {
		positions[i].Moment = now
	}

	return positions, nil
}
