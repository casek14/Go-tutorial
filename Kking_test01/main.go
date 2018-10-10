package main


import (
	"google.golang.org/api/container/v1"
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	"log"
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"fmt"
	)

const PROJECT_ID  = "kking-213316"
const ZONE = "europe-north1-a"
const CLUSTER_NAME = "lokart-test"
const JSON_FILE = "/home/casek/Documents/kking-4f149755e825.json"



func main() {

	gce, err := newContainerClient()

	if err != nil{
		log.Fatal(err)
	}

	/*
	err = createCluster(gce,PROJECT_ID,ZONE,CLUSTER_NAME)
	if err != nil {
		log.Fatal(err)
	}
	*/

	cluster, err := getCluster(gce,PROJECT_ID,ZONE,CLUSTER_NAME)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("Here is the info about cluster: %s\n",cluster.Name)

	fmt.Printf("Node count is: %d\n",cluster.InitialNodeCount)

	fmt.Printf("Cluster ipv4 CIDR: %s\n",cluster.ClusterIpv4Cidr)

	fmt.Printf("Stack was created at %s\n",cluster.CreateTime)
	fmt.Printf("Version of the master endpoint:  %s\n",cluster.CurrentMasterVersion)
	for _,location := range cluster.Locations{
		fmt.Printf("Cluster is located in: %s\n",location)
	}
	fmt.Printf("Status of the cluster is: %s\n",cluster.Status)


}

func newContainerClient() (*container.Service, error){
	ctx := context.Background()

	data, err := ioutil.ReadFile(JSON_FILE)
	if err != nil {
		log.Fatal("Cannot load file ",JSON_FILE)
	}

	creds, err := google.CredentialsFromJSON(ctx, data,"https://www.googleapis.com/auth/cloud-platform")

	o := []option.ClientOption{
		option.WithEndpoint("https://container.googleapis.com/"),
		option.WithCredentials(creds),
	}

	httpClient, endpoint, err := transport.NewHTTPClient(ctx, o...)

	if err != nil {
		log.Fatal(err)
	}

	client, err := container.New(httpClient)

	if err != nil {
		return nil, err
	}

	client.BasePath = endpoint

	return client, nil
}

func createCluster(gce *container.Service, project string, zone string, cluster string) error{

	req := &container.CreateClusterRequest{}
	req.Cluster = &container.Cluster{Name: cluster, InitialNodeCount: 3}
	//_, err := gce.Projects.Zones.Clusters.Create(project, zone,req).Do()


	return nil
}

func getCluster(gce *container.Service, project string, zone string, cluster string) (*container.Cluster, error){

	return gce.Projects.Zones.Clusters.Get(project,zone,cluster).Do()
}