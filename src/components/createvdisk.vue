<template>
<div>
      <div class="col-sm-12 form-group" style="border-bottom: 1px green solid;margin-top:30px">
                <h4>创建云盘</h4>
            </div>
		<div class="col-sm-8 col-sm-offset-1" style="margin-top:30px;margin-bottom:30px">
			<div class="col-sm-12">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>数据中心*</label>
				</div>
				<div class="col-sm-8">
				    <select class="col-sm-12" v-model="centervalue" @change="getstorage(centervalue)">
						 <option value="">--请选择--</option>
                        <option  v-for="c in datacenter" :value="c.Datacenter">
                            {{ c.Datacenter }}
                        </option>
                    </select>
				</div>
				</div>
    		</div>
				<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>存储集群*</label>
				</div>
				<div class="col-sm-8">
				    <select class="col-sm-12" v-model="storagevalue" @change="getpool">
						 <option value="">--请选择--</option>
                        <option  v-for="c in storage" :value="c.Uuid">
                            {{ c.Name }} (size: {{c.Remainder}} / {{c.Contain}} GB)
                        </option>
                    </select>
				</div>
				</div>
    		</div>
				<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>存储池*</label>
				</div>
				<div class="col-sm-8">
				    <select class="col-sm-12" v-model="poolvalue">
						 <option value="">--请选择--</option>
                        <option  v-for="p in pool" :value="p.Pool">
                            {{ p.Pool }}
                        </option>
                    </select>
				</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>容量*</label>
				</div>
				<div class="col-sm-8">
				    <select class="col-sm-12" v-model="containvalue">
						 <option value="">--请选择--</option>
                        <option  v-for="c in contain" :value="c">
                            {{ c }}G
                        </option>
                    </select>
				</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-2 col-sm-offset-2">
        			<label>备注</label>
				</div>
				<div class="col-sm-8">
                    <form role="form">
                        <div class="form-group">
                            <input type="text" class="form-control" v-model="comment" placeholder="">
                        </div>
                    </form>
				</div>
				</div>
    		</div>
			<div class="col-sm-12" style="margin-top:20px">
				<div class="form-group">
					<div class="col-sm-2 col-sm-offset-4">
  						<button type="submit" @click="createvdisk" class="btn btn-success">提交</button>
					</div>
				</div>
			</div>
	</div>
</div>
</template>
<script>

export default {
    data () {
        return {
            comment: "",

           	centervalue: "",
            datacenter: [],

			storage : [],
            storagevalue: "",

			poolvalue: "",
			pool: [],
			
			containvalue: "", 
			contain: [
				100, 200, 500,1024 
				], 
        }
    },

    created: function () {
		this.getdatacenter()
    },

    methods: {
		getpool: function () {
            var apiurl = `/api/storage/getpool`

            this.$http.get(apiurl, {params: { datacenter: this.centervalue, storage: this.storagevalue}}).then(response => {
                if (response.data.err === null) {
                    this.pool = response.data.res
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
            },

        getdatacenter: function () {
            var apiurl = `/api/datacenter/getdatacenter`

            this.$http.get(apiurl).then(response => {
                if (response.data.err === null) {
                    this.datacenter = response.data.res
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
            },

       getstorage: function (centervalue) {
            var apiurl = `/api/storage/get`
            this.$http.get(apiurl, { params: { datacenter: centervalue}}).then(response => {
                if (response.data.err === null) {
                    this.storage = response.data.res
                } else {
                    alert("获取数据失败(" + response.data.err.Message+ ")" )
                    }
            })
        },


		check: function (centervalue, storagevalue, poolvalue, containvalue) {
			if (typeof  centervalue=== 'undefined' ||  centervalue=== null ||  centervalue=== ''|| typeof  storagevalue=== 'undefined' || storagevalue === null || storagevalue === '' || typeof poolvalue=== 'undefined' || poolvalue=== null || poolvalue=== '' ||typeof containvalue === 'undefined' || containvalue === null || containvalue === '') {
				alert("缺少信息")
                return true
            } else {
				return false
				}
			},

		createvdisk: function () {
			if (this.check(this.centervalue, this.storagevalue, this.poolvalue, this.containvalue)) {
				return 
				}

            var apiurl = `/api/vdisk/createvdisk`

            this.$http.post(apiurl,  this.$qs.stringify({ datacenter: this.centervalue, storage: this.storagevalue, pool: this.poolvalue, contain: this.containvalue, comment: this.comment})).then(response => {
				if (response.data.err === null) {
					alert("创建成功!")
					this.$emit("toParent", "disk");
				} else {
					alert("创建失败(" + response.data.err.Message+ ")" )
					}
			})
			},

        }
  }
</script>
<style scoped>

label {
	float: right;
}

select{
	height:30px;
    font-family: "微软雅黑";
    border: 1px #ccc solid;
    border-radius: 5px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>
