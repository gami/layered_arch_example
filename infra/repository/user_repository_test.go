package repository_test

import (
	"context"
	"testing"

	"app/domain/user"
	"app/infra/repository"

	"github.com/google/go-cmp/cmp"
)

func TestUser_FindByID(t *testing.T) {
	type args struct {
		id user.ID
	}
	tests := []struct {
		name    string
		args    args
		want    *user.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			r := repository.NewUser(testDB)
			got, err := r.FindByID(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("User.FindByID() is not match (-got +want):\n%s", diff)
			}
		})
	}
}
