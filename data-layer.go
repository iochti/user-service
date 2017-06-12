package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/namsral/flag"

	"github.com/iochti/user-service/models"
	_ "github.com/lib/pq"
)

// DataLayerInterface is here to abstract DB usage in our code
// so we can use any DB and not break other functions
// used for testing too
type DataLayerInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	DeleteUser(id string) error
}

// MgoDL implements DataLayerInterface
type MgoDL struct {
	DBName  string
	Session *mgo.Session
}

const USER_COLLECTION = "account"

var (
	mainSession *mgo.Session
	mainDB      *mgo.Database
	DBName      string
)

// Init inits the DB
func (m *MgoDL) Init() error {
	mHost := flag.String("mhost", "localhost", "MongoDB database host")
	mPort := flag.String("mport", "27017", "MongoDB's port")
	mName := flag.String("mname", "crm", "MongoDB's name")
	flag.Parse()
	mainSession, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s", *mHost, *mPort))
	if err != nil {
		panic(err)
	}
	mainDB = mainSession.DB(*mName)

	m.Session = mainSession.Copy()

	s := m.Session.Copy()
	defer s.Close()
	return nil
}

// CreateUser creates a user passed as parameter
func (m *MgoDL) CreateUser(user *models.User) error {
	timeCreated := time.Now()
	sess := m.Session.Copy()
	defer sess.Close()
	user.ID = bson.NewObjectId()
	user.Created = timeCreated
	user.Updated = timeCreated
	if err := sess.DB(m.DBName).C(USER_COLLECTION).Insert(&user); err != nil {
		return err
	}
	return nil
}

// GetUserByID fetch a user by its user id
func (m *MgoDL) GetUserByID(id string) (*models.User, error) {
	sess := m.Session.Copy()
	defer sess.Close()
	var user models.User
	if err := sess.DB(m.DBName).C(USER_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByLogin fetch a user by its github login
func (m *MgoDL) GetUserByLogin(login string) (*models.User, error) {
	sess := m.Session.Copy()
	defer sess.Close()
	// Check ID values
	if login == "" {
		return nil, fmt.Errorf("Error, invalid search: login must not be empty")
	}
	user := new(models.User)
	if err := sess.DB(m.DBName).C(USER_COLLECTION).Find(bson.M{"login": bson.M{"$eq": login}}).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail fetch a user by its github email
func (m *MgoDL) GetUserByEmail(email string) (*models.User, error) {
	// Check ID values
	if email == "" {
		return nil, fmt.Errorf("Error, invalid search: login must not be empty")
	}
	sess := m.Session.Copy()
	defer sess.Close()
	user := new(models.User)
	if err := sess.DB(m.DBName).C(USER_COLLECTION).Find(bson.M{"email": bson.M{"$eq": email}}).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser delete a user identified by its id
func (m *MgoDL) DeleteUser(id string) error {
	if id == "" {
		return fmt.Errorf("Error, invalid argument: id must be > 0")
	}

	sess := m.Session.Copy()
	defer sess.Close()
	if err := sess.DB(m.DBName).C(USER_COLLECTION).RemoveId(id); err != nil {
		return err
	}
	return nil
}
