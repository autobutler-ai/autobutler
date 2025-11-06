package util

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// FilterPhotoFiles filters a list of files to only include photo files
func FilterPhotoFiles(files []fs.FileInfo) []fs.FileInfo {
	photoFiles := make([]fs.FileInfo, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileType := DetermineFileTypeFromPath(file.Name())
		if fileType == FileTypeImage {
			photoFiles = append(photoFiles, file)
		}
	}
	return photoFiles
}

// RecursivePhotoInfo stores a photo with its relative path
type RecursivePhotoInfo struct {
	FileInfo fs.FileInfo
	RelPath  string
}

// FindAllPhotosRecursively finds all photo files in a directory and its subdirectories
func FindAllPhotosRecursively(rootDir string) ([]RecursivePhotoInfo, error) {
	photos := make([]RecursivePhotoInfo, 0)

	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fileType := DetermineFileTypeFromPath(info.Name())
		if fileType == FileTypeImage {
			// Get relative path from rootDir
			relPath, err := filepath.Rel(rootDir, path)
			if err != nil {
				return err
			}
			photos = append(photos, RecursivePhotoInfo{
				FileInfo: info,
				RelPath:  relPath,
			})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory %s: %w", rootDir, err)
	}

	return photos, nil
}
