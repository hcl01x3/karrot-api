package batch

import (
	"context"
	"fmt"

	dl "github.com/graph-gophers/dataloader/v6"
	"github.com/octo-5/karrot-api/model"
	"xorm.io/xorm"
)

func UserById(db *xorm.Engine) dl.BatchFunc {
	return func(ctx context.Context, ids dl.Keys) []*dl.Result {
		results := make([]*dl.Result, len(ids))

		var users []*model.User
		err := db.Context(ctx).In("id", ids).Find(&users)
		if err != nil {
			for i := range results {
				results[i] = &dl.Result{
					Error: err,
				}
			}
			return results
		}

		userById := map[string]*model.User{}
		for _, user := range users {
			id := fmt.Sprintf("%d", user.Id)
			userById[id] = user
		}

		for i, id := range ids {
			results[i] = &dl.Result{
				Data: userById[id.String()],
			}
		}
		return results
	}
}
