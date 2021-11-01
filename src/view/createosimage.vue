<template>
	<div>
	<nicloudhead></nicloudhead>
	<vmleft></vmleft>
  	<div class="content whisper-content leacots-content details-content col-md-11 col-md-offset-2" style="background-color:white; float:left">
		<div class="col-sm-8 col-sm-offset-1" style="margin-top:20px">
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>镜像名称</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="osimage" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>ceph块</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="cephblockdevice" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>快照名称</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<input type="text" class="form-control" v-model="snapimage" placeholder="">
  						</div>
					</form>
				</div>
    		</div>
	 		<div class="form-group">
				<div class="col-sm-4">
        			<label>xml</label>
				</div>
				<div class="col-sm-8">
					<form role="form">
  						<div class="form-group">
    						<textarea class="form-control" v-model="xml" rows="16"></textarea>
  						</div>
					</form>
				</div>
    		</div>
		<div class="form-group" style="margin-top:20px" >
			<div class="col-sm-2 col-sm-offset-4">
  				<button type="submit" @click="createosimage" class="btn btn-info">提交</button>
			</div>
		</div>
		</div>
	</div>		
	</div>		
</template>
<script>
import foot from '@/components/footer'
import nicloudhead from '@/components/nicloudhead'
import vmleft from '@/components/vmleft'


export default {
    data () {
        return {
			osimage: "",
			cephblockdevice: "",
			snapimage: "",
			xml: "",
        }
    },

    components: {
        foot, nicloudhead, vmleft
    },

    created: function () {
        this.vlaninfo()
    },

    methods: {
		vlaninfo: function () {
            this.osimage = this.$route.query.osimage
            this.cephblockdevice = this.$route.query.cephblockdevice
            this.snapimage = this.$route.query.snapimage
            this.xml = this.$route.query.xml
            },

		check: function (osname, cephblockdevice, snapimage, xml) {
			if (typeof osname === 'undefined' || osname === null || osname === ''|| typeof cephblockdevice === 'undefined' || cephblockdevice === null || cephblockdevice === '' || typeof snapimage === 'undefined' || snapimage === null || snapimage === '' ||typeof xml === 'undefined' || xml === null || xml === '') {
				alert("缺少信息")
                return true
            } else {
				return false
				}
			},

		createosimage: function () {
			if (this.check(this.osimage, this.cephblockdevice, this.snapimage, this.xml)) {
				return 
				}
		

            var apiurl = `/api/osimage/createimage`

            this.$http.get(apiurl, { params: { osname: this.osimage, cephblockdevice: this.cephblockdevice, snapimage: this.snapimage, xml: this.xml} }).then(response => {
				if (response.data.res === null) {
					alert("创建成功!")
					this.$router.push('/osimage')
				} else {
					alert("创建失败(" + response.data.res.Message+ ")" )
					}
			})
			},

        }
  }
</script>
<style scoped>

.col-sm-4 label{
	float: right;
}
select{
    font-family: "微软雅黑";
    border: 1px #1a1a1a solid;
    border-radius: 5px;
}

.content {
    box-shadow: 0 0 10px rgba(0,0,0,8);
    border-radius: 10px/10px;
    z-index: -1;
    padding: 100px 0px 100px 0px;
    margin-left: 0px;
    margin-TOP: 50px;
}

.details-content .article-cont p {
    padding:30px 0 0 5px
}

label {
	font-weight : 400;
	margin-top: 5px;
}
</style>
