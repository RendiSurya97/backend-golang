package controller

import (
	"errors"
	"reflect"
	"testing"

	"github.com/backendgolang/searchapp/internal/movie/ext_module"
	mockCore "github.com/backendgolang/searchapp/mocks/internalPkg/movie/core"
	mockExtMod "github.com/backendgolang/searchapp/mocks/internalPkg/movie/ext_module"
	"github.com/golang/mock/gomock"
)

func Test_ctrl_GetByTitle(t *testing.T) {
	c := gomock.NewController(t)
	mockMovie := mockCore.NewMockMovie(c)
	mockOmdb := mockExtMod.NewMockIOmdb(c)
	movieCore = mockMovie
	omdbExt = mockOmdb

	type args struct {
		title string
		page  int64
	}
	tests := []struct {
		name       string
		c          *ctrl
		args       args
		wantResult MovieResult
		wantErr    bool
		mockFunc   func()
	}{
		{
			name: "valid",
			args: args{
				title: "spiderman",
				page:  1,
			},
			wantResult: MovieResult{
				TotalResults: "3",
				Response:     "True",
				MovieDetails: []MovieDetail{
					{
						Title:  "Spiderman 1",
						Year:   "2000",
						ImdbID: "1",
						Type:   "movie",
					},
					{
						Title:  "Spiderman 2",
						Year:   "2001",
						ImdbID: "2",
						Type:   "movie",
					},
					{
						Title:  "Spiderman 3",
						Year:   "2002",
						ImdbID: "3",
						Type:   "movie",
					},
				},
			},
			wantErr: false,
			mockFunc: func() {
				mockOmdb.EXPECT().GetMovieByTitle("spiderman", int64(1)).Return(ext_module.MovieResult{
					TotalResults: "3",
					Response:     "True",
					MovieDetails: []ext_module.MovieDetail{
						{
							Title:  "Spiderman 1",
							Year:   "2000",
							ImdbID: "1",
							Type:   "movie",
						},
						{
							Title:  "Spiderman 2",
							Year:   "2001",
							ImdbID: "2",
							Type:   "movie",
						},
						{
							Title:  "Spiderman 3",
							Year:   "2002",
							ImdbID: "3",
							Type:   "movie",
						},
					},
				}, nil)

				mockMovie.EXPECT().SaveLog("spiderman").Return(nil)
			},
		},
		{
			name: "error get movie",
			args: args{
				title: "spiderman",
				page:  1,
			},
			wantResult: MovieResult{
				TotalResults: "",
				Response:     "",
				MovieDetails: nil,
			},
			wantErr: true,
			mockFunc: func() {
				mockOmdb.EXPECT().GetMovieByTitle("spiderman", int64(1)).Return(ext_module.MovieResult{}, errors.New("error happen"))

				mockMovie.EXPECT().SaveLog("spiderman").Return(nil)
			},
		},
		{
			name: "error save log",
			args: args{
				title: "spiderman",
				page:  1,
			},
			wantResult: MovieResult{
				TotalResults: "3",
				Response:     "True",
				MovieDetails: []MovieDetail{
					{
						Title:  "Spiderman 1",
						Year:   "2000",
						ImdbID: "1",
						Type:   "movie",
					},
					{
						Title:  "Spiderman 2",
						Year:   "2001",
						ImdbID: "2",
						Type:   "movie",
					},
					{
						Title:  "Spiderman 3",
						Year:   "2002",
						ImdbID: "3",
						Type:   "movie",
					},
				},
			},
			wantErr: false,
			mockFunc: func() {
				mockOmdb.EXPECT().GetMovieByTitle("spiderman", int64(1)).Return(ext_module.MovieResult{
					TotalResults: "3",
					Response:     "True",
					MovieDetails: []ext_module.MovieDetail{
						{
							Title:  "Spiderman 1",
							Year:   "2000",
							ImdbID: "1",
							Type:   "movie",
						},
						{
							Title:  "Spiderman 2",
							Year:   "2001",
							ImdbID: "2",
							Type:   "movie",
						},
						{
							Title:  "Spiderman 3",
							Year:   "2002",
							ImdbID: "3",
							Type:   "movie",
						},
					},
				}, nil)

				mockMovie.EXPECT().SaveLog("spiderman").Return(errors.New("error happen"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ctrl{}
			tt.mockFunc()
			gotResult, err := c.GetByTitle(tt.args.title, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("ctrl.GetByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ctrl.GetByTitle() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
