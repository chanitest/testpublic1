package featureflag

//import (
//	"context"
//	"fmt"
//	"github.com/checkmarxDev/sca-grpc/connmanager"
//	"strconv"
//	"time"
//
//	"github.com/CheckmarxDev/feature-flag/pkg/api/grpc/featureflag"
//	lumologgerabs "github.com/checkmarxDev/lumologger/abstraction"
//	grpcclient "github.com/checkmarxDev/sca-grpc/client"
//	"github.com/pkg/errors"
//	"google.golang.org/grpc"
//)
//
//type GrpcFeatureFlagFetch struct {
//	client *grpcclient.GrpcClient[featureflag.FeatureFlagServiceClient, featureflag.GetFlagRequest, featureflag.GetFlagResponse]
//}
//
//func NewGrpcFeatureFlagFetcher(ctx context.Context, address string, opts []grpc.DialOption, timeout time.Duration, logger lumologgerabs.Logger) (*GrpcFeatureFlagFetch, error) {
//	err := connmanager.AddConnectionToScaGrpcConnectionFactory(ctx, logger, featureflag.NewFeatureFlagServiceClient, address, timeout, opts...)
//	if err != nil {
//		return nil, err
//	}
//	return &GrpcFeatureFlagFetch{client: grpcclient.NewGrpcClient[featureflag.FeatureFlagServiceClient, featureflag.GetFlagRequest, featureflag.GetFlagResponse](ctx, logger, timeout)}, nil
//}
//
//func (f *GrpcFeatureFlagFetch) GetFeatureFlag(ctx context.Context, tenantID, flagName string, logger lumologgerabs.Logger) (bool, error) {
//	logger.AddEnrichersFromContext(ctx)
//
//	ffClient, err := f.client.GetClientFromClientInterface(featureflag.NewFeatureFlagServiceClient)
//	if err != nil {
//		return false, errors.Wrap(err, "failed connect to feature flag grpc")
//	}
//
//	response, err := f.client.Execute(ffClient.GetFlag, &featureflag.GetFlagRequest{Name: flagName, Filter: &tenantID})
//
//	if err != nil {
//		return false, errors.Wrap(err, fmt.Sprintf("error calling FeatureFlagService.GetFlag grpc with flag name: %s", flagName))
//	}
//
//	logger.DebugWithAdditionalInfo("Feature flag %s status is: %t", map[string]string{
//		"FlagName": flagName, "TenantID": tenantID, "Status": strconv.FormatBool(response.Flag.Status)}, flagName, response.Flag.Status)
//
//	return response.Flag.Status, nil
//}
