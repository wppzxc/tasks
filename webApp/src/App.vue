<template>
    <div id="app">
        <router-view></router-view>
        <br/>
        <van-tabs :border="false"/>
        <van-tabbar v-model="active" route>
            <van-tabbar-item icon="apps-o" name="tasks" to="/taskInfo">首页</van-tabbar-item>
            <van-tabbar-item icon="todo-list-o" name="commits" to="/commits">历史</van-tabbar-item>
            <van-tabbar-item icon="contact" name="settings" to="/settings">我的</van-tabbar-item>
        </van-tabbar>
    </div>
</template>

<script>
    import {mapActions} from 'vuex';
    import {USERS} from './js/const/const';

    export default {
        data() {
            return {
                active: "tasks"
            }
        },
        mounted() {
            // if user exist ?
            if (this.$route.name !== "register") {
                this.autoLogin();
            }
        },
        methods: {
            ...mapActions({
                loginFunc: 'loginFunc'
            }),
            autoLogin: function () {
                let that = this;
                let username = localStorage.getItem("username");
                if (username === null) {
                    return that.generateUser("test");
                }
                that.$axios.get(USERS + username).then((resp)=>{
                    that.loginFunc(resp.data);
                }).catch((err)=>{
                    console.log(err);
                    console.log(err)
                })
            },
            generateUser: function (parent) {
                let that = this;
                let username = new Date().getTime();
                let url = USERS + username;
                if (parent !== null && parent.length > 0 ) {
                    url  = url + "?parent=" + parent
                }
                that.$axios.post(url).then((resp)=>{
                    console.log(resp);
                    let user = resp.data;
                    that.loginFunc(user);
                    localStorage.setItem("username", user.name);
                }).catch((err)=>{
                    console.log(err);
                    that.$toast("注册失败：" + err)
                })
            }
        }
    }
</script>

<style>
    #app {
        font-family: Helvetica, sans-serif;
        text-align: center;
    }
</style>
