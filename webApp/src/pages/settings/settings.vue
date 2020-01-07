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
                  is-link @click="showInvite"/>
        <van-field center disabled label="累积收益" left-icon="balance-o" :value="this.user.balance | transMoney">
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
        <van-popup v-model="showQrCode">
<!--            width="360"-->
<!--            height="640"-->
            <van-image
                    width="1000"
                    height="1750"
                    :src="final_img"
                    v-on:click="showQrCode = false"/>
        </van-popup>
        <img :src="final_img" class="result-img" v-show="false">
        <div class="result-img" v-show="false">
            <canvas id="my_canvas" width="1100" height="1750"></canvas>
        </div>
    </div>
</template>

<script>
    import {mapActions, mapState} from 'vuex'
    import {INVITE_URL} from '../../js/const/const'
    import  QRCode from 'qrcode'
    export default {
        data() {
            return {
                showPasswordInput: false,
                genQrCode: true,
                showQrCode: false,
                checkPassword: "",
                inviteURL: "",
                final_img: "",
            }
        },
        computed: {
            ...mapState({
                user: state => state.user
            })
        },
        mounted() {
            let that = this;
            QRCode.toDataURL(INVITE_URL + that.user.name).then(function (url) {
                that.inviteURL = INVITE_URL + that.user.name;
                let canvas = document.getElementById('my_canvas');
                let ctx = canvas.getContext('2d');
                let img1 = new Image();
                let img2 = new Image();
                // 处理跨域
                img1.crossOrigin = 'anonymous';
                img2.crossOrigin = 'anonymous';
                img1.src = require('../../assets/img/invite.jpg'); // 背景图路经
                img2.src = url; // 生成的二维码base64
                img1.onload = function () {
                    ctx.drawImage(img1, -70, 0, 1250, 1750); // 背景图载入画板
                    ctx.drawImage(img2, 330, 1210, 440, 420);
                    that.final_img = canvas.toDataURL('image/jpeg', 0.5)
                }
            });
        },
        methods: {
            ...mapActions({
                logoutFunc: 'logoutFunc'
            }),
            showInvite: function () {
                let that = this;
                that.showQrCode = true;
                if (that.genQrCode) {
                    setTimeout(function () {
                        that.inviteURL = INVITE_URL + that.user.name;
                        that.genQrCode = false;
                    },100);
                }
            },
            onLogout: function () {
                sessionStorage.clear();
                localStorage.clear();
                this.logoutFunc();
                this.$toast("退出成功！");
                this.$router.push("/login")
            },
            onInput: function (key) {
                this.checkPassword = (this.checkPassword + key).slice(0, 6);
            },
            onDelete: function () {
                this.checkPassword = this.checkPassword.slice(0, this.checkPassword.length - 1);
            },
            withdrawals: function () {
                if (this.checkPassword.length < 6) {
                    this.$toast("密码错误!")
                } else {
                    this.$toast("提现成功")
                }
                this.showPasswordInput = false;
            },
        },
        filters: {
            transMoney: function (money) {
                return (money/100).toFixed(2)
            }
        },
    }
</script>
<style>
    #qrcode > img {
        z-index: 3003;
        position: absolute;
        height: 25rem;
        margin-top: -50%;
        margin-left: 30%;
    }
</style>