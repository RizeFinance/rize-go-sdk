package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// Expand paginated results
func (e Example) ExampleClientService_WithPagination(rc *rize.Client) {
	params := &rize.CustomerListParams{
		Limit:  100,
		Offset: 0,
		Sort:   "first_name_asc",
	}

	var (
		count int
		list  *rize.CustomerListResponse
		data  []*rize.Customer
	)

	for {
		resp, err := rc.Customers.List(context.Background(), params)
		if err != nil {
			log.Fatal("Error fetching customers\n", err)
		}

		data = append(data, resp.Data...)
		count += resp.Count

		if count >= resp.TotalCount {
			list = &rize.CustomerListResponse{
				ListResponse: rize.ListResponse{
					TotalCount: resp.TotalCount,
					Count:      count,
					Limit:      resp.Limit,
					Offset:     params.Offset,
				},
				Data: data,
			}
			break
		}

		params.Offset += count
	}

	output, _ := json.MarshalIndent(list, "", "\t")
	log.Println("Expanded Customer List:", string(output))
}
