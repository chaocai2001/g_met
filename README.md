# GMet
GMet is to solve the following pains:
* Various metrics solution in a group bring the hard problems of maintaining. Even, the whole group uses unified data aggregation and demonstration solution such as: ELK (ElasticSearch LogStach Kibana) or EFK (ElasticSearch FluentD Kibana). Because, the way to output metric log/data and the format of log/data are normally different. 

* The ineffcient metric data creating and extracting ways take huge CPU resouce.

* So many different configurations (fluentD configurations, elasticsearch mappings), which are hard to maintain.

* Creating metric monitor is a time-consuming task, which needs to write and debug the configurations for the related systems (fluentD, elastic search)


GMet is to help you create the program running metrics quickly and easily.
* Easy to use. To send the metrics, just one line code is needed.
* Easy to extend. GMet is based on the plugin structure, it is open to extend and by ioslated with the abstract layer, the changes are transparent to the clients.
* Easy to maintain. Use the unified API to avoid varied and messy implementations. Deal with different strategies on log and metric data by separating them.


```Golang
//The following example is to demonstrate how to use GMet.
func main() {
	//create GMet instance by given default writer and the formatter
	gmet := CreateGMetInstanceByDefault("../configs/g_met_seelog/g_met_log.xml")
	//Create a metric item of host IP, the name of the metric item is "HostAddr"
	addr, _ := IpAddress()
	for i := 0; i < 100; i++ {
		gmet.Send(addr, Metric("input_bytes", strconv.Itoa(rand.Intn(100))),
			Metric("output_bytes", strconv.Itoa(rand.Intn(100))))
		gmet.Flush() //in your real case, DON'T flush for each sending.
		//For seelog writer, the auto-flushing can be set in the log configuration
		time.Sleep(time.Second * 1)

	}
}
```
For the example, please, refer to https://github.com/chaocai2001/g_met_examples
