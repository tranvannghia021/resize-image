package main

import (
	"bufio"
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"os"
)

type Blog struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	CreatedAt   string             `bson:"created_at,omitempty"`
	UpdatedAt   string             `bson:"updated_at,omitempty"`
}

var (
	projectId = "golang-dev-1"
	topicId   = "my-topic"
)

func main() {
	/**
	connect database cluster scylla
	*/
	//var cluster = gocql.NewCluster("node-0.aws-us-east-1.e634da48f786d4ba37ae.clusters.scylla.cloud", "node-1.aws-us-east-1.e634da48f786d4ba37ae.clusters.scylla.cloud", "node-2.aws-us-east-1.e634da48f786d4ba37ae.clusters.scylla.cloud")
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "KXAwMNJ1VIf4U5g"}
	//cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")
	//
	//var session, err = cluster.CreateSession()
	//if err != nil {
	//	panic("Failed to connect to cluster")
	//}
	//
	//defer session.Close()
	//
	//var query = session.Query("SELECT * FROM system.clients")
	//
	//if rows, err := query.Iter().SliceMap(); err == nil {
	//	log.Println(rows)
	//}

	/**
	connect mongo db
	//*/
	//serverApi := options.ServerAPI(options.ServerAPIVersion1)
	//// init option prepare
	//opts := options.Client().ApplyURI(configs.GetMongoDns()).SetServerAPIOptions(serverApi).SetConnectTimeout(2 * time.Second)
	//client, err := mongo.Connect(context.TODO(), opts)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var collection string = "blogs"
	//// init context connection
	//connect := client.Database(configs.GetMongoDB()).Collection(collection)
	//
	//// insert table blogs (id ,title , description ,created_at,updated_at)
	////
	////var document = map[string]string{
	////	"title":       "update new title",
	//	"description": "SocketTimeout, wTimeout, MaxTimevà MaxCommitTime sẽ không được dùng nữa trong bản phát hành sắp tới. Trình điều khiển sẽ bỏ qua MaxTimevà MaxCommitTimenếu bạn đặt Timeout. Trình điều khiển vẫn tôn trọng SocketTimeoutvà wTimeout, nhưng những cài đặt này có thể dẫn đến hành vi không xác định. Thay vào đó, hãy cân nhắc chỉ sử dụng tùy chọn hết thời gian chờ duy nhất.",
	//	"created_at":  time.Now().String(),
	//	"updated_at":  time.Now().String(),
	//}
	//_id, _ := primitive.ObjectIDFromHex("65a89b50ffef506811900b6e")
	//var blogs = []interface{}{
	//	Blog{
	//		Title:       "title 1",
	//		Description: "description 1",
	//		CreatedAt:   time.Now().String(),
	//		UpdatedAt:   time.Now().String(),
	//	},
	//	Blog{
	//		Title:       "title 2",
	//		Description: "description 2",
	//		CreatedAt:   time.Now().String(),
	//		UpdatedAt:   time.Now().String(),
	//	},
	//}

	// search like in sql!
	//var result []Blog
	//var filter = bson.M{"title": bson.M{"$regex": "title"}}
	//var option = options.Find().SetProjection(bson.M{"title": 1, "description": 1})
	//var cursor, err = connect.Find(context.Background(), filter, option)
	//filter := bson.D{
	//	{"title", bson.D{{"$regex", "1"}}},
	//}
	//
	//result, err := connect.DeleteMany(context.TODO(), filter)
	//
	//if err != nil {
	//	log.Fatal("[insert]", err)
	//}
	//
	//log.Println(result.DeletedCount)

	/**
	connect redis
	*/
	//
	//client := redis.NewClient(&redis.Options{
	//	Addr:     ":6379",
	//	Password: "", // no password set
	//	DB:       3,  // use default DB
	//})
	//ctx := context.Background()
	////session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	////for k, v := range session {
	////	err := client.HSet(ctx, "user-session:123", k, v).Err()
	////	if err != nil {
	////		panic(err)
	////	}
	////}
	//userSession := client.HGetAll(ctx, "user-session:123").Val()
	//fmt.Println(userSession)

	/**
	pubsub
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	if err := publish(os.Stdout, projectId, topicId, text); err != nil {
		panic(err)
	}
}

func publish(w io.Writer, projectID, topicID, msg string) error {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %w", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
