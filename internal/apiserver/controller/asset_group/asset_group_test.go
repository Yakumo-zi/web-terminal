package asset_group

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

// MockAssetGroupService 是 AssetGroupService 的 mock 实现
type MockAssetGroupService struct {
	mock.Mock
}

func (m *MockAssetGroupService) Create(ctx context.Context, group *domain.AssetGroup) error {
	args := m.Called(ctx, group)
	return args.Error(0)
}

func (m *MockAssetGroupService) Update(ctx context.Context, group *domain.AssetGroup) error {
	args := m.Called(ctx, group)
	return args.Error(0)
}

func (m *MockAssetGroupService) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockAssetGroupService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *MockAssetGroupService) Get(ctx context.Context, id uuid.UUID) (*domain.AssetGroup, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.AssetGroup), args.Error(1)
}

func (m *MockAssetGroupService) List(ctx context.Context, options *service.ListOptions) ([]*domain.AssetGroup, int, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]*domain.AssetGroup), args.Int(1), args.Error(2)
}

func (m *MockAssetGroupService) AddMembers(ctx context.Context, groupID uuid.UUID, memberIDs []uuid.UUID) error {
	args := m.Called(ctx, groupID, memberIDs)
	return args.Error(0)
}

// 创建测试用的 Echo 实例
func setupEcho() *echo.Echo {
	e := echo.New()
	return e
}

// 创建测试用的 Service mock
func createMockService(mockAssetGroupService *MockAssetGroupService) *service.Service {
	return &service.Service{
		AssetGroupService: mockAssetGroupService,
	}
}

func TestAssetGroupController_Create(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	tests := []struct {
		name           string
		requestBody    CreateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功创建资产组",
			requestBody: CreateRequest{
				Name: "Test Group",
			},
			mockSetup: func() {
				mockService.On("Create", mock.Anything, mock.AnythingOfType("*domain.AssetGroup")).Return(nil)
			},
			expectedStatus: 200,
		},
		{
			name: "空名称",
			requestBody: CreateRequest{
				Name: "",
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
			req := httptest.NewRequest(http.MethodPost, "/Api/V1/AssetGroups", bytes.NewReader(body))
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

func TestAssetGroupController_Get(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	groupID := uuid.New()
	mockGroup := &domain.AssetGroup{
		Id:        groupID,
		Name:      "Test Group",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name           string
		groupID        string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:    "成功获取资产组",
			groupID: groupID.String(),
			mockSetup: func() {
				mockService.On("Get", mock.Anything, groupID).Return(mockGroup, nil)
			},
			expectedStatus: 200,
		},
		{
			name:           "无效的UUID",
			groupID:        "invalid-uuid",
			mockSetup:      func() {},
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/AssetGroups/%s", tt.groupID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.groupID)
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

func TestAssetGroupController_List(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	mockGroups := []*domain.AssetGroup{
		{
			Id:   uuid.New(),
			Name: "Group 1",
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功获取资产组列表",
			queryParams: map[string]string{
				"limit":  "10",
				"offset": "0",
			},
			mockSetup: func() {
				mockService.On("List", mock.Anything, mock.AnythingOfType("*service.ListOptions")).Return(mockGroups, 1, nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/Api/V1/AssetGroups", nil)
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
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetGroupController_Update(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	groupID := uuid.New()

	tests := []struct {
		name           string
		requestBody    UpdateRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功更新资产组",
			requestBody: UpdateRequest{
				Id:   groupID.String(),
				Name: "Updated Group",
			},
			mockSetup: func() {
				mockService.On("Update", mock.Anything, mock.AnythingOfType("*domain.AssetGroup")).Return(nil)
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
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/AssetGroups/%s", tt.requestBody.Id), bytes.NewReader(body))
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

func TestAssetGroupController_Delete(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	groupID := uuid.New()

	tests := []struct {
		name           string
		groupID        string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:    "成功删除资产组",
			groupID: groupID.String(),
			mockSetup: func() {
				mockService.On("Delete", mock.Anything, groupID).Return(nil)
			},
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			tt.mockSetup()

			// Create request
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/AssetGroups/%s", tt.groupID), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.groupID)

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

func TestAssetGroupController_DeleteCollection(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
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
			name: "成功批量删除资产组",
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
			req := httptest.NewRequest(http.MethodDelete, "/Api/V1/AssetGroups/Collection", bytes.NewReader(body))
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
			mockService.ExpectedCalls = nil
			mockService.Calls = nil
		})
	}
}

func TestAssetGroupController_AddMembers(t *testing.T) {
	// Setup
	mockService := new(MockAssetGroupService)
	mockSvc := createMockService(mockService)
	controller := NewController(mockSvc)
	e := setupEcho()

	groupID := uuid.New()
	memberIDs := []uuid.UUID{uuid.New(), uuid.New()}

	tests := []struct {
		name           string
		requestBody    AddMembersRequest
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "成功添加成员到资产组",
			requestBody: AddMembersRequest{
				GroupID:  groupID.String(),
				AssetIDs: []string{memberIDs[0].String(), memberIDs[1].String()},
			},
			mockSetup: func() {
				mockService.On("AddMembers", mock.Anything, groupID, memberIDs).Return(nil)
			},
			expectedStatus: 200,
		},
		{
			name: "空的资产ID列表",
			requestBody: AddMembersRequest{
				GroupID:  groupID.String(),
				AssetIDs: []string{},
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
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/AssetGroups/%s/AddMembers", tt.requestBody.GroupID), bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute
			err := controller.AddMembers(c)
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
