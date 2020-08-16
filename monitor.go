package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"time"
)

var HealthMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "nigha",
		Name:      "health_metric",
		Help:      "nigha health metric @todo add doc here",
	},
	[]string{"component", "env_name", "type"},
)

func HealthMetricMetric(component, env_name, type1 string) {
	HealthMetric.With(prometheus.Labels{"component": component, "env_name": env_name, "type": type1}).Set(1.0)
}

func init() {
	fmt.Println("Initializing the prometheus")
	prometheus.MustRegister(HealthMetric)
}

func PromHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	Demo(c)
}

func Demo(c *gin.Context) {
	//for {
		HealthMetricMetric("prometheus", "nigha_ford.us.hopkinton.dellemc.com", "dev")
		HealthMetricMetric("alertmanager", "nigha_ford.us.hopkinton.dellemc.com", "dev")
		HealthMetricMetric("grafana", "nigha_ford.us.hopkinton.dellemc.com", "dev")
		HealthMetricMetric("elasticsearch", "nigha_ford.us.hopkinton.dellemc.com", "dev")
		HealthMetricMetric("kibana", "nigha_ford.us.hopkinton.dellemc.com", "dev")

		HealthMetricMetric("prometheus", "nigha_benz.us.hopkinton.dellemc.com", "test")
		HealthMetricMetric("alertmanager", "nigha_benz.us.hopkinton.dellemc.com", "test")
		HealthMetricMetric("grafana", "nigha_benz.us.hopkinton.dellemc.com", "test")
		HealthMetricMetric("elasticsearch", "nigha_benz.us.hopkinton.dellemc.com", "test")
		HealthMetricMetric("kibana", "nigha_benz.us.hopkinton.dellemc.com", "test")

		HealthMetricMetric("prometheus", "nigha_benz.us.hopkinton.dellemc.com", "prod")
		HealthMetricMetric("alertmanager", "nigha_benz.us.hopkinton.dellemc.com", "prod")
		HealthMetricMetric("grafana", "nigha_benz.us.hopkinton.dellemc.com", "prod")
		HealthMetricMetric("elasticsearch", "nigha_benz.us.hopkinton.dellemc.com", "prod")
		HealthMetricMetric("kibana", "nigha_benz.us.hopkinton.dellemc.com", "prod")
		time.Sleep(5 * time.Second)
	//}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func StartServer() {
	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("/")
	{
		v1.GET("/metrics", PromHandler)
		v1.GET("/demo", Demo)
	}
	log.Println("Starting REST server")

	r.Run(":9360")
}
