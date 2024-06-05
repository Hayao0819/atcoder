package flist

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type options = struct {
	maxDepth int
	minDepth int
	fileOnly bool
	dirOnly  bool
	extOnly  string
	filename bool
	extsOnly []string
	relpath  bool
}

type Option func(*options)

func WithMaxDepth(depth int) Option {
	return func(opt *options) {
		opt.maxDepth = depth
	}
}

func WithMinDepth(depth int) Option {
	return func(opt *options) {
		opt.minDepth = depth
	}
}

func WithExactDepth(depth int) Option {
	return func(opt *options) {
		opt.maxDepth = depth
		opt.minDepth = depth
	}
}

func WithExtOnly(ext string) Option {
	return func(opt *options) {
		opt.extOnly = ext
	}
}

func WithExtsOnly(exts []string) Option {
	return func(opt *options) {
		opt.extsOnly = exts
	}
}

func WithFileOnly() func(*options) {
	return func(opt *options) {
		opt.fileOnly = true
	}
}

func WithFileName() Option {
	return func(opt *options) {
		opt.filename = true
	}
}

func WithDirOnly() Option {
	return func(opt *options) {
		opt.dirOnly = true
	}
}

func WithRelPath() Option {
	return func(opt *options) {
		opt.relpath = true
	}
}

func Get(dir string, opts ...Option) (*[]string, error) {
	opt := options{
		maxDepth: -1,
		minDepth: 0,
		fileOnly: false,
		dirOnly:  false,
		extOnly:  "",
		extsOnly: []string{},
	}

	for _, o := range opts {
		o(&opt)
	}

	rtn := []string{}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// handle error
		if err != nil {
			return err
		}

		// フィルター

		// ファイルのみでディレクトリは除外
		if opt.fileOnly && d.IsDir() {
			return nil
		}

		// ディレクトリのみでファイルは除外
		if opt.dirOnly && !d.IsDir() {
			return nil
		}

		// 階層の深さ
		depthBasis := len(strings.Split(dir, string(os.PathSeparator)))
		depth := len(strings.Split(path, string(os.PathSeparator)))
		depthDiff := depth - depthBasis
		if !(opt.maxDepth < 0) && depthDiff > opt.maxDepth {
			return nil
		}
		if depthDiff < opt.minDepth {
			return nil
		}

		// 拡張子
		if opt.extOnly != "" {
			ext := filepath.Ext(path)
			if ext != opt.extOnly && !slices.Contains(opt.extsOnly, ext) {
				return nil
			}
		}

		// 出力形式
		rtnPath := path
		if opt.relpath {
			rtnPath = strings.Replace(rtnPath, dir, "", 1)
		}
		if opt.filename {
			rtnPath = filepath.Base(rtnPath)
		}

		rtn = append(rtn, rtnPath)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
