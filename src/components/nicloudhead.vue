<template>
<div class="abc">
	<ul class="breadcrumb">
    	<li>
			<router-link :to="{name:'nicloud'}">
				<span>NICLOUD</span>
			</router-link>
		</li>
    	<li>
			<router-link :to="{name:'serveroom'}">
				<span>SERVEROOM</span>
			</router-link>
		</li>
	</ul>
</div>

</template>

<style scoped>
a {
color : #FFF;
}

.breadcrumb {
	color: #FFF;
	background-color: #5B5B5B;
}
.abc {
	background-color: #e3e3e3;
}
</style>
