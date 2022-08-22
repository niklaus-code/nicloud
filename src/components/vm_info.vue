<template>
<div>
    <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top: 20px">
        <div style="width: 200px; float: left;">
            <h3>云主机 & 详情</h3>
        </div>
    </div>
    
    <div class="col-sm-12 choose" >
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>UUID :</label>
            </div>
            <div class="col-sm-10" >
                {{uuid}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>IP :</label>
            </div>
            <div class="col-sm-10" >
               {{ip}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>所属者 :</label>
            </div>
            <div class="col-sm-10" >
               {{owner}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>宿主机 :</label>
            </div>
            <div class="col-sm-10" >
               {{host}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>存储集群 :</label>
            </div>
            <div class="col-sm-10" >
                {{storage}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>数据中心 :</label>
            </div>
            <div class="col-sm-10" >
                {{datacenter}}
            </div>
        </div>
        <div class="col-sm-12" >
            <div class="col-sm-2" >
                <label>备注 :</label>
            </div>
            <div class="col-sm-10" >
                {{comment}}
            </div>
        </div>
    </div>

    <div class="col-sm-12" style="margin-top: 50px">
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
        ip: "",
        datacenter: "",
        owner: "",
        storage: "",
        comment: "",
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
        }, 2000);

    this.getmemChart();
    this.memtimer = setInterval(() => {
        this.getmemChart();
        }, 2000);

    this.getChart();
    this.nettimer = setInterval(() => {
        this.getChart();
        }, 2000);

    this.getdiskChart();
    this.disktimer = setInterval(() => {
        this.getdiskChart();
        }, 2000);
  },

    //离开当前路由, 停止请求
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
        var v = this.$store.state.vm.uuid
        if (v === null || typeof v === 'undefined' || v === '' || v === "undefined") {
            this.uuid = sessionStorage.getItem('uuid')
            this.host = sessionStorage.getItem('host')
            this.ip = sessionStorage.getItem('ip')
            this.os = sessionStorage.getItem('os')
            this.datacenter = sessionStorage.getItem('datacenter')
            this.storage = sessionStorage.getItem('storage')
            this.owner = sessionStorage.getItem('owner')
            this.comment = sessionStorage.getItem('comment')

            } else {
                this.uuid = this.$store.state.vm.uuid
                this.host = this.$store.state.vm.host
                this.ip = this.$store.state.vm.ip
                this.os = this.$store.state.vm.os
                this.datacenter = this.$store.state.vm.datacenter
                this.storage = this.$store.state.vm.storage
                this.owner = this.$store.state.vm.owner
                sessionStorage.setItem('uuid', this.$store.state.vm.uuid)
                sessionStorage.setItem('ip', this.$store.state.vm.ip)
                sessionStorage.setItem('os', this.$store.state.vm.os)
                sessionStorage.setItem('datacenter', this.$store.state.vm.datacenter)
                sessionStorage.setItem('storage', this.$store.state.vm.storage)
                sessionStorage.setItem('owner', this.$store.state.vm.owner)
                sessionStorage.setItem('comment', this.$store.state.vm.comment)
                sessionStorage.setItem('host', this.$store.state.vm.host)
                }
        },

    getcpuChart() {
        var apiurl = `/api/vm/details/cpuinfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 console.log(response.data.err.Message)
             } else {
                 this.cpuchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.cpuchart.xAxis.categories.length > 60 ) {
                    this.cpuchart.xAxis.categories.splice(0, 1)
                    }

                 this.cpuchart.series[0].data.push(response.data.res["Load"])
                 if (this.cpuchart.series[0].data.length > 60 ) {
                    this.cpuchart.series[0].data.splice(0, 1)
                    }
                }
            })
        },

    getmemChart() {
        var apiurl = `/api/vm/details/meminfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 console.log(response.data.err.Message)
             } else {
                 this.memchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.memchart.xAxis.categories.length > 60 ) {
                    this.memchart.xAxis.categories.splice(0, 1)
                    }

                 this.memchart.series[0].data.push(response.data.res["Mem_total"])
                 if (this.memchart.series[0].data.length > 60 ) {
                    this.memchart.series[0].data.splice(0, 1)
                    }

                 this.memchart.series[1].data.push(response.data.res["Mem_used"])
                 if (this.memchart.series[1].data.length > 60 ) {
                    this.memchart.series[1].data.splice(0, 1)
                    }

                 this.memchart.series[2].data.push(response.data.res["Mem_availabled"])
                 if (this.memchart.series[2].data.length > 60 ) {
                    this.memchart.series[2].data.splice(0, 1)
                    }
                }
            })
        },

    getdiskChart() {
        var apiurl = `/api/vm/details/diskinfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 console.log(response.data.err.Message)
             } else {
                 this.diskchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.diskchart.xAxis.categories.length > 60 ) {
                    this.diskchart.xAxis.categories.splice(0, 1)
                    }

                 this.diskchart.series[0].data.push(response.data.res["Read"])
                 if (this.diskchart.series[0].data.length > 60 ) {
                    this.diskchart.series[0].data.splice(0, 1)
                    }

                 this.diskchart.series[1].data.push(response.data.res["Write"])
                 if (this.diskchart.series[1].data.length > 60 ) {
                    this.diskchart.series[1].data.splice(0, 1)
                    }
                }
            })
        },

    getChart() {
        var apiurl = `/api/vm/details/netinfo`
        this.$http.get(apiurl, { params: {uuid: this.uuid, host: this.host} }).then(response => {
             if (response.data.err != null) {
                 console.log(response.data.err.Message)
             } else {
                 this.netchart.xAxis.categories.push(response.data.res["Ctime"])
                 if (this.netchart.xAxis.categories.length > 60 ) {
                    this.netchart.xAxis.categories.splice(0, 1)
                    }

                 this.netchart.series[0].data.push(response.data.res["Rx"])
                 if (this.netchart.series[0].data.length > 60 ) {
                    this.netchart.series[0].data.splice(0, 1)
                    }

                 this.netchart.series[1].data.push(response.data.res["Tx"])
                 if (this.netchart.series[1].data.length > 60 ) {
                    this.netchart.series[1].data.splice(0, 1)
                    }
                }
            })
        }
    }
 }
</script>

<style>
.col-sm-2 label {
    float: right
}

.choose {
    margin-top: 30px;
    padding-bottom: 30px;
    border-bottom: solid #ddd 1px;
    border-radius: 4px 4px 0 0;
}
</style>
