import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
	token: "",
	vm: {
		uuid: "",
		host: "",
		cpu: "",
		mem: "",
		owner: "",
		os: "",
		comment: "",
		ip: ""
		},
	network: {
		vlan: "",
		bridge: "",
		network: "",
		prefix: "",
		gateway: "",
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
	set_token(state, token) {
		state.token = token
		sessionStorage.token = token
  		}
	}
})

export default store
