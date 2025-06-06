package credential

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

// MockCredentialService 是 CredentialService 的 mock 实现
type MockCredentialService struct {
	mock.Mock
}

func (m *MockCredentialService) Create(ctx context.Context, credential *domain.Credential) error {
	args := m.Called(ctx, credential)
	return args.Error(0)
}

func (m *MockCredentialService) Update(ctx context.Context, credential *domain.Credential) error {
	args := m.Called(ctx, credential)
	return args.Error(0)
}

func (m *MockCredentialService) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCredentialService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *MockCredentialService) Get(ctx context.Context, id uuid.UUID) (*domain.Credential, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Credential), args.Error(1)
}

func (m *MockCredentialService) List(ctx context.Context, options *service.ListOptions) ([]*domain.Credential, int, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]*domain.Credential), args.Int(1), args.Error(2)
}

func (m *MockCredentialService) GetByAsset(ctx context.Context, assetID uuid.UUID, limit int, offset int) ([]*domain.Credential, int, error) {
	args := m.Called(ctx, assetID, limit, offset)
	return args.Get(0).([]*domain.Credential), args.Int(1), args.Error(2)
}

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

func (m *MockAssetService) GetWithoutGroup(ctx context.Context, limit int, offset int) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

func (m *MockAssetService) GetByGroup(ctx context.Context, groupID uuid.UUID, limit int, offset int) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, groupID, limit, offset)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

func (m *MockAssetService) Get(ctx context.Context, id uuid.UUID) (*domain.Asset, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Asset), args.Error(1)
}

func (m *MockAssetService) List(ctx context.Context, options *service.ListOptions) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

// 创建测试用的 Echo 实例
func setupEcho() *echo.Echo {
	e := echo.New()
	return e
}

// 创建测试用的 Service mock
func createMockService(mockCredentialService *MockCredentialService, mockAssetService *MockAssetService) *service.Service {
	return &service.Service{
		CredentialService: mockCredentialService,
		AssetService:      mockAssetService,
	}
}

func TestCredentialController_Create(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	tests := []struct {
		name           string
		requestBody    CreateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功创建凭据",
			requestBody: CreateRequest{
				AssetId:  uuid.New().String(),
				Type:     "password",
				Username: "testuser",
				Secret:   "testpass",
			},
			mockSetup: func() {
				mockService.On("Create", mock.Anything, mock.AnythingOfType("*domain.Credential")).Return(nil)
				mockAssetService.On("Get", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&domain.Asset{
					Id:   uuid.New(),
					Name: "testasset",
				}, nil)
			},
			expectedStatus: 200,
		},
		{
			name: "空类型",
			requestBody: CreateRequest{
				AssetId:  uuid.New().String(),
				Type:     "",
				Username: "testuser",
				Secret:   "testpass",
			},
			mockSetup: func() {
				mockAssetService.On("Get", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("asset not found"))
			},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/Api/V1/Credentials", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.Create(c)
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

func TestCredentialController_Get(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	credentialID := uuid.New()
	mockCredential := &domain.Credential{
		Id:        credentialID,
		Type:      "password",
		Username:  "testuser",
		Secret:    "testpass",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name           string
		credentialID   string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:         "成功获取凭据",
			credentialID: credentialID.String(),
			mockSetup: func() {
				mockService.On("Get", mock.Anything, credentialID).Return(mockCredential, nil)
			},
			expectedStatus: 200,
		},
		{
			name:           "无效的UUID",
			credentialID:   "invalid-uuid",
			mockSetup:      func() {},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Credentials/"+tt.credentialID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.credentialID)

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

func TestCredentialController_List(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	mockCredentials := []*domain.Credential{
		{
			Id:       uuid.New(),
			Type:     "password",
			Username: "testuser",
			Secret:   "testpass",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取凭据列表",
			queryParams: map[string]string{
				"limit":  "10",
				"offset": "0",
			},
			mockSetup: func() {
				mockService.On("List", mock.Anything, mock.AnythingOfType("*service.ListOptions")).Return(mockCredentials, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Credentials", nil)
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

func TestCredentialController_Update(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	tests := []struct {
		name           string
		requestBody    UpdateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功更新凭据",
			requestBody: UpdateRequest{
				Id:       uuid.New().String(),
				Type:     "password",
				Username: "updateduser",
				Secret:   "updatedpass",
			},
			mockSetup: func() {
				mockService.On("Update", mock.Anything, mock.AnythingOfType("*domain.Credential")).Return(nil)
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
			req := httptest.NewRequest(http.MethodPost, "/Api/V1/Credentials/"+tt.requestBody.Id, bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.requestBody.Id)

			// Execute
			err := controller.Update(c)

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

func TestCredentialController_Delete(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	credentialID := uuid.New()

	tests := []struct {
		name           string
		credentialID   string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:         "成功删除凭据",
			credentialID: credentialID.String(),
			mockSetup: func() {
				mockService.On("Delete", mock.Anything, credentialID).Return(nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodDelete, "/Api/V1/Credentials/"+tt.credentialID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.credentialID)

			// Execute
			err := controller.Delete(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestCredentialController_DeleteCollection(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	ids := []string{uuid.New().String(), uuid.New().String()}

	tests := []struct {
		name           string
		requestBody    DeleteCollectionRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功批量删除凭据",
			requestBody: DeleteCollectionRequest{
				Ids: ids,
			},
			mockSetup: func() {
				mockService.On("DeleteCollection", mock.Anything, mock.AnythingOfType("[]uuid.UUID")).Return(nil)
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
			req := httptest.NewRequest(http.MethodDelete, "/Api/V1/Credentials", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.DeleteCollection(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestCredentialController_ByAsset(t *testing.T) {
	// Setup
	mockService := new(MockCredentialService)
	mockAssetService := new(MockAssetService)
	mockSvc := createMockService(mockService, mockAssetService)
	controller := NewController(mockSvc)
	e := setupEcho()

	assetID := uuid.New()
	mockCredentials := []*domain.Credential{
		{
			Id:       uuid.New(),
			Type:     "password",
			Username: "testuser",
			Secret:   "testpass",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取资产的凭据",
			queryParams: map[string]string{
				"id":     assetID.String(),
				"limit":  "20",
				"offset": "0",
			},
			mockSetup: func() {
				mockService.On("GetByAsset", mock.Anything, assetID, 20, 0).Return(mockCredentials, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Credentials/ByAsset", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.ByAsset(c)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}
