<template>
  <div class="col-sm-12" style="margin-top: 10px; padding-right:0; padding-left:0">
    <div class="col-md-5" style="float: right">
      <button class="btn btn-success btn-sm" type="button" @click="toParent" style="float: right; margin-bottom: 12px; margin-right: 5px">
        创建镜像配置文件xml<span class="glyphicon glyphicon-plus" style="margin-left: 5px"></span>
      </button>
    </div>
    <table class="table table-hover" style="text-align: center;">
      <thead>
        <tr>
          <th>ID</th>
          <th>名称</th>
          <th>配置文件类型</th>
          <th width="45%">XML</th>
          <th>操作</th>
        </tr>
      </thead>

      <tbody v-for="(item, index) in xmllist" :key="item.id">
        <tr class="table-dark text-dark" :id="item.Uuid">
          <td>{{ item.Id }}</td>
          <td>{{ item.Comment }}</td>
          <td>{{ item.Sort.Tag }}</td>
          <td class="checkxml">
            <textarea v-if="item.checkxml" rows="20" width="30" class="form-control" v-model="item.Xml" />
            <p v-else>{{ item.Xml }}</p>
          </td>
          <td style="min-width: 125px">
            <button class="btn btn-success btn-xs" type="button" @click="xmlinfo(index)">
              <span class="glyphicon glyphicon-info-sign"></span>
              查看详情
            </button>
            <button class="btn btn-danger btn-xs" type="button" @click="delxml(item.Id, index)">
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
      xmllist: [],
      data: [],
      osimagesort: []
    }
  },

  mounted: function() {
    this.getosxml()
  },

  methods: {
    xmlinfo: function(index) {
      if (this.xmllist[index].checkxml === false) {
        this.xmllist[index].checkxml = true
      } else {
        this.xmllist[index].checkxml = false
      }
    },

    getosxml: function() {
      var apiurl = `/api/osimage/getosimagexml`

      this.$http.get(apiurl).then(response => {
        if (response.data.err === null) {
          for (var i in response.data.res) {
            response.data.res[i]['checkxml'] = false
            this.xmllist.push(response.data.res[i])
          }
        } else {
          alert(response.data.err.Message)
        }
      })
    },

    toParent: function(item) {
      this.$emit('toParent', 'createxml')
    },

    delxml: function(xmlid, index) {
      var apiurl = `/api/osimage/delxml`
      this.$http.get(apiurl, { params: { id: xmlid } }).then(response => {
        if (response.data.err === null) {
          alert('删除成功')
          this.getosxml()
        } else {
          alert(response.data.err.Message)
        }
      })
    }
  }
}
</script>
<style scoped>
table {
  table-layout: fixed;
}

.table thead tr th {
  border-bottom: 2px solid #846d6d;
}

select {
  font-family: '微软雅黑';
  border: 1px #1a1a1a solid;
  border-radius: 5px;
}

.checkbox-inline {
  margin-bottom: 30px;
}

.details-content .article-cont p {
  padding: 30px 0 0 5px;
}

input {
  margin-top: 2px;
}

label {
  font-weight: 400;
}

.table tbody tr td {
  vertical-align: 'middle';
}

.tdxml {
  max-width: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

th {
  background-color: #e8d18d;
  font-weight: bold;
  color: black;
  text-align: center;
}

.breadcrumb {
  background-color: #fff;
  margin-bottom: 0;
  padding-left: 10px;
}

.checkxml {
  word-break: keep-all;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
