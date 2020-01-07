<template>
    <div>
        <van-nav-bar
                title="请登录"
                left-text="返回"
                left-arrow
                @click-left="goBack"/>
        <van-cell-group>
            <van-field
                    v-model="username"
                    required
                    clearable
                    left-icon="contact"
                    label="用户名"
                    placeholder="请输入用户名"
            />

            <van-field
                    v-model="password"
                    type="password"
                    left-icon="closed-eye"
                    label="密码"
                    placeholder="请输入密码"
                    required
            />
        </van-cell-group>
        <van-cell-group title=" ">
            <van-button type="primary" round size="large" v-on:click="onLogin">登录</van-button>
            <br/><br/>
        </van-cell-group>
    </div>
</template>

<script>
    import {mapActions} from 'vuex';
    import {USERS} from "../../js/const/const";

    export default {
        data() {
            return {
                username: "",
                password: "",
            }
        },
        mounted() {
            this.$notify({
                message: "请登录！",
                type: "warning",
            })
        },
        methods: {
            onLogin: function () {
                let that = this;
                let url = USERS + that.username + "/" + that.password;
                that.$axios.get(url).then((resp)=>{
                    that.loginFunc(resp.data);
                    localStorage.setItem("username", that.username);
                    that.$notify({
                        message: "登录成功！",
                        type: "success"
                    });
                    that.$router.go(-1)
                }).catch((err)=>{
                    console.log(err)
                });
            },
            ...mapActions({
                loginFunc: 'loginFunc'
            }),
            goBack: function () {
                this.$router.go(-2)
            }
        }
    }
</script>