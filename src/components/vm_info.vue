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
        cputime: "",
        nettimer: "",
        disktimer: "",
        memtimer: "",

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
                    },
                    {
                        name: '可用内存',
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
                text: '硬盘IO（KB）'
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
                        name: 'READ IO',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    },
                    {
                        name: 'WRITE IO',
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
                text: '网卡速率（KB）'
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
                        name: 'RX IO',
                        data: ["", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
                    },
                    {
                        name: 'TX IO',
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
    this.cputimer = setInterval(() => {
        this.getcpuChart();
        }, 1000);

    this.getmemChart();
    this.memtimer = setInterval(() => {
        this.getmemChart();
        }, 1000);

    this.getChart();
    this.nettimer = setInterval(() => {
        this.getChart();
        }, 1000);

    this.getdiskChart();
    this.disktimer = setInterval(() => {
        this.getdiskChart();
        }, 1000);
  },

    beforeDestroy() {
        clearInterval(this.nettimer);        
        this.nettimer = null;
        clearInterval(this.disktimer);        
        this.nettimer = null;
        clearInterval(this.memtimer);        
        this.nettimer = null;
        clearInterval(this.cputimer);        
        this.nettimer = null;
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

                 this.memchart.series[0].data.push(response.data.res["Mem_total"])
                 if (this.memchart.series[0].data.length > 20 ) {
                    this.memchart.series[0].data.splice(0, 1)
                    }

                 this.memchart.series[1].data.push(response.data.res["Mem_used"])
                 if (this.memchart.series[1].data.length > 20 ) {
                    this.memchart.series[1].data.splice(0, 1)
                    }

                 this.memchart.series[2].data.push(response.data.res["Mem_availabled"])
                 if (this.memchart.series[2].data.length > 20 ) {
                    this.memchart.series[2].data.splice(0, 1)
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
