<template>
<input type="text" value="2012-05-15 21:05" id="datetimepicker">
</template>


<script>
import "../../static/js/bootstrap-datetimepicker.min.js"

export default {

mounted () {
		$('#datetimepicker').datetimepicker({
    		format: 'yyyy-mm-dd hh:ii'
		});
	}
}
</script>
