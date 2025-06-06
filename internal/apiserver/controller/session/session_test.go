package session

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

// MockSessionService 是 SessionService 的 mock 实现
type MockSessionService struct {
	mock.Mock
}

func (m *MockSessionService) Create(ctx context.Context, session *domain.Session) error {
	args := m.Called(ctx, session)
	return args.Error(0)
}

func (m *MockSessionService) Update(ctx context.Context, session *domain.Session) error {
	args := m.Called(ctx, session)
	return args.Error(0)
}

func (m *MockSessionService) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSessionService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *MockSessionService) Get(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Session), args.Error(1)
}

func (m *MockSessionService) List(ctx context.Context, options *service.ListOptions) ([]*domain.Session, int, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]*domain.Session), args.Int(1), args.Error(2)
}

func (m *MockSessionService) GetByAsset(ctx context.Context, assetID uuid.UUID, limit int, offset int) ([]*domain.Session, int, error) {
	args := m.Called(ctx, assetID, limit, offset)
	return args.Get(0).([]*domain.Session), args.Int(1), args.Error(2)
}

type MockAssetService struct {
	mock.Mock
}

var _ service.AssetService = (*MockAssetService)(nil)

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

func (m *MockAssetService) GetByGroup(ctx context.Context, groupID uuid.UUID, limit int, offset int) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, groupID, limit, offset)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

func (m *MockAssetService) GetWithoutGroup(ctx context.Context, limit int, offset int) ([]*domain.Asset, int, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.Asset), args.Int(1), args.Error(2)
}

type MockCredentialService struct {
	mock.Mock
}

var _ service.CredentialService = (*MockCredentialService)(nil)

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

func createMockService(mockSessionService *MockSessionService, mockAssetService *MockAssetService, mockCredentialService *MockCredentialService) *service.Service {
	return &service.Service{
		SessionService:    mockSessionService,
		AssetService:      mockAssetService,
		CredentialService: mockCredentialService,
	}
}

// 创建测试用的 Echo 实例
func setupEcho() *echo.Echo {
	e := echo.New()
	return e
}

// 创建测试用的 Service mock

func TestSessionController_Create(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
	controller := NewController(mockSvc)
	e := setupEcho()

	tests := []struct {
		name           string
		requestBody    CreateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功创建会话",
			requestBody: CreateRequest{
				AssetId: uuid.New().String(),
				CredId:  uuid.New().String(),
				Type:    "ssh",
				Status:  "active",
			},
			mockSetup: func() {
				mockSessionService.On("Create", mock.Anything, mock.AnythingOfType("*domain.Session")).Return(nil)
				mockAssetService.On("Get", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&domain.Asset{}, nil)
				mockCredentialService.On("Get", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&domain.Credential{}, nil)
			},
			expectedStatus: 200,
		},
		{
			name: "空类型",
			requestBody: CreateRequest{
				AssetId: uuid.New().String(),
				CredId:  uuid.New().String(),
				Type:    "",
				Status:  "running",
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
			req := httptest.NewRequest(http.MethodPost, "/Api/V1/Sessions", bytes.NewReader(body))
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
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}

func TestSessionController_Get(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
	controller := NewController(mockSvc)
	e := setupEcho()

	sessionID := uuid.New()
	mockSession := &domain.Session{
		Id:        sessionID,
		Type:      "ssh",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name           string
		sessionID      string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:      "成功获取会话",
			sessionID: sessionID.String(),
			mockSetup: func() {
				mockSessionService.On("Get", mock.Anything, sessionID).Return(mockSession, nil)
			},
			expectedStatus: 200,
		},
		{
			name:           "无效的UUID",
			sessionID:      "invalid-uuid",
			mockSetup:      func() {},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/Sessions/%s", tt.sessionID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.sessionID)

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
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}

func TestSessionController_List(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
	controller := NewController(mockSvc)
	e := setupEcho()

	mockSessions := []*domain.Session{
		{
			Id:     uuid.New(),
			Type:   "ssh",
			Status: "active",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取会话列表",
			queryParams: map[string]string{
				"limit":  "10",
				"offset": "0",
			},
			mockSetup: func() {
				mockSessionService.On("List", mock.Anything, mock.AnythingOfType("*service.ListOptions")).Return(mockSessions, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Sessions", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.List(c)
			if he, ok := err.(*echo.HTTPError); ok {
				rec.Code = he.Code
				rec.Body.WriteString(he.Message.(string))
			}
			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}

func TestSessionController_Update(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
	controller := NewController(mockSvc)
	e := setupEcho()

	sessionID := uuid.New()

	tests := []struct {
		name           string
		requestBody    UpdateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功更新会话",
			requestBody: UpdateRequest{
				Id:     sessionID.String(),
				Type:   "ssh",
				Status: "inactive",
			},
			mockSetup: func() {
				mockSessionService.On("Update", mock.Anything, mock.AnythingOfType("*domain.Session")).Return(nil)
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
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/Sessions/%s", tt.requestBody.Id), bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

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
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}

func TestSessionController_Delete(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
	controller := NewController(mockSvc)
	e := setupEcho()

	sessionID := uuid.New()

	tests := []struct {
		name           string
		sessionID      string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:      "成功删除会话",
			sessionID: sessionID.String(),
			mockSetup: func() {
				mockSessionService.On("Delete", mock.Anything, sessionID).Return(nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/Sessions/%s", tt.sessionID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.sessionID)

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
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}

func TestSessionController_DeleteCollection(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
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
			name: "成功批量删除会话",
			requestBody: DeleteCollectionRequest{
				Ids: ids,
			},
			mockSetup: func() {
				mockSessionService.On("DeleteCollection", mock.Anything, mock.AnythingOfType("[]uuid.UUID")).Return(nil)
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
			req := httptest.NewRequest(http.MethodDelete, "/Api/V1/Sessions/Collection", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.DeleteCollection(c)
			if he, ok := err.(*echo.HTTPError); ok {
				rec.Code = he.Code
				rec.Body.WriteString(he.Message.(string))
			}
			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}

func TestSessionController_ByAsset(t *testing.T) {
	// Setup
	mockSessionService := new(MockSessionService)
	mockAssetService := new(MockAssetService)
	mockCredentialService := new(MockCredentialService)
	mockSvc := createMockService(mockSessionService, mockAssetService, mockCredentialService)
	controller := NewController(mockSvc)
	e := setupEcho()

	assetID := uuid.New()
	mockSessions := []*domain.Session{
		{
			Id:     uuid.New(),
			Type:   "ssh",
			Status: "active",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取资产的会话",
			queryParams: map[string]string{
				"id":     assetID.String(),
				"limit":  "20",
				"offset": "0",
			},
			mockSetup: func() {
				mockSessionService.On("GetByAsset", mock.Anything, assetID, 20, 0).Return(mockSessions, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/Sessions/ByAsset", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.ByAsset(c)
			if he, ok := err.(*echo.HTTPError); ok {
				rec.Code = he.Code
				rec.Body.WriteString(he.Message.(string))
			}
			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Reset mock
			mockSessionService.ExpectedCalls = nil
			mockSessionService.Calls = nil
		})
	}
}
