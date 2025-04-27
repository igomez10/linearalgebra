package main

import (
	"image"
	"image/color"
	"testing"
)

func TestDraw2DVector(t *testing.T) {
	type args struct {
		x     float64
		y     float64
		img   *image.RGBA
		color *color.RGBA
	}
	tests := []struct {
		name   string
		args   args
		expect func(*image.RGBA) error
	}{
		{
			name: "simple case x=1 y=1",
			args: args{
				x:     1,
				y:     1,
				img:   image.NewRGBA(image.Rect(0, 0, 1, 1)),
				color: &blackColor,
			},
			expect: func(img *image.RGBA) error {
				img.Set(0, 1, whiteColor)
				img.Set(1, 0, whiteColor)
				img.Set(0, 0, blackColor)
				img.Set(1, 1, blackColor)
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Draw2DVector(tt.args.x, tt.args.y, tt.args.img, tt.args.color)
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		imgA *image.RGBA
		imgB *image.RGBA
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty image",
			args: args{
				imgA: image.NewRGBA(image.Rect(0, 0, 0, 0)),
				imgB: image.NewRGBA(image.Rect(0, 0, 0, 0)),
			},
			want: true,
		},
		{
			name: "single pixel is black",
			args: args{
				imgA: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, blackColor)
					return img
				}(),
				imgB: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, blackColor)
					return img
				}(),
			},
			want: true,
		},
		{
			name: "single pixel is black in one image",
			args: args{
				imgA: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, blackColor)
					return img
				}(),
				imgB: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					return img
				}(),
			},
			want: false,
		},
		{
			name: "different dimensions",
			args: args{
				imgA: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, blackColor)
					return img
				}(),
				imgB: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 2, 2))
					return img
				}(),
			},
			want: false,
		},
		{
			name: "last color wins",
			args: args{
				imgA: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, redColor)
					img.Set(0, 0, grayColor)
					img.Set(0, 0, blackColor)
					return img
				}(),
				imgB: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, blackColor)
					return img
				}(),
			},
			want: true,
		},
		{
			name: "different pixels but same color",
			args: args{
				imgA: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 2, 2))
					img.Set(0, 0, blackColor)
					return img
				}(),
				imgB: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 2, 2))
					img.Set(1, 1, blackColor)
					return img
				}(),
			},
			want: false,
		},
		{
			name: "different colors",
			args: args{
				imgA: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, redColor)
					return img
				}(),
				imgB: func() *image.RGBA {
					img := image.NewRGBA(image.Rect(0, 0, 1, 1))
					img.Set(0, 0, grayColor)
					return img
				}(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.imgA, tt.args.imgB); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualColor(t *testing.T) {
	type args struct {
		colorA color.Color
		colorB color.Color
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same color",
			args: args{
				colorA: blackColor,
				colorB: blackColor,
			},
			want: true,
		},
		{
			name: "different colors",
			args: args{
				colorA: blackColor,
				colorB: redColor,
			},
			want: false,
		},
		{
			name: "almost same color",
			args: args{
				colorA: color.RGBA{0, 0, 0, 0},
				colorB: color.RGBA{1, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "same red but different",
			args: args{
				colorA: color.RGBA{1, 1, 0, 0},
				colorB: color.RGBA{1, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "same green but different",
			args: args{
				colorA: color.RGBA{1, 1, 0, 0},
				colorB: color.RGBA{0, 1, 0, 0},
			},
			want: false,
		},
		{
			name: "same blue but different",
			args: args{
				colorA: color.RGBA{0, 1, 1, 0},
				colorB: color.RGBA{0, 0, 1, 0},
			},
			want: false,
		},
		{
			name: "same alpha but different",
			args: args{
				colorA: color.RGBA{0, 0, 1, 1},
				colorB: color.RGBA{0, 0, 0, 1},
			},
			want: false,
		},
		{
			name: "different alpha",
			args: args{
				colorA: color.RGBA{0, 0, 0, 1},
				colorB: color.RGBA{0, 0, 0, 2},
			},
			want: false,
		},
		{
			name: "completly different",
			args: args{
				colorA: color.RGBA{1, 2, 3, 4},
				colorB: color.RGBA{5, 6, 7, 8},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualColor(tt.args.colorA, tt.args.colorB); got != tt.want {
				t.Errorf("EqualColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
