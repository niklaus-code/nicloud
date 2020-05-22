<template>
    <div id="app">
	<headd></headd>
	<div class="row">
        <VueDragResize :isActive="false" :x="600" y="200" :w="200" :h="200" style="background-color: rgb(174, 213, 129);">
            <h3>可以移动哦</h3>
        </VueDragResize>
        <VueDragResize :isActive="false" :x="200" y="200":w="200" :h="200" style="background-color: rgb(129, 212, 250);">
            <h3>可以变大哦</h3>
        </VueDragResize>
        <VueDragResize :isActive="false" :x="400" :y="400" :w="200" :h="200" style="background-color: rgb(239, 154, 154);">
            <h3>可以缩小哦</h3>
        </VueDragResize>
	</div>
    </div>
</template>

<script>
    import foot from '@/components/footer'
    import headd from '@/components/head'
    import VueDragResize from 'vue-drag-resize';

    export default {
        name: 'app',

        components: {
            VueDragResize, foot, headd
        },

        data() {
            return {
                width: 100,
                height: 200,
                top: 100,
                left: 100
            }
        },

        methods: {
            resize(newRect) {
                this.width = newRect.width;
                this.height = newRect.height;
                this.top = newRect.top;
                this.left = newRect.left;
            },
	    blog: function (id, sig) {
      		this.$http.post('/api/xianyu/get_user/get_blog_by_id', {id: id, category: this.category}).then(response => {
      		})
    	    }
        }
    }
</script>
