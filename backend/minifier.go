package backend

import (
	"bufio"
	"bytes"
	"github.com/tdewolff/minify/v2"
	_ "github.com/tdewolff/minify/v2/css"
	_ "github.com/tdewolff/minify/v2/js"
	"os"
)

func Minifier(fp, mime string, dir ...string) error {
	buf := bytes.NewBuffer(nil)
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	c := make(chan []byte)
	e := make(chan error)
	n, _ := f.Stat()
	go walkDir(dir, n.Name(), c, e)

	min := minify.New()
	m, _, h := min.Match(mime)
	min.Add(m, h)
wait:
	for {
		select {
		case b, ok := <-c:
			_, err = buf.Write(b)
			if err != nil {
				return err
			}
			if !ok {
				break wait
			}
		case err = <-e:
			panic(err.Error())
		}
	}

	tmp, err := min.Bytes(m, buf.Bytes())
	if err != nil {
		return err
	}
	if _, err = w.Write(tmp); err != nil {
		return err
	}
	w.Flush()
	return nil
}

func walkDir(dir []string, ex string, b chan []byte, e chan error) {
	defer close(b)
	if len(dir) < 1 {
		return
	}

	for _, d := range dir {
		info, err := os.Stat(d)
		if err != nil {
			e <- err
		}
		if info.Name() == ex {
			continue
		}
		if info.IsDir() {
			nd, _ := os.ReadDir(d)
			p := handleDir(d, ex, nd)

			c := make(chan []byte)
			e2 := make(chan error)
			go walkDir(p, ex, b, e)
		wait2:
			for {
				select {
				case x, ok := <-c:
					b <- x
					if !ok {
						break wait2
					}
				case err = <-e2:
					e <- err
					return
				}
			}
			continue
		}

		tmp, err := os.ReadFile(d)
		if err != nil {
			e <- err
		}
		b <- tmp
	}
}

func handleDir(d, ex string, c []os.DirEntry) (paths []string) {
	for _, de := range c {
		if de.Name() == ex {
			continue
		}
		p := d + "/" + de.Name()
		paths = append(paths, p)
	}
	return
}
