<template>
  <div class="col-sm-12" style="margin-top: 10px; padding-right:0; padding-left:0">
    <button class="btn btn-success btn-sm" @click="createhost" type="button" style="margin-bottom: 12px; margin-right: 20px; float: right">
      创建数据中心<span class="glyphicon glyphicon-plus" style="margin-left: 5px"></span>
    </button>
    <table class="table table-hover" style="text-align: center;">
      <thead>
        <tr>
          <th>
            <label class="checkbox-inline" style="border:red 1px">
              <input type="checkbox" v-model="checkvalue" @click="checkbox()" />
            </label>
          </th>
          <th>数据中心</th>
          <th>备注</th>
          <th>状态</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody v-for="(item, index) in data" :key="item.id">
        <tr v-if="item.Status" class="table-dark text-dark" :id="item.Uuid">
          <label class="checkbox-inline" style="width:10px">
            <input type="checkbox" v-model="item.Checkout" />
          </label>
          <td>{{ item.Datacenter }}</td>
          <td>{{ item.Comment }}</td>
          <td>
            <span class="glyphicon glyphicon-ok"></span>
          </td>

          <td>
            <button class="btn btn-danger btn-xs" type="button" @click="deletedatacenter(item.Datacenter, index)">
              <span class="glyphicon glyphicon-trash"></span>
              删除
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
export default {
  data() {
    return {
      data: []
    }
  },

  mounted: function() {
    this.getdatacenter()
  },

  methods: {
    createhost: function() {
      this.$emit('toParent', 'createdatacenter')
    },

    deletedatacenter: function(d, index) {
      var apiurl = `/api/datacenter/deldatacenter`
      this.$http.get(apiurl, { params: { datacenter: d } }).then(response => {
        if (response.data.err === null) {
          alert('删除成功')
          this.data[index].Status = 0
        } else {
          alert(response.data.err.Message)
        }
      })
    },

    getdatacenter: function(ip) {
      var apiurl = `/api/datacenter/getdatacenter`
      this.$http.get(apiurl).then(response => {
        if (response.data.err === null) {
          this.data = response.data.res
        } else {
          alert(response.data.err.Message)
        }
      })
    }
  }
}
</script>
<style scoped>
.checkbox-inline {
  margin-bottom: 30px;
}

.table thead tr th {
  border-bottom: 2px solid #846d6d;
}

.details-content .article-cont p {
  padding: 30px 0 0 5px;
}

label {
  font-weight: 400;
}

.table tbody tr td {
  vertical-align: 'middle';
}

th {
  background-color: #e8d18d;
  font-weight: bold;
  color: black;
  text-align: center;
}
</style>
