<template>
  <div style="color:#FFA500; margin-top: 2px">
    <div style="margin-left: 201px">
      <strong @click="vmlist"> <span class="glyphicon glyphicon-tasks"></span><span> Dashboard Overview </span> </strong>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      username: '',
      charge: '',
      active: ''
    }
  },
  created: function() {
    this.getuser()
  },
  methods: {
    serveroom: function() {
      sessionStorage.setItem('router', 'server')

      this.$router.push('serveroom')
    },

    getuser: function() {
      var u = this.$store.state.username
      if (u === null || typeof u === 'undefined' || u === '' || u === 'undefined') {
        this.username = sessionStorage.getItem('username')
      } else {
        sessionStorage.setItem('username', this.$store.state.username)
        this.username = this.$store.state.username
      }
    },

    leave: function() {
      this.active = 'color: #FFF'
    },

    vmlist: function() {
      this.$emit('toParent', 'vm')
    },

    mouseOver: function() {
      this.active = 'color: #3090C7'
    }
  }
}
</script>

<style scoped>
a {
  color: #ffa500;
}

.glyphicon {
  top: 2px;
}
</style>
