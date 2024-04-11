package grafana

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/lahsivjar/grafonnet-playground/config"
)

// Service exposes methods to handle grafana dashboards
type Service interface {
	SetupCleanerJob(ctx context.Context) error
	CreateDashboard(*CreateRequest) (*CreateResponse, error)
	DeleteDashboard(string) error
}

type grafana struct {
	cfg *config.Config
	pq  *PriorityQueue
}

// NewService returns a new service with default implementations for grafana service
// It also starts the cleaner job if cleaup is enabled in the configuration
func NewService(cfg *config.Config) Service {
	g := &grafana{
		cfg: cfg,
	}
	if cfg.AutoCleanup {
		g.pq = NewPriorityQueue()
	}
	return g
}

// SetupCleanerJob starts the cleaner job if cleanup is enabled in the configurations
func (g *grafana) SetupCleanerJob(ctx context.Context) error {
	if g.pq == nil {
		return errors.New("Priority queue for cleanup job is nil")
	}

	startCleaner(ctx, g.pq, g, g.cfg)
	return nil
}

// CreateDashboard creates a new grafana dashboard
func (g *grafana) CreateDashboard(cr *CreateRequest) (*CreateResponse, error) {
	reqBody, err := getRequestBody(cr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.cfg.GrafanaPostURL+"/api/dashboards/db", reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set(g.cfg.GrafanaAPIKeyHeaderName, "Bearer "+g.cfg.GrafanaAPIKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		var gRes CreateResponse
		err = json.NewDecoder(resp.Body).Decode(&gRes)
		if err != nil {
			return nil, err
		}

		if g.cfg.AutoCleanup {
			g.pq.Push(&Item{Key: gRes.UID, ProcessAt: time.Now().Add(g.cfg.CleanupAfter)})
		}
		return &gRes, nil
	}

	errorMsg, err := ioutil.ReadAll(resp.Body)
	return nil, fmt.Errorf("Error occurred while creating graph: %s", errorMsg)
}

func (g *grafana) DeleteDashboard(uid string) error {
	req, err := http.NewRequest("DELETE", g.cfg.GrafanaPostURL+"/api/dashboards/uid/"+uid, nil)
	if err != nil {
		return err
	}
	req.Header.Set(g.cfg.GrafanaAPIKeyHeaderName, "Bearer "+g.cfg.GrafanaAPIKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}

	errorMsg, err := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("Error occurred while deleting graph: %s", errorMsg)
}

func getRequestBody(req *CreateRequest) (io.Reader, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(b)), nil
}
