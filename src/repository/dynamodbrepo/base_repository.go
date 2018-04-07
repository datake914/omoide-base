package dynamodbrepo

// AWS Session.
import (
	"errors"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var sess = session.Must(session.NewSession())

type Config struct {
	DynamoDBEndpoint    string
	DynamoDBTablePrefix string
}

type BaseRepository struct {
	Cfg Config
	DB  *dynamodb.DynamoDB
}

func NewBaseRepository(cfg Config) BaseRepository {
	return BaseRepository{
		Cfg: cfg,
		DB:  dynamodb.New(sess, new(aws.Config).WithEndpoint(cfg.DynamoDBEndpoint)),
	}
}

func (repo BaseRepository) Table(s string) *string {
	return aws.String(repo.Cfg.DynamoDBTablePrefix + "." + s)
}

func (repo BaseRepository) Index(s string) *string {
	return aws.String(repo.Cfg.DynamoDBTablePrefix + "." + s + ".GSI")
}

func (repo BaseRepository) UpdateBuilder(domain interface{}) (expression.UpdateBuilder, error) {
	update := expression.UpdateBuilder{}
	rv := reflect.Indirect(reflect.ValueOf(domain))
	if rv.Kind() != reflect.Struct {
		return expression.UpdateBuilder{}, errors.New("domain must be a struct.")
	}
	for i, rt := 0, rv.Type(); i < rv.NumField(); i++ {
		name, val, tag := rt.Field(i).Name, rv.Field(i).Interface(), rt.Field(i).Tag
		if _, ok := tag.Lookup("dynamodbkey"); ok {
			continue
		}
		if v, ok := tag.Lookup("dynamodbav"); ok {
			name = v
		}
		update = update.Set(expression.Name(name), expression.Value(val))
	}
	return update, nil
}

func (repo BaseRepository) NotKeyAttributes(domain interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return repo.Attributes(domain, func(tag reflect.StructTag) bool {
		_, ok := tag.Lookup("dynamodbkey")
		return !ok
	})
}

func (repo BaseRepository) KeyAttributes(domain interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return repo.Attributes(domain, func(tag reflect.StructTag) bool {
		_, ok := tag.Lookup("dynamodbkey")
		return ok
	})
}

func (repo BaseRepository) Attributes(domain interface{}, condition func(tag reflect.StructTag) bool) (map[string]*dynamodb.AttributeValue, error) {
	rv := reflect.Indirect(reflect.ValueOf(domain))
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("domain must be a struct.")
	}
	// Create dynamodb attributes.
	attr, err := dynamodbattribute.MarshalMap(rv.Interface())
	if err != nil {
		return nil, err
	}
	// Delete except key attributes.
	//TODO: フィールド名とタグ名が一致しない場合にエラー
	for i, rt := 0, rv.Type(); i < rv.NumField(); i++ {
		name, tag := rt.Field(i).Name, rt.Field(i).Tag
		if !condition(tag) {
			delete(attr, name)
		}
	}
	return attr, nil
}
