package model

import (
	"flag"
	"os"
	"testing"

	"github.com/frm-adiputra/qspec/examples/model/blog"
	"github.com/stretchr/testify/assert"
)

var blogFixtures = []struct{ ID, Title, Content string }{
	{"first", "My First Blog", "My First Blog content"},
	{"second", "My Second Blog", "My Second Blog content"},
	{"third", "My Third Blog", "My Third Blog content"},
}

func TestMain(m *testing.M) {
	flag.Parse()

	err := OpenDB(":memory:")
	if err != nil {
		panic(err)
	}

	err = InitDB()
	if err != nil {
		panic(err)
	}

	ret := m.Run()

	err = CloseDB()
	if err != nil {
		panic(err)
	}

	os.Exit(ret)
}

func TestBlogSelectByID(t *testing.T) {
	assert := assert.New(t)

	b, err := blog.SelectByID("first")
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.Equal("first", b.ID)
	assert.Equal("My First Blog", b.Title)
	assert.Equal("My First Blog content", b.Content)
}

func TestBlogPreparedSelectByID(t *testing.T) {
	assert := assert.New(t)

	stmt, err := blog.PrepareSelectByID()
	if !assert.NoError(err) {
		t.FailNow()
	}

	for _, fx := range blogFixtures {
		var b blog.SelectByIDResult
		b, err = stmt.Query(fx.ID)
		if !assert.NoError(err) {
			t.FailNow()
		}

		assert.Equal(fx.ID, b.ID)
		assert.Equal(fx.Title, b.Title)
		assert.Equal(fx.Content, b.Content)
	}

	err = stmt.Close()
	if !assert.NoError(err) {
		t.FailNow()
	}
}

func TestBlogAllTitles(t *testing.T) {
	assert := assert.New(t)

	a, err := blog.AllTitles()
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.Equal(3, len(a))

	for i, r := range a {
		assert.Equal(blogFixtures[i].ID, r.ID)
		assert.Equal(blogFixtures[i].Title, r.Title)
		assert.Equal(blogFixtures[i].Content, r.Content)
	}
}

func TestBlogPreparedAllTitles(t *testing.T) {
	assert := assert.New(t)

	stmt, err := blog.PrepareAllTitles()
	if !assert.NoError(err) {
		t.FailNow()
	}

	a, err := stmt.Query()
	if !assert.NoError(err) {
		t.FailNow()
	}

	err = stmt.Close()
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.Equal(3, len(a))

	for i, r := range a {
		assert.Equal(blogFixtures[i].ID, r.ID)
		assert.Equal(blogFixtures[i].Title, r.Title)
		assert.Equal(blogFixtures[i].Content, r.Content)
	}
}

func TestBlogUpdateComment(t *testing.T) {
	assert := assert.New(t)

	_, err := blog.UpdateComment("UsernameUpdated", "CommentUpdated", "1")
	if !assert.NoError(err) {
		t.FailNow()
	}

	v, err := blog.SelectCommentByID("1")
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.Equal("UsernameUpdated", v.Username)
	assert.Equal("CommentUpdated", v.Comment)
}

func TestBlogPreparedUpdateComment(t *testing.T) {
	assert := assert.New(t)

	stmt, err := blog.PrepareUpdateComment()
	if !assert.NoError(err) {
		t.FailNow()
	}

	_, err = stmt.Exec("UsernameUpdated", "CommentUpdated", "2")
	if !assert.NoError(err) {
		t.FailNow()
	}

	err = stmt.Close()
	if !assert.NoError(err) {
		t.FailNow()
	}

	v, err := blog.SelectCommentByID("2")
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.Equal("UsernameUpdated", v.Username)
	assert.Equal("CommentUpdated", v.Comment)
}

func TestBlogCountComment(t *testing.T) {
	assert := assert.New(t)

	n, err := blog.CountComment()
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.EqualValues(2, n)
}

func TestBlogPreparedCountComment(t *testing.T) {
	assert := assert.New(t)

	stmt, err := blog.PrepareCountComment()
	if !assert.NoError(err) {
		t.FailNow()
	}

	n, err := stmt.Query()
	if !assert.NoError(err) {
		t.FailNow()
	}

	err = stmt.Close()
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.EqualValues(2, n)
}

func TestBlogCountBlogByID(t *testing.T) {
	assert := assert.New(t)

	n, err := blog.CountBlogByID("second")
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.EqualValues(1, n)
}

func TestBlogPreparedCountBlogByID(t *testing.T) {
	assert := assert.New(t)

	stmt, err := blog.PrepareCountBlogByID()
	if !assert.NoError(err) {
		t.FailNow()
	}

	n, err := stmt.Query("third")
	if !assert.NoError(err) {
		t.FailNow()
	}

	err = stmt.Close()
	if !assert.NoError(err) {
		t.FailNow()
	}

	assert.EqualValues(1, n)
}
