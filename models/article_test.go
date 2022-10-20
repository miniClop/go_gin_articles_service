package models

import (
	"reflect"
	"testing"
)

func Test_allArticles(t *testing.T) {
	tests := []struct {
		name string
		want []article
	}{
		{
			name: "Article test",
			want: []article{
				{1, "Article 1", "Article Content 1"},
				{2, "Article 2", "Article Content 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllArticles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
