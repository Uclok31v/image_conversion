package convert

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Cli struct {
	DirectoryPath string
	Before        *string
	After         *string
}

func (c *Cli) Execute() int {
	// 引数で指定したディレクトリ配下を走査
	err := filepath.Walk(c.DirectoryPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fmt.Sprintf(".%s", *c.Before) {
			err := convert(path, *c.After)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

// 画像形式を変換
func convert(path string, after string) error {
	// ファイルを開く
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ファイルを開けませんでした")
		return err
	}
	defer file.Close()

	// 画像をデコードする
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("画像のデコードに失敗しました")
		return err
	}

	// 出力ファイルを生成
	out, err := os.Create(getConvertedFilePath(path, after))
	if err != nil {
		fmt.Println("ファイルの作成に失敗しました")
		return err
	}
	defer out.Close()

	// 画像ファイル出力
	switch after {
	case "jpg", "jpeg":
		err := jpeg.Encode(out, img, nil)
		if err != nil {
			fmt.Println("画像のエンコードに失敗しました")
			return err
		}
	case "png":
		err := png.Encode(out, img)
		if err != nil {
			fmt.Println("画像のエンコードに失敗しました")
			return err
		}
	default:
		fmt.Println("画像のエンコードに失敗しました")
		return errors.New("Invalid format")
	}

	return nil
}

// 出力ファイルのパス
func getConvertedFilePath(path string, after string) string {
	return path[:len(path)-len(filepath.Ext(path))] + fmt.Sprintf(".%s", after)
}
