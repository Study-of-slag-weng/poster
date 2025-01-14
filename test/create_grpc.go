package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/Study-of-slag-weng/poster/proto"
	grpc "google.golang.org/grpc"
)

// 测试grpc调用
func main() {
	conn, err := grpc.Dial("127.0.0.1:10280", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewPosterClient(conn)

	// 图片使用file
	// testByFile(client)

	// 图片使用url
	testByUrl(client)
}

// 图片信息放在请求体
func testByFile(client proto.PosterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 读取背景图片
	backgroundImg, err := ioutil.ReadFile("./background.jpg")
	if err != nil {
		log.Println(err)
		return
	}
	// 二维码
	qrcode, err := ioutil.ReadFile("./qrcode.png")
	if err != nil {
		log.Println(err)
		return
	}
	// 文本
	texts := make([]*proto.Text, 0)
	texts = append(texts, &proto.Text{
		Top:        176,
		Left:       166,
		Width:      600,
		Height:     56,
		LineCount:  20,
		Content:    "爱 就 大 声 说 出 来",
		FontSize:   50,
		LineHeight: 1.5,
		FontColor:  "#ff0000",
	})
	// 子图片
	subImages := make([]*proto.Image, 0)
	subImages = append(subImages, &proto.Image{
		Top:       940,
		Left:      23,
		Width:     170,
		Height:    170,
		Padding:   30,
		ImageType: "png",
		Image:     qrcode,
	})
	// 调用
	rsp, err := client.CreatePoster(ctx, &proto.CreatePosterRequest{
		Width:  720,
		Height: 1280,
		Background: &proto.Background{
			Image: backgroundImg,
		},
		Texts:     texts,
		SubImages: subImages,
	})
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("out.jpg", rsp.GetImage(), 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

// 图片信息为url
func testByUrl(client proto.PosterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// 文本
	texts := make([]*proto.Text, 0)
	texts = append(texts, &proto.Text{
		Top:        176,
		Left:       166,
		Width:      600,
		Height:     56,
		LineCount:  20,
		Content:    "爱 就 大 声 说 出 来",
		FontSize:   50,
		LineHeight: 1.5,
		FontColor:  "#ff0000",
	})
	// 子图片
	subImages := make([]*proto.Image, 0)
	subImages = append(subImages, &proto.Image{
		Top:       940,
		Left:      23,
		Width:     170,
		Height:    170,
		Padding:   30,
		ImageType: "png",
		ImageUrl:  "http://127.0.0.1:8000/qrcode.png",
	})
	// 二维码
	qrCode := make([]*proto.QrCode, 0)
	qrCode = append(qrCode, &proto.QrCode{
		Top:     660,
		Left:    40,
		Width:   150,
		Angle:   60,
		Content: "https://github.com/Study-of-slag-weng",
	})
	// 小程序码
	wxQrCode := make([]*proto.WxQrCode, 0)
	wxQrCode = append(wxQrCode, &proto.WxQrCode{
		Top:         360,
		Left:        40,
		Width:       130,
		Angle:       60,
		AccessToken: "21_0srT_h4jGvBGst77nQYgA9gIq39unyPUKIvnEiRncHmll3SDdVZUT9JElRX15g8OUr4TRIsBvOxrk-yFvX5w1_Q6nOmIcta3UPyBFT2IfCjjIR12DUT5niz1kw7YFq5E5XSaxtycUzgpdYoYKEJgAJAGAF",
		Scene:       "hehe",
	})
	// 调用
	rsp, err := client.CreatePoster(ctx, &proto.CreatePosterRequest{
		Width:  720,
		Height: 1280,
		Background: &proto.Background{
			ImageUrl: "http://127.0.0.1:8000/background.jpg",
		},
		Texts:       texts,
		SubImages:   subImages,
		SubQrCode:   qrCode,
		SubWxQrCode: wxQrCode,
	})
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("out.jpg", rsp.GetImage(), 0644)
	if err != nil {
		log.Println(err)
		return
	}
}
