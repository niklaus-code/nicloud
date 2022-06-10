<template>
<div class="col-sm-12">
    <div class="col-sm-6">
        <highcharts :options="cpuchart"></highcharts>
    </div>
    <div class="col-sm-6">
        <highcharts :options="memchart"></highcharts>
    </div>
    <div class="col-sm-6">
        <highcharts :options="netchart"></highcharts>
    </div>
    <div class="col-sm-6">
        <highcharts :options="diskchart"></highcharts>
    </div>
</div>
</template>

<script>
import {Chart} from 'highcharts-vue'

export default {
  data() {
    return {
        uuid: "",
        host: "",
        cpuchart: {
            chart: {
                type: 'areaspline'
                },
        
            plotOptions: {
		        areaspline: {
                    fillOpacity: 0.5
		            }
	            },

            title: {
                text: 'CPU使用率'
                },

            yAxis: {
		        title: {
			        text: ''
		        }
	        },

            xAxis: {
                type: 'datetime',
                categories: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                },

                series: [
                    {
                        name: '使用率',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    }
                ]
            },

        memchart: {
            chart: {
                type: 'areaspline'
                },

            plotOptions: {
		        areaspline: {
                    fillOpacity: 0.5
		            }
	            },

            title: {
                text: '内存使用统计'
                },

            yAxis: {
		        title: {
			        text: ''
		        }
	        },

            xAxis: {
                type: 'datetime',
                categories: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                },

                series: [
                    {
                        name: '总内存',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    },
                    {
                        name: '已用内存',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    }
                ]
            },

        diskchart: {
            chart: {
                type: 'areaspline'
                },

            plotOptions: {
		        areaspline: {
                    fillOpacity: 0.5
		            }
	            },

            title: {
                text: '硬盘IO'
                },

            yAxis: {
		        title: {
			        text: ''
		        }
	        },

            xAxis: {
                type: 'datetime',
                categories: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                },

                series: [
                    {
                        name: '读取速度',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    },
                    {
                        name: '写速度',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    }
                ]
            },

        netchart: {
            chart: {
                type: 'areaspline'
                },

            plotOptions: {
		        areaspline: {
                    fillOpacity: 0.5
		            }
	            },

            title: {
                text: '网卡速率'
                },

            yAxis: {
		        title: {
			        text: ''
		        }
	        },

            xAxis: {
                type: 'datetime',
                categories: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                },

                series: [
                    {
                        name: '接收速度',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    },
                    {
                        name: '发送速度',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    }
                ]
            }
        };
  },

  components: {
      highcharts: Chart
  },

  created(){
    this.monitor()
    },

  mounted(){
    this.getcpuChart();
    this.times = setInterval(() => {
        this.getcpuChart();
        }, 1000);

    this.getmemChart();
    this.times = setInterval(() => {
        this.getmemChart();
        }, 1000);

    this.getChart();
    this.times = setInterval(() => {
        this.getChart();
        }, 1000);

    this.getdiskChart();
    this.times = setInterval(() => {
        this.getdiskChart();
        }, 1000);
  },

  methods:{
    monitor() {
        var v = this.$store.state.monitor.uuid
        if (v === null || typeof v === 'undefined' || v === '' || v === "undefined") {
            this.uuid = sessionStorage.getItem('uuid')
            this.host = sessionStorage.getItem('host')
            } else {
                this.uuid = this.$store.state.monitor.uuid
                this.host = this.$store.state.monitor.host
                sessionStorage.setItem('uuid', this.$store.state.monitor.uuid)
                sessionStorage.setItem('host', this.$store.state.monitor.host)
                }
        },

    getcpuChart() {
        var apiurl = `/api/vm/details/cpuinfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 alert(response.data.err.Message)
             } else {
                 this.cpuchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.cpuchart.xAxis.categories.length > 20 ) {
                    this.cpuchart.xAxis.categories.splice(0, 1)
                    }

                 this.cpuchart.series[0].data.push(response.data.res["Load"])
                 if (this.cpuchart.series[0].data.length > 20 ) {
                    this.cpuchart.series[0].data.splice(0, 1)
                    }
                }
            })
        },

    getmemChart() {
        var apiurl = `/api/vm/details/meminfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 alert(response.data.err.Message)
             } else {
                 this.memchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.memchart.xAxis.categories.length > 20 ) {
                    this.memchart.xAxis.categories.splice(0, 1)
                    }

                 this.memchart.series[0].data.push(response.data.res["Mem_available"])
                 if (this.memchart.series[0].data.length > 20 ) {
                    this.memchart.series[0].data.splice(0, 1)
                    }

                 this.memchart.series[1].data.push(response.data.res["Mem_used"])
                 if (this.memchart.series[1].data.length > 20 ) {
                    this.memchart.series[1].data.splice(0, 1)
                    }
                }
            })
        },

    getdiskChart() {
        var apiurl = `/api/vm/details/diskinfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 alert(response.data.err.Message)
             } else {
                 this.diskchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.diskchart.xAxis.categories.length > 20 ) {
                    this.diskchart.xAxis.categories.splice(0, 1)
                    }

                 this.diskchart.series[0].data.push(response.data.res["Read"])
                 if (this.diskchart.series[0].data.length > 20 ) {
                    this.diskchart.series[0].data.splice(0, 1)
                    }

                 this.diskchart.series[1].data.push(response.data.res["Write"])
                 if (this.diskchart.series[1].data.length > 20 ) {
                    this.diskchart.series[1].data.splice(0, 1)
                    }
                }
            })
        },

    getChart() {
        var apiurl = `/api/vm/details/netinfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 alert(response.data.err.Message)
             } else {
                 this.netchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.netchart.xAxis.categories.length > 20 ) {
                    this.netchart.xAxis.categories.splice(0, 1)
                    }

                 this.netchart.series[0].data.push(response.data.res["Rx"])
                 if (this.netchart.series[0].data.length > 20 ) {
                    this.netchart.series[0].data.splice(0, 1)
                    }

                 this.netchart.series[1].data.push(response.data.res["Tx"])
                 if (this.netchart.series[1].data.length > 20 ) {
                    this.netchart.series[1].data.splice(0, 1)
                    }
                }
            })
        }
    }
 }
</script>
