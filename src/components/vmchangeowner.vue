<template>
  <div>
    <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top: 20px">
      <div style="width: 150px; float: left;">
        <h3>修改所属用户</h3>
      </div>
      <div style="margin-left: 30px; width: 160px; float: left; color: red">
        <h5>*挂载云盘会一并改变</h5>
      </div>
    </div>

    <div class="col-sm-8 col-sm-offset-2 choose" style="margin-top:30px;">
      <ul class="nav nav-pills nav-stacked">
        <li><strong>uuid</strong>:{{ uuid }}</li>
        <li><strong>ip</strong>:{{ ip }}</li>
        <li><strong>os</strong>:{{ os }}</li>
        <li><strong>host</strong>:{{ host }}</li>
        <li><strong>cpu</strong>:{{ cpu }}核</li>
        <li><strong>mem</strong>:{{ mem }}G</li>
        <li><strong>owner</strong>:{{ owner }}</li>
        <li><strong>comment</strong>:{{ comment }}</li>
      </ul>
    </div>
    <div class="col-sm-8 col-sm-offset-2 choose" style="margin-top:30px; margin-bottom:30px">
      <div class="col-sm-8">
        <div class="col-sm-2" style="margin-top:1px">选择用户</div>
        <div class="col-sm-9" style="padding-left:0">
          <select class="col-sm-3" v-model="uservalue">
            <option v-for="c in user" :value="c" :key="c.id">
              {{ c.Username }}
            </option>
          </select>
          <button @click="commit" style="margin-left:40px;margin-top:1px" type="button" class="btn btn-success btn-xs">提交</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      uuid: '',
      ip: '',
      os: '',
      host: '',
      cpu: '',
      mem: '',
      owner: '',
      comment: '',
      uservalue: {},
      user: []
    }
  },

  mounted: function() {
    this.vminfo()
    this.getuser()
  },

  methods: {
    getuser: function() {
      var apiurl = `/api/user/getuser`
      this.$http.get(apiurl).then(response => {
        let rr = response.data.res
        if (response.data.err === null) {
          var d = []
          for (var v in rr) {
            rr[v]['status'] = true
            d.push(rr[v])
          }
          this.user = rr
        } else {
          alert(response.data.err.Message)
        }
      })
    },

    vminfo: function() {
      var v = this.$store.state.editsetting.uuid
      if (v === null || typeof v === 'undefined' || v === '' || v === 'undefined') {
        this.uuid = sessionStorage.getItem('uuid')
        this.ip = sessionStorage.getItem('ip')
        this.os = sessionStorage.getItem('os')
        this.host = sessionStorage.getItem('host')
        this.cpu = sessionStorage.getItem('cpu')
        this.mem = sessionStorage.getItem('mem')
        this.owner = sessionStorage.getItem('owner')
        this.comment = sessionStorage.getItem('comment')
      } else {
        this.uuid = this.$store.state.editsetting.uuid
        this.ip = this.$store.state.editsetting.ip
        this.os = this.$store.state.editsetting.os
        this.host = this.$store.state.editsetting.host
        this.cpu = this.$store.state.editsetting.cpu
        this.mem = this.$store.state.editsetting.mem
        this.owner = this.$store.state.editsetting.owner
        this.comment = this.$store.state.editsetting.comment
        sessionStorage.setItem('uuid', this.$store.state.editsetting.uuid)
        sessionStorage.setItem('ip', this.$store.state.editsetting.ip)
        sessionStorage.setItem('os', this.$store.state.editsetting.os)
        sessionStorage.setItem('host', this.$store.state.editsetting.host)
        sessionStorage.setItem('cpu', this.$store.state.editsetting.cpu)
        sessionStorage.setItem('mem', this.$store.state.editsetting.mem)
        sessionStorage.setItem('owner', this.$store.state.editsetting.owner)
        sessionStorage.setItem('comment', this.$store.state.editsetting.comment)
      }
    },

    commit: function() {
      var apiurl = `/api/vm/vmchangeowner`
      this.$http.post(apiurl, this.$qs.stringify({ userid: this.uservalue.Id, uuid: this.uuid, ip: this.ip })).then(response => {
        if (response.data.res === null) {
          alert('修改成功')
        } else {
          alert(response.data.res.Message)
        }
      })
    }
  }
}
</script>
<style scoped>
.createip {
  font-weight: 500;
}

.vlaninfo {
  font-weight: 501;
}
.col-sm-2 {
  padding-left: 0;
}

.choose {
  padding: 10px;
  border-style: solid;
  border-color: #ddd;
  border-width: 1px;
  border-radius: 4px 4px 0 0;
}

.col-sm-6 {
  padding: 10px;
  border-style: solid;
  border-color: #ddd;
  border-width: 1px;
  border-radius: 4px 4px 0 0;
}

.startip {
  margin-top: 10px;
  padding-right: 0px;
}

.endip {
  margin-top: 10px;
  padding-right: 0px;
  padding-left: 0px;
}

.col-sm-4 label {
  float: right;
}
select {
  font-family: '微软雅黑';
  border: 1px #ccc solid;
  border-radius: 5px;
}

.details-content .article-cont p {
  padding: 30px 0 0 5px;
}

label {
  font-weight: 400;
  margin-top: 5px;
}

h5 {
  margin-top: 29px;
  margin-bottom: 0;
}
</style>
