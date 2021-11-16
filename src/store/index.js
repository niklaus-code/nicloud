import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
	vm: {
		uuid: "",
		host: ""
		},
	network: {
		vlan: "",
		},
	vdisk: {
    	vdiskid: "",
    	storage: "",
    	pool: "",
		},
	osimage: {
		id: 0,
		osname: "",
		cephblockdevice: "",
		snap: "",
		xml: "",
		},
  },
  mutations: {
    increment (state) {
      state.count++
    }
  }
})

export default store
