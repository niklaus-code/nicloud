<template>
  <div style="margin-top: 10px;padding-right:0; padding-left:0">
    <div style="float: right;  margin-bottom: 12px" class="col-md-5">
      <button style="float: right; margin-right: 5px; display:inline;" class="btn btn-success btn-sm" @click="createhost" type="button">
        创建宿主机<span class="glyphicon glyphicon-plus" style="margin-left: 5px"></span>
      </button>
    </div>
    <div class="col-md-12" style="border-bottom: 1px solid #9f0303;  border-top: 1px solid #9f0303; padding-top: 5px">
      <h4 style="margin-left: 10px">数据中心</h4>
      <table class="table1 table table-hover" style="text-align: center;">
        <thead>
          <tr>
            <th>数据中心</th>
            <th>总物理机数</th>
            <th>总CPU数（核）</th>
            <th>总内存数（G）</th>
            <th>已分配CPU（核）</th>
            <th>已分配内存（G）</th>
            <th>cpu使用率</th>
            <th>内存使用率</th>
          </tr>
        </thead>

        <tbody>
          <tr class="table-dark text-dark">
            <td>{{ counthosts.Datacenter }}</td>
            <td>{{ counthosts.Counthosts }}</td>
            <td>{{ counthosts.Cpu }}核</td>
            <td>{{ counthosts.Mem }}G</td>
            <td>{{ counthosts.Usedcpu }}核</td>
            <td>{{ counthosts.Usedmem }}G</td>
            <td>{{ counthosts.Cpu_percent }}%</td>
            <td>{{ counthosts.Mem_percent }}%</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="col-md-12" style="border-bottom: 1px solid green; border-top: 1px solid green; padding-top: 5px; margin-bottom: 30px">
      <h4 style="margin-left: 10px">宿主机</h4>
      <table class="table2 table table-hover" style="text-align: center;">
        <thead>
          <tr>
            <th>
              <label class="checkbox-inline" style="border:red 1px">
                <input type="checkbox" v-model="checkvalue" @click="checkbox()" />
              </label>
            </th>
            <th>IP地址</th>
            <th>VLAN子网</th>
            <th>已分配 / cpu（核）</th>
            <th>已分配 / 内存（G）</th>
            <th>可创建数</th>
            <th>实际虚拟机数</th>
            <th>数据中心</th>
            <th>备注</th>
            <th>操作</th>
          </tr>
        </thead>

        <tbody v-for="(item, index) in data" :key="item.id">
          <tr class="table-dark text-dark" :id="item.Uuid" v-if="item.Status">
            <label class="checkbox-inline" style="width:10px">
              <input type="checkbox" v-model="item.Checkout" />
            </label>
            <td>{{ item.Ipv4 }}</td>
            <td>
              <ul>
                <li v-for="k in item.vlan" :key="k.id">
                  {{ k.Vlan }}
                </li>
              </ul>
            </td>
            <td>{{ item.Usedcpu }}核/{{ item.Cpu }}核</td>
            <td>{{ item.Usedmem }}G/{{ item.Mem }}G</td>
            <td>{{ item.count }}/{{ item.Max_vms }}</td>
            <td>{{ item.vmnum }}</td>
            <td>{{ item.Datacenter }}</td>
            <td>
              <span v-if="item.flag2" @click="c(index)">
                {{ item.Comment }}
              </span>
              <li v-if="item.flag"><span class="glyphicon glyphicon-calendar" @click="edit(index)"></span></li>
              <div v-if="item.flag1">
                <div><input type="text" v-model="comments" /></div>
                <div><span @click="input(index, item.Ipv4)" class="glyphicon glyphicon-calendar"></span></div>
              </div>
            </td>

            <td>
              <button class="btn btn-primary btn-xs" type="button" @click="updatehost(item.Ipv4, index)">
                <span class="glyphicon glyphicon-edit"></span>
                修改
              </button>
              <button class="btn btn-danger btn-xs" type="button" @click="deletehost(item.Ipv4, index)">
                <span class="glyphicon glyphicon-trash"></span>
                删除
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      counthosts: {},
      data: [],
      cpu: '',
      mem: '',
      ip: '',
      num: '',
      comments: ''
    }
  },

  mounted: function() {
    this.gethost()
    this.counthost()
  },

  methods: {
    createhost: function() {
      this.$emit('toParent', 'createhost')
    },

    updatehost: function(ip, index) {
      this.$emit('toParent', 'updatehost')
      this.$store.state.host.ip = ip
    },

    deletehost: function(ip, index) {
      var apiurl = `/api/hosts/delete`
      this.$http.get(apiurl, { params: { ip: ip } }).then(response => {
        if (response.data.err === null) {
          alert('删除成功')
          this.data[index].Status = 0
        } else {
          alert(response.data.err.Message)
        }
      })
    },

    c: function(index) {
      this.data[index].flag2 = false
      this.data[index].flag1 = true
      this.comments = this.data[index].Comment
    },

    edit: function(index) {
      this.data[index].flag = false
      this.data[index].flag1 = true
    },

    input: function(index, ip) {
      var apiurl = `/api/hosts/addcomment`
      this.$http.get(apiurl, { params: { ip: ip, comment: this.comments } }).then(response => {
        if (response.data) {
          this.data[index].Comment = this.comments
        }
      })
      this.data[index].flag = false
      this.data[index].flag1 = false
      this.data[index].flag2 = true
    },

    getvmstatus: function(uuid, host) {
      var apiurl = `/api/vm/getstatus`
      return this.$http.get(apiurl, { params: { uuid: uuid, host: host } }).then(response => {
        return response.data.res
      })
    },

    comment: function(res) {
      var d = []
      for (var v in res) {
        if (res[v]['Comment'].length > 0) {
          res[v]['flag'] = false
          res[v]['flag2'] = true
        } else {
          res[v]['flag2'] = false
          res[v]['flag'] = true
        }
        res[v]['flag1'] = false
        d.push(res[v])
        this.data = d
      }
      for (let v in this.data) {
        // var r = this.getvmstatus(this.data[v].Uuid, this.data[v].Host)
        // r.then(value => {
        //  this.data[v].Status = value
        //  })

        var c = this.gethostvm(this.data[v].Ipv4)
        c.then(value => {
          this.data[v].vmnum = value
        })
      }
    },

    gethostvm: function(hostip) {
      var apiurl = `/api/hosts/countdomains`
      return this.$http.get(apiurl, { params: { host: hostip } }).then(response => {
        if (response.data.err === null) {
          return response.data.res
        } else {
          return 0
        }
      })
    },

    counthost: function() {
      var apiurl = `/api/hosts/counthosts`
      this.$http.get(apiurl).then(response => {
        if (response.data.err === null) {
          this.counthosts = response.data.res
        } else {
          alert(response.data.err.Message)
        }
      })
    },

    gethost: function(ip) {
      var apiurl = `/api/hosts/gethosts`
      this.$http.get(apiurl).then(response => {
        if (response.data.err === null) {
          this.comment(response.data.res)
        } else {
          alert(response.data.err.Message)
        }
      })
    }
  }
}
</script>
<style scoped>
.table thead tr th {
  border-bottom: 2px solid #846d6d;
}

.table {
  margin-bottom: 10px;
}

.col-md-12 {
  padding-right: 0;
  padding-left: 0;
}

select {
  font-family: '微软雅黑';
  border: 1px #1a1a1a solid;
  border-radius: 5px;
}

.content {
  box-shadow: 0 0 10px rgba(0, 0, 0, 8);
  border-radius: 10px/10px;
  z-index: -1;
  padding: 50px 0px 50px 0px;
  margin-left: 0px;
  margin-top: 50px;
}
.checkbox-inline {
  margin-bottom: 30px;
}

.details-content .article-cont p {
  padding: 30px 0 0 5px;
}

label {
  font-weight: 400;
}

.table tbody tr td {
  vertical-align: 'middle';
  border: none;
  padding-bottom: 0;
}

.table1 th {
  background-color: #e8d18d;
}

.table2 th {
  background-color: #e8d18d;
  font-weight: bold;
  color: black;
  text-align: center;
  border-bottom: none;
}

.table th {
  font-weight: bold;
  color: black;
  text-align: center;
  border-bottom: none;
}
</style>
