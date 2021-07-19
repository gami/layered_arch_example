package user_test

import (
	"context"
	"testing"

	"github.com/gami/layered_arch_example/domain/user"
	"github.com/google/go-cmp/cmp"
)

func Test_service_FindByID(t *testing.T) {
	type args struct {
		id uint64
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
				FindByIDFunc: func(ctx context.Context, id uint64) (*user.User, error) {
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
			s := user.NewService(tt.repo)
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