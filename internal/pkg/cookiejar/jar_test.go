package cookiejar_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/toolset/fiddle/internal/pkg/cookiejar"
)

func TestJar(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("test"); err != nil {
			http.SetCookie(w, &http.Cookie{Name: "test", Value: "init"})
		} else {
			cookie.Value = "work"
			http.SetCookie(w, cookie)
		}
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	require.NoError(t, err)

	jar, err := New(new(Options).StoreToMemory().AsFile("testdata/cookies.json"))
	require.NoError(t, err)

	client := &http.Client{Jar: jar}
	for _, cookies := range [][]*http.Cookie{
		{{Name: "test", Value: "init"}},
		{{Name: "test", Value: "work"}},
	} {
		_, err = client.Get(u.String())
		assert.NoError(t, err)
		assert.Equal(t, cookies, jar.Cookies(u))
	}
	require.NoError(t, jar.Dump())
}
