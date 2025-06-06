package asset

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAssetService 是 AssetService 的 mock 实现
type MockAssetService struct {
	mock.Mock
}

func (m *MockAssetService) Create(ctx context.Context, asset *domain.Asset) error {
	args := m.Called(ctx, asset)
	return args.Error(0)
}

func (m *MockAssetService) Update(ctx context.Context, asset *domain.Asset) error {
	args := m.Called(ctx, asset)
	return args.Error(0)
}

func (m *MockAssetService) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssetService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *MockAssetService) Get(ctx context.Context, id uuid.UUID) (*domain.Asset, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Asset), args.Error(1)
}

func (m *MockAssetService) List(ctx context.Context, options *service.ListOptions) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

func (m *MockAssetService) GetByGroup(ctx context.Context, groupID uuid.UUID, limit, offset int) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, groupID, limit, offset)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

func (m *MockAssetService) GetWithoutGroup(ctx context.Context, limit, offset int) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

// 创建测试用的 Echo 实例
func setupEcho() *echo.Echo {
	e := echo.New()
	return e
}

// 创建测试用的 Service mock
func createMockService(mockAssetService *MockAssetService) *service.Service {
	return &service.Service{
		AssetService: mockAssetService,
	}
}

func TestAssetController_Create(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	tests := []struct {
		name           string
		requestBody    CreateRequest
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "成功创建资产",
			requestBody: CreateRequest{
				Name: "Test Asset",
				Ip:   "192.168.1.1",
				Port: 22,
				Type: "host",
			},
			mockSetup: func() {
				mockService.On("Create", mock.Anything, mock.AnythingOfType("*domain.Asset")).Return(nil)
			},
			expectedStatus: 200,
		},
		{
			name: "无效的IP地址",
			requestBody: CreateRequest{
				Name: "Test Asset",
				Ip:   "invalid-ip",
				Port: 22,
				Type: "host",
			},
			mockSetup:      func() {},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/Api/V1/Assets", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.Create(c)

			// Assert
			if tt.expectedStatus < 400 {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetController_Get(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	assetID := uuid.New()
	mockAsset := &domain.Asset{
		Id:        assetID,
		Name:      "Test Asset",
		Ip:        "192.168.1.1",
		Port:      22,
		Type:      "host",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name           string
		assetID        string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:    "成功获取资产",
			assetID: assetID.String(),
			mockSetup: func() {
				mockService.On("Get", mock.Anything, assetID).Return(mockAsset, nil)
			},
			expectedStatus: 200,
		},
		{
			name:           "无效的UUID",
			assetID:        "invalid-uuid",
			mockSetup:      func() {},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/Assets/%s", tt.assetID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.assetID)
			// Execute
			err := controller.Get(c)

			if he, ok := err.(*echo.HTTPError); ok {
				rec.Code = he.Code
				rec.Body.WriteString(he.Message.(string))
			}

			// Assert
			if tt.expectedStatus < 400 {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetController_List(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	mockAssets := []*domain.Asset{
		{
			Id:   uuid.New(),
			Name: "Asset 1",
			Ip:   "192.168.1.1",
			Port: 22,
			Type: "host",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取资产列表",
			queryParams: map[string]string{
				"limit":  "10",
				"offset": "0",
			},
			mockSetup: func() {
				mockService.On("List", mock.Anything, mock.AnythingOfType("*service.ListOptions")).Return(mockAssets, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Assets", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.List(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetController_Update(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	assetID := uuid.New()

	tests := []struct {
		name           string
		requestBody    UpdateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功更新资产",
			requestBody: UpdateRequest{
				Id:   assetID.String(),
				Name: "Updated Asset",
				Ip:   "192.168.1.2",
				Port: 23,
				Type: "host",
			},
			mockSetup: func() {
				mockService.On("Update", mock.Anything, mock.AnythingOfType("*domain.Asset")).Return(nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/Assets/%s", tt.requestBody.Id), bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.requestBody.Id)

			// Execute
			err := controller.Update(c)
			if he, ok := err.(*echo.HTTPError); ok {
				rec.Code = he.Code
				rec.Body.WriteString(he.Message.(string))
			}
			// Assert
			if tt.expectedStatus < 400 {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetController_Delete(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	assetID := uuid.New()

	tests := []struct {
		name           string
		assetID        string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:    "成功删除资产",
			assetID: assetID.String(),
			mockSetup: func() {
				mockService.On("Delete", mock.Anything, assetID).Return(nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/Assets/%s", tt.assetID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.assetID)

			// Execute
			err := controller.Delete(c)
			if he, ok := err.(*echo.HTTPError); ok {
				rec.Code = he.Code
				rec.Body.WriteString(he.Message.(string))
			}
			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetController_ByGroup(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	groupID := uuid.New()
	mockAssets := []*domain.Asset{
		{
			Id:   uuid.New(),
			Name: "Asset 1",
			Ip:   "192.168.1.1",
			Port: 22,
			Type: "host",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取资产组的资产",
			queryParams: map[string]string{
				"id":     groupID.String(),
				"limit":  "20",
				"offset": "0",
			},
			mockSetup: func() {
				mockService.On("GetByGroup", mock.Anything, groupID, 0, 20).Return(mockAssets, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Assets/ByGroup", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.ByGroup(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetController_WithoutGroup(t *testing.T) {
	// Setup
	mockService := new(MockAssetService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	mockAssets := []*domain.Asset{
		{
			Id:   uuid.New(),
			Name: "Asset 1",
			Ip:   "192.168.1.1",
			Port: 22,
			Type: "host",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取未分组的资产",
			queryParams: map[string]string{
				"limit":  "20",
				"offset": "0",
			},
			mockSetup: func() {
				mockService.On("GetWithoutGroup", mock.Anything, 20, 0).Return(mockAssets, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Assets/WithoutGroup", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.WithoutGroup(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}
