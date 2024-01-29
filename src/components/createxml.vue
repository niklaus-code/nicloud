<template>
  <div>
    <div class="col-sm-12 form-group" style="border-bottom: 1px green solid; margin-top:30px">
      <h4>创建镜像配置文件</h4>
    </div>

    <div class="col-sm-8 col-sm-offset-1">
      <div class="col-sm-12" style="margin-top:20px">
        <div class="form-group">
          <div class="col-sm-3 col-sm-offset-1">
            <label>名称</label>
          </div>
          <div class="col-sm-3">
            <form role="form">
              <div class="form-group">
                <input type="text" class="form-control" v-model="comment" placeholder="" />
              </div>
            </form>
          </div>
        </div>
      </div>
      <div class="col-sm-12">
        <div class="form-group">
          <div class="col-sm-3 col-sm-offset-1">
            <label>配置文件标签</label>
          </div>
          <div class="col-sm-3">
            <select class="col-sm-12" v-model="tagvalue">
              <option value="">--请选择--</option>
              <option v-for="s in ostags" :key="s.id" :value="s">
                {{ s.Tag }}
              </option>
            </select>
          </div>
        </div>
      </div>
      <div class="col-sm-12" style="margin-top: 10px">
        <div class="form-group">
          <div class="col-sm-3 col-sm-offset-1">
            <label>镜像xml</label>
          </div>
          <div class="col-sm-8">
            <form role="form">
              <div class="form-group">
                <textarea class="form-control" v-model="xml" rows="16"></textarea>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div class="col-sm-12" style=" margin-bottom:20px">
        <div class="form-group" style="margin-top:20px;">
          <div class="col-sm-2 col-sm-offset-2"></div>
          <div class="col-sm-8" style="margin:0 auto; text-align: center;">
            <button type="submit" style="margin:0 auto" @click="createosimage" class="btn btn-success btn-sm">提交</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      ostags: [],
      tagvalue: '',

      xml: '',
      comment: ''
    }
  },

  created: function() {
    this.getostag()
  },

  methods: {
    getostag: function() {
      var apiurl = `/api/osimage/getiostags`

      this.$http.get(apiurl).then(response => {
        if (response.data.err === null) {
          this.ostags = response.data.res
        } else {
          alert(response.data.err.Message)
        }
      })
    },

    createosimage: function() {
      var apiurl = `/api/osimage/createxml`

      this.$http.post(apiurl, this.$qs.stringify({ xml: this.xml, comment: this.comment, tag: this.tagvalue.Id })).then(response => {
        if (response.data.err === null) {
          alert('创建成功!')
          this.$emit('toParent', 'xml')
        } else {
          alert(response.data.err.Message)
        }
      })
    }
  }
}
</script>
<style scoped>
select {
  height: 30px;
  font-family: '微软雅黑';
  border: 1px #ccc solid;
  border-radius: 5px;
}

label {
  float: right;
  font-weight: 400;
  margin-top: 5px;
}
</style>
