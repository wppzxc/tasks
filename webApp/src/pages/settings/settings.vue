<template>
    <div>
        <van-nav-bar title="系统设置"/>
        <van-cell icon="user-circle-o"
                  :title="this.user.name"
                  value=" " size="large"
                  is-link to="/userInfo"/>
        <van-cell icon="flag-o"
                  title="我的邀请码"
                  value=" " size="large"
                  is-link to="/inviteInfo"/>
        <van-field center disabled label="累积收益" left-icon="balance-o" v-model="this.user.balance">
            <van-button slot="button" size="small" type="danger" @click="showPasswordInput=true">提现</van-button>
        </van-field>
        <van-cell icon="user-o" title="一级用户" :value="this.user.subUsers" size="large"/>
        <van-cell icon="friends-o" title="二级用户" :value="this.user.subTwoUsers" size="large"/>
        <van-cell icon="service-o"
                  title="客服"
                  value=" " size="large"
                  is-link to="/service"/>
        <br/>
        <van-row>
            <van-col offset="2" span="20">
                <van-button type="danger" round size="large" v-on:click="onLogout">切换用户</van-button>
            </van-col>
        </van-row>
        <van-popup
                v-model="showPasswordInput"
                position="bottom"
                :style="{ height: '60%' }">
            <van-password-input
                    :value="checkPassword"
                    info="请输入密码"
                    :focused="showPasswordInput"
                    @focus="showPasswordInput = true"/>
            <van-number-keyboard
                    :show="showPasswordInput"
                    extra-key="."
                    close-button-text="确认"
                    @input="onInput"
                    @delete="onDelete"
                    @close="withdrawals"
                    @blur="showPasswordInput = false"/>
        </van-popup>
    </div>
</template>

<script>
    import {mapActions, mapState} from 'vuex'
    import {INVITE_URL} from '../../js/const/const'
    export default {
        data() {
            return {
                showPasswordInput: false,
                checkPassword: "",
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        methods: {
            ...mapActions({
                logoutFunc: 'logoutFunc'
            }),
            copy: function () {
                let that = this;
                let url = INVITE_URL + that.user.name;
                that.$copyText(url).then(function (e) {
                    that.$toast("已复制邀请链接，去邀请吧");
                }, function (e) {
                    that.$toast("无法获取邀请链接！");
                })
            },
            onLogout: function () {
                sessionStorage.clear();
                localStorage.clear();
                this.logoutFunc();
                this.$toast("退出成功！");
                this.$router.push("/login")
            },
            onInput: function(key) {
                this.checkPassword = (this.checkPassword + key).slice(0, 6);
            },
            onDelete: function() {
                this.checkPassword = this.checkPassword.slice(0, this.checkPassword.length - 1);
            },
            withdrawals:function () {
                if (this.checkPassword.length < 6) {
                    this.$toast("密码错误!")
                } else {
                    this.$toast("提现成功")
                }
                this.showPasswordInput = false;
            },
        }
    }
</script>