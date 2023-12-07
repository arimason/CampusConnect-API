package productappl

import (
	"campusconnect-api/internal/domain/product"
	"context"
	"testing"
)

func Test_productApplicationImpl_Create(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		e *product.Entity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *product.Entity
		wantErr bool
	}{
		{
			name: "Create Product Sucess",
			fields: fields{
				ctx: context.TODO(),
			},
			args: args{
				&product.Entity{
					Name:  "teste",
					Price: 27.30,
				},
			},
			want: &product.Entity{
				Name:  "teste",
				Price: 27.30,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &productApplicationImpl{
				ctx: tt.fields.ctx,
			}
			got, err := s.Create(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("productApplicationImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.Name != got.Name || tt.want.Price != got.Price || got.ID == "" {
				t.Errorf("productApplicationImpl.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
