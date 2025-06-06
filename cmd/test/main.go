package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/asset"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/asset_group"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/credential"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/session"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/Yakumo-zi/web-terminal/pkg/logger"
	"github.com/Yakumo-zi/web-terminal/pkg/web/middlewares"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// random ip
func randomIP() string {
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	return ip
}
func randomName() string {
	return fmt.Sprintf("测试资产%d", rand.Intn(1000000))
}

func main() {
	svc := service.NewService()
	e := echo.New()
	e.Use(middlewares.LoggerWithSlog(logger.Log()))
	asset.RegisterRoutes(e, svc)
	asset_group.RegisterRoutes(e, svc)
	credential.RegisterRoutes(e, svc)
	session.RegisterRoutes(e, svc)

	// 测试 Asset
	fmt.Println("测试 Asset...")
	assetCreateReq := asset.CreateRequest{
		Name: randomName(),
		Ip:   randomIP(),
		Port: 22,
		Type: "host",
	}
	assetCreateJSON, _ := json.Marshal(assetCreateReq)
	assetCreateRec := httptest.NewRecorder()
	assetCreateReqObj := httptest.NewRequest(http.MethodPost, "/Api/V1/Assets", bytes.NewReader(assetCreateJSON))
	assetCreateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(assetCreateRec, assetCreateReqObj)
	if assetCreateRec.Code != http.StatusOK {
		log.Fatalf("创建资产失败: %s", assetCreateRec.Body.String())
	}
	logger.Log().Info("创建资产", "assetCreateReqObj", assetCreateRec.Body.String())
	fmt.Println("创建资产成功")
	var createAssetResp asset.CreateResponse
	json.Unmarshal(assetCreateRec.Body.Bytes(), &createAssetResp)
	assetID := uuid.MustParse(createAssetResp.Id)

	assetGetRec := httptest.NewRecorder()
	assetGetReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/Assets/%s", assetID), nil)
	logger.Log().Info("获取资产", "assetGetReq", assetGetReq.URL.String())
	e.ServeHTTP(assetGetRec, assetGetReq)
	if assetGetRec.Code != http.StatusOK {
		log.Fatalf("获取资产失败: %s", assetGetRec.Body.String())
	}
	logger.Log().Info("获取资产", "assetGetRec", assetGetRec.Body.String())
	fmt.Println("获取资产成功")

	assetListRec := httptest.NewRecorder()
	assetListReq := httptest.NewRequest(http.MethodGet, "/Api/V1/Assets?offset=0&limit=10", nil)
	e.ServeHTTP(assetListRec, assetListReq)
	if assetListRec.Code != http.StatusOK {
		log.Fatalf("获取资产列表失败: %s", assetListRec.Body.String())
	}
	fmt.Println("获取资产列表成功")

	assetUpdateReq := asset.UpdateRequest{
		Id:   assetID.String(),
		Name: randomName(),
		Ip:   randomIP(),
		Port: 22,
		Type: "host",
	}
	assetUpdateJSON, _ := json.Marshal(assetUpdateReq)
	assetUpdateRec := httptest.NewRecorder()
	assetUpdateReqObj := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/Assets/%s", assetID), bytes.NewReader(assetUpdateJSON))
	assetUpdateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(assetUpdateRec, assetUpdateReqObj)
	if assetUpdateRec.Code != http.StatusOK {
		log.Fatalf("更新资产失败: %s", assetUpdateRec.Body.String())
	}
	fmt.Println("更新资产成功")

	// assetDeleteRec := httptest.NewRecorder()
	// assetDeleteReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/Assets/%s", assetID), nil)
	// e.ServeHTTP(assetDeleteRec, assetDeleteReq)
	// if assetDeleteRec.Code != http.StatusOK {
	// 	log.Fatalf("删除资产失败: %s", assetDeleteRec.Body.String())
	// }
	// fmt.Println("删除资产成功")

	// assetDeleteCollectionReq := asset.DeleteCollectionRequest{
	// 	Ids: []string{assetID.String()},
	// }
	// assetDeleteCollectionJSON, _ := json.Marshal(assetDeleteCollectionReq)
	// assetDeleteCollectionRec := httptest.NewRecorder()
	// assetDeleteCollectionReqObj := httptest.NewRequest(http.MethodDelete, "/Api/V1/Assets/Collection", bytes.NewReader(assetDeleteCollectionJSON))
	// assetDeleteCollectionReqObj.Header.Set("Content-Type", "application/json")
	// e.ServeHTTP(assetDeleteCollectionRec, assetDeleteCollectionReqObj)
	// if assetDeleteCollectionRec.Code != http.StatusOK {
	// 	log.Fatalf("批量删除资产失败: %s", assetDeleteCollectionRec.Body.String())
	// }
	// fmt.Println("批量删除资产成功")

	// 测试 AssetGroup
	fmt.Println("\n测试 AssetGroup...")
	groupCreateReq := asset_group.CreateRequest{
		Name: randomName(),
	}
	groupCreateJSON, _ := json.Marshal(groupCreateReq)
	groupCreateRec := httptest.NewRecorder()
	groupCreateReqObj := httptest.NewRequest(http.MethodPost, "/Api/V1/AssetGroups", bytes.NewReader(groupCreateJSON))
	groupCreateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(groupCreateRec, groupCreateReqObj)
	if groupCreateRec.Code != http.StatusOK {
		log.Fatalf("创建组失败: %s", groupCreateRec.Body.String())
	}
	fmt.Println("创建组成功")
	var createGroupResp asset_group.CreateResponse
	json.Unmarshal(groupCreateRec.Body.Bytes(), &createGroupResp)
	groupID := uuid.MustParse(createGroupResp.Id)
	groupGetRec := httptest.NewRecorder()
	groupGetReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/AssetGroups/%s", groupID), nil)
	e.ServeHTTP(groupGetRec, groupGetReq)
	if groupGetRec.Code != http.StatusOK {
		log.Fatalf("获取组失败: %s", groupGetRec.Body.String())
	}
	fmt.Println("获取组成功")

	groupListRec := httptest.NewRecorder()
	groupListReq := httptest.NewRequest(http.MethodGet, "/Api/V1/AssetGroups?offset=0&limit=10", nil)
	e.ServeHTTP(groupListRec, groupListReq)
	if groupListRec.Code != http.StatusOK {
		log.Fatalf("获取组列表失败: %s", groupListRec.Body.String())
	}
	fmt.Println("获取组列表成功")

	groupUpdateReq := asset_group.UpdateRequest{
		Id:   groupID.String(),
		Name: randomName(),
	}
	groupUpdateJSON, _ := json.Marshal(groupUpdateReq)
	groupUpdateRec := httptest.NewRecorder()
	groupUpdateReqObj := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/AssetGroups/%s", groupID), bytes.NewReader(groupUpdateJSON))
	groupUpdateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(groupUpdateRec, groupUpdateReqObj)
	if groupUpdateRec.Code != http.StatusOK {
		log.Fatalf("更新组失败: %s", groupUpdateRec.Body.String())
	}
	fmt.Println("更新组成功")

	groupDeleteRec := httptest.NewRecorder()
	groupDeleteReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/AssetGroups/%s", groupID), nil)
	e.ServeHTTP(groupDeleteRec, groupDeleteReq)
	if groupDeleteRec.Code != http.StatusOK {
		log.Fatalf("删除组失败: %s", groupDeleteRec.Body.String())
	}
	fmt.Println("删除组成功")

	groupDeleteCollectionReq := asset_group.DeleteCollectionRequest{
		Ids: []string{groupID.String()},
	}
	groupDeleteCollectionJSON, _ := json.Marshal(groupDeleteCollectionReq)
	groupDeleteCollectionRec := httptest.NewRecorder()
	groupDeleteCollectionReqObj := httptest.NewRequest(http.MethodDelete, "/Api/V1/AssetGroups/Collection", bytes.NewReader(groupDeleteCollectionJSON))
	groupDeleteCollectionReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(groupDeleteCollectionRec, groupDeleteCollectionReqObj)
	if groupDeleteCollectionRec.Code != http.StatusOK {
		log.Fatalf("批量删除组失败: %s", groupDeleteCollectionRec.Body.String())
	}
	fmt.Println("批量删除组成功")

	// 测试 Credential
	fmt.Println("\n测试 Credential...")
	credCreateReq := credential.CreateRequest{
		Type:     "password",
		Username: randomName(),
		Secret:   "123456",
		AssetId:  assetID.String(),
	}
	credCreateJSON, _ := json.Marshal(credCreateReq)
	credCreateRec := httptest.NewRecorder()
	credCreateReqObj := httptest.NewRequest(http.MethodPost, "/Api/V1/Credentials", bytes.NewReader(credCreateJSON))
	credCreateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(credCreateRec, credCreateReqObj)
	if credCreateRec.Code != http.StatusOK {
		log.Fatalf("创建凭证失败: %s", credCreateRec.Body.String())
	}
	fmt.Println("创建凭证成功")
	var createCredResp credential.CreateResponse
	json.Unmarshal(credCreateRec.Body.Bytes(), &createCredResp)
	credID := uuid.MustParse(createCredResp.Id)

	credGetRec := httptest.NewRecorder()
	credGetReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/Credentials/%s", credID), nil)
	e.ServeHTTP(credGetRec, credGetReq)
	if credGetRec.Code != http.StatusOK {
		log.Fatalf("获取凭证失败: %s", credGetRec.Body.String())
	}
	fmt.Println("获取凭证成功")

	credListRec := httptest.NewRecorder()
	credListReq := httptest.NewRequest(http.MethodGet, "/Api/V1/Credentials?offset=0&limit=10", nil)
	e.ServeHTTP(credListRec, credListReq)
	if credListRec.Code != http.StatusOK {
		log.Fatalf("获取凭证列表失败: %s", credListRec.Body.String())
	}
	fmt.Println("获取凭证列表成功")

	credUpdateReq := credential.UpdateRequest{
		Id:       credID.String(),
		Type:     "password",
		Username: randomName(),
		Secret:   "654321",
	}
	credUpdateJSON, _ := json.Marshal(credUpdateReq)
	credUpdateRec := httptest.NewRecorder()
	credUpdateReqObj := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/Credentials/%s", credID), bytes.NewReader(credUpdateJSON))
	credUpdateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(credUpdateRec, credUpdateReqObj)
	if credUpdateRec.Code != http.StatusOK {
		log.Fatalf("更新凭证失败: %s", credUpdateRec.Body.String())
	}
	fmt.Println("更新凭证成功")

	// credDeleteRec := httptest.NewRecorder()
	// credDeleteReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/Credentials/%s", credID), nil)
	// e.ServeHTTP(credDeleteRec, credDeleteReq)
	// if credDeleteRec.Code != http.StatusOK {
	// 	log.Fatalf("删除凭证失败: %s", credDeleteRec.Body.String())
	// }
	// fmt.Println("删除凭证成功")

	// credDeleteCollectionReq := credential.DeleteCollectionRequest{
	// 	Ids: []string{credID.String()},
	// }
	// credDeleteCollectionJSON, _ := json.Marshal(credDeleteCollectionReq)
	// credDeleteCollectionRec := httptest.NewRecorder()
	// credDeleteCollectionReqObj := httptest.NewRequest(http.MethodDelete, "/Api/V1/Credentials/Collection", bytes.NewReader(credDeleteCollectionJSON))
	// credDeleteCollectionReqObj.Header.Set("Content-Type", "application/json")
	// e.ServeHTTP(credDeleteCollectionRec, credDeleteCollectionReqObj)
	// if credDeleteCollectionRec.Code != http.StatusOK {
	// 	log.Fatalf("批量删除凭证失败: %s", credDeleteCollectionRec.Body.String())
	// }
	// fmt.Println("批量删除凭证成功")

	// 测试 Session
	fmt.Println("\n测试 Session...")
	sessionCreateReq := session.CreateRequest{
		Type:    "ssh",
		Status:  "active",
		AssetId: assetID.String(),
		CredId:  credID.String(),
	}
	sessionCreateJSON, _ := json.Marshal(sessionCreateReq)
	sessionCreateRec := httptest.NewRecorder()
	sessionCreateReqObj := httptest.NewRequest(http.MethodPost, "/Api/V1/Sessions", bytes.NewReader(sessionCreateJSON))
	sessionCreateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(sessionCreateRec, sessionCreateReqObj)
	if sessionCreateRec.Code != http.StatusOK {
		log.Fatalf("创建会话失败: %s", sessionCreateRec.Body.String())
	}
	fmt.Println("创建会话成功")
	var createSessionResp session.CreateResponse
	json.Unmarshal(sessionCreateRec.Body.Bytes(), &createSessionResp)
	sessionID := uuid.MustParse(createSessionResp.Id)

	sessionGetRec := httptest.NewRecorder()
	sessionGetReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/Api/V1/Sessions/%s", sessionID), nil)
	e.ServeHTTP(sessionGetRec, sessionGetReq)
	if sessionGetRec.Code != http.StatusOK {
		log.Fatalf("获取会话失败: %s", sessionGetRec.Body.String())
	}
	fmt.Println("获取会话成功")

	sessionListRec := httptest.NewRecorder()
	sessionListReq := httptest.NewRequest(http.MethodGet, "/Api/V1/Sessions?offset=0&limit=10", nil)
	e.ServeHTTP(sessionListRec, sessionListReq)
	if sessionListRec.Code != http.StatusOK {
		log.Fatalf("获取会话列表失败: %s", sessionListRec.Body.String())
	}
	fmt.Println("获取会话列表成功")

	sessionUpdateReq := session.UpdateRequest{
		Id:     sessionID.String(),
		Type:   "ssh",
		Status: "inactive",
	}
	sessionUpdateJSON, _ := json.Marshal(sessionUpdateReq)
	sessionUpdateRec := httptest.NewRecorder()
	sessionUpdateReqObj := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/Api/V1/Sessions/%s", sessionID), bytes.NewReader(sessionUpdateJSON))
	sessionUpdateReqObj.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(sessionUpdateRec, sessionUpdateReqObj)
	if sessionUpdateRec.Code != http.StatusOK {
		log.Fatalf("更新会话失败: %s", sessionUpdateRec.Body.String())
	}
	fmt.Println("更新会话成功")

	// sessionDeleteRec := httptest.NewRecorder()
	// sessionDeleteReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/Api/V1/Sessions/%s", sessionID), nil)
	// e.ServeHTTP(sessionDeleteRec, sessionDeleteReq)
	// if sessionDeleteRec.Code != http.StatusOK {
	// 	log.Fatalf("删除会话失败: %s", sessionDeleteRec.Body.String())
	// }
	// fmt.Println("删除会话成功")

	// sessionDeleteCollectionReq := session.DeleteCollectionRequest{
	// 	Ids: []string{sessionID.String()},
	// }
	// sessionDeleteCollectionJSON, _ := json.Marshal(sessionDeleteCollectionReq)
	// sessionDeleteCollectionRec := httptest.NewRecorder()
	// sessionDeleteCollectionReqObj := httptest.NewRequest(http.MethodDelete, "/Api/V1/Sessions/Collection", bytes.NewReader(sessionDeleteCollectionJSON))
	// sessionDeleteCollectionReqObj.Header.Set("Content-Type", "application/json")
	// e.ServeHTTP(sessionDeleteCollectionRec, sessionDeleteCollectionReqObj)
	// if sessionDeleteCollectionRec.Code != http.StatusOK {
	// 	log.Fatalf("批量删除会话失败: %s", sessionDeleteCollectionRec.Body.String())
	// }
	// fmt.Println("批量删除会话成功")
}
