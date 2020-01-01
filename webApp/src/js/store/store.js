import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        user: {
            name: "",
            alipayAccount: "",
            status: "",
            registered: false,
            parent: "",
            subUsers: 0,
            subTwoUsers: 0,
            balance: 0
        },
    },
    getters: {
        getUser: function (state) {
            return state.user
        },
    },
    mutations: {
        setUser: function (state, user) {
            state.user = user
        },
    },
    actions: {
        loginFunc(context, user) {
            context.commit("setUser", user);
        },
        logoutFunc(context) {
            context.commit("setUser", {
                name: "",
                alipayAccount: "",
                status: "",
                registered: false,
                parent: "",
                subUsers: 0,
                subTwoUsers: 0,
                balance: 0
            });
        }
    }
});

export default store;