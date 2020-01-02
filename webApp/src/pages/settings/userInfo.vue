<template>
    <div>
        <van-nav-bar
                title="用户信息"
                left-text="返回"
                left-arrow
                @click-left="goBack">
            <van-icon name="edit" slot="right" v-on:click="showPasswordInput = true"/>
        </van-nav-bar>
        <van-field v-model="this.user.name" disabled label="用户ID ："/>
        <van-field value="********" disabled label="密 码 ："/>
        <van-field
                v-model="this.user.alipayAccount" center disabled label="支付宝：">
        </van-field>
        <van-popup
                v-model="showPasswordInput"
                position="bottom"
                :style="{ height: '50%' }">
            <van-cell title="修改密码或支付宝" :border="false"/>
            <van-field required v-model="updateInfo.oldPassword" type="password" label="旧密码 ："/>
            <van-field v-model="updateInfo.newPassword" type="password" label="新密码 ："/>
            <van-field v-model="updateInfo.checkPassword" type="password" label="确认密码 ："/>
            <van-field v-model="updateInfo.newAlipayAccount" label="新支付宝："/>
            <van-button type="primary" round v-on:click="update">确认修改</van-button>
        </van-popup>
    </div>
</template>
<script>
    import {mapActions, mapState} from 'vuex'
    import {USERS} from '../../js/const/const'

    export default {
        data() {
            return {
                checkPassword: "",
                showPasswordInput: false,
                alipayAccount: "",
                updateInfo: {
                    oldPassword: "",
                    newPassword: "",
                    checkPassword: "",
                    newAlipayAccount: ""
                }
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        methods: {
            ...mapActions({
                loginFunc: 'loginFunc'
            }),
            changeAlipayAccount: function () {
                if (this.checkPassword.length < 6) {
                    this.$toast("密码错误!")
                } else {
                    this.$toast("修改成功")
                }
                this.showPasswordInput = false;
            },
            goBack: function () {
                this.$router.go(-1)
            },
            update: function () {
                let that = this;
                if (that.updateInfo.oldPassword.length === 0 || that.updateInfo.newPassword !== that.updateInfo.checkPassword) {
                    return that.$toast("密码为空或两次密码不一致!");
                }
                let url = USERS + that.user.name + "/" + that.updateInfo.oldPassword;
                that.$axios.get(url).then(function (resp) {
                    let newUserInfo = {
                        name: that.user.name,
                        password: that.updateInfo.newPassword,
                        alipayAccount: that.updateInfo.newAlipayAccount,
                    };
                    let putUrl = USERS + that.user.name;
                    console.log(that.user.registered);
                    console.log(!that.user.registered);
                    if (!that.user.registered) {
                        putUrl = putUrl + "?register=true"
                    }
                    that.$axios.put(putUrl,
                        JSON.stringify(newUserInfo),
                        {headers: {'Content-Type': 'application/json'}}
                        ).then(function (resp) {
                        that.loginFunc(resp.data);
                        that.showPasswordInput = false;
                        that.updateInfo = {
                            oldPassword: "",
                            newPassword: "",
                            checkPassword: "",
                            newAlipayAccount: ""
                        }
                    }).catch(function (err) {
                        that.$toast(err)
                    })
                }).catch(function (err) {
                    that.$toast(err)
                })
            }
        }
    }
</script>