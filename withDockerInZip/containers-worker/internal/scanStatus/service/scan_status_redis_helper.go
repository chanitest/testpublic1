package scan_status_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	"github.com/go-redis/redis/v8"
)

func UpdateScanStatusToRedis(ctx context.Context, client redis.UniversalClient, status scan_status_service.ScanStatus) error {
	key := fmt.Sprintf("%s:%s", scan_status_service.ContainersScanStatusKeyPrefix, status.Id)
	data, err := json.Marshal(status)
	if err != nil {
		return err
	}

	_, err = client.Set(ctx, key, data, 0).Result()

	return err
}

func GetScanStatusFromRedis(ctx context.Context, client redis.UniversalClient, scanId string) (*scan_status_service.ScanStatus, error) {
	key := fmt.Sprintf("%s:%s", scan_status_service.ContainersScanStatusKeyPrefix, scanId)
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if exists == 0 {
		return nil, nil
	}
	data, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var status scan_status_service.ScanStatus
	err = json.Unmarshal([]byte(data), &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
