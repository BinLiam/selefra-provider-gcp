package gcp_client

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExtractorProject() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).ProjectId, nil
	})
}

func ExtractorProtoTimestamp(path string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		v, err := column_value_extractor.StructSelector(path).Extract(ctx, clientMeta, client, task, row, column, result)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		ts, ok := v.(*timestamppb.Timestamp)
		if !ok {
			return nil, schema.NewDiagnosticsAddErrorMsg(fmt.Sprintf("unextected type, wanted \"*timestamppb.Timestamp\", have \"%T\"", v))
		}
		return ts.AsTime(), nil
	})
}

func ExtractorProtoEtag() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		v, err := column_value_extractor.StructSelector("Etag").Extract(ctx, clientMeta, client, task, row, column, result)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		switch data := v.(type) {
		case []uint8:
			return base64.StdEncoding.EncodeToString(data), nil
		default:
			return data, nil
		}
	})
}

func ExtractorOrganization() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).OrgId, nil
	})
}
