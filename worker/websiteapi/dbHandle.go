package websiteapi

import (
	"context"
	"fmt"
)

type types struct {
	TypeId   int    `json:"typeId"`
	TypeName string `json:"typeName"`
}

func dbGetTypes(ctx context.Context) ([]*types, error) {
	cursor, err := db.Client.Query(ctx,
		`
			FOR doc IN Coolpc
			COLLECT typeId = doc.typeId, typeName = doc.typeName
			RETURN { typeId:typeId, typeName:typeName }
		`,
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	payloads := []*types{}
	for {
		doc := &types{}
		_, err = cursor.ReadDocument(ctx, doc)
		if err != nil {
			break
		}

		payloads = append(payloads, doc)
	}

	return payloads, nil
}

type items struct {
	TypeId     int    `json:"typeId"`
	UpdateTime int    `json:"updateTime"`
	Price      int    `json:"price"`
	TypeName   string `json:"typeName"`
	Name       string `json:"name"`
	PriceTag   string `json:"priceTag"`
	Date       string `json:"date"`
}

func dbGetItems(ctx context.Context, typeId int) ([]string, error) {
	cursor, err := db.Client.Query(ctx,
		`
			FOR doc IN Coolpc
			FILTER doc.typeId == @typeId
			COLLECT Name = doc.name
			RETURN Name
		`,
		map[string]interface{}{
			"typeId": typeId,
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	payloads := []string{}
	for {
		var doc string
		_, err = cursor.ReadDocument(ctx, &doc)
		if err != nil {
			break
		}

		payloads = append(payloads, doc)
	}

	return payloads, nil
}

func dbGetNameLike(ctx context.Context, typeId int, nameLike string) ([]*items, error) {
	nameLike = fmt.Sprintf("%%%v%%", nameLike)
	cursor, err := db.Client.Query(ctx,
		`
			FOR doc IN Coolpc
			FILTER doc.typeId == @typeId && doc.name like @nameLike
			RETURN doc
		`,
		map[string]interface{}{
			"typeId":   typeId,
			"nameLike": nameLike,
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	payloads := []*items{}
	for {
		doc := &items{}
		_, err = cursor.ReadDocument(ctx, &doc)
		if err != nil {
			break
		}

		payloads = append(payloads, doc)
	}

	return payloads, nil
}
