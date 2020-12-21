package db

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"path"
	"github.com/boltdb/bolt"
)

//get the database absolute path
func DBPATH() string {
	return  path.Join(os.Getenv("GOPATH") ,"src/github.com/chuhongwei/BlogServer/source/Blog.db")
}

//create the bucket for article and user
func Init() {
	db, err := bolt.Open(DBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("article"))
			if err != nil {
				log.Fatal(err)
			}
		}

		b = tx.Bucket([]byte("user"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("user"))
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// PutArticle : put one article to blog.db
func PutArticle(article Article) error {
	db, err := bolt.Open(DBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b != nil {
				key := make([]byte, 8)
				binary.LittleEndian.PutUint64(key, uint64(article.Id))
				data, _ := json.Marshal(article)
				b.Put(key, data)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}


func PutUser(user User) error {
	db, err := bolt.Open(DBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b != nil {
				data, _ := json.Marshal(user)
				b.Put([]byte(user.Username), data)
			}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Get Atricle with ID,ID=-1 means get all articles
func GetArticles(id int64, page int64) []Article {
	db, err := bolt.Open(DBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	articles := make([]Article, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b != nil && id >= 0 {
			key := make([]byte, 8)
			binary.LittleEndian.PutUint64(key, uint64(id))
			data := b.Get(key)
			if data != nil {
				atc := Article{}
				err := json.Unmarshal(data, &atc)
				if err != nil {
					log.Fatal(err)
				}
				articles = append(articles, atc)
			}

		} else if b != nil && id == -1 {
			cursor := b.Cursor()
			nPerPage := 4
			fromKey := make([]byte, 8)
			binary.LittleEndian.PutUint64(fromKey, uint64(page-1)*(uint64)(nPerPage))

			for k, v := cursor.Seek(fromKey); k != nil && nPerPage > 0; k, v = cursor.Next() {
				atc := Article{}
				err := json.Unmarshal(v, &atc)
				if err != nil {
					log.Fatal(err)
				}
				articles = append(articles, atc)
				nPerPage--
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return articles
}

//  put articles to blog.db
func PutArticles(articles []Article) error {
	for i := 0; i < len(articles); i++ {
		PutArticle(articles[i])
	}
	return nil;
}

// put users to blog.db
func PutUsers(users []User) error {
	for i := 0; i < len(users); i++ {
		PutUser(users[i])
	}
	return nil
}

//Get user information with the username
func GetUser(username string) User {

	db, err := bolt.Open(DBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := User{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		if b != nil {
			data := b.Get([]byte(username))
			if data != nil {
				err := json.Unmarshal(data, &user)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return user
}
