package main

import (
	"strings"
	"testing"
)

func simplifyPath(path string) string {
	segments := strings.Split(path, "/")
	ans := make([]string, 0)
	for _, seg := range segments {
		if seg == "." || seg == "" {
			continue
		} else if seg == ".." {
			if len(ans) >= 1 {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans = append(ans, seg)
		}
	}

	ret := "/" + strings.Join(ans, "/")
	return ret
}

func TestSimplifyPath(t *testing.T) {
	if simplifyPath("/home/") != "/home" {
		t.Errorf("TEST 1 FAILED")
	}

	if simplifyPath("/home//foo/") != "/home/foo" {
		t.Errorf("TEST 2 FAILED")
	}

	if simplifyPath("/home/user/Documents/../Pictures") != "/home/user/Pictures" {
		t.Errorf("TEST 3 FAILED")
	}

	if simplifyPath("/../") != "/" {
		t.Errorf("TEST 4 FAILED")
	}

	if simplifyPath("/.../a/../b/c/../d/./") != "/.../b/d" {
		t.Errorf("TEST 5 FAILED")
	}

	if simplifyPath("/a/./b/../../c/") != "/c" {
		t.Errorf("TEST 5 FAILED")
	}

}
