package authappl

import (
	"campusconnect-api/internal/domain/auth"
	"context"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_userApplicationImpl_Create(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		e *auth.Entity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.Entity
		wantErr bool
	}{
		{
			name: "Create user Sucess",
			fields: fields{
				ctx: context.TODO(),
			},
			args: args{
				&auth.Entity{
					Name:     "junior",
					Email:    "junior@email.com",
					Password: "123456",
				},
			},
			want: &auth.Entity{
				Name:  "junior",
				Email: "junior@email.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authApplicationImpl{
				ctx: tt.fields.ctx,
			}
			got, err := s.Create(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("userApplicationImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("userApplicationImpl.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userApplicationImpl_validatePassword(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		hashPassword string
		password     string
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Validate Sucess",
			fields: fields{
				ctx: context.TODO(),
			},
			args: args{
				hashPassword: string(hash),
				password:     "123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authApplicationImpl{
				ctx: tt.fields.ctx,
			}
			if err := s.validatePassword(tt.args.hashPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("userApplicationImpl.validatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
