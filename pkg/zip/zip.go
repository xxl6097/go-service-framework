package zip

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Zip struct {
	*zip.Writer
}

func IsZip(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".zip":
		return true
	case ".gz":
		return true
	}
	return false
}

func UnPack(filename, dest string) (bool, error) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".zip":
		return true, UnzipFile(filename, dest)
	case ".gz":
		return true, UnTarGzFile(filename, dest)
	}
	return false, nil
}

func UnTarGzFile(filename, dest string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if dest == "" {
		dest = filepath.Dir(filename)
	}

	// 创建 gzip.Reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// 创建 tar.Reader
	tarReader := tar.NewReader(gzipReader)

	// 解压 tar 包中的文件
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // 读取完所有文件
		}
		if err != nil {
			return err
		}

		// 获取文件信息
		path := filepath.Join(dest, header.Name)
		info := header.FileInfo()

		// 创建目录或文件
		if info.IsDir() {
			if err := os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
		} else {
			dir := filepath.Dir(path)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, info.Mode())
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	log.Println("解压完成")
	return nil
}

func UnzipFile(filename, dest string) error {
	zr, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer zr.Close()

	if dest == "" {
		dest = filepath.Dir(filename)
	}

	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// ignore .ghs.yml
		filename := sanitizedName(f.Name)
		if filepath.Base(filename) == ".ghs.yml" {
			continue
		}
		fpath := filepath.Join(dest, filename)

		// filename maybe GBK or UTF-8
		// Ref: https://studygolang.com/articles/3114
		if f.Flags&(1<<11) == 0 { // GBK
			tr := simplifiedchinese.GB18030.NewDecoder()
			fpathUtf8, err := tr.String(fpath)
			if err == nil {
				fpath = fpathUtf8
			}
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

// GetRootDir 获取解压目录的根目录名称
func GetRootDir(dirPath string) (string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.HasSuffix(".", file.Name()) || strings.EqualFold(file.Name(), "__MACOSX") {
			continue
		}
		if file.IsDir() {
			return file.Name(), nil
		}
	}
	// 假设ZIP文件中只有一个根目录
	//if len(files) == 1 && files[0].IsDir() {
	//	return files[0].Name(), nil
	//}

	return "", fmt.Errorf("expected only one root directory")
}

func sanitizedName(filename string) string {
	if len(filename) > 1 && filename[1] == ':' &&
		runtime.GOOS == "windows" {
		filename = filename[2:]
	}
	filename = strings.TrimLeft(strings.Replace(filename, `\`, "/", -1), `/`)
	filename = filepath.ToSlash(filename)
	filename = filepath.Clean(filename)
	return filename
}

func statFile(filename string) (info os.FileInfo, reader io.ReadCloser, err error) {
	info, err = os.Lstat(filename)
	if err != nil {
		return
	}
	// content
	if info.Mode()&os.ModeSymlink != 0 {
		var target string
		target, err = os.Readlink(filename)
		if err != nil {
			return
		}
		reader = ioutil.NopCloser(bytes.NewBuffer([]byte(target)))
	} else if !info.IsDir() {
		reader, err = os.Open(filename)
		if err != nil {
			return
		}
	} else {
		reader = ioutil.NopCloser(bytes.NewBuffer(nil))
	}
	return
}

func (z *Zip) Add(relpath, abspath string) error {
	info, rdc, err := statFile(abspath)
	if err != nil {
		return err
	}
	defer rdc.Close()

	hdr, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	hdr.Name = sanitizedName(relpath)
	if info.IsDir() {
		hdr.Name += "/"
	}
	hdr.Method = zip.Deflate // compress method
	writer, err := z.CreateHeader(hdr)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, rdc)
	return err
}

func CompressToZip(w http.ResponseWriter, rootDir string) {
	rootDir = filepath.Clean(rootDir)
	zipFileName := filepath.Base(rootDir) + ".zip"

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="`+zipFileName+`"`)

	zw := &Zip{Writer: zip.NewWriter(w)}
	defer zw.Close()

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		zipPath := path[len(rootDir):]
		return zw.Add(zipPath, path)
	})
}
