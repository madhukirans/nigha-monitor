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

var DriverHealthMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "csi_driver",
		Name:      "health_metric",
		Help:      "nigha health metric @todo add doc here",
	},
	[]string{"driver", "dtype", "env_name", "type"},
)

func HealthMetricMetric(value float64, component, env_name, type1 string) {
	HealthMetric.With(prometheus.Labels{"component": component, "env_name": env_name, "type": type1}).Set(value)
}

func DriverHealthMetricMetric(value float64, driver, dtype, env_name, type1 string) {
	DriverHealthMetric.With(prometheus.Labels{"driver": driver, "dtype": dtype, "env_name": env_name, "type": type1}).Set(value)
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
	HealthMetricMetric(1.0, "prometheus", "nigha_ford.us.hopkinton.dellemc.com", "dev")
	HealthMetricMetric(1.0, "alertmanager", "nigha_ford.us.hopkinton.dellemc.com", "dev")
	HealthMetricMetric(1.0, "grafana", "nigha_ford.us.hopkinton.dellemc.com", "dev")
	HealthMetricMetric(1.0, "elasticsearch", "nigha_ford.us.hopkinton.dellemc.com", "dev")
	HealthMetricMetric(1.0, "kibana", "nigha_ford.us.hopkinton.dellemc.com", "dev")

	HealthMetricMetric(1.0, "prometheus", "nigha_benz.us.hopkinton.dellemc.com", "test")
	HealthMetricMetric(1.0, "alertmanager", "nigha_benz.us.hopkinton.dellemc.com", "test")
	HealthMetricMetric(1.0, "grafana", "nigha_benz.us.hopkinton.dellemc.com", "test")
	HealthMetricMetric(1.0, "elasticsearch", "nigha_benz.us.hopkinton.dellemc.com", "test")
	HealthMetricMetric(1.0, "kibana", "nigha_benz.us.hopkinton.dellemc.com", "test")

	HealthMetricMetric(1.0, "prometheus", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	HealthMetricMetric(1.0, "alertmanager", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	HealthMetricMetric(1.0, "grafana", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	HealthMetricMetric(1.0, "elasticsearch", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	HealthMetricMetric(1.0, "kibana", "nigha_benz.us.hopkinton.dellemc.com", "prod")

	DriverHealthMetricMetric(1.0, "csi-unity", "controller", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	DriverHealthMetricMetric(1.0, "csi-unity", "node", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	DriverHealthMetricMetric(1.0, "csi-powermax", "controller", "nigha_benz.us.hopkinton.dellemc.com", "prod")
	DriverHealthMetricMetric(1.0, "csi-powermax", "node", "nigha_benz.us.hopkinton.dellemc.com", "prod")
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
