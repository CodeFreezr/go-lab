/* Copyright (C) 2018 by Detlef Burkhardt */

/* This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

func format(n int64) string {
	in := strconv.FormatInt(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func du(currentPath string, info os.FileInfo) int64 {

	size := info.Size()
	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(currentPath)
	if err != nil {
		log.Print(err)
		return size
	}
	defer dir.Close()

	fis, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fis {
		if fi.Name() == "." || fi.Name() == ".." {
			continue
		}
		size += du(currentPath+"/"+fi.Name(), fi)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 14, 8, 0, '-', tabwriter.AlignRight)

	t := format(size) + "\t " + currentPath
	fmt.Fprintln(w, t)

	w.Flush()

	return size
}

func main() {
	log.SetFlags(log.Lshortfile)
	dir := os.Args[1]

	info, err := os.Lstat(dir)
	if err != nil {
		log.Fatal(err)
	}

	du(dir, info)
}
