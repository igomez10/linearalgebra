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
		expect func() *image.RGBA
	}{
		{
			name: "set origin to black",
			args: args{
				x:     0,
				y:     0,
				img:   image.NewRGBA(image.Rect(0, 0, 1, 1)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, blackColor)
				return img
			},
		},
		{
			name: "x=1 y=1",
			args: args{
				x:     1,
				y:     1,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX+1, originY-1, blackColor)
				return img
			},
		},
		{
			name: "x=2 y=2",
			args: args{
				x:     2,
				y:     2,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX+1, originY-1, blackColor)
				img.Set(originX+2, originY-2, blackColor)
				return img
			},
		},
		{
			name: "x=5 y=5",
			args: args{
				x:     5,
				y:     5,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX+1, originY-1, blackColor)
				img.Set(originX+2, originY-2, blackColor)
				img.Set(originX+3, originY-3, blackColor)
				img.Set(originX+4, originY-4, blackColor)
				img.Set(originX+5, originY-5, blackColor)
				return img
			},
		},
		{
			name: "x=-1 y=1",
			args: args{
				x:     -1,
				y:     1,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX-1, originY-1, blackColor)
				return img
			},
		},
		{
			name: "x=1 y=-1",
			args: args{
				x:     1,
				y:     -1,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX+1, originY+1, blackColor)
				return img
			},
		},
		{
			name: "x=-1 y=-1",
			args: args{
				x:     -1,
				y:     -1,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX-1, originY+1, blackColor)
				return img
			},
		},
		{
			name: "x=-5 y=-5",
			args: args{
				x:     -5,
				y:     -5,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX-1, originY+1, blackColor)
				img.Set(originX-2, originY+2, blackColor)
				img.Set(originX-3, originY+3, blackColor)
				img.Set(originX-4, originY+4, blackColor)
				img.Set(originX-5, originY+5, blackColor)
				return img
			},
		},
		{
			name: "x=0 y=-5",
			args: args{
				x:     0,
				y:     -5,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX, originY+1, blackColor)
				img.Set(originX, originY+2, blackColor)
				img.Set(originX, originY+3, blackColor)
				img.Set(originX, originY+4, blackColor)
				img.Set(originX, originY+5, blackColor)
				return img
			},
		},
		{
			name: "x=5 y=0",
			args: args{
				x:     5,
				y:     0,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX+1, originY, blackColor)
				img.Set(originX+2, originY, blackColor)
				img.Set(originX+3, originY, blackColor)
				img.Set(originX+4, originY, blackColor)
				img.Set(originX+5, originY, blackColor)
				return img
			},
		},
		{
			name: "x=-5 y=0",
			args: args{
				x:     -5,
				y:     0,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX-1, originY, blackColor)
				img.Set(originX-2, originY, blackColor)
				img.Set(originX-3, originY, blackColor)
				img.Set(originX-4, originY, blackColor)
				img.Set(originX-5, originY, blackColor)
				return img
			},
		},
		{
			name: "x=1 y=2",
			args: args{
				x:     1,
				y:     2,
				img:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				color: &blackColor,
			},
			expect: func() *image.RGBA {
				img := image.NewRGBA(image.Rect(0, 0, 100, 100))

				originX := img.Bounds().Max.X / 2
				originY := img.Bounds().Max.Y / 2

				img.Set(originX, originY, blackColor)
				img.Set(originX+1, originY-2, blackColor)
				img.Set(originX+2, originY-4, blackColor)
				return img
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Draw2DVector(tt.args.x, tt.args.y, tt.args.img, tt.args.color)
			if !Equal(tt.args.img, tt.expect()) {
				t.Errorf("expected \n%+v \nbut got \n%+v\n", tt.expect(), tt.args.img)
			}
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

func TestMax(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "a is bigger",
			args: args{
				a: 1,
				b: 0,
			},
			want: 1,
		},
		{
			name: "b is bigger",
			args: args{
				a: 0,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
