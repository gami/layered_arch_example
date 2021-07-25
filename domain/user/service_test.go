package user_test

import (
	"context"
	"testing"

	"app/domain/user"

	"github.com/google/go-cmp/cmp"
)

func Test_service_FindByID(t *testing.T) {
	type args struct {
		id user.ID
	}

	tests := []struct {
		name    string
		repo    user.Repository
		args    args
		want    *user.User
		wantErr bool
	}{
		{
			name: "Found",
			repo: &MockRepository{
				FindByIDFunc: func(ctx context.Context, id user.ID) (*user.User, error) {
					return &user.User{
						ID:   1,
						Name: "gami",
					}, nil
				},
			},
			args: args{
				id: 1,
			},
			want: &user.User{
				ID:   1,
				Name: "gami",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			s := user.NewService(tt.repo, nil)
			got, err := s.FindByID(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.FindByID() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("service.FindByID() is not match (-got +want):\n%s", diff)
			}
		})
	}
}
