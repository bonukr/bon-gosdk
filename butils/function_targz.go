package butils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateTarGz(destFile string, srcDir string) error {
	// 출력 파일 생성
	outFile, err := os.Create(destFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	// gzip writer → tar writer 생성
	gzw := gzip.NewWriter(outFile)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	// srcDir 기준으로 모든 파일 순회
	return filepath.Walk(srcDir, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// tar 헤더 생성
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		// 경로 상대화 (압축 파일 내에 경로 보존)
		relPath, err := filepath.Rel(srcDir, file)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 헤더 기록
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// 디렉토리는 내용 없음
		if fi.IsDir() {
			return nil
		}

		// 파일 내용 복사
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(tw, f)
		return err
	})
}

// createTarGz creates a tar.gz archive with the given paths
func CreateTarGzFromList(output string, paths []string, useStopOnErro bool) error {
	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	gw := gzip.NewWriter(outFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, path := range paths {
		err = addToTar(tw, path)
		if err != nil {
			fmt.Printf("could't add to tar.gz: %s -> %s: %s", path, output, err)
			if useStopOnErro {
				return err
			} else {
				continue
			}
		}
	}
	return nil
}

// addToTar adds the file or directory to the tar writer
func addToTar(tw *tar.Writer, src string) error {
	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		// Ensure the header name uses the full path (not just basename)
		relPath := strings.TrimPrefix(file, "/")
		header.Name = relPath

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if fi.Mode().IsRegular() {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(tw, f)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func ExtractTarGz(tarGzPath, destDir string) error {
	// 파일 열기
	f, err := os.Open(tarGzPath)
	if err != nil {
		return fmt.Errorf("failed to open tar.gz file: %w", err)
	}
	defer f.Close()

	// gzip 리더
	gzr, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzr.Close()

	// tar 리더
	tr := tar.NewReader(gzr)

	// tar 파일 내부를 순회
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // 끝
		}
		if err != nil {
			return fmt.Errorf("error reading tar entry: %w", err)
		}

		targetPath := filepath.Join(destDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			// 디렉토리 먼저 생성
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return fmt.Errorf("failed to write file: %w", err)
			}
			outFile.Close()
		default:
			// 생략: 심볼릭 링크 등은 필요 시 추가 처리
			fmt.Printf("Skipping unsupported type: %s\n", header.Name)
		}
	}

	return nil
}
