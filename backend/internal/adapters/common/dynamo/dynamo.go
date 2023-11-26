package dynamo

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/realfabecker/wallet/internal/adapters/common/lock"
)

// DecodedPageToken
type DecodedPageToken struct {
	PK struct {
		Value string `json:"Value"`
	} `json:"PK"`
	SK struct {
		Value string `json:"Value"`
	} `json:"SK"`
}

// EncodePageToken
func EncodePageToken(atts map[string]types.AttributeValue, key string) (string, error) {
	lsk, err := json.Marshal(atts)
	if err != nil {
		return "", err
	}
	crp, err := lock.Encrypt(lsk, key)
	if err != nil {
		return "", err
	}
	return lock.Base64Encode(crp), nil
}

// DecodePageToken
func DecodePageToken(atts string, key string) (map[string]types.AttributeValue, error) {
	var k DecodedPageToken
	if atts == "" {
		return nil, nil
	}

	token, err := lock.Base64Decode(atts)
	if err != nil {
		return nil, err
	}

	dtoken, err := lock.Decrypt(token, key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dtoken, &k); err != nil {
		return nil, err
	}

	t := map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{
			Value: k.PK.Value,
		},
		"SK": &types.AttributeValueMemberS{
			Value: k.SK.Value,
		},
	}
	return t, nil
}
